package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{32768, 2147483647}
    var v2 int16 = -286
    var v3 int = -879
    var v4 int64 = 0
    var v5 uint = 550
    switch "\n\t" {
    case "hello":
        v1 = append(v1, (v3 + v3))
    case "\n\t":
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: (-(v2))}
            _ = fi6
            _ = v7
        }
    }
    var v8 int32 = -884
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int16 = v2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v11 := GenS1{F0: -310}
        _ = fi9
        _ = v10
        _ = v11
    }
    if (false || false) {
        if (!true) {
            var v12 uint32 = 760
            if ((v5 / (((v5 * (v5 - v5))) | 1)) == 18446744073709551615) {
                switch "foo bar" {
                case "":
                    var v13 int = (v3 * (v3 ^ (v3 - (v3 * 0))))
                    _ = v13
                case "foo bar":
                    var v14 int64 = 1
                    _ = v14
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                if ((4294967295 > 2) || true) {
                    fmt.Println(v2, v8)
                }
                var v15 uint32 = 4294967295
                _ = v15
            }
            _ = v12
        }
        switch v3 {
        case 568:
            v8 += (v8 * 0)
        default:
            _ = len(v1)
        }
    } else {
        v3 = v3
    }
    var v16 uint = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        if (!false) {
            var v18 float64 = 0.1
            _ = v18
        } else {
            fmt.Println(fi17, v4, v8)
            v3 += (fi17 - v3)
            for fi19 := 0; fi19 < 18; fi19++ {
                v3 -= (v3 & -28)
                if (2 > 255) {
                    v16 = v16
                    switch (((("" + "hello") + "\n\t") + "a") + "foo bar") {
                    case "":
                        v2 = (v2 / (((-(v2))) | 1))
                    case "foo bar":
                        v4 += (v4 >> ((v16) & 63))
                    }
                    _ = len(v1)
                    var v20 uint64 = 1
                    _ = v20
                } else {
                    v8 += v8
                }
                v4 -= (-(v4))
                var v21 int16 = 0
                _ = fi19
                _ = v21
            }
        }
        v1 = append(v1, v3)
        _ = fi17
    }
    v5 += v16
    var v22 uint8 = 2
    v23 := GenS2{F0: (-(v8))}
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println(v16, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
