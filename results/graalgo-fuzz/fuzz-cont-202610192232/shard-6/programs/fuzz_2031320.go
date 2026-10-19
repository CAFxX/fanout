package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 115
    var v3 uint64 = (v2 | v2)
    v4 := []string{}
    for fi5 := 0; fi5 < 3; fi5++ {
        if (fi5 <= -911) {
            var v6 int = (fi5 + (-(fi5)))
            _ = v6
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                v3 -= v3
                _ = fi7
            }
            v2 = (v3 * (v3 % ((259) | 1)))
        }
        _ = fi5
    }
    var v8 uint = 18446744073709551615
    switch "" {
    case "hello":
        var v9 int64 = int64(v8)
        _ = v9
    default:
        var v10 float32 = 1e10
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v3 = (^(v2))
        if false {
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := v1[0]
                if false {
                    var v14 float64 = (0.0 * float64(1e10))
                    if v13 {
                        var v15 float64 = v14
                        var v16 bool = (!v13)
                        _ = v15
                        _ = v16
                    } else {
                        switch (-(fi12)) {
                        case -73:
                            fmt.Println(v8, fi12, v13)
                        default:
                            v17 := GenS1{F0: 1, F1: -128}
                            _ = v17
                        }
                    }
                    _ = v14
                } else {
                    var v18 uint32 = 348
                    if (v13 || (v13 && false)) {
                        v19 := GenS1{F0: -864, F1: 0}
                        var v20 int16 = 93
                        _ = v19
                        _ = v20
                    }
                    switch "" {
                    case "a":
                        var v21 float64 = 3.14159
                        _ = v21
                    case "hello":
                        v18 += v18
                    case "":
                        var v22 uint32 = 65535
                        _ = v22
                    default:
                        v23 := GenS1{F0: 317, F1: 42}
                        _ = v23
                    }
                    switch (fi12 << (((v8 * v8)) & 63)) {
                    case 128:
                        v18 = v18
                        var v24 uint64 = (v2 / ((v3) | 1))
                        fmt.Println(fi11)
                        _ = v24
                    default:
                        v2 = v3
                    }
                    var v25 uint8 = uint8(v8)
                    delete(v1, 65535)
                    _ = v18
                    _ = v25
                }
                _ = fi12
                _ = v13
            }
        }
        _ = fi11
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 bool = ((2 == 42) || true)
        switch (fi26 << ((v8) & 63)) {
        case 128:
            if v27 {
                var v28 int8 = 0
                var v29 int = fi26
                _ = v28
                _ = v29
            } else {
                if v27 {
                    fmt.Println(v3, v2, v2)
                    var v30 uint8 = uint8(v2)
                    _ = v30
                }
                switch 1 {
                case -400:
                    if true {
                        fmt.Println(fi26, v27, v2)
                    }
                }
                v27 = false
            }
        case 32768:
            v8 = 4294967295
        }
        _ = fi26
        _ = v27
    }
    v31 := GenS1{F0: -1, F1: 1}
    switch ("a" + "\n\t") {
    case "":
        fmt.Println(v2)
    default:
        var v32 int8 = -1
        _ = v32
    }
    var v33 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v8
}
