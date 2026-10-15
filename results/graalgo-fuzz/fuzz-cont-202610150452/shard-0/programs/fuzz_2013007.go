package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 int64) uint64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        p0 -= 655
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 bool = false
            _ = fi5
            _ = v6
        }
        v7 := []int{127, fi4, (fi4 + fi4), fi4}
        _ = fi4
        _ = v7
    }
    v8 := GenS1{F0: -32768, F1: "a"}
    _ = p0
    _ = v8
    return 1
}

func main() {
    var v1 int16 = -1
    var v2 int = 32767
    var v3 uint8 = 128
    v9 := GenS1{F0: 42, F1: "a"}
    v2 = 9223372036854775807
    var v10 uint32 = 1
    switch "a" {
    case "hello":
        v11 := GenS1{F0: (v1 | v1), F1: ("a" + "")}
        _ = v11
    case "\n\t":
        var v12 uint32 = v10
        _ = v12
    case "":
        v13 := GenS2{F0: 65535}
        _ = v13
    default:
        v14 := []int{-494, (v2 + v2)}
        _ = v14
    }
    fmt.Println(v1, v10)
    v15 := []int{v2, (v2 >> ((4294967295) & 63))}
    if false {
        fmt.Println(v2, v2)
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int64 = 296
            v18 := GenS1{F0: int16(v3), F1: ""}
            _ = fi16
            _ = v17
            _ = v18
        }
        v2 -= 41
    } else {
        _ = len(v15)
    }
    v19 := GenS3{F0: 42, F1: (-(v1))}
    var v20 uint8 = (v3 >> ((uint(v3)) & 63))
    fmt.Println(v1, v3)
    var v21 uint16 = 42
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(fi22)
        for fi23 := 0; fi23 < 6; fi23++ {
            switch (((("hello" + (("hello" + "hello") + (("\n\t" + ("a" + "hello")) + "foo bar"))) + "\n\t") + "hello") + "a") {
            case "":
                switch ("a" + "a") {
                case "hello":
                    v1 += v1
                default:
                    v15 = append(v15, 40)
                }
            case "\n\t":
                switch int(v21) {
                case 132:
                    switch "hello" {
                    case "\n\t":
                        _ = len(v15)
                    case "foo bar":
                        var v24 int8 = -1
                        _ = v24
                    }
                default:
                    var v25 uint = 2
                    _ = v25
                }
            case "hello":
                _ = len(v15)
            default:
                var v26 int = fi23
                _ = v26
            }
            _ = fi23
        }
        fmt.Println(v20)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
