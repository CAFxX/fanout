package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 uint32) int16 {
    v5 := GenS1{F0: -32768}
    var v6 string = "\n\t"
    v6 = v6
    var v7 int64 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := []string{"", ("foo bar" + "hello")}
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 uint16 = 1
    v4 -= v4
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 144
        _ = fi8
        _ = v9
    }
    if (v4 > (^(v4))) {
        fmt.Println(v4)
        var v10 float64 = 1.5
        if true {
            var v11 uint64 = 923
            switch 474 {
            case 831:
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 int = (fi12 % ((fi12) | 1))
                    fmt.Println(v10)
                    var v14 float32 = 1e10
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 9223372036854775807:
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint8 = 152
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v11 = (v11 | (v11 | (v11 ^ (v11 / ((v11) | 1)))))
                        v18 := GenS1{F0: -32768}
                        switch "foo bar" {
                        case "\n\t":
                            v1 = append(v1, "hello")
                        case "":
                            var v19 string = ("\n\t" + "a")
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
            case 0:
                if ((v11 | v11) < (v11 % (((^(v11))) | 1))) {
                    v10 -= ((0.0 / -2.25) / v10)
                    fmt.Println(v10, v11, v11)
                }
            default:
                for fi20 := 0; fi20 < 12; fi20++ {
                    v11 -= v11
                    var v21 int = (fi20 | (fi20 & 297))
                    v2[("a" + "\n\t")] = "\n\t"
                    _ = fi20
                    _ = v21
                }
            }
            v22 := GenS1{F0: -625}
            _ = v11
            _ = v22
        } else {
            v4 = 1
        }
        _ = v10
    }
    var v23 string = ("hello" + (("foo bar" + "") + "hello"))
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 bool = true
        if (!(-2.25 <= ((0.1 / 0.1) + ((0.1 / 1.0) + 1e10)))) {
            var v26 bool = (v25 && ((v4 + v4) >= (v4 | (v4 & (v4 ^ (v4 >> ((4294967295) & 63)))))))
            _ = v26
        } else {
            fmt.Println(fi24)
            for fi27 := 0; fi27 < 1; fi27++ {
                fmt.Println(fi27)
                v28 := GenS1{F0: 90}
                if true {
                    if v25 {
                        switch v23 {
                        case "a":
                            v4 -= v4
                        }
                    }
                    switch v23 {
                    case "a":
                        v29 := GenS1{F0: 42}
                        _ = v29
                    case "foo bar":
                        _ = len(v3)
                    default:
                        v25 = v25
                    }
                    var v30 int8 = 100
                    _ = v30
                } else {
                    v31 := GenS2{F0: v4, F1: (42 <= 18446744073709551615)}
                    var v32 float32 = ((1.5 - 3.14159) * (0.0 / -2.25))
                    v23 = v23
                    _ = v31
                    _ = v32
                }
                _ = fi27
                _ = v28
            }
            if false {
                fmt.Println(v4, fi24)
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS3{F0: v23}
                    _ = fi33
                    _ = v34
                }
                var v35 string = "\n\t"
                _ = v35
            }
        }
        v36 := GenS3{F0: v23}
        _ = fi24
        _ = v25
        _ = v36
    }
    var v37 uint = 717
    var v38 uint32 = 1
    v39 := v3[(v38 | v38)]
    v38 += 430
    v39 += (-(v39))
    v4 = (v4 << (((v37 ^ (v37 & (v37 - 42)))) & 63))
    var v40 bool = (-904 > int16(v38))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
