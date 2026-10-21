package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int, p1 string) uint {
    var v4 int8 = -1
    if (!(0 != 255)) {
        var v5 float32 = 0.0
        _ = v5
    }
    if (0 == (p0 * (p0 ^ -9223372036854775808))) {
        switch (p0 >> ((649) & 63)) {
        case 42:
            v6 := []int{}
            _ = v6
        default:
            for fi7 := 0; fi7 < 7; fi7++ {
                p0 = (fi7 * p0)
                _ = fi7
            }
        }
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v4
    return 662
}
func genFn1() uint {
    var v8 float32 = (3.14159 * 1.5)
    var v9 uint = 42
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 231
    var v3 float64 = float64((-2.25 + 1.5))
    fmt.Println(v2, v2)
    switch -1 {
    case 32768:
        v3 += v3
    default:
        v10 := v1["a"]
        _ = v10
    }
    var v11 float64 = v3
    v1["hello"] = (false || true)
    v11 += 1e10
    v12 := GenS1{F0: 730, F1: 18446744073709551615, F2: ("hello" + (("foo bar" + "") + (("hello" + "\n\t") + "hello")))}
    v13 := GenS2{F0: 0}
    switch ("hello" + "a") {
    case "":
        v14 := GenS2{F0: -1}
        _ = v14
    default:
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(fi15, v3)
            v16 := GenS1{F0: 0, F1: 2, F2: "hello"}
            _ = fi15
            _ = v16
        }
    }
    if false {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 uint64 = 4294967295
            v11 += v11
            v2 -= v2
            _ = fi17
            _ = v18
        }
        var v19 uint64 = 246
        _ = v19
    } else {
        v20 := []int{-2147483648, int(v2), 494}
        if len(v20) > 0 {
            v20[0] = 0
        }
        _ = v20
    }
    v21 := GenS3{F0: 908, F1: genFn1()}
    if (42 < 18446744073709551615) {
        var v22 uint32 = 65535
        _ = v22
    }
    v2 += v2
    _ = len(v1)
    fmt.Println(v11)
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v3)
        var v24 uint = 397
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 897, F1: 229, F2: ("foo bar" + "\n\t")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v25
    _ = v3
}
