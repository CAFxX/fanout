package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
func genFn0() int {
    var v6 float64 = 0.1
    fmt.Println(v6)
    _ = v6
    return 127
}
func genFn1() uint16 {
    if true {
        var v7 float32 = 3.14159
        fmt.Println(v7)
        _ = v7
    }
    v8 := GenS1{F0: 42}
    v9 := GenS2{F0: 242, F1: 2}
    var v10 string = "a"
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint64 = 4294967295
        v10 = (v10 + v10)
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v8
    _ = v9
    return 999
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{("foo bar" + "a"), "\n\t", ("hello" + ("\n\t" + ""))}
    var v3 int16 = -32768
    v4 := []uint32{87, 118, 65535, 2147483647}
    v5 := make(map[string]string)
    var v13 int8 = 0
    _ = len(v2)
    v13 = (v13 & (v13 / ((v13) | 1)))
    var v14 float64 = (1.5 / (1e10 - -2.25))
    var v15 uint32 = 2147483647
    var v16 uint64 = 413
    if (2147483647 <= 1) {
        var v17 uint16 = 1
        _ = v17
    } else {
        var v18 int16 = v3
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v3)
            v4 = append(v4, v15)
            _ = fi19
        }
        var v20 uint8 = 1
        _ = v18
        _ = v20
    }
    var v21 float64 = v14
    for fi22 := 0; fi22 < 1; fi22++ {
        if (v3 >= (v3 / ((v3) | 1))) {
            if (!(967 < 1)) {
                for fi23 := 0; fi23 < 2; fi23++ {
                    fmt.Println(v13)
                    if (!((884 > 2) || false)) {
                        if true {
                            var v24 int32 = -1
                            v25 := GenS3{F0: (v1 << ((24) & 63))}
                            fmt.Println(v15, v1, v24)
                            _ = v24
                            _ = v25
                        }
                        fmt.Println(v3)
                        if ("foo bar" != "\n\t") {
                            var v26 int64 = (v1 | v1)
                            if (!false) {
                                var v27 uint64 = 18446744073709551615
                                fmt.Println(v14, v3)
                                _ = v27
                            } else {
                                var v28 float32 = (1e10 / (1e10 + -2.25))
                                _ = v28
                            }
                            _ = v26
                        }
                    }
                    _ = fi23
                }
                v15 -= v15
                v13 = (v13 >> ((2) & 63))
            }
            switch "\n\t" {
            case "\n\t":
                var v29 int64 = 9223372036854775807
                _ = v29
            case "hello":
                if (v21 < 0.1) {
                    var v30 int32 = -248
                    v31 := GenS2{F0: 255, F1: 1}
                    _ = v30
                    _ = v31
                }
            case "foo bar":
                v32 := v5["hello"]
                _ = v32
            }
        }
        if true {
            fmt.Println(v13)
            var v33 float32 = 3.14159
            fmt.Println(v16, v33, v16)
            _ = v33
        }
        _ = fi22
    }
    fmt.Println(v16)
    v34 := GenS4{F0: v16, F1: 0, F2: 395}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
