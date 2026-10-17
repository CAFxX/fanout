package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 887
    v3 := make(map[int]int)
    v4 := []int{-806, -186, 127}
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: true, F1: 127}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
        _ = v6
    }
    v2 -= (v2 >> ((0) & 63))
    fmt.Println(v2)
    v7 := GenS1{F0: (!false), F1: 0}
    var v8 uint8 = 0
    v2 -= 2147483647
    switch (("\n\t" + (("\n\t" + "foo bar") + "")) + ("\n\t" + "a")) {
    case "\n\t":
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint8 = v8
            v11 := GenS2{F0: (v10 >> ((2) & 63)), F1: v10, F2: "hello"}
            v12 := GenS3{F0: v10, F1: 127}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    case "a":
        v13 := v1[32768]
        _ = v13
    case "hello":
        v14 := GenS2{F0: 1, F1: (v8 >> ((2) & 63)), F2: "hello"}
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        _ = len(v4)
        if (((fi15 % ((fi15) | 1)) == fi15) || (-2.25 >= 1.5)) {
            v16 := GenS3{F0: (v8 * (v8 - (v8 - (v8 - (v8 % ((v8) | 1)))))), F1: -81}
            fmt.Println(fi15, fi15, v2)
            _ = v16
        } else {
            var v17 int = (fi15 | -2147483648)
            fmt.Println(v17, v8, v17)
            _ = v17
        }
        _ = fi15
    }
    v4 = append(v4, -9223372036854775808)
    if true {
        var v18 float64 = 0.0
        for fi19 := 0; fi19 < 4; fi19++ {
            v8 = v8
            var v20 int64 = 0
            var v21 uint16 = 2
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    } else {
        v22 := GenS4{F0: (^(v2)), F1: 0, F2: (v8 * 2)}
        v8 = v8
        _ = v22
    }
    if true {
        fmt.Println(v2)
        for fi23 := 0; fi23 < 10; fi23++ {
            if false {
                for fi24 := 0; fi24 < 20; fi24++ {
                    v25 := v3[(fi24 << ((255) & 63))]
                    _ = fi24
                    _ = v25
                }
            }
            _ = fi23
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 int32 = -876
        if len(v4) > 0 {
            v4[0] = (fi26 - (fi26 | -572))
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
