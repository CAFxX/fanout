package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0(p0 int64, p1 bool, p2 int) string {
    fmt.Println(p2)
    var v7 int32 = -2147483648
    v8 := []int{}
    fmt.Println(p0, p2, p2)
    if (p1 && (-32768 == 0)) {
        if (65535 < 159) {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int16 = 32767
                _ = fi9
                _ = v10
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 bool = (905 >= 929)
    v2 := []int{}
    v3 := []string{"foo bar", ("hello" + "a"), "a", ("\n\t" + ("hello" + "hello"))}
    v4 := []uint32{42, 915, 4294967295, 2}
    var v5 uint = 0
    var v6 uint32 = 228
    v11 := GenS1{F0: (v6 / (((v6 ^ (v6 / (((v6 - 2147483648)) | 1)))) | 1)), F1: 42, F2: 42}
    v1 = (v1 && true)
    v12 := GenS1{F0: uint32(v5), F1: 1000, F2: 9223372036854775807}
    var v13 uint64 = 4294967295
    fmt.Println(v1)
    fmt.Println(v6)
    for fi14 := 0; fi14 < 15; fi14++ {
        if (!(!v1)) {
            v6 = (v6 / (((v6 + (v6 & 65535))) | 1))
            switch fi14 {
            case -2147483648:
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 int16 = 1000
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi15
                    _ = v16
                }
            case -488:
                if (42 == 0) {
                    if false {
                        var v17 int16 = 0
                        var v18 uint64 = (v13 >> ((v5) & 63))
                        _ = v17
                        _ = v18
                    }
                } else {
                    v1 = (0.1 >= 0.0)
                    var v19 uint16 = 2
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 int16 = int16(v5)
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            case 0:
                var v22 uint8 = 0
                _ = v22
            default:
                var v23 int8 = 42
                _ = v23
            }
            v6 += v6
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v3) > 0 {
                v3[0] = ""
            }
            v5 += (v5 >> ((1) & 63))
        }
        for fi24 := 0; fi24 < 13; fi24++ {
            v25 := GenS1{F0: v6, F1: 1000, F2: -1}
            if (!(fi14 != 304)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi24
            _ = v25
        }
        v1 = (127 > 2)
        _ = fi14
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS1{F0: (v6 << ((1) & 63)), F1: 1000, F2: -724}
            _ = len(v2)
            _ = fi27
            _ = v28
        }
        fmt.Println(v5)
        for fi29 := 0; fi29 < 9; fi29++ {
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            if v1 {
                v13 = (v13 / ((v13) | 1))
            }
            v30 := GenS1{F0: 570, F1: 0, F2: -9223372036854775808}
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    v1 = ((true || true) && (255 > 65535))
    fmt.Println(v13, v5)
    v13 += v13
    v6 = 2
    v6 -= 168
    if false {
        if true {
            for fi31 := 0; fi31 < 2; fi31++ {
                v13 = 47
                var v32 float64 = 0.1
                _ = fi31
                _ = v32
            }
            switch "foo bar" {
            case "hello":
                fmt.Println(v5, v1)
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        var v33 uint64 = (v13 >> ((v5) & 63))
        _ = v33
    }
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
