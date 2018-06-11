var webui=angular.module("webui","webui.services.utils webui.services.deps webui.services.base64 webui.services.configuration webui.services.rpc webui.services.modals webui.services.alerts webui.services.settings webui.services.settings.filters webui.filters.bytes webui.filters.url webui.directives.chunkbar webui.directives.dgraph webui.directives.fselect webui.directives.fileselect webui.ctrls.download webui.ctrls.nav webui.ctrls.modal webui.ctrls.alert webui.ctrls.props ui.bootstrap pascalprecht.translate".split(" "));
function mergeTranslation(b,a){for(var c in a)a.hasOwnProperty(c)&&(b[c]&&b[c].length||(b[c]=a[c]));return b}
webui.config(function ($translateProvider, $locationProvider) {
  $translateProvider
      .translations('en_US', translations.en_US)
      .translations('zh_CN', mergeTranslation(translations.zh_CN, translations.en_US))
      .useSanitizeValueStrategy('escapeParameters')
      .determinePreferredLanguage();

      $locationProvider.html5Mode({
        enabled: true,
        requireBase: false
      });
});
$(function(){String.prototype.startsWith||Object.defineProperty(String.prototype,"startsWith",{enumerable:!1,configurable:!1,writable:!1,value:function(b,a){a=a||0;return this.indexOf(b,a)===a}});angular.bootstrap(document,["webui"])});
