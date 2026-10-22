package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 float32 = (0.1 + 1e10)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v4 := GenS1{F0: 42}
    var v5 float64 = 1.5
    var v6 uint16 = 672
    v6 -= v6
    var v7 int = 2147483647
    switch "foo bar" {
    case "a":
        switch "" {
        case "foo bar":
            fmt.Println(v5)
        default:
            var v8 float32 = float32(((0.0 + (v5 - v5)) * (0.1 - (float64((0.0 - v1)) * v5))))
            _ = v8
        }
    default:
        fmt.Println(v5)
    }
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{(fi9 * fi9)}
            _ = len(v10)
            switch "a" {
            case "":
                switch "hello" {
                case "a":
                    switch "foo bar" {
                    case "hello":
                        if (128 >= 127) {
                            v6 -= v6
                        }
                    case "foo bar":
                        _ = len(v2)
                    default:
                        for fi11 := 0; fi11 < 19; fi11++ {
                            v5 -= 1.5
                            v5 -= v5
                            _ = fi11
                        }
                    }
                case "\n\t":
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 uint32 = 4294967295
                        v13 = v13
                        if false {
                            var v14 int32 = 0
                            v2[v13] = v7
                            v1 += v1
                            var v15 bool = (!(int(v14) != fi9))
                            _ = v14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            case "a":
                var v16 int32 = 564
                _ = v16
            default:
                if true {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        var v18 int8 = -1
                        switch ("a" + "a") {
                        case "a":
                            var v19 float32 = 0.1
                            _ = v19
                        case "\n\t":
                            v20 := GenS2{F0: -819, F1: (v18 << ((452) & 63))}
                            _ = v20
                        case "hello":
                            v1 = v1
                        }
                        _ = fi17
                        _ = v18
                    }
                    var v21 string = "a"
                    var v22 int16 = -32768
                    _ = v21
                    _ = v22
                }
            }
            _ = fi9
            _ = v10
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            _ = len(v2)
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v1, fi24)
                _ = fi24
            }
            v7 += (^(fi23))
            _ = fi23
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS2{F0: -32768, F1: -96}
                fmt.Println(v6, fi25)
                var v28 int64 = 251
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi25
        }
    } else {
        v1 += v1
    }
    if false {
        for fi29 := 0; fi29 < 3; fi29++ {
            v7 = (fi29 * 9223372036854775807)
            var v30 int64 = 1
            fmt.Println(v6)
            _ = fi29
            _ = v30
        }
        var v31 string = "\n\t"
        var v32 int8 = 1
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
