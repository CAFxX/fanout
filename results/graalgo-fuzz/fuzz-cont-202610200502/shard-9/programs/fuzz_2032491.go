package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]bool)
    var v3 int16 = v1
    v4 := []string{("" + "\n\t")}
    fmt.Println(v3)
    _ = len(v4)
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            v1 = v3
            var v7 uint8 = 242
            var v8 int32 = 628
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    v1 += 0
    var v9 string = "hello"
    v3 = -32768
    if (v3 > v1) {
        switch "hello" {
        case "":
            if (1 <= 6) {
                if (((0.1 == float64(0.0)) && false) || (0.0 <= (1e10 * (1.5 * 1e10)))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    v4 = append(v4, v9)
                }
                v3 = 0
            }
        case "a":
            v3 -= (v3 - -1)
        case "\n\t":
            if false {
                v9 = v9
            } else {
                var v10 int16 = (v1 >> ((807) & 63))
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                fmt.Println(v10, v9, v3)
                _ = v10
            }
        }
    } else {
        switch v9 {
        case "hello":
            v9 = ("hello" + v9)
        }
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v9 = "\n\t"
            _ = fi12
        }
        v3 -= -32768
        var v13 string = (((v9 + ("foo bar" + (v9 + "hello"))) + v9) + ((((v9 + v9) + (v9 + v9)) + v9) + (("" + "hello") + v9)))
        _ = fi11
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if false {
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: (true && (-9223372036854775808 < 9223372036854775807))}
                v9 = v9
                switch v9 {
                case "hello":
                    v3 = v1
                default:
                    v2[fi15] = (!false)
                }
                _ = fi15
                _ = v16
            }
            delete(v2, (-(fi14)))
        } else {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        switch v9 {
        case "":
            for fi17 := 0; fi17 < 7; fi17++ {
                v2[(fi17 % ((fi14) | 1))] = (-1 >= 0)
                _ = fi17
            }
        default:
            v18 := GenS1{F0: false}
            _ = v18
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 bool = (((0.1 >= -2.25) || false) || false)
        _ = fi19
        _ = v20
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
