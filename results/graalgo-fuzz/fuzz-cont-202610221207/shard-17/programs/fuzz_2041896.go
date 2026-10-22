package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0() uint32 {
    var v6 string = (("foo bar" + ("" + "")) + ("" + "foo bar"))
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 1}
        _ = fi7
        _ = v8
    }
    _ = v6
    return 193
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    var v5 uint32 = 1
    v5 -= genFn0()
    v9 := []int{}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float32 = float32((1e10 - (1.5 + (0.0 / 1.0))))
        fmt.Println(v5, v2, v11)
        var v12 int16 = 32767
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = len(v9)
    fmt.Println(v2, v2)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v2, v2, fi13)
        var v14 int = (fi13 + int(v2))
        v14 += -2147483648
        _ = fi13
        _ = v14
    }
    if (true || ((^(v2)) < 952)) {
        v5 -= (v5 + 42)
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
    } else {
        if (!true) {
            if (!true) {
                v15 := GenS2{F0: ("\n\t" + ("hello" + ("\n\t" + (("foo bar" + "\n\t") + ("\n\t" + "\n\t")))))}
                var v16 uint = uint(v2)
                v17 := GenS3{F0: 0, F1: -128}
                _ = v15
                _ = v16
                _ = v17
            }
        }
        v5 = 2147483647
        switch -2147483648 {
        case 127:
            v2 += v2
        case 32768:
            var v18 int16 = 1000
            _ = v18
        default:
            fmt.Println(v5)
        }
    }
    var v19 int32 = 59
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
