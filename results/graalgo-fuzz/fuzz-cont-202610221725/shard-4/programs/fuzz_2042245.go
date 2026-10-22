package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int = 1
    var v2 uint8 = 224
    v3 := make(map[string]int)
    var v4 float32 = 3.14159
    var v5 int16 = 0
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int64 = 1
        switch "hello" {
        case "hello":
            switch ("" + ((("hello" + "\n\t") + "") + "\n\t")) {
            case "a":
                v8 := GenS1{F0: (^(v2))}
                _ = v8
            case "":
                v3["hello"] = (v1 >> ((287) & 63))
            }
        case "":
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 0}
                _ = fi9
                _ = v10
            }
        default:
            var v11 uint64 = 18446744073709551615
            _ = v11
        }
        switch "hello" {
        case "a":
            var v12 int64 = (-(v7))
            _ = v12
        case "":
            v3["\n\t"] = -858
        }
        _ = fi6
        _ = v7
    }
    v13 := GenS2{F0: -1, F1: -906}
    if (42 == 2147483647) {
        switch v1 {
        case 0:
            var v14 float32 = v4
            _ = v14
        default:
            if ((^(v5)) < v5) {
                v1 = v1
                var v15 uint = 4294967295
                fmt.Println(v15, v1, v2)
                _ = v15
            } else {
                if true {
                    v5 = (^(v5))
                } else {
                    v4 -= 0.1
                    v2 = (^(v2))
                    var v16 int = 297
                    _ = v16
                }
                _ = len(v3)
                v17 := GenS3{F0: 182}
                _ = v17
            }
        }
    } else {
        delete(v3, "foo bar")
        v1 -= v1
        switch "" {
        case "hello":
            v3[((("\n\t" + ("hello" + "foo bar")) + "foo bar") + "\n\t")] = (v1 - (^(v1)))
        default:
            v2 = (v2 - v2)
        }
    }
    if ("foo bar" != "") {
        v18 := []int{v1, -17}
        switch "" {
        case "hello":
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(v2)
                _ = fi19
            }
        default:
            v20 := GenS1{F0: 240}
            _ = v20
        }
        _ = v18
    }
    if (("\n\t" + "") >= ("" + (("\n\t" + "hello") + "hello"))) {
        var v21 int8 = 7
        _ = v21
    }
    v1 = (v1 & 1)
    v4 = (v4 / v4)
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := []int{}
            v2 = v2
            _ = fi23
            _ = v24
        }
        v1 += 42
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
