package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 string }
type GenS5 struct { F0 int8; F1 int16; F2 uint16 }
func genFn0() int64 {
    var v7 int32 = -1
    if (false && false) {
        fmt.Println(v7)
        var v8 string = "a"
        if ("a" > ("hello" + "\n\t")) {
            v9 := GenS1{F0: (false && (!(!(!false)))), F1: 692, F2: -914}
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(fi10)
                if (true || (fi10 >= fi10)) {
                    var v11 uint32 = 2147483648
                    _ = v11
                } else {
                    v12 := GenS1{F0: false, F1: 32767, F2: 42}
                    _ = v12
                }
                v8 = v8
                _ = fi10
            }
            v7 = -846
            _ = v9
        } else {
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := GenS1{F0: (((true && false) || (42 < 65535)) && (false || (7 <= 32767))), F1: 494, F2: 1}
                var v15 bool = true
                v8 = v8
                _ = fi13
                _ = v14
                _ = v15
            }
            switch v8 {
            case "foo bar":
                for fi16 := 0; fi16 < 10; fi16++ {
                    if (false && (true && false)) {
                        v17 := []int{fi16, (fi16 << ((4294967295) & 63)), (fi16 | (-(fi16))), (fi16 / ((fi16) | 1))}
                        v18 := GenS2{F0: -359}
                        v8 = v8
                        _ = v17
                        _ = v18
                    }
                    v7 -= v7
                    var v19 int32 = 1
                    _ = fi16
                    _ = v19
                }
            case "hello":
                v20 := []int{386, -678}
                _ = v20
            case "":
                v7 = v7
            default:
                for fi21 := 0; fi21 < 9; fi21++ {
                    if ((!(!(false && (127 != -50)))) && (false && (((0.0 / 0.1) + 3.14159) != (((-2.25 * 1e10) / 3.14159) / (3.14159 / -2.25))))) {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 uint8 = 255
                            _ = fi22
                            _ = v23
                        }
                        fmt.Println(v8)
                        v8 = v8
                    }
                    v7 -= v7
                    _ = fi21
                }
            }
            if (1 > 2) {
                var v24 uint32 = 880
                if true {
                    v7 += (v7 ^ (v7 / ((v7) | 1)))
                } else {
                    fmt.Println(v24, v8)
                    v7 += 2147483647
                    v25 := GenS1{F0: false, F1: 2, F2: 1}
                    _ = v25
                }
                fmt.Println(v24, v7)
                _ = v24
            }
        }
        _ = v8
    } else {
        var v26 bool = (32767 > 2)
        var v27 int64 = 0
        var v28 float32 = 1.5
        _ = v26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        v7 = v7
        fmt.Println(fi29, v7)
        _ = fi29
    }
    v30 := GenS3{F0: 435, F1: 0, F2: 948}
    _ = v30
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-2147483648, 32767, 127}
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    var v5 uint = 622
    var v6 int32 = -1
    for fi31 := 0; fi31 < 19; fi31++ {
        for fi32 := 0; fi32 < 17; fi32++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v33 int64 = genFn0()
            v5 = (v5 >> ((4294967295) & 63))
            v34 := GenS1{F0: (9223372036854775807 > fi31), F1: 564, F2: 1}
            fmt.Println(v5)
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 7; fi35++ {
            v6 = -393
            if (-213 > 957) {
                v36 := GenS1{F0: true, F1: 0, F2: int16(v6)}
                var v37 float32 = 1e10
                _ = v36
                _ = v37
            } else {
                v6 = -2147483648
                if true {
                    v5 -= v5
                    if true {
                        v6 = v6
                    }
                    switch ("\n\t" + (("\n\t" + "") + (((("foo bar" + "hello") + "a") + "\n\t") + ""))) {
                    case "a":
                        fmt.Println(fi35)
                    default:
                        if (42 >= 166) {
                            var v38 uint16 = 744
                            v39 := v3[("\n\t" + "\n\t")]
                            _ = v38
                            _ = v39
                        } else {
                            var v40 float64 = 1e10
                            _ = v40
                        }
                    }
                } else {
                    switch 98 {
                    case -1:
                        if ("a" == ("" + "foo bar")) {
                            _ = len(v2)
                        }
                    case -2147483648:
                        v41 := GenS3{F0: 42, F1: fi35, F2: (^(v5))}
                        _ = v41
                    }
                    _ = len(v2)
                    v5 -= (v5 | 255)
                    if (fi35 >= fi31) {
                        if false {
                            v5 = (v5 + (v5 | 4294967295))
                            fmt.Println(fi31)
                            var v42 int64 = 9223372036854775807
                            _ = v42
                        }
                    } else {
                        switch fi31 {
                        case 876:
                            var v43 int64 = genFn0()
                            _ = v43
                        case 1:
                            v44 := GenS1{F0: false, F1: 728, F2: 1}
                            _ = v44
                        default:
                            fmt.Println(v5, v6)
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi35
        }
        v45 := GenS4{F0: 4294967295, F1: 2, F2: "foo bar"}
        _ = fi31
        _ = v45
    }
    var v46 uint16 = 32767
    v1[65535] = ((0 < 2) && false)
    if ((0.0 * (1.5 * float64(1e10))) <= float64(-2.25)) {
        if (!false) {
            v47 := GenS2{F0: 924}
            var v48 int = 885
            v1[467] = true
            _ = v47
            _ = v48
        }
        if true {
            v6 = v6
            v6 = v6
        } else {
            var v49 bool = true
            _ = v49
        }
    } else {
        v5 += v5
    }
    switch 793 {
    case 42:
        if (100 < 127) {
            var v50 uint64 = 4294967295
            if false {
                if false {
                    var v51 uint16 = (v46 & v46)
                    if (1 > 100) {
                        switch 127 {
                        case -1:
                            if true {
                                if len(v2) > 0 {
                                    v2[0] = int(v50)
                                }
                            }
                        }
                    }
                    var v52 uint = v5
                    _ = v51
                    _ = v52
                }
            } else {
                switch 127 {
                case -9223372036854775808:
                    var v53 uint16 = (v46 % (((v46 * v46)) | 1))
                    _ = v53
                case 1:
                    var v54 uint64 = (v50 * (v50 - v50))
                    _ = v54
                }
                if (1.5 >= float32(((0.0 - (1e10 - 0.1)) / 1e10))) {
                    var v55 uint64 = (v50 & v50)
                    v2 = append(v2, 0)
                    _ = v55
                }
            }
            _ = v50
        } else {
            v46 -= v46
        }
    case -444:
        for fi56 := 0; fi56 < 9; fi56++ {
            var v57 string = "foo bar"
            v46 += v46
            for fi58 := 0; fi58 < 5; fi58++ {
                v59 := GenS1{F0: false, F1: (v46 & 2), F2: 32767}
                var v60 int32 = -701
                var v61 float64 = -2.25
                _ = fi58
                _ = v59
                _ = v60
                _ = v61
            }
            _ = fi56
            _ = v57
        }
    case 9223372036854775807:
        v6 += v6
    }
    v1[371] = true
    v62 := GenS5{F0: -1, F1: -839, F2: v46}
    if false {
        for fi63 := 0; fi63 < 20; fi63++ {
            for fi64 := 0; fi64 < 3; fi64++ {
                v5 += 18446744073709551615
                v5 = (v5 & v5)
                _ = fi64
            }
            var v65 string = ("\n\t" + "\n\t")
            _ = fi63
            _ = v65
        }
        if false {
            v5 = (^(v5))
            var v66 uint8 = 0
            v66 += v66
            _ = v66
        }
        var v67 uint16 = (v46 + 255)
        _ = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v62
}
