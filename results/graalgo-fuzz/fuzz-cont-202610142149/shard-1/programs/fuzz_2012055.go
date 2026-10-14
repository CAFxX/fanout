package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 uint32, p1 uint32, p2 uint) uint {
    v5 := []int{}
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        if false {
            var v6 float32 = 1e10
            var v7 uint = (^(p2))
            p2 -= 42
            _ = v6
            _ = v7
        }
    case "\n\t":
        if false {
            p0 += 2147483648
        } else {
            if (false && (!(!(-615 <= 32767)))) {
                p1 -= p1
                if len(v5) > 0 {
                    v5[0] = 128
                }
            } else {
                if (-79 >= -127) {
                    v8 := GenS1{F0: "\n\t", F1: -350}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v8
                }
                p0 = (p1 * 2)
                var v9 int16 = 32767
                _ = v9
            }
            p2 += (p2 * p2)
            var v10 uint32 = uint32(p2)
            _ = v10
        }
    }
    v5 = append(v5, 849)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 42
    var v3 int64 = 1
    v4 := make(map[string]int)
    v11 := GenS1{F0: "\n\t", F1: v3}
    _ = len(v1)
    var v12 float64 = float64(3.14159)
    switch (v2 - v2) {
    case -888:
        v2 = (^(v2))
    }
    var v13 bool = true
    switch ("\n\t" + (((("hello" + "foo bar") + "\n\t") + "\n\t") + "\n\t")) {
    case "\n\t":
        switch (v2 | (-(v2))) {
        case 0:
            if v13 {
                v3 += (v3 * 1)
                if (((3.14159 * -2.25) - 0.0) != 1.5) {
                    v14 := []int{(v2 << ((4294967295) & 63)), v2, -2147483648}
                    v12 += v12
                    _ = v14
                } else {
                    var v15 uint8 = 205
                    switch "foo bar" {
                    case "":
                        v13 = (v13 && v13)
                    default:
                        v12 += float64(0.1)
                    }
                    var v16 uint32 = 2147483647
                    _ = v15
                    _ = v16
                }
            }
        }
    default:
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 uint32 = 356
                _ = fi18
                _ = v19
            }
            v20 := GenS2{F0: ("hello" + "foo bar"), F1: (683 <= 2), F2: 127}
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 int64 = (-(v3))
                _ = fi21
                _ = v22
            }
            _ = fi17
            _ = v20
        }
    }
    v23 := GenS2{F0: "hello", F1: v13, F2: 68}
    for fi24 := 0; fi24 < 17; fi24++ {
        switch "" {
        case "foo bar":
            if (4294967295 < 65535) {
                fmt.Println(v3, v2, v2)
                fmt.Println(v13)
                fmt.Println(fi24)
            } else {
                v25 := GenS2{F0: "a", F1: (!false), F2: 21}
                _ = v25
            }
        case "a":
            v26 := GenS2{F0: "\n\t", F1: true, F2: -128}
            _ = v26
        default:
            v12 = (float64(-2.25) * v12)
        }
        v27 := []int{v2, v2, v2}
        v3 = -1
        _ = fi24
        _ = v27
    }
    fmt.Println(v2, v13)
    v13 = ((("" < ("\n\t" + ((("\n\t" + ((((("foo bar" + "a") + (("foo bar" + "a") + ("" + "hello"))) + (("" + "a") + "a")) + "\n\t") + ("foo bar" + ""))) + "") + "hello"))) || v13) && v13)
    v28 := GenS2{F0: ((("foo bar" + "") + "a") + "a"), F1: (2147483647 == -573), F2: 51}
    v2 -= v2
    var v29 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
