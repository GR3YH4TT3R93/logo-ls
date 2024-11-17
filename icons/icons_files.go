package icons

var IconFileName = map[string]*IconInfo{
	".babelrc":                            IconSet["babel"],
	".babelrc.js":                         IconSet["babel"],
	".babelrc.json":                       IconSet["babel"],
	".bash_history":                       IconSet["console"],
	".bowerrc":                            IconSet["bower"],
	".buildignore":                        IconSet["settings"],
	".clang-format":                       IconSet["settings"],
	".clang-tidy":                         IconSet["settings"],
	".codeclimate.yml":                    IconSet["code-climate"],
	".codecov.yml":                        IconSet["codecov"],
	".commitlint.yaml":                    IconSet["commitlint"],
	".commitlint.yml":                     IconSet["commitlint"],
	".commitlintrc":                       IconSet["commitlint"],
	".commitlintrc.js":                    IconSet["commitlint"],
	".commitlintrc.json":                  IconSet["commitlint"],
	".editorconfig":                       IconSet["editorconfig"],
	".env.defaults":                       IconSet["tune"],
	".env.dev":                            IconSet["tune"],
	".env.development":                    IconSet["tune"],
	".env.development.local":              IconSet["tune"],
	".env.example":                        IconSet["tune"],
	".env.local":                          IconSet["tune"],
	".env.preview":                        IconSet["tune"],
	".env.prod":                           IconSet["tune"],
	".env.production":                     IconSet["tune"],
	".env.production.local":               IconSet["tune"],
	".env.qa":                             IconSet["tune"],
	".env.qa.local":                       IconSet["tune"],
	".env.sample":                         IconSet["tune"],
	".env.schema":                         IconSet["tune"],
	".env.staging":                        IconSet["tune"],
	".env.staging.local":                  IconSet["tune"],
	".env.test":                           IconSet["tune"],
	".env.test.local":                     IconSet["tune"],
	".env.testing":                        IconSet["tune"],
	".esformatter":                        IconSet["json"],
	".eslintcache":                        IconSet["eslint"],
	".eslintignore":                       IconSet["eslint"],
	".eslintrc":                           IconSet["eslint"],
	".eslintrc.cjs":                       IconSet["eslint"],
	".eslintrc.js":                        IconSet["eslint"],
	".eslintrc.json":                      IconSet["eslint"],
	".eslintrc.yaml":                      IconSet["eslint"],
	".eslintrc.yml":                       IconSet["eslint"],
	".esmrc":                              IconSet["nodejs"],
	".firebaserc":                         IconSet["firebase"],
	".gcloudignore":                       IconSet["gcp"],
	".gitattributes":                      IconSet["git"],
	".gitconfig":                          IconSet["git"],
	".gitignore":                          IconSet["git"],
	".gitkeep":                            IconSet["git"],
	".gitmodules":                         IconSet["git"],
	".graphqlconfig":                      IconSet["graphql"],
	".helmignore":                         IconSet["helm"],
	".htaccess":                           IconSet["xml"],
	".htpasswd":                           IconSet["key"],
	".huskyrc":                            IconSet["husky"],
	".huskyrc.js":                         IconSet["husky"],
	".huskyrc.json":                       IconSet["husky"],
	".huskyrc.yaml":                       IconSet["husky"],
	".huskyrc.yml":                        IconSet["husky"],
	".io-config.json":                     IconSet["ionic"],
	".jestrc":                             IconSet["jest"],
	".jestrc.js":                          IconSet["jest"],
	".jestrc.json":                        IconSet["jest"],
	".jsbeautifyrc":                       IconSet["json"],
	".jscsrc":                             IconSet["json"],
	".jshintignore":                       IconSet["settings"],
	".jshintrc":                           IconSet["json"],
	".luacheckrc":                         IconSet["lua"],
	".mailmap":                            IconSet["email"],
	".mjmlconfig":                         IconSet["json"],
	".mocharc.js":                         IconSet["mocha"],
	".mocharc.json":                       IconSet["mocha"],
	".mocharc.jsonc":                      IconSet["mocha"],
	".mocharc.yaml":                       IconSet["mocha"],
	".mocharc.yml":                        IconSet["mocha"],
	".modernizrrc":                        IconSet["modernizr"],
	".modernizrrc.js":                     IconSet["modernizr"],
	".modernizrrc.json":                   IconSet["modernizr"],
	".mrconfig":                           IconSet["settings"],
	".netrc":                              IconSet["gnu"],
	".node-version":                       IconSet["nodejs"],
	".nowignore":                          IconSet["vercel"],
	".npmignore":                          IconSet["npm"],
	".npmrc":                              IconSet["npm"],
	".nuspec":                             IconSet["nuget"],
	".nuxtrc":                             IconSet["nuxt"],
	".nvmrc":                              IconSet["nodejs"],
	".postcssrc":                          IconSet["postcss"],
	".postcssrc.js":                       IconSet["postcss"],
	".postcssrc.json":                     IconSet["postcss"],
	".postcssrc.yml":                      IconSet["postcss"],
	".prettierignore":                     IconSet["prettier"],
	".prettierrc":                         IconSet["prettier"],
	".prettierrc.js":                      IconSet["prettier"],
	".prettierrc.json":                    IconSet["prettier"],
	".prettierrc.yaml":                    IconSet["prettier"],
	".prettierrc.yml":                     IconSet["prettier"],
	".pug-lintrc":                         IconSet["pug"],
	".pug-lintrc.js":                      IconSet["pug"],
	".pug-lintrc.json":                    IconSet["pug"],
	".python-version":                     IconSet["python-misc"],
	".releaserc":                          IconSet["semantic-release"],
	".releaserc.js":                       IconSet["semantic-release"],
	".releaserc.json":                     IconSet["semantic-release"],
	".releaserc.yaml":                     IconSet["semantic-release"],
	".releaserc.yml":                      IconSet["semantic-release"],
	".Rhistory":                           IconSet["r"],
	".stylelintignore":                    IconSet["stylelint"],
	".stylelintrc":                        IconSet["stylelint"],
	".stylelintrc.js":                     IconSet["stylelint"],
	".stylelintrc.json":                   IconSet["stylelint"],
	".stylelintrc.yaml":                   IconSet["stylelint"],
	".stylelintrc.yml":                    IconSet["stylelint"],
	".travis.yml":                         IconSet["travis"],
	".vercelignore":                       IconSet["vercel"],
	".wgetrc":                             IconSet["gnu"],
	".wget-hsts":                          IconSet["gnu"],
	".yardopts":                           IconSet["settings"],
	".yarn-integrity":                     IconSet["yarn"],
	".yarnclean":                          IconSet["yarn"],
	".yarnrc":                             IconSet["yarn"],
	".yarnrc.yaml":                        IconSet["yarn"],
	".yarnrc.yml":                         IconSet["yarn"],
	".zhistory":                           IconSet["console"],
	".zlogin":                             IconSet["console"],
	".zlogout":                            IconSet["console"],
	".zprofile":                           IconSet["console"],
	".zshenv":                             IconSet["console"],
	".zshrc":                              IconSet["console"],
	".zsh_aliases":                        IconSet["console"],
	".zsh_history":                        IconSet["console"],
	"androidmanifest.xml":                 IconSet["android"],
	"appfile":                             IconSet["fastlane"],
	"authors":                             IconSet["authors"],
	"authors.md":                          IconSet["authors"],
	"authors.txt":                         IconSet["authors"],
	"azure-pipelines.yaml":                IconSet["azure-pipelines"],
	"azure-pipelines.yml":                 IconSet["azure-pipelines"],
	"babel.config.js":                     IconSet["babel"],
	"babel.config.json":                   IconSet["babel"],
	"bitbucket-pipelines.yaml":            IconSet["bitbucket"],
	"bitbucket-pipelines.yml":             IconSet["bitbucket"],
	"bower.json":                          IconSet["bower"],
	"bun.lockb":                           IconSet["bun"],
	"cdp.pid":                             IconSet["json"],
	"changelog":                           IconSet["changelog"],
	"changelog.md":                        IconSet["changelog"],
	"changelog.txt":                       IconSet["changelog"],
	"changes":                             IconSet["changelog"],
	"changes.md":                          IconSet["changelog"],
	"changes.txt":                         IconSet["changelog"],
	"cmakecache.txt":                      IconSet["cmake"],
	"cmakelists.txt":                      IconSet["cmake"],
	"CNAME":                               IconSet["http"],
	"codecov.yml":                         IconSet["codecov"],
	"codeowners":                          IconSet["codeowners"],
	"code_of_conduct.md":                  IconSet["conduct"],
	"code_of_conduct.txt":                 IconSet["conduct"],
	"commitlint.config.js":                IconSet["commitlint"],
	"commitlint.config.ts":                IconSet["commitlint"],
	"composer.lock":                       IconSet["json"],
	"contributing.md":                     IconSet["contributing"],
	"credits":                             IconSet["credits"],
	"credits.md":                          IconSet["credits"],
	"credits.txt":                         IconSet["credits"],
	"docker-compose.ci.yml":               IconSet["docker"],
	"docker-compose.dev.yml":              IconSet["docker"],
	"docker-compose.local.yml":            IconSet["docker"],
	"docker-compose.override.yml":         IconSet["docker"],
	"docker-compose.prod.yml":             IconSet["docker"],
	"docker-compose.production.yml":       IconSet["docker"],
	"docker-compose.staging.yml":          IconSet["docker"],
	"docker-compose.test.yml":             IconSet["docker"],
	"docker-compose.yaml":                 IconSet["docker"],
	"docker-compose.yml":                  IconSet["docker"],
	"dockerfile":                          IconSet["docker"],
	"dockerfile.prod":                     IconSet["docker"],
	"dockerfile.production":               IconSet["docker"],
	"dune":                                IconSet["dune"],
	"dune-project":                        IconSet["dune"],
	"eslint.config.cjs":                   IconSet["eslint"],
	"eslint.config.d.mts":                 IconSet["eslint"],
	"eslint.config.js":                    IconSet["eslint"],
	"eslint.config.mjs":                   IconSet["eslint"],
	"eslint.config.ts":                    IconSet["eslint"],
	"fastfile":                            IconSet["fastlane"],
	"favicon.ico":                         IconSet["favicon"],
	"firebase.json":                       IconSet["firebase"],
	"firestore.indexes.json":              IconSet["firebase"],
	"firestore.rules":                     IconSet["firebase"],
	"gatsby-browser.js":                   IconSet["gatsby"],
	"gatsby-config.js":                    IconSet["gatsby"],
	"gatsby-node.js":                      IconSet["gatsby"],
	"gatsby-ssr.js":                       IconSet["gatsby"],
	"gatsby.config.js":                    IconSet["gatsby"],
	"gemfile":                             IconSet["gemfile"],
	"git-history":                         IconSet["git"],
	"go.mod":                              IconSet["go-mod"],
	"go.sum":                              IconSet["go-mod"],
	"gradle-wrapper.properties":           IconSet["gradle"],
	"gradle.properties":                   IconSet["gradle"],
	"gradlew":                             IconSet["gradle"],
	"gruntfile.babel.coffee":              IconSet["grunt"],
	"gruntfile.babel.js":                  IconSet["grunt"],
	"gruntfile.babel.ts":                  IconSet["grunt"],
	"gruntfile.coffee":                    IconSet["grunt"],
	"gruntfile.js":                        IconSet["grunt"],
	"gruntfile.ts":                        IconSet["grunt"],
	"gulpfile.babel.js":                   IconSet["gulp"],
	"gulpfile.js":                         IconSet["gulp"],
	"gulpfile.mjs":                        IconSet["gulp"],
	"gulpfile.ts":                         IconSet["gulp"],
	"husky.config.js":                     IconSet["husky"],
	"i18n.config.js":                      IconSet["i18n"],
	"i18n.config.ts":                      IconSet["i18n"],
	"i18n.options.mjs":                    IconSet["i18n"],
	"ionic.config.json":                   IconSet["ionic"],
	"jenkinsfile":                         IconSet["jenkins"],
	"jest.config.cjs":                     IconSet["jest"],
	"jest.config.js":                      IconSet["jest"],
	"jest.config.json":                    IconSet["jest"],
	"jest.config.mjs":                     IconSet["jest"],
	"jest.config.ts":                      IconSet["jest"],
	"jest.e2e.config.cjs":                 IconSet["jest"],
	"jest.e2e.config.js":                  IconSet["jest"],
	"jest.e2e.config.json":                IconSet["jest"],
	"jest.e2e.config.mjs":                 IconSet["jest"],
	"jest.e2e.config.ts":                  IconSet["jest"],
	"jest.json":                           IconSet["jest"],
	"jest.setup.js":                       IconSet["jest"],
	"jest.setup.ts":                       IconSet["jest"],
	"jest.teardown.js":                    IconSet["jest"],
	"karma-main.js":                       IconSet["karma"],
	"karma-main.ts":                       IconSet["karma"],
	"karma.conf.coffee":                   IconSet["karma"],
	"karma.conf.js":                       IconSet["karma"],
	"karma.conf.ts":                       IconSet["karma"],
	"karma.config.js":                     IconSet["karma"],
	"karma.config.ts":                     IconSet["karma"],
	"licence":                             IconSet["certificate"],
	"licence.md":                          IconSet["certificate"],
	"licence.txt":                         IconSet["certificate"],
	"license":                             IconSet["certificate"],
	"license.md":                          IconSet["certificate"],
	"license.txt":                         IconSet["certificate"],
	"makefile":                            IconSet["makefile"],
	"manifest.in":                         IconSet["python-misc"],
	"manifest.mf":                         IconSet["settings"],
	"milestones.md":                       IconSet["roadmap"],
	"milestones.txt":                      IconSet["roadmap"],
	"mocha.opts":                          IconSet["mocha"],
	"now.json":                            IconSet["vercel"],
	"nuget.config":                        IconSet["nuget"],
	"nuget.exe":                           IconSet["nuget"],
	"nuxt.config.js":                      IconSet["nuxt"],
	"nuxt.config.ts":                      IconSet["nuxt"],
	"nuxt.d.ts":                           IconSet["nuxt"],
	"package-lock.json":                   IconSet["nodejs"],
	"package.json":                        IconSet["nodejs"],
	"pipfile":                             IconSet["python-misc"],
	"pnpm-lock.yaml":                      IconSet["nodejs"],
	"postcss.config.js":                   IconSet["postcss"],
	"prettier.config.js":                  IconSet["prettier"],
	"prisma.yml":                          IconSet["prisma"],
	"procfile":                            IconSet["heroku"],
	"procfile.windows":                    IconSet["heroku"],
	"protractor.conf.coffee":              IconSet["protractor"],
	"protractor.conf.js":                  IconSet["protractor"],
	"protractor.conf.ts":                  IconSet["protractor"],
	"protractor.config.js":                IconSet["protractor"],
	"protractor.config.ts":                IconSet["protractor"],
	"readme":                              IconSet["readme"],
	"readme.md":                           IconSet["readme"],
	"readme.txt":                          IconSet["readme"],
	"release.config.js":                   IconSet["semantic-release"],
	"requirements.txt":                    IconSet["python-misc"],
	"roadmap.md":                          IconSet["roadmap"],
	"roadmap.txt":                         IconSet["roadmap"],
	"rollup-config.js":                    IconSet["rollup"],
	"rollup-config.ts":                    IconSet["rollup"],
	"rollup.config.base.js":               IconSet["rollup"],
	"rollup.config.base.ts":               IconSet["rollup"],
	"rollup.config.common.js":             IconSet["rollup"],
	"rollup.config.common.ts":             IconSet["rollup"],
	"rollup.config.dev.js":                IconSet["rollup"],
	"rollup.config.dev.ts":                IconSet["rollup"],
	"rollup.config.js":                    IconSet["rollup"],
	"rollup.config.prod.js":               IconSet["rollup"],
	"rollup.config.prod.ts":               IconSet["rollup"],
	"rollup.config.prod.vendor.js":        IconSet["rollup"],
	"rollup.config.prod.vendor.ts":        IconSet["rollup"],
	"rollup.config.ts":                    IconSet["rollup"],
	"routes.js":                           IconSet["routing"],
	"routes.jsx":                          IconSet["routing"],
	"routes.ts":                           IconSet["routing"],
	"routes.tsx":                          IconSet["routing"],
	"routing.js":                          IconSet["routing"],
	"routing.jsx":                         IconSet["routing"],
	"routing.ts":                          IconSet["routing"],
	"routing.tsx":                         IconSet["routing"],
	"security":                            IconSet["lock"],
	"security.md":                         IconSet["lock"],
	"security.txt":                        IconSet["lock"],
	"stryker.conf.js":                     IconSet["stryker"],
	"stryker.conf.json":                   IconSet["stryker"],
	"stylelint.config.js":                 IconSet["stylelint"],
	"svelte.config.js":                    IconSet["svelte"],
	"svelte.config.ts":                    IconSet["svelte"],
	"tailwind.config.cjs":                 IconSet["tailwindcss"],
	"tailwind.config.js":                  IconSet["tailwindcss"],
	"tailwind.config.mjs":                 IconSet["tailwindcss"],
	"tailwind.config.ts":                  IconSet["tailwindcss"],
	"tailwind.js":                         IconSet["tailwindcss"],
	"timeline.md":                         IconSet["roadmap"],
	"timeline.txt":                        IconSet["roadmap"],
	"tsconfig.json":                       IconSet["typescript-def"],
	"tsconfig.server.json":                IconSet["typescript-def"],
	"unlicense":                           IconSet["certificate"],
	"unlicense.md":                        IconSet["certificate"],
	"unlicense.txt":                       IconSet["certificate"],
	"vagrantfile":                         IconSet["vagrant"],
	"vercel.json":                         IconSet["vercel"],
	"vue.config.js":                       IconSet["vue-config"],
	"vue.config.ts":                       IconSet["vue-config"],
	"webpack.base.js":                     IconSet["webpack"],
	"webpack.base.ts":                     IconSet["webpack"],
	"webpack.client.js":                   IconSet["webpack"],
	"webpack.client.ts":                   IconSet["webpack"],
	"webpack.common.js":                   IconSet["webpack"],
	"webpack.common.ts":                   IconSet["webpack"],
	"webpack.config.babel.js":             IconSet["webpack"],
	"webpack.config.babel.ts":             IconSet["webpack"],
	"webpack.config.base.babel.js":        IconSet["webpack"],
	"webpack.config.base.babel.ts":        IconSet["webpack"],
	"webpack.config.base.js":              IconSet["webpack"],
	"webpack.config.base.ts":              IconSet["webpack"],
	"webpack.config.client.js":            IconSet["webpack"],
	"webpack.config.client.ts":            IconSet["webpack"],
	"webpack.config.coffee":               IconSet["webpack"],
	"webpack.config.common.babel.js":      IconSet["webpack"],
	"webpack.config.common.babel.ts":      IconSet["webpack"],
	"webpack.config.common.js":            IconSet["webpack"],
	"webpack.config.common.ts":            IconSet["webpack"],
	"webpack.config.dev.babel.js":         IconSet["webpack"],
	"webpack.config.dev.babel.ts":         IconSet["webpack"],
	"webpack.config.dev.js":               IconSet["webpack"],
	"webpack.config.dev.ts":               IconSet["webpack"],
	"webpack.config.js":                   IconSet["webpack"],
	"webpack.config.prod.babel.js":        IconSet["webpack"],
	"webpack.config.prod.babel.ts":        IconSet["webpack"],
	"webpack.config.prod.js":              IconSet["webpack"],
	"webpack.config.prod.ts":              IconSet["webpack"],
	"webpack.config.production.babel.js":  IconSet["webpack"],
	"webpack.config.production.babel.ts":  IconSet["webpack"],
	"webpack.config.production.js":        IconSet["webpack"],
	"webpack.config.production.ts":        IconSet["webpack"],
	"webpack.config.server.js":            IconSet["webpack"],
	"webpack.config.server.ts":            IconSet["webpack"],
	"webpack.config.staging.babel.js":     IconSet["webpack"],
	"webpack.config.staging.babel.ts":     IconSet["webpack"],
	"webpack.config.staging.js":           IconSet["webpack"],
	"webpack.config.staging.ts":           IconSet["webpack"],
	"webpack.config.test.js":              IconSet["webpack"],
	"webpack.config.test.ts":              IconSet["webpack"],
	"webpack.config.ts":                   IconSet["webpack"],
	"webpack.config.vendor.js":            IconSet["webpack"],
	"webpack.config.vendor.production.js": IconSet["webpack"],
	"webpack.config.vendor.production.ts": IconSet["webpack"],
	"webpack.config.vendor.ts":            IconSet["webpack"],
	"webpack.dev.js":                      IconSet["webpack"],
	"webpack.dev.ts":                      IconSet["webpack"],
	"webpack.development.js":              IconSet["webpack"],
	"webpack.development.ts":              IconSet["webpack"],
	"webpack.dist.js":                     IconSet["webpack"],
	"webpack.dist.ts":                     IconSet["webpack"],
	"webpack.js":                          IconSet["webpack"],
	"webpack.prod.js":                     IconSet["webpack"],
	"webpack.prod.ts":                     IconSet["webpack"],
	"webpack.production.js":               IconSet["webpack"],
	"webpack.production.ts":               IconSet["webpack"],
	"webpack.server.js":                   IconSet["webpack"],
	"webpack.server.ts":                   IconSet["webpack"],
	"webpack.test.js":                     IconSet["webpack"],
	"webpack.test.ts":                     IconSet["webpack"],
	"webpack.ts":                          IconSet["webpack"],
	"webpackfile.js":                      IconSet["webpack"],
	"webpackfile.ts":                      IconSet["webpack"],
	"yarn-error.log":                      IconSet["yarn"],
	"yarn.lock":                           IconSet["yarn"],
}