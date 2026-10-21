package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 uint64 = 2
    var v2 int8 = -1
    var v3 int64 = 1
    var v4 uint16 = 255
    var v5 bool = false
    var v6 bool = ((!("\n\t" < "\n\t")) || ((-2.25 * float32(float64(-2.25))) == (-2.25 * ((float32(1.5) - (-2.25 + 0.1)) - ((1e10 / 0.1) * 0.0)))))
    fmt.Println(v3, v1)
    v6 = v6
    for fi7 := 0; fi7 < 9; fi7++ {
        switch "a" {
        case "a":
            var v8 uint32 = 2147483648
            _ = v8
        default:
            v9 := GenS1{F0: (("foo bar" + "a") + ("\n\t" + "a")), F1: v1, F2: 42}
            _ = v9
        }
        switch (fi7 & (-(fi7))) {
        case -2147483648:
            var v10 float32 = 1.5
            _ = v10
        case -9223372036854775808:
            v2 = v2
        case 371:
            v11 := GenS1{F0: "\n\t", F1: (v1 + 421), F2: (v4 * (v4 ^ v4))}
            _ = v11
        default:
            if (int8(v3) < 100) {
                v1 = v1
            }
        }
        v6 = (!(!(!(!true))))
        _ = fi7
    }
    v12 := []int{1, -1, 9223372036854775807, 384}
    var v13 uint = 2
    if false {
        v1 += 1
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint16 = v4
            var v16 int32 = 779
            fmt.Println(v4, v15)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = len(v12)
    }
    var v17 float64 = (1.5 / ((3.14159 - (3.14159 - 1e10)) - (1.5 + -2.25)))
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS1{F0: "foo bar", F1: uint64(v13), F2: 65535}
        v20 := GenS2{F0: -608, F1: v3}
        _ = fi18
        _ = v19
        _ = v20
    }
    switch "hello" {
    case "\n\t":
        var v21 int64 = -30
        _ = v21
    case "hello":
        var v22 uint64 = (v1 % ((18446744073709551615) | 1))
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
