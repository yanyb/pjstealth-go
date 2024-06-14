package pjstealth_go

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/yanyb/pjstealth-go/feature"
	"github.com/yanyb/pjstealth-go/utils/random"
)

type KwArgs map[string]any

var scripts = map[string]string{
	"chrome_csi":                    feature.FromFile("chrome.csi.js"),
	"chrome_app":                    feature.FromFile("chrome.app.js"),
	"chrome_runtime":                feature.FromFile("chrome.runtime.js"),
	"chrome_load_times":             feature.FromFile("chrome.load.times.js"),
	"chrome_hairline":               feature.FromFile("chrome.hairline.js"),
	"generate_magic_arrays":         feature.FromFile("generate.magic.arrays.js"),
	"iframe_content_window":         feature.FromFile("iframe.contentWindow.js"),
	"media_codecs":                  feature.FromFile("media.codecs.js"),
	"navigator_vendor":              feature.FromFile("navigator.vendor.js"),
	"navigator_plugins":             feature.FromFile("navigator.plugins.js"),
	"navigator_permissions":         feature.FromFile("navigator.permissions.js"),
	"navigator_languages":           feature.FromFile("navigator.languages.js"),
	"navigator_platform":            feature.FromFile("navigator.platform.js"),
	"navigator_user_agent":          feature.FromFile("navigator.userAgent.js"),
	"navigator_hardwareConcurrency": feature.FromFile("navigator.hardwareConcurrency.js"),
	"outerdimensions":               feature.FromFile("window.outerdimensions.js"),
	"utils":                         feature.FromFile("utils.js"),
	"webdriver":                     feature.FromFile("chrome.webdriver.js"),
	"webgl_vendor":                  feature.FromFile("webgl.vendor.js"),
	"navigator_appVersion":          feature.FromFile("navigator.appVersion.js"),
	"navigator_deviceMemory":        feature.FromFile("navigator.deviceMemory.js"),
	"navigator_language":            feature.FromFile("navigator.language.js"),
	"navigator_userAgentData":       feature.FromFile("navigator.userAgentData.js"),
	"chrome_canvasfeature":          feature.FromFile("chrome.canvasfeature.js"),
	"chrome_clientrectfeature":      feature.FromFile("chrome.clientrectfeature.js"),
	"chrome_cssfeature":             feature.FromFile("chrome.cssfeature.js"),
	"chrome_fontsfeature":           feature.FromFile("chrome.fontsfeature.js"),
	"chrome_webrtc":                 feature.FromFile("chrome.webrtc.js"),
	"hookfuc_headless":              feature.FromFile("hookfuc.headless.js"),
	"chrome_videofeature":           feature.FromFile("chrome.videofeature.js"),
	"chrome_canvasfeature2":         feature.FromFile("chrome.canvasfeature2.js"),
	"chrome_screen_colordepth":      feature.FromFile("chrome.screen.colordepth.js"),
}

type StealthConfig struct {
	vendor               string
	renderer             string
	navVendor            string
	runOnInsecureOrigins any
	mouseDetail          int
	mouseButton          int
	mouseButtons         int
	webDriver            bool
	webglVendor          bool
	navigatorPlugins     bool
	navigatorPermissions bool
	mediaCodecs          bool
	iframeContentWindow  bool
	chromeRuntime        bool
	chromeLoadTimes      bool
	chromeCsi            bool
	chromeApp            bool
	outerDimensions      bool
	hairline             bool
	// 随机特征开启默认为false
	randomFeature bool
	//local variable
	navigatorUserAgent           string
	navigatorPlatform            string
	navigatorLanguages           []string
	navigatorLanguage            string
	browserVersion               string
	sysPlatform                  interface{}
	navigatorHardwareConcurrency int
	deviceMemory                 int
	cssFeature                   KwArgs
	fontsFeature                 KwArgs
	webRTC                       bool
	canvasFeature                KwArgs
	videoFeature                 KwArgs
	clientRectFeature            bool
	headlessCheck                bool
	isMobile                     bool
	screenColorDepth             int
	opts                         KwArgs
}

func (kwArgs KwArgs) get(kInput string) any {
	for k, v := range kwArgs {
		if k == kInput {
			return v
		}
	}
	return nil
}

