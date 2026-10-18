package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int; F1 uint32; F2 int }
type GenS5 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint8 = 137
    v2 := make(map[string]string)
    var v3 float64 = (1e10 + 3.14159)
    var v4 bool = true
    var v5 int32 = 1
    v2["hello"] = "foo bar"
    var v6 uint32 = 65535
    v3 = float64(-2.25)
    v7 := []int{32768, 9223372036854775807}
    v8 := GenS1{F0: 127, F1: 1, F2: false}
    switch 32768 {
    case -2147483648:
        var v9 string = "foo bar"
        _ = v9
    case 2147483647:
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 float32 = 1e10
            v12 := GenS2{F0: 1, F1: v4, F2: v1}
            var v13 uint16 = 65535
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 2; fi14++ {
            if v4 {
                v3 = (v3 / v3)
                v15 := GenS3{F0: v5, F1: 752}
                _ = v15
            } else {
                v16 := GenS4{F0: (fi14 >> ((255) & 63)), F1: v6, F2: fi14}
                _ = v16
            }
            v4 = false
            _ = fi14
        }
    }
    if (("a" + "foo bar") > "foo bar") {
        switch 32768 {
        case -750:
            if v4 {
                var v17 bool = (false || true)
                switch "a" {
                case "a":
                    var v18 int8 = 42
                    _ = v18
                default:
                    var v19 int = 4294967295
                    _ = v19
                }
                _ = v17
            } else {
                fmt.Println(v3)
                var v20 bool = (((-303 != 127) && v4) && v4)
                if true {
                    var v21 float32 = 1.5
                    var v22 uint64 = 4294967295
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS5{F0: 32767, F1: 42}
                    _ = v23
                }
                _ = v20
            }
        }
    } else {
        if v4 {
            v1 = (v1 & v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
