package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint64, p1 uint64, p2 int8) bool {
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 2; fi7++ {
        if false {
            v8 := GenS1{F0: -731}
            _ = v8
        }
        p0 = 2
        _ = fi7
    }
    var v9 int16 = 1000
    fmt.Println(p2, v9, p1)
    p1 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return true
}

func main() {
    var v1 bool = true
    var v2 int64 = 0
    var v3 int16 = 32767
    v4 := []uint32{}
    var v5 int8 = 1
    var v10 bool = (42 != 4294967295)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint64 = 888
        var v13 int = fi11
        var v14 int = (^(fi11))
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint64 = 17
    var v16 float32 = (float32((3.14159 - 1.5)) - (-2.25 * 0.1))
    v10 = (v2 == (v2 << ((65535) & 63)))
    v1 = v1
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int = (fi17 - (fi17 - fi17))
        v16 = v16
        _ = fi17
        _ = v18
    }
    v3 += v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 2147483647 {
    case -781:
        var v19 uint = 65535
        _ = v19
    case -1:
        v16 = (v16 * v16)
    }
    var v20 int = 32768
    v15 = (v15 % ((v15) | 1))
    v5 = (v5 - int8(v15))
    v1 = ((!(!false)) || v1)
    var v21 uint32 = 365
    switch 0 {
    case 128:
        var v22 int8 = v5
        _ = v22
    default:
        v23 := GenS1{F0: v20}
        _ = v23
    }
    if v10 {
        v5 -= v5
        var v24 int32 = 0
        var v25 int64 = v2
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
