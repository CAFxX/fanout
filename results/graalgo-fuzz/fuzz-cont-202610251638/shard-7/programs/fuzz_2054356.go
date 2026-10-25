package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
func genFn0(p0 uint64) int32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 = (p0 ^ (p0 & p0))
            p0 = (p0 / ((1) | 1))
            v9 := []int{(fi8 ^ -9223372036854775808), 128, (fi8 ^ fi7), (fi8 & fi8)}
            _ = fi8
            _ = v9
        }
        if false {
            for fi10 := 0; fi10 < 11; fi10++ {
                switch (fi7 % (((fi10 % (((fi7 % ((fi7) | 1))) | 1))) | 1)) {
                case 127:
                    var v11 uint32 = 2147483648
                    _ = v11
                case -2147483648:
                    var v12 float32 = (((-2.25 - 1.5) / (1e10 / 1.5)) - 0.0)
                    p0 -= p0
                    _ = v12
                case 32768:
                    v13 := []int{fi10, fi10, (fi10 % ((177) | 1)), (-(fi7))}
                    _ = v13
                }
                _ = fi10
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                p0 = p0
                if (!(2147483647 == 42)) {
                    v15 := []int{(fi7 + (fi14 << ((255) & 63))), 2147483647, (fi7 % ((4294967295) | 1))}
                    fmt.Println(p0, p0)
                    _ = v15
                } else {
                    var v16 int64 = 1
                    _ = v16
                }
                var v17 string = "\n\t"
                _ = fi14
                _ = v17
            }
        }
        var v18 bool = (!false)
        _ = fi7
        _ = v18
    }
    v19 := GenS1{F0: 100, F1: 1, F2: ((32767 == 2) && (!false))}
    switch -817 {
    case 144:
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS1{F0: 50, F1: 42, F2: false}
            _ = fi20
            _ = v21
        }
    case 0:
        v22 := []int{-172, -9223372036854775808}
        _ = v22
    default:
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint64 = 42
            for fi25 := 0; fi25 < 15; fi25++ {
                v26 := GenS1{F0: 91, F1: 542, F2: true}
                v27 := GenS1{F0: 127, F1: 1, F2: (42 <= 629)}
                v24 = 18446744073709551615
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    }
    _ = p0
    _ = v19
    return 0
}

func main() {
    var v1 bool = true
    var v2 int32 = 2147483647
    var v3 uint = 0
    v4 := make(map[string]string)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]int)
    v28 := GenS1{F0: 1, F1: -9223372036854775808, F2: v1}
    var v29 bool = v1
    switch 9223372036854775807 {
    case -640:
        var v30 uint = v3
        _ = v30
    case 1:
        switch 60 {
        case -1:
            v3 = 783
        default:
            var v31 uint = (v3 * v3)
            _ = v31
        }
    case -110:
        for fi32 := 0; fi32 < 15; fi32++ {
            v3 = v3
            v6[816] = fi32
            var v33 int = (fi32 << (((v3 >> ((v3) & 63))) & 63))
            _ = fi32
            _ = v33
        }
    default:
        v34 := []int{32767, 155, -9223372036854775808, 1}
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            switch "" {
            case "a":
                v37 := GenS2{F0: 42, F1: 2147483648}
                _ = v37
            case "foo bar":
                v1 = v1
            }
            v1 = (!true)
            v38 := []int{fi35}
            _ = fi36
            _ = v38
        }
        fmt.Println(v3)
        if (((((float64(0.1) * 0.1) / (float64((1.5 / 1e10)) + -2.25)) + 0.0) - 3.14159) != 0.1) {
            var v39 uint32 = 42
            v40 := GenS1{F0: -1, F1: -842, F2: true}
            _ = v39
            _ = v40
        } else {
            v41 := GenS2{F0: 42, F1: 959}
            _ = v41
        }
        _ = fi35
    }
    switch "" {
    case "":
        fmt.Println(v2, v29, v3)
    case "\n\t":
        for fi42 := 0; fi42 < 5; fi42++ {
            if ((-875 <= 1) || (v3 == (v3 << (((v3 ^ (^(v3)))) & 63)))) {
                v43 := GenS2{F0: 205, F1: 1}
                _ = v43
            }
            switch "a" {
            case "":
                v44 := GenS2{F0: 1, F1: 42}
                _ = v44
            case "foo bar":
                v2 = (v2 - (v2 % ((1) | 1)))
            default:
                for fi45 := 0; fi45 < 12; fi45++ {
                    if v29 {
                        var v46 uint = 42
                        _ = v46
                    }
                    fmt.Println(fi45, fi45)
                    v29 = (42 >= 1000)
                    _ = fi45
                }
            }
            _ = fi42
        }
    case "hello":
        var v47 int = -684
        _ = v47
    default:
        var v48 string = "foo bar"
        _ = v48
    }
    var v49 int64 = 42
    v50 := GenS3{F0: 42, F1: (v2 % (((v2 >> ((0) & 63))) | 1)), F2: v3}
    for fi51 := 0; fi51 < 16; fi51++ {
        var v52 uint = (v3 / ((v3) | 1))
        _ = fi51
        _ = v52
    }
    var v53 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v6
}
