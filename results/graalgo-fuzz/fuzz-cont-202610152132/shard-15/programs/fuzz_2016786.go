package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
func genFn0() uint32 {
    var v7 uint64 = 2
    var v8 uint64 = v7
    v9 := []int{-9223372036854775808, 55}
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 uint64 = 0
    v5 := []string{"\n\t", "a", (("a" + ("hello" + "hello")) + ("hello" + "hello")), "hello"}
    v6 := make(map[int]string)
    var v10 float64 = ((((0.0 * float64(1.5)) / 1.5) * (0.0 + (3.14159 + (-2.25 * float64(float32(0.0)))))) + 3.14159)
    switch ((("foo bar" + "") + ("foo bar" + "")) + ("\n\t" + ("foo bar" + ""))) {
    case "\n\t":
        var v11 int32 = -845
        _ = v11
    case "":
        v4 = (v4 << ((4294967295) & 63))
    default:
        v10 += v10
    }
    if false {
        if false {
            v12 := GenS1{F0: -2147483648, F1: 128, F2: 0}
            _ = len(v5)
            v4 -= v4
            _ = v12
        } else {
            switch 1 {
            case -58:
                v4 -= 2
            case 42:
                _ = len(v5)
            case 32767:
                if false {
                    v13 := GenS1{F0: 90, F1: 128, F2: 42}
                    var v14 float32 = float32(v10)
                    _ = v13
                    _ = v14
                } else {
                    v10 = 1.5
                    v4 -= v4
                }
            }
            fmt.Println(v10, v4)
            if len(v5) > 0 {
                v5[0] = "a"
            }
        }
    } else {
        v15 := v1[4294967295]
        _ = v15
    }
    var v16 int32 = 0
    if false {
        v16 -= (v16 * (v16 % ((2147483647) | 1)))
        var v17 bool = false
        if (0.1 != -2.25) {
            if v17 {
                fmt.Println(v17, v17)
                if (!true) {
                    v17 = (((("a" + ("hello" + "foo bar")) + (("a" + "hello") + ("hello" + "\n\t"))) + "\n\t") <= "foo bar")
                    var v18 int8 = 0
                    _ = v18
                } else {
                    _ = len(v3)
                }
            } else {
                v17 = v17
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            v17 = (0 >= -1)
            if (v17 || false) {
                v5 = append(v5, "foo bar")
            }
            v17 = v17
        }
        _ = v17
    }
    v19 := GenS1{F0: v16, F1: 106, F2: -32768}
    v5 = append(v5, ("" + "a"))
    v20 := GenS1{F0: 0, F1: 128, F2: -623}
    v16 += (v16 / ((v16) | 1))
    for fi21 := 0; fi21 < 1; fi21++ {
        if true {
            fmt.Println(v10)
        } else {
            var v22 int16 = 290
            if ((!true) || ((89 >= 255) && false)) {
                v16 -= v16
                v22 = (v22 << ((2) & 63))
            } else {
                fmt.Println(v10)
            }
            v10 = v10
            _ = v22
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v10)
        _ = fi23
    }
    if (false || true) {
        fmt.Println(v4, v4)
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 bool = false
            _ = fi24
            _ = v25
        }
        v26 := GenS1{F0: (v16 / ((v16) | 1)), F1: 1, F2: 1}
        _ = v26
    } else {
        var v27 uint = 2
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
