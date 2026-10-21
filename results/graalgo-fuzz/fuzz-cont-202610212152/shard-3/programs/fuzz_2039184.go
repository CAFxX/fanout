package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = 0.1
    var v2 int32 = -1
    var v3 string = "hello"
    var v4 float64 = float64(((v1 - 1.5) * (v1 / (v1 * (v1 / v1)))))
    v5 := make(map[string]bool)
    v6 := []string{v3}
    v4 += v4
    var v7 uint32 = 2
    if true {
        switch 0 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 15; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    v2 = int32(fi8)
                    v3 = v3
                    v1 -= v1
                    _ = fi9
                }
                var v10 int8 = 32
                _ = fi8
                _ = v10
            }
        case 128:
            v11 := GenS1{F0: true, F1: 65535, F2: (("hello" + v3) + "foo bar")}
            _ = v11
        }
        fmt.Println(v1, v3, v1)
    } else {
        if true {
            if true {
                v1 = 1.5
                if true {
                    v4 = 3.14159
                    for fi12 := 0; fi12 < 20; fi12++ {
                        switch v3 {
                        case "a":
                            if false {
                                v13 := GenS2{F0: 520}
                                v14 := GenS2{F0: 0}
                                v6 = append(v6, (v3 + ("" + "a")))
                                _ = v13
                                _ = v14
                            }
                        default:
                            v4 = (float64(v1) / v4)
                        }
                        if true {
                            var v15 int = fi12
                            _ = len(v6)
                            _ = v15
                        }
                        var v16 bool = true
                        _ = fi12
                        _ = v16
                    }
                } else {
                    fmt.Println(v1, v1)
                }
                var v17 int16 = -710
                _ = v17
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if (false || true) {
                            v2 = 2147483647
                        }
                        v2 += v2
                        var v20 uint64 = 42
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
            var v21 int = 32768
            _ = v21
        }
    }
    v4 = v4
    var v22 float32 = (v1 / v1)
    v23 := GenS1{F0: false, F1: 1, F2: v3}
    switch -2147483648 {
    case 580:
        var v24 int8 = -1
        _ = v24
    }
    v25 := GenS2{F0: 1}
    var v26 int64 = -745
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS1{F0: (true && false), F1: 32767, F2: (("foo bar" + ("\n\t" + (("\n\t" + v3) + ""))) + "")}
        if false {
            v5[v3] = (!(73 == 1))
            v7 -= (v7 % (((v7 ^ v7)) | 1))
        } else {
            v29 := GenS3{F0: (v2 | (^(v2)))}
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    v30 := GenS3{F0: v2}
    switch "a" {
    case "foo bar":
        var v31 float64 = float64((1.5 * float32((-2.25 + (v4 + v4)))))
        _ = v31
    case "a":
        v22 = ((1e10 * float32((v4 + -2.25))) + 1e10)
    }
    _ = len(v6)
    fmt.Println(v26, v4)
    v32 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
