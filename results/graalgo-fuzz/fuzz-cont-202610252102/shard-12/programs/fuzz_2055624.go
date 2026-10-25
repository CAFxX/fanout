package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0() uint16 {
    var v6 uint16 = 1
    v7 := GenS1{F0: 742}
    var v8 uint64 = 661
    _ = v6
    _ = v7
    _ = v8
    return v6
}
func genFn1(p0 string, p1 uint32, p2 bool) bool {
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int32 = 1
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 0
    var v12 uint64 = uint64(v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (462 < -1)
}

func main() {
    v1 := []int{}
    var v2 float32 = (0.1 / (float32(((1e10 + 1e10) + 1e10)) * 1.5))
    var v3 uint64 = 42
    var v4 float64 = (((0.1 / 0.1) - -2.25) / (3.14159 * 1e10))
    var v5 int16 = 1
    v13 := GenS1{F0: 2147483647}
    fmt.Println(v2)
    switch 425 {
    case 128:
        if len(v1) > 0 {
            v1[0] = -1
        }
    case -572:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if true {
            var v14 uint64 = v3
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float64 = ((v4 - 0.0) + (0.0 - v4))
                v5 += (v5 * v5)
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 int8 = 127
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            _ = len(v1)
            _ = v14
        }
    }
    var v19 int64 = 738
    switch 2147483647 {
    case 0:
        v3 = (^(v3))
    }
    if false {
        if (!(v3 == (v3 - v3))) {
            var v20 string = (("a" + "") + ("\n\t" + "a"))
            v21 := GenS2{F0: 2147483647, F1: 1, F2: 4294967295}
            _ = v20
            _ = v21
        }
    }
    v19 -= v19
    if false {
        if (((v3 | v3) > v3) && (v3 == 4294967295)) {
            _ = len(v1)
            var v22 float32 = v2
            v5 += (v5 - (v5 + (v5 / ((1) | 1))))
            _ = v22
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v5)
            switch "\n\t" {
            case "hello":
                v23 := GenS3{F0: 32767}
                _ = v23
            case "a":
                switch 32768 {
                case 128:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v24 int64 = v19
                    _ = v24
                }
            default:
                if (-1 <= (^(v5))) {
                    var v25 string = "foo bar"
                    for fi26 := 0; fi26 < 5; fi26++ {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            _ = len(v1)
                            var v28 int8 = 45
                            v29 := GenS1{F0: 1}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS2{F0: fi26, F1: -1, F2: 641}
                        fmt.Println(v5, v3, v5)
                        _ = fi26
                        _ = v30
                    }
                    fmt.Println(v4)
                    _ = v25
                }
            }
        }
        fmt.Println(v3, v3)
        var v31 int = int(v19)
        _ = v31
    }
    fmt.Println(v19)
    if (!(3.14159 == (float64(v2) - (float64(v2) * v4)))) {
        var v32 uint16 = 255
        _ = v32
    } else {
        var v33 uint = 42
        if true {
            var v34 int = -78
            v1 = append(v1, 839)
            fmt.Println(v3)
            _ = v34
        }
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
