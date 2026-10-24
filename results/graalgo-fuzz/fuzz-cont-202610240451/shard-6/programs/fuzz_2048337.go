package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
type GenS4 struct { F0 uint64; F1 uint8; F2 bool }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint, p2 uint64) string {
    v4 := []int{32767, 128, 0}
    p0 -= p0
    if len(v4) > 0 {
        v4[0] = -1
    }
    var v5 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return v5
}
func genFn1() string {
    switch 42 {
    case -635:
        var v6 uint64 = 2
        _ = v6
    case 43:
        if (-33 != 100) {
            var v7 uint8 = 1
            v8 := []int{32767, 1, 16, -1}
            _ = v7
            _ = v8
        }
    }
    if ((0 >= -1) || (2 <= 127)) {
        var v9 int = -345
        v10 := GenS1{F0: (("hello" + "\n\t") + "foo bar")}
        for fi11 := 0; fi11 < 9; fi11++ {
            v9 = (-(fi11))
            var v12 int8 = 12
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    } else {
        var v13 int8 = 100
        if false {
            v14 := []int{-2147483648, 1}
            v13 += v13
            _ = v14
        }
        _ = v13
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS1{F0: "hello"}
        fmt.Println(fi15)
        _ = fi15
        _ = v16
    }
    var v17 bool = (18446744073709551615 <= 18446744073709551615)
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float64 = 0.1
        _ = fi18
        _ = v19
    }
    _ = v17
    return (("hello" + ("hello" + "a")) + (("a" + "") + "a"))
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{2147483648, 114, 0}
    var v3 uint16 = 51
    for fi20 := 0; fi20 < 15; fi20++ {
        v3 -= (^(v3))
        _ = fi20
    }
    switch ("hello" + ("" + "a")) {
    case "\n\t":
        var v21 uint8 = 3
        _ = v21
    case "":
        var v22 int = 128
        _ = v22
    }
    fmt.Println(v3)
    v3 -= v3
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS2{F0: 1, F1: 17, F2: 0}
        _ = fi23
        _ = v24
    }
    if false {
        if (55 == 4294967295) {
            v3 -= (v3 / ((v3) | 1))
        } else {
            fmt.Println(v3)
            switch 42 {
            case 571:
                v3 += 42
            case 230:
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS2{F0: 2, F1: -1, F2: 0}
                    v3 = (^(v3))
                    var v27 float64 = ((1e10 + float64(float32(0.0))) + 3.14159)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            default:
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 int8 = 42
                    _ = fi28
                    _ = v29
                }
            }
            var v30 uint64 = 253
            _ = v30
        }
        v3 -= (v3 ^ (v3 | v3))
    } else {
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 int32 = 2147483647
            if true {
                var v33 int32 = 0
                _ = v33
            }
            var v34 float64 = (1.5 + ((1.5 / 1e10) * float64((0.1 - (0.1 + (float32((0.1 + 3.14159)) / 0.1))))))
            _ = fi31
            _ = v32
            _ = v34
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            fmt.Println(fi35, v3)
            v3 += (^(v3))
            _ = fi35
        }
    }
    v3 -= v3
    if true {
        if ((1 > 2) || false) {
            if (!((!(0.0 < (0.1 + float64(((1.5 + 0.1) + ((((1e10 * 3.14159) * float32((float64(-2.25) + 3.14159))) / 1.0) * 0.1)))))) || (1.5 < 0.0))) {
                v3 -= v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi36 := 0; fi36 < 4; fi36++ {
                    v1 = append(v1, 469)
                    v3 += v3
                    _ = fi36
                }
            }
            switch "a" {
            case "hello":
                v37 := GenS2{F0: 800, F1: 42, F2: 42}
                _ = v37
            case "a":
                switch 1 {
                case 32767:
                    switch 33 {
                    case 1:
                        v3 += (v3 | v3)
                    }
                case 128:
                    v3 = (v3 << ((1) & 63))
                }
            }
        } else {
            var v38 uint = 65535
            var v39 uint16 = (^(v3))
            _ = v38
            _ = v39
        }
    } else {
        v3 += 338
    }
    switch 32768 {
    case 648:
        fmt.Println(v3)
    case 42:
        fmt.Println(v3)
    default:
        v3 = v3
    }
    v40 := GenS3{F0: ((2147483648 >= 2147483648) && true), F1: 755, F2: false}
    v3 = (^(v3))
    if true {
        var v41 int16 = 511
        v41 -= v41
        var v42 int = 128
        _ = v41
        _ = v42
    } else {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        var v43 uint16 = (^(v3))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 int8 = -79
        for fi46 := 0; fi46 < 7; fi46++ {
            v47 := GenS4{F0: 4294967295, F1: 128, F2: (-32768 <= 1000)}
            _ = fi46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    _ = len(v2)
    if ((float64(3.14159) == (0.1 - 1e10)) || true) {
        fmt.Println(v3)
        v48 := GenS5{F0: 1}
        v3 -= 719
        _ = v48
    }
    var v49 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v49
}
