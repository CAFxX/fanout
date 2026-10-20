package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
func genFn0() uint32 {
    if (("a" + ("" + ("foo bar" + ("\n\t" + "\n\t")))) > (("" + "") + "foo bar")) {
        var v6 uint32 = 65535
        v7 := GenS1{F0: -9223372036854775808}
        _ = v6
        _ = v7
    } else {
        v8 := GenS2{F0: 127}
        var v9 float64 = 0.0
        v10 := []int{0, -77, 128, 41}
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{-2147483648, 2147483647, 4294967295, 200}
    var v12 uint32 = 65535
    fmt.Println(v12)
    _ = v11
    _ = v12
    return (v12 << ((2) & 63))
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 128
    var v4 uint8 = (v3 - (v3 - v3))
    var v5 float64 = 1e10
    v4 += v3
    var v13 float64 = (float64(1.5) + v5)
    switch 439 {
    case -935:
        if ((0 <= -128) || true) {
            if (0.1 < -2.25) {
                v5 = (v5 / (float64(-2.25) * 1e10))
                fmt.Println(v3, v4)
            }
            var v14 float64 = ((v5 + float64((1e10 - ((1.5 + 0.0) - 3.14159)))) / (1e10 + (v13 / (v13 * (((v13 * 0.1) / (v5 * 1.5)) + -2.25)))))
            _ = v14
        } else {
            fmt.Println(v1)
            v15 := []int{-738, 620, 0}
            v3 = v4
            _ = v15
        }
    default:
        fmt.Println(v5, v1)
    }
    var v16 int = 42
    if (0.1 != v5) {
        v17 := GenS1{F0: 9223372036854775807}
        v4 = (v4 << ((255) & 63))
        v5 = 0.0
        _ = v17
    }
    v4 -= v3
    if false {
        v18 := GenS1{F0: 0}
        switch (("\n\t" + v1) + v1) {
        case "hello":
            v4 -= v3
        case "a":
            v19 := GenS1{F0: 9223372036854775807}
            _ = v19
        default:
            for fi20 := 0; fi20 < 3; fi20++ {
                if (1e10 >= 3.14159) {
                    var v21 uint32 = 777
                    v22 := GenS1{F0: -332}
                    _ = v21
                    _ = v22
                }
                switch fi20 {
                case -2147483648:
                    v2[0] = 1
                case -9223372036854775808:
                    v4 = 24
                }
                _ = fi20
            }
        }
        var v23 float64 = ((1.5 - -2.25) / (1e10 + 0.1))
        _ = v18
        _ = v23
    } else {
        for fi24 := 0; fi24 < 20; fi24++ {
            v13 += ((v5 * v13) + v5)
            fmt.Println(v4)
            _ = fi24
        }
    }
    v13 = -2.25
    v5 += (v5 / (3.14159 + v5))
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
