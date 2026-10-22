package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
func genFn0(p0 string, p1 int32) int8 {
    var v7 int = 907
    p1 -= (p1 % (((^(p1))) | 1))
    v8 := []int{(v7 % (((v7 ^ (v7 - (v7 & v7)))) | 1)), v7}
    for fi9 := 0; fi9 < 19; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        p1 -= p1
        var v10 uint32 = 900
        _ = fi9
        _ = v10
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 35
}

func main() {
    var v1 int32 = 789
    var v2 string = "a"
    var v3 int8 = -1
    v4 := []int{128, 1}
    v5 := make(map[string]string)
    var v6 uint = 4294967295
    if len(v4) > 0 {
        v4[0] = 313
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (v2 + ((((v2 + "a") + (v2 + (v2 + (("\n\t" + (v2 + v2)) + "\n\t")))) + "a") + v2)) {
    case "a":
        var v11 int = 458
        _ = v11
    case "foo bar":
        v2 = v2
    case "hello":
        v12 := GenS1{F0: 1}
        _ = v12
    default:
        switch "" {
        case "\n\t":
            var v13 uint32 = 2
            _ = v13
        default:
            v2 = "\n\t"
        }
    }
    _ = len(v5)
    var v14 bool = (483 < 42)
    v14 = (255 <= 255)
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: 195}
        v17 := GenS1{F0: 128}
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 932, F1: 32767, F2: 98}
    for fi19 := 0; fi19 < 11; fi19++ {
        v6 -= (v6 / ((255) | 1))
        _ = fi19
    }
    _ = len(v5)
    switch v2 {
    case "":
        v2 = "hello"
    case "foo bar":
        if ((v6 ^ v6) <= v6) {
            switch "a" {
            case "foo bar":
                var v20 uint32 = 2147483648
                _ = v20
            }
            _ = len(v5)
        } else {
            if false {
                var v21 float32 = float32(3.14159)
                _ = v21
            }
            var v22 uint16 = 42
            _ = v22
        }
    case "\n\t":
        v14 = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
