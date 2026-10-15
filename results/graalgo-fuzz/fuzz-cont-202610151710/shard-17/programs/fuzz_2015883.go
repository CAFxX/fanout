package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0() int64 {
    var v4 int64 = 648
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(v4)
        v6 := GenS1{F0: 1, F1: v4, F2: 248}
        v7 := GenS1{F0: 1, F1: (v4 << ((255) & 63)), F2: 255}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v4)
    var v8 uint32 = 42
    var v9 uint = 42
    _ = v4
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint8 = 128
    v2 := []int{1}
    v3 := []string{"foo bar", ("foo bar" + ("a" + ("a" + ("hello" + ("hello" + (("" + "a") + ""))))))}
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int16 = 1
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint8 = (v1 ^ 2)
            fmt.Println(fi12, fi12)
            var v14 uint = 70
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    var v15 uint16 = 413
    for fi16 := 0; fi16 < 15; fi16++ {
        if (false || (3.14159 != 3.14159)) {
            v17 := GenS1{F0: 989, F1: 871, F2: (v1 >> ((482) & 63))}
            if (true || false) {
                _ = len(v3)
            }
            var v18 uint8 = 1
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        if true {
            v20 := GenS2{F0: 42}
            _ = v20
        }
        _ = fi19
    }
    v1 -= 42
    var v21 int8 = -128
    fmt.Println(v21, v1)
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int64 = 9223372036854775807
        if (float64(0.0) >= 3.14159) {
            if (!true) {
                var v24 int32 = 0
                var v25 bool = (float64(0.0) != (0.0 * (0.0 / ((1e10 + 0.0) * -2.25))))
                var v26 int64 = (v23 + v23)
                _ = v24
                _ = v25
                _ = v26
            } else {
                v27 := GenS1{F0: 127, F1: v23, F2: (v1 ^ v1)}
                v3 = append(v3, (("\n\t" + "foo bar") + (("hello" + "foo bar") + "")))
                _ = v27
            }
            var v28 int8 = -72
            var v29 float64 = ((1e10 * -2.25) - (float64((((1e10 + 0.1) - 0.0) + 3.14159)) - (-2.25 * 0.1)))
            _ = v28
            _ = v29
        } else {
            var v30 int64 = v23
            _ = v30
        }
        _ = fi22
        _ = v23
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "")
    }
    v31 := GenS3{F0: 2147483647, F1: v1}
    v1 = v1
    fmt.Println(v21, v15, v1)
    var v32 float32 = (0.1 - 3.14159)
    var v33 int64 = 1
    v34 := GenS4{F0: (v15 | v15), F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
