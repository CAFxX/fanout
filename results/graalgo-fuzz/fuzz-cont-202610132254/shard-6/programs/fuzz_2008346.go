package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 int16 = 56
    var v2 bool = true
    var v3 uint8 = 1
    v4 := make(map[string]uint32)
    var v5 int32 = -2147483648
    v6 := []int{-683}
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 65535}
        fmt.Println(v3)
        fmt.Println(v7)
        _ = fi8
        _ = v9
    }
    fmt.Println(v3, v2, v2)
    var v10 float32 = v7
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 -= (-(v1))
    _ = len(v6)
    v7 = v10
    v10 = (3.14159 - v7)
    v5 = int32(v1)
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 2}
        var v13 bool = (-1 <= 0)
        _ = fi11
        _ = v12
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint = 255
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v17 int64 = -655
                _ = fi15
                _ = v16
                _ = v17
            }
            v1 = v1
            _ = fi14
        }
        var v18 uint16 = 65535
        v5 += v5
        _ = v18
    } else {
        if false {
            v2 = (65535 == 598)
            v10 -= v10
        } else {
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 float32 = ((((1.5 + (1.5 * (v10 * 1.5))) / 1e10) * v10) * ((v10 - 1.5) * ((1.5 / -2.25) / 1.0)))
                _ = fi19
                _ = v20
            }
            var v21 uint = 4294967295
            _ = v21
        }
        if true {
            var v22 int8 = 0
            var v23 uint8 = (^(v3))
            var v24 uint16 = 896
            _ = v22
            _ = v23
            _ = v24
        }
    }
    var v25 int32 = (v5 * v5)
    if ((v10 + -2.25) < (v7 * v7)) {
        if (true && (!v2)) {
            v26 := GenS1{F0: 2}
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS2{F0: 948, F1: 345, F2: 18446744073709551615}
                if v2 {
                    v29 := GenS3{F0: ("foo bar" + (("a" + "foo bar") + "foo bar"))}
                    if v2 {
                        delete(v4, ("foo bar" + "\n\t"))
                    }
                    if len(v6) > 0 {
                        v6[0] = 128
                    }
                    _ = v29
                }
                _ = fi27
                _ = v28
            }
            if (32767 > 32768) {
                var v30 int = 32768
                _ = v30
            } else {
                for fi31 := 0; fi31 < 9; fi31++ {
                    v1 += 1
                    var v32 int16 = -1
                    _ = fi31
                    _ = v32
                }
            }
            _ = v26
        }
        if (v2 && v2) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(v7, v7)
                _ = fi33
            }
        }
        switch ("" + "a") {
        case "\n\t":
            switch 4294967295 {
            case -1:
                v10 = (v10 / (((((float32(0.1) - v7) * 0.1) - v10) - v7) - v7))
            case 1:
                for fi34 := 0; fi34 < 7; fi34++ {
                    var v35 float64 = 1e10
                    v1 = -32768
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v1 -= v1
                        _ = fi36
                    }
                    _ = fi34
                    _ = v35
                }
            }
        case "hello":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "a":
            var v37 uint64 = 496
            _ = v37
        default:
            v5 = (v25 * (v5 % ((v5) | 1)))
        }
    } else {
        v2 = (!(2 == v3))
        fmt.Println(v2, v3, v25)
    }
    var v38 uint16 = 178
    for fi39 := 0; fi39 < 19; fi39++ {
        v6 = append(v6, fi39)
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
