package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint32 }
func genFn0() uint32 {
    var v5 uint64 = 42
    v5 = v5
    if ((1 != 0) || false) {
        v5 = (v5 / ((v5) | 1))
    }
    if false {
        var v6 bool = (3.14159 != (1e10 * (float32(float64(3.14159)) + float32(1.5))))
        var v7 int32 = 606
        v6 = ((!false) && (2147483648 < 868))
        _ = v6
        _ = v7
    }
    _ = v5
    return 2147483648
}

func main() {
    var v1 int8 = -29
    var v2 int16 = 32767
    var v3 int16 = v2
    v4 := make(map[string]bool)
    v1 += (v1 | (^(v1)))
    var v8 uint8 = 128
    fmt.Println(v8, v8, v8)
    v9 := GenS1{F0: (-(v3)), F1: 0, F2: (3.14159 >= 0.1)}
    if ((0.0 / (float32((3.14159 * ((0.1 + 3.14159) / 0.1))) / ((3.14159 * 0.1) - 0.0))) < (1.5 - 0.1)) {
        if (!true) {
            v10 := []int{2147483647, 4294967295, -872}
            switch ("hello" + ("a" + ("a" + ""))) {
            case "hello":
                var v11 int = 32767
                _ = v11
            case "\n\t":
                var v12 bool = true
                _ = v12
            case "foo bar":
                fmt.Println(v2, v1)
            default:
                v8 = (v8 / (((v8 >> ((255) & 63))) | 1))
            }
            if false {
                if false {
                    delete(v4, "a")
                }
                _ = len(v10)
                v10 = append(v10, 32767)
            }
            _ = v10
        }
        var v13 float32 = 3.14159
        v13 -= (v13 + v13)
        _ = v13
    }
    switch ("\n\t" + (("foo bar" + "hello") + "\n\t")) {
    case "\n\t":
        fmt.Println(v1, v2, v8)
    case "foo bar":
        if (446 == 255) {
            v2 += (v3 >> ((815) & 63))
            for fi14 := 0; fi14 < 18; fi14++ {
                v2 -= 1
                v15 := GenS1{F0: v3, F1: 255, F2: false}
                v16 := GenS2{F0: (fi14 + fi14), F1: uint64(v8)}
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            delete(v4, ("foo bar" + "hello"))
        }
    case "":
        v2 = (v3 << ((65535) & 63))
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 float32 = -2.25
        v19 := GenS2{F0: (^(fi17)), F1: 42}
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        fmt.Println(v2, fi20, v8)
        if ((v2 << ((200) & 63)) > 0) {
            fmt.Println(v2, v3, v2)
        } else {
            if true {
                v21 := v4[("foo bar" + "")]
                _ = v21
            } else {
                for fi22 := 0; fi22 < 17; fi22++ {
                    v8 -= v8
                    v23 := GenS1{F0: 831, F1: 0, F2: true}
                    _ = fi22
                    _ = v23
                }
            }
        }
        var v24 int16 = -964
        _ = fi20
        _ = v24
    }
    var v25 float64 = float64((0.0 + -2.25))
    var v26 uint32 = 65535
    v27 := GenS3{F0: 32767, F1: 127, F2: v26}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
