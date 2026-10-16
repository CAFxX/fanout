package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 int) uint32 {
    var v6 float64 = -2.25
    v6 -= (((float64((float32(float64(1.5)) / 3.14159)) / float64(3.14159)) * ((1.5 / (3.14159 + v6)) - v6)) - (v6 / (v6 - (float64(1e10) * v6))))
    v7 := []int{(p0 | (p0 | 32767)), 32767}
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint32) bool {
    var v8 uint32 = (^(p0))
    var v9 int8 = 100
    if true {
        var v10 string = (((("hello" + "foo bar") + ("\n\t" + "hello")) + (("a" + (("hello" + "\n\t") + "\n\t")) + "")) + (("" + ("hello" + "\n\t")) + ""))
        fmt.Println(v10)
        var v11 int64 = -773
        _ = v10
        _ = v11
    }
    if (-2147483648 <= 4294967295) {
        v12 := GenS1{F0: (1 > 0)}
        _ = v12
    }
    var v13 int8 = 1
    _ = p0
    _ = v13
    _ = v8
    _ = v9
    return (0.1 <= (1.5 / float32(float64((0.1 / 1.0)))))
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 780
    var v3 int32 = v1
    var v4 string = ("hello" + "")
    v5 := make(map[uint32]bool)
    v14 := []int{128, 0, 1}
    var v15 int8 = 127
    switch v4 {
    case "\n\t":
        if (!true) {
            fmt.Println(v3, v4, v15)
            if false {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            } else {
                v3 = (v1 << ((1) & 63))
                v16 := GenS2{F0: (v2 | genFn0(-1)), F1: ""}
                var v17 float64 = float64((3.14159 + ((((1e10 * 0.1) * 3.14159) * float32(1.5)) / 1.0)))
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                switch (v4 + v4) {
                case "a":
                    delete(v5, (^(v2)))
                case "":
                    v3 = 42
                }
                var v19 uint64 = 202
                if true {
                    var v20 int = fi18
                    _ = v20
                } else {
                    if len(v14) > 0 {
                        v14[0] = (fi18 * fi18)
                    }
                    fmt.Println(v2, v1, v2)
                    _ = len(v14)
                }
                _ = fi18
                _ = v19
            }
        }
    }
    if (!(!(3.14159 <= (0.1 - (0.1 - ((1.5 + -2.25) / 1.0)))))) {
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS2{F0: (v2 - 639), F1: ""}
            _ = fi21
            _ = v22
        }
        v1 -= -2147483648
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            v1 = (-(v1))
            v3 += -2147483648
            _ = fi23
        }
        var v24 int32 = v1
        _ = v24
    }
    v25 := GenS1{F0: (0 < (v15 ^ (v15 - v15)))}
    fmt.Println(v4, v2)
    v14 = append(v14, 127)
    for fi26 := 0; fi26 < 3; fi26++ {
        v1 = 2147483647
        v27 := GenS1{F0: (!false)}
        v5[(v2 >> ((339) & 63))] = (v1 >= v1)
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: ((65535 > 1) && true)}
    var v29 uint8 = 52
    for fi30 := 0; fi30 < 2; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            if false {
                v4 = v4
                v1 = (v1 * (-(v3)))
                v15 -= (v15 & (v15 / (((^(v15))) | 1)))
            } else {
                v32 := GenS2{F0: (v2 + (v2 | v2)), F1: v4}
                _ = v32
            }
            _ = fi31
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
