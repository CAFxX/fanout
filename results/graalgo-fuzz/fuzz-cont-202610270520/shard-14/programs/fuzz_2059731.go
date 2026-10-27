package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int16 = 42
    var v2 uint = 18446744073709551615
    v3 := make(map[string]string)
    var v4 float64 = ((0.1 * 3.14159) + 0.0)
    v5 := []uint32{42, 65535}
    v6 := []int{-1, 0}
    var v7 uint = (^(v2))
    var v8 int = 9223372036854775807
    v9 := GenS1{F0: 127}
    var v10 int32 = 617
    v11 := GenS2{F0: 127, F1: (v1 | (-(v1)))}
    _ = len(v6)
    for fi12 := 0; fi12 < 9; fi12++ {
        delete(v3, "foo bar")
        _ = fi12
    }
    if (false || false) {
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi13
        }
    } else {
        v14 := GenS3{F0: ((0.1 + 3.14159) == 0.1)}
        _ = v14
    }
    v3["hello"] = ""
    _ = len(v5)
    var v15 bool = ((1 > 2) || (-2.25 != (1.5 - ((0.0 + 0.0) + -2.25))))
    if v15 {
        delete(v3, "\n\t")
        v16 := GenS4{F0: (^(v8)), F1: 819}
        _ = v16
    } else {
        v15 = ((42 >= 42) || ((^(v8)) != (v8 ^ v8)))
        for fi17 := 0; fi17 < 13; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS4{F0: 2147483647, F1: 761}
                _ = fi18
                _ = v19
            }
            var v20 int16 = v1
            v21 := GenS5{F0: 24}
            _ = fi17
            _ = v20
            _ = v21
        }
    }
    var v22 int32 = (^(v10))
    _ = len(v5)
    var v23 int32 = v10
    var v24 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
