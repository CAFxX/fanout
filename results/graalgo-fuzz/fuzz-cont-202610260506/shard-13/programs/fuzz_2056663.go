package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0() uint16 {
    v5 := GenS1{F0: 519, F1: 18446744073709551615}
    var v6 float32 = 3.14159
    var v7 float32 = v6
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string) uint8 {
    if (0 == 0) {
        fmt.Println(p0)
    }
    var v8 uint16 = 2
    p0 = ((p0 + p0) + (("foo bar" + ((p0 + (p0 + "\n\t")) + "\n\t")) + (p0 + "\n\t")))
    switch 32768 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int8 = -1
                v11 -= v11
                var v12 uint8 = 1
                var v13 int64 = 979
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi9
        }
    case 0:
        if (false || true) {
            if (-1 >= 2147483647) {
                v14 := []int{127, -1, 32767, 32767}
                v8 += v8
                _ = v14
            }
            var v15 uint64 = 4294967295
            _ = v15
        } else {
            switch "" {
            case "":
                switch p0 {
                case "hello":
                    var v16 int8 = 1
                    _ = v16
                }
            case "hello":
                var v17 uint64 = 0
                _ = v17
            default:
                p0 = (("" + "foo bar") + "hello")
            }
            var v18 int16 = 881
            for fi19 := 0; fi19 < 2; fi19++ {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v8 = v8
                    fmt.Println(v18, fi19, fi20)
                    var v21 float64 = -2.25
                    _ = fi20
                    _ = v21
                }
                switch (fi19 | (-(fi19))) {
                case 1:
                    p0 = p0
                case -512:
                    if (1 <= 55) {
                        var v22 uint = 494
                        if false {
                            var v23 int32 = 333
                            fmt.Println(p0, v22, v23)
                            _ = v23
                        }
                        v8 = 2
                        _ = v22
                    }
                }
                _ = fi19
            }
            _ = v18
        }
    case 2147483647:
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint32 = 316
            _ = fi24
            _ = v25
        }
    default:
        if true {
            var v26 uint16 = genFn0()
            v8 = v26
            _ = v26
        }
    }
    _ = p0
    _ = v8
    return 149
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 string = (("hello" + "hello") + "foo bar")
    var v27 int8 = 0
    v28 := v3[65]
    var v29 int32 = 0
    delete(v2, 0)
    var v30 int32 = -1
    v31 := []int{110, 4294967295}
    v32 := GenS1{F0: 2, F1: 4294967295}
    switch (((v28 + (v28 + ("hello" + "a"))) + ("a" + v28)) + (("a" + "hello") + v28)) {
    case "":
        v33 := GenS1{F0: 0, F1: 263}
        _ = v33
    case "a":
        if (4294967295 > 255) {
            var v34 uint8 = 95
            if true {
                if len(v31) > 0 {
                    v31[0] = 42
                }
            } else {
                if false {
                    fmt.Println(v1, v34)
                    v27 -= 42
                    v28 = (v4 + "")
                } else {
                    v30 = (v30 / ((1) | 1))
                }
            }
            for fi35 := 0; fi35 < 6; fi35++ {
                if len(v31) > 0 {
                    v31[0] = fi35
                }
                _ = fi35
            }
            _ = v34
        }
    }
    switch v4 {
    case "":
        var v36 uint8 = 75
        _ = v36
    case "a":
        v4 = v28
    default:
        var v37 uint32 = 533
        _ = v37
    }
    var v38 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v4
}
