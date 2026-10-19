package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42, 4294967295, 456}
    var v3 uint64 = 4294967295
    var v4 float64 = ((-2.25 / float64(1.5)) / 1e10)
    var v5 uint64 = (v3 % (((v3 / (((v3 % ((875) | 1))) | 1))) | 1))
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi6)
            _ = fi7
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            v3 = (v5 | (^(v3)))
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS2{F0: -41}
    fmt.Println(v5)
    v3 = (v3 * (v5 ^ v3))
    _ = len(v2)
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v3, v5)
            var v13 int16 = 32767
            var v14 uint64 = v3
            v5 -= 18446744073709551615
            var v15 bool = (false && (65535 != 670))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        v3 = (v3 & 4294967295)
        var v16 int = fi11
        _ = fi11
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: -2147483648}
        fmt.Println(v5)
        switch -9223372036854775808 {
        case 127:
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 int8 = -1
                var v21 string = ""
                _ = fi19
                _ = v20
                _ = v21
            }
        case 42:
            if (!(v4 >= ((0.0 - v4) - (v4 + float64(float32(float64(float32(1.5)))))))) {
                v5 = (^(v5))
                for fi22 := 0; fi22 < 4; fi22++ {
                    v5 -= v3
                    v5 = v5
                    _ = fi22
                }
            } else {
                var v23 float64 = float64(0.0)
                v4 += 1.5
                _ = len(v1)
                _ = v23
            }
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v5, v3)
    switch ("" + "") {
    case "hello":
        if (-758 <= 42) {
            var v24 int8 = 124
            _ = v24
        }
    case "foo bar":
        switch 127 {
        case 42:
            delete(v1, 330)
        case 4294967295:
            switch (("" + ("" + ("\n\t" + ("hello" + "a")))) + "a") {
            case "":
                v3 = 42
            default:
                v25 := GenS2{F0: -1}
                _ = v25
            }
        }
    case "a":
        v5 = (^(v3))
    default:
        v4 = (v4 - v4)
    }
    v4 -= ((v4 * 0.1) / (3.14159 - v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
