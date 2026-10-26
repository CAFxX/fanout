package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int16; F2 uint }
type GenS6 struct { F0 bool }
func genFn0() uint16 {
    v4 := GenS1{F0: 42}
    var v5 uint8 = 152
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 int8, p2 bool) bool {
    v6 := GenS2{F0: 32767, F1: p1}
    switch "hello" {
    case "":
        var v7 int16 = 1
        _ = v7
    case "\n\t":
        v8 := []int{}
        _ = v8
    default:
        v9 := GenS3{F0: 102}
        _ = v9
    }
    v10 := []int{127, 42, 2147483647, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return p2
}

func main() {
    v1 := []int{128, 396, -299}
    v2 := []string{}
    var v3 int64 = 9223372036854775807
    switch "" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = ""
        }
    case "hello":
        if true {
            v3 = (v3 << ((65535) & 63))
            if len(v2) > 0 {
                v2[0] = ""
            }
            var v11 uint16 = 65535
            _ = v11
        }
    }
    v2 = append(v2, ("" + "hello"))
    var v12 int = 590
    v13 := GenS4{F0: true}
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS5{F0: 42, F1: 705, F2: 972}
        fmt.Println(v12, v3)
        v12 = fi14
        _ = fi14
        _ = v15
    }
    switch v12 {
    case 677:
        _ = len(v1)
    case -2147483648:
        if true {
            if false {
                v16 := GenS3{F0: 644}
                for fi17 := 0; fi17 < 10; fi17++ {
                    v18 := GenS5{F0: 715, F1: 42, F2: 42}
                    if (!true) {
                        var v19 float32 = (-2.25 + -2.25)
                        var v20 string = "a"
                        var v21 int16 = 1000
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        v1 = append(v1, (fi17 >> ((255) & 63)))
                    }
                    var v22 int16 = 42
                    _ = fi17
                    _ = v18
                    _ = v22
                }
                var v23 uint32 = 270
                _ = v16
                _ = v23
            }
            v24 := GenS6{F0: false}
            _ = v24
        }
    }
    if (4294967295 == 42) {
        fmt.Println(v12, v12)
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
