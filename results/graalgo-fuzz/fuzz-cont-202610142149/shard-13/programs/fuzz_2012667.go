package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }

func main() {
    v1 := []uint32{140}
    var v2 int64 = -406
    var v3 int32 = 0
    var v4 float64 = (float64(1.5) - 1.5)
    var v5 uint16 = 255
    v6 := make(map[string]uint32)
    if ("\n\t" < ("foo bar" + "a")) {
        fmt.Println(v5)
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            if (!(!(true && true))) {
                _ = len(v1)
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                var v8 int = 32767
                _ = v8
            } else {
                var v9 int64 = 0
                _ = v9
            }
            var v10 uint16 = v5
            _ = fi7
            _ = v10
        }
    }
    v5 = (^(v5))
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := v6[("hello" + ("a" + ("hello" + (("hello" + ("foo bar" + "hello")) + "\n\t"))))]
        _ = fi11
        _ = v12
    }
    v5 -= v5
    v13 := GenS1{F0: 0, F1: "foo bar"}
    v4 = (1e10 + 1.5)
    var v14 int8 = 42
    switch (("\n\t" + "foo bar") + "hello") {
    case "a":
        v15 := GenS1{F0: 153, F1: ("foo bar" + "\n\t")}
        _ = v15
    case "foo bar":
        switch "" {
        case "":
            v4 -= (v4 * (3.14159 + v4))
        case "\n\t":
            switch 9223372036854775807 {
            case 32767:
                var v16 bool = (false || true)
                _ = v16
            case -1:
                fmt.Println(v4, v14)
            case -502:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if ("" < "a") {
                    v14 -= (v14 >> ((2) & 63))
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if (false || (65535 == 540)) {
                            v3 = (v3 % ((307) | 1))
                        }
                        v2 -= 1
                        var v18 int16 = 1000
                        _ = fi17
                        _ = v18
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v2 = v2
                }
            }
        default:
            var v19 uint8 = 2
            _ = v19
        }
    default:
        v6["\n\t"] = 65535
    }
    if (false || true) {
        v4 += float64(0.0)
        fmt.Println(v3, v5, v5)
        fmt.Println(v14, v4)
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 740
        }
        switch ("a" + "\n\t") {
        case "a":
            delete(v6, "\n\t")
        default:
            var v20 uint32 = 2
            _ = v20
        }
    }
    v3 = (v3 - -1)
    v21 := GenS1{F0: 127, F1: "a"}
    var v22 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
