package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 42
    v3 := []int{127}
    v4 := []string{"", "foo bar"}
    if (42 <= -1) {
        v2 = -29
    }
    fmt.Println(v2)
    v5 := GenS1{F0: 0, F1: (false || (42 >= 9223372036854775807)), F2: 42}
    fmt.Println(v2)
    v2 = (v2 / ((100) | 1))
    delete(v1, "hello")
    if (1 == 42) {
        if (42 == 18446744073709551615) {
            if true {
                v6 := GenS2{F0: -9223372036854775808, F1: false, F2: true}
                v7 := GenS2{F0: 9223372036854775807, F1: true, F2: true}
                _ = v6
                _ = v7
            } else {
                v8 := GenS2{F0: 4294967295, F1: ("hello" > "\n\t"), F2: true}
                v9 := GenS3{F0: (("foo bar" <= ("\n\t" + ((((("foo bar" + "") + ("hello" + "a")) + ("" + ("\n\t" + ("hello" + (("" + "hello") + "a"))))) + "a") + "\n\t"))) && false), F1: -9223372036854775808, F2: 0}
                _ = v8
                _ = v9
            }
        } else {
            v2 = (-(v2))
            if len(v4) > 0 {
                v4[0] = ("a" + "")
            }
        }
        v2 += 100
        if (1 >= 42) {
            for fi10 := 0; fi10 < 14; fi10++ {
                if (((!false) || (644 != 557)) && (-577 > 42)) {
                    var v11 uint8 = 0
                    v11 += v11
                    _ = v11
                }
                switch (fi10 >> ((714) & 63)) {
                case 127:
                    switch -9223372036854775808 {
                    case 127:
                        v2 += -33
                    case 52:
                        var v12 uint64 = 375
                        _ = v12
                    case -9223372036854775808:
                        var v13 int8 = (^(v2))
                        _ = v13
                    default:
                        switch (fi10 + fi10) {
                        case -9223372036854775808:
                            var v14 int16 = -32768
                            _ = v14
                        case 4294967295:
                            v2 -= 99
                        default:
                            v1[("" + "foo bar")] = 2
                        }
                    }
                default:
                    fmt.Println(fi10, v2)
                }
                var v15 int32 = 863
                _ = fi10
                _ = v15
            }
        }
    }
    switch 1 {
    case 1:
        var v16 uint8 = 71
        _ = v16
    case -151:
        v2 -= (v2 >> ((42) & 63))
    default:
        v2 = 42
    }
    v2 = (v2 % (((v2 * (-(v2)))) | 1))
    var v17 uint8 = 128
    var v18 int64 = -301
    var v19 float32 = (0.0 / 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
