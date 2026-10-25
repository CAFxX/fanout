package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    var v1 float32 = 0.1
    v2 := []string{"foo bar"}
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: "hello", F1: 18446744073709551615}
    _ = len(v2)
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 -= v6
        switch "hello" {
        case "\n\t":
            var v9 int16 = 0
            _ = v9
        }
        var v10 uint16 = 2
        _ = fi8
        _ = v10
    }
    v3["hello"] = -2147483648
    var v11 bool = (!true)
    v12 := GenS1{F0: ("\n\t" + (((("hello" + "") + ("a" + ("a" + "a"))) + ("a" + "\n\t")) + "a")), F1: 139}
    v13 := GenS1{F0: "hello", F1: 623}
    for fi14 := 0; fi14 < 6; fi14++ {
        if (-191 >= -182) {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            v11 = true
        } else {
            _ = len(v2)
            v15 := v3[(("foo bar" + "foo bar") + ("\n\t" + "\n\t"))]
            if (!(118 >= 42)) {
                v6 -= v6
            }
            _ = v15
        }
        _ = fi14
    }
    fmt.Println(v11, v11, v11)
    var v16 uint8 = 42
    v17 := GenS1{F0: "\n\t", F1: 292}
    v1 += 0.0
    switch 4294967295 {
    case -9223372036854775808:
        var v18 uint8 = v16
        _ = v18
    case 0:
        v6 += v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
