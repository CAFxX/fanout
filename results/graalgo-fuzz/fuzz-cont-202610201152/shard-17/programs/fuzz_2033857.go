package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint64; F1 int32; F2 int16 }
func genFn0(p0 uint64) uint16 {
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 string = ("foo bar" + "foo bar")
            var v7 int64 = 9223372036854775807
            p0 += p0
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 101
}

func main() {
    var v1 float64 = (float64((-2.25 + -2.25)) / 1.5)
    v2 := make(map[string]uint32)
    var v3 uint64 = 951
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: -2147483648, F1: false}
        _ = fi8
        _ = v9
    }
    var v10 int = -758
    if false {
        switch "" {
        case "\n\t":
            var v11 int32 = -1
            _ = v11
        case "":
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := []int{v10, (-(v10)), v10, (v10 + (v10 | (v10 - fi12)))}
                v13 = append(v13, (v10 * (fi12 % (((v10 >> ((931) & 63))) | 1))))
                _ = fi12
                _ = v13
            }
        default:
            v14 := GenS1{F0: 840, F1: false}
            _ = v14
        }
    } else {
        if false {
            if (2147483648 <= 42) {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 int32 = -2147483648
                    _ = fi15
                    _ = v16
                }
                var v17 bool = (false && false)
                var v18 int64 = -255
                _ = v17
                _ = v18
            } else {
                switch 32767 {
                case -261:
                    fmt.Println(v3, v1)
                case 2147483647:
                    v19 := GenS1{F0: -2147483648, F1: (-743 < 42)}
                    _ = v19
                case 4294967295:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := GenS1{F0: -404, F1: true}
                        v3 = (v3 ^ 1)
                        _ = fi20
                        _ = v21
                    }
                }
                v22 := GenS2{F0: 189, F1: (((false || ((v10 & (v10 >> ((1) & 63))) > v10)) || (genFn0(759) <= genFn0(v3))) || (("a" + "foo bar") < (("foo bar" + "a") + "\n\t")))}
                var v23 string = "a"
                _ = v22
                _ = v23
            }
        } else {
            v24 := GenS2{F0: 42, F1: true}
            _ = v24
        }
        v3 = v3
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 int8 = 1
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 string = "a"
        v29 := GenS3{F0: 42, F1: 1}
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        var v30 uint8 = 128
        _ = v30
    } else {
        var v31 float64 = v1
        v32 := GenS2{F0: 9223372036854775807, F1: (int(v3) >= v10)}
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int16 = 1000
        v3 -= (v3 / ((v3) | 1))
        _ = fi33
        _ = v34
    }
    if false {
        for fi35 := 0; fi35 < 6; fi35++ {
            for fi36 := 0; fi36 < 18; fi36++ {
                v3 = v3
                _ = fi36
            }
            for fi37 := 0; fi37 < 8; fi37++ {
                v10 = 128
                _ = fi37
            }
            _ = fi35
        }
    }
    _ = len(v2)
    if ("\n\t" > "a") {
        v3 = v3
    }
    switch "a" {
    case "foo bar":
        if false {
            if (int64(v3) < 711) {
                var v38 uint64 = v3
                _ = v38
            } else {
                v10 -= (v10 | v10)
                var v39 uint64 = 736
                _ = v39
            }
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint = 18446744073709551615
                var v42 int64 = 0
                var v43 bool = true
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        }
    case "\n\t":
        v44 := GenS4{F0: (v3 / (((v3 + 4294967295)) | 1)), F1: -172, F2: 32767}
        _ = v44
    case "hello":
        if true {
            v45 := v2["a"]
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
