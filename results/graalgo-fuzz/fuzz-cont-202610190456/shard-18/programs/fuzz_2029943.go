package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []int{0, 2147483647, 42, 32767}
    v2 := []string{("\n\t" + "\n\t"), ((("\n\t" + (("" + "a") + ("foo bar" + "foo bar"))) + "hello") + "\n\t"), "\n\t", "foo bar"}
    v3 := make(map[uint32]string)
    var v4 uint16 = 1
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if (!false) {
        for fi5 := 0; fi5 < 1; fi5++ {
            if (0 != 255) {
                for fi6 := 0; fi6 < 6; fi6++ {
                    switch (fi6 >> ((306) & 63)) {
                    case 1:
                        v7 := GenS1{F0: 1, F1: 4294967295}
                        _ = v7
                    case -658:
                        var v8 float32 = (1e10 - ((1.5 - 0.1) * 1e10))
                        _ = v8
                    }
                    _ = fi6
                }
                if true {
                    v9 := GenS1{F0: 689, F1: 422}
                    v4 = 209
                    delete(v3, 65535)
                    _ = v9
                }
                v10 := GenS2{F0: 42, F1: (v4 % ((v4) | 1))}
                _ = v10
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                if true {
                    var v12 int8 = 0
                    var v13 float64 = 3.14159
                    fmt.Println(v13)
                    _ = v12
                    _ = v13
                }
                var v14 int8 = 45
                var v15 float64 = (-2.25 / 1.0)
                _ = fi11
                _ = v14
                _ = v15
            }
            v4 = (v4 & (v4 & (v4 - v4)))
            _ = fi5
        }
    }
    v16 := GenS2{F0: 1, F1: (^(v4))}
    v4 -= (v4 ^ (v4 >> ((18446744073709551615) & 63)))
    var v17 int = 1
    v18 := v3[2]
    if (4294967295 <= 4294967295) {
        switch "" {
        case "":
            fmt.Println(v18, v4)
        default:
            switch v18 {
            case "foo bar":
                delete(v3, 2)
            case "\n\t":
                var v19 string = v18
                _ = v19
            case "":
                fmt.Println(v4)
            }
        }
    } else {
        v17 = (^(v17))
        v2 = append(v2, "foo bar")
        v3[42] = v18
    }
    var v20 int16 = 0
    var v21 uint = 4294967295
    var v22 uint = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
