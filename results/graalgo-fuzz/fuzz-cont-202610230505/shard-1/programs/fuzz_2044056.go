package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 int) int32 {
    var v6 float32 = 1e10
    var v7 int16 = 0
    v8 := GenS1{F0: (p1 << ((42) & 63))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = (1e10 + (0.0 * ((0.1 / 1.0) - -2.25)))
    var v2 int64 = 317
    var v3 uint = 2
    v4 := make(map[string]string)
    var v5 bool = (v3 > (v3 << ((v3) & 63)))
    for fi9 := 0; fi9 < 5; fi9++ {
        if (!v5) {
            v10 := GenS1{F0: fi9}
            fmt.Println(v5)
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v2 -= v2
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(fi12, fi12, fi11)
            v2 += (v2 >> ((4294967295) & 63))
            var v13 uint8 = 2
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 int = 0
    fmt.Println(v5)
    var v15 float64 = v1
    var v16 uint16 = 65535
    var v17 uint64 = 1
    fmt.Println(v5)
    v5 = (!(v2 >= 9223372036854775807))
    if v5 {
        var v18 string = ("hello" + "a")
        _ = v18
    }
    fmt.Println(v5)
    fmt.Println(v2, v16, v14)
    var v19 string = ("hello" + "")
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 float64 = float64((-2.25 + 1.5))
        if (!false) {
            var v22 int32 = 0
            v3 = (v3 ^ uint(v16))
            var v23 string = ((v19 + "hello") + v19)
            _ = v22
            _ = v23
        } else {
            var v24 float32 = 1e10
            var v25 uint16 = (v16 ^ v16)
            var v26 int = v14
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
