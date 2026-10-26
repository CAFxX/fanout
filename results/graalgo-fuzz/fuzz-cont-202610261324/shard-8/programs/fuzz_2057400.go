package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int32, p1 uint) uint8 {
    var v6 int32 = (p0 >> ((p1) & 63))
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 += -263
        _ = fi7
    }
    switch -682 {
    case -92:
        p0 = p0
    default:
        v8 := GenS1{F0: "", F1: -32768}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := []int{}
    v2 := []uint32{2147483647, 0, 397}
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    var v9 int32 = 89
    delete(v4, 127)
    v9 = (v9 % (((v9 & (v9 + (v9 % ((v9) | 1))))) | 1))
    for fi10 := 0; fi10 < 2; fi10++ {
        v9 += (v9 ^ v9)
        _ = fi10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int16 = 42
        _ = fi11
        _ = v12
    }
    delete(v3, (((("foo bar" + ("foo bar" + "a")) + ("\n\t" + "foo bar")) + (("a" + "foo bar") + ("a" + "a"))) + "foo bar"))
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int32 = (v9 / (((v9 ^ v9)) | 1))
        v9 = v14
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: ("foo bar" + ("\n\t" + "")), F1: 314}
    v16 := GenS1{F0: "", F1: -32768}
    v9 = v9
    v17 := GenS2{F0: ((4294967295 != 65535) && (-2147483648 > v9)), F1: 0}
    v9 -= v9
    fmt.Println(v9)
    fmt.Println(v9)
    if false {
        var v18 float64 = (-2.25 / ((1e10 + 0.1) * 1.5))
        _ = v18
    }
    switch "" {
    case "foo bar":
        v19 := GenS3{F0: (-(v9)), F1: 127, F2: 0}
        _ = v19
    default:
        fmt.Println(v9)
    }
    v20 := GenS1{F0: (("a" + "\n\t") + ""), F1: -32768}
    v1 = append(v1, 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
