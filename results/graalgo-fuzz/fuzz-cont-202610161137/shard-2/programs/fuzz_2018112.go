package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 bool }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 uint16, p2 string) uint16 {
    v6 := GenS1{F0: 255, F1: 4294967295, F2: true}
    if false {
        var v7 uint = 4294967295
        fmt.Println(p2, v7)
        var v8 uint64 = 2
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(p0, p1)
            v10 := GenS1{F0: 1, F1: 2, F2: ("\n\t" <= (p2 + ("foo bar" + p2)))}
            v11 := []int{fi9, fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint16 = p1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 42
    var v3 int8 = 127
    var v4 uint8 = 104
    var v5 uint8 = 255
    v5 = (v5 - v5)
    v13 := []int{128, 893, -645, -847}
    v2 = v2
    if len(v13) > 0 {
        v13[0] = -332
    }
    if true {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint = 0
            v3 = (v3 >> ((v15) & 63))
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS1{F0: 936, F1: v2, F2: false}
            _ = fi16
            _ = v17
        }
    } else {
        _ = len(v13)
        if true {
            var v18 int = int(v4)
            var v19 uint8 = 252
            _ = v18
            _ = v19
        }
    }
    if (32767 >= 23) {
        v2 = (v2 << ((786) & 63))
    }
    if (!(!false)) {
        if true {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 int = fi20
                var v22 uint32 = 2147483648
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            var v23 uint8 = (v5 << ((65535) & 63))
            _ = v23
        }
        v3 = v3
    }
    v24 := GenS1{F0: 65535, F1: (v2 << ((401) & 63)), F2: (((float64(-2.25) - (((float64(3.14159) * (3.14159 * float64((1.5 + 1e10)))) + (((float64(0.0) * (float64((0.1 + (((0.1 + 0.0) + (3.14159 + 3.14159)) + -2.25))) * 1.5)) + 1.5) / 1e10)) + 0.1)) >= 0.1) && (429 > 0))}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = len(v13)
    v25 := GenS2{F0: ("\n\t" + "\n\t")}
    if false {
        for fi26 := 0; fi26 < 6; fi26++ {
            if false {
                for fi27 := 0; fi27 < 9; fi27++ {
                    switch -1 {
                    case 32768:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        if (false || false) {
                            v13 = append(v13, 1)
                        }
                        v5 += v5
                    case 42:
                        v13 = append(v13, fi26)
                    }
                    _ = fi27
                }
                v3 += (v3 % ((v3) | 1))
                for fi28 := 0; fi28 < 3; fi28++ {
                    var v29 int64 = -1
                    _ = fi28
                    _ = v29
                }
            }
            v30 := GenS1{F0: 65535, F1: 2, F2: (true || true)}
            fmt.Println(fi26, v4, v3)
            _ = fi26
            _ = v30
        }
        var v31 uint8 = v5
        _ = v31
    }
    switch ("hello" + "") {
    case "a":
        v4 = v5
    case "hello":
        var v32 float64 = 0.0
        _ = v32
    default:
        if true {
            v33 := GenS1{F0: 4294967295, F1: 42, F2: (v1 > (v1 * 32767))}
            var v34 float32 = 1e10
            _ = v33
            _ = v34
        }
    }
    if false {
        var v35 uint64 = (v2 * 899)
        var v36 string = "a"
        _ = v35
        _ = v36
    } else {
        v13 = append(v13, int(v3))
    }
    v3 = v3
    for fi37 := 0; fi37 < 13; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            if true {
                var v39 uint8 = (v5 | v5)
                _ = v39
            }
            _ = fi38
        }
        v2 -= v2
        var v40 int8 = v3
        _ = fi37
        _ = v40
    }
    switch "" {
    case "a":
        fmt.Println(v2, v4, v3)
    case "hello":
        v41 := GenS1{F0: 4294967295, F1: v2, F2: (false || false)}
        _ = v41
    default:
        fmt.Println(v4)
    }
    switch "\n\t" {
    case "foo bar":
        for fi42 := 0; fi42 < 7; fi42++ {
            for fi43 := 0; fi43 < 19; fi43++ {
                v4 += (^(v4))
                _ = fi43
            }
            _ = fi42
        }
    case "a":
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 int32 = int32(v1)
            var v46 int16 = (v1 - 32767)
            _ = fi44
            _ = v45
            _ = v46
        }
    default:
        v2 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
