package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint; F1 int; F2 uint8 }

func main() {
    var v1 int32 = 679
    v2 := make(map[int]bool)
    var v3 int32 = v1
    delete(v2, -1)
    var v4 int16 = -109
    v5 := GenS1{F0: 1, F1: "hello"}
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 9, F1: ("foo bar" + "")}
        var v8 uint64 = 18446744073709551615
        v9 := []int{(fi6 % (((fi6 - (fi6 / ((fi6) | 1)))) | 1)), (fi6 ^ (fi6 - -9223372036854775808)), -2147483648, fi6}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v4 += 807
        var v11 uint16 = 681
        if false {
            switch (fi10 >> ((129) & 63)) {
            case 432:
                v12 := GenS1{F0: 128, F1: ("" + "foo bar")}
                _ = v12
            case 130:
                fmt.Println(v4, fi10)
            case 32768:
                v13 := GenS1{F0: 127, F1: "a"}
                _ = v13
            default:
                fmt.Println(v3, v3, fi10)
            }
            switch (fi10 & 42) {
            case 32767:
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 float64 = 0.1
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v3, v1)
    v16 := v2[-9223372036854775808]
    for fi17 := 0; fi17 < 6; fi17++ {
        v2[-665] = (316 == 176)
        fmt.Println(v1, fi17, v1)
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(fi17, v4)
            _ = fi18
        }
        _ = fi17
    }
    v19 := []int{32768, 1, 128}
    v3 = (v3 % ((42) | 1))
    v4 += 861
    if (65535 >= 65535) {
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS2{F0: 42, F1: (fi20 & (-(fi20))), F2: 79}
            v22 := GenS2{F0: 1, F1: int(v1), F2: 129}
            _ = fi20
            _ = v21
            _ = v22
        }
        if (0 > 42) {
            if (726 == 2147483647) {
                var v23 float32 = ((3.14159 / float32(0.1)) * ((0.0 * 0.0) - 3.14159))
                _ = v23
            }
            if v16 {
                var v24 int8 = -97
                _ = v24
            } else {
                var v25 uint16 = 282
                switch 41 {
                case 9223372036854775807:
                    v16 = v16
                case 361:
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                case -2147483648:
                    v25 += v25
                }
                var v26 uint64 = 42
                _ = v25
                _ = v26
            }
        }
        if ((42 <= 18446744073709551615) && true) {
            fmt.Println(v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
