package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint32 }
func genFn0() int32 {
    if false {
        var v4 uint64 = 1
        _ = v4
    }
    var v5 int16 = 32767
    _ = v5
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{"a", "hello", ("a" + ("\n\t" + "hello"))}
    var v3 int = 563
    fmt.Println(v1, v1)
    var v6 string = "a"
    v7 := GenS1{F0: false, F1: 119}
    v8 := GenS1{F0: false, F1: 42}
    switch v3 {
    case 32768:
        var v9 float64 = 0.0
        _ = v9
    default:
        v10 := GenS1{F0: false, F1: 106}
        _ = v10
    }
    var v11 uint16 = 1
    var v12 uint32 = 65535
    var v13 int = (v3 & (v3 >> ((42) & 63)))
    var v14 uint32 = 2
    if true {
        v1 += v12
    } else {
        v1 += (^(v12))
        v11 += (v11 / (((v11 ^ v11)) | 1))
        var v15 int16 = -749
        _ = v15
    }
    var v16 float32 = 0.0
    if false {
        if false {
            if (!(!(v16 != v16))) {
                v3 += 9223372036854775807
                if true {
                    if true {
                        v17 := GenS1{F0: true, F1: 255}
                        v11 = v11
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 uint32 = v1
                            _ = fi18
                            _ = v19
                        }
                    }
                    v20 := GenS1{F0: false, F1: 2}
                    _ = v20
                }
                var v21 uint = 18446744073709551615
                _ = v21
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v6 = v6
        } else {
            v1 += v14
            if (-9223372036854775808 >= 9223372036854775807) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    _ = len(v2)
                    _ = fi22
                }
            } else {
                for fi23 := 0; fi23 < 11; fi23++ {
                    fmt.Println(v14)
                    v2 = append(v2, "hello")
                    var v24 uint8 = 255
                    _ = fi23
                    _ = v24
                }
            }
        }
        v25 := GenS2{F0: v14, F1: 255}
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = v25
    } else {
        for fi26 := 0; fi26 < 15; fi26++ {
            v16 = v16
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi27 := 0; fi27 < 4; fi27++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v14 = v12
                var v28 int64 = 945
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS1{F0: true, F1: 255}
            for fi31 := 0; fi31 < 5; fi31++ {
                fmt.Println(v14, v1)
                fmt.Println(v6, v1)
                v16 = float32((3.14159 * 0.0))
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
        if (!((v11 & (v11 - (^(v11)))) > (v11 + v11))) {
            var v32 string = v6
            _ = v32
        } else {
            if (102 <= -536) {
                switch 127 {
                case -946:
                    v14 -= v1
                case -993:
                    v14 -= v12
                case 1:
                    v33 := GenS3{F0: 1, F1: true}
                    _ = v33
                }
            } else {
                var v34 uint = 255
                _ = v34
            }
            v2 = append(v2, v6)
        }
    }
    switch v6 {
    case "":
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 uint64 = 1
            v2 = append(v2, "")
            _ = fi35
            _ = v36
        }
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v13 += (v13 * v13)
    if (9223372036854775807 <= 42) {
        for fi37 := 0; fi37 < 16; fi37++ {
            switch (^(v13)) {
            case 32767:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -1:
                for fi38 := 0; fi38 < 17; fi38++ {
                    switch v3 {
                    case 629:
                        switch "a" {
                        case "hello":
                            switch fi38 {
                            case -9223372036854775808:
                                v3 = (fi38 - (v3 ^ (v13 - (fi37 - (fi37 | v3)))))
                            }
                        case "foo bar":
                            v11 += (^(v11))
                        default:
                            if (v6 == v6) {
                                v6 = ("" + (v6 + "foo bar"))
                                v2 = append(v2, ((v6 + "\n\t") + (v6 + "hello")))
                                var v39 uint64 = 349
                                _ = v39
                            } else {
                                _ = len(v2)
                                if false {
                                    v6 = v6
                                    switch "" {
                                    case "\n\t":
                                        v3 = (fi38 ^ (v3 & 32767))
                                    case "a":
                                        v40 := GenS3{F0: 42, F1: (!false)}
                                        v41 := GenS3{F0: 9223372036854775807, F1: true}
                                        if true {
                                            var v42 int32 = 0
                                            var v43 uint32 = uint32(v11)
                                            _ = v42
                                            _ = v43
                                        }
                                        fmt.Println(v13, v6, v11)
                                        if false {
                                            var v44 int32 = 0
                                            if true {
                                                v1 = (^(v1))
                                            }
                                            _ = v44
                                        } else {
                                            fmt.Println(fi38)
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = v6
                                                }
                                            } else {
                                                var v45 uint16 = v11
                                                _ = v45
                                            }
                                        }
                                        v14 = v12
                                        var v46 bool = false
                                        switch "\n\t" {
                                        case "hello":
                                            v1 = 2
                                        case "foo bar":
                                            v47 := GenS4{F0: (v11 * v11)}
                                            _ = v47
                                        default:
                                            v11 = v11
                                        }
                                        var v48 float64 = 3.14159
                                        fmt.Println(v11, fi38, v11)
                                        _ = v40
                                        _ = v41
                                        _ = v46
                                        _ = v48
                                    default:
                                        if (42 < 0) {
                                            if len(v2) > 0 {
                                                v2[0] = ("hello" + v6)
                                            }
                                        } else {
                                            var v49 uint64 = 4294967295
                                            v6 = ("hello" + (v6 + v6))
                                            _ = v49
                                        }
                                    }
                                } else {
                                    fmt.Println(v13, v13, v14)
                                }
                                v13 = (fi37 - (v13 << ((1) & 63)))
                            }
                        }
                    case -9223372036854775808:
                        v50 := GenS4{F0: v11}
                        _ = v50
                    default:
                        fmt.Println(v1, v12)
                    }
                    _ = fi38
                }
            case 42:
                switch v6 {
                case "\n\t":
                    v51 := GenS2{F0: 1, F1: 0}
                    _ = v51
                }
            default:
                for fi52 := 0; fi52 < 6; fi52++ {
                    var v53 uint64 = 18446744073709551615
                    v54 := GenS5{F0: (v13 * (v3 + -1))}
                    v11 += (v11 | (^(v11)))
                    _ = fi52
                    _ = v53
                    _ = v54
                }
            }
            for fi55 := 0; fi55 < 3; fi55++ {
                var v56 int8 = 127
                v3 = (fi37 - fi55)
                _ = fi55
                _ = v56
            }
            _ = fi37
        }
        var v57 float64 = ((1e10 - ((-2.25 + 1e10) - (((3.14159 - (1e10 - float64(float32(3.14159)))) / ((-2.25 / ((-2.25 / 1.0) - (-2.25 - ((3.14159 * 3.14159) / 1.0)))) - ((0.0 / (0.1 / (1e10 + (0.1 / 3.14159)))) * float64((0.0 + float32(float64((0.0 * (v16 * 1.5))))))))) - (1e10 * float64(0.1))))) / 3.14159)
        _ = v57
    } else {
        var v58 int64 = 1
        _ = v58
    }
    v59 := GenS6{F0: 699, F1: (v12 * v12)}
    var v60 int64 = 200
    if true {
        v61 := GenS2{F0: v1, F1: 127}
        var v62 int = 325
        if false {
            fmt.Println(v1)
            for fi63 := 0; fi63 < 7; fi63++ {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                _ = fi63
            }
        } else {
            v6 = (v6 + v6)
            var v64 uint64 = 18446744073709551615
            _ = v64
        }
        _ = v61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v59
    _ = v6
    _ = v60
    _ = v7
    _ = v8
}
