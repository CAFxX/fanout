package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
func genFn0(p0 int64, p1 bool, p2 uint32) uint {
    var v4 uint64 = 0
    if p1 {
        var v5 uint16 = 828
        v5 = 427
        v6 := []int{-2147483648, 453}
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1(p0 string, p1 int64) string {
    p0 = "\n\t"
    v7 := GenS1{F0: -2147483648}
    var v8 int8 = -107
    var v9 uint64 = 720
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p0 + (p0 + (p0 + "a")))
}

func main() {
    var v1 int = -407
    var v2 uint = 412
    v3 := []int{}
    switch v1 {
    case 1:
        for fi10 := 0; fi10 < 9; fi10++ {
            v1 = 128
            v3 = append(v3, (fi10 % ((fi10) | 1)))
            _ = fi10
        }
    case 32767:
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v1 = (fi11 & (fi11 - 32768))
                fmt.Println(fi12, v1)
                _ = fi12
            }
            v1 = (fi11 * 32767)
            fmt.Println(fi11)
            _ = fi11
        }
    default:
        var v13 uint64 = 18446744073709551615
        _ = v13
    }
    switch "" {
    case "":
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch 330 {
            case -948:
                fmt.Println(v2, v2)
            default:
                for fi14 := 0; fi14 < 11; fi14++ {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        if (!(true && false)) {
                            v16 := GenS1{F0: (fi14 & (v1 & (fi15 * v1)))}
                            var v17 bool = false
                            if v17 {
                                var v18 int8 = 0
                                _ = v18
                            }
                            var v19 int32 = -1
                            _ = v16
                            _ = v17
                            _ = v19
                        }
                        var v20 float64 = float64((0.0 / 0.1))
                        var v21 float64 = v20
                        _ = fi15
                        _ = v20
                        _ = v21
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi14
                }
            }
            if (42 <= 127) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "" {
                case "\n\t":
                    var v22 float64 = ((((-2.25 / float64((float32(3.14159) * float32(1.5)))) * 0.0) + 0.1) / float64(1e10))
                    _ = v22
                case "":
                    v2 -= 4294967295
                }
                var v23 uint64 = 1
                _ = v23
            }
        }
    case "foo bar":
        v1 -= (v1 << ((v2) & 63))
    }
    if true {
        v2 += v2
        v1 += 0
        var v24 float32 = (float32(float64((0.1 / -2.25))) + 1e10)
        _ = v24
    } else {
        if ((v2 * genFn0(-516, true, 323)) == 18446744073709551615) {
            var v25 int64 = 42
            fmt.Println(v25, v1, v2)
            _ = v25
        } else {
            v26 := GenS1{F0: v1}
            for fi27 := 0; fi27 < 2; fi27++ {
                v2 += (v2 % ((v2) | 1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi27
            }
            _ = v26
        }
    }
    switch "a" {
    case "\n\t":
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 int = int(v2)
            _ = fi28
            _ = v29
        }
    case "foo bar":
        var v30 int32 = 490
        _ = v30
    default:
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 uint64 = 1
            v32 = v32
            _ = fi31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int8 = -1
        if true {
            if true {
                if len(v3) > 0 {
                    v3[0] = (fi33 + (fi33 - fi33))
                }
                for fi35 := 0; fi35 < 3; fi35++ {
                    v2 = (v2 & v2)
                    var v36 float32 = ((-2.25 / 1.5) + 0.1)
                    _ = fi35
                    _ = v36
                }
                if true {
                    v1 -= (v1 * fi33)
                }
            }
            var v37 uint32 = 2147483648
            if len(v3) > 0 {
                v3[0] = (fi33 << ((uint(v37)) & 63))
            }
            _ = v37
        } else {
            v3 = append(v3, (v1 >> ((v2) & 63)))
            if false {
                v38 := GenS1{F0: fi33}
                v39 := GenS1{F0: v1}
                var v40 int8 = -4
                _ = v38
                _ = v39
                _ = v40
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        fmt.Println(v2, v1)
        _ = fi33
        _ = v34
    }
    var v41 int32 = 1
    if (!true) {
        v42 := GenS2{F0: (v1 / ((v1) | 1)), F1: 627, F2: -730}
        var v43 int8 = 127
        fmt.Println(v41, v1, v2)
        _ = v42
        _ = v43
    } else {
        var v44 int = (v1 << ((v2) & 63))
        fmt.Println(v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v44
    }
    v3 = append(v3, v1)
    var v45 bool = true
    if v45 {
        var v46 uint32 = 2
        v47 := GenS2{F0: -2147483648, F1: 757, F2: -588}
        _ = v46
        _ = v47
    }
    v2 = v2
    var v48 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v45
    _ = v48
}
