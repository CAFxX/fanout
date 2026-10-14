package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    var v1 int16 = -784
    var v2 string = (("hello" + (("foo bar" + "foo bar") + "\n\t")) + "\n\t")
    var v3 uint64 = 184
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    switch 1 {
    case 1:
        v1 = -105
    case 9223372036854775807:
        fmt.Println(v3)
    default:
        v6 := GenS1{F0: v2, F1: -1}
        _ = v6
    }
    if ((!(!false)) || (v2 <= "")) {
        v7 := GenS1{F0: v2, F1: 127}
        switch -503 {
        case 9223372036854775807:
            v3 -= v3
        }
        _ = v7
    } else {
        fmt.Println(v2)
        var v8 uint16 = 921
        for fi9 := 0; fi9 < 16; fi9++ {
            delete(v4, (fi9 * (fi9 >> ((255) & 63))))
            _ = fi9
        }
        _ = v8
    }
    v10 := GenS1{F0: (((v2 + v2) + v2) + (v2 + "")), F1: -516}
    delete(v5, 1)
    v11 := GenS2{F0: (v2 + v2)}
    v1 = (v1 - v1)
    v12 := v5[429]
    v13 := GenS2{F0: (v2 + v2)}
    switch 991 {
    case -2147483648:
        v2 = (("" + v2) + "hello")
    case -1:
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint64 = 2
            if (true && (true || true)) {
                v2 = ("hello" + "\n\t")
            } else {
                v16 := []int{fi14}
                v15 += v15
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    default:
        v17 := GenS3{F0: v2, F1: 2147483647}
        _ = v17
    }
    v2 = (("a" + (v2 + (v2 + "a"))) + (((((("hello" + v2) + v2) + (("foo bar" + "a") + ((v2 + v2) + (v2 + "\n\t")))) + v2) + v2) + ("a" + ("" + v2))))
    var v18 int16 = v1
    v19 := GenS4{F0: 65535, F1: v1}
    if (false && (!false)) {
        v3 -= (v3 % (((v3 - (v3 % ((v3) | 1)))) | 1))
    } else {
        var v20 int8 = 111
        fmt.Println(v20, v20)
        v18 = 385
        _ = v20
    }
    if false {
        v18 -= -32768
        var v21 uint64 = v3
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint64 = 4294967295
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
