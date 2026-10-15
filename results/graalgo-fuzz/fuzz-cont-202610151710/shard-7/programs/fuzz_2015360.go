package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 uint = uint(v2)
    var v4 uint32 = v2
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 0
        var v8 int8 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    v1["\n\t"] = 2
    var v9 float32 = (-2.25 - (1.5 * (1e10 - (1.5 + 1e10))))
    switch -164 {
    case 0:
        var v10 float64 = 1e10
        _ = v10
    case -634:
        fmt.Println(v5)
    case -753:
        v4 = 0
    default:
        v11 := []int{42, 127, 353}
        _ = v11
    }
    switch 0 {
    case -453:
        _ = len(v1)
    case -2147483648:
        switch "hello" {
        case "foo bar":
            if (0 > -2147483648) {
                v3 += v3
            }
        }
    case 9223372036854775807:
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
    default:
        fmt.Println(v3, v4, v4)
    }
    delete(v1, ("" + ("" + "")))
    var v14 float64 = 1e10
    v15 := GenS1{F0: "hello"}
    var v16 float64 = ((float64(v9) - ((-2.25 + v14) * v14)) - v14)
    v17 := v1[""]
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v3)
        v20 := GenS1{F0: "a"}
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: "\n\t"}
    v22 := []int{-1}
    var v23 int = 2147483647
    var v24 int = 127
    v9 -= v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
