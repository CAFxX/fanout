package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32, p1 uint64, p2 int16) bool {
    var v7 uint32 = (^(p0))
    p0 -= p0
    var v8 uint8 = 90
    v7 += 272
    fmt.Println(v8, v7, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 182
    var v3 uint8 = v2
    v4 := []string{"\n\t", "hello", ("" + "foo bar")}
    v5 := make(map[string]int)
    var v6 string = "foo bar"
    var v9 int64 = -1
    if true {
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                switch v6 {
                case "":
                    fmt.Println(fi10, v3)
                case "hello":
                    v6 = v6
                case "\n\t":
                    var v11 int64 = v9
                    _ = v11
                default:
                    var v12 float32 = 3.14159
                    _ = v12
                }
                var v13 bool = false
                _ = fi10
                _ = v13
            }
            if ((true || ((1000 == 1) || (102 != (v2 | (^(v3)))))) && (true && true)) {
                v9 -= (v9 | v9)
                v4 = append(v4, "")
                var v14 uint8 = (v2 * v3)
                _ = v14
            } else {
                v2 = v2
            }
        }
        var v15 string = "\n\t"
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if (-2147483648 != 0) {
            delete(v5, (((v6 + "a") + (("\n\t" + "hello") + v6)) + ((v6 + v6) + (v6 + ("hello" + (v6 + ("a" + (v6 + v6))))))))
            var v17 float64 = 0.0
            _ = v17
        } else {
            v6 = v6
        }
        _ = fi16
    }
    v1 += 552
    if true {
        if false {
            v18 := GenS1{F0: 65535, F1: v6}
            v3 -= v3
            v9 += 1
            _ = v18
        }
        if true {
            switch v6 {
            case "hello":
                _ = len(v4)
            default:
                var v19 int = 4294967295
                _ = v19
            }
            fmt.Println(v3)
            v5[(v6 + "\n\t")] = 32768
        }
        v20 := v5[((v6 + (v6 + (v6 + ("a" + "\n\t")))) + "a")]
        _ = v20
    } else {
        v21 := GenS1{F0: 0, F1: v6}
        _ = v21
    }
    var v22 float64 = ((3.14159 + ((1.5 - 1.5) - ((0.0 / float64(1.5)) - -2.25))) * 0.0)
    v4 = append(v4, v6)
    var v23 uint = 42
    var v24 uint64 = 840
    v5[v6] = -863
    var v25 uint = v23
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS1{F0: v25, F1: ("foo bar" + v6)}
        if false {
            if (fi26 < 95) {
                v9 += v9
                v25 = (v25 ^ (v25 % ((v25) | 1)))
                if (((0.1 <= v22) && (v25 == v23)) && false) {
                    v2 = uint8(v23)
                    v5[((v6 + v6) + "")] = 42
                    var v28 uint16 = 2
                    _ = v28
                } else {
                    switch v6 {
                    case "hello":
                        if (0.1 >= 0.0) {
                            v6 = v6
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                }
            } else {
                v29 := GenS2{F0: (v6 + v6), F1: int32(v1), F2: 1}
                _ = v29
            }
        }
        switch "" {
        case "":
            var v30 uint32 = v1
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    if false {
        for fi31 := 0; fi31 < 10; fi31++ {
            v1 = 0
            _ = fi31
        }
        if false {
            if true {
                switch 42 {
                case 991:
                    if (424 >= (^(v25))) {
                        _ = len(v4)
                        v25 = v23
                    } else {
                        v4 = append(v4, "foo bar")
                        if (217 > 0) {
                            v32 := GenS1{F0: (v23 >> ((v25) & 63)), F1: ((v6 + v6) + (v6 + "\n\t"))}
                            v6 = v6
                            var v33 float32 = float32(v22)
                            _ = v32
                            _ = v33
                        } else {
                            fmt.Println(v23, v6)
                            if true {
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    v4 = append(v4, v6)
                                    v35 := GenS1{F0: 2, F1: (v6 + (("" + v6) + ((v6 + "a") + v6)))}
                                    _ = fi34
                                    _ = v35
                                }
                                v36 := GenS3{F0: -967}
                                v37 := GenS4{F0: v9, F1: 1}
                                _ = v36
                                _ = v37
                            } else {
                                fmt.Println(v24, v2)
                            }
                        }
                        for fi38 := 0; fi38 < 1; fi38++ {
                            var v39 int8 = -1
                            _ = fi38
                            _ = v39
                        }
                    }
                case 0:
                    var v40 uint = 65535
                    _ = v40
                case 34:
                    fmt.Println(v25, v6)
                }
                v41 := GenS5{F0: v9, F1: 65535}
                if false {
                    v3 = (v2 << ((uint(v24)) & 63))
                    fmt.Println(v1, v23, v3)
                }
                _ = v41
            } else {
                v42 := GenS3{F0: -993}
                _ = v42
            }
            for fi43 := 0; fi43 < 2; fi43++ {
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 uint64 = v24
                    _ = fi44
                    _ = v45
                }
                v1 -= 2147483647
                v3 += 42
                _ = fi43
            }
            var v46 int16 = 42
            _ = v46
        }
    }
    v4 = append(v4, "\n\t")
    for fi47 := 0; fi47 < 12; fi47++ {
        v48 := GenS6{F0: 2147483648, F1: (v1 + v1)}
        _ = fi47
        _ = v48
    }
    if (-2147483648 <= -1) {
        switch ("\n\t" + ("a" + ("\n\t" + ((v6 + "\n\t") + "a")))) {
        case "a":
            v22 = v22
        case "\n\t":
            v2 += (v2 + v3)
        default:
            fmt.Println(v2, v22, v25)
        }
        var v49 int8 = 1
        if ((v9 % ((726) | 1)) != v9) {
            var v50 int32 = -2147483648
            _ = v50
        }
        _ = v49
    } else {
        switch 9223372036854775807 {
        case 127:
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 bool = true
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi51
                _ = v52
            }
        case 32767:
            var v53 float32 = 3.14159
            _ = v53
        case 2147483647:
            v54 := GenS6{F0: (v1 - (v1 / (((^(v1))) | 1))), F1: (v1 + 335)}
            _ = v54
        default:
            if true {
                v4 = append(v4, (v6 + v6))
                if (v23 != (v23 % ((1) | 1))) {
                    for fi55 := 0; fi55 < 13; fi55++ {
                        var v56 float32 = 1e10
                        _ = fi55
                        _ = v56
                    }
                }
                var v57 int8 = 85
                _ = v57
            }
        }
        var v58 uint32 = v1
        var v59 uint16 = 1
        _ = v58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
