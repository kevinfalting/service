// Code generated by 'ccgo stdio/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o stdio/stdio_windows_386.go -pkgname stdio', DO NOT EDIT.

package stdio

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	BUFSIZ                                                 = 512    // stdio.h:17:1:
	DUMMYSTRUCTNAME                                        = 0      // _mingw.h:519:1:
	DUMMYSTRUCTNAME1                                       = 0      // _mingw.h:520:1:
	DUMMYSTRUCTNAME2                                       = 0      // _mingw.h:521:1:
	DUMMYSTRUCTNAME3                                       = 0      // _mingw.h:522:1:
	DUMMYSTRUCTNAME4                                       = 0      // _mingw.h:523:1:
	DUMMYSTRUCTNAME5                                       = 0      // _mingw.h:524:1:
	DUMMYUNIONNAME                                         = 0      // _mingw.h:497:1:
	DUMMYUNIONNAME1                                        = 0      // _mingw.h:498:1:
	DUMMYUNIONNAME2                                        = 0      // _mingw.h:499:1:
	DUMMYUNIONNAME3                                        = 0      // _mingw.h:500:1:
	DUMMYUNIONNAME4                                        = 0      // _mingw.h:501:1:
	DUMMYUNIONNAME5                                        = 0      // _mingw.h:502:1:
	DUMMYUNIONNAME6                                        = 0      // _mingw.h:503:1:
	DUMMYUNIONNAME7                                        = 0      // _mingw.h:504:1:
	DUMMYUNIONNAME8                                        = 0      // _mingw.h:505:1:
	DUMMYUNIONNAME9                                        = 0      // _mingw.h:506:1:
	EOF                                                    = -1     // stdio.h:21:1:
	FILENAME_MAX                                           = 260    // stdio.h:61:1:
	FOPEN_MAX                                              = 20     // stdio.h:62:1:
	MINGW_DDK_H                                            = 0      // _mingw_ddk.h:2:1:
	MINGW_HAS_DDK_H                                        = 1      // _mingw_ddk.h:4:1:
	MINGW_HAS_SECURE_API                                   = 1      // _mingw.h:602:1:
	MINGW_SDK_INIT                                         = 0      // _mingw.h:598:1:
	SEEK_CUR                                               = 1      // stdio.h:53:1:
	SEEK_END                                               = 2      // stdio.h:54:1:
	SEEK_SET                                               = 0      // stdio.h:55:1:
	STDERR_FILENO                                          = 2      // stdio.h:59:1:
	STDIN_FILENO                                           = 0      // stdio.h:57:1:
	STDOUT_FILENO                                          = 1      // stdio.h:58:1:
	SYS_OPEN                                               = 20     // stdio.h:1486:1:
	TMP_MAX                                                = 32767  // stdio.h:64:1:
	TMP_MAX_S                                              = 32767  // stdio_s.h:27:3:
	UNALIGNED                                              = 0      // _mingw.h:384:1:
	USE___UUIDOF                                           = 0      // _mingw.h:77:1:
	WIN32                                                  = 1      // <predefined>:258:1:
	WINNT                                                  = 1      // <predefined>:306:1:
	X_AGLOBAL                                              = 0      // _mingw.h:346:1:
	X_ANONYMOUS_STRUCT                                     = 0      // _mingw.h:474:1:
	X_ANONYMOUS_UNION                                      = 0      // _mingw.h:473:1:
	X_ARGMAX                                               = 100    // _mingw.h:402:1:
	X_CONST_RETURN                                         = 0      // _mingw.h:377:1:
	X_CRTNOALIAS                                           = 0      // corecrt.h:29:1:
	X_CRTRESTRICT                                          = 0      // corecrt.h:33:1:
	X_CRT_ALTERNATIVE_IMPORTED                             = 0      // _mingw.h:313:1:
	X_CRT_DIRECTORY_DEFINED                                = 0      // stdio.h:728:1:
	X_CRT_INTERNAL_LOCAL_PRINTF_OPTIONS                    = 4      // corecrt_stdio_config.h:23:1:
	X_CRT_INTERNAL_LOCAL_SCANF_OPTIONS                     = 2      // corecrt_stdio_config.h:27:1:
	X_CRT_INTERNAL_PRINTF_LEGACY_MSVCRT_COMPATIBILITY      = 0x0008 // corecrt_stdio_config.h:15:1:
	X_CRT_INTERNAL_PRINTF_LEGACY_THREE_DIGIT_EXPONENTS     = 0x0010 // corecrt_stdio_config.h:16:1:
	X_CRT_INTERNAL_PRINTF_LEGACY_VSPRINTF_NULL_TERMINATION = 0x0001 // corecrt_stdio_config.h:12:1:
	X_CRT_INTERNAL_PRINTF_LEGACY_WIDE_SPECIFIERS           = 0x0004 // corecrt_stdio_config.h:14:1:
	X_CRT_INTERNAL_PRINTF_STANDARD_SNPRINTF_BEHAVIOR       = 0x0002 // corecrt_stdio_config.h:13:1:
	X_CRT_INTERNAL_SCANF_LEGACY_MSVCRT_COMPATIBILITY       = 0x0004 // corecrt_stdio_config.h:20:1:
	X_CRT_INTERNAL_SCANF_LEGACY_WIDE_SPECIFIERS            = 0x0002 // corecrt_stdio_config.h:19:1:
	X_CRT_INTERNAL_SCANF_SECURECRT                         = 0x0001 // corecrt_stdio_config.h:18:1:
	X_CRT_MANAGED_HEAP_DEPRECATE                           = 0      // _mingw.h:361:1:
	X_CRT_PACKING                                          = 8      // corecrt.h:14:1:
	X_CRT_PERROR_DEFINED                                   = 0      // stdio.h:712:1:
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES                 = 0      // _mingw_secapi.h:34:1:
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES_MEMORY          = 0      // _mingw_secapi.h:35:1:
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES               = 0      // _mingw_secapi.h:36:1:
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT         = 0      // _mingw_secapi.h:37:1:
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_MEMORY        = 0      // _mingw_secapi.h:38:1:
	X_CRT_USE_WINAPI_FAMILY_DESKTOP_APP                    = 0      // corecrt.h:501:1:
	X_CRT_WPERROR_DEFINED                                  = 0      // stdio.h:1423:1:
	X_DLL                                                  = 0      // _mingw.h:326:1:
	X_ERRCODE_DEFINED                                      = 0      // corecrt.h:117:1:
	X_FILE_DEFINED                                         = 0      // stdio.h:35:1:
	X_FILE_OFFSET_BITS                                     = 64     // <builtin>:25:1:
	X_FILE_OFFSET_BITS_SET_FSEEKO                          = 0      // stdio.h:692:1:
	X_FILE_OFFSET_BITS_SET_FTELLO                          = 0      // stdio.h:699:1:
	X_FILE_OFFSET_BITS_SET_OFFT                            = 0      // _mingw_off_t.h:21:1:
	X_FPOS_T_DEFINED                                       = 0      // stdio.h:98:1:
	X_ILP32                                                = 1      // <predefined>:211:1:
	X_INC_CORECRT                                          = 0      // corecrt.h:8:1:
	X_INC_CRTDEFS_MACRO                                    = 0      // _mingw_mac.h:8:1:
	X_INC_MINGW_SECAPI                                     = 0      // _mingw_secapi.h:8:1:
	X_INC_STDIO                                            = 0      // stdio.h:7:1:
	X_INC_STDIO_S                                          = 0      // stdio_s.h:7:1:
	X_INC_SWPRINTF_INL                                     = 0      // swprintf.inl:8:1:
	X_INC_VADEFS                                           = 0      // vadefs.h:7:1:
	X_INC__MINGW_H                                         = 0      // _mingw.h:8:1:
	X_INT128_DEFINED                                       = 0      // _mingw.h:237:1:
	X_INTEGRAL_MAX_BITS                                    = 64     // <predefined>:320:1:
	X_INTPTR_T_DEFINED                                     = 0      // corecrt.h:62:1:
	X_IOB_ENTRIES                                          = 20     // stdio.h:20:1:
	X_IOEOF                                                = 0x0010 // stdio.h:127:1:
	X_IOERR                                                = 0x0020 // stdio.h:128:1:
	X_IOFBF                                                = 0x0000 // stdio.h:122:1:
	X_IOLBF                                                = 0x0040 // stdio.h:123:1:
	X_IOMYBUF                                              = 0x0008 // stdio.h:126:1:
	X_IONBF                                                = 0x0004 // stdio.h:124:1:
	X_IOREAD                                               = 0x0001 // stdio.h:119:1:
	X_IORW                                                 = 0x0080 // stdio.h:130:1:
	X_IOSTRG                                               = 0x0040 // stdio.h:129:1:
	X_IOWRT                                                = 0x0002 // stdio.h:120:1:
	X_MT                                                   = 0      // _mingw.h:330:1:
	X_M_IX86                                               = 600    // _mingw_mac.h:54:1:
	X_NFILE                                                = 512    // stdio.h:18:1:
	X_NSTREAM_                                             = 512    // stdio.h:19:1:
	X_OFF64_T_DEFINED                                      = 0      // _mingw_off_t.h:12:1:
	X_OFF_T_                                               = 0      // _mingw_off_t.h:4:1:
	X_OFF_T_DEFINED                                        = 0      // _mingw_off_t.h:2:1:
	X_OLD_P_OVERLAY                                        = 2      // stdio.h:1557:1:
	X_PGLOBAL                                              = 0      // _mingw.h:342:1:
	X_PTRDIFF_T_                                           = 0      // corecrt.h:90:1:
	X_PTRDIFF_T_DEFINED                                    = 0      // corecrt.h:88:1:
	X_P_DETACH                                             = 4      // stdio.h:1559:1:
	X_P_NOWAIT                                             = 1      // stdio.h:1556:1:
	X_P_NOWAITO                                            = 3      // stdio.h:1558:1:
	X_P_OVERLAY                                            = 2      // stdio.h:1560:1:
	X_P_WAIT                                               = 0      // stdio.h:1555:1:
	X_P_tmpdir                                             = "\\"   // stdio.h:42:1:
	X_RSIZE_T_DEFINED                                      = 0      // corecrt.h:58:1:
	X_SECURECRT_FILL_BUFFER_PATTERN                        = 0xFD   // _mingw.h:349:1:
	X_SIZE_T_DEFINED                                       = 0      // corecrt.h:37:1:
	X_SPAWNV_DEFINED                                       = 0      // stdio.h:1567:1:
	X_SSIZE_T_DEFINED                                      = 0      // corecrt.h:47:1:
	X_STDIO_CONFIG_DEFINED                                 = 0      // corecrt_stdio_config.h:8:1:
	X_STDIO_DEFINED                                        = 0      // stdio.h:1453:1:
	X_STDIO_S_DEFINED                                      = 0      // stdio_s.h:25:1:
	X_STDSTREAM_DEFINED                                    = 0      // stdio.h:112:1:
	X_SYS_OPEN                                             = 20     // stdio.h:63:1:
	X_TAGLC_ID_DEFINED                                     = 0      // corecrt.h:447:1:
	X_THREADLOCALEINFO                                     = 0      // corecrt.h:456:1:
	X_TIME32_T_DEFINED                                     = 0      // corecrt.h:122:1:
	X_TIME64_T_DEFINED                                     = 0      // corecrt.h:127:1:
	X_TIME_T_DEFINED                                       = 0      // corecrt.h:139:1:
	X_TWO_DIGIT_EXPONENT                                   = 0x1    // stdio.h:135:1:
	X_UINTPTR_T_DEFINED                                    = 0      // corecrt.h:75:1:
	X_USE_32BIT_TIME_T                                     = 0      // _mingw.h:372:1:
	X_VA_LIST_DEFINED                                      = 0      // <builtin>:55:1:
	X_W64                                                  = 0      // _mingw.h:296:1:
	X_WAIT_CHILD                                           = 0      // stdio.h:1562:1:
	X_WAIT_GRANDCHILD                                      = 1      // stdio.h:1563:1:
	X_WCHAR_T_DEFINED                                      = 0      // corecrt.h:101:1:
	X_WCTYPE_T_DEFINED                                     = 0      // corecrt.h:108:1:
	X_WIN32                                                = 1      // <predefined>:164:1:
	X_WIN32_WINNT                                          = 0x502  // _mingw.h:233:1:
	X_WINT_T                                               = 0      // corecrt.h:110:1:
	X_WSPAWN_DEFINED                                       = 0      // stdio.h:1543:1:
	X_WSTDIO_DEFINED                                       = 0      // stdio.h:958:1:
	X_WSTDIO_S_DEFINED                                     = 0      // stdio_s.h:579:1:
	X_X86_                                                 = 1      // <predefined>:169:1:
	I386                                                   = 1      // <predefined>:171:1:
)

