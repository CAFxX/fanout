package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int16 }
func genFn0() int16 {
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 = (v7 - (v7 >> ((255) & 63)))
        _ = fi8
    }
    if false {
        var v9 uint64 = 0
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(v7)
            if (-9223372036854775808 == 0) {
                fmt.Println(v7, v7, fi10)
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint64 = (v9 / (((^(v9))) | 1))
                    _ = fi11
                    _ = v12
                }
                var v13 int64 = 0
                _ = v13
            }
            _ = fi10
        }
        fmt.Println(v7, v7)
        _ = v9
    } else {
        var v14 uint = 753
        _ = v14
    }
    var v15 int32 = -2147483648
    var v16 uint8 = 128
    _ = v15
    _ = v16
    _ = v7
    return -275
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{898, 2}
    v3 := make(map[int]bool)
    var v4 uint16 = 32767
    var v5 uint16 = (v4 << ((42) & 63))
    v6 := make(map[string]string)
    v4 = v4
    switch 9223372036854775807 {
    case 32767:
        var v17 int16 = -1
        _ = v17
    case 32768:
        var v18 uint = 4294967295
        _ = v18
    default:
        if ((false || false) && (33 != (v1 + (^(v1))))) {
            var v19 uint8 = 2
            fmt.Println(v4, v1, v4)
            var v20 int8 = 1
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 9223372036854775807, F1: 946}
            v1 += (v1 >> ((4294967295) & 63))
            if (false && (2 > 65535)) {
                v4 -= v5
            } else {
                var v22 int = 4294967295
                _ = v22
            }
            _ = v21
        }
    }
    v23 := GenS1{F0: -1, F1: -495}
    if true {
        for fi24 := 0; fi24 < 3; fi24++ {
            if true {
                v2 = append(v2, uint32(v4))
                if (!true) {
                    v1 = v1
                    var v25 int16 = 248
                    var v26 float32 = 1.5
                    _ = v25
                    _ = v26
                } else {
                    var v27 int32 = 350
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v4 = 65535
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 uint8 = 1
                        switch 1 {
                        case 983:
                            v31 := GenS2{F0: 18446744073709551615, F1: 0, F2: 919}
                            _ = v31
                        case 42:
                            v1 += v1
                        default:
                            var v32 float32 = -2.25
                            _ = v32
                        }
                        switch "hello" {
                        case "\n\t":
                            v30 -= v30
                        case "foo bar":
                            _ = len(v6)
                        case "":
                            fmt.Println(v5, v4, v27)
                        default:
                            if true {
                                var v33 int = fi24
                                _ = v33
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v27
                }
            } else {
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 int = (fi24 ^ (fi24 & (^(fi34))))
                    switch int(v4) {
                    case -537:
                        var v36 int64 = -931
                        _ = v36
                    case 127:
                        v37 := GenS1{F0: v35, F1: fi34}
                        _ = v37
                    case 32768:
                        v4 = (v4 * (v5 % ((0) | 1)))
                    default:
                        var v38 bool = (576 < -1)
                        _ = v38
                    }
                    v39 := GenS3{F0: (true || (!(!(1 >= 42)))), F1: 533, F2: v4}
                    _ = fi34
                    _ = v35
                    _ = v39
                }
                if true {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    v1 += v1
                }
            }
            if true {
                var v40 uint8 = 128
                _ = v40
            }
            _ = fi24
        }
    } else {
        v41 := GenS3{F0: true, F1: 197, F2: v5}
        switch "a" {
        case "":
            var v42 int32 = 1
            _ = v42
        case "hello":
            v6["\n\t"] = (("" + "") + ("\n\t" + ("\n\t" + "")))
        }
        _ = v41
    }
    fmt.Println(v1, v4, v4)
    var v43 uint8 = 127
    v44 := GenS1{F0: 1, F1: 680}
    var v45 int = 1
    if (((v43 - (v43 % ((v43) | 1))) < (^(v43))) && (!true)) {
        var v46 int = v45
        fmt.Println(v45)
        v45 = v45
        _ = v46
    } else {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi47 := 0; fi47 < 3; fi47++ {
                var v48 int32 = -1
                fmt.Println(v5)
                _ = fi47
                _ = v48
            }
        } else {
            if false {
                var v49 float32 = ((0.1 - 3.14159) - float32(-2.25))
                if ((!true) || (!false)) {
                    delete(v3, (v45 / (((v45 / (((v45 + v45)) | 1))) | 1)))
                    fmt.Println(v45)
                }
                if true {
                    v50 := GenS1{F0: v45, F1: (v45 ^ v45)}
                    v51 := GenS4{F0: 1}
                    _ = v50
                    _ = v51
                } else {
                    var v52 string = "a"
                    v1 = 0
                    _ = v52
                }
                _ = v49
            } else {
                v5 -= 272
            }
        }
        fmt.Println(v1, v45)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v53 uint64 = 905
    v53 -= (^(v53))
    for fi54 := 0; fi54 < 3; fi54++ {
        _ = len(v6)
        _ = len(v2)
        _ = fi54
    }
    var v55 int16 = int16(v53)
    var v56 uint16 = v5
    switch "" {
    case "":
        var v57 int8 = (v1 | (v1 << ((634) & 63)))
        _ = v57
    case "\n\t":
        delete(v3, 1)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((-(v55)) < (v55 % ((v55) | 1))) {
        switch v45 {
        case 32768:
            if ((3.14159 * ((-2.25 - 1e10) * 0.0)) > 1e10) {
                switch "" {
                case "hello":
                    for fi58 := 0; fi58 < 10; fi58++ {
                        for fi59 := 0; fi59 < 15; fi59++ {
                            v55 = int16(v43)
                            v60 := v6[("\n\t" + "\n\t")]
                            var v61 bool = false
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        switch "foo bar" {
                        case "a":
                            if (!(false || (!true))) {
                                var v62 bool = false
                                for fi63 := 0; fi63 < 7; fi63++ {
                                    var v64 uint16 = v4
                                    fmt.Println(v4)
                                    var v65 bool = ((^(v64)) > v64)
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                }
                                _ = v62
                            }
                        case "foo bar":
                            var v66 int = fi58
                            _ = v66
                        default:
                            for fi67 := 0; fi67 < 3; fi67++ {
                                var v68 uint32 = uint32(v5)
                                fmt.Println(v45, v4, fi67)
                                v56 += (v4 << ((104) & 63))
                                _ = fi67
                                _ = v68
                            }
                        }
                        v55 = v55
                        _ = fi58
                    }
                case "foo bar":
                    var v69 uint32 = 65535
                    _ = v69
                default:
                    for fi70 := 0; fi70 < 7; fi70++ {
                        var v71 float32 = 0.1
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        _ = fi70
                        _ = v71
                    }
                }
                delete(v3, -1)
            }
        default:
            v53 = 42
        }
        v43 = v43
        var v72 uint64 = (v53 % ((v53) | 1))
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v53
    _ = v55
    _ = v56
    _ = v6
}
