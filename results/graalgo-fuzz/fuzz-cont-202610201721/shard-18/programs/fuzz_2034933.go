package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = float32(3.14159)
    v3 := make(map[int]int)
    var v4 int32 = -2147483648
    v5 := []uint32{42, 4294967295, 2147483648, 2147483647}
    var v6 uint64 = 292
    for fi7 := 0; fi7 < 13; fi7++ {
        if true {
            for fi8 := 0; fi8 < 9; fi8++ {
                v9 := v1[(fi7 & fi8)]
                switch 0 {
                case 4294967295:
                    v4 += (v4 - v4)
                case 32768:
                    v10 := v3[(-(fi8))]
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            for fi11 := 0; fi11 < 7; fi11++ {
                v5 = append(v5, 879)
                _ = fi11
            }
        } else {
            if (127 <= uint8(v6)) {
                for fi12 := 0; fi12 < 5; fi12++ {
                    if ((v2 <= v2) && (!(!(!false)))) {
                        v4 = v4
                        v13 := GenS1{F0: -640}
                        _ = v13
                    }
                    _ = fi12
                }
            } else {
                v2 += (3.14159 - v2)
                var v14 uint64 = v6
                _ = v14
            }
            v4 += v4
        }
        _ = fi7
    }
    if false {
        delete(v1, int(v6))
        if (v6 < v6) {
            var v15 int8 = 42
            if (595 >= 410) {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := GenS2{F0: 48, F1: v4}
                    var v18 int32 = (v4 % (((v4 >> ((18446744073709551615) & 63))) | 1))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v15 += v15
            } else {
                var v19 uint8 = 127
                _ = v19
            }
            v6 += v6
            _ = v15
        }
    } else {
        if true {
            var v20 int = 770
            _ = v20
        }
        switch "foo bar" {
        case "hello":
            for fi21 := 0; fi21 < 10; fi21++ {
                switch "foo bar" {
                case "hello":
                    v6 += v6
                case "":
                    fmt.Println(fi21, v2, fi21)
                case "foo bar":
                    fmt.Println(v4, v2)
                default:
                    var v22 uint = 943
                    _ = v22
                }
                v5 = append(v5, 4294967295)
                _ = fi21
            }
        }
        fmt.Println(v4, v4, v2)
    }
    if true {
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 bool = false
            var v25 uint8 = 114
            v3[32767] = (fi23 >> ((2) & 63))
            _ = fi23
            _ = v24
            _ = v25
        }
        v2 = (v2 + (1.5 * 0.1))
        if false {
            v6 -= 1
            if true {
                fmt.Println(v2, v6)
                fmt.Println(v2)
                if true {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        var v27 float64 = (3.14159 - 1e10)
                        v6 += 1
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = len(v1)
                }
            }
        } else {
            delete(v3, -211)
            v28 := v3[127]
            if ((-717 < int16(v4)) || true) {
                fmt.Println(v6, v2, v6)
            }
            _ = v28
        }
    } else {
        if (128 <= 128) {
            v2 += v2
            v4 -= v4
        }
        v29 := GenS1{F0: -9223372036854775808}
        switch 0 {
        case 901:
            v2 += v2
        case 32767:
            switch 12 {
            case 0:
                for fi30 := 0; fi30 < 13; fi30++ {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v32 := GenS1{F0: 1}
                        _ = fi31
                        _ = v32
                    }
                    var v33 bool = ((fi30 << ((4294967295) & 63)) != (fi30 | (fi30 % (((-(fi30))) | 1))))
                    v4 += 2147483647
                    _ = fi30
                    _ = v33
                }
            case -829:
                var v34 int8 = 42
                _ = v34
            }
        case -9223372036854775808:
            switch 32768 {
            case 647:
                delete(v1, 54)
            case 278:
                var v35 float64 = 1e10
                _ = v35
            case 4294967295:
                v36 := GenS1{F0: 0}
                _ = v36
            }
        }
        _ = v29
    }
    if (-230 == 88) {
        v37 := GenS2{F0: 389, F1: (v4 ^ 2147483647)}
        _ = v37
    }
    fmt.Println(v2, v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v38 := v3[-730]
    v39 := GenS1{F0: 9223372036854775807}
    fmt.Println(v6)
    if len(v5) > 0 {
        v5[0] = 830
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