type Ptrdiff_t = int32 /* <builtin>:3:26 */

type Size_t = uint32 /* <builtin>:9:23 */

type Wchar_t = uint16 /* <builtin>:15:24 */

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

type Va_list = X__builtin_va_list /* <builtin>:50:27 */

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the Wine project.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// This macro holds an monotonic increasing value, which indicates
//    a specific fix/patch is present on trunk.  This value isn't related to
//    minor/major version-macros.  It is increased on demand, if a big
//    fix was applied to trunk.  This macro gets just increased on trunk.  For
//    other branches its value won't be modified.

// mingw.org's version macros: these make gcc to define
//    MINGW32_SUPPORTS_MT_EH and to use the _CRT_MT global
//    and the __mingwthr_key_dtor() function from the MinGW
//    CRT in its private gthr-win32.h header.

// Set VC specific compiler target macros.

// For x86 we have always to prefix by underscore.

// Special case nameless struct/union.

// MinGW-w64 has some additional C99 printf/scanf feature support.
//    So we add some helper macros to ease recognition of them.

// If _FORTIFY_SOURCE is enabled, some inline functions may use
//    __builtin_va_arg_pack().  GCC may report an error if the address
//    of such a function is used.  Set _FORTIFY_VA_ARG=0 in this case.

// Enable workaround for ABI incompatibility on affected platforms

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// http://msdn.microsoft.com/en-us/library/ms175759%28v=VS.100%29.aspx
// Templates won't work in C, will break if secure API is not enabled, disabled

