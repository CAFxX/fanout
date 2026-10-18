package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 0
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    _ = len(v1)
    if true {
        v5 := GenS1{F0: 233}
        var v6 uint16 = 42
        _ = v5
        _ = v6
    }
    fmt.Println(v2)
    if ("" >= "a") {
        fmt.Println(v2)
        switch ("" + "hello") {
        case "a":
            v2 = v2
        case "hello":
            fmt.Println(v2)
        case "foo bar":
            v7 := GenS2{F0: 9223372036854775807}
            _ = v7
        }
        v2 = (v2 & -9223372036854775808)
    } else {
        v2 = (^(v2))
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                if false {
                    v10 := GenS2{F0: (^(fi8))}
                    _ = v10
                } else {
                    v11 := []int{fi8, (^(fi8))}
                    fmt.Println(fi9, fi9)
                    v12 := GenS2{F0: (fi9 << ((1) & 63))}
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
            switch fi8 {
            case 1:
                if (65535 >= 1) {
                    var v13 uint = 18446744073709551615
                    var v14 int = -2147483648
                    _ = v13
                    _ = v14
                }
            case 32767:
                v15 := GenS3{F0: ""}
                _ = v15
            case 2147483647:
                for fi16 := 0; fi16 < 4; fi16++ {
                    switch ((((("a" + ("a" + "\n\t")) + "\n\t") + (("a" + "") + "a")) + "a") + ("a" + ("foo bar" + "a"))) {
                    case "foo bar":
                        var v17 int8 = -1
                        _ = v17
                    case "":
                        delete(v1, -835)
                        v2 = (v2 * (v2 | -9223372036854775808))
                        var v18 int8 = -112
                        _ = v18
                    case "a":
                        var v19 float32 = 0.1
                        _ = v19
                    }
                    var v20 string = ("a" + "a")
                    _ = fi16
                    _ = v20
                }
            default:
                v2 = (v2 + -1)
            }
            _ = fi8
        }
        var v21 float64 = (1.5 - (1e10 - (0.0 / 0.1)))
        _ = v21
    }
    v2 = (v2 | 42)
    v2 = 1
    v2 = 164
    var v22 uint = 65535
    v2 = (v2 ^ (v2 ^ v2))
    v1[32767] = -405
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
