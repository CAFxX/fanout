package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 369
    v3 := make(map[uint32]bool)
    if (1 != 247) {
        var v4 float64 = 1.5
        _ = v4
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            if false {
                v7 := []int{}
                _ = len(v3)
                _ = v7
            }
            if (((("\n\t" + ("a" + "foo bar")) + ((((("\n\t" + "a") + "") + ("foo bar" + ("hello" + ("\n\t" + "")))) + ("foo bar" + ("a" + "\n\t"))) + "")) + (("a" + "\n\t") + ("\n\t" + ((((("\n\t" + (("\n\t" + "foo bar") + "\n\t")) + "") + "foo bar") + (("a" + "a") + ("foo bar" + "foo bar"))) + "a")))) != ("hello" + "")) {
                var v8 uint64 = 997
                _ = v8
            }
            fmt.Println(v2)
            var v9 int32 = -56
            _ = fi6
            _ = v9
        }
        var v10 uint = 0
        if ((v2 + 0) < (v2 << ((v10) & 63))) {
            v2 -= (^(v2))
            var v11 uint8 = 1
            _ = v11
        } else {
            v12 := []int{(fi5 << ((v10) & 63)), (-(fi5)), (fi5 / ((42) | 1))}
            v13 := GenS1{F0: 58}
            _ = v12
            _ = v13
        }
        _ = fi5
        _ = v10
    }
    v2 -= v2
    if (!false) {
        v14 := GenS1{F0: -128}
        _ = v14
    }
    var v15 uint16 = 92
    v15 = v15
    fmt.Println(v2, v15)
    for fi16 := 0; fi16 < 10; fi16++ {
        _ = len(v3)
        v17 := GenS1{F0: 1}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 0, F1: -9223372036854775808}
    switch "" {
    case "a":
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 += v2
            _ = fi19
        }
    case "foo bar":
        v15 -= v15
    }
    switch "" {
    case "\n\t":
        v15 = (^(v15))
    case "hello":
        v20 := v3[(v2 + v2)]
        _ = v20
    default:
        if (0 == 42) {
            v21 := GenS3{F0: (("\n\t" + "hello") + "a")}
            _ = v21
        } else {
            v15 += 65535
            var v22 uint16 = v15
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
