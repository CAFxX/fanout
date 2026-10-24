package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint; F2 string }

func main() {
    var v1 int = -2147483648
    var v2 uint8 = 127
    var v3 int = v1
    var v4 int = v1
    v5 := []uint32{4294967295, 82, 529, 547}
    v6 := make(map[int]bool)
    v4 -= v3
    v3 -= v3
    v7 := GenS1{F0: -1, F1: ("" + ((((("a" + "foo bar") + "") + "foo bar") + "") + "")), F2: 42}
    if true {
        fmt.Println(v3, v4, v4)
        fmt.Println(v3, v3)
    }
    v4 = 1
    var v8 uint16 = 582
    fmt.Println(v8, v3)
    for fi9 := 0; fi9 < 4; fi9++ {
        delete(v6, v4)
        v1 += (-(v1))
        _ = fi9
    }
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS1{F0: 42, F1: "foo bar", F2: 137}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        v13 := GenS2{F0: 782, F1: 0, F2: ("" + "")}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(fi14, fi14)
            _ = fi14
        }
    }
    if (!(!false)) {
        v3 = v3
        fmt.Println(v8, v2)
        var v15 float32 = -2.25
        _ = v15
    }
    v6[(v4 + v1)] = (32767 > 1)
    switch "" {
    case "":
        var v16 uint32 = 2
        _ = v16
    case "hello":
        v2 += v2
    case "a":
        v17 := GenS1{F0: 9223372036854775807, F1: "a", F2: -32768}
        _ = v17
    }
    v18 := v6[32768]
    _ = len(v5)
    if false {
        v3 -= -2147483648
        switch "hello" {
        case "hello":
            var v19 string = "foo bar"
            _ = v19
        case "\n\t":
            if (4294967295 <= 4294967295) {
                switch "foo bar" {
                case "foo bar":
                    delete(v6, -1)
                case "a":
                    _ = len(v5)
                case "\n\t":
                    switch (v1 ^ (^(v4))) {
                    case 32768:
                        delete(v6, v1)
                    default:
                        v20 := GenS2{F0: 0, F1: 4294967295, F2: ""}
                        _ = v20
                    }
                default:
                    v4 += 128
                }
                var v21 int64 = -1
                _ = v21
            }
        case "":
            fmt.Println(v8)
        default:
            switch (v3 & 32767) {
            case 1:
                if (false || v18) {
                    var v22 uint32 = 42
                    if ((v8 % (((v8 % ((v8) | 1))) | 1)) < v8) {
                        v4 -= (v3 >> ((42) & 63))
                        var v23 int16 = 1000
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v25 := GenS1{F0: 42, F1: (("foo bar" + "hello") + ""), F2: 32767}
                            _ = fi24
                            _ = v25
                        }
                        v5 = append(v5, (v22 % ((65535) | 1)))
                    }
                    v8 -= 681
                    _ = v22
                } else {
                    switch "\n\t" {
                    case "":
                        v4 -= v3
                    case "a":
                        v6[v3] = (!false)
                    case "foo bar":
                        delete(v6, v1)
                    }
                    for fi26 := 0; fi26 < 10; fi26++ {
                        var v27 uint32 = 4294967295
                        _ = fi26
                        _ = v27
                    }
                    v8 = v8
                }
            case -9223372036854775808:
                fmt.Println(v1)
            default:
                v18 = true
            }
        }
        var v28 int64 = -9223372036854775808
        _ = v28
    } else {
        var v29 string = ("" + "")
        if v18 {
            v30 := GenS2{F0: 42, F1: 2, F2: ("a" + (v29 + (v29 + v29)))}
            v31 := GenS1{F0: -9223372036854775808, F1: "foo bar", F2: 1000}
            var v32 int64 = 9223372036854775807
            _ = v30
            _ = v31
            _ = v32
        } else {
            var v33 float64 = -2.25
            _ = v33
        }
        _ = v29
    }
    var v34 uint32 = 372
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
