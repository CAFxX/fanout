package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int64 {
    var v7 int32 = 42
    v7 += v7
    _ = v7
    return 42
}

func main() {
    var v1 uint = 410
    v2 := make(map[uint32]string)
    var v3 string = "foo bar"
    var v4 int32 = -1
    var v5 uint = v1
    var v6 float32 = ((1.5 * 0.1) - 1e10)
    fmt.Println(v6)
    var v8 int8 = -41
    var v9 uint8 = 104
    var v10 uint64 = 2
    var v11 uint8 = (v9 - 42)
    for fi12 := 0; fi12 < 17; fi12++ {
        v4 += (v4 / (((v4 | (^(v4)))) | 1))
        v13 := v2[0]
        _ = fi12
        _ = v13
    }
    if (false || (2 >= (v11 ^ (v9 * v9)))) {
        v14 := GenS1{F0: 712}
        if true {
            var v15 bool = false
            v3 = (("foo bar" + (v3 + (v3 + ("a" + "foo bar")))) + "\n\t")
            _ = v15
        }
        v2[43] = v3
        _ = v14
    } else {
        v16 := GenS1{F0: int16(v8)}
        _ = v16
    }
    var v17 int32 = 1
    var v18 bool = (!false)
    v19 := []int{}
    var v20 float32 = v6
    switch "a" {
    case "":
        var v21 uint64 = (v10 | 0)
        _ = v21
    case "hello":
        v4 -= int32(v9)
    case "foo bar":
        for fi22 := 0; fi22 < 8; fi22++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            v6 -= 1.5
            _ = fi22
        }
    default:
        v18 = (232 > 811)
    }
    var v23 float64 = float64(v20)
    v4 = (v4 >> ((v1) & 63))
    var v24 int64 = int64(v9)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
