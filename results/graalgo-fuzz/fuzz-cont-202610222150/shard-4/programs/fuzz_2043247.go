package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }

func main() {
    var v1 int8 = 63
    v2 := []string{}
    var v3 uint64 = 42
    v4 := []string{}
    v5 := make(map[string]bool)
    var v6 int16 = -1
    v6 = (v6 / ((-279) | 1))
    for fi7 := 0; fi7 < 11; fi7++ {
        v3 = 4294967295
        v5["a"] = (!(int64(v3) != -9223372036854775808))
        _ = fi7
    }
    v8 := GenS1{F0: (v6 / ((-32768) | 1))}
    switch 127 {
    case 32767:
        var v9 uint16 = 350
        _ = v9
    case 128:
        switch 1 {
        case 9223372036854775807:
            _ = len(v4)
        case 1:
            v1 += (v1 / (((-(v1))) | 1))
        case -2147483648:
            v6 = (v6 ^ -1)
        default:
            v1 = (v1 ^ 127)
        }
    default:
        if ((("a" + "hello") + "\n\t") >= "\n\t") {
            switch ("a" + ("hello" + (("" + ("foo bar" + (("a" + "hello") + "\n\t"))) + (("\n\t" + "foo bar") + "foo bar")))) {
            case "":
                v10 := GenS1{F0: 466}
                _ = v10
            case "hello":
                switch -243 {
                case -9223372036854775808:
                    _ = len(v5)
                case 2147483647:
                    fmt.Println(v1, v1, v6)
                case 32767:
                    v6 += v6
                default:
                    var v11 bool = true
                    _ = v11
                }
            default:
                var v12 uint = 42
                _ = v12
            }
        } else {
            var v13 uint64 = v3
            for fi14 := 0; fi14 < 2; fi14++ {
                v13 = (v13 & v3)
                var v15 float64 = -2.25
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    }
    v16 := GenS1{F0: -1}
    var v17 int8 = (v1 | 1)
    var v18 int32 = 0
    v18 = 2147483647
    var v19 int16 = 157
    switch ("foo bar" + ("a" + "")) {
    case "foo bar":
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 float64 = (float64(0.0) - 3.14159)
            for fi22 := 0; fi22 < 5; fi22++ {
                delete(v5, "\n\t")
                var v23 int = -1
                v5["\n\t"] = true
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    case "":
        delete(v5, "")
    case "a":
        switch 9223372036854775807 {
        case -2147483648:
            fmt.Println(v17)
        case 4294967295:
            v24 := GenS2{F0: (v17 << ((121) & 63)), F1: 688, F2: 42}
            _ = v24
        default:
            v5["\n\t"] = true
        }
    default:
        switch 32767 {
        case 32767:
            v18 = 743
        case -1:
            if (110 < 53) {
                v19 = -513
            }
        case 2147483647:
            fmt.Println(v1, v18)
        }
    }
    v6 = (-(v6))
    v6 = 534
    var v25 int32 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
