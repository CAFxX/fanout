package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
func genFn0() bool {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint32 = 2
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 2
    v8 = (v8 / ((v8) | 1))
    switch "foo bar" {
    case "\n\t":
        for fi9 := 0; fi9 < 18; fi9++ {
            v8 = (v8 & v8)
            _ = fi9
        }
    case "hello":
        v10 := []int{}
        _ = v10
    }
    _ = v8
    return (!true)
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (0.1 - 1e10)
    v3 := []int{}
    var v4 uint64 = 0
    var v5 int = -232
    _ = len(v1)
    var v11 bool = (v2 >= 0.1)
    v11 = (!("\n\t" > "foo bar"))
    fmt.Println(v4, v11, v5)
    v12 := GenS1{F0: -65, F1: "foo bar"}
    for fi13 := 0; fi13 < 7; fi13++ {
        if false {
            var v14 int8 = 100
            fmt.Println(v14, v2)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 17; fi15++ {
                fmt.Println(v5, v5, v5)
                var v16 uint = 1
                delete(v1, (("foo bar" + "a") + "\n\t"))
                _ = fi15
                _ = v16
            }
        }
        _ = fi13
    }
    switch v5 {
    case 32767:
        if (179 >= uint8(v4)) {
            switch (v5 * 4294967295) {
            case 641:
                var v17 int32 = 42
                _ = v17
            case 32768:
                var v18 int64 = 1
                _ = v18
            }
            v4 = (v4 & (v4 - v4))
        } else {
            v19 := GenS1{F0: 2147483647, F1: ""}
            if (!v11) {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v2 = v2
                    switch "\n\t" {
                    case "foo bar":
                        v21 := GenS2{F0: "", F1: 1}
                        _ = v21
                    case "":
                        fmt.Println(v2, v2, v4)
                    default:
                        fmt.Println(fi20, v11, v5)
                    }
                    v22 := GenS2{F0: (("a" + "\n\t") + ""), F1: 1}
                    _ = fi20
                    _ = v22
                }
            }
            _ = v19
        }
    case -513:
        var v23 string = "foo bar"
        _ = v23
    case -300:
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS3{F0: 250}
            if len(v3) > 0 {
                v3[0] = v5
            }
            _ = fi24
            _ = v25
        }
    default:
        v26 := GenS4{F0: int16(v4)}
        _ = v26
    }
    v11 = true
    v27 := GenS2{F0: "hello", F1: 42}
    var v28 uint8 = 255
    var v29 bool = v11
    var v30 int16 = 0
    fmt.Println(v2, v30)
    var v31 uint8 = (v28 / ((v28) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
