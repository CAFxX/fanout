package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0() uint32 {
    if (!(127 == -538)) {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = 3.14159
            _ = fi7
            _ = v8
        }
        if (!true) {
            switch ("a" + ("" + ("a" + "foo bar"))) {
            case "a":
                var v9 uint = 669
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: 2147483647}
                _ = v10
            default:
                var v11 int64 = 42
                _ = v11
            }
            var v12 int16 = -1
            v13 := GenS2{F0: 724, F1: 559}
            _ = v12
            _ = v13
        }
        var v14 int32 = 42
        _ = v14
    }
    var v15 int64 = -1
    _ = v15
    return 673
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483647
    var v4 uint8 = 2
    var v5 uint = 239
    v6 := []int{32768}
    _ = len(v6)
    if true {
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    switch "" {
    case "\n\t":
        var v17 uint = v5
        _ = v17
    case "foo bar":
        if true {
            if false {
                fmt.Println(v5)
                var v18 uint8 = (v4 | v4)
                var v19 uint = v5
                _ = v18
                _ = v19
            }
            v20 := GenS3{F0: 0}
            var v21 int = 901
            _ = v20
            _ = v21
        } else {
            if true {
                delete(v2, 2147483647)
            } else {
                if false {
                    v3 = v3
                }
                v22 := GenS2{F0: 18446744073709551615, F1: 767}
                _ = v22
            }
        }
    default:
        var v23 int = -1
        _ = v23
    }
    v1 = v1
    fmt.Println(v3, v4, v1)
    fmt.Println(v5, v5, v3)
    for fi24 := 0; fi24 < 17; fi24++ {
        v1 = (((("\n\t" + (v1 + (v1 + ""))) + ((v1 + (((v1 + "hello") + "foo bar") + "a")) + v1)) + ((v1 + (v1 + v1)) + (v1 + ("" + "\n\t")))) + "hello")
        _ = fi24
    }
    var v25 uint = v5
    var v26 int16 = -190
    v27 := v2[32768]
    _ = len(v6)
    var v28 uint16 = 660
    var v29 uint32 = uint32(v28)
    var v30 uint32 = 2
    var v31 uint = (v25 ^ 0)
    v32 := GenS4{F0: v4}
    fmt.Println(v3, v26, v25)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
