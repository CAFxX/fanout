package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 string; F2 int64 }
type GenS8 struct { F0 int16; F1 uint }
func genFn0() int16 {
    var v5 int8 = 100
    fmt.Println(v5)
    v5 = 42
    if (1 < 127) {
        v6 := GenS1{F0: 203, F1: "hello", F2: 730}
        _ = v6
    } else {
        fmt.Println(v5)
    }
    var v7 uint8 = 0
    _ = v5
    _ = v7
    return int16(v5)
}
func genFn1(p0 int64, p1 bool, p2 uint64) uint16 {
    v8 := []int{}
    if (65535 >= 1) {
        v9 := GenS1{F0: 42, F1: "foo bar", F2: 1}
        if (p1 || ((!(171 >= genFn0())) && (!p1))) {
            p2 += p2
        } else {
            var v10 bool = (!true)
            _ = v10
        }
        v11 := GenS2{F0: 127, F1: 78, F2: 133}
        _ = v11
        _ = v9
    }
    if true {
        p2 = p2
    } else {
        v12 := GenS3{F0: 4294967295, F1: genFn0()}
        p0 = p0
        _ = v12
    }
    if ((true && p1) && ("\n\t" > "hello")) {
        fmt.Println(p0, p1, p0)
        p0 += p0
        if p1 {
            p2 = (p2 * (p2 - 18446744073709551615))
            if p1 {
                p1 = (!p1)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            fmt.Println(p0, p1, p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 496
}

func main() {
    var v1 int8 = 42
    var v2 uint = 42
    v3 := []string{"", ("a" + "")}
    v4 := make(map[uint32]bool)
    var v13 int64 = -9223372036854775808
    if ((((255 < 65535) && true) || false) || (v2 > 798)) {
        var v14 float64 = 3.14159
        _ = v14
    }
    if ((!true) && (!((-6 != (v1 / ((0) | 1))) || false))) {
        if false {
            var v15 string = "hello"
            _ = v15
        }
        v2 = (^(v2))
    } else {
        v16 := GenS4{F0: 2}
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v2, v1)
        _ = fi17
    }
    switch 9223372036854775807 {
    case -2147483648:
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 string = "hello"
            switch v19 {
            case "\n\t":
                v20 := v4[2]
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    case -33:
        v1 -= -1
    }
    v1 = v1
    if true {
        var v21 float32 = -2.25
        _ = v21
    }
    var v22 int = 2147483647
    v22 = 576
    if (1 >= 127) {
        var v23 uint32 = 2
        _ = v23
    } else {
        fmt.Println(v22, v13, v13)
        if false {
            if true {
                if true {
                    v24 := GenS5{F0: ""}
                    fmt.Println(v2, v1, v13)
                    v2 = (v2 | (v2 ^ v2))
                    _ = v24
                }
                v25 := GenS6{F0: -1}
                v13 = (v13 & v13)
                _ = v25
            } else {
                var v26 float32 = 0.0
                v27 := GenS7{F0: 856, F1: (("" + (("hello" + (("a" + "") + "a")) + (("\n\t" + ("hello" + (("" + "foo bar") + ("foo bar" + "hello")))) + ""))) + "\n\t"), F2: (v13 % ((9223372036854775807) | 1))}
                _ = len(v4)
                _ = v26
                _ = v27
            }
            v28 := GenS3{F0: 42, F1: 188}
            _ = v28
        } else {
            v1 = v1
        }
        if false {
            fmt.Println(v13, v13)
        } else {
            var v29 uint = v2
            var v30 string = (("hello" + "foo bar") + "foo bar")
            delete(v4, 42)
            _ = v29
            _ = v30
        }
    }
    fmt.Println(v1)
    v2 = v2
    for fi31 := 0; fi31 < 19; fi31++ {
        if (!(v13 > (v13 << ((v2) & 63)))) {
            v32 := GenS6{F0: v1}
            switch (v22 % ((fi31) | 1)) {
            case 4294967295:
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := GenS5{F0: ("foo bar" + ("" + "a"))}
                    v22 += fi31
                    _ = fi33
                    _ = v34
                }
            case 42:
                var v35 string = ""
                _ = v35
            }
            _ = v32
        } else {
            v1 += (v1 % (((v1 - (v1 ^ 100))) | 1))
        }
        v22 -= (v22 + (-(fi31)))
        for fi36 := 0; fi36 < 12; fi36++ {
            switch (fi31 - 32768) {
            case -834:
                switch (fi36 ^ (fi36 / (((fi36 / ((v22) | 1))) | 1))) {
                case -444:
                    fmt.Println(fi36)
                case 1:
                    fmt.Println(v13)
                case 128:
                    if false {
                        var v37 uint16 = 379
                        v22 -= (-(fi36))
                        _ = v37
                    } else {
                        fmt.Println(fi31)
                    }
                default:
                    if false {
                        switch ("foo bar" + "a") {
                        case "a":
                            v38 := v4[65535]
                            _ = v38
                        case "":
                            var v39 int32 = -1
                            _ = v39
                        case "\n\t":
                            var v40 int16 = -240
                            _ = v40
                        }
                        _ = len(v4)
                        switch "hello" {
                        case "\n\t":
                            switch (("\n\t" + "a") + "foo bar") {
                            case "hello":
                                v4[385] = (1000 <= 0)
                            default:
                                fmt.Println(fi31)
                            }
                        }
                        v2 = v2
                        v2 = (v2 - 255)
                        v41 := GenS8{F0: 1000, F1: v2}
                        fmt.Println(v22)
                        switch ("a" + "\n\t") {
                        case "":
                            switch v22 {
                            case -973:
                                v2 = (v2 << ((1) & 63))
                            case 4294967295:
                                v42 := GenS4{F0: genFn1((v13 | v13), true, 18446744073709551615)}
                                _ = v42
                            }
                        default:
                            v22 = fi36
                        }
                        _ = v41
                    }
                }
            case -612:
                v13 += 1
            case 42:
                var v43 int32 = 2147483647
                _ = v43
            }
            fmt.Println(v1, v2)
            _ = fi36
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
