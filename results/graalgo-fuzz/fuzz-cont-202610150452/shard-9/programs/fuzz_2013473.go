package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int16; F1 string }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]uint32)
    var v3 string = "a"
    var v4 int16 = -839
    var v5 int32 = 0
    v6 := make(map[string]int)
    switch v3 {
    case "a":
        v1 += (v1 + 610)
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v1)
            v8 := GenS1{F0: -1, F1: 2}
            for fi9 := 0; fi9 < 17; fi9++ {
                v5 = (v5 + (-(v5)))
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v1, v5)
    fmt.Println(v3, v3, v5)
    v10 := v2[42]
    if ((v4 % ((0) | 1)) < (v4 << ((4294967295) & 63))) {
        v11 := GenS2{F0: 42, F1: v4, F2: 1}
        if true {
            v12 := []int{456}
            _ = v12
        } else {
            v13 := []int{}
            var v14 float32 = 1e10
            _ = v13
            _ = v14
        }
        v15 := GenS3{F0: 100, F1: (true && false)}
        _ = v11
        _ = v15
    }
    var v16 uint16 = (v1 << ((782) & 63))
    if (!false) {
        v10 += 658
        fmt.Println(v4)
    }
    fmt.Println(v1, v1, v10)
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS4{F0: v4, F1: "a"}
        var v19 float64 = 3.14159
        switch 0 {
        case 2147483647:
            switch ("a" + ((v3 + ("a" + v3)) + ("a" + (((v3 + "\n\t") + v3) + "")))) {
            case "hello":
                v20 := []int{}
                _ = v20
            case "\n\t":
                var v21 int8 = 100
                _ = v21
            case "":
                fmt.Println(v5, v16)
            }
        case 669:
            fmt.Println(v5, v10)
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    switch -888 {
    case -167:
        v16 = (v1 + v16)
    case 1:
        v4 = 32767
    }
    var v22 int8 = -128
    v4 = v4
    _ = len(v6)
    if (int8(v1) >= 96) {
        v10 -= (v10 >> ((18446744073709551615) & 63))
    }
    v22 -= (v22 >> ((324) & 63))
    v6[("\n\t" + ("foo bar" + v3))] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
