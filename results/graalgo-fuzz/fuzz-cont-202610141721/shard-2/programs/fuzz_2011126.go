package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 int; F1 int8; F2 int8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    var v5 uint8 = 1
    v6 := make(map[string]string)
    v5 -= v5
    v5 -= v5
    v7 := []int{42, -923}
    switch "hello" {
    case "foo bar":
        if true {
            var v8 bool = ((42 <= -1) && (455 < 18446744073709551615))
            v9 := GenS1{F0: (v4 & v4)}
            switch "hello" {
            case "":
                switch ((("hello" + "hello") + ((("a" + "hello") + "") + "a")) + "a") {
                case "":
                    var v10 uint8 = v5
                    _ = v10
                default:
                    fmt.Println(v4)
                }
            default:
                if v8 {
                    v11 := GenS2{F0: 4294967295, F1: v1, F2: 76}
                    v1 -= v1
                    v8 = true
                    _ = v11
                }
            }
            _ = v8
            _ = v9
        }
    default:
        fmt.Println(v4, v1, v1)
    }
    var v12 int32 = -61
    var v13 uint8 = v5
    var v14 float64 = 3.14159
    switch -911 {
    case 32767:
        v15 := GenS1{F0: v4}
        _ = v15
    case 4294967295:
        var v16 uint8 = (v13 << ((v1) & 63))
        _ = v16
    case 1:
        var v17 uint64 = 1
        _ = v17
    }
    var v18 int16 = 1
    for fi19 := 0; fi19 < 13; fi19++ {
        v12 += 42
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v20 := GenS3{F0: fi19, F1: 1, F2: -41}
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
