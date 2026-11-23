package main

import (
    "fmt"
)

func genFn0() uint8 {
    var v5 int = -600
    fmt.Println(v5)
    v5 = (-(v5))
    v5 = (v5 ^ (v5 << ((18446744073709551615) & 63)))
    _ = v5
    return 127
}

func main() {
    var v1 int64 = 327
    var v2 string = "hello"
    var v3 int = 127
    var v4 uint8 = 42
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int = (v3 << ((921) & 63))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{(fi8 / ((0) | 1)), (fi8 << ((236) & 63))}
        fmt.Println(v4, v1)
        _ = fi8
        _ = v9
    }
    if (1 >= 18446744073709551615) {
        v4 = v4
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v1)
            _ = fi10
        }
        var v11 int32 = 202
        _ = v11
    } else {
        var v12 float32 = (3.14159 + 3.14159)
        _ = v12
    }
    v3 += v3
    var v13 float32 = ((3.14159 * 3.14159) * float32(1e10))
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint64 = 0
        switch v3 {
        case -358:
            v13 -= v13
        case 285:
            fmt.Println(v15, fi14)
        case 127:
            v4 += v4
        }
        _ = fi14
        _ = v15
    }
    v2 = "hello"
    v16 := []int{v3, (v3 + v3), (v3 + v3)}
    for fi17 := 0; fi17 < 20; fi17++ {
        if ((!false) && (1 > 752)) {
            var v18 uint = 255
            _ = v18
        }
        v3 += (-(v3))
        var v19 uint8 = (v4 * 41)
        _ = fi17
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
