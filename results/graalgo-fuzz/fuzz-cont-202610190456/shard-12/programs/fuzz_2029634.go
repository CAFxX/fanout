package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int) uint64 {
    if (1 >= 1) {
        for fi4 := 0; fi4 < 8; fi4++ {
            v5 := GenS1{F0: (fi4 - (^(fi4)))}
            p0 += (fi4 << ((559) & 63))
            _ = fi4
            _ = v5
        }
    } else {
        fmt.Println(p0)
        v6 := GenS1{F0: p0}
        for fi7 := 0; fi7 < 14; fi7++ {
            switch 1000 {
            case 4294967295:
                var v8 uint8 = 127
                _ = v8
            case 1:
                v9 := GenS2{F0: "foo bar", F1: -448}
                _ = v9
            default:
                p0 = p0
            }
            v10 := GenS3{F0: 201}
            var v11 uint64 = 2
            _ = fi7
            _ = v10
            _ = v11
        }
        _ = v6
    }
    var v12 uint8 = 128
    v12 -= 127
    if (!true) {
        v12 = v12
        v12 = v12
    } else {
        if true {
            var v13 int64 = -784
            fmt.Println(p0, v13, v12)
            _ = v13
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := []int{(^(p0)), p0, (fi14 ^ (fi14 >> ((211) & 63))), p0}
                _ = fi14
                _ = v15
            }
            p0 -= (p0 | p0)
        }
    }
    v16 := []int{}
    _ = p0
    _ = v12
    _ = v16
    return 4294967295
}

func main() {
    var v1 uint = 65535
    var v2 int32 = 266
    var v3 int64 = 0
    var v17 int = 42
    var v18 uint16 = 1
    v3 -= 1
    v3 += v3
    v17 = v17
    for fi19 := 0; fi19 < 3; fi19++ {
        switch "\n\t" {
        case "\n\t":
            for fi20 := 0; fi20 < 5; fi20++ {
                if true {
                    v1 = (v1 + v1)
                    if (v3 >= (v3 & (v3 << ((18446744073709551615) & 63)))) {
                        fmt.Println(v3)
                    }
                    if (false || true) {
                        v21 := []int{fi20, -672}
                        var v22 bool = ((!true) && (42 <= 602))
                        fmt.Println(v17, v18)
                        v22 = false
                        _ = len(v21)
                        var v23 int8 = 1
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        v24 := []int{fi19, (^(fi19)), 127}
                        switch "\n\t" {
                        case "\n\t":
                            if len(v24) > 0 {
                                v24[0] = int(v3)
                            }
                        default:
                            var v25 uint16 = 65535
                            var v26 uint = 2
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS2{F0: "foo bar", F1: (v3 & 1)}
                        fmt.Println(fi19)
                        _ = v24
                        _ = v27
                    }
                }
                if ((((("a" + "") + ("\n\t" + (("hello" + "a") + ("a" + ("foo bar" + "foo bar"))))) + (("a" + "foo bar") + "\n\t")) < "hello") || (true || ((v3 ^ 1) == v3))) {
                    switch "\n\t" {
                    case "hello":
                        v28 := GenS1{F0: int(v18)}
                        _ = v28
                    }
                }
                var v29 int64 = v3
                v30 := GenS4{F0: 260}
                switch ("" + ((("hello" + "") + "a") + ("a" + "foo bar"))) {
                case "foo bar":
                    var v31 float64 = (-2.25 * 1e10)
                    _ = v31
                case "a":
                    v1 -= v1
                default:
                    v32 := []int{fi19}
                    _ = v32
                }
                _ = fi20
                _ = v29
                _ = v30
            }
        case "":
            for fi33 := 0; fi33 < 15; fi33++ {
                if (false || (((("foo bar" + "") + "") + ("a" + "\n\t")) <= "hello")) {
                    v3 -= v3
                    var v34 int = v17
                    _ = v34
                } else {
                    v18 -= (v18 - (v18 & v18))
                    v1 += 121
                    var v35 float64 = ((1e10 - 1e10) / 3.14159)
                    _ = v35
                }
                if true {
                    v36 := GenS1{F0: (v17 | v17)}
                    var v37 int8 = -73
                    _ = v36
                    _ = v37
                }
                _ = fi33
            }
        default:
            v38 := []int{-281, (-(fi19)), 216, (v17 % ((131) | 1))}
            _ = v38
        }
        _ = fi19
    }
    var v39 uint8 = 127
    v3 = (v3 & 549)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
}