func NewStealthConfig(navigatorUserAgent string, navigatorPlatform string, kwArgs KwArgs) *StealthConfig {
	envData := NewEnvData()
	sc := &StealthConfig{
		vendor:               "Intel Inc.",
		renderer:             "Intel Iris OpenGL Engine",
		navVendor:            "Google Inc.",
		runOnInsecureOrigins: nil,
		webDriver:            true,
		webglVendor:          true,
		navigatorPlugins:     true,
		navigatorPermissions: true,
		mediaCodecs:          true,
		iframeContentWindow:  true,
		chromeRuntime:        true,
		chromeLoadTimes:      true,
		chromeCsi:            true,
		chromeApp:            true,
		outerDimensions:      true,
		hairline:             true,
		randomFeature:        true,
	}

	sc.navigatorUserAgent = navigatorUserAgent
	sc.navigatorPlatform = navigatorPlatform
	if v := kwArgs.get("random_feature"); v != nil {
		sc.randomFeature = v.(bool)
	}
	if sc.randomFeature {
		sc.navigatorLanguages = envData.get("languages").([]string)
		sc.navigatorLanguage = envData.get("language").(string)

		//user-agent mac(m系列和非m系列)， windows版对应
		if v := envData.get("user_agent"); v != nil {
			sc.navigatorUserAgent = v.(string)
		}
		if v := envData.get("browser_version"); v != nil {
			sc.browserVersion = envData.get("browser_version").(string)
		}

		if sc.navigatorUserAgent != "" {
			if strings.Contains(strings.ToLower(sc.navigatorUserAgent), "mac") {
				sc.navigatorPlatform = "MacIntel"

				if !strings.Contains(strings.ToLower(sc.navigatorUserAgent), "intel") {
					for {
						tmpWebglInfo := random.Choice(envData.get(sc.navigatorPlatform).(KwArgs).get("webgl_infos").([][]string))
						if strings.Contains(strings.ToLower(strings.Join(tmpWebglInfo, "")), "m1") ||
							strings.Contains(strings.ToLower(strings.Join(tmpWebglInfo, "")), "m2") {
							sc.vendor = tmpWebglInfo[0]
							sc.renderer = tmpWebglInfo[1]
							break
						}
					}
				} else {
					for {
						tmpWebglInfo := random.Choice(envData.get(sc.navigatorPlatform).(KwArgs).get("webgl_infos").([][]string))
						if strings.Contains(strings.ToLower(strings.Join(tmpWebglInfo, "")), "m1") ||
							strings.Contains(strings.ToLower(strings.Join(tmpWebglInfo, "")), "m2") {
							continue
						}
						sc.vendor = tmpWebglInfo[0]
						sc.renderer = tmpWebglInfo[1]
						break
					}
				}
			}

			if strings.Contains(strings.ToLower(sc.navigatorUserAgent), "windows") {
				sc.navigatorPlatform = "Win64"
				tmpWebglInfo := envData.get(sc.navigatorPlatform).(KwArgs).get("webgl_infos").([]string)
				sc.vendor = tmpWebglInfo[0]
				sc.renderer = tmpWebglInfo[1]
			}

			if strings.Contains(strings.ToLower(sc.navigatorUserAgent), "linux") {
				sc.navigatorPlatform = "Linux x86_64"
				tmpWebglInfo := envData.get(sc.navigatorPlatform).(KwArgs).get("webgl_infos").([]string)
				sc.vendor = tmpWebglInfo[0]
				sc.renderer = tmpWebglInfo[1]
			}

			r := regexp.MustCompile(`Chrome/(\d+)`)
			sc.browserVersion = r.FindStringSubmatch(sc.navigatorUserAgent)[1]
		}

		if sc.navigatorPlatform == "" {
			sc.navigatorPlatform = random.Choice([]string{"MacIntel", "Win64"})
		}

		sc.sysPlatform = envData.get(sc.navigatorPlatform).(KwArgs).get("sys_platform").(string)
		sc.navigatorHardwareConcurrency = envData.get("navigator_hardware_concurrency").(int)
		sc.deviceMemory = envData.get("device_memory").(int)
		sc.cssFeature = envData.get("cssfeature").(KwArgs)
		sc.fontsFeature = envData.get("fontsfeature").(KwArgs)
		sc.webRTC = envData.get("webrtc").(bool)
		sc.canvasFeature = envData.get("canvasfeature").(KwArgs)
		sc.videoFeature = envData.get("videofeature").(KwArgs)
		sc.clientRectFeature = envData.get("clientrectfeature").(bool)
		sc.headlessCheck = envData.get("headless_check").(bool)
		sc.isMobile = false
		sc.screenColorDepth = envData.get("screen_color_depth").(int)
	}

	if v := kwArgs.get("navigator_languages"); v != nil {
		sc.navigatorLanguages = v.([]string)
	}

	if v := kwArgs.get("navigator_language"); v != nil {
		sc.navigatorLanguage = v.(string)
	}

	if v := kwArgs.get("navigator_platform"); v != nil {
		sc.navigatorPlatform = v.(string)
	}

	if v := kwArgs.get("user_agent"); v != nil {
		sc.navigatorUserAgent = v.(string)
	}

	if sc.navigatorPlatform != "" {
		if sc.navigatorPlatform[0:1] == "W" {
			sc.sysPlatform = "Windows"
		} else {
			sc.sysPlatform = "macOS"
		}
	} else {
		sc.sysPlatform = nil
	}

	if v := kwArgs.get("navigator_hardware_concurrency"); v != nil {
		sc.navigatorHardwareConcurrency = v.(int)
	}

	if v := kwArgs.get("device_memory"); v != nil {
		sc.deviceMemory = v.(int)
	}

	if v := kwArgs.get("is_mobile"); v != nil {
		sc.isMobile = v.(bool)
	}

	if v := kwArgs.get("browser_version"); v != nil {
		sc.browserVersion = v.(string)
	}

	if v := kwArgs.get("screen_color_depth"); v != nil {
		sc.screenColorDepth = v.(int)
	}

	if v := kwArgs.get("vendor"); v != nil {
		sc.vendor = v.(string)
	}

	if v := kwArgs.get("renderer"); v != nil {
		sc.renderer = v.(string)
	}

	if v := kwArgs.get("nav_vendor"); v != nil {
		sc.navVendor = v.(string)
	}

	if v := kwArgs.get("runOnInsecureOrigins"); v != nil {
		sc.runOnInsecureOrigins = v
	}

	if v := kwArgs.get("webdriver"); v != nil {
		sc.webDriver = v.(bool)
	}

	if v := kwArgs.get("webgl_vendor"); v != nil {
		sc.webglVendor = v.(bool)
	}

	if v := kwArgs.get("navigator_plugins"); v != nil {
		sc.navigatorPlugins = v.(bool)
	}

	if v := kwArgs.get("navigator_permissions"); v != nil {
		sc.navigatorPermissions = v.(bool)
	}

	if v := kwArgs.get("media_codecs"); v != nil {
		sc.mediaCodecs = v.(bool)
	}

	if v := kwArgs.get("iframe_content_window"); v != nil {
		sc.iframeContentWindow = v.(bool)
	}

	if v := kwArgs.get("chrome_runtime"); v != nil {
		sc.chromeRuntime = v.(bool)
	}

	if v := kwArgs.get("chrome_load_times"); v != nil {
		sc.chromeLoadTimes = v.(bool)
	}

	if v := kwArgs.get("chrome_csi"); v != nil {
		sc.chromeCsi = v.(bool)
	}

	if v := kwArgs.get("chrome_app"); v != nil {
		sc.chromeApp = v.(bool)
	}

	if v := kwArgs.get("outerdimensions"); v != nil {
		sc.outerDimensions = v.(bool)
	}

	if v := kwArgs.get("hairline"); v != nil {
		sc.hairline = v.(bool)
	}

	if v := kwArgs.get("cssfeature"); v != nil {
		sc.cssFeature = v.(KwArgs)
	}

	if v := kwArgs.get("fontsfeature"); v != nil {
		sc.fontsFeature = v.(KwArgs)
	}

	if v := kwArgs.get("webrtc"); v != nil {
		sc.webRTC = v.(bool)
	}

	if v := kwArgs.get("canvasfeature"); v != nil {
		sc.canvasFeature = v.(KwArgs)
	}

	if v := kwArgs.get("videofeature"); v != nil {
		sc.videoFeature = v.(KwArgs)
	}

	if v := kwArgs.get("clientrectfeature"); v != nil {
		sc.clientRectFeature = v.(bool)
	}

	if v := kwArgs.get("headless_check"); v != nil {
		sc.headlessCheck = v.(bool)
	}

	sc.opts = KwArgs{
		"languages":                      sc.navigatorLanguages,
		"language":                       sc.navigatorLanguage,
		"webgl_vendor":                   sc.vendor,
		"webgl_renderer":                 sc.renderer,
		"navigator_vendor":               sc.navVendor,
		"navigator_platform":             sc.navigatorPlatform,
		"navigator_user_agent":           sc.navigatorUserAgent,
		"navigator_app_version":          strings.ReplaceAll(sc.navigatorUserAgent, "Mozilla/", ""),
		"runOnInsecureOrigins":           sc.runOnInsecureOrigins,
		"navigator_hardware_concurrency": sc.navigatorHardwareConcurrency,
		"device_memory":                  sc.deviceMemory,
		"user_agent_data": KwArgs{
			"brands": []KwArgs{
				{"brand": "Not)A;Brand", "version": "24"},
				{"brand": "Chromium", "version": sc.browserVersion},
				{"brand": "Google Chrome", "version": sc.browserVersion},
			},
			"mobile":   sc.isMobile,
			"platform": sc.sysPlatform,
		},
		"cssfeature":         sc.cssFeature,
		"fontsfeature":       sc.fontsFeature,
		"webrtc":             sc.webRTC,
		"canvasfeature":      sc.canvasFeature,
		"videofeature":       sc.videoFeature,
		"clientrectfeature":  sc.clientRectFeature,
		"headless_check":     sc.headlessCheck,
		"fonts_start":        0,
		"screen_color_depth": sc.screenColorDepth,
		"mouse_event": KwArgs{
			"detail":  sc.mouseDetail,
			"button":  sc.mouseButton,
			"buttons": sc.mouseButtons,
		},
	}

	return sc
}

