package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool, p1 int64) uint64 {
    fmt.Println(p1, p1)
    var v5 int = 607
    fmt.Println(p0)
    v6 := []int{}
    if p0 {
        p0 = ((v5 - (v5 * (-(v5)))) != v5)
    } else {
        v7 := GenS1{F0: 1, F1: 2}
        v5 -= v5
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 bool = (-405 <= -2147483648)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 uint16 = 658
    v1 = (!(("" + ("a" + "foo bar")) == ("\n\t" + "foo bar")))
    v8 := []int{32767, -580, -185, -65}
    var v9 uint = 255
    v4 += 2
    if v1 {
        fmt.Println(v1)
        if true {
            if ((0 < 18446744073709551615) && (-284 > 0)) {
                if (v1 && true) {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 uint32 = 0
                        _ = fi10
                        _ = v11
                    }
                    v8 = append(v8, 880)
                }
            } else {
                var v12 uint = v9
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 float32 = 3.14159
                v15 := GenS1{F0: 42, F1: v4}
                if (255 != v4) {
                    var v16 string = ""
                    v17 := GenS2{F0: 554}
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            switch ((("hello" + "\n\t") + "hello") + "\n\t") {
            case "hello":
                fmt.Println(v9, v9)
            case "\n\t":
                fmt.Println(v4, v4)
            case "foo bar":
                var v18 float32 = (float32((1.5 * (((1e10 * 1.5) + (1.5 - 3.14159)) / -2.25))) * (1.5 * float32(0.0)))
                _ = v18
            default:
                v9 += 0
            }
        } else {
            var v19 int8 = 42
            _ = v19
        }
        switch 2147483647 {
        case -2147483648:
            var v20 uint64 = 2
            _ = v20
        case -469:
            v4 += v4
        default:
            var v21 uint32 = 2147483647
            _ = v21
        }
    }
    v4 = (v4 ^ v4)
    var v22 uint16 = (v4 % ((v4) | 1))
    for fi23 := 0; fi23 < 20; fi23++ {
        v4 = 314
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := v2["\n\t"]
            v26 := v2[""]
            fmt.Println(v26)
            _ = fi24
            _ = v25
            _ = v26
        }
        if (3.14159 < (0.1 - 1.5)) {
            if (42 > -1) {
                for fi27 := 0; fi27 < 11; fi27++ {
                    v1 = v1
                    var v28 int16 = 32767
                    fmt.Println(v9, v9)
                    _ = fi27
                    _ = v28
                }
            } else {
                v1 = v1
            }
        }
        _ = fi23
    }
    v8 = append(v8, 32768)
    v29 := GenS2{F0: -661}
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 uint8 = 238
        var v32 string = "hello"
        v33 := GenS2{F0: -623}
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
