package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    v1 := []uint32{65535}
    var v2 int = 840
    v3 := make(map[int]int)
    var v4 int8 = 1
    var v5 float64 = (1e10 - 3.14159)
    v6 := GenS1{F0: (v4 + v4), F1: 18446744073709551615}
    var v7 int16 = -1
    var v8 uint32 = 2147483647
    fmt.Println(v8, v4)
    if true {
        var v9 int16 = -1
        if (float32(float64(float32(((v5 / v5) - v5)))) != ((3.14159 / -2.25) / (float32((1.5 / 3.14159)) / (1e10 / (0.0 + (1e10 - -2.25)))))) {
            v4 = v4
            for fi10 := 0; fi10 < 9; fi10++ {
                if true {
                    _ = len(v1)
                } else {
                    var v11 float64 = float64(0.1)
                    _ = v11
                }
                var v12 string = ("a" + "")
                v12 = v12
                _ = fi10
                _ = v12
            }
        } else {
            switch "a" {
            case "a":
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 int8 = v4
                    _ = fi13
                    _ = v14
                }
            case "foo bar":
                switch -2147483648 {
                case 4294967295:
                    v15 := GenS1{F0: v4, F1: 4294967295}
                    _ = v15
                case -536:
                    var v16 uint = 2
                    _ = v16
                case 1:
                    v5 -= 1.5
                }
            }
        }
        _ = v9
    } else {
        fmt.Println(v4, v8, v7)
        v2 = (v2 % ((-9223372036854775808) | 1))
        fmt.Println(v5)
    }
    _ = len(v1)
    if (int16(v2) >= v7) {
        var v17 int64 = -1
        v7 = v7
        v18 := GenS2{F0: 2, F1: v7}
        _ = v17
        _ = v18
    }
    switch ("" + "hello") {
    case "":
        v4 = v4
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := v3[v2]
        v7 = 256
        var v21 int = 4294967295
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint8 = 54
    var v23 string = ("\n\t" + "hello")
    v24 := GenS1{F0: (^(v4)), F1: 0}
    fmt.Println(v4, v23, v5)
    if true {
        v8 = v8
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
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
}
