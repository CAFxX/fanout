package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32) bool {
    p0 = 42
    v7 := GenS1{F0: 886}
    var v8 uint8 = 255
    _ = p0
    _ = v7
    _ = v8
    return (3.14159 > (-2.25 * -2.25))
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 int16 = 181
    var v5 float64 = -2.25
    var v6 string = "hello"
    var v9 string = v6
    var v10 uint8 = 94
    var v11 int16 = (v4 ^ (v4 & 32767))
    v6 = v6
    var v12 int64 = 42
    var v13 uint8 = v10
    var v14 int32 = 519
    v13 = 171
    if ((v11 | -788) < (v11 - v11)) {
        v4 = v11
        var v15 float32 = 1e10
        _ = v15
    }
    v6 = (v6 + "hello")
    v16 := []int{}
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 uint32 = 2
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        v14 += v14
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint = 42
        v11 = 42
        fmt.Println(v1, v14)
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 1}
    var v22 uint8 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
