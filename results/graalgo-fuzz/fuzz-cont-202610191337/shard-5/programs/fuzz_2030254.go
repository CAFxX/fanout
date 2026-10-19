package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 int64; F1 string; F2 int8 }
func genFn0(p0 int, p1 int16, p2 int) int32 {
    p2 += p2
    p0 += p2
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 uint8, p1 uint16, p2 int32) uint {
    var v7 uint8 = p0
    var v8 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = (-2.25 + 1.5)
    v2 := []int{-2147483648, 953}
    var v3 string = "foo bar"
    var v4 int8 = 1
    var v5 uint = 42
    v6 := []uint32{983, 1, 0}
    switch 4294967295 {
    case -98:
        var v9 float32 = v1
        _ = v9
    case -415:
        if false {
            if len(v6) > 0 {
                v6[0] = 2
            }
            for fi10 := 0; fi10 < 10; fi10++ {
                switch (v3 + "") {
                case "foo bar":
                    v3 = v3
                case "a":
                    var v11 int32 = -2147483648
                    _ = v11
                default:
                    var v12 string = "\n\t"
                    _ = v12
                }
                _ = fi10
            }
            switch "foo bar" {
            case "a":
                var v13 int64 = -1
                _ = v13
            case "foo bar":
                v5 += genFn1(151, 739, genFn0(-9223372036854775808, 0, 32768))
            }
        }
    case -9223372036854775808:
        if true {
            v14 := GenS1{F0: 1, F1: -811, F2: 127}
            var v15 int32 = 725
            _ = v14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v5, v4, fi16)
            var v17 float32 = (v1 * 0.0)
            switch 0 {
            case 128:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 109:
                for fi18 := 0; fi18 < 17; fi18++ {
                    if (!(false || false)) {
                        var v19 int32 = -2147483648
                        v1 += v17
                        var v20 float32 = (v1 - v1)
                        _ = v19
                        _ = v20
                    }
                    v17 = float32(float64((1e10 + ((1.5 + (v17 * 1e10)) / float32((0.1 / -2.25))))))
                    _ = fi18
                }
            case 32768:
                v2 = append(v2, -667)
            default:
                fmt.Println(v4, v1)
            }
            _ = fi16
            _ = v17
        }
    }
    v3 = v3
    v1 -= ((v1 - -2.25) / (0.1 / ((v1 - v1) * v1)))
    v1 = v1
    v1 += 0.0
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v4 -= (v4 - 0)
            _ = fi22
        }
        _ = fi21
    }
    var v23 float64 = 1.5
    var v24 uint8 = 2
    var v25 int = 2147483647
    if (!(v25 != 1)) {
        for fi26 := 0; fi26 < 11; fi26++ {
            v4 += (v4 / ((v4) | 1))
            var v27 string = v3
            v28 := GenS1{F0: -104, F1: 9223372036854775807, F2: fi26}
            _ = fi26
            _ = v27
            _ = v28
        }
        if true {
            var v29 int16 = 1
            fmt.Println(v25)
            v25 += v25
            _ = v29
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                v4 = v4
                v1 -= v1
                _ = fi30
            }
        }
        v23 += v23
    }
    v25 -= (v25 + v25)
    v31 := GenS2{F0: -1, F1: v3, F2: (v4 ^ v4)}
    switch ("\n\t" + v3) {
    case "foo bar":
        var v32 int16 = -32768
        _ = v32
    case "hello":
        v33 := GenS1{F0: -1, F1: 0, F2: v25}
        _ = v33
    }
    if true {
        for fi34 := 0; fi34 < 17; fi34++ {
            v5 = v5
            _ = fi34
        }
    } else {
        var v35 int32 = 42
        switch "\n\t" {
        case "foo bar":
            v3 = v3
        case "":
            for fi36 := 0; fi36 < 17; fi36++ {
                switch (fi36 * (fi36 >> (((v5 & (v5 << ((v5) & 63)))) & 63))) {
                case -1:
                    v4 = (v4 * -88)
                case -132:
                    fmt.Println(v5, v24)
                }
                _ = fi36
            }
        case "a":
            v2 = append(v2, (v25 | v25))
        }
        _ = v35
    }
    v37 := GenS2{F0: -1, F1: "foo bar", F2: v4}
    if (!(float64(v1) < (v23 / float64((((v1 - v1) + v1) + ((((1e10 / 0.1) - -2.25) + v1) - float32(0.0))))))) {
        var v38 int64 = -910
        v39 := GenS1{F0: 0, F1: -9223372036854775808, F2: v25}
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
