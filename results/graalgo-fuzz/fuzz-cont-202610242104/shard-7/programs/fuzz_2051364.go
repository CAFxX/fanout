package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }

func main() {
    var v1 int64 = 42
    var v2 float32 = 1.5
    var v3 float64 = (((-2.25 - 0.0) * 0.0) / 1.0)
    var v4 float64 = v3
    fmt.Println(v2)
    if false {
        v5 := []int{}
        switch 2147483647 {
        case 9223372036854775807:
            if ((true && (!(!((v1 << ((153) & 63)) < 176)))) && (1 > 1)) {
                v5 = append(v5, 97)
                if len(v5) > 0 {
                    v5[0] = -1
                }
            }
        case 364:
            if false {
                if (!true) {
                    v3 = 0.1
                    var v6 string = "a"
                    v3 = v3
                    _ = v6
                }
                for fi7 := 0; fi7 < 7; fi7++ {
                    fmt.Println(v1, v4)
                    if len(v5) > 0 {
                        v5[0] = fi7
                    }
                    _ = fi7
                }
            }
        case -9223372036854775808:
            if false {
                var v8 string = ("hello" + ("hello" + "\n\t"))
                v4 = (v3 * (v3 * float64((v2 + (v2 / float32((0.0 - ((v4 * float64(v2)) + (0.1 / v4)))))))))
                _ = v8
            } else {
                _ = len(v5)
                if len(v5) > 0 {
                    v5[0] = 50
                }
            }
        default:
            var v9 uint16 = 0
            _ = v9
        }
        var v10 uint32 = 1
        _ = v10
        _ = v5
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v1 -= (^(v1))
        _ = fi11
    }
    v12 := GenS1{F0: 147}
    fmt.Println(v2, v3)
    v13 := []int{-9223372036854775808, 128}
    switch "foo bar" {
    case "hello":
        var v14 int = 32768
        _ = v14
    case "foo bar":
        var v15 float32 = (v2 * v2)
        _ = v15
    }
    fmt.Println(v4, v4, v1)
    for fi16 := 0; fi16 < 7; fi16++ {
        v1 -= (-(v1))
        fmt.Println(v1, v3)
        var v17 float32 = v2
        _ = fi16
        _ = v17
    }
    if (65535 > 2147483648) {
        for fi18 := 0; fi18 < 12; fi18++ {
            v13 = append(v13, fi18)
            for fi19 := 0; fi19 < 13; fi19++ {
                v2 += v2
                v4 += (((((float64(1e10) / (v4 / v4)) / -2.25) + 3.14159) * float64(v2)) + float64(v2))
                _ = fi19
            }
            _ = fi18
        }
        v20 := GenS2{F0: true, F1: (-1 != 32768), F2: 42}
        _ = v20
    } else {
        fmt.Println(v1)
    }
    if (true && false) {
        var v21 uint8 = 42
        v3 -= 1e10
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v21
    }
    if false {
        var v22 float64 = (float64(1.5) - (v4 / 1.5))
        var v23 float64 = v22
        for fi24 := 0; fi24 < 19; fi24++ {
            if true {
                var v25 int = fi24
                v26 := GenS1{F0: 1}
                _ = v25
                _ = v26
            } else {
                var v27 uint64 = 2
                v2 += v2
                _ = v27
            }
            if false {
                v23 = v23
                var v28 int = (fi24 | fi24)
                switch (fi24 + 1) {
                case 2147483647:
                    var v29 uint16 = 65535
                    _ = v29
                }
                _ = v28
            }
            _ = fi24
        }
        _ = v22
        _ = v23
    } else {
        if (true || (-106 != -32768)) {
            fmt.Println(v1, v3, v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
