package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int32 = 42
    var v2 int = -262
    var v3 string = ("\n\t" + "\n\t")
    v4 := []int{}
    v5 := []string{(v3 + ((("" + ((v3 + v3) + v3)) + ("" + v3)) + v3)), ""}
    v1 -= (-(v1))
    switch v2 {
    case 0:
        v6 := GenS1{F0: 1, F1: 32767, F2: 1000}
        _ = v6
    case -1:
        var v7 string = v3
        _ = v7
    default:
        v4 = append(v4, (v2 + v2))
    }
    switch 886 {
    case -786:
        if false {
            if true {
                switch (v2 + v2) {
                case 42:
                    fmt.Println(v1)
                case -470:
                    var v8 int8 = -1
                    _ = v8
                case 32767:
                    var v9 int16 = 543
                    _ = v9
                }
            }
            v2 = (v2 * (^(v2)))
            var v10 float32 = 1e10
            _ = v10
        } else {
            v5 = append(v5, ("foo bar" + v3))
        }
    case -749:
        v11 := GenS2{F0: 4294967295}
        _ = v11
    }
    switch v3 {
    case "":
        v1 = v1
    case "foo bar":
        var v12 int16 = 42
        _ = v12
    case "\n\t":
        v2 -= (v2 >> ((4294967295) & 63))
    default:
        v2 = v2
    }
    v13 := GenS3{F0: 65535, F1: 498}
    v4 = append(v4, v2)
    v14 := GenS1{F0: 162, F1: int16(v1), F2: 1000}
    var v15 float32 = (3.14159 - -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
