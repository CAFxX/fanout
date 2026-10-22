package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64, p2 int8) int16 {
    v7 := GenS1{F0: (p1 >> ((255) & 63))}
    if p0 {
        p2 += (p2 - (p2 - p2))
        v8 := GenS1{F0: p1}
        p2 = (p2 % (((p2 + p2)) | 1))
        _ = v8
    } else {
        var v9 int16 = 42
        v9 = (-(v9))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 int64, p1 int32, p2 bool) string {
    switch ("" + ("foo bar" + "a")) {
    case "":
        if p2 {
            var v10 string = ("foo bar" + ("" + ""))
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 bool = true
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := []int{(^(fi13)), (fi13 << ((18446744073709551615) & 63))}
                    p1 = p1
                    if len(v14) > 0 {
                        v14[0] = (fi13 | (fi13 << ((918) & 63)))
                    }
                    _ = fi13
                    _ = v14
                }
                v15 := GenS1{F0: 986}
                _ = fi11
                _ = v12
                _ = v15
            }
            fmt.Println(v10, p2, v10)
            _ = v10
        }
    case "\n\t":
        v16 := []int{}
        _ = v16
    case "foo bar":
        v17 := []int{128, -9223372036854775808, 1, 760}
        _ = v17
    }
    if ((1e10 == (3.14159 * 3.14159)) && (1e10 >= 1.5)) {
        var v18 int32 = p1
        _ = v18
    } else {
        if p2 {
            p1 += p1
        } else {
            var v19 uint16 = 42
            _ = v19
        }
    }
    p0 += p0
    v20 := []int{0}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return "foo bar"
}

func main() {
    v1 := []uint32{65535}
    var v2 uint16 = 65535
    var v3 uint64 = 181
    v4 := []uint32{}
    var v5 int64 = -9223372036854775808
    v6 := []string{"foo bar", "foo bar", (("hello" + "\n\t") + (("\n\t" + (("a" + "") + "a")) + (("" + "a") + "")))}
    var v21 int64 = v5
    v3 -= 42
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 float64 = (3.14159 * 0.0)
        _ = fi22
        _ = v23
    }
    v5 += v5
    v5 -= (v5 ^ v5)
    var v24 bool = false
    switch 2147483647 {
    case 42:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v25 := GenS2{F0: v2, F1: 255, F2: 21}
        _ = v25
    }
    v5 -= (v21 << ((739) & 63))
    for fi26 := 0; fi26 < 12; fi26++ {
        fmt.Println(v2)
        v27 := GenS3{F0: 111}
        switch fi26 {
        case -1:
            for fi28 := 0; fi28 < 10; fi28++ {
                if len(v4) > 0 {
                    v4[0] = 740
                }
                v29 := GenS1{F0: v3}
                _ = len(v1)
                var v30 uint = 4294967295
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            v2 = (v2 ^ v2)
        }
        _ = fi26
        _ = v27
    }
    if ((-(v21)) < int64(v3)) {
        var v31 float64 = 0.0
        var v32 uint = 65535
        if (!v24) {
            var v33 int16 = 42
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v33
        } else {
            var v34 string = "foo bar"
            v32 = (v32 - 195)
            v35 := GenS4{F0: 910, F1: 900}
            _ = v34
            _ = v35
        }
        _ = v31
        _ = v32
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        v37 := GenS3{F0: 1}
        var v38 int32 = 1
        _ = len(v6)
        _ = fi36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 uint32 = 671
        v24 = (42 >= genFn0((255 != 18446744073709551615), (^(v3)), -75))
        switch 9223372036854775807 {
        case 4294967295:
            if v24 {
                v41 := GenS5{F0: 4294967295}
                if len(v6) > 0 {
                    v6[0] = ""
                }
                fmt.Println(fi39, v40)
                _ = v41
            }
        case -2147483648:
            for fi42 := 0; fi42 < 11; fi42++ {
                if true {
                    v43 := GenS5{F0: v3}
                    if v24 {
                        var v44 int16 = genFn0((!v24), v3, 0)
                        _ = v44
                    }
                    _ = v43
                }
                var v45 float32 = float32(((3.14159 + ((-2.25 - (0.1 - -2.25)) + (-2.25 - 3.14159))) / 1.0))
                switch fi42 {
                case -407:
                    var v46 int8 = 0
                    _ = v46
                case 1:
                    v3 -= (v3 >> ((712) & 63))
                default:
                    v2 -= v2
                }
                _ = fi42
                _ = v45
            }
        default:
            v3 -= (^(v3))
        }
        _ = fi39
        _ = v40
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        var v48 int64 = v5
        _ = fi47
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
