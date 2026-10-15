package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    var v4 int16 = -1
    v5 := []string{}
    delete(v3, "foo bar")
    _ = len(v3)
    _ = len(v1)
    var v6 int8 = 127
    fmt.Println(v4, v4, v4)
    if (100 != (v6 + v6)) {
        v2 = (v2 * ((-2.25 * (((0.1 + 3.14159) + 1e10) + v2)) / v2))
        if false {
            v4 -= (v4 | v4)
            for fi7 := 0; fi7 < 2; fi7++ {
                switch (("" + ("" + "")) + "foo bar") {
                case "":
                    v2 = 0.0
                case "hello":
                    v8 := GenS1{F0: 0, F1: 18446744073709551615}
                    _ = v8
                }
                _ = fi7
            }
            if (!false) {
                var v9 int8 = v6
                fmt.Println(v2, v4)
                _ = v9
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v10 float32 = ((1.5 + 3.14159) - 1e10)
    fmt.Println(v6, v4)
    v4 = (v4 % (((^(v4))) | 1))
    if (false && false) {
        v2 = (1.5 - (1e10 - float64(v10)))
        fmt.Println(v2, v2, v4)
        fmt.Println(v10)
    }
    v11 := GenS2{F0: (("" + "a") + (((("a" + ("a" + ("\n\t" + ("a" + "foo bar")))) + "") + ("" + "hello")) + ("hello" + ("a" + ("" + ("hello" + ("foo bar" + "a"))))))), F1: (!true)}
    if ((-2147483648 < 1) && (!(v6 != 0))) {
        if false {
            v10 = 0.1
        } else {
            if true {
                var v12 float32 = ((v10 - 1.5) * (float32(v2) + v10))
                _ = v12
            }
            var v13 string = ""
            if len(v5) > 0 {
                v5[0] = ("a" + v13)
            }
            _ = v13
        }
        v6 -= (v6 >> ((196) & 63))
        v10 += 1.5
    } else {
        v10 -= ((v10 - 0.1) + v10)
        v14 := GenS3{F0: ("" == "a"), F1: 9223372036854775807}
        _ = v14
    }
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = ("" + "a")
    }
    fmt.Println(v10, v4, v10)
    switch "hello" {
    case "":
        fmt.Println(v6, v2, v6)
    }
    if ((v4 % (((v4 + (v4 * 249))) | 1)) >= 1) {
        fmt.Println(v4, v2, v10)
        if (0.1 == float32(float64(float32(0.0)))) {
            var v15 uint64 = 1
            _ = v15
        }
        switch "" {
        case "hello":
            v3["\n\t"] = 696
        case "a":
            if (!false) {
                v16 := v1[42]
                _ = v16
            }
        default:
            var v17 int64 = -1
            _ = v17
        }
    } else {
        if ((v6 ^ v6) > 42) {
            for fi18 := 0; fi18 < 8; fi18++ {
                if (!("hello" >= ("\n\t" + "a"))) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if true {
                            v6 -= (-(v6))
                        } else {
                            v2 += v2
                        }
                        v3[(((("" + "\n\t") + ("a" + "")) + "hello") + "")] = 0
                        v4 -= -197
                        _ = fi19
                    }
                }
                v10 -= (1.5 / float32((v2 + 0.0)))
                v4 = 42
                _ = fi18
            }
            var v20 float64 = v2
            _ = len(v3)
            _ = v20
        }
    }
    v10 -= (-2.25 / float32(-2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
