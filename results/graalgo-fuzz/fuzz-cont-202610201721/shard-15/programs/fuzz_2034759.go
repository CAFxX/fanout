package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint }
func genFn0() int16 {
    var v6 uint64 = 1
    var v7 uint = 4294967295
    _ = v6
    _ = v7
    return 551
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 0.1
    var v3 bool = ((2147483648 < 831) || (("a" + "hello") == ("foo bar" + "")))
    v4 := make(map[uint32]int)
    v5 := []uint32{2, 2, 0, 1}
    if v3 {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v3, v3)
            v9 := GenS1{F0: -1, F1: (("a" + "foo bar") + "a"), F2: 42}
            if v3 {
                if v3 {
                    v5 = append(v5, 2147483648)
                }
                _ = len(v5)
                var v10 string = ""
                _ = v10
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v2 = -2.25
                    v1 = (v1 ^ (v1 % (((v1 << ((42) & 63))) | 1)))
                    _ = fi11
                }
                for fi12 := 0; fi12 < 20; fi12++ {
                    fmt.Println(fi12, v1, v2)
                    v4[4294967295] = 283
                    v1 += v1
                    _ = fi12
                }
            }
            _ = fi8
            _ = v9
        }
        if (1e10 >= 1e10) {
            var v13 uint32 = 110
            _ = v13
        }
        v2 += v2
    }
    v2 += (v2 * float64((float32(0.1) / 0.1)))
    v3 = v3
    v3 = (v3 && v3)
    var v14 int32 = -1
    var v15 float32 = (((1.5 + (float32(v2) - ((3.14159 * 1e10) * 3.14159))) - 3.14159) - ((float32((v2 * 1e10)) - 1.5) * 0.1))
    switch 32768 {
    case -979:
        var v16 uint = 255
        _ = v16
    case -9223372036854775808:
        v17 := GenS2{F0: v14}
        _ = v17
    case -390:
        v18 := v4[0]
        _ = v18
    }
    v14 = (v14 >> ((173) & 63))
    v2 += v2
    if (((9223372036854775807 == 1) && v3) && false) {
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(v3, v3, v14)
            var v20 uint64 = 18446744073709551615
            _ = fi19
            _ = v20
        }
        if v3 {
            switch -2147483648 {
            case 9223372036854775807:
                v15 = 0.0
            }
        }
    } else {
        v14 = -658
        v21 := GenS2{F0: v14}
        _ = v21
    }
    v3 = (!v3)
    switch ("" + (("\n\t" + (("a" + "a") + "a")) + ("hello" + "foo bar"))) {
    case "foo bar":
        v1 -= (v1 / ((v1) | 1))
    case "":
        switch ("\n\t" + "hello") {
        case "foo bar":
            switch "foo bar" {
            case "hello":
                v22 := v4[2147483647]
                _ = v22
            case "":
                var v23 uint8 = 72
                _ = v23
            case "a":
                v3 = (v3 && (v3 || false))
            }
        }
    default:
        for fi24 := 0; fi24 < 6; fi24++ {
            if false {
                v14 = v14
                switch "" {
                case "\n\t":
                    v15 = v15
                case "hello":
                    fmt.Println(v3, v1)
                case "foo bar":
                    switch "hello" {
                    case "hello":
                        v4[962] = (fi24 - 0)
                    case "":
                        for fi25 := 0; fi25 < 20; fi25++ {
                            fmt.Println(fi25)
                            if (v1 >= 1000) {
                                _ = len(v5)
                            } else {
                                v5 = append(v5, 1)
                            }
                            _ = fi25
                        }
                    default:
                        fmt.Println(v15, v14, v15)
                    }
                }
                fmt.Println(v14, v15)
            } else {
                if v3 {
                    fmt.Println(v14, v3)
                    v14 -= (v14 << ((65535) & 63))
                    if len(v5) > 0 {
                        v5[0] = 2147483648
                    }
                }
            }
            var v26 int64 = 0
            _ = fi24
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS2{F0: v14}
        for fi29 := 0; fi29 < 10; fi29++ {
            if (v3 || (((!((!(127 >= 0)) || v3)) && (1 < 259)) || false)) {
                switch "" {
                case "hello":
                    var v30 float32 = v15
                    _ = v30
                case "":
                    v31 := v4[691]
                    _ = v31
                case "foo bar":
                    v1 += v1
                default:
                    v32 := GenS3{F0: 1, F1: 65535}
                    _ = v32
                }
                var v33 float64 = v2
                _ = v33
            } else {
                var v34 int = (fi27 % (((fi29 ^ (fi29 + 1))) | 1))
                _ = v34
            }
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "" {
    case "foo bar":
        v35 := v4[2147483648]
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
