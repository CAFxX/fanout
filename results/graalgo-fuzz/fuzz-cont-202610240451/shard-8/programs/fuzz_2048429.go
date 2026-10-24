package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
func genFn0() uint {
    var v5 uint32 = 2
    fmt.Println(v5)
    v5 += v5
    var v6 int8 = -4
    _ = v5
    _ = v6
    return 530
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 128
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint32 = 2147483647
        var v9 uint = 255
        if (false || false) {
            v10 := v1["hello"]
            _ = v10
        } else {
            if (false && (!(521 < int16(v3)))) {
                v3 += v3
            }
            v11 := GenS1{F0: 42}
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 float32 = 1.5
        v3 -= v3
        fmt.Println(v2, v2)
        _ = fi12
        _ = v13
    }
    v2 = v4
    if (false || true) {
        if false {
            v2 = 38
        } else {
            switch -9223372036854775808 {
            case 32768:
                v4 += (v4 * (^(v2)))
            }
            var v14 float32 = 3.14159
            _ = v14
        }
    }
    switch 128 {
    case -9223372036854775808:
        switch (("\n\t" + "") + "a") {
        case "\n\t":
            v3 -= 1
        case "":
            var v15 int8 = 127
            _ = v15
        default:
            v3 = (v3 ^ (v3 & (v3 ^ (v3 << ((18446744073709551615) & 63)))))
        }
    case 127:
        var v16 string = "foo bar"
        _ = v16
    default:
        if (127 <= (^(v2))) {
            v17 := GenS1{F0: int(v2)}
            _ = v17
        }
    }
    v18 := GenS2{F0: 2, F1: 49, F2: uint8(v3)}
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 uint64 = v3
        v20 -= (v3 % ((v3) | 1))
        _ = fi19
        _ = v20
    }
    v4 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
