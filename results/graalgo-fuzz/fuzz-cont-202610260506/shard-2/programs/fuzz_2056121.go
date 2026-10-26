package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }

func main() {
    var v1 uint8 = 246
    var v2 int = 831
    var v3 string = ((("a" + "hello") + "a") + "")
    var v4 float64 = -2.25
    var v5 uint8 = 127
    v4 += v4
    v3 = ((v3 + v3) + ((((v3 + (v3 + v3)) + (v3 + v3)) + v3) + (v3 + ((v3 + (((v3 + "foo bar") + (v3 + (v3 + "hello"))) + "hello")) + v3))))
    for fi6 := 0; fi6 < 1; fi6++ {
        if false {
            v1 = (v1 | v5)
            v1 += v5
            v7 := GenS1{F0: ((787 > 42) && false)}
            _ = v7
        } else {
            var v8 float32 = float32(float64(((float32(float64(float32(v4))) / 0.1) + 0.0)))
            _ = v8
        }
        v9 := GenS2{F0: "foo bar"}
        _ = fi6
        _ = v9
    }
    var v10 int32 = 506
    v11 := []int{(v2 * v2), (v2 | (v2 - int(v1))), v2}
    var v12 uint16 = 255
    var v13 uint = 18446744073709551615
    v11 = append(v11, v2)
    fmt.Println(v12, v10)
    v1 += (^(v1))
    if false {
        v1 = (v5 << ((v13) & 63))
        for fi14 := 0; fi14 < 20; fi14++ {
            v1 -= v5
            v11 = append(v11, 32767)
            var v15 uint64 = 42
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int16 = 1000
            _ = fi16
            _ = v17
        }
    } else {
        var v18 uint64 = 0
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v10)
            _ = fi19
        }
        _ = v18
    }
    v11 = append(v11, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
