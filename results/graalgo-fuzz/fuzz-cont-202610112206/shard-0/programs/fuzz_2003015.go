package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int16 {
    var v5 int64 = 0
    v5 = (v5 | -9223372036854775808)
    v5 -= (v5 - v5)
    _ = v5
    return 32767
}

func main() {
    var v1 uint8 = 127
    var v2 int64 = 9223372036854775807
    var v3 int32 = 0
    var v4 uint = 1
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 4; fi7++ {
        switch fi7 {
        case -9223372036854775808:
            var v8 uint8 = (v1 % (((v1 | 0)) | 1))
            _ = v8
        case 9223372036854775807:
            switch (fi7 / ((fi7) | 1)) {
            case 127:
                v9 := []int{(fi7 ^ fi7), (fi7 * (fi7 + 1)), 9223372036854775807, (fi7 ^ (fi7 & (fi7 + fi7)))}
                _ = v9
            case 32767:
                var v10 int32 = (v3 * (v3 | v3))
                _ = v10
            }
        }
        v11 := []int{42}
        v3 += 794
        _ = fi7
        _ = v11
    }
    if true {
        var v12 int8 = -128
        _ = v12
    }
    var v13 float32 = 1.5
    fmt.Println(v13)
    v14 := []int{32767, 9223372036854775807}
    if (false || false) {
        v13 = (v13 - float32(0.0))
        var v15 int8 = -89
        _ = len(v14)
        _ = v15
    } else {
        v3 = v3
    }
    v16 := GenS1{F0: (false && (!(!true)))}
    v13 = (3.14159 + float32((0.0 - (((3.14159 + 0.1) / 1.5) - (1.5 + (0.0 + 0.1))))))
    v4 += (v4 >> (((^(v4))) & 63))
    var v17 uint16 = 529
    fmt.Println(v2, v1, v6)
    if ((18446744073709551615 != 18446744073709551615) || false) {
        v17 -= v17
        switch ("a" + ("" + "a")) {
        case "foo bar":
            v3 -= 42
        case "":
            if len(v14) > 0 {
                v14[0] = 4294967295
            }
        case "a":
            v3 = v3
        }
    } else {
        v4 += (^(v4))
    }
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
