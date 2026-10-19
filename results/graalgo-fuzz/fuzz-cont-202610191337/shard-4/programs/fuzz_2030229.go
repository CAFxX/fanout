package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32, p1 string) string {
    if ((p0 ^ (p0 ^ (^(p0)))) >= p0) {
        p1 = (("" + p1) + p1)
        var v5 uint = 42
        _ = v5
    } else {
        var v6 uint8 = 1
        v7 := GenS1{F0: ((0 >= 32768) || true), F1: 1}
        _ = v6
        _ = v7
    }
    p0 -= -2147483648
    for fi8 := 0; fi8 < 14; fi8++ {
        switch fi8 {
        case 80:
            var v9 uint = 96
            _ = v9
        case 32768:
            p1 = (p1 + p1)
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    return (((p1 + p1) + "") + p1)
}

func main() {
    var v1 uint8 = 127
    var v2 uint8 = (^(v1))
    var v3 uint32 = uint32(v1)
    var v4 uint16 = 65535
    v1 += v2
    v3 -= v3
    var v10 float64 = (float64(1.5) - -2.25)
    var v11 uint64 = 795
    var v12 uint64 = v11
    v13 := []int{1}
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: false, F1: 929}
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint8 = (v2 + (v2 >> ((2) & 63)))
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        switch "foo bar" {
        case "a":
            v11 = 4294967295
        case "":
            if (1 < 1) {
                v19 := GenS1{F0: false, F1: 75}
                v12 = v11
                var v20 uint8 = 105
                _ = v19
                _ = v20
            }
        }
        fmt.Println(fi18)
        v10 += ((v10 + (v10 - v10)) + ((1e10 / ((0.1 + v10) - v10)) + v10))
        _ = fi18
    }
    var v21 uint = 4294967295
    switch -845 {
    case 544:
        if (true || ((^(v12)) != 42)) {
            var v22 uint16 = 42
            if true {
                v23 := GenS2{F0: 168}
                var v24 float32 = -2.25
                _ = v23
                _ = v24
            } else {
                switch 9223372036854775807 {
                case 32767:
                    if (528 >= v3) {
                        var v25 int16 = 32767
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v1 -= (v2 >> ((uint(v22)) & 63))
                            for fi27 := 0; fi27 < 10; fi27++ {
                                v28 := GenS1{F0: (uint16(v3) != (v4 * 1)), F1: v21}
                                var v29 float32 = float32((v10 + v10))
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            _ = fi26
                        }
                        _ = v25
                    }
                case -1:
                    var v30 int = -1
                    _ = v30
                case 4294967295:
                    fmt.Println(v1)
                }
                var v31 float64 = 0.1
                _ = v31
            }
            _ = v22
        } else {
            if ((!(0.0 > ((((1.5 + v10) + v10) * float64(-2.25)) + (v10 + v10)))) || false) {
                if true {
                    v1 += v1
                    switch int(v4) {
                    case 1:
                        switch 209 {
                        case 32768:
                            for fi32 := 0; fi32 < 12; fi32++ {
                                v33 := GenS2{F0: uint32(v11)}
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    var v35 uint = (v21 ^ 1)
                                    _ = fi34
                                    _ = v35
                                }
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                _ = fi32
                                _ = v33
                            }
                        case 409:
                            v4 += 0
                        }
                    case 128:
                        if len(v13) > 0 {
                            v13[0] = 128
                        }
                    default:
                        fmt.Println(v12)
                    }
                    for fi36 := 0; fi36 < 10; fi36++ {
                        v13 = append(v13, 9223372036854775807)
                        _ = fi36
                    }
                }
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 int64 = 0
                    if len(v13) > 0 {
                        v13[0] = (fi37 ^ fi37)
                    }
                    _ = fi37
                    _ = v38
                }
            } else {
                var v39 int32 = -1
                _ = v39
            }
            v40 := GenS1{F0: ((v11 <= 2) || true), F1: 0}
            var v41 uint8 = 1
            _ = v40
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
