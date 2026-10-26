package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 int16 }
func genFn0(p0 uint16, p1 uint) string {
    v5 := []int{1, -2147483648, 0, -261}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return "\n\t"
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 1
    var v3 uint = 42
    var v4 uint = (v3 & 65535)
    if ((!false) || ((v2 << ((1) & 63)) >= v2)) {
        var v6 uint64 = 218
        _ = v6
    }
    if (1 < uint32(v2)) {
        v7 := []int{}
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint = v4
            v10 := GenS1{F0: 32767, F1: 161, F2: 127}
            _ = fi8
            _ = v10
            _ = v9
        }
        if true {
            fmt.Println(v1, v2)
            v11 := GenS1{F0: 1, F1: 537, F2: 0}
            if false {
                for fi12 := 0; fi12 < 15; fi12++ {
                    switch (-(fi12)) {
                    case 2147483647:
                        var v13 int16 = 0
                        _ = v13
                    case 128:
                        if true {
                            for fi14 := 0; fi14 < 6; fi14++ {
                                var v15 int = (fi14 * -43)
                                _ = fi14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 14; fi16++ {
                                fmt.Println(fi12, v1)
                                if (!(59 <= 1)) {
                                    v2 += v2
                                }
                                var v17 float32 = (1.5 / 1.0)
                                _ = fi16
                                _ = v17
                            }
                        }
                    case 4294967295:
                        v4 = (v3 * v4)
                    default:
                        for fi18 := 0; fi18 < 2; fi18++ {
                            var v19 float32 = (0.1 * (0.1 - 3.14159))
                            if true {
                                var v20 string = ("\n\t" + "\n\t")
                                _ = v20
                            } else {
                                var v21 int16 = -1
                                var v22 bool = true
                                _ = v21
                                _ = v22
                            }
                            var v23 bool = (((209 != 18) && (!(v3 <= v4))) && ("\n\t" <= ("a" + "")))
                            switch "hello" {
                            case "\n\t":
                                var v24 uint8 = 255
                                _ = v24
                            case "foo bar":
                                fmt.Println(v1, v4, v1)
                            case "":
                                v23 = ((v1 % ((v1) | 1)) == (v1 + 0))
                            }
                            fmt.Println(v19)
                            _ = fi18
                            _ = v19
                            _ = v23
                        }
                    }
                    fmt.Println(v3, v1, fi12)
                    _ = fi12
                }
            }
            _ = v11
        }
        _ = v7
    } else {
        var v25 int32 = 338
        v26 := GenS2{F0: 17, F1: 334}
        _ = v25
        _ = v26
    }
    var v27 int32 = 2147483647
    fmt.Println(v27)
    v27 = (v27 >> (((v3 * 65535)) & 63))
    if (false && (!(((v27 - (v27 - (v27 % (((v27 / ((v27) | 1))) | 1)))) < (v27 & v27)) && (304 >= 179)))) {
        v1 = v1
        v27 -= v27
    }
    v28 := GenS1{F0: -1, F1: 725, F2: 121}
    switch 9223372036854775807 {
    case -2147483648:
        if (false && (128 <= 145)) {
            v29 := []int{127, 32768}
            v30 := GenS3{F0: 42, F1: 65535, F2: 434}
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 uint32 = 13
                v2 = (^(v2))
                if (0 <= 484) {
                    v3 += (v3 - 0)
                    for fi33 := 0; fi33 < 9; fi33++ {
                        v32 = 0
                        _ = fi33
                    }
                    var v34 int8 = 63
                    _ = v34
                } else {
                    v3 = (^(v3))
                    fmt.Println(v1, v2, v2)
                    for fi35 := 0; fi35 < 6; fi35++ {
                        var v36 float64 = 0.1
                        _ = fi35
                        _ = v36
                    }
                }
                _ = fi31
                _ = v32
            }
            _ = v29
            _ = v30
        }
    case 765:
        v37 := []int{}
        _ = v37
    default:
        v2 += v2
    }
    var v38 int32 = (-(v27))
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
}
