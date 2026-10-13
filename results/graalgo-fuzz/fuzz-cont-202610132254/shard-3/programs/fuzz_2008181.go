package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint8, p1 int, p2 uint64) string {
    fmt.Println(p0, p0)
    p2 -= 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return (("foo bar" + ("a" + ("a" + "a"))) + "\n\t")
}
func genFn1() int64 {
    v7 := []int{}
    var v8 bool = false
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 255
    var v2 string = ""
    v3 := make(map[string]bool)
    var v4 int64 = -9223372036854775808
    var v5 int = 9223372036854775807
    v6 := make(map[int]uint32)
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int8 = 1
        var v11 int16 = 1
        var v12 uint32 = 65535
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{(v5 ^ (v5 - v5)), -2147483648, v5, 127}
    var v14 uint = 255
    if false {
        v13 = append(v13, (v5 | v5))
        var v15 int = v5
        var v16 float64 = float64((3.14159 + ((((((1.5 - -2.25) - 1.5) * (0.0 + ((0.1 * float32(1.5)) / 3.14159))) / (0.1 / (3.14159 * 1.5))) / (0.1 / 1.0)) - 0.1)))
        _ = v15
        _ = v16
    } else {
        switch v2 {
        case "hello":
            switch v5 {
            case 42:
                v14 = (v14 % (((v14 << (((^(v14))) & 63))) | 1))
            default:
                v17 := GenS1{F0: 42, F1: 255}
                _ = v17
            }
        case "\n\t":
            var v18 int = v5
            _ = v18
        default:
            var v19 int32 = -292
            _ = v19
        }
        v20 := GenS1{F0: 42, F1: 112}
        _ = v20
    }
    if (942 < 0) {
        v21 := GenS2{F0: 2, F1: 5, F2: ((v2 + ("" + v2)) + ((v2 + "") + v2))}
        if true {
            _ = len(v13)
            fmt.Println(v5)
        } else {
            v2 = "hello"
        }
        v22 := GenS1{F0: 100, F1: 94}
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: 63}
    var v24 int64 = v4
    var v25 int = v5
    if (0.1 != (0.0 + -2.25)) {
        v14 -= v14
        if true {
            if false {
                fmt.Println(v1)
                v13 = append(v13, 9223372036854775807)
                switch -724 {
                case 0:
                    var v26 uint32 = 867
                    _ = v26
                case 1:
                    _ = len(v6)
                case 489:
                    v1 = (v1 % ((32767) | 1))
                default:
                    if true {
                        v27 := GenS2{F0: 283, F1: 1, F2: (v2 + (v2 + (v2 + "")))}
                        var v28 float32 = -2.25
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint16 = 255
                        v30 := v3[v2]
                        _ = v29
                        _ = v30
                    }
                }
            }
            if true {
                var v31 uint32 = 610
                v32 := GenS3{F0: 127}
                for fi33 := 0; fi33 < 4; fi33++ {
                    var v34 uint8 = 169
                    _ = len(v6)
                    _ = fi33
                    _ = v34
                }
                _ = v31
                _ = v32
            }
            if true {
                if (((v2 + ("a" + (v2 + v2))) + "a") <= v2) {
                    v35 := GenS2{F0: 0, F1: 42, F2: ("foo bar" + (v2 + "hello"))}
                    _ = v35
                } else {
                    v5 = (v5 * (^(v25)))
                    v1 = v1
                }
                if true {
                    if true {
                        if ((^(v1)) != (v1 >> ((255) & 63))) {
                            var v36 uint64 = 0
                            v14 = (^(v14))
                            if len(v13) > 0 {
                                v13[0] = v5
                            }
                            _ = v36
                        }
                        var v37 uint = (^(v14))
                        _ = v37
                    }
                    delete(v6, (v25 | (^(v25))))
                } else {
                    var v38 bool = (1 < 1000)
                    var v39 int64 = v24
                    var v40 float32 = 1e10
                    _ = v38
                    _ = v39
                    _ = v40
                }
                for fi41 := 0; fi41 < 16; fi41++ {
                    var v42 uint8 = 2
                    v5 -= (-(fi41))
                    switch 128 {
                    case 42:
                        v1 = (v1 | (v1 & (v1 * v1)))
                    case 9223372036854775807:
                        v43 := GenS4{F0: (v1 % ((0) | 1)), F1: 1000}
                        _ = v43
                    case 69:
                        if true {
                            v44 := GenS1{F0: 1, F1: uint8(v1)}
                            v2 = "foo bar"
                            delete(v6, 2147483647)
                            _ = v44
                        } else {
                            v24 = (v24 & v4)
                        }
                    }
                    _ = fi41
                    _ = v42
                }
            } else {
                var v45 string = v2
                _ = v45
            }
        } else {
            var v46 float32 = 1.5
            if true {
                var v47 int16 = -687
                var v48 string = "foo bar"
                for fi49 := 0; fi49 < 11; fi49++ {
                    v14 = 596
                    _ = fi49
                }
                _ = v47
                _ = v48
            } else {
                var v50 bool = true
                _ = v50
            }
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
