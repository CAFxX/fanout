package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 bool, p1 int32) uint64 {
    var v7 uint = 651
    v8 := []int{127}
    v9 := GenS1{F0: uint8(v7)}
    var v10 uint8 = 100
    var v11 int64 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = v1
    var v3 int16 = 32767
    var v4 int = -2147483648
    var v5 float32 = 0.0
    v6 := make(map[string]bool)
    v12 := GenS2{F0: int64(v1), F1: v2, F2: 2}
    if false {
        v5 -= (((v5 * float32(0.1)) * 1.5) * (v5 / 1.0))
    } else {
        v3 += -679
        var v13 int64 = 42
        switch ("hello" + ((("" + "hello") + "") + ("\n\t" + ""))) {
        case "hello":
            var v14 uint = 0
            _ = v14
        case "":
            if (false || true) {
                for fi15 := 0; fi15 < 7; fi15++ {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v3 = (v3 + -32768)
                        _ = fi16
                    }
                    v5 = (v5 + ((float32((float64(((((1.5 + 0.0) - float32((-2.25 - 0.0))) + 0.0) - (((((v5 + 1e10) / v5) + (v5 * v5)) + (v5 / v5)) * 3.14159))) * 3.14159)) - v5) - v5))
                    _ = fi15
                }
                v17 := GenS1{F0: 0}
                _ = v17
            }
        }
        _ = v13
    }
    switch v4 {
    case -314:
        var v18 int = -937
        _ = v18
    case -868:
        fmt.Println(v5, v4)
    case -1:
        var v19 bool = true
        _ = v19
    default:
        v20 := GenS3{F0: 42, F1: "a"}
        _ = v20
    }
    fmt.Println(v5, v5, v3)
    v1 -= v2
    if (!((!true) || true)) {
        v5 = ((v5 * (0.1 * (1e10 - v5))) - ((v5 + (((1e10 * float32(-2.25)) + v5) - (v5 - (v5 / v5)))) / v5))
        if false {
            _ = len(v6)
            v21 := GenS2{F0: 42, F1: (v1 - 255), F2: 127}
            v4 = (-(v4))
            _ = v21
        } else {
            var v22 float64 = -2.25
            fmt.Println(v5, v4)
            _ = v22
        }
    } else {
        v23 := GenS4{F0: v4, F1: v4}
        switch (v4 + (v4 - v4)) {
        case 2147483647:
            switch -2147483648 {
            case 2147483647:
                var v24 int32 = -2147483648
                _ = v24
            case 127:
                v4 = 482
            }
        case -1:
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(v3, fi25, fi25)
                fmt.Println(fi25)
                _ = fi25
            }
        case 32768:
            fmt.Println(v5)
        default:
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := v6[("foo bar" + ("foo bar" + (("hello" + "foo bar") + ("" + ("hello" + "hello")))))]
                _ = fi26
                _ = v27
            }
        }
        _ = v23
    }
    var v28 bool = true
    var v29 uint32 = 1
    v6["hello"] = (290 >= 2)
    v30 := []int{4294967295, v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
