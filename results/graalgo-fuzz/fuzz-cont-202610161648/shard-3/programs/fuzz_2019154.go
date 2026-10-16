package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0(p0 int8, p1 uint, p2 int16) string {
    p1 = p1
    p1 = (p1 % (((p1 | (p1 * (p1 ^ (p1 << (((p1 + p1)) & 63)))))) | 1))
    switch 220 {
    case 468:
        p1 = p1
    case 609:
        v5 := GenS1{F0: 1}
        _ = v5
    case 0:
        for fi6 := 0; fi6 < 1; fi6++ {
            switch "" {
            case "hello":
                for fi7 := 0; fi7 < 10; fi7++ {
                    switch fi6 {
                    case 1:
                        fmt.Println(p1, p2, p2)
                    case 921:
                        p1 -= (^(p1))
                    default:
                        var v8 uint = p1
                        _ = v8
                    }
                    p2 -= p2
                    p2 += (p2 ^ (p2 + -32768))
                    _ = fi7
                }
            case "":
                var v9 uint16 = 5
                _ = v9
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (("\n\t" + "foo bar") + "hello")
}

func main() {
    var v1 uint8 = 17
    var v2 float32 = 1e10
    var v3 uint16 = 12
    var v4 int64 = 773
    fmt.Println(v2)
    switch "foo bar" {
    case "hello":
        var v10 uint8 = 1
        _ = v10
    default:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := []int{(-(fi11))}
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v1)
    if false {
        if true {
            if false {
                v13 := GenS2{F0: 484}
                fmt.Println(v2)
                _ = v13
            } else {
                var v14 uint32 = 2
                switch "hello" {
                case "foo bar":
                    var v15 int = 2147483647
                    _ = v15
                case "a":
                    if (((1.5 * -2.25) != (0.1 - 1.5)) && (true || (((v1 >> ((1) & 63)) > uint8(v3)) && (false && true)))) {
                        v1 = uint8(v14)
                    }
                case "hello":
                    fmt.Println(v4)
                }
                _ = v14
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{867, fi16, 128}
                _ = fi16
                _ = v17
            }
        } else {
            fmt.Println(v4, v1)
            if (32767 <= v3) {
                fmt.Println(v3, v3)
            }
        }
    }
    v18 := []int{950, -597}
    var v19 int32 = 42
    var v20 string = ("" + ("foo bar" + ""))
    if false {
        var v21 uint32 = 954
        v22 := GenS1{F0: (v21 / (((^(v21))) | 1))}
        _ = v21
        _ = v22
    }
    v3 = 42
    if (!true) {
        v19 -= (v19 << ((80) & 63))
        if false {
            if true {
                v23 := GenS1{F0: 258}
                var v24 string = (v20 + v20)
                _ = v23
                _ = v24
            } else {
                switch (v20 + (((("\n\t" + "hello") + v20) + v20) + v20)) {
                case "foo bar":
                    var v25 int64 = (v4 & v4)
                    _ = v25
                default:
                    if (true || false) {
                        v26 := GenS3{F0: (v19 + v19), F1: -1}
                        v3 = v3
                        _ = v26
                    } else {
                        v4 = v4
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            _ = fi27
                        }
                    }
                }
                var v28 uint8 = v1
                _ = v28
            }
        }
        switch v20 {
        case "hello":
            var v29 float32 = ((float32(-2.25) - v2) * (v2 / v2))
            _ = v29
        case "foo bar":
            if len(v18) > 0 {
                v18[0] = -9223372036854775808
            }
        case "\n\t":
            var v30 int16 = -1
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
