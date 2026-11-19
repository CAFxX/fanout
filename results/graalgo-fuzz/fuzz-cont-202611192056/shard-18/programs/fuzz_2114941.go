package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }
func genFn0(p0 uint16) int16 {
    switch "foo bar" {
    case "hello":
        p0 = p0
    case "foo bar":
        p0 += p0
    }
    v6 := []int{32767, -9223372036854775808}
    p0 += p0
    var v7 uint64 = 1
    _ = p0
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 int8, p1 int64, p2 uint16) uint16 {
    var v8 int = 32767
    if (((float32(1e10) - (1.5 / -2.25)) > float32(0.0)) && (!false)) {
        v9 := GenS1{F0: 1, F1: ("" + ("a" + "hello"))}
        _ = v9
    } else {
        var v10 uint32 = 42
        var v11 float32 = 0.1
        _ = v10
        _ = v11
    }
    if false {
        var v12 int16 = -32768
        switch (((("foo bar" + "hello") + "") + ((("a" + "\n\t") + "hello") + "hello")) + ("\n\t" + (("foo bar" + "hello") + "a"))) {
        case "":
            switch (v8 / (((v8 / ((v8) | 1))) | 1)) {
            case -343:
                var v13 uint8 = 42
                _ = v13
            case 4294967295:
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v12)
                    _ = fi14
                }
            default:
                if (!true) {
                    if (3.14159 != float64(1e10)) {
                        var v15 string = "\n\t"
                        _ = v15
                    }
                    var v16 int32 = -321
                    if false {
                        p0 = -1
                        for fi17 := 0; fi17 < 8; fi17++ {
                            fmt.Println(fi17)
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 uint32 = 65535
                                _ = fi18
                                _ = v19
                            }
                            v20 := GenS1{F0: 42, F1: "a"}
                            _ = fi17
                            _ = v20
                        }
                        fmt.Println(p2)
                    }
                    _ = v16
                } else {
                    var v21 bool = false
                    var v22 uint = 388
                    _ = v21
                    _ = v22
                }
            }
        }
        _ = v12
    }
    var v23 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v8
    return p2
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 0
    v3 := make(map[int]int)
    v4 := []int{0, -241, 9223372036854775807, -2147483648}
    var v5 int16 = 0
    if (4294967295 >= 2) {
        fmt.Println(v2, v5)
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v5 += v5
        v4 = append(v4, 2147483647)
        var v25 string = (("foo bar" + "foo bar") + "a")
        _ = fi24
        _ = v25
    }
    v5 += v5
    var v26 uint64 = 20
    fmt.Println(v2)
    v27 := GenS1{F0: 255, F1: "foo bar"}
    if true {
        if false {
            v28 := GenS2{F0: 2147483648, F1: 65535, F2: 255}
            _ = v28
        }
        if (!(!(!((0.1 == float64((float32(float64(0.0)) - (1e10 * 0.0)))) || true)))) {
            v5 -= (v5 / ((genFn0(32767)) | 1))
            var v29 uint16 = 2
            _ = v29
        }
        fmt.Println(v5, v5, v2)
    }
    v30 := v3[-863]
    fmt.Println(v2)
    v4 = append(v4, v30)
    var v31 int8 = -10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
