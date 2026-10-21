package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int8) string {
    var v5 float64 = (0.0 / 3.14159)
    if false {
        v6 := GenS1{F0: (p0 * -1)}
        p0 = p0
        _ = v6
    }
    v7 := GenS1{F0: (p0 | 1)}
    var v8 int = 9223372036854775807
    v5 = 0.0
    _ = p0
    _ = v5
    _ = v7
    _ = v8
    return ("\n\t" + "\n\t")
}
func genFn1(p0 int16) uint8 {
    p0 += p0
    switch (((("foo bar" + "foo bar") + ("a" + (("foo bar" + "hello") + (("a" + ("foo bar" + "hello")) + "hello")))) + "\n\t") + ("hello" + "foo bar")) {
    case "foo bar":
        fmt.Println(p0)
    case "\n\t":
        var v9 bool = false
        _ = v9
    }
    _ = p0
    return 128
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 55
    var v3 float64 = (1.5 - 0.0)
    v4 := make(map[int]uint32)
    fmt.Println(v2, v2, v2)
    switch 32768 {
    case 384:
        fmt.Println(v3, v1, v3)
    case 269:
        var v10 int32 = 42
        _ = v10
    case 9223372036854775807:
        if (!(false || (((1.5 < ((v3 * float64(-2.25)) * 3.14159)) || true) || (2147483647 < -1)))) {
            v1 += (v1 << ((42) & 63))
            var v11 int16 = 39
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint16 = 560
        _ = fi12
        _ = v13
    }
    var v14 uint32 = (v2 & 0)
    var v15 int16 = 0
    if ((v15 + (v15 ^ v15)) < 1000) {
        v3 = (1.5 - 3.14159)
    }
    var v16 string = "a"
    fmt.Println(v3)
    if (!(0 != 1)) {
        _ = len(v4)
    } else {
        if (0 < -128) {
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := []int{128, fi17}
                switch fi17 {
                case 356:
                    v4[127] = (v2 & (^(v2)))
                case 128:
                    v1 = v1
                }
                _ = fi17
                _ = v18
            }
            var v19 float64 = (1.5 * ((((3.14159 - (v3 + float64(1.5))) / v3) - 3.14159) - v3))
            v19 -= v19
            _ = v19
        }
        v20 := GenS1{F0: 127}
        switch 2147483647 {
        case -2147483648:
            var v21 uint8 = uint8(v2)
            _ = v21
        case 127:
            v4[-9223372036854775808] = 42
        default:
            fmt.Println(v2, v15, v1)
        }
        _ = v20
    }
    var v22 float32 = (0.1 / ((-2.25 - 1e10) * (3.14159 / 1.5)))
    v3 -= float64(v22)
    v23 := []int{-9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
