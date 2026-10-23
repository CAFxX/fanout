package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0(p0 uint, p1 uint64, p2 uint) string {
    v6 := GenS1{F0: -643, F1: -1}
    v7 := GenS1{F0: -32768, F1: 42}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    v1 := []string{"foo bar", (("foo bar" + "hello") + ""), (("hello" + "") + "foo bar"), "\n\t"}
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := []string{"hello", "\n\t"}
    v5 := make(map[uint32]int)
    var v9 uint64 = 97
    switch ("" + "foo bar") {
    case "hello":
        for fi10 := 0; fi10 < 6; fi10++ {
            v9 = 2
            var v11 int16 = -679
            _ = fi10
            _ = v11
        }
    case "\n\t":
        delete(v3, "a")
    }
    switch 0 {
    case 2147483647:
        if (1 > 0) {
            if ((1e10 * float32(0.1)) != 0.1) {
                fmt.Println(v9)
                switch 4294967295 {
                case 782:
                    for fi12 := 0; fi12 < 17; fi12++ {
                        switch fi12 {
                        case 2147483647:
                            _ = len(v5)
                        default:
                            var v13 float32 = 0.1
                            _ = v13
                        }
                        var v14 string = ("\n\t" + (("a" + "") + "\n\t"))
                        _ = fi12
                        _ = v14
                    }
                case -9223372036854775808:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 uint8 = 255
                        _ = fi15
                        _ = v16
                    }
                case 32767:
                    var v17 bool = (!false)
                    _ = v17
                default:
                    v9 = (v9 * v9)
                }
            }
            var v18 int32 = -394
            v19 := GenS1{F0: 32767, F1: v18}
            _ = v18
            _ = v19
        }
    case 772:
        v20 := GenS1{F0: 0, F1: -2147483648}
        _ = v20
    case -973:
        if (!(("a" + "foo bar") > "foo bar")) {
            for fi21 := 0; fi21 < 18; fi21++ {
                if (!(false || (0 > 127))) {
                    var v22 int = 361
                    _ = v22
                }
                switch (fi21 % (((fi21 % ((fi21) | 1))) | 1)) {
                case -116:
                    fmt.Println(v9)
                }
                _ = fi21
            }
        }
    default:
        var v23 uint32 = 1
        _ = v23
    }
    v2[42] = ("\n\t" + ("hello" + "\n\t"))
    if (((-2.25 - 1.5) < float32((0.0 * (((3.14159 / (float64((0.0 - (1e10 * 0.1))) / 1.0)) * 0.0) + -2.25)))) || (-1 != 0)) {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 uint64 = (v9 - (^(v9)))
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            fmt.Println(v25)
            _ = fi24
            _ = v25
        }
    } else {
        fmt.Println(v9)
    }
    v9 -= v9
    var v26 int8 = -1
    v9 += v9
    var v27 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
