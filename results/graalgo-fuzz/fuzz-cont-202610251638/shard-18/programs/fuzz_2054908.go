package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 int16, p1 string, p2 uint) int32 {
    v6 := []int{}
    p2 = (p2 + 42)
    v7 := GenS1{F0: p0}
    var v8 int64 = 0
    var v9 uint = 846
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    v5 := []uint32{117, 65535, 0, 249}
    v2 = v2
    var v10 uint8 = 0
    if true {
        v10 = (v10 / (((v10 | v10)) | 1))
    }
    switch ("foo bar" + "a") {
    case "":
        var v11 uint64 = (v2 ^ (v2 << ((910) & 63)))
        _ = v11
    case "foo bar":
        var v12 string = "foo bar"
        _ = v12
    default:
        switch (("a" + "a") + "hello") {
        case "hello":
            if false {
                var v13 float64 = -2.25
                _ = v13
            } else {
                v2 -= 18446744073709551615
            }
        }
    }
    v14 := GenS2{F0: v2}
    for fi15 := 0; fi15 < 8; fi15++ {
        switch "" {
        case "hello":
            v16 := GenS2{F0: (v2 ^ v2)}
            _ = v16
        case "foo bar":
            v2 = v2
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            v17 := GenS1{F0: -32768}
            _ = v17
        }
        if (-125 < 248) {
            switch fi15 {
            case -2147483648:
                v18 := GenS1{F0: 1}
                _ = v18
            case 32768:
                v2 += 18446744073709551615
            }
            v3[755] = false
            v10 += (v10 / ((42) | 1))
        } else {
            if (true && ("hello" == "hello")) {
                v19 := GenS1{F0: 1000}
                _ = len(v1)
                if ((3.14159 - -2.25) <= 1e10) {
                    v20 := GenS1{F0: -32768}
                    var v21 bool = false
                    v10 -= v10
                    _ = v20
                    _ = v21
                } else {
                    switch "\n\t" {
                    case "hello":
                        var v22 bool = (0 < 419)
                        _ = v22
                    case "\n\t":
                        var v23 float32 = (-2.25 * 1e10)
                        _ = v23
                    }
                }
                _ = v19
            }
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS1{F0: -1}
            var v26 int8 = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v27 uint8 = 42
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi15
    }
    v28 := GenS3{F0: 4294967295, F1: -1}
    v29 := GenS2{F0: (^(v2))}
    switch 0 {
    case 128:
        switch "foo bar" {
        case "a":
            v5 = append(v5, 4294967295)
        case "hello":
            v30 := GenS4{F0: (v10 ^ 233), F1: false}
            _ = v30
        case "":
            v31 := GenS3{F0: -9223372036854775808, F1: 1000}
            _ = v31
        }
    case 603:
        var v32 int64 = 9223372036854775807
        _ = v32
    default:
        var v33 uint8 = v10
        _ = v33
    }
    v4["hello"] = 4294967295
    if true {
        v34 := GenS5{F0: 290, F1: (^(v10))}
        if false {
            v35 := GenS2{F0: (v2 - 303)}
            _ = v35
        }
        if false {
            _ = len(v3)
            v36 := GenS4{F0: (v10 ^ (^(v10))), F1: false}
            _ = v36
        }
        _ = v34
    } else {
        v5 = append(v5, 963)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
