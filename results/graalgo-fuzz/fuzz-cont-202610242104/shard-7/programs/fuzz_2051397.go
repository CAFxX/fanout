package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
func genFn0() uint16 {
    var v7 int16 = 0
    var v8 int16 = v7
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float64 = 3.14159
    var v2 int = 202
    var v3 uint8 = 70
    v4 := []uint32{4294967295, 1, 4294967295}
    v5 := []string{"\n\t", ("a" + ("foo bar" + ("hello" + "foo bar")))}
    v6 := make(map[uint32]bool)
    fmt.Println(v1, v3, v2)
    fmt.Println(v3, v2)
    v1 -= -2.25
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 541, F1: v2}
        _ = fi9
        _ = v10
    }
    v5 = append(v5, "foo bar")
    var v11 int32 = 210
    v3 -= (v3 % (((^(v3))) | 1))
    v12 := GenS1{F0: 884, F1: 4294967295}
    if true {
        if (127 < 114) {
            var v13 uint32 = 42
            v1 = 0.0
            _ = v13
        }
        if (4294967295 < v2) {
            var v14 int = v2
            _ = v14
        } else {
            switch (-(v2)) {
            case 9223372036854775807:
                fmt.Println(v11, v2, v2)
            case 0:
                v3 -= 127
            }
            switch "a" {
            case "\n\t":
                _ = len(v6)
            }
            v2 = v2
        }
    }
    var v15 int32 = (v11 + (v11 * -2147483648))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
