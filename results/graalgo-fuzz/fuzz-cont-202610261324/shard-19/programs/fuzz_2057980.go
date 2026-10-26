package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }

func main() {
    var v1 int = 128
    v2 := []string{"a", ""}
    v3 := make(map[int]string)
    _ = len(v3)
    switch 1 {
    case -2147483648:
        switch 1 {
        case 1:
            v4 := GenS1{F0: true, F1: 42, F2: (v1 ^ 714)}
            _ = v4
        case 128:
            v1 -= 470
        }
    default:
        v1 -= 128
    }
    var v5 bool = (0.0 <= (1e10 / (0.1 * float32(3.14159))))
    switch v1 {
    case -827:
        v5 = (42 > -153)
    case -388:
        for fi6 := 0; fi6 < 20; fi6++ {
            switch "a" {
            case "":
                switch ("a" + "a") {
                case "\n\t":
                    v5 = v5
                }
            case "foo bar":
                var v7 float64 = ((1.5 + 0.1) - ((-2.25 * -2.25) + -2.25))
                _ = v7
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if v5 {
                    if ((0.1 - float64(3.14159)) <= (1e10 - -2.25)) {
                        v5 = (true && (65535 != 0))
                    }
                }
            }
            switch ((("hello" + "foo bar") + ("a" + "foo bar")) + "a") {
            case "hello":
                v3[-2147483648] = (("a" + "") + "\n\t")
            default:
                var v8 bool = v5
                _ = v8
            }
            _ = fi6
        }
    }
    var v9 int8 = 10
    var v10 uint8 = 42
    var v11 uint32 = 0
    if (!(762 > uint16(v11))) {
        if v5 {
            var v12 int32 = 2147483647
            var v13 uint16 = 255
            _ = v12
            _ = v13
        } else {
            var v14 int = int(v9)
            _ = v14
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                if (true || v5) {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                } else {
                    _ = len(v3)
                }
                switch 9223372036854775807 {
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -59:
                    if ((255 != uint16(v11)) || (!(!true))) {
                        v2 = append(v2, "a")
                        if false {
                            var v17 int32 = 64
                            _ = len(v2)
                            v3[(fi16 | fi16)] = ("foo bar" + (("a" + "hello") + ("hello" + "\n\t")))
                            _ = v17
                        }
                        v5 = false
                        v5 = v5
                        var v18 uint8 = v10
                        _ = v18
                    }
                }
                _ = fi16
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                if v5 {
                    var v20 int8 = (v9 - int8(v10))
                    _ = v20
                }
                _ = fi19
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                delete(v3, v1)
                v1 -= (fi21 * fi15)
                _ = fi21
            }
            _ = fi15
        }
    }
    v22 := GenS1{F0: (42 <= 174), F1: 42, F2: (v1 >> ((2) & 63))}
    var v23 int64 = 0
    fmt.Println(v9)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v24 int = (v1 & (v1 / ((-9223372036854775808) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v5
    _ = v9
}
