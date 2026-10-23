package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []string{("a" + "foo bar")}
    var v2 int8 = 1
    var v3 float32 = ((1.5 - (((3.14159 / 1.5) * (0.0 + 3.14159)) + float32(0.0))) + (-2.25 - 0.1))
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "foo bar")
    }
    for fi4 := 0; fi4 < 15; fi4++ {
        v2 = (v2 ^ 0)
        v3 = -2.25
        _ = fi4
    }
    v5 := GenS1{F0: 1000, F1: "hello"}
    _ = len(v1)
    if true {
        if len(v1) > 0 {
            v1[0] = ("hello" + "a")
        }
    } else {
        v3 -= v3
        var v6 uint = 312
        _ = len(v1)
        _ = v6
    }
    fmt.Println(v2, v3)
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 = 126
        switch ("" + "\n\t") {
        case "hello":
            switch (fi7 | 9223372036854775807) {
            case 891:
                v2 = 1
            case -2147483648:
                var v8 float64 = 0.1
                _ = v8
            case 4294967295:
                v3 += (((v3 * (v3 * (v3 * float32(3.14159)))) * v3) / float32((0.0 - (-2.25 + (((-2.25 / 1e10) * -2.25) * float64(3.14159))))))
            default:
                for fi9 := 0; fi9 < 18; fi9++ {
                    v10 := GenS2{F0: 2}
                    _ = fi9
                    _ = v10
                }
            }
        }
        _ = fi7
    }
    if true {
        var v11 int16 = 0
        if (1 >= -9223372036854775808) {
            var v12 uint8 = 255
            for fi13 := 0; fi13 < 7; fi13++ {
                switch (-(fi13)) {
                case 351:
                    v12 = 127
                case 127:
                    var v14 float32 = ((v3 - v3) - float32((0.1 + 1.5)))
                    _ = v14
                case 42:
                    var v15 uint32 = 2147483647
                    _ = v15
                default:
                    if (!true) {
                        v11 -= v11
                    } else {
                        _ = len(v1)
                        if (false && true) {
                            v1 = append(v1, "a")
                        }
                    }
                }
                _ = fi13
            }
            _ = v12
        }
        v16 := GenS2{F0: 563}
        _ = v11
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 uint = 271
    var v18 uint = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
}
