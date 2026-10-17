package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint64 }
func genFn0(p0 uint8, p1 int8) uint {
    var v7 uint64 = 444
    v8 := GenS1{F0: 2147483647, F1: (p0 ^ p0), F2: v7}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 978
}

func main() {
    var v1 float32 = (float32(1.5) - (-2.25 - 3.14159))
    var v2 uint32 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    v5 := []string{(("hello" + "\n\t") + (("hello" + "a") + "foo bar")), "hello", "a", ("a" + "hello")}
    var v6 float64 = float64((1e10 / (-2.25 + (1e10 + float32(1.5)))))
    var v9 float64 = (((0.0 / v6) / v6) * 0.1)
    v6 = v9
    fmt.Println(v2, v6, v1)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 string = ("foo bar" + "foo bar")
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v2 -= (v2 - (v2 / ((v2) | 1)))
        _ = fi12
    }
    v13 := GenS1{F0: 0, F1: 255, F2: 792}
    var v14 string = ("a" + (("hello" + "foo bar") + "foo bar"))
    v6 -= 1.5
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            v14 = ("foo bar" + (v14 + "a"))
            _ = fi16
        }
        var v17 uint16 = 689
        _ = fi15
        _ = v17
    }
    if (0 != -134) {
        v1 -= (v1 - v1)
        switch 127 {
        case 7:
            v2 = v2
        case 9223372036854775807:
            v18 := v3[798]
            _ = v18
        default:
            if (float64(3.14159) > v6) {
                fmt.Println(v14)
            }
        }
        if false {
            v1 = v1
            v5 = append(v5, (("\n\t" + v14) + v14))
        } else {
            var v19 string = ("foo bar" + (v14 + v14))
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 int16 = 1
                _ = fi20
                _ = v21
            }
            v2 = (^(v2))
            _ = v19
        }
    }
    v22 := v3[v2]
    switch 4294967295 {
    case 128:
        delete(v4, "a")
    case 0:
        v5 = append(v5, "hello")
    }
    var v23 float32 = v1
    var v24 int = 641
    v25 := GenS1{F0: int32(v24), F1: 2, F2: 2}
    v6 = float64(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