// https://blogs.msdn.com/b/sdl/archive/2010/02/16/vc-2010-and-memcpy.aspx?Redirected=true
// fallback on default implementation if we can't know the size of the destination

// Include _cygwin.h if we're building a Cygwin application.

// Target specific macro replacement for type "long".  In the Windows API,
//    the type long is always 32 bit, even if the target is 64 bit (LLP64).
//    On 64 bit Cygwin, the type long is 64 bit (LP64).  So, to get the right
//    sized definitions and declarations, all usage of type long in the Windows
//    headers have to be replaced by the below defined macro __LONG32.

// C/C++ specific language defines.

// Note the extern. This is needed to work around GCC's
// limitations in handling dllimport attribute.

// Attribute `nonnull' was valid as of gcc 3.3.  We don't use GCC's
//    variadiac macro facility, because variadic macros cause syntax
//    errors with  --traditional-cpp.

//  High byte is the major version, low byte is the minor.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// for backward compatibility

type X__gnuc_va_list = X__builtin_va_list /* vadefs.h:24:29 */

type Ssize_t = int32 /* corecrt.h:52:13 */

type Rsize_t = Size_t /* corecrt.h:57:16 */

type Intptr_t = int32 /* corecrt.h:69:13 */

type Uintptr_t = uint32 /* corecrt.h:82:22 */

