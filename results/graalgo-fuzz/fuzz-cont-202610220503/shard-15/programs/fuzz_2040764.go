package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8; F1 int; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint16) int32 {
    v6 := GenS1{F0: 42, F1: (0.1 >= 1.5)}
    var v7 uint8 = 42
    p0 -= (p0 - p0)
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 uint32, p1 int8) uint8 {
    p1 -= p1
    if false {
        v8 := GenS2{F0: 2, F1: -955, F2: (p0 - p0)}
        p0 = p0
        _ = v8
    } else {
        p1 -= p1
        var v9 float32 = 1.5
        _ = v9
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 float32 = (0.0 + float32((1e10 - 3.14159)))
    v4 := make(map[int]int)
    var v5 uint32 = 65535
    for fi10 := 0; fi10 < 17; fi10++ {
        if (1 >= -1) {
            fmt.Println(fi10, v1, v3)
            switch (^(fi10)) {
            case -728:
                switch fi10 {
                case -1:
                    var v11 float64 = (1.5 / 1e10)
                    _ = v11
                case 32767:
                    fmt.Println(v1, fi10, v1)
                default:
                    var v12 uint8 = 1
                    _ = v12
                }
            case 4294967295:
                v5 -= v5
            }
        }
        _ = fi10
    }
    v13 := GenS2{F0: 127, F1: -518, F2: v5}
    v3 = 3.14159
    v14 := []int{}
    fmt.Println(v3, v3, v5)
    for fi15 := 0; fi15 < 11; fi15++ {
        v3 = v3
        if (255 < 2) {
            if false {
                v1 = v1
                v16 := GenS3{F0: -1, F1: 1}
                _ = v16
            } else {
                var v17 string = ("a" + "a")
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 int64 = 0
                    if false {
                        if (!true) {
                            v1 = (v1 + ((v17 + v1) + ("foo bar" + v1)))
                            v19 = v19
                            _ = len(v14)
                        } else {
                            switch ("" + v1) {
                            case "hello":
                                v1 = ("hello" + v17)
                            case "\n\t":
                                var v20 uint = 4294967295
                                var v21 int = fi18
                                _ = v20
                                _ = v21
                            case "a":
                                switch "hello" {
                                case "":
                                    var v22 uint8 = 2
                                    _ = v22
                                default:
                                    if len(v14) > 0 {
                                        v14[0] = fi18
                                    }
                                }
                            }
                            var v23 float32 = (v3 - 1.5)
                            _ = v23
                        }
                        v19 -= -1
                        v1 = (v17 + ("\n\t" + v17))
                    }
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
        }
        _ = fi15
    }
    fmt.Println(v5, v1, v1)
    v14 = append(v14, 9223372036854775807)
    v5 = 4294967295
    for fi24 := 0; fi24 < 1; fi24++ {
        v5 += v5
        var v25 int64 = 42
        _ = fi24
        _ = v25
    }
    if ((false && (!(4294967295 >= (v5 * v5)))) || false) {
        var v26 int = 4294967295
        delete(v2, "foo bar")
        var v27 int64 = 112
        _ = v26
        _ = v27
    } else {
        var v28 uint16 = 65535
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
