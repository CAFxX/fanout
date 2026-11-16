package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
func genFn0(p0 bool) string {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{fi4, (fi4 & fi4), fi4, fi4}
        if ((-1 <= 100) || (!((("" + ((("" + ("a" + ("\n\t" + "a"))) + ("foo bar" + ("foo bar" + "foo bar"))) + "")) + ("foo bar" + "a")) == (("\n\t" + ("a" + "hello")) + "")))) {
            var v6 int16 = -452
            v6 = 42
            _ = v6
        } else {
            var v7 uint16 = 32767
            _ = v7
        }
        v8 := GenS1{F0: 0, F1: 2, F2: 42}
        _ = fi4
        _ = v5
        _ = v8
    }
    p0 = p0
    _ = p0
    return "a"
}
func genFn1(p0 uint64, p1 string, p2 uint32) int8 {
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint = 42
        _ = fi9
        _ = v10
    }
    var v11 int = -308
    v12 := GenS2{F0: p2, F1: 65535, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 bool = (false || ((-1 == -69) && true))
    var v2 string = ""
    var v3 int16 = 1
    var v13 uint8 = 128
    v14 := []int{1, -935, 1}
    fmt.Println(v13, v1, v13)
    switch v2 {
    case "a":
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    default:
        v13 -= v13
    }
    v13 -= 28
    _ = len(v14)
    fmt.Println(v1, v3)
    v13 = (v13 | v13)
    if v1 {
        var v15 float64 = (-2.25 - (1e10 * 1.5))
        v13 = v13
        _ = v15
    }
    var v16 uint8 = v13
    var v17 int16 = (v3 & v3)
    var v18 uint16 = 147
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int32 = 0
            v13 += 2
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v22 := GenS1{F0: 0, F1: 0, F2: 42}
    _ = len(v14)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi23
    }
    var v24 uint16 = (v18 + (v18 + v18))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v24
    _ = v3
}
