package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0() int32 {
    var v7 uint8 = 42
    v8 := []int{}
    var v9 uint32 = 418
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v10 uint8 = 42
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint32) uint32 {
    switch 1 {
    case 32768:
        p0 += (p0 % ((42) | 1))
    case 9223372036854775807:
        p0 -= (p0 * p0)
    case -85:
        for fi11 := 0; fi11 < 20; fi11++ {
            p0 = 2
            var v12 uint16 = 42
            p0 += 1
            _ = fi11
            _ = v12
        }
    }
    var v13 uint64 = 780
    if ((((p0 << ((684) & 63)) != p0) || (3.14159 < ((1.5 + -2.25) - float64(0.1)))) || false) {
        v13 -= 555
    }
    _ = p0
    _ = v13
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{0}
    var v3 string = "foo bar"
    var v4 float32 = 0.0
    var v5 uint32 = 42
    var v6 float64 = float64(float32((1e10 * (-2.25 - -2.25))))
    switch (v3 + ((v3 + "a") + v3)) {
    case "foo bar":
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS1{F0: 2}
            _ = fi14
            _ = v15
        }
    case "a":
        v4 -= (float32(1e10) - (-2.25 + float32(0.1)))
    }
    if false {
        v16 := GenS2{F0: 275, F1: genFn1(v5)}
        var v17 uint = 4294967295
        _ = v16
        _ = v17
    }
    if (-9223372036854775808 != -9223372036854775808) {
        v6 = v6
        delete(v1, "foo bar")
        fmt.Println(v6, v5, v3)
    } else {
        v18 := GenS1{F0: 157}
        var v19 uint16 = 0
        _ = v18
        _ = v19
    }
    v5 += (v5 >> ((4294967295) & 63))
    v4 = v4
    v4 = v4
    v20 := v1[v3]
    v20 = v3
    fmt.Println(v5, v20, v5)
    var v21 uint = 891
    v22 := GenS1{F0: 255}
    var v23 uint = (v21 ^ (v21 | (v21 * v21)))
    var v24 string = (v20 + ("hello" + "hello"))
    if (float32(0.1) <= v4) {
        fmt.Println(v5, v4, v4)
    } else {
        switch v24 {
        case "a":
            if len(v2) > 0 {
                v2[0] = (v5 / ((genFn1((^(v5)))) | 1))
            }
        default:
            _ = len(v1)
        }
        if false {
            v25 := GenS2{F0: -1, F1: v5}
            switch v20 {
            case "a":
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            }
            fmt.Println(v23)
            _ = v25
        }
    }
    v21 = 1
    fmt.Println(v3)
    var v26 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
