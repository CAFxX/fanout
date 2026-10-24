package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 string }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 int8) int {
    v6 := GenS1{F0: 481, F1: 2}
    v7 := GenS1{F0: 18446744073709551615, F1: 127}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 bool = (108 <= 255)
    var v2 string = "\n\t"
    v3 := []uint32{541, 1, 536}
    v4 := make(map[int]uint32)
    var v5 uint = 42
    var v8 int64 = 0
    var v9 uint8 = 2
    var v10 float64 = 1e10
    switch v2 {
    case "a":
        switch "" {
        case "foo bar":
            var v11 uint64 = 4294967295
            _ = v11
        }
    default:
        v5 = (v5 | 355)
    }
    if false {
        v5 = 840
        if v1 {
            for fi12 := 0; fi12 < 19; fi12++ {
                switch 128 {
                case 32767:
                    v13 := GenS2{F0: 32767, F1: -1, F2: (v2 + "a")}
                    _ = v13
                case 128:
                    var v14 uint8 = v9
                    _ = v14
                case 187:
                    var v15 string = ""
                    _ = v15
                }
                _ = fi12
            }
            fmt.Println(v9, v1, v2)
            v2 = "a"
        }
        var v16 float64 = 0.1
        _ = v16
    }
    v17 := GenS3{F0: 128}
    var v18 uint = 65535
    delete(v4, 42)
    v8 = v8
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int32 = 1
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
