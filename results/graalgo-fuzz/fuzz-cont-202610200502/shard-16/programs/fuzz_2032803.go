package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint64) int8 {
    if (true || (!(p0 && false))) {
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(p1, p1)
            _ = fi6
        }
        v7 := []int{-725, 32767, -1, -9223372036854775808}
        _ = v7
    }
    v8 := GenS1{F0: 0}
    var v9 float64 = 3.14159
    v10 := []int{-1, 0}
    switch 0 {
    case 127:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case -245:
        if len(v10) > 0 {
            v10[0] = 42
        }
    default:
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 int64 = 9223372036854775807
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 int32 = 0
    var v3 int8 = -1
    v4 := make(map[string]string)
    v5 := []string{"hello", "hello", "\n\t"}
    fmt.Println(v2, v1)
    var v14 uint8 = 255
    var v15 uint8 = 42
    var v16 bool = true
    v4["hello"] = "hello"
    if (959 == 42) {
        fmt.Println(v2, v14)
        v17 := GenS2{F0: (^(v1))}
        v3 = v3
        _ = v17
    } else {
        var v18 uint8 = 2
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := v4[("" + "hello")]
            var v21 float32 = 3.14159
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    v22 := GenS1{F0: v1}
    var v23 uint32 = (^(v1))
    var v24 uint8 = (v14 | (v15 << ((951) & 63)))
    v25 := v4["\n\t"]
    v16 = (false && v16)
    switch 127 {
    case 1:
        v26 := v4[""]
        _ = v26
    case 127:
        var v27 uint64 = 42
        _ = v27
    default:
        _ = len(v4)
    }
    v28 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