type Wint_t = uint16   /* corecrt.h:111:24 */
type Wctype_t = uint16 /* corecrt.h:112:24 */

type Errno_t = int32 /* corecrt.h:118:13 */

type X__time32_t = int32 /* corecrt.h:123:14 */

type X__time64_t = int64 /* corecrt.h:128:35 */

type Time_t = X__time32_t /* corecrt.h:141:20 */

type Threadlocaleinfostruct = struct {
	Frefcount      int32
	Flc_codepage   uint32
	Flc_collate_cp uint32
	Flc_handle     [6]uint32
	Flc_id         [6]LC_ID
	Flc_category   [6]struct {
		Flocale    uintptr
		Fwlocale   uintptr
		Frefcount  uintptr
		Fwrefcount uintptr
	}
	Flc_clike            int32
	Fmb_cur_max          int32
	Flconv_intl_refcount uintptr
	Flconv_num_refcount  uintptr
	Flconv_mon_refcount  uintptr
	Flconv               uintptr
	Fctype1_refcount     uintptr
	Fctype1              uintptr
	Fpctype              uintptr
	Fpclmap              uintptr
	Fpcumap              uintptr
	Flc_time_curr        uintptr
} /* corecrt.h:435:1 */

type Pthreadlocinfo = uintptr /* corecrt.h:437:39 */
type Pthreadmbcinfo = uintptr /* corecrt.h:438:36 */

type Localeinfo_struct = struct {
	Flocinfo Pthreadlocinfo
	Fmbcinfo Pthreadmbcinfo
} /* corecrt.h:441:9 */

type X_locale_tstruct = Localeinfo_struct /* corecrt.h:444:3 */
type X_locale_t = uintptr                 /* corecrt.h:444:19 */

type TagLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
} /* corecrt.h:435:1 */

type LC_ID = TagLC_ID  /* corecrt.h:452:3 */
type LPLC_ID = uintptr /* corecrt.h:452:9 */

type Threadlocinfo = Threadlocaleinfostruct /* corecrt.h:487:3 */

type X_iobuf = struct {
	F_ptr      uintptr
	F_cnt      int32
	F_base     uintptr
	F_flag     int32
	F_file     int32
	F_charbuf  int32
	F_bufsiz   int32
	F_tmpfname uintptr
} /* stdio.h:24:3 */

type FILE = X_iobuf /* stdio.h:34:25 */

type X_off_t = int32 /* _mingw_off_t.h:5:16 */
type Off32_t = int32 /* _mingw_off_t.h:7:16 */

type X_off64_t = int64 /* _mingw_off_t.h:13:39 */
type Off64_t = int64   /* _mingw_off_t.h:15:39 */

type Off_t = Off64_t /* _mingw_off_t.h:24:17 */ // A pointer to an array of FILE

type Fpos_t = int64 /* stdio.h:102:37 */

var _ int8 /* gen.c:2:13: */