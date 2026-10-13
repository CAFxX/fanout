package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint }
type GenS3 struct { F0 string; F1 int16 }
func genFn0() uint {
    var v6 string = ""
    var v7 float32 = 1e10
    fmt.Println(v6, v6)
    var v8 uint64 = 4294967295
    v9 := GenS1{F0: -2147483648, F1: 308}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]string)
    var v3 int = 32767
    v4 := make(map[string]bool)
    var v5 uint32 = 65535
    var v10 uint16 = 148
    v11 := v2[v3]
    var v12 bool = false
    for fi13 := 0; fi13 < 8; fi13++ {
        v5 = v5
        var v14 float32 = 1.5
        var v15 float32 = (v14 * v14)
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v12, v1)
    switch (v3 << ((42) & 63)) {
    case -1:
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 float32 = (1e10 / (float32(0.1) / 1.5))
                switch ("" + ("a" + (v11 + v11))) {
                case "\n\t":
                    var v19 uint = 18446744073709551615
                    _ = v19
                case "a":
                    var v20 string = v11
                    _ = v20
                default:
                    v21 := GenS2{F0: ("\n\t" + "\n\t"), F1: -2147483648, F2: 267}
                    _ = v21
                }
                v5 += (v5 - v5)
                var v22 uint32 = (v5 | 2)
                v23 := []int{-2147483648}
                v11 = v1
                _ = fi17
                _ = v18
                _ = v22
                _ = v23
            }
            _ = fi16
        }
    case 16:
        v24 := GenS1{F0: (v3 * (v3 ^ v3)), F1: 2}
        _ = v24
    }
    v25 := GenS1{F0: 282, F1: 218}
    var v26 int64 = 42
    fmt.Println(v3)
    v27 := GenS3{F0: (v11 + v1), F1: -493}
    switch "hello" {
    case "foo bar":
        if true {
            v28 := []int{v3, 0, (v3 + (v3 >> ((4294967295) & 63)))}
            v26 = v26
            _ = v28
        } else {
            var v29 uint32 = v5
            v5 = v5
            delete(v4, v1)
            _ = v29
        }
    case "a":
        _ = len(v4)
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 uint64 = 0
        var v32 float64 = ((float64(float32(1.5)) / 0.1) * (0.0 + -2.25))
        v33 := []int{4294967295}
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    var v34 int = (v3 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
