'use strict';

// 使用 https://www.npmjs.com/package/lei-run 命令运行

const pkg = 'github.com/leizongmin/gogo';

const pkgParent = path.dirname(pkg);
const workspace = path.resolve(pwd, '_workspace');
const out = path.resolve(pwd, 'bin', path.basename(pkg));

function mkdir(dir) {
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir);
  }
}

function autoInitWorkspace() {
  if (!fs.existsSync(workspace)) run('workspace');
}

register('workspace', 'create virtual workspace', function () {
  env.GOPATH = workspace;
  mexec([
    `mkdir -p ${workspace}/src/${pkgParent}`,
    `ln -s ${pwd} ${workspace}/src/${pkg}`,
    `mkdir -p ${pwd}/vendor`,
    `mkdir -p ${workspace}/vendor`,
    `ln -s ${pwd}/vendor ${workspace}/vendor/src`,
    `mkdir -p -p ${workspace}/vendor/pkg`,
    `ln -s ${workspace}/vendor/pkg ${workspace}/pkg`,
  ]);
});

register('clean', 'clean virtual workspace and vendor', function () {
  mexec([
    `rm -rf ${workspace}`,
    `rm -rf ${pwd}/vendor`,
    `rm -rf ${pwd}/bin`,
  ]);
});

register('build', 'build project', function () {
  autoInitWorkspace();
  env.GOPATH = workspace;
  exec(`go build -o ${out} ${pkg}`);
});

register('bin', 'run program', function () {
   exec(`${out}`);
});

register('vendor', 'add dependencies', function () {
  autoInitWorkspace();
  if (argv.length < 1) return exit(1, 'Usage: run vendor packages');
  const gopath = path.resolve(workspace, 'vendor');
  env.GOPATH = gopath;
  exec(`go get ${argv.join(' ')}`);
});

register('go', 'run go command in GOPATH', function () {
  autoInitWorkspace();
  env.GOPATH = workspace;
  exec(`go ${argv.join(' ')}`);
});
