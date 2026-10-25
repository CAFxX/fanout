package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool, p1 int64) uint32 {
    v4 := []int{}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 float64 = 0.1
        fmt.Println(p0)
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{128, -518}
    var v3 int8 = 100
    var v8 int = -75
    v8 += 2147483647
    for fi9 := 0; fi9 < 20; fi9++ {
        v8 -= (v8 * v8)
        var v10 uint8 = 154
        _ = fi9
        _ = v10
    }
    v8 = 127
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int32 = 42
        if false {
            v13 := GenS1{F0: 42, F1: v3, F2: 40}
            v3 -= v3
            _ = v13
        }
        v14 := GenS2{F0: (-(v12)), F1: 2147483647, F2: v8}
        _ = fi11
        _ = v12
        _ = v14
    }
    v15 := GenS2{F0: -373, F1: genFn0(true, -947), F2: v8}
    v3 = v3
    var v16 string = ""
    v3 = v3
    var v17 uint32 = 2147483647
    fmt.Println(v16, v3)
    delete(v1, (v16 + ""))
    v18 := GenS1{F0: genFn0(false, 42), F1: v3, F2: v17}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v19 int32 = -2147483648
        v19 = v19
        _ = v19
    } else {
        v20 := GenS3{F0: 2147483647}
        _ = v20
    }
    v16 = v16
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint32 = genFn0((("hello" + ("\n\t" + "hello")) > v16), 462)
        switch v16 {
        case "\n\t":
            switch v8 {
            case 2147483647:
                var v23 float32 = 1.5
                _ = v23
            case 711:
                if (!true) {
                    var v24 uint16 = 1
                    var v25 uint16 = v24
                    var v26 bool = true
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    fmt.Println(v16, v17, v16)
                    v8 -= (^(fi21))
                }
            }
        case "":
            fmt.Println(v16, v17)
        case "foo bar":
            v27 := GenS1{F0: v17, F1: v3, F2: v22}
            _ = v27
        }
        var v28 bool = (false || true)
        _ = fi21
        _ = v22
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
