package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 bool; F1 string }
func genFn0() int32 {
    var v6 string = "hello"
    v6 = v6
    v6 = ((("a" + v6) + "") + "\n\t")
    _ = v6
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{((("" + "hello") + "hello") + "a")}
    var v3 uint8 = 54
    var v4 int64 = 42
    var v5 uint32 = 2
    var v7 string = "foo bar"
    if false {
        switch 287 {
        case 865:
            var v8 float64 = float64(3.14159)
            _ = v8
        case 1:
            var v9 uint16 = 1
            _ = v9
        case 795:
            v5 += 776
        default:
            if false {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 bool = (1e10 > float32(0.1))
                    if (!(v7 > v7)) {
                        var v12 float32 = float32(0.0)
                        _ = v12
                    }
                    v13 := GenS1{F0: v3, F1: -1}
                    _ = fi10
                    _ = v11
                    _ = v13
                }
                v3 += (v3 - v3)
                var v14 uint64 = 1
                _ = v14
            } else {
                fmt.Println(v7, v4)
                v7 = "foo bar"
                if false {
                    v15 := GenS2{F0: 2, F1: 17}
                    _ = v15
                }
            }
        }
        if false {
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint32 = (^(v5))
                if true {
                    if (-1 == 0) {
                        fmt.Println(v17, fi16, v7)
                    }
                    v18 := GenS1{F0: (v3 << ((370) & 63)), F1: -32768}
                    v3 += (v3 | 128)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 int16 = -32768
                        _ = fi19
                        _ = v20
                    }
                    if (((v3 >> ((77) & 63)) >= v3) || false) {
                        if (1.5 >= 3.14159) {
                            for fi21 := 0; fi21 < 1; fi21++ {
                                fmt.Println(v17)
                                v17 -= 42
                                _ = fi21
                            }
                            var v22 uint8 = (v3 % ((v3) | 1))
                            _ = v22
                        }
                        v23 := GenS2{F0: uint(v17), F1: 255}
                        var v24 int64 = v4
                        _ = v23
                        _ = v24
                    }
                    var v25 float64 = (3.14159 - 0.1)
                    _ = v25
                }
                v1[v7] = (v5 & 42)
                _ = fi16
                _ = v17
            }
            v4 += v4
        } else {
            if (-1 <= 1) {
                _ = len(v1)
                var v26 float64 = 3.14159
                var v27 int8 = 42
                _ = v26
                _ = v27
            }
        }
        v7 = "foo bar"
    }
    var v28 uint8 = v3
    v29 := GenS2{F0: 255, F1: 236}
    v1[v7] = v5
    var v30 float64 = 0.1
    var v31 uint32 = 2147483647
    v32 := GenS3{F0: (!true), F1: v7}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
}
