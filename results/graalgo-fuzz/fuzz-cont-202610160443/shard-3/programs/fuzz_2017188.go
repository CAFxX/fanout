package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint; F2 uint }
type GenS8 struct { F0 uint; F1 int32; F2 bool }
func genFn0(p0 uint32) int32 {
    p0 += p0
    p0 += (p0 & p0)
    p0 += p0
    var v5 int = -921
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 0
    v4 := make(map[int]bool)
    v3 += v3
    v6 := []int{2147483647, 9223372036854775807, 42, 729}
    if (32767 != 65535) {
        v7 := GenS1{F0: 42, F1: 742, F2: 39}
        v8 := GenS2{F0: 42, F1: 1000}
        var v9 int64 = 9223372036854775807
        _ = v7
        _ = v8
        _ = v9
    }
    switch -607 {
    case -976:
        fmt.Println(v3, v3)
    case -872:
        for fi10 := 0; fi10 < 13; fi10++ {
            if false {
                switch fi10 {
                case 128:
                    fmt.Println(v3, fi10)
                }
                var v11 uint32 = 2147483647
                v1 = v1
                _ = v11
            } else {
                if ((("foo bar" == ("hello" + "foo bar")) || false) || (42 != 965)) {
                    delete(v2, 2147483647)
                    var v12 string = ("a" + "foo bar")
                    _ = v12
                }
                var v13 uint16 = 42
                delete(v4, (fi10 | fi10))
                _ = v13
            }
            v1 = v1
            _ = fi10
        }
    }
    v3 = (^(v3))
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v3, v1, v1)
        v15 := v2[735]
        _ = fi14
        _ = v15
    }
    if false {
        v1 = 141
        switch 32768 {
        case 778:
            for fi16 := 0; fi16 < 3; fi16++ {
                switch (fi16 * (fi16 / ((fi16) | 1))) {
                case -522:
                    v17 := GenS2{F0: 0, F1: 32767}
                    _ = v17
                case -2147483648:
                    fmt.Println(v1)
                case 727:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (!(!(false && false))) {
                            v3 -= (v3 << ((4294967295) & 63))
                            _ = len(v2)
                            fmt.Println(fi18, fi18)
                        }
                        var v19 int = fi16
                        if true {
                            if false {
                                var v20 int64 = 1
                                _ = v20
                            } else {
                                switch fi18 {
                                case 127:
                                    v3 = 128
                                }
                                var v21 string = "foo bar"
                                _ = v21
                            }
                            var v22 int16 = 373
                            switch 153 {
                            case 4294967295:
                                v23 := GenS2{F0: 1, F1: (-(v22))}
                                _ = v23
                            }
                            _ = v22
                        } else {
                            v19 = 2147483647
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                var v24 uint = 967
                if (int(v3) >= fi16) {
                    switch (-(fi16)) {
                    case 355:
                        var v25 int = (-(fi16))
                        _ = v25
                    }
                } else {
                    v26 := GenS1{F0: 0, F1: 855, F2: 255}
                    var v27 uint32 = 65535
                    switch 128 {
                    case -9223372036854775808:
                        if false {
                            _ = len(v6)
                            for fi28 := 0; fi28 < 1; fi28++ {
                                if ((float64(3.14159) <= 0.1) || (1 >= -510)) {
                                    v1 += (v1 - (v1 + uint64(v27)))
                                    var v29 int64 = -143
                                    v30 := GenS2{F0: 1, F1: 32767}
                                    _ = v29
                                    _ = v30
                                } else {
                                    fmt.Println(v3)
                                    var v31 uint = 8
                                    v3 += v3
                                    fmt.Println(fi28, v31)
                                    _ = v31
                                }
                                v24 = 255
                                v32 := GenS2{F0: 127, F1: -923}
                                _ = fi28
                                _ = v32
                            }
                            v33 := GenS3{F0: (v27 / ((v27) | 1)), F1: fi16, F2: true}
                            _ = v33
                        }
                    case 2147483647:
                        v1 -= uint64(v24)
                    case 4294967295:
                        switch ((("a" + "foo bar") + "foo bar") + "\n\t") {
                        case "\n\t":
                            if false {
                                v34 := GenS4{F0: uint16(v27)}
                                for fi35 := 0; fi35 < 14; fi35++ {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi35
                                }
                                _ = v34
                            }
                        case "hello":
                            fmt.Println(v1)
                        case "a":
                            v6 = append(v6, fi16)
                        }
                    }
                    _ = v26
                    _ = v27
                }
                _ = fi16
                _ = v24
            }
        case 42:
            for fi36 := 0; fi36 < 19; fi36++ {
                if false {
                    v3 = (v3 * v3)
                } else {
                    v37 := GenS1{F0: 43, F1: 0, F2: v3}
                    _ = v37
                }
                _ = fi36
            }
        }
        var v38 int64 = 9223372036854775807
        _ = v38
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        if false {
            switch fi39 {
            case -1:
                v40 := GenS3{F0: 73, F1: 9223372036854775807, F2: (((!(499 == (fi39 % ((32767) | 1)))) && (true || ((32767 <= 65535) && (4294967295 > v1)))) || true)}
                _ = v40
            default:
                for fi41 := 0; fi41 < 1; fi41++ {
                    v42 := GenS5{F0: (!false)}
                    switch fi39 {
                    case -9223372036854775808:
                        var v43 uint = 250
                        _ = v43
                    default:
                        if (!(!(!true))) {
                            var v44 uint = uint(v3)
                            _ = v44
                        } else {
                            var v45 bool = true
                            v46 := GenS6{F0: v1}
                            _ = v45
                            _ = v46
                        }
                    }
                    if false {
                        if false {
                            v47 := v4[fi39]
                            v3 = 42
                            _ = v47
                        } else {
                            v1 = (v1 / (((v1 * 247)) | 1))
                        }
                        if true {
                            v1 = v1
                            var v48 int = 128
                            _ = v48
                        }
                    } else {
                        _ = len(v6)
                        v1 = v1
                        if true {
                            if false {
                                switch fi41 {
                                case -2147483648:
                                    v49 := GenS7{F0: -32768, F1: 1, F2: 4294967295}
                                    _ = v49
                                default:
                                    var v50 uint8 = 128
                                    _ = v50
                                }
                            }
                            var v51 int16 = 1
                            _ = v51
                        } else {
                            v52 := GenS8{F0: 175, F1: 42, F2: false}
                            _ = v52
                        }
                        var v53 string = "a"
                        _ = v53
                    }
                    _ = fi41
                    _ = v42
                }
            }
        }
        _ = fi39
    }
    if (!false) {
        for fi54 := 0; fi54 < 19; fi54++ {
            for fi55 := 0; fi55 < 20; fi55++ {
                v3 += (v3 >> ((277) & 63))
                var v56 int16 = 1000
                _ = fi55
                _ = v56
            }
            var v57 string = ("foo bar" + "")
            v1 = (v1 | (v1 % (((^(v1))) | 1)))
            _ = fi54
            _ = v57
        }
        v3 = 0
        if ("a" != "\n\t") {
            var v58 int = 32
            _ = v58
        } else {
            switch -716 {
            case 32767:
                for fi59 := 0; fi59 < 11; fi59++ {
                    fmt.Println(v1)
                    v60 := GenS8{F0: 2, F1: -359, F2: (0 < -939)}
                    _ = fi59
                    _ = v60
                }
            case 88:
                switch 2147483647 {
                case -473:
                    var v61 float32 = -2.25
                    _ = v61
                case 128:
                    var v62 int16 = 1
                    _ = v62
                default:
                    var v63 int16 = 1
                    _ = v63
                }
            default:
                v1 -= v1
            }
        }
    } else {
        fmt.Println(v3)
        var v64 uint64 = (^(v1))
        _ = v64
    }
    v3 += 0
    if false {
        fmt.Println(v1)
        if len(v6) > 0 {
            v6[0] = -212
        }
    } else {
        fmt.Println(v1, v3)
        if false {
            fmt.Println(v1, v1)
            if (-1 <= -110) {
                v65 := GenS2{F0: 68, F1: 522}
                v66 := GenS2{F0: 0, F1: 286}
                _ = v65
                _ = v66
            }
            if false {
                v6 = append(v6, 4294967295)
            }
        } else {
            var v67 int64 = -1
            _ = v67
        }
    }
    fmt.Println(v3, v3)
    if false {
        delete(v2, 65535)
        v68 := GenS3{F0: 92, F1: 237, F2: false}
        v3 -= v3
        _ = v68
    } else {
        if (1 <= (v3 + (v3 >> ((935) & 63)))) {
            var v69 int32 = -25
            switch (("\n\t" + "hello") + "foo bar") {
            case "hello":
                v1 = v1
            }
            _ = v69
        } else {
            v4[-9223372036854775808] = false
        }
        if false {
            for fi70 := 0; fi70 < 1; fi70++ {
                if false {
                    for fi71 := 0; fi71 < 2; fi71++ {
                        var v72 uint16 = 32767
                        _ = fi71
                        _ = v72
                    }
                    delete(v2, 2147483647)
                } else {
                    var v73 bool = true
                    if ((4294967295 != uint32(v1)) && (0 < (^(v3)))) {
                        v1 += v1
                    }
                    var v74 int = (fi70 * 4294967295)
                    _ = v73
                    _ = v74
                }
                _ = fi70
            }
            for fi75 := 0; fi75 < 16; fi75++ {
                if (-1 != 1) {
                    v3 = (v3 ^ (^(v3)))
                }
                _ = fi75
            }
        }
        v6 = append(v6, -1)
    }
    v1 = (v1 * 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
