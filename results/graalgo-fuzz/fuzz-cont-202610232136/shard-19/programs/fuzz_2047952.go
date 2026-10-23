package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }

func main() {
    v1 := []uint32{574, 2147483648, 1, 65535}
    var v2 uint64 = 823
    var v3 uint = 42
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    fmt.Println(v3, v2)
    switch (("a" + ("" + "")) + "hello") {
    case "\n\t":
        var v6 string = "\n\t"
        _ = v6
    case "hello":
        var v7 uint = v3
        _ = v7
    case "":
        var v8 uint = (^(v3))
        _ = v8
    default:
        if ("a" < "") {
            for fi9 := 0; fi9 < 8; fi9++ {
                _ = len(v1)
                fmt.Println(v3)
                _ = fi9
            }
            _ = len(v4)
        } else {
            var v10 int8 = 42
            _ = v10
        }
    }
    var v11 int = 42
    delete(v4, v11)
    v12 := GenS1{F0: int(v2), F1: v3}
    var v13 float64 = 1.5
    v11 = (v11 & v11)
    v13 = (3.14159 + 3.14159)
    var v14 uint16 = 942
    v15 := GenS2{F0: "foo bar", F1: 4294967295, F2: 1}
    v1 = append(v1, 1)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float32 = 1e10
        var v18 int32 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint64 = (v2 - (^(v2)))
        if (!(true && true)) {
            v20 = v20
            var v21 uint16 = (^(v14))
            switch 127 {
            case 32768:
                v14 -= (v14 | uint16(v3))
            case 128:
                fmt.Println(v14, v2, v21)
            case -455:
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS2{F0: "", F1: 0, F2: 273}
                    var v24 uint8 = 2
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                var v25 int = (v11 >> ((v3) & 63))
                _ = v25
            }
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v5[((("a" + ("a" + "a")) + "hello") + "foo bar")] = (((v2 & v2) != (v2 >> ((42) & 63))) && false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
