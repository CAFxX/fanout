package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 int32) int32 {
    if false {
        var v5 int32 = p1
        _ = v5
    } else {
        var v6 int8 = (^(p0))
        fmt.Println(p1, p0, p0)
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{fi7, (fi7 % ((fi7) | 1))}
        p0 = p0
        v9 := GenS1{F0: (p0 << ((187) & 63)), F1: -66}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 float32 = 1e10
            fmt.Println(v12, fi11)
            p1 -= (p1 << ((27) & 63))
            _ = fi11
            _ = v12
        }
        fmt.Println(p1)
        _ = fi10
    }
    var v13 int = 127
    _ = p0
    _ = p1
    _ = v13
    return (p1 * p1)
}
func genFn1(p0 int64, p1 int32, p2 int) uint8 {
    var v14 uint16 = 129
    var v15 string = ((("hello" + "a") + ("" + "\n\t")) + "foo bar")
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int16 = 0
        fmt.Println(p1, p2, v17)
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 int8 = 84
    var v2 uint8 = 127
    v3 := make(map[int]string)
    var v4 string = ""
    v18 := GenS2{F0: 570}
    var v19 int = 1
    if (2 <= 0) {
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS1{F0: v1, F1: 42}
            v22 := GenS2{F0: 42}
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            v1 = (v1 & v1)
            var v24 float64 = 0.0
            _ = fi23
            _ = v24
        }
    }
    if false {
        for fi25 := 0; fi25 < 4; fi25++ {
            if (-2.25 != float64(3.14159)) {
                v26 := GenS1{F0: 1, F1: 0}
                _ = v26
            }
            fmt.Println(v1, v2)
            var v27 int16 = 846
            _ = fi25
            _ = v27
        }
        var v28 uint64 = 0
        v29 := []int{(^(v19)), v19}
        _ = v28
        _ = v29
    }
    v19 -= (v19 - (v19 << ((1) & 63)))
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 uint32 = 0
        _ = fi30
        _ = v31
    }
    v32 := GenS1{F0: v1, F1: -32768}
    delete(v3, (v19 - (v19 - (v19 & 128))))
    var v33 uint64 = 4294967295
    var v34 uint8 = (^(v2))
    var v35 float64 = (0.0 * ((3.14159 * (1.5 + 1.5)) - 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
