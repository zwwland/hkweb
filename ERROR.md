|通用错误码|-|-|
| --- | --- | --- |
|NET_DVR_NOERROR|0|没有错误|
|NET_DVR_PASSWORD_ERROR|1|用户名密码错误|
|NET_DVR_NOENOUGHPRI|2|权限不足|
|NET_DVR_NOINIT|3|没有初始化|
|NET_DVR_CHANNEL_ERROR|4|通道号错误|
|NET_DVR_OVER_MAXLINK|5|连接到DVR的客户端个数超过最大|
|NET_DVR_VERSIONNOMATCH|6|版本不匹配|
|NET_DVR_NETWORK_FAIL_CONNECT|7|连接服务器失败|
|NET_DVR_NETWORK_SEND_ERROR|8|向服务器发送失败|
|NET_DVR_NETWORK_RECV_ERROR|9|从服务器接收数据失败|
|NET_DVR_NETWORK_RECV_TIMEOUT|10|从服务器接收数据超时|
|NET_DVR_NETWORK_ERRORDATA|11|传送的数据有误|
|NET_DVR_ORDER_ERROR|12|调用次序错误|
|NET_DVR_OPERNOPERMIT|13|无此权限|
|NET_DVR_COMMANDTIMEOUT|14|DVR命令执行超时|
|NET_DVR_ERRORSERIALPORT|15|串口号错误|
|NET_DVR_ERRORALARMPORT|16|报警端口错误|
|NET_DVR_PARAMETER_ERROR|17|参数错误|
|NET_DVR_CHAN_EXCEPTION|18|服务器通道处于错误状态|
|NET_DVR_NODISK|19|没有硬盘|
|NET_DVR_ERRORDISKNUM|20|硬盘号错误|
|NET_DVR_DISK_FULL|21|服务器硬盘满|
|NET_DVR_DISK_ERROR|22|服务器硬盘出错|
|NET_DVR_NOSUPPORT|23|服务器不支持|
|NET_DVR_BUSY|24|服务器忙|
|NET_DVR_MODIFY_FAIL|25|服务器修改不成功|
|NET_DVR_PASSWORD_FORMAT_ERROR|26|密码输入格式不正确|
|NET_DVR_DISK_FORMATING|27|硬盘正在格式化，不能启动操作|
|NET_DVR_DVRNORESOURCE|28|DVR资源不足|
|NET_DVR_DVROPRATEFAILED|29|DVR操作失败|
|NET_DVR_OPENHOSTSOUND_FAIL|30|打开PC声音失败|
|NET_DVR_DVRVOICEOPENED|31|服务器语音对讲被占用|
|NET_DVR_TIMEINPUTERROR|32|时间输入不正确|
|NET_DVR_NOSPECFILE|33|回放时服务器没有指定的文件|
|NET_DVR_CREATEFILE_ERROR|34|创建文件出错|
|NET_DVR_FILEOPENFAIL|35|打开文件出错|
|NET_DVR_OPERNOTFINISH|36|上次的操作还没有完成|
|NET_DVR_GETPLAYTIMEFAIL|37|获取当前播放的时间出错|
|NET_DVR_PLAYFAIL|38|播放出错|
|NET_DVR_FILEFORMAT_ERROR|39|文件格式不正确|
|NET_DVR_DIR_ERROR|40|路径错误|
|NET_DVR_ALLOC_RESOURCE_ERROR|41|资源分配错误|
|NET_DVR_AUDIO_MODE_ERROR|42|声卡模式错误|
|NET_DVR_NOENOUGH_BUF|43|缓冲区太小|
|NET_DVR_CREATESOCKET_ERROR|44|创建SOCKET出错|
|NET_DVR_SETSOCKET_ERROR|45|设置SOCKET出错|
|NET_DVR_MAX_NUM|46|个数达到最大|
|NET_DVR_USERNOTEXIST|47|用户不存在|
|NET_DVR_WRITEFLASHERROR|48|写FLASH出错|
|NET_DVR_UPGRADEFAIL|49|DVR升级失败|
|NET_DVR_CARDHAVEINIT|50|解码卡已经初始化过|
|NET_DVR_PLAYERFAILED|51|调用播放库中某个函数失败|
|NET_DVR_MAX_USERNUM|52|设备端用户数达到最大|
|NET_DVR_GETLOCALIPANDMACFAIL|53|获得客户端的IP地址或物理地址失败|
|NET_DVR_NOENCODEING|54|该通道没有编码|
|NET_DVR_IPMISMATCH|55|IP地址不匹配|
|NET_DVR_MACMISMATCH|56|MAC地址不匹配|
|NET_DVR_UPGRADELANGMISMATCH|57|升级文件语言不匹配|
|NET_DVR_MAX_PLAYERPORT|58|播放器路数达到最大|
|NET_DVR_NOSPACEBACKUP|59|备份设备中没有足够空间进行备份|
|NET_DVR_NODEVICEBACKUP|60|没有找到指定的备份设备|
|NET_DVR_PICTURE_BITS_ERROR|61|图像素位数不符，限24色|
|NET_DVR_PICTURE_DIMENSION_ERROR|62|图片高宽超限，|限128256|
|NET_DVR_PICTURE_SIZ_ERROR|63|图片大小超限，限100K|
|NET_DVR_LOADPLAYERSDKFAILED|64|载入当前目录下Player|Sdk出错|
|NET_DVR_LOADPLAYERSDKPROC_ERROR|65|找不到Player|Sdk中某个函数入口|
|NET_DVR_LOADDSSDKFAILED|66|载入当前目录下DSsdk出错|
|NET_DVR_LOADDSSDKPROC_ERROR|67|找不到DsSdk中某个函数入口|
|NET_DVR_DSSDK_ERROR|68|调用硬解码库DsSdk中某个函数失败|
|NET_DVR_VOICEMONOPOLIZE|69|声卡被独占|
|NET_DVR_JOINMULTICASTFAILED|70|加入多播组失败|
|NET_DVR_CREATEDIR_ERROR|71|建立日志文件目录失败|
|NET_DVR_BINDSOCKET_ERROR|72|绑定套接字失败|
|NET_DVR_SOCKETCLOSE_ERROR|73|socket连接中断，此错误通常是由于连接中断或目的地不可达|
|NET_DVR_USERID_ISUSING|74|注销时用户ID正在进行某操作|
|NET_DVR_SOCKETLISTEN_ERROR|75|监听失败|
|NET_DVR_PROGRAM_EXCEPTION|76|程序异常|
|NET_DVR_WRITEFILE_FAILED|77|写文件失败|
|NET_DVR_FORMAT_READONLY|78|禁止格式化只读硬盘|
|NET_DVR_WITHSAMEUSERNAME|79|用户配置结构中存在相同的用户名|
|NET_DVR_DEVICETYPE_ERROR|80|导入参数时设备型号不匹配|
|NET_DVR_LANGUAGE_ERROR|81|导入参数时语言不匹配|
|NET_DVR_PARAVERSION_ERROR|82|导入参数时软件版本不匹配|
|NET_DVR_IPCHAN_NOTALIVE|83|预览时外接IP通道不在线|
|NET_DVR_RTSP_SDK_ERROR|84|加载高清IPC通讯库StreamTransClient.dll失败|
|NET_DVR_CONVERT_SDK_ERROR|85|加载转码库失败|
|NET_DVR_IPC_COUNT_OVERFLOW|86|超出最大的ip接入通道数|
|NET_DVR_MAX_ADD_NUM|87|添加标签(一个文件片段64)等个数达到最大|
|NET_DVR_PARAMMODE_ERROR|88|图像增强仪，参数模式错误（用于硬件设置时，客户端进行软件设置时错误值）|
|NET_DVR_CODESPITTER_OFFLINE|89|视频综合平台，码分器不在线|
|NET_DVR_BACKUP_COPYING|90|设备正在备份|
|NET_DVR_CHAN_NOTSUPPORT|91|通道不支持该操作|
|NET_DVR_CALLINEINVALID|92|高度线位置太集中或长度线不够倾斜|
|NET_DVR_CALCANCELCONFLICT|93|取消标定冲突，如果设置了规则及全局的实际大小尺寸过滤|
|NET_DVR_CALPOINTOUTRANGE|94|标定点超出范围|
|NET_DVR_FILTERRECTINVALID|95|尺寸过滤器不符合要求|
|NET_DVR_DDNS_DEVOFFLINE|96|设备没有注册到ddns上|
|NET_DVR_DDNS_INTER_ERROR|97|DDNS|服务器内部错误|
|NET_DVR_FUNCTION_NOT_SUPPORT_OS|98|此功能不支持该操作系统|
|NET_DVR_DEC_CHAN_REBIND|99|解码通道绑定显示输出次数受限|
|NET_DVR_INTERCOM_SDK_ERROR|100|加载当前目录下的语音对讲库失败|
|NET_DVR_NO_CURRENT_UPDATEFILE|101|没有正确的升级包|
|NET_DVR_USER_NOT_SUCC_LOGIN|102|用户还没登陆成功|
|NET_DVR_USE_LOG_SWITCH_FILE|103|正在使用日志开关文件|
|NET_DVR_POOL_PORT_EXHAUST|104|端口池中用于绑定的端口已耗尽|
|NET_DVR_PACKET_TYPE_NOT_SUPPORT|105|码流封装格式错误|
|NET_DVR_IPPARA_IPID_ERROR|106|IP接入配置时IPID有误|
|NET_DVR_LOAD_HCPREVIEW_SDK_ERROR|107|预览组件加载失败|
|NET_DVR_LOAD_HCVOICETALK_SDK_ERROR|108|语音组件加载失败|
|NET_DVR_LOAD_HCALARM_SDK_ERROR|109|报警组件加载失败|
|NET_DVR_LOAD_HCPLAYBACK_SDK_ERROR|110|回放组件加载失败|
|NET_DVR_LOAD_HCDISPLAY_SDK_ERROR|111|显示组件加载失败|
|NET_DVR_LOAD_HCINDUSTRY_SDK_ERROR|112|行业应用组件加载失败|
|NET_DVR_LOAD_HCGENERALCFGMGR_SDK_ERROR|113|通用配置管理组件加载失败|
|NET_DVR_CORE_VER_MISMATCH|121|单独加载组件时，组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCPREVIEW|122|预览组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCVOICETALK|123|语音组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCALARM|124|报警组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCPLAYBACK|125|回放组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCDISPLAY|126|显示组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCINDUSTRY|127|行业应用组件与core版本不匹配|
|NET_DVR_CORE_VER_MISMATCH_HCGENERALCFGMGR|128|通用配置管理组件与core版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCPREVIEW|136|预览组件与HCNetSDK版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCVOICETALK|137|语音组件与HCNetSDK版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCALARM|138|报警组件与HCNetSDK版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCPLAYBACK|139|回放组件与HCNetSDK版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCDISPLAY|140|显示组件与HCNetSDK版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCINDUSTRY|141|行业应用组件与HCNetSDK版本不匹配|
|NET_DVR_COM_VER_MISMATCH_HCGENERALCFGMGR|142|通用配置管理组件与HCNetSDK版本不匹配|
|NET_DVR_ALIAS_DUPLICATE|150|别名重复，通过别名或者序列号来访问设备的新版本ddns的配置|
|NET_DVR_INVALID_COMMUNICATION|151|无效通信|
|NET_DVR_USERNAME_NOT_EXIST|152|用户名不存在|
|NET_DVR_USER_LOCKED|153|用户被锁定|
|NET_DVR_INVALID_USERID|154|无效用户ID|
|NET_DVR_LOW_LOGIN_VERSION|155|登录版本低|
|NET_DVR_LOAD_LIBEAY32_DLL_ERROR|156|加载libeay32.dll库失败|
|NET_DVR_LOAD_SSLEAY32_DLL_ERROR|157|加载ssleay32.dll库失败|
|NET_DVR_COUNTRYID_NOT_EXIST|164|国家编号不存在,|hiddns上根据国家编号查询服务器地址时不存在这个国家编号|
|NET_DVR_TEST_SERVER_FAIL_CONNECT|165|连接测试服务器失败|
|NET_DVR_NAS_SERVER_INVALID_DIR|166|NAS服务器挂载目录失败，目录无效|
|NET_DVR_NAS_SERVER_NOENOUGH_PRI|167|NAS服务器挂载目录失败，没有权限|
|NET_DVR_EMAIL_SERVER_NOT_CONFIG_DNS|168|服务器使用域名，但是没有配置DNS，可能造成域名无效。|
|NET_DVR_EMAIL_SERVER_NOT_CONFIG_GATEWAY|169|没有配置网关，可能造成发送邮件失败。|
|NET_DVR_TEST_SERVER_PASSWORD_ERROR|170|用户名密码不正确，测试服务器的用户名或密码错误|
|NET_DVR_EMAIL_SERVER_CONNECT_EXCEPTION_WITH_SMTP|171|设备和smtp服务器交互异常|
|NET_DVR_FTP_SERVER_FAIL_CREATE_DIR|172|FTP服务器创建目录失败|
|NET_DVR_FTP_SERVER_NO_WRITE_PIR|173|FTP服务器没有写入权限|
|NET_DVR_IP_CONFLICT|174|IP冲突|
|NET_DVR_INSUFFICIENT_STORAGEPOOL_SPACE|175|存储池空间已满|
|NET_DVR_STORAGEPOOL_INVALID|176|云服务器存储池无效,没有配置存储池或者存储池ID错误|
|NET_DVR_EFFECTIVENESS_REBOOT|177|生效需要重启|
|阵列错误码|||
|RAID_ERROR_INDEX|200|阵列错误码索引|
|NET_DVR_NAME_NOT_ONLY|(RAID_ERROR_INDEX|+|0)|名称已存在|
|NET_DVR_OVER_MAX_ARRAY|(RAID_ERROR_INDEX|+|1|)|阵列达到上限|
|NET_DVR_OVER_MAX_VD|(RAID_ERROR_INDEX|+|2|)|虚拟磁盘达到上限|
|NET_DVR_VD_SLOT_EXCEED|(RAID_ERROR_INDEX|+|3|)|虚拟磁盘槽位已满|
|NET_DVR_PD_STATUS_INVALID|(RAID_ERROR_INDEX|+|4|)|重建阵列所需物理磁盘状态错误|
|NET_DVR_PD_BE_DEDICATE_SPARE|(RAID_ERROR_INDEX|+|5|)|重建阵列所需物理磁盘为指定热备|
|NET_DVR_PD_NOT_FREE|(RAID_ERROR_INDEX|+|6|)|重建阵列所需物理磁盘非空闲|
|NET_DVR_CANNOT_MIG2NEWMODE|(RAID_ERROR_INDEX|+|7|)|不能从当前的阵列类型迁移到新的阵列类型|
|NET_DVR_MIG_PAUSE|(RAID_ERROR_INDEX|+|8|)|迁移操作已暂停|
|NET_DVR_MIG_CANCEL|(RAID_ERROR_INDEX|+|9|)|正在执行的迁移操作已取消|
|NET_DVR_EXIST_VD|(RAID_ERROR_INDEX|+|10)|阵列上阵列上存在虚拟磁盘，无法删除阵列|
|NET_DVR_TARGET_IN_LD_FUNCTIONAL|(RAID_ERROR_INDEX|+|11)|对象物理磁盘为虚拟磁盘组成部分且工作正常|
|NET_DVR_HD_IS_ASSIGNED_ALREADY|(RAID_ERROR_INDEX|+|12)|指定的物理磁盘被分配为虚拟磁盘|
|NET_DVR_INVALID_HD_COUNT|(RAID_ERROR_INDEX|+|13)|物理磁盘数量与指定的RAID等级不匹配|
|NET_DVR_LD_IS_FUNCTIONAL|(RAID_ERROR_INDEX|+|14)|阵列正常，无法重建|
|NET_DVR_BGA_RUNNING|(RAID_ERROR_INDEX|+|15)|存在正在执行的后台任务|
|NET_DVR_LD_NO_ATAPI|(RAID_ERROR_INDEX|+|16)|无法用ATAPI盘创建虚拟磁盘|
|NET_DVR_MIGRATION_NOT_NEED|(RAID_ERROR_INDEX|+|17)|阵列无需迁移|
|NET_DVR_HD_TYPE_MISMATCH|(RAID_ERROR_INDEX|+|18)|物理磁盘不属于同意类型|
|NET_DVR_NO_LD_IN_DG|(RAID_ERROR_INDEX|+|19)|无虚拟磁盘，无法进行此项操作|
|NET_DVR_NO_ROOM_FOR_SPARE|(RAID_ERROR_INDEX|+|20)|磁盘空间过小，无法被指定为热备盘|
|NET_DVR_SPARE_IS_IN_MULTI_DG|(RAID_ERROR_INDEX|+|21)|磁盘已被分配为某阵列热备盘|
|NET_DVR_DG_HAS_MISSING_PD|(RAID_ERROR_INDEX|+|22)|阵列缺少盘|
|NET_DVR_NAME_EMPTY|(RAID_ERROR_INDEX|+|23)|名称为空|
|NET_DVR_INPUT_PARAM|(RAID_ERROR_INDEX|+|24)|输入参数有误|
|NET_DVR_PD_NOT_AVAILABLE|(RAID_ERROR_INDEX|+|25)|物理磁盘不可用|
|NET_DVR_ARRAY_NOT_AVAILABLE|(RAID_ERROR_INDEX|+|26)|阵列不可用|
|NET_DVR_PD_COUNT|(RAID_ERROR_INDEX|+|27)|物理磁盘数不正确|
|NET_DVR_VD_SMALL|(RAID_ERROR_INDEX|+|28)|虚拟磁盘太小|
|NET_DVR_NO_EXIST|(RAID_ERROR_INDEX|+|29)|不存在|
|NET_DVR_NOT_SUPPORT|(RAID_ERROR_INDEX|+|30)|不支持该操作|
|NET_DVR_NOT_FUNCTIONAL|(RAID_ERROR_INDEX|+|31)|阵列状态不是正常状态|
|NET_DVR_DEV_NODE_NOT_FOUND|(RAID_ERROR_INDEX|+|32)|虚拟磁盘设备节点不存在|
|NET_DVR_SLOT_EXCEED|(RAID_ERROR_INDEX|+|33)|槽位达到上限|
|NET_DVR_NO_VD_IN_ARRAY|(RAID_ERROR_INDEX|+|34)|阵列上不存在虚拟磁盘|
|NET_DVR_VD_SLOT_INVALID|(RAID_ERROR_INDEX|+|35)|虚拟磁盘槽位无效|
|NET_DVR_PD_NO_ENOUGH_SPACE|(RAID_ERROR_INDEX|+|36)|所需物理磁盘空间不足|
|NET_DVR_ARRAY_NONFUNCTION|(RAID_ERROR_INDEX|+|37)|只有处于正常状态的阵列才能进行迁移|
|NET_DVR_ARRAY_NO_ENOUGH_SPACE|(RAID_ERROR_INDEX|+|38)|阵列空间不足|
|NET_DVR_STOPPING_SCANNING_ARRAY|(RAID_ERROR_INDEX|+|39)|正在执行安全拔盘或重新扫描|
|NET_DVR_NOT_SUPPORT_16T|(RAID_ERROR_INDEX|+|40)|不支持创建大于16T的阵列|
|NET_DVR_ARRAY_FORMATING|(RAID_ERROR_INDEX|+|41)|正在执行格式化的阵列无法删除|
|NET_DVR_QUICK_SETUP_PD_COUNT|(RAID_ERROR_INDEX|+|42)|一键配置至少需要三块空闲盘|
|NET_DVR_ERROR_DEVICE_NOT_ACTIVATED|250|设备未激活|
|NET_DVR_ERROR_RISK_PASSWORD|251|老SDK接新设备，有风险的密码|
|NET_DVR_ERROR_DEVICE_HAS_ACTIVATED|252|设备已激活|
|智能错误码|||
|VCA_ERROR_INDEX|300|智能错误码索引|
|NET_DVR_ID_ERROR|(VCA_ERROR_INDEX|+|0)|配置ID不合理|
|NET_DVR_POLYGON_ERROR|(VCA_ERROR_INDEX|+|1)|多边形不符合要求|
|NET_DVR_RULE_PARAM_ERROR|(VCA_ERROR_INDEX|+|2)|规则参数不合理|
|NET_DVR_RULE_CFG_CONFLICT|(VCA_ERROR_INDEX|+|3)|配置信息冲突|
|NET_DVR_CALIBRATE_NOT_READY|(VCA_ERROR_INDEX|+|4)|当前没有标定信息|
|NET_DVR_CAMERA_DATA_ERROR|(VCA_ERROR_INDEX|+|5)|摄像机参数不合理|
|NET_DVR_CALIBRATE_DATA_UNFIT|(VCA_ERROR_INDEX|+|6)|长度不够倾斜，不利于标定|
|NET_DVR_CALIBRATE_DATA_CONFLICT|(VCA_ERROR_INDEX|+|7)|标定出错，以为所有点共线或者位置太集中|
|NET_DVR_CALIBRATE_CALC_FAIL|(VCA_ERROR_INDEX|+|8)|摄像机标定参数值计算失败|
|NET_DVR_CALIBRATE_LINE_OUT_RECT|(VCA_ERROR_INDEX|+|9)|输入的样本标定线超出了样本外接矩形框|
|NET_DVR_ENTER_RULE_NOT_READY|(VCA_ERROR_INDEX|+|10)|没有设置进入区域|
|NET_DVR_AID_RULE_NO_INCLUDE_LANE|(VCA_ERROR_INDEX|+|11)|交通事件规则中没有包括车道（特值拥堵和逆行）|
|NET_DVR_LANE_NOT_READY|(VCA_ERROR_INDEX|+|12)|当前没有设置车道|
|NET_DVR_RULE_INCLUDE_TWO_WAY|(VCA_ERROR_INDEX|+|13)|事件规则中包含2种不同方向|
|NET_DVR_LANE_TPS_RULE_CONFLICT|(VCA_ERROR_INDEX|+|14)|车道和数据规则冲突|
|NET_DVR_NOT_SUPPORT_EVENT_TYPE|(VCA_ERROR_INDEX|+|15)|不支持的事件类型|
|NET_DVR_LANE_NO_WAY|(VCA_ERROR_INDEX|+|16)|车道没有方向|
|NET_DVR_SIZE_FILTER_ERROR|(VCA_ERROR_INDEX|+|17)|尺寸过滤框不合理|
|NET_DVR_LIB_FFL_NO_FACE|(VCA_ERROR_INDEX|+|18)|特征点定位时输入的图像没有人脸|
|NET_DVR_LIB_FFL_IMG_TOO_SMALL|(VCA_ERROR_INDEX|+|19)|特征点定位时输入的图像太小|
|NET_DVR_LIB_FD_IMG_NO_FACE|(VCA_ERROR_INDEX|+|20)|单张图像人脸检测时输入的图像没有人脸|
|NET_DVR_LIB_FACE_TOO_SMALL|(VCA_ERROR_INDEX|+|21)|建模时人脸太小|
|NET_DVR_LIB_FACE_QUALITY_TOO_BAD|(VCA_ERROR_INDEX|+|22)|建模时人脸图像质量太差|
|NET_DVR_KEY_PARAM_ERR|(VCA_ERROR_INDEX|+|23)|高级参数设置错误|
|NET_DVR_CALIBRATE_DATA_ERR|(VCA_ERROR_INDEX|+|24)|标定样本数目错误，或数据值错误，或样本点超出地平线|
|NET_DVR_CALIBRATE_DISABLE_FAIL|(VCA_ERROR_INDEX|+|25)|所配置规则不允许取消标定|
|NET_DVR_VCA_LIB_FD_SCALE_OUTRANGE|(VCA_ERROR_INDEX|+|26)|最大过滤框的宽高最小值超过最小过滤框的宽高最大值两倍以上|
|NET_DVR_LIB_FD_REGION_TOO_LARGE|(VCA_ERROR_INDEX|+|27)|当前检测区域范围过大。检测区最大为图像的2/3|
|NET_DVR_TRIAL_OVERDUE|(VCA_ERROR_INDEX|+|28)|试用版评估期已结束|
|NET_DVR_CONFIG_FILE_CONFLICT|(VCA_ERROR_INDEX|+|29)|设备类型与配置文件冲突（加密狗类型与现有分析仪配置不符错误码提示）|
|NET_DVR_FR_FPL_FAIL|(VCA_ERROR_INDEX|+|30)|人脸特征点定位失败|
|NET_DVR_FR_IQA_FAIL|(VCA_ERROR_INDEX|+|31)|人脸评分失败|
|NET_DVR_FR_FEM_FAIL|(VCA_ERROR_INDEX|+|32)|人脸特征提取失败|
|NET_DVR_FPL_DT_CONF_TOO_LOW|(VCA_ERROR_INDEX|+|33)|特征点定位时人脸检测置信度过低|
|NET_DVR_FPL_CONF_TOO_LOW|(VCA_ERROR_INDEX|+|34)|特征点定位置信度过低|
|NET_DVR_E_DATA_SIZE|(VCA_ERROR_INDEX|+|35)|数据长度不匹配|
|NET_DVR_FR_MODEL_VERSION_ERR|(VCA_ERROR_INDEX|+|36)|人脸模型数据中的模型版本错误|
|NET_DVR_FR_FD_FAIL|(VCA_ERROR_INDEX|+|37)|识别库中人脸检测失败|
|NET_DVR_FA_NORMALIZE_ERR|(VCA_ERROR_INDEX|+|38)|人脸归一化出错|
|NET_DVR_DOG_PUSTREAM_NOT_MATCH|(VCA_ERROR_INDEX|+|39)|加密狗与前端取流设备类型不匹配|
|NET_DVR_DEV_PUSTREAM_NOT_MATCH|(VCA_ERROR_INDEX|+|40)|前端取流设备版本不匹配|
|NET_DVR_PUSTREAM_ALREADY_EXISTS|(VCA_ERROR_INDEX|+|41)|设备的其他通道已经添加过该前端设备|
|NET_DVR_SEARCH_CONNECT_FAILED|(VCA_ERROR_INDEX|+|42)|连接检索服务器失败|
|NET_DVR_INSUFFICIENT_DISK_SPACE|(VCA_ERROR_INDEX|+|43)|可存储的硬盘空间不足|
|NET_DVR_DATABASE_CONNECTION_FAILED|(VCA_ERROR_INDEX|+|44)|数据库连接失败|
|NET_DVR_DATABASE_ADM_PW_ERROR|(VCA_ERROR_INDEX|+|45)|数据库用户名、密码错误|
|NET_DVR_DECODE_YUV|(VCA_ERROR_INDEX|+|46)|解码失败|
|RTSP|错误码|||
|NET_DVR_RTSP_ERROR_NOENOUGHPRI|401|无权限：服务器返回401时，转成这个错误码|
|NET_DVR_RTSP_ERROR_ALLOC_RESOURCE|402|分配资源失败|
|NET_DVR_RTSP_ERROR_PARAMETER|403|参数错误|
|NET_DVR_RTSP_ERROR_NO_URL|404|指定的URL地址不存在：服务器返回404时，转成这个错误码|
|NET_DVR_RTSP_ERROR_FORCE_STOP|406|用户中途强行退出|
|NET_DVR_RTSP_GETPORTFAILED|407|rtsp|得到端口错误|
|NET_DVR_RTSP_DESCRIBERROR|410|rtsp|decribe|交互错误|
|NET_DVR_RTSP_DESCRIBESENDTIMEOUT|411|rtsp|decribe|发送超时|
|NET_DVR_RTSP_DESCRIBESENDERROR|412|rtsp|decribe|发送失败|
|NET_DVR_RTSP_DESCRIBERECVTIMEOUT|413|rtsp|decribe|接收超时|
|NET_DVR_RTSP_DESCRIBERECVDATALOST|414|rtsp|decribe|接收数据错误|
|NET_DVR_RTSP_DESCRIBERECVERROR|415|rtsp|decribe|接收失败|
|NET_DVR_RTSP_DESCRIBESERVERERR|416|rtsp|decribe|服务器返回错误状态|
|NET_DVR_RTSP_SETUPERROR|420|rtsp|setup|交互错误|
|NET_DVR_RTSP_SETUPSENDTIMEOUT|421|rtsp|setup|发送超时|
|NET_DVR_RTSP_SETUPSENDERROR|422|rtsp|setup|发送错误|
|NET_DVR_RTSP_SETUPRECVTIMEOUT|423|rtsp|setup|接收超时|
|NET_DVR_RTSP_SETUPRECVDATALOST|424|rtsp|setup|接收数据错误|
|NET_DVR_RTSP_SETUPRECVERROR|425|rtsp|setup|接收失败|
|NET_DVR_RTSP_OVER_MAX_CHAN|426|超过服务器最大连接数，或者服务器资源不足，服务器返回453时，转成这个错误码。|
|NET_DVR_RTSP_SETUPSERVERERR|427|rtsp|setup|服务器返回错误状态|
|NET_DVR_RTSP_PLAYERROR|430|rtsp|play|交互错误|
|NET_DVR_RTSP_PLAYSENDTIMEOUT|431|rtsp|play|发送超时|
|NET_DVR_RTSP_PLAYSENDERROR|432|rtsp|play|发送错误|
|NET_DVR_RTSP_PLAYRECVTIMEOUT|433|rtsp|play|接收超时|
|NET_DVR_RTSP_PLAYRECVDATALOST|434|rtsp|play|接收数据错误|
|NET_DVR_RTSP_PLAYRECVERROR|435|rtsp|play|接收失败|
|NET_DVR_RTSP_PLAYSERVERERR|436|rtsp|play|服务器返回错误状态|
|NET_DVR_RTSP_TEARDOWNERROR|440|rtsp|teardown|交互错误|
|NET_DVR_RTSP_TEARDOWNSENDTIMEOUT|441|rtsp|teardown|发送超时|
|NET_DVR_RTSP_TEARDOWNSENDERROR|442|rtsp|teardown|发送错误|
|NET_DVR_RTSP_TEARDOWNRECVTIMEOUT|443|rtsp|teardown|接收超时|
|NET_DVR_RTSP_TEARDOWNRECVDATALOST|444|rtsp|teardown|接收数据错误|
|NET_DVR_RTSP_TEARDOWNRECVERROR|445|rtsp|teardown|接收失败|
|NET_DVR_RTSP_TEARDOWNSERVERERR|446|rtsp|teardown|服务器返回错误状态|
|播放错误码|||
|NET_PLAYM4_NOERROR|500|no|error|
|NET_PLAYM4_PARA_OVER|501|input|parameter|is|invalid;|
|NET_PLAYM4_ORDER_ERROR|502|The|order|of|the|function|to|be|called|is|error.|
|NET_PLAYM4_TIMER_ERROR|503|Create|multimedia|clock|failed;|
|NET_PLAYM4_DEC_VIDEO_ERROR|504|Decode|video|data|failed.|
|NET_PLAYM4_DEC_AUDIO_ERROR|505|Decode|audio|data|failed.|
|NET_PLAYM4_ALLOC_MEMORY_ERROR|506|Allocate|memory|failed.|
|NET_PLAYM4_OPEN_FILE_ERROR|507|Open|the|file|failed.|
|NET_PLAYM4_CREATE_OBJ_ERROR|508|Create|thread|or|event|failed|
|NET_PLAYM4_CREATE_DDRAW_ERROR|509|Create|DirectDraw|object|failed.|
|NET_PLAYM4_CREATE_OFFSCREEN_ERROR|510|failed|when|creating|off-screen|surface.|
|NET_PLAYM4_BUF_OVER|511|buffer|is|overflow|
|NET_PLAYM4_CREATE_SOUND_ERROR|512|failed|when|creating|audio|device.|
|NET_PLAYM4_SET_VOLUME_ERROR|513|Set|volume|failed|
|NET_PLAYM4_SUPPORT_FILE_ONLY|514|The|function|only|support|play|file.|
|NET_PLAYM4_SUPPORT_STREAM_ONLY|515|The|function|only|support|play|stream.|
|NET_PLAYM4_SYS_NOT_SUPPORT|516|System|not|support.|
|NET_PLAYM4_FILEHEADER_UNKNOWN|517|No|file|header.|
|NET_PLAYM4_VERSION_INCORRECT|518|The|version|of|decoder|and|encoder|is|not|adapted.|
|NET_PALYM4_INIT_DECODER_ERROR|519|Initialize|decoder|failed.|
|NET_PLAYM4_CHECK_FILE_ERROR|520|The|file|data|is|unknown.|
|NET_PLAYM4_INIT_TIMER_ERROR|521|Initialize|multimedia|clock|failed.|
|NET_PLAYM4_BLT_ERROR|522|Blt|failed.|
|NET_PLAYM4_UPDATE_ERROR|523|Update|failed.|
|NET_PLAYM4_OPEN_FILE_ERROR_MULTI|524|openfile|error,|streamtype|is|multi|
|NET_PLAYM4_OPEN_FILE_ERROR_VIDEO|525|openfile|error,|streamtype|is|video|
|NET_PLAYM4_JPEG_COMPRESS_ERROR|526|JPEG|compress|error|
|NET_PLAYM4_EXTRACT_NOT_SUPPORT|527|Don't|support|the|version|of|this|file.|
|NET_PLAYM4_EXTRACT_DATA_ERROR|528|extract|video|data|failed.|
|语音对讲库错误码|||
|NET_AUDIOINTERCOM_OK|600|无错误|
|NET_AUDIOINTECOM_ERR_NOTSUPORT|601|不支持|
|NET_AUDIOINTECOM_ERR_ALLOC_MEMERY|602|内存申请错误|
|NET_AUDIOINTECOM_ERR_PARAMETER|603|参数错误|
|NET_AUDIOINTECOM_ERR_CALL_ORDER|604|调用次序错误|
|NET_AUDIOINTECOM_ERR_FIND_DEVICE|605|未发现设备|
|NET_AUDIOINTECOM_ERR_OPEN_DEVICE|606|不能打开设备诶|
|NET_AUDIOINTECOM_ERR_NO_CONTEXT|607|设备上下文出错|
|NET_AUDIOINTECOM_ERR_NO_WAVFILE|608|WAV文件出错|
|NET_AUDIOINTECOM_ERR_INVALID_TYPE|609|无效的WAV参数类型|
|NET_AUDIOINTECOM_ERR_ENCODE_FAIL|610|编码失败|
|NET_AUDIOINTECOM_ERR_DECODE_FAIL|611|解码失败|
|NET_AUDIOINTECOM_ERR_NO_PLAYBACK|612|播放失败|
|NET_AUDIOINTECOM_ERR_DENOISE_FAIL|613|降噪失败|
|NET_AUDIOINTECOM_ERR_UNKOWN|619|未知错误|
|QOS|错误码|||
|NET_QOS_OK|700|no|error|
|NET_QOS_ERROR|(NET_QOS_OK|-|1)|qos|error|
|NET_QOS_ERR_INVALID_ARGUMENTS|(NET_QOS_OK|-|2)|invalid|arguments|
|NET_QOS_ERR_SESSION_NOT_FOUND|(NET_QOS_OK|-|3)|session|net|found|
|NET_QOS_ERR_LIB_NOT_INITIALIZED|(NET_QOS_OK|-|4)|lib|not|initialized|
|NET_QOS_ERR_OUTOFMEM|(NET_QOS_OK|-|5)|outtofmem|
|NET_QOS_ERR_PACKET_UNKNOW|(NET_QOS_OK|-|10)|packet|unknow|
|NET_QOS_ERR_PACKET_VERSION|(NET_QOS_OK|-|11)|packet|version|error|
|NET_QOS_ERR_PACKET_LENGTH|(NET_QOS_OK|-|12)|packet|length|error|
|NET_QOS_ERR_PACKET_TOO_BIG|(NET_QOS_OK|-|13)|packet|too|big|
|NET_QOS_ERR_SCHEDPARAMS_INVALID_BANDWIDTH|(NET_QOS_OK|-|20)|schedparams|invalid|bandwidth|
|NET_QOS_ERR_SCHEDPARAMS_BAD_FRACTION|(NET_QOS_OK|-|21)|schedparams|bad|fraction|
|NET_QOS_ERR_SCHEDPARAMS_BAD_MINIMUM_INTERVAL|(NET_QOS_OK|-|22)|schedparams|bad|minimum|interval|
|其它错误码|||
|NET_ERROR_TRUNK_LINE|711|子系统已被配成干线|
|NET_ERROR_MIXED_JOINT|712|不能进行混合拼接|
|NET_ERROR_DISPLAY_SWITCH|713|不能进行显示通道切换|
|NET_ERROR_USED_BY_BIG_SCREEN|714|解码资源被大屏占用|
|NET_ERROR_USE_OTHER_DEC_RESOURCE|715|不能使用其他解码子系统资源|
|NET_ERROR_DISP_MODE_SWITCH|716|显示通道显示状态切换中|
|NET_ERROR_SCENE_USING|717|场景正在使用|
|NET_ERR_NO_ENOUGH_DEC_RESOURCE|718|解码资源不足|
|NET_ERR_NO_ENOUGH_FREE_SHOW_RESOURCE|719|畅显资源不足|
|NET_ERR_NO_ENOUGH_VIDEO_MEMORY|720|显存资源不足|
|NET_ERR_MAX_VIDEO_NUM|721|一拖多资源不足|
|NET_ERR_WIN_COVER_FREE_SHOW_AND_NORMAL|722|窗口跨越了畅显输出口和非畅显输出口|
|NET_ERR_FREE_SHOW_WIN_SPLIT|723|畅显窗口不支持分屏|
|NET_ERR_INAPPROPRIATE_WIN_FREE_SHOW|724|不是输出口整数倍的窗口不支持开启畅显|
|NET_DVR_TRANSPARENT_WIN_NOT_SUPPORT_SPLIT|725|开启透明度的窗口不支持分屏|
|NET_DVR_SPLIT_WIN_NOT_SUPPORT_TRANSPARENT|726|开启多分屏的窗口不支持透明度设置|
|NET_ERR_MAX_LOGO_NUM|727|logo数达到上限|
|NET_ERR_MAX_WIN_LOOP_NUM|728|轮巡窗口数达到上限|
|NET_DVR_DEV_NET_OVERFLOW|800|网络流量超过设备能力上限|
|NET_DVR_STATUS_RECORDFILE_WRITING_NOT_LOCK|801|录像文件在录像，无法被锁定|
|NET_DVR_STATUS_CANT_FORMAT_LITTLE_DISK|802|由于硬盘太小无法格式化|
|N+1|错误码|||
|NET_SDK_ERR_REMOTE_DISCONNECT|803|远端无法连接|
|NET_SDK_ERR_RD_ADD_RD|804|备机不能添加备机|
|NET_SDK_ERR_BACKUP_DISK_EXCEPT|805|备份盘异常|
|NET_SDK_ERR_RD_LIMIT|806|备机数已达上限|
|NET_SDK_ERR_ADDED_RD_IS_WD|807|添加的备机是工作机|
|NET_SDK_ERR_ADD_ORDER_WRONG|808|添加顺序出错，比如没有被工作机添加为备机，就添加工作机|
|NET_SDK_ERR_WD_ADD_WD|809|工作机不能添加工作机|
|NET_SDK_ERR_WD_SERVICE_EXCETP|810|工作机CVR服务异常|
|NET_SDK_ERR_RD_SERVICE_EXCETP|811|备机CVR服务异常|
|NET_SDK_ERR_ADDED_WD_IS_RD|812|添加的工作机是备机|
|NET_SDK_ERR_PERFORMANCE_LIMIT|813|性能达到上限|
|NET_SDK_ERR_ADDED_DEVICE_EXIST|814|添加的设备已经存在|
|审讯机错误码|||
|NET_SDK_ERR_INQUEST_RESUMING|815|审讯恢复中|
|NET_SDK_ERR_RECORD_BACKUPING|816|审讯备份中|
|NET_SDK_ERR_DISK_PLAYING|817|光盘回放中|
|NET_SDK_ERR_INQUEST_STARTED|818|审讯已开启|
|NET_SDK_ERR_LOCAL_OPERATING|819|本地操作进行中|
|NET_SDK_ERR_INQUEST_NOT_START|820|审讯未开启|
|Netra|3.1.0|错误码|||
|NET_SDK_ERR_CHAN_AUDIO_BIND|821|通道未绑定或绑定语音对讲失败|
|云存储错误码|||
|NET_DVR_N_PLUS_ONE_MODE|822|设备当前处于N+1模式|
|NET_DVR_CLOUD_STORAGE_OPENED|823|云存储模式已开启|
|庭审主机错误码|||
|NET_SDK_ERR_IR_PORT_ERROR|830|红外输出口错误|
|NET_SDK_ERR_IR_CMD_ERROR|831|红外输出口的命令号错误|
|多屏控制器错误码（900-950）|||
|NET_ERR_WINCHAN_IDX|901|开窗通道号错误|
|NET_ERR_WIN_LAYER|902|窗口层数错误，单个屏幕上最多覆盖的窗口层数|
|NET_ERR_WIN_BLK_NUM|903|窗口的块数错误，单个窗口可覆盖的屏幕个数|
|NET_ERR_OUTPUT_RESOLUTION|904|输出分辨率错误|
|NET_ERR_LAYOUT|905|布局号错误|
|NET_ERR_INPUT_RESOLUTION|906|输入分辨率不支持|
|NET_ERR_SUBDEVICE_OFFLINE|907|子设备不在线|
|NET_ERR_NO_DECODE_CHAN|908|没有空闲解码通道|
|NET_ERR_MAX_WINDOW_ABILITY|909|开窗能力上限,|分布式多屏控制器中解码子设备能力上限或者显示处理器能力上限导致|
|NET_ERR_ORDER_ERROR|910|调用顺序有误|
|NET_ERR_PLAYING_PLAN|911|正在执行预案|
|NET_ERR_DECODER_USED|912|解码板正在使用|
|NET_ERR_OUTPUT_BOARD_DATA_OVERFLOW|913|输出板数据量超限|
|NET_ERR_SAME_USER_NAME|914|用户名相同|
|NET_ERR_INVALID_USER_NAME|915|无效用户名|
|NET_ERR_MATRIX_USING|916|输入矩阵正在使用|
|NET_ERR_DIFFERENT_CHAN_TYPE|917|通道类型不同（矩阵输出通道和控制器的输入为不同的类型）|
|NET_ERR_INPUT_CHAN_BINDED|918|输入通道已经被其他矩阵绑定|
|NET_ERR_BINDED_OUTPUT_CHAN_OVERFLOW|919|正在使用的矩阵输出通道个数超过矩阵与控制器绑定的通道个数|
|NET_ERR_MAX_SIGNAL_NUM|920|输入信号源个数达到上限|
|NET_ERR_INPUT_CHAN_USING|921|输入通道正在使用|
|NET_ERR_MANAGER_LOGON|922|管理员已经登陆，操作失败|
|NET_ERR_USERALREADY_LOGON|923|该用户已经登陆，操作失败|
|NET_ERR_LAYOUT_INIT|924|布局正在初始化，操作失败|
|NET_ERR_BASEMAP_SIZE_NOT_MATCH|925|底图大小不符|
|NET_ERR_WINDOW_OPERATING|926|窗口正在执行其他操作，本次操作失败|
|NET_ERR_SIGNAL_UPLIMIT|927|信号源开窗个数达到上限|
|NET_ERR_SIGNAL_MAX_ENLARGE_TIMES|928|信号源放大倍数超限|
|NET_ERR_ONE_SIGNAL_MULTI_CROSS|929|单个信号源不能多次跨屏|
|NET_ERR_ULTRA_HD_SIGNAL_MULTI_WIN|930|超高清信号源不能重复开窗|
|NET_ERR_MAX_VIRTUAL_LED_WIDTH|931|虚拟LED宽度大于限制值|
|NET_ERR_MAX_VIRTUAL_LED_WORD_LEN|932|虚拟LED字符数大于限制值|
|NET_ERR_SINGLE_OUTPUTPARAM_CONFIG|933|不支持单个显示输出参数设置|
|NET_ERR_MULTI_WIN_BE_COVER|934|多分屏窗口被覆盖|
|NET_ERR_WIN_NOT_EXIST|935|窗口不存在|
|NET_ERR_WIN_MAX_SIGNALSOURCE|936|窗口信号源数超过限制值|
|NET_ERR_MULTI_WIN_MOVE|937|对多分屏窗口移动|
|NET_ERR_MULTI_WIN_YPBPR_SDI|938|YPBPR|和SDI信号源不支持9/16分屏|
|NET_ERR_DIFF_TYPE_OUTPUT_MIXUSE|939|不同类型输出板混插|
|NET_ERR_SPLIT_WIN_CROSS|940|对跨屏窗口分屏|
|NET_ERR_SPLIT_WIN_NOT_FULL_SCREEN|941|对未满屏窗口分屏|
|NET_ERR_SPLIT_WIN_MANY_WIN|942|对单个输出口上有多个窗口的窗口分屏|
|NET_ERR_LED_RESOLUTION|946|LED|分辨率大于输出分辨率|
|解码器错误码（951-999）|||
|NET_ERR_MAX_WIN_OVERLAP|951|达到最大窗口重叠数|
|NET_ERR_STREAMID_CHAN_BOTH_VALID|952|stream|ID和通道号同时有效|
|NET_ERR_NO_ZERO_CHAN|953|设备无零通道|
|NEED_RECONNECT|955|需要重定向（转码子系统使用）|
|NET_ERR_NO_STREAM_ID|956|流ID不存在|
|NET_DVR_TRANS_NOT_START|957|转码未启动|
|NET_ERR_MAXNUM_STREAM_ID|958|流ID数达到上限|
|NET_ERR_WORKMODE_MISMATCH|959|工作模式不匹配|
|NET_ERR_MODE_IS_USING|960|已工作在当前模式|
|NET_ERR_DEV_PROGRESSING|961|设备正在处理中|
|NET_ERR_PASSIVE_TRANSCODING|962|正在被动转码|
|能力集解析库错误码|||
|XML_ABILITY_NOTSUPPORT|1000|不支持能力节点获取|
|XML_ANALYZE_NOENOUGH_BUF|1001|输出内存不足|
|XML_ANALYZE_FIND_LOCALXML_ERROR|1002|无法找到对应的本地xml|
|XML_ANALYZE_LOAD_LOCALXML_ERROR|1003|加载本地xml出错|
|XML_NANLYZE_DVR_DATA_FORMAT_ERROR|1004|设备能力数据格式错误|
|XML_ANALYZE_TYPE_ERROR|1005|能力集类型错误|
|XML_ANALYZE_XML_NODE_ERROR|1006|XML能力节点格式错误|
|XML_INPUT_PARAM_ERROR|1007|输入的能力XML节点值错误|
|民用错误码（1100～1200）|||
|NET_ERR_PLT_USERID|1100|验证平台userid错误|
|NET_ERR_TRANS_CHAN_START|1101|透明通道已打开，当前操作无法完成|
|NET_ERR_DEV_UPGRADING|1102|设备正在升级|
|NET_ERR_MISMATCH_UPGRADE_PACK_TYPE|1103|升级包类型不匹配|
|NET_ERR_DEV_FORMATTING|1104|设备正在格式化|
|NET_ERR_MISMATCH_UPGRADE_PACK_VERSION|1105|升级包版本不匹配|
|报警设备错误码（1200~1300）|||
|NET_ERR_SEARCHING_MODULE|1201|正在搜索外接模块|
|NET_ERR_REGISTERING_MODULE|1202|正在注册外接模块|
|NET_ERR_GETTING_ZONES|1203|正在获取防区参数|
|NET_ERR_GETTING_TRIGGERS|1204|正在获取触发器|
|NET_ERR_ARMED_STATUS|1205|系统处于布防状态|
|NET_ERR_PROGRAM_MODE_STATUS|1206|系统处于编程模式|
|NET_ERR_WALK_TEST_MODE_STATUS|1207|系统处于步测模式|
|NET_ERR_BYPASS_STATUS|1208|旁路状态|
|NET_ERR_DISABLED_MODULE_STATUS|1209|功能未使能|
|NET_ERR_NOT_SUPPORT_OPERATE_ZONE|1210|防区不支持该操作|
|NET_ERR_NOT_SUPPORT_MOD_MODULE_ADDR|1211|模块地址不能被修改|
|NET_ERR_UNREGISTERED_MODULE|1212|模块未注册|
|NET_ERR_PUBLIC_SUBSYSTEM_ASSOCIATE_SELF|1213|公共子系统关联自身|
|NET_ERR_EXCEEDS_ASSOCIATE_SUBSYSTEM_NUM|1214|超过公共子系统最大关联个数|
|NET_ERR_BE_ASSOCIATED_BY_PUBLIC_SUBSYSTEM|1215|子系统被其他公共子系统关联|
|NET_ERR_ZONE_FAULT_STATUS|1216|防区处于故障状态|
|NET_ERR_SAME_EVENT_TYPE|1217|事件触发报警输出开启和事件触发报警输出关闭中有相同事件类型|
|NET_ERR_ZONE_ALARM_STATUS|1218|防区处于报警状态|
|NET_ERR_EXPANSION_BUS_SHORT_CIRCUIT|1219|扩展总线短路|
|NET_ERR_PWD_CONFLICT|1220|密码冲突|
|抓拍机错误码（1400-1499）|||
|NET_DVR_ERR_LANENUM_EXCEED|1400|车道数超出能力|
|NET_DVR_ERR_PRAREA_EXCEED|1401|牌识区域过大|
|NET_DVR_ERR_LIGHT_PARAM|1402|信号灯接入参数错误|
|NET_DVR_ERR_LANE_LINE_INVALID|1403|车道线配置错误|
|NET_DVR_ERR_STOP_LINE_INVALID|1404|停止线配置错误|
|NET_DVR_ERR_LEFTORRIGHT_LINE_INVALID|1405|左/右转分界线配置错误|
|NET_DVR_ERR_LANE_NO_REPEAT|1406|叠加车道号重复|
|NET_DVR_ERR_PRAREA_INVALID|1407|牌识多边形不符合要求|
|NET_DVR_ERR_LIGHT_NUM_EXCEED|1408|视频检测交通灯信号灯数目超出最大值|
|NET_DVR_ERR_SUBLIGHT_NUM_INVALID|1409|视频检测交通灯信号灯子灯数目不合法|
|NET_DVR_ERR_LIGHT_AREASIZE_INVALID|1410|视频检测交通灯输入信号灯框大小不合法|
|NET_DVR_ERR_LIGHT_COLOR_INVALID|1411|视频检测交通灯输入信号灯颜色不合法|
|NET_DVR_ERR_LIGHT_DIRECTION_INVALID|1412|视频检测交通灯输入灯方向属性不合法|
|NET_DVR_ERR_LACK_IOABLITY|1413|IO口实际支持的能力不足|
|NET_DVR_ERR_FTP_PORT|1414|FTP端口号非法（端口号重复或者异常）|
|NET_DVR_ERR_FTP_CATALOGUE|1415|FTP目录名非法（启用多级目录，多级目录传值为空）|
|NET_DVR_ERR_FTP_UPLOAD_TYPE|1416|FTP上传类型非法（单ftp只支持全部/双ftp只支持卡口和违章）|
|NET_DVR_ERR_FLASH_PARAM_WRITE|1417|配置参数时写FLASH失败|
|NET_DVR_ERR_FLASH_PARAM_READ|1418|配置参数时读FLASH失败|
|NET_DVR_ERR_PICNAME_DELIMITER|1419|FTP图片命名分隔符非法|
|NET_DVR_ERR_PICNAME_ITEM|1420|FTP图片命名项非法（例如|分隔符）|
|NET_DVR_ERR_PLATE_RECOGNIZE_TYPE|1421|牌识区域类型非法|（矩形和多边形有效性校验）|
|NET_DVR_ERR_CAPTURE_TIMES|1422|抓拍次数非法|（有效值是0～5）|
|NET_DVR_ERR_LOOP_DISTANCE|1423|线圈距离非法|（有效值是0～2000ms）|
|NET_DVR_ERR_LOOP_INPUT_STATUS|1424|线圈输入状态非法|（有效值）|
|NET_DVR_ERR_RELATE_IO_CONFLICT|1425|测速组IO关联冲突|
|NET_DVR_ERR_INTERVAL_TIME|1426|连拍间隔时间非法|（0～6000ms）|
|NET_DVR_ERR_SIGN_SPEED|1427|标志限速值非法（大车标志限速不能大于小车标志限速|）|
|NET_DVR_ERR_PIC_FLIP|1428|图像配置翻转|（配置交互影响）|
|NET_DVR_ERR_RELATE_LANE_NUMBER|1429|关联车道数错误|(重复|有效值校验1～99)|
|NET_DVR_ERR_TRIGGER_MODE|1430|配置抓拍机触发模式非法|
|NET_DVR_ERR_DELAY_TIME|1431|触发延时时间错误(2000ms)|
|NET_DVR_ERR_EXCEED_RS485_COUNT|1432|超过最大485个数限制|
|NET_DVR_ERR_RADAR_TYPE|1433|雷达类型错误|
|NET_DVR_ERR_RADAR_ANGLE|1434|雷达角度错误|
|NET_DVR_ERR_RADAR_SPEED_VALID_TIME|1435|雷达有效时间错误|
|NET_DVR_ERR_RADAR_LINE_CORRECT|1436|雷达线性矫正参数错误|
|NET_DVR_ERR_RADAR_CONST_CORRECT|1437|雷达常量矫正参数错误|
|NET_DVR_ERR_RECORD_PARAM|1438|录像参数无效（预录时间不超过10s）|
|NET_DVR_ERR_LIGHT_WITHOUT_COLOR_AND_DIRECTION|1439|视频检测信号灯配置信号灯个数，但是没有勾选信号灯方向和颜色的|
|NET_DVR_ERR_LIGHT_WITHOUT_DETECTION_REGION|1440|视频检测信号灯配置信号灯个数，但是没有画检测区域|
|NET_DVR_ERR_RECOGNIZE_PROVINCE_PARAM|1441|牌识参数省份参数的合法性|
|NET_DVR_ERR_SPEED_TIMEOUT|1442|IO测速超时时间非法（有效值大于0）|
|NET_DVR_ERR_NTP_TIMEZONE|1443|ntp时区参数错误|
|NET_DVR_ERR_NTP_INTERVAL_TIME|1444|ntp校时间隔错误|
|NET_DVR_ERR_NETWORK_CARD_NUM|1445|可配置网卡数目错误|
|NET_DVR_ERR_DEFAULT_ROUTE|1446|默认路由错误|
|NET_DVR_ERR_BONDING_WORK_MODE|1447|bonding网卡工作模式错误|
|NET_DVR_ERR_SLAVE_CARD|1448|slave网卡错误|
|NET_DVR_ERR_PRIMARY_CARD|1449|Primary网卡错误|
|NET_DVR_ERR_DHCP_PPOE_WORK|1450|dhcp和pppoE不能同时启动|
|NET_DVR_ERR_NET_INTERFACE|1451|网络接口错误|
|NET_DVR_ERR_MTU|1452|MTU错误|
|NET_DVR_ERR_NETMASK|1453|子网掩码错误|
|NET_DVR_ERR_IP_INVALID|1454|IP地址不合法|
|NET_DVR_ERR_MULTICAST_IP_INVALID|1455|多播地址不合法|
|NET_DVR_ERR_GATEWAY_INVALID|1456|网关不合法|
|NET_DVR_ERR_DNS_INVALID|1457|DNS不合法|
|NET_DVR_ERR_ALARMHOST_IP_INVALID|1458|告警主机地址不合法|
|NET_DVR_ERR_IP_CONFLICT|1459|IP冲突|
|NET_DVR_ERR_NETWORK_SEGMENT|1460|IP不支持同网段|
|NET_DVR_ERR_NETPORT|1461|端口错误|
|NET_DVR_ERR_PPPOE_NOSUPPORT|1462|PPPOE不支持|
|NET_DVR_ERR_DOMAINNAME_NOSUPPORT|1463|域名不支持|
|NET_DVR_ERR_NO_SPEED|1464|未启用测速功能|
|NET_DVR_ERR_IOSTATUS_INVALID|1465|IO状态错误|
|NET_DVR_ERR_BURST_INTERVAL_INVALID|1466|连拍间隔非法|
|NET_DVR_ERR_RESERVE_MODE|1467|备用模式错误|
|NET_DVR_ERR_LANE_NO|1468|叠加车道号错误|
|NET_DVR_ERR_COIL_AREA_TYPE|1469|线圈区域类型错误|
|NET_DVR_ERR_TRIGGER_AREA_PARAM|1470|触发区域参数错误|
|NET_DVR_ERR_SPEED_LIMIT_PARAM|1471|违章限速参数错误|
|NET_DVR_ERR_LANE_PROTOCOL_TYPE|1472|车道关联协议类型错误|
|NET_DVR_ERR_INTERVAL_TYPE|1473|连拍间隔类型非法|
|NET_DVR_ERR_INTERVAL_DISTANCE|1474|连拍间隔距离非法|
|NET_DVR_ERR_RS485_ASSOCIATE_DEVTYPE|1475|RS485关联类型非法|
|NET_DVR_ERR_RS485_ASSOCIATE_LANENO|1476|RS485关联车道号非法|
|NET_DVR_ERR_LANENO_ASSOCIATE_MULTIRS485|1477|车道号关联多个RS485口|
|NET_DVR_ERR_LIGHT_DETECTION_REGION|1478|视频检测信号灯配置信号灯个数，但是检测区域宽或高为0|
|NET_DVR_ERR_DN2D_NOSUPPORT|1479|不支持抓拍帧2D降噪|
|NET_DVR_ERR_IRISMODE_NOSUPPORT|1480|不支持的镜头类型|
|NET_DVR_ERR_WB_NOSUPPORT|1481|不支持的白平衡模式|
|NET_DVR_ERR_IO_EFFECTIVENESS|1482|IO口的有效性|
|NET_DVR_ERR_LIGHTNO_MAX|1483|信号灯检测器接入红/黄灯超限(16)|
|NET_DVR_ERR_LIGHTNO_CONFLICT|1484|信号灯检测器接入红/黄灯冲突|
|NET_DVR_ERR_CANCEL_LINE|1485|直行触发线|
|NET_DVR_ERR_STOP_LINE|1486|待行区停止线|
|NET_DVR_ERR_RUSH_REDLIGHT_LINE|1487|闯红灯触发线|
|NET_DVR_ERR_IOOUTNO_MAX|1488|IO输出口编号越界|
|NET_DVR_ERR_IOOUTNO_AHEADTIME_MAX|1489|IO输出口提前时间超限|
|NET_DVR_ERR_IOOUTNO_IOWORKTIME|1490|IO输出口有效持续时间超限|
|NET_DVR_ERR_IOOUTNO_FREQMULTI|1491|IO输出口脉冲模式下倍频出错|
|NET_DVR_ERR_IOOUTNO_DUTYRATE|1492|IO输出口脉冲模式下占空比出错|
|NET_DVR_ERR_VIDEO_WITH_EXPOSURE|1493|以曝闪起效，工作方式不支持视频|
|NET_DVR_ERR_PLATE_BRIGHTNESS_WITHOUT_FLASHDET|1494|车牌亮度自动使能闪光灯仅在车牌亮度补偿模式下起效|
|NET_DVR_ERR_RECOGNIZE_TYPE_PARAM|1495|识别类型非法|车牌识别参数（如大车、小车、背向、正向、车标识别等）|
|NET_DVR_ERR_PALTE_RECOGNIZE_AREA_PARAM|1496|牌识参数非法|牌识区域配置时判断出错|
|NET_DVR_ERR_PORT_CONFLICT|1497|端口有冲突|
|NET_DVR_ERR_LOOP_IP|1498|IP不能设置为回环地址|
|NET_DVR_ERR_DRIVELINE_SENSITIVE|1499|压线灵敏度出错(视频电警模式下)|
|VQD错误码（1500～1550）|||
|NET_ERR_VQD_TIME_CONFLICT|1500|VQD诊断时间段冲突|
|NET_ERR_VQD_PLAN_NO_EXIST|1501|VQD诊断计划不存在|
|NET_ERR_VQD_CHAN_NO_EXIST|1502|VQD监控点不存在|
|NET_ERR_VQD_CHAN_MAX|1503|VQD计划数已达上限|
|NET_ERR_VQD_TASK_MAX|1504|VQD任务数已达上限|
|抓拍机错误码新增扩展(1600~1900)|||
|NET_DVR_ERR_EXCEED_MAX_CAPTURE_TIMES|1600|抓拍模式为频闪时最大抓拍张数为2张(IVT模式下)|
|NET_DVR_ERR_REDAR_TYPE_CONFLICT|1601|相同485口关联雷达类型冲突|
|NET_DVR_ERR_LICENSE_PLATE_NULL|1602|车牌号为空|
|NET_DVR_ERR_WRITE_DATABASE|1603|写入数据库失败|
|NET_DVR_ERR_LICENSE_EFFECTIVE_TIME|1604|车牌有效时间错误|
|门禁主机错误码|||
|NET_ERR_TIME_OVERLAP|1900|时间段重叠|
|NET_ERR_HOLIDAY_PLAN_OVERLAP|1901|假日计划重叠|
|NET_ERR_CARDNO_NOT_SORT|1902|卡号未排序|
|NET_ERR_CARDNO_NOT_EXIST|1903|卡号不存在|
|NET_ERR_ILLEGAL_CARDNO|1904|卡号错误|
|NET_ERR_ZONE_ALARM|1905|防区处于布防状态(参数修改不允许)|
|NET_ERR_ZONE_OPERATION_NOT_SUPPORT|1906|防区不支持该操作|
|NET_ERR_INTERLOCK_ANTI_CONFLICT|1907|多门互锁和反潜回同时配置错误|
|NET_ERR_DEVICE_CARD_FULL|1908|卡已满（卡达到10W后返回）|
|民用错误码|(5001~6000)|||
|NET_DVR_EZVIZ_NO_ENOUGH_DATA|5001|接收的数据长度不够|
|NET_DVR_EZVIZ_SDK_ERROR|5002|民用通信库加载失败|
|NET_DVR_EZVIZ_GENERAL_UNKNOW_ERROR|5003|未知错误|
|NET_DVR_EZVIZ_GENERAL_SERIAL_NOT_FOR_CIVIL|5004|序列号解析失败|
|NET_DVR_EZVIZ_GENERAL_SERIAL_FORBIDDEN|5005|序列号被禁止|
|NET_DVR_EZVIZ_GENERAL_SERIAL_DUPLICATE|5006|序列号重复|
|NET_DVR_EZVIZ_GENERAL_SERIAL_FLUSHED_IN_A_SECOND|5007|相同序列号短时间内大量请求|
|NET_DVR_EZVIZ_GENERAL_SERIAL_NO_LONGER_SUPPORTED|5008|序列号不再支持|
|NET_DVR_EZVIZ_PLATFORM_CLIENT_REQUEST_NO_PU_FOUNDED|5009|请求的设备不在线|
|NET_DVR_EZVIZ_PLAYFORM_CLIENT_REQUEST_REFUSED_TO_PROTECT_PU|5010|为了保护设备，拒绝请求|
|NET_DVR_EZVIZ_PLAYFORM_CLIENT_REQUEST_PU_LIMIT_REACHED|5011|设备达到链接的客户端上限|
|NET_DVR_EZVIZ_PLAYFORM_CLIENT_TEARDOWN_PU_CONNECTION|5012|服务器要求客户端断开与设备的连接|
|NET_DVR_EZVIZ_PLAYFORM_CLIENT_VERIFY_SESSION_ERROR|5013|设备拒绝平台发送的客户端连接请求|
