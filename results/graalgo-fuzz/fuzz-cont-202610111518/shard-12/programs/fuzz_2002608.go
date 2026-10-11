package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
func genFn0() int32 {
    var v7 int8 = 105
    fmt.Println(v7)
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -2
    var v3 uint64 = 2
    var v4 int16 = -1
    v5 := make(map[string]bool)
    var v6 float64 = ((0.1 + 3.14159) - (1e10 * ((0.1 / (((1e10 / 1.0) / (((3.14159 + float64(3.14159)) - (((((1e10 - -2.25) / 1.5) + 3.14159) / 0.1) * 1e10)) / 1e10)) + (0.1 + 0.0))) + 0.1)))
    var v8 int32 = 42
    v6 = v6
    v9 := GenS1{F0: 42, F1: genFn0(), F2: -1}
    if (false || (77 < 127)) {
        for fi10 := 0; fi10 < 13; fi10++ {
            v8 += v8
            _ = fi10
        }
        if (!(true && false)) {
            var v11 int32 = -486
            fmt.Println(v11, v8, v8)
            _ = v11
        } else {
            fmt.Println(v3)
        }
        fmt.Println(v8, v4, v2)
    }
    fmt.Println(v2, v2)
    for fi12 := 0; fi12 < 4; fi12++ {
        v4 = -32768
        if true {
            v3 = v3
            var v13 uint16 = 2
            _ = v13
        }
        fmt.Println(v3, v4, v2)
        _ = fi12
    }
    switch 42 {
    case -512:
        if false {
            var v14 float32 = 3.14159
            _ = v14
        }
    case 1:
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 float64 = v6
                fmt.Println(v6, v3, v6)
                _ = fi16
                _ = v17
            }
            v3 += 18446744073709551615
            _ = fi15
        }
    case 32767:
        v3 = (v3 * v3)
    }
    var v18 uint16 = 32767
    v19 := GenS1{F0: (v18 | 770), F1: 542, F2: -70}
    v20 := v5[("foo bar" + "hello")]
    var v21 uint64 = v3
    if v20 {
        v22 := []int{2147483647, 4294967295, 967}
        _ = v22
    }
    v4 += (v4 << ((42) & 63))
    delete(v5, ((("" + ("hello" + "foo bar")) + ("hello" + "hello")) + (("\n\t" + "") + "")))
    delete(v5, "hello")
    v8 -= 125
    v23 := GenS1{F0: v18, F1: -328, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
