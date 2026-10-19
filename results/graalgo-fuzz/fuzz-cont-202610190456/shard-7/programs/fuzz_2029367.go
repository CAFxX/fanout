package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
func genFn0(p0 bool) uint32 {
    p0 = p0
    var v7 string = ""
    _ = p0
    _ = v7
    return 870
}
func genFn1() uint {
    var v8 int16 = -32768
    var v9 uint32 = genFn0(((255 >= 32767) && (-9223372036854775808 < 32767)))
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 string = ("hello" + "")
        v11 = "\n\t"
        var v12 int32 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{32767, 9223372036854775807}
    for fi14 := 0; fi14 < 4; fi14++ {
        switch 32768 {
        case -832:
            v9 = v9
        default:
            var v15 float64 = (float64(1e10) / -2.25)
            _ = v15
        }
        v9 += (v9 - (v9 * v9))
        if true {
            if false {
                var v16 uint8 = 164
                v17 := GenS1{F0: 0}
                v8 += 0
                _ = v16
                _ = v17
            }
            if len(v13) > 0 {
                v13[0] = fi14
            }
            fmt.Println(fi14, v9, fi14)
        }
        _ = fi14
    }
    _ = v13
    _ = v8
    _ = v9
    return 281
}

func main() {
    var v1 uint = 255
    v2 := []string{"", ((("a" + ((("a" + "") + "foo bar") + ("hello" + ""))) + "") + ("\n\t" + "hello")), ("\n\t" + "a")}
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v5 int = -489
    var v6 uint8 = 159
    switch v5 {
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -9223372036854775808:
        v5 = (^(v5))
    case -467:
        v1 -= v1
    }
    fmt.Println(v6)
    if true {
        v1 = 4294967295
        for fi18 := 0; fi18 < 11; fi18++ {
            delete(v4, "foo bar")
            _ = fi18
        }
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int = (fi19 * fi19)
            v21 := GenS2{F0: -1, F1: true, F2: genFn0((65535 > uint16(v6)))}
            _ = fi19
            _ = v20
            _ = v21
        }
        v5 = (-(v5))
    }
    var v22 float64 = (float64(0.1) / (((1.5 + -2.25) + 1e10) - (1.5 / -2.25)))
    v23 := GenS1{F0: 791}
    for fi24 := 0; fi24 < 2; fi24++ {
        _ = len(v2)
        _ = fi24
    }
    var v25 bool = (292 < -32768)
    _ = len(v4)
    if (v1 <= (v1 | 0)) {
        v25 = (32767 < 42)
        fmt.Println(v6, v1)
    }
    v5 = v5
    var v26 int64 = int64(v5)
    v1 = 186
    var v27 uint8 = v6
    delete(v3, "a")
    if (!((4294967295 == 2) && false)) {
        var v28 float32 = (1e10 - (1e10 / 1e10))
        v29 := GenS1{F0: 0}
        _ = v28
        _ = v29
    } else {
        v1 = (v1 << ((genFn1()) & 63))
    }
    v4["\n\t"] = ((v5 << ((v1) & 63)) != 0)
    var v30 int = (v5 | (v5 << ((v1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
