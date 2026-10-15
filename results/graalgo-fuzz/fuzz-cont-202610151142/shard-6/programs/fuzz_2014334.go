package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint64 = 887
    v2 := make(map[string]uint32)
    v3 := []uint32{65535, 1, 1, 42}
    var v4 uint8 = 255
    var v5 uint32 = 59
    if (true && (((^(v1)) >= (v1 * 0)) && ((991 != -1) || (592 >= -9223372036854775808)))) {
        v5 = v5
        v4 = v4
    }
    v5 = (^(v5))
    v1 = (v1 << ((42) & 63))
    v5 -= v5
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 = v1
        _ = fi6
    }
    _ = len(v3)
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (fi7 + fi7) {
        case 4294967295:
            var v8 float64 = 0.1
            _ = v8
        default:
            if (fi7 < fi7) {
                v1 += v1
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                var v9 string = "hello"
                _ = v9
            }
        }
        if ((1 >= 0) || (!(true && false))) {
            switch (("hello" + "hello") + "\n\t") {
            case "":
                v2["foo bar"] = (v5 >> ((uint(v5)) & 63))
            case "foo bar":
                v10 := GenS1{F0: 0, F1: 1000}
                _ = v10
            case "a":
                v11 := GenS1{F0: 72, F1: 65535}
                _ = v11
            default:
                var v12 string = "hello"
                _ = v12
            }
            v13 := GenS1{F0: 1, F1: 1}
            v5 -= (v5 >> ((18446744073709551615) & 63))
            _ = v13
        } else {
            var v14 string = ("foo bar" + "hello")
            _ = v14
        }
        _ = fi7
    }
    v4 = 42
    fmt.Println(v1)
    v15 := GenS1{F0: 0, F1: uint16(v1)}
    if (-628 != 42) {
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v1, v4)
            _ = fi16
        }
        v4 -= 158
    }
    var v17 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
