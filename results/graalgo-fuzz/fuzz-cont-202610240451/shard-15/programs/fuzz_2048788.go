package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{4294967295, 609}
    var v2 float32 = 0.1
    var v3 uint32 = 4294967295
    var v4 uint16 = 209
    v5 := make(map[int]int)
    var v6 int64 = 396
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint = 255
            _ = len(v5)
            v4 += (v4 & v4)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if (127 > 202) {
            switch 128 {
            case 42:
                v11 := GenS1{F0: 709}
                _ = v11
            default:
                var v12 string = "a"
                _ = v12
            }
            v13 := GenS2{F0: 2}
            if false {
                v1 = append(v1, v3)
                v2 = (float32(0.0) / v2)
                v6 -= int64(fi10)
            }
            _ = v13
        }
        var v14 bool = true
        _ = fi10
        _ = v14
    }
    fmt.Println(v3)
    switch (("foo bar" + "hello") + "hello") {
    case "a":
        for fi15 := 0; fi15 < 12; fi15++ {
            if true {
                v4 = (v4 + 684)
            }
            v6 = (v6 / ((v6) | 1))
            _ = fi15
        }
    case "\n\t":
        v6 -= v6
    case "foo bar":
        v6 = 42
    default:
        if false {
            v4 = (v4 ^ v4)
        }
    }
    var v16 uint32 = v3
    v2 -= -2.25
    if (false || false) {
        v2 = 3.14159
    } else {
        v17 := GenS2{F0: 0}
        switch 128 {
        case -596:
            if true {
                for fi18 := 0; fi18 < 5; fi18++ {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v16 = (v3 * (v16 % ((v16) | 1)))
                        _ = fi19
                    }
                    var v20 uint = 841
                    v21 := GenS1{F0: v4}
                    _ = fi18
                    _ = v20
                    _ = v21
                }
            } else {
                var v22 bool = true
                var v23 uint32 = v3
                _ = v22
                _ = v23
            }
        default:
            v24 := GenS2{F0: 1}
            _ = v24
        }
        _ = v17
    }
    v16 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
