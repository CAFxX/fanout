package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint8 }

func main() {
    v1 := []string{"a", "hello", "foo bar", "foo bar"}
    var v2 uint8 = 98
    var v3 uint64 = 4294967295
    v4 := []string{"a", "a", "foo bar", "\n\t"}
    fmt.Println(v2, v3)
    fmt.Println(v2, v3)
    _ = len(v4)
    switch ("\n\t" + ("a" + "hello")) {
    case "":
        if false {
            var v5 float32 = 3.14159
            fmt.Println(v3, v3)
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS1{F0: -9223372036854775808, F1: 2, F2: 1}
                v5 = 1e10
                _ = fi6
                _ = v7
            }
            _ = v5
        }
    default:
        var v8 uint = 268
        _ = v8
    }
    fmt.Println(v3, v2)
    if true {
        v3 -= v3
        v1 = append(v1, "foo bar")
    } else {
        var v9 float64 = (0.1 * 0.0)
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v1 = append(v1, "hello")
        _ = fi10
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    var v11 int = -1
    var v12 uint8 = (v2 >> ((69) & 63))
    v13 := GenS1{F0: -539, F1: v3, F2: 32767}
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: true, F1: v3, F2: v2}
        for fi16 := 0; fi16 < 20; fi16++ {
            v2 = (v12 / ((v2) | 1))
            var v17 float64 = 1.5
            v2 += (v2 + (v12 ^ 73))
            _ = fi16
            _ = v17
        }
        if true {
            _ = len(v4)
            var v18 string = (("hello" + "\n\t") + ("a" + ("foo bar" + "\n\t")))
            if len(v1) > 0 {
                v1[0] = (v18 + v18)
            }
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    v19 := GenS1{F0: 455, F1: (v3 ^ v3), F2: 0}
    switch v11 {
    case 1:
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                fmt.Println(fi21, v2, fi20)
                _ = fi21
            }
            var v22 string = ("\n\t" + "\n\t")
            var v23 int16 = -423
            _ = fi20
            _ = v22
            _ = v23
        }
    case -1:
        var v24 float64 = 0.1
        _ = v24
    case 32767:
        if (true && ((-24 >= 100) && (((!true) || ("hello" == "a")) && (v3 < (v3 & 0))))) {
            var v25 int64 = 9223372036854775807
            _ = v25
        } else {
            v4 = append(v4, "a")
            v2 -= uint8(v3)
        }
    }
    v12 += v12
    fmt.Println(v11, v12, v12)
    v12 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
