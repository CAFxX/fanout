package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1000
    var v3 bool = (!(!false))
    v1 += (v1 ^ v1)
    var v4 float32 = 1e10
    switch "" {
    case "\n\t":
        v5 := []int{0}
        _ = v5
    case "foo bar":
        var v6 uint32 = 250
        _ = v6
    case "a":
        v7 := []int{2147483647, -524, -646, 32767}
        _ = v7
    default:
        v8 := GenS1{F0: 10, F1: "a"}
        _ = v8
    }
    switch ("a" + ("foo bar" + "foo bar")) {
    case "":
        if true {
            switch -1 {
            case 0:
                v9 := []int{128}
                _ = v9
            case 4294967295:
                v10 := GenS2{F0: 660, F1: (v2 / ((v2) | 1)), F2: -405}
                _ = v10
            default:
                v11 := GenS3{F0: 1, F1: 0, F2: v2}
                _ = v11
            }
            switch 128 {
            case 9223372036854775807:
                for fi12 := 0; fi12 < 17; fi12++ {
                    v13 := []int{}
                    v1 = (v1 - (v1 * (v1 >> ((4294967295) & 63))))
                    _ = fi12
                    _ = v13
                }
            case 32767:
                fmt.Println(v2, v3)
            default:
                var v14 uint8 = 255
                _ = v14
            }
            if v3 {
                fmt.Println(v1)
                var v15 int32 = -2147483648
                v2 = v2
                _ = v15
            }
        } else {
            v3 = v3
        }
    case "hello":
        var v16 int8 = int8(v2)
        _ = v16
    case "a":
        v17 := GenS2{F0: (v1 >> ((260) & 63)), F1: (v2 % ((v2) | 1)), F2: 42}
        _ = v17
    }
    v18 := []int{651, -1, 42, -9223372036854775808}
    v3 = v3
    v2 -= (-(v2))
    fmt.Println(v1, v3)
    var v19 uint16 = 1
    var v20 int16 = (^(v2))
    v21 := GenS4{F0: (v3 || (!v3))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
