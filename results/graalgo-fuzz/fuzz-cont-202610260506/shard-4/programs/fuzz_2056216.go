package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int8; F2 int }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 42
    var v3 string = "\n\t"
    v1[v2] = ("foo bar" + "hello")
    v3 = "a"
    switch ("foo bar" + "\n\t") {
    case "":
        v3 = (v3 + "")
    case "\n\t":
        var v4 float64 = (0.1 * 0.1)
        _ = v4
    case "foo bar":
        var v5 float32 = 1e10
        _ = v5
    }
    v2 = (v2 | 127)
    v2 += v2
    v3 = v3
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            if true {
                v7 := []int{(v2 - (v2 + fi6)), fi6, -9223372036854775808, 127}
                if true {
                    v2 = (fi6 ^ -457)
                } else {
                    _ = len(v1)
                    if false {
                        if true {
                            v8 := GenS1{F0: (-1 <= 1)}
                            if true {
                                if (1 <= 0) {
                                    v3 = ""
                                    v3 = v3
                                    switch (-(v2)) {
                                    case 2147483647:
                                        switch fi6 {
                                        case 580:
                                            if (-500 <= 1) {
                                                v2 -= (^(fi6))
                                                var v9 uint64 = 18446744073709551615
                                                for fi10 := 0; fi10 < 2; fi10++ {
                                                    switch v3 {
                                                    case "\n\t":
                                                        if true {
                                                            v7 = append(v7, v2)
                                                            switch (-(fi6)) {
                                                            case 32768:
                                                                v9 = (v9 * 0)
                                                            }
                                                        } else {
                                                            v11 := v1[(fi6 / (((fi10 * (v2 | (^(v2))))) | 1))]
                                                            var v12 int16 = 823
                                                            v3 = v11
                                                            _ = v11
                                                            _ = v12
                                                        }
                                                    case "":
                                                        if (0 != 0) {
                                                            _ = len(v1)
                                                            var v13 uint16 = 42
                                                            _ = v13
                                                        } else {
                                                            v9 -= v9
                                                            var v14 int8 = 1
                                                            _ = v14
                                                        }
                                                        v15 := GenS1{F0: (v3 == (v3 + v3))}
                                                        var v16 string = (("" + "") + (("\n\t" + (v3 + "")) + (v3 + "\n\t")))
                                                        _ = v15
                                                        _ = v16
                                                    default:
                                                        switch (fi6 | fi10) {
                                                        case 819:
                                                            var v17 uint16 = 255
                                                            var v18 int = 32767
                                                            var v19 uint = 4
                                                            _ = v17
                                                            _ = v18
                                                            _ = v19
                                                        default:
                                                            var v20 int16 = 0
                                                            _ = v20
                                                        }
                                                    }
                                                    var v21 int32 = -647
                                                    fmt.Println(fi6)
                                                    delete(v1, -817)
                                                    v22 := v1[fi6]
                                                    _ = fi10
                                                    _ = v21
                                                    _ = v22
                                                }
                                                _ = v9
                                            }
                                        case -1:
                                            switch "a" {
                                            case "\n\t":
                                                switch v3 {
                                                case "hello":
                                                    v2 += v2
                                                case "a":
                                                    v23 := GenS1{F0: (float64(3.14159) >= ((((1e10 * (0.1 + (float64(1.5) * 1.5))) / 1e10) - (0.0 - 0.0)) * (-2.25 * 1e10)))}
                                                    _ = v23
                                                default:
                                                    if false {
                                                        for fi24 := 0; fi24 < 17; fi24++ {
                                                            v25 := GenS2{F0: 2, F1: 18446744073709551615}
                                                            _ = fi24
                                                            _ = v25
                                                        }
                                                    } else {
                                                        v2 -= 1
                                                        fmt.Println(v3, fi6, v3)
                                                        switch v2 {
                                                        case -9223372036854775808:
                                                            var v26 string = "\n\t"
                                                            _ = v26
                                                        case 352:
                                                            v2 = v2
                                                        default:
                                                            if ((42 >= 946) || (!true)) {
                                                                for fi27 := 0; fi27 < 20; fi27++ {
                                                                    if (0 > 725) {
                                                                        v28 := GenS2{F0: 4294967295, F1: 4294967295}
                                                                        fmt.Println(fi27, fi27)
                                                                        _ = v28
                                                                    }
                                                                    if false {
                                                                        fmt.Println(v3, v3, fi6)
                                                                    } else {
                                                                        var v29 int64 = 0
                                                                        v3 = v3
                                                                        _ = v29
                                                                    }
                                                                    v1[(fi6 % (((v2 & 1)) | 1))] = v3
                                                                    _ = fi27
                                                                }
                                                                if ((v3 >= (v3 + "foo bar")) || (0 < -1)) {
                                                                    v3 = (v3 + v3)
                                                                }
                                                                v30 := v1[32768]
                                                                _ = v30
                                                            }
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                fmt.Println(v2)
                                            case "":
                                                v31 := GenS1{F0: (9223372036854775807 > 9223372036854775807)}
                                                _ = v31
                                            }
                                        case 0:
                                            var v32 uint16 = 1
                                            _ = v32
                                        default:
                                            var v33 uint32 = 42
                                            _ = v33
                                        }
                                    default:
                                        fmt.Println(v2, fi6, v3)
                                    }
                                } else {
                                    v3 = ((v3 + ("a" + (("" + ("" + "foo bar")) + (v3 + v3)))) + ("\n\t" + v3))
                                    v2 = fi6
                                }
                            }
                            if false {
                                v34 := GenS3{F0: -32768}
                                if (false || (!(false || (false && (42 <= -681))))) {
                                    fmt.Println(v2)
                                }
                                _ = v34
                            } else {
                                if false {
                                    if true {
                                        for fi35 := 0; fi35 < 9; fi35++ {
                                            var v36 string = v3
                                            switch (fi6 + fi35) {
                                            case -1:
                                                fmt.Println(v2, fi6)
                                                fmt.Println(v36, fi35)
                                            default:
                                                v7 = append(v7, v2)
                                            }
                                            var v37 int8 = 100
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                        }
                                    } else {
                                        if true {
                                            v38 := GenS3{F0: 55}
                                            var v39 string = v3
                                            if false {
                                                switch (v3 + "foo bar") {
                                                case "foo bar":
                                                    v40 := GenS4{F0: 0}
                                                    _ = v40
                                                }
                                            } else {
                                                v2 = fi6
                                                v41 := GenS4{F0: 2}
                                                _ = v41
                                            }
                                            _ = v38
                                            _ = v39
                                        }
                                        v42 := GenS5{F0: "a", F1: -1, F2: (v2 - fi6)}
                                        _ = v42
                                    }
                                    if ((-2.25 - 1.5) < 0.0) {
                                        v3 = v3
                                        v2 += 2147483647
                                    } else {
                                        for fi43 := 0; fi43 < 14; fi43++ {
                                            _ = len(v1)
                                            var v44 uint16 = 42
                                            var v45 int32 = 1
                                            _ = fi43
                                            _ = v44
                                            _ = v45
                                        }
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                } else {
                                    fmt.Println(v2)
                                    v2 -= (fi6 * -338)
                                }
                                v46 := GenS5{F0: v3, F1: -45, F2: -980}
                                for fi47 := 0; fi47 < 2; fi47++ {
                                    var v48 uint = 42
                                    v48 = v48
                                    v48 = 0
                                    v49 := GenS1{F0: true}
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                                _ = v46
                            }
                            _ = v8
                        } else {
                            v50 := GenS6{F0: 0}
                            _ = v50
                        }
                        v2 -= -139
                    } else {
                        var v51 uint8 = 1
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = v51
                    }
                }
                _ = v7
            } else {
                var v52 uint = 4294967295
                fmt.Println(fi6, fi6, v2)
                v52 += (v52 * (v52 - (v52 - v52)))
                _ = v52
            }
            var v53 string = ("" + (v3 + v3))
            _ = v53
        } else {
            if (false && true) {
                var v54 uint64 = 553
                for fi55 := 0; fi55 < 8; fi55++ {
                    v3 = v3
                    v2 -= (v2 % (((v2 ^ fi55)) | 1))
                    _ = fi55
                }
                _ = v54
            } else {
                for fi56 := 0; fi56 < 15; fi56++ {
                    fmt.Println(v2, fi6)
                    _ = fi56
                }
            }
            var v57 uint16 = 507
            var v58 int16 = -32768
            _ = v57
            _ = v58
        }
        v59 := []int{(fi6 % ((-1) | 1)), fi6}
        _ = fi6
        _ = v59
    }
    v60 := GenS1{F0: (0 < 1)}
    var v61 float32 = (0.0 / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v61
}
