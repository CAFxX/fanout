package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint; F1 int64 }
func genFn0(p0 int, p1 int16, p2 int16) uint8 {
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int16 = p1
            _ = fi7
            _ = v8
        }
        v9 := []int{}
        if len(v9) > 0 {
            v9[0] = (p0 * p0)
        }
        _ = v9
    }
    if ((true || true) && false) {
        if true {
            if false {
                var v10 int = 70
                _ = v10
            }
            switch "foo bar" {
            case "foo bar":
                p2 -= p2
            }
        } else {
            var v11 uint16 = 1
            if false {
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        var v12 float32 = (3.14159 / (1.5 * 1.5))
                        _ = v12
                    case "hello":
                        for fi13 := 0; fi13 < 3; fi13++ {
                            for fi14 := 0; fi14 < 4; fi14++ {
                                switch 42 {
                                case 762:
                                    var v15 float32 = 0.0
                                    _ = v15
                                case 97:
                                    switch -1 {
                                    case 4294967295:
                                        v16 := GenS1{F0: (-(fi14)), F1: ("hello" + ("\n\t" + ("foo bar" + "a"))), F2: 2}
                                        _ = v16
                                    case -110:
                                        v17 := []int{(-(fi14))}
                                        _ = v17
                                    default:
                                        p1 = p1
                                    }
                                case 9223372036854775807:
                                    var v18 string = ("a" + "\n\t")
                                    _ = v18
                                }
                                v19 := GenS1{F0: (fi13 | (-(p0))), F1: "a", F2: v11}
                                v20 := []int{}
                                _ = fi14
                                _ = v19
                                _ = v20
                            }
                            p1 -= (-(p2))
                            _ = fi13
                        }
                    }
                    v21 := GenS2{F0: ("hello" + (("foo bar" + "hello") + "")), F1: (v11 / ((707) | 1))}
                    _ = v21
                }
            }
            _ = v11
        }
    } else {
        v22 := GenS2{F0: ("" + ((("hello" + "foo bar") + ("a" + ((("" + "\n\t") + "foo bar") + ("\n\t" + ("" + ("foo bar" + ("\n\t" + "\n\t"))))))) + "a")), F1: 65535}
        v23 := []int{p0, (^(p0))}
        _ = v22
        _ = v23
    }
    switch 9223372036854775807 {
    case 42:
        if (!(!(true && (true && false)))) {
            var v24 int16 = 515
            fmt.Println(p2)
            _ = v24
        } else {
            v25 := []int{(p0 >> ((65535) & 63)), p0, p0}
            for fi26 := 0; fi26 < 5; fi26++ {
                _ = len(v25)
                p1 += (p2 << ((866) & 63))
                var v27 uint32 = 2147483647
                _ = fi26
                _ = v27
            }
            var v28 int = (p0 << ((166) & 63))
            _ = v25
            _ = v28
        }
    case 2147483647:
        fmt.Println(p0)
    }
    switch (p0 << ((4294967295) & 63)) {
    case 451:
        if true {
            var v29 uint16 = 2
            v30 := GenS2{F0: "a", F1: v29}
            v29 += v29
            _ = v29
            _ = v30
        }
    default:
        v31 := []int{(p0 | (p0 / ((p0) | 1))), -663, p0}
        _ = v31
    }
    switch ("a" + "hello") {
    case "":
        var v32 bool = false
        _ = v32
    case "foo bar":
        v33 := GenS1{F0: (p0 * 0), F1: "hello", F2: 0}
        _ = v33
    default:
        p0 = -351
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int16) uint {
    for fi34 := 0; fi34 < 4; fi34++ {
        p0 += p0
        _ = fi34
    }
    p0 = (p0 | -584)
    var v35 uint64 = 654
    v36 := []int{-9223372036854775808, 32768}
    _ = p0
    _ = v35
    _ = v36
    return 4294967295
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{0, 795}
    var v3 int = -937
    var v4 uint8 = 42
    var v5 int64 = -9223372036854775808
    var v6 float32 = float32(float64((0.0 + 0.0)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= v3
    var v37 uint32 = 68
    var v38 int64 = v5
    v6 = v6
    v38 = 9223372036854775807
    v39 := GenS3{F0: 65535, F1: "\n\t", F2: v4}
    if true {
        if (2147483647 < -653) {
            v40 := GenS4{F0: false, F1: -1}
            _ = v40
        } else {
            for fi41 := 0; fi41 < 3; fi41++ {
                var v42 uint64 = 42
                v2 = append(v2, 42)
                v42 -= v42
                _ = fi41
                _ = v42
            }
        }
        var v43 uint8 = 110
        var v44 int8 = -128
        _ = v43
        _ = v44
    } else {
        v45 := GenS4{F0: (true && (false && false)), F1: 806}
        _ = v45
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        if len(v2) > 0 {
            v2[0] = fi46
        }
        if true {
            fmt.Println(v3, v3)
            v1 -= v1
            v3 += 2147483647
        }
        _ = fi46
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        var v48 uint32 = v37
        _ = fi47
        _ = v48
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            v4 -= (v4 | (v4 >> ((2) & 63)))
            switch v3 {
            case 127:
                v49 := GenS5{F0: v3, F1: 4294967295, F2: 100}
                _ = v49
            case 2147483647:
                var v50 int64 = v5
                _ = v50
            case 588:
                fmt.Println(v38)
            }
        }
    case "hello":
        if true {
            if true {
                v51 := GenS6{F0: ("hello" + ""), F1: v37}
                _ = v51
            }
            fmt.Println(v5)
        }
    }
    v37 -= (v37 << ((genFn1(0)) & 63))
    var v52 uint = uint(v37)
    _ = len(v2)
    v37 = 186
    switch "\n\t" {
    case "a":
        var v53 bool = false
        _ = v53
    case "\n\t":
        v6 += float32(-2.25)
    case "foo bar":
        var v54 int8 = 1
        _ = v54
    }
    v5 += (v5 + (v38 << ((690) & 63)))
    if (genFn1(340) < genFn1(447)) {
        v55 := GenS7{F0: (v52 & v52), F1: -764}
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
