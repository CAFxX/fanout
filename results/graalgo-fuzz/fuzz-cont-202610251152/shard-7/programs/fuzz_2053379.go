package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{(("foo bar" + "a") + "foo bar"), "a"}
    var v3 float32 = 1.5
    v4 := GenS1{F0: 18446744073709551615, F1: (("foo bar" + "foo bar") == "a"), F2: (((("a" + "a") + "") + "hello") + "\n\t")}
    switch 42 {
    case 446:
        v5 := GenS1{F0: 0, F1: (true || (false || false)), F2: "a"}
        _ = v5
    case -1:
        fmt.Println(v3)
    case -734:
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 0, F1: false, F2: "foo bar"}
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int64 = 9223372036854775807
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    v10 := GenS1{F0: 42, F1: ((-9223372036854775808 <= 0) || (0 < -76)), F2: "\n\t"}
    v3 = (v3 / v3)
    v11 := GenS2{F0: 2147483648, F1: 42}
    v1[("foo bar" + ("foo bar" + "a"))] = ((true && (2 >= 1)) || (42 != 127))
    v3 -= v3
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            switch fi12 {
            case -243:
                var v14 int8 = 42
                v3 = (1.5 + (v3 * ((v3 - v3) * (v3 * v3))))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v14
            case 9223372036854775807:
                v3 = ((v3 / v3) + (v3 - float32(((1e10 / (1e10 * -2.25)) * (1e10 + -2.25)))))
            }
            var v15 uint8 = 2
            var v16 string = ("hello" + "\n\t")
            _ = fi13
            _ = v15
            _ = v16
        }
        fmt.Println(fi12)
        var v17 int = fi12
        _ = fi12
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi19
        }
        var v20 uint16 = 796
        _ = fi18
        _ = v20
    }
    var v21 int8 = 116
    _ = len(v2)
    v3 += (v3 + (v3 * ((float32((3.14159 + 3.14159)) + v3) + v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
