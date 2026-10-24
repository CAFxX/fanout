package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint16) int64 {
    if false {
        v6 := GenS1{F0: p0}
        _ = v6
    }
    p0 -= p0
    switch 1 {
    case -990:
        v7 := []int{0, 0}
        _ = v7
    case -2147483648:
        var v8 int16 = -1
        _ = v8
    }
    _ = p0
    return -1
}
func genFn1(p0 int) uint16 {
    v9 := GenS1{F0: 345}
    v10 := []int{p0, p0}
    p0 -= 129
    v10 = append(v10, 127)
    var v11 uint32 = 1
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 382
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    var v3 int16 = -903
    v4 := make(map[uint32]bool)
    v5 := []string{"\n\t", ("a" + "foo bar"), "a"}
    var v12 int = int(v3)
    v13 := GenS1{F0: 0}
    v12 = v12
    v2 = append(v2, "")
    if ((false || false) && false) {
        var v14 int8 = 42
        v3 -= (v3 | v3)
        _ = v14
    }
    v12 -= v12
    if (!false) {
        switch ((("foo bar" + "a") + "foo bar") + "hello") {
        case "hello":
            var v15 float32 = ((((((1e10 / 1.0) - ((0.1 + 1.5) + 1e10)) - 3.14159) - 1.5) - (0.1 / (1e10 * ((0.1 + 1e10) * 1e10)))) / 1.0)
            _ = v15
        }
    }
    v12 = 183
    if len(v5) > 0 {
        v5[0] = ""
    }
    v3 -= (-(v3))
    if (false || (3.14159 > 0.0)) {
        v12 = v12
        var v16 float32 = ((1.5 / (0.1 - ((1.5 / 3.14159) * (float32(1.5) + (0.1 + (float32(0.1) * 3.14159)))))) - float32(3.14159))
        if (true && (!true)) {
            if (true && (-856 <= 1)) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int32 = 489
                    if true {
                        if false {
                            v19 := GenS2{F0: 32767, F1: 2, F2: 1}
                            var v20 float64 = (0.1 + 1.5)
                            v16 += 0.0
                            _ = v19
                            _ = v20
                        }
                    } else {
                        var v21 float64 = (0.0 + 0.0)
                        _ = v21
                    }
                    _ = len(v2)
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v12, v16, v3)
            } else {
                v22 := GenS3{F0: 686, F1: -1}
                var v23 uint8 = 172
                _ = v22
                _ = v23
            }
            var v24 bool = true
            _ = v24
        }
        _ = v16
    } else {
        if ((!(true || false)) || (true || false)) {
            for fi25 := 0; fi25 < 7; fi25++ {
                v12 = (v12 >> ((864) & 63))
                _ = fi25
            }
            fmt.Println(v3, v12)
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(1 == 1)) {
        v12 = v12
        v26 := GenS1{F0: 42}
        _ = v26
    } else {
        v12 = v12
        v3 = 42
    }
    var v27 bool = (32767 >= genFn1(128))
    v27 = v27
    if v27 {
        v3 -= (v3 - (-(v3)))
        fmt.Println(v3, v3, v12)
        for fi28 := 0; fi28 < 5; fi28++ {
            v12 -= fi28
            v3 = (v3 ^ v3)
            _ = fi28
        }
    } else {
        var v29 uint8 = 205
        v3 = v3
        _ = v29
    }
    v12 = int(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
