package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 float32 = float32(-2.25)
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 12; fi6++ {
        if (!(!false)) {
            var v7 uint = 42
            _ = v7
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            switch ("a" + "") {
            case "hello":
                var v9 uint32 = 486
                _ = v9
            case "a":
                v10 := []int{fi6, (fi8 % (((fi6 | -9223372036854775808)) | 1))}
                _ = v10
            }
            _ = fi8
        }
        _ = fi6
    }
    var v11 int64 = 133
    var v12 bool = false
    v11 -= v11
    v13 := []int{869, 32768, 0, -664}
    var v14 int = -9223372036854775808
    var v15 bool = (42 >= 1)
    fmt.Println(v11)
    switch (-(v14)) {
    case -554:
        v3 = -2.25
    case -1:
        v12 = v15
    default:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    var v16 int = (v14 + -1)
    v13 = append(v13, -2147483648)
    for fi17 := 0; fi17 < 18; fi17++ {
        if v12 {
            _ = len(v13)
            var v18 float32 = v3
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint64 = v4
                switch v16 {
                case -1:
                    fmt.Println(v15, v14)
                case 128:
                    v4 = v20
                    var v21 bool = (-456 != (-(fi17)))
                    _ = v21
                case 2147483647:
                    v22 := GenS1{F0: ((1e10 - (0.1 / 3.14159)) != 0.0), F1: 2}
                    _ = v22
                default:
                    var v23 uint = 2
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        _ = fi17
    }
    fmt.Println(v16)
    for fi24 := 0; fi24 < 15; fi24++ {
        if true {
            fmt.Println(v12)
        }
        _ = fi24
    }
    var v25 float32 = v3
    v25 = (float32(0.0) * v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
