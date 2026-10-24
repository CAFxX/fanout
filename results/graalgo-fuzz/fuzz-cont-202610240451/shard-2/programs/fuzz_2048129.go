package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
func genFn0() uint {
    var v6 int8 = 42
    v7 := GenS1{F0: -593}
    v8 := GenS2{F0: 120}
    v6 = v6
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{-1, 127}
    v2 := []string{("hello" + "\n\t"), "", ("\n\t" + ("" + ""))}
    var v3 float32 = 0.0
    v4 := []int{}
    var v5 int = 127
    _ = len(v2)
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v5, fi9)
            fmt.Println(v3, v5)
            v2 = append(v2, (("hello" + (("foo bar" + ("" + ("\n\t" + ("foo bar" + "a")))) + "hello")) + "a"))
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(v3, v3)
    fmt.Println(v5, v3)
    v11 := GenS1{F0: (v5 ^ v5)}
    var v12 float64 = 3.14159
    var v13 int = (v5 >> ((65535) & 63))
    var v14 float32 = (v3 - float32(1.5))
    if (!true) {
        for fi15 := 0; fi15 < 18; fi15++ {
            v12 = -2.25
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 int32 = -1
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        var v18 uint = 4294967295
        _ = v18
    } else {
        v5 = v13
        v3 += (v14 + float32(float64(v14)))
        v5 += v5
    }
    var v19 uint = 18446744073709551615
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int8 = -47
        _ = len(v4)
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 2
    var v23 bool = ((v19 / ((v19) | 1)) != uint(v22))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
