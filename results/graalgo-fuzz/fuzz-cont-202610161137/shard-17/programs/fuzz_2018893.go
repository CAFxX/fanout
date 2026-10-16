package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int32) uint16 {
    p0 = (p0 % ((p0) | 1))
    var v4 uint8 = 42
    _ = p0
    _ = v4
    return 668
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    if ((((1.5 * (0.0 + 0.1)) - 1e10) + (1.5 - 0.1)) <= (((0.1 * 0.1) * 0.0) * ((1e10 / 0.1) + 3.14159))) {
        switch (("a" + "foo bar") + "\n\t") {
        case "":
            var v5 float32 = (((0.0 / 1.0) - float32(3.14159)) * float32(3.14159))
            _ = v5
        default:
            v6 := GenS1{F0: 127, F1: -9223372036854775808}
            _ = v6
        }
    } else {
        switch (("" + ("foo bar" + ("" + "foo bar"))) + ("" + "")) {
        case "a":
            v1 = (v1 << ((102) & 63))
        default:
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint64 = 489
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 uint8 = 42
                    v11 := GenS1{F0: 0, F1: fi7}
                    fmt.Println(v1)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 int16 = (v3 % ((510) | 1))
                _ = fi7
                _ = v12
                _ = v8
            }
        }
    }
    v3 += -5
    switch "a" {
    case "\n\t":
        v13 := v2["a"]
        _ = v13
    case "":
        fmt.Println(v3, v3)
    }
    v1 += v1
    if (false && ((v1 + (v1 - (^(v1)))) <= v1)) {
        if (0.0 <= 1e10) {
            switch "foo bar" {
            case "a":
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v14 := []int{9223372036854775807, 1, 29, 527}
                    _ = v14
                case 42:
                    var v15 int8 = 35
                    _ = v15
                }
            case "\n\t":
                if false {
                    v1 = v1
                    if true {
                        v3 = v3
                    }
                    v1 = v1
                } else {
                    if (4294967295 == 435) {
                        v1 = (^(v1))
                        v16 := v2[(("" + "foo bar") + "hello")]
                        for fi17 := 0; fi17 < 19; fi17++ {
                            _ = len(v2)
                            _ = fi17
                        }
                        _ = v16
                    }
                }
            case "foo bar":
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 uint8 = 255
                    _ = fi18
                    _ = v19
                }
            default:
                if false {
                    v3 = 1
                    var v20 int8 = 42
                    _ = v20
                } else {
                    v21 := GenS2{F0: 42, F1: 2}
                    _ = v21
                }
            }
            v1 = -676
        }
    } else {
        fmt.Println(v1, v1)
        v1 += (v1 * -2147483648)
    }
    if false {
        switch "a" {
        case "\n\t":
            var v22 bool = (false && (42 >= 1))
            _ = v22
        case "":
            switch "" {
            case "hello":
                var v23 uint16 = 255
                _ = v23
            }
        default:
            switch 873 {
            case 32767:
                v1 += (v1 * v1)
            default:
                v24 := GenS1{F0: 68, F1: 32767}
                _ = v24
            }
        }
    }
    v2["\n\t"] = false
    fmt.Println(v3, v1)
    v3 += (v3 & v3)
    v3 = (v3 / ((v3) | 1))
    var v25 int8 = -18
    v26 := GenS2{F0: -9223372036854775808, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
