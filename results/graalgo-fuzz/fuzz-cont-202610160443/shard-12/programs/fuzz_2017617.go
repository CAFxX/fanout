package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int; F2 int16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint64) int16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 889, F1: 127}
        var v8 uint8 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = (p0 - (p0 + 65535))
    if true {
        v9 := GenS1{F0: 804, F1: 20}
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    return -32768
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = 1
    var v3 uint16 = 255
    var v4 uint8 = 42
    var v5 int64 = int64(v4)
    v11 := []int{489, 128, 127, 88}
    var v12 float32 = (0.1 * 0.1)
    fmt.Println(v5)
    v2 += 42
    if (42 >= -32768) {
        if (-908 < 2147483647) {
            v3 = 0
            var v13 uint8 = v4
            v12 -= 0.1
            _ = v13
        } else {
            if (-32768 >= 42) {
                if (!true) {
                    fmt.Println(v2, v3)
                    var v14 uint32 = v1
                    _ = v14
                } else {
                    var v15 uint = 1
                    v16 := GenS1{F0: v3, F1: -28}
                    _ = v15
                    _ = v16
                }
            } else {
                if len(v11) > 0 {
                    v11[0] = -1
                }
            }
            fmt.Println(v5, v12)
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int16 = 396
            switch fi17 {
            case 128:
                v2 -= v2
            case 647:
                var v19 int = 0
                _ = v19
            case -2147483648:
                if (false || (true || (("foo bar" + "foo bar") <= "\n\t"))) {
                    if true {
                        v18 -= v18
                    } else {
                        v18 -= v18
                        var v20 bool = true
                        var v21 int = fi17
                        _ = v20
                        _ = v21
                    }
                    v5 -= -975
                }
            default:
                var v22 int = fi17
                _ = v22
            }
            var v23 int = 1
            _ = fi17
            _ = v18
            _ = v23
        }
    } else {
        v24 := GenS2{F0: v1}
        if len(v11) > 0 {
            v11[0] = -202
        }
        _ = v24
    }
    v1 -= v1
    for fi25 := 0; fi25 < 8; fi25++ {
        if (841 >= 18446744073709551615) {
            var v26 uint32 = 42
            v27 := GenS3{F0: 18446744073709551615, F1: -2147483648, F2: (v26 >> ((0) & 63))}
            _ = v26
            _ = v27
        } else {
            fmt.Println(v5, v5)
        }
        v28 := GenS4{F0: v4, F1: (fi25 - fi25), F2: genFn0((^(v1)), 18446744073709551615)}
        _ = fi25
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        switch (("a" + ("" + ("a" + "foo bar"))) + ("foo bar" + "hello")) {
        case "foo bar":
            switch "" {
            case "foo bar":
                v3 = (^(v3))
            case "":
                for fi30 := 0; fi30 < 3; fi30++ {
                    v1 -= (^(v1))
                    v2 += v2
                    v31 := GenS5{F0: 251}
                    _ = fi30
                    _ = v31
                }
            case "a":
                var v32 string = ("foo bar" + "hello")
                _ = v32
            }
        }
        _ = fi29
    }
    var v33 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
