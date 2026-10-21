package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint) int8 {
    if true {
        v5 := []int{}
        if ((65535 > 255) && true) {
            fmt.Println(p0)
        } else {
            v6 := GenS1{F0: false, F1: ("hello" != ""), F2: "hello"}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            p0 = p0
            _ = v6
        }
        p0 -= p0
        _ = v5
    }
    v7 := GenS1{F0: false, F1: (((((-2.25 + 1e10) / -2.25) / 0.1) >= 0.1) && false), F2: "\n\t"}
    v8 := GenS1{F0: (false || true), F1: true, F2: ("a" + (("a" + ("" + (("\n\t" + "hello") + ""))) + "\n\t"))}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint16, p1 uint8) bool {
    p1 -= 2
    fmt.Println(p1)
    p0 += p0
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := []string{"", "foo bar", (("\n\t" + "") + ("\n\t" + "hello"))}
    v2 := make(map[string]int)
    var v3 uint16 = 345
    v4 := make(map[string]string)
    v3 += v3
    fmt.Println(v3)
    v3 += v3
    fmt.Println(v3)
    fmt.Println(v3)
    var v9 bool = (32767 >= 9223372036854775807)
    var v10 int = -9223372036854775808
    fmt.Println(v9, v10)
    var v11 float64 = (0.0 / 0.1)
    v3 -= v3
    v12 := GenS1{F0: (false || false), F1: ((-649 == int64(v10)) && (!(200 <= 4294967295))), F2: "foo bar"}
    v13 := GenS2{F0: 0}
    fmt.Println(v3, v9)
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS3{F0: 42, F1: 466}
        v3 = (^(v3))
        var v16 int32 = int32(v3)
        _ = fi14
        _ = v15
        _ = v16
    }
    switch (^(v10)) {
    case -9223372036854775808:
        _ = len(v2)
    case 43:
        var v17 int = v10
        _ = v17
    case 1:
        v3 -= (v3 | v3)
    default:
        v9 = (!(!v9))
    }
    var v18 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
