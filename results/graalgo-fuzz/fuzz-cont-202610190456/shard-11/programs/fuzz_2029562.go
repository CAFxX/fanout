package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 bool; F2 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 2147483648
    v4 := make(map[int]bool)
    var v5 int64 = 1
    if true {
        var v6 uint8 = 210
        _ = v6
    } else {
        var v7 uint = v1
        v8 := []int{32767}
        v9 := v2[(v3 / (((^(v3))) | 1))]
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int64 = 42
    fmt.Println(v5)
    v11 := []int{128, -119, int(v3)}
    var v12 uint16 = 1
    v13 := GenS1{F0: 0, F1: 42, F2: ("" + "foo bar")}
    var v14 uint16 = (v12 / ((32767) | 1))
    if ("hello" < "") {
        for fi15 := 0; fi15 < 14; fi15++ {
            switch ("" + "\n\t") {
            case "foo bar":
                v16 := GenS2{F0: 2, F1: ((!("" == "a")) && false), F2: 2}
                _ = v16
            }
            fmt.Println(v10, v10)
            _ = fi15
        }
        v10 = v10
        v14 = (v14 + v14)
    }
    v17 := GenS2{F0: 18446744073709551615, F1: false, F2: 255}
    var v18 uint8 = 128
    if true {
        var v19 int16 = 32767
        v1 = (v1 * (v1 + v1))
        _ = v19
    } else {
        var v20 uint16 = (v12 << ((1) & 63))
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 float32 = 3.14159
            switch "hello" {
            case "foo bar":
                var v23 uint8 = v18
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v24 int8 = 51
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
