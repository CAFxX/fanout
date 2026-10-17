package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint32) int {
    for fi4 := 0; fi4 < 15; fi4++ {
        fmt.Println(fi4)
        v5 := GenS1{F0: 42}
        _ = fi4
        _ = v5
    }
    p0 -= p0
    p0 = p0
    for fi6 := 0; fi6 < 7; fi6++ {
        if (42 <= 2) {
            var v7 int16 = -900
            var v8 int8 = 1
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: -1}
            p0 = p0
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 uint8 = 42
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    return -461
}

func main() {
    var v1 bool = (2 >= 0)
    var v2 uint64 = 429
    var v3 int = 32768
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS1{F0: 0}
        v2 += 403
        v2 -= v2
        _ = fi12
        _ = v13
    }
    v14 := []int{}
    v15 := GenS1{F0: -869}
    if false {
        v14 = append(v14, v3)
        var v16 string = "hello"
        _ = v16
    }
    var v17 uint64 = (v2 | 907)
    v3 -= (v3 << ((4294967295) & 63))
    if (!(0.0 <= (-2.25 + 3.14159))) {
        var v18 float32 = 1e10
        if false {
            var v19 int32 = 1
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v19
        } else {
            if len(v14) > 0 {
                v14[0] = (v3 >> ((65535) & 63))
            }
            fmt.Println(v3, v3, v2)
        }
        _ = v18
    } else {
        var v20 uint = 4294967295
        _ = v20
    }
    if len(v14) > 0 {
        v14[0] = (v3 | (v3 / ((v3) | 1)))
    }
    if (1.5 != 3.14159) {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    } else {
        v21 := GenS1{F0: -765}
        fmt.Println(v1, v17, v17)
        v17 += (v17 & 42)
        _ = v21
    }
    switch -1 {
    case 42:
        var v22 int16 = 60
        _ = v22
    }
    fmt.Println(v3)
    var v23 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