func (sc *StealthConfig) enabledScripts() []string {
	var scriptsContent []string

	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(sc.opts)
	if err != nil {
		panic(err)
	}
	scriptsContent = append(scriptsContent, "const opts = "+buf.String())
	scriptsContent = append(scriptsContent, scripts["utils"])
	scriptsContent = append(scriptsContent, scripts["generate_magic_arrays"])
	scriptsContent = append(scriptsContent, scripts["webgl_vendor"])

	if sc.chromeApp {
		scriptsContent = append(scriptsContent, scripts["chrome_app"])
	}
	if sc.chromeRuntime {
		scriptsContent = append(scriptsContent, scripts["chrome_runtime"])
	}
	if sc.chromeLoadTimes {
		scriptsContent = append(scriptsContent, scripts["chrome_load_times"])
	}
	if sc.chromeCsi {
		scriptsContent = append(scriptsContent, scripts["chrome_csi"])
	}
	if sc.iframeContentWindow {
		scriptsContent = append(scriptsContent, scripts["iframe_content_window"])
	}
	if sc.mediaCodecs {
		scriptsContent = append(scriptsContent, scripts["media_codecs"])
	}
	if sc.navigatorPlugins {
		scriptsContent = append(scriptsContent, scripts["navigator_plugins"])
	}
	if sc.navigatorPermissions {
		scriptsContent = append(scriptsContent, scripts["navigator_permissions"])
	}
	if sc.webDriver {
		scriptsContent = append(scriptsContent, scripts["webdriver"])
	}
	if sc.outerDimensions {
		scriptsContent = append(scriptsContent, scripts["outerdimensions"])
	}
	if sc.hairline {
		scriptsContent = append(scriptsContent, scripts["chrome_hairline"])
	}
	if sc.opts.get("navigator_languages") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_languages"])
	}
	if sc.opts.get("navigator_vendor") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_vendor"])
	}
	if sc.opts.get("navigator_platform") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_platform"])
	}
	if sc.opts.get("navigator_user_agent") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_user_agent"])
		scriptsContent = append(scriptsContent, scripts["navigator_appVersion"])
	}
	if sc.opts.get("language") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_language"])
	}
	if sc.opts.get("user_agent_data") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_userAgentData"])
	}
	if sc.opts.get("navigator_hardware_concurrency") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_hardwareConcurrency"])
	}
	if sc.opts.get("device_memory") != nil {
		scriptsContent = append(scriptsContent, scripts["navigator_deviceMemory"])
	}
	if sc.opts.get("cssfeature") != nil {
		scriptsContent = append(scriptsContent, scripts["chrome_cssfeature"])
	}
	if sc.opts.get("fontsfeature") != nil {
		scriptsContent = append(scriptsContent, scripts["chrome_fontsfeature"])
	}
	if sc.opts.get("webrtc") != nil {
		scriptsContent = append(scriptsContent, scripts["chrome_webrtc"])
	}
	if sc.opts.get("headless_check") != nil {
		scriptsContent = append(scriptsContent, scripts["hookfuc_headless"])
	}
	if sc.opts.get("canvasfeature") != nil {
		//scriptsContent = append(scriptsContent, scripts["chrome_canvasfeature"])
		//scriptsContent = append(scriptsContent, scripts["chrome_canvasfeature2"])
	}
	if sc.opts.get("videofeature") != nil {
		scriptsContent = append(scriptsContent, scripts["chrome_videofeature"])
	}
	if sc.opts.get("clientrectfeature") != nil {
		scriptsContent = append(scriptsContent, scripts["chrome_clientrectfeature"])
	}
	if sc.opts.get("screen_color_depth") != nil {
		scriptsContent = append(scriptsContent, scripts["chrome_screen_colordepth"])
	}
	return scriptsContent
}
