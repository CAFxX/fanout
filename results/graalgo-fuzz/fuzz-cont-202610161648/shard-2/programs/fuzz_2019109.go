package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    var v4 string = ""
    _ = p0
    _ = v4
    return 127
}
func genFn1(p0 string, p1 int8, p2 uint) string {
    p0 = ("a" + p0)
    v5 := []int{}
    v6 := GenS1{F0: "", F1: p0, F2: int(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := []string{}
    var v2 float64 = (3.14159 - 0.0)
    v3 := []string{"\n\t", "a", (("a" + "hello") + (("" + ((("hello" + (("a" + "hello") + ((("foo bar" + "hello") + "\n\t") + ("foo bar" + "")))) + "\n\t") + "foo bar")) + "a")), "a"}
    v1 = append(v1, ("hello" + "a"))
    fmt.Println(v2)
    v2 = float64((1.5 * (0.0 - 1e10)))
    _ = len(v3)
    var v7 int = 653
    fmt.Println(v2, v2)
    switch ("" + "hello") {
    case "hello":
        fmt.Println(v2, v7)
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        var v8 float64 = (-2.25 - v2)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            if false {
                if (42 >= 127) {
                    v10 := GenS1{F0: "", F1: (("\n\t" + "") + ("" + "hello")), F2: (fi9 / (((fi9 / (((-(v7))) | 1))) | 1))}
                    _ = v10
                }
                fmt.Println(fi9, fi9, fi9)
            }
            v1 = append(v1, ("foo bar" + "hello"))
            var v11 int32 = 2147483647
            _ = fi9
            _ = v11
        }
    }
    v2 -= 1.5
    for fi12 := 0; fi12 < 16; fi12++ {
        switch "a" {
        case "a":
            var v13 float64 = v2
            _ = v13
        default:
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS1{F0: ("\n\t" + ""), F1: "hello", F2: (v7 % ((128) | 1))}
                _ = fi14
                _ = v15
            }
        }
        v2 = v2
        v7 += -305
        _ = fi12
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(fi16, fi16, v7)
        fmt.Println(fi16, fi16, v7)
        var v17 float32 = 1e10
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 1
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            if true {
                v21 := GenS1{F0: (("hello" + "foo bar") + ("\n\t" + ((("\n\t" + ("foo bar" + "")) + "hello") + "\n\t"))), F1: ("foo bar" + "foo bar"), F2: fi20}
                var v22 uint64 = 18446744073709551615
                var v23 string = "\n\t"
                _ = v21
                _ = v22
                _ = v23
            } else {
                v24 := GenS1{F0: ("foo bar" + "a"), F1: "hello", F2: 327}
                _ = v24
            }
            fmt.Println(v18, fi20)
            if false {
                if false {
                    v25 := GenS1{F0: "\n\t", F1: ("foo bar" + ((("foo bar" + "hello") + "\n\t") + (("foo bar" + "a") + ("\n\t" + ("\n\t" + "foo bar"))))), F2: (fi19 + (fi20 ^ -482))}
                    _ = v25
                } else {
                    fmt.Println(fi19)
                    switch (("hello" + "foo bar") + ("a" + ("hello" + "hello"))) {
                    case "":
                        var v26 uint8 = 127
                        _ = v26
                    case "a":
                        switch fi19 {
                        case 32767:
                            var v27 int8 = -128
                            _ = v27
                        case -2147483648:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v1) > 0 {
                                v1[0] = ("foo bar" + "foo bar")
                            }
                            var v28 int32 = 525
                            _ = v28
                        case -342:
                            if false {
                                v29 := GenS1{F0: "", F1: "", F2: -807}
                                v7 -= fi20
                                _ = v29
                            }
                        default:
                            v18 = v18
                        }
                    case "hello":
                        _ = len(v1)
                        var v30 bool = false
                        _ = v30
                    }
                }
                var v31 bool = (!((float64((1.5 * (1e10 * float32(v2)))) - v2) < (float64((float32(0.0) + 1e10)) - (((v2 / ((v2 + v2) * v2)) - v2) * float64(0.0)))))
                v7 -= (v7 << ((1) & 63))
                _ = v31
            }
            _ = fi20
        }
        v2 -= 0.1
        fmt.Println(v18, v18)
        _ = fi19
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        _ = len(v1)
        if (!true) {
            fmt.Println(v18, v2, fi32)
            switch (v7 << ((723) & 63)) {
            case 1:
                var v33 uint16 = 65535
                _ = v33
            case -1:
                var v34 bool = (!(!(false || false)))
                _ = v34
            case 127:
                var v35 int = (-(v7))
                _ = v35
            default:
                for fi36 := 0; fi36 < 10; fi36++ {
                    var v37 float64 = v2
                    if (!false) {
                        v38 := GenS2{F0: -128, F1: 1}
                        fmt.Println(v7)
                        _ = v38
                    } else {
                        var v39 int8 = 0
                        v2 -= v37
                        switch 0 {
                        case -9223372036854775808:
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        case 32768:
                            if false {
                                var v40 int64 = int64(fi32)
                                var v41 string = ("hello" + ("a" + ""))
                                _ = v40
                                _ = v41
                            }
                        }
                        _ = v39
                    }
                    v42 := GenS3{F0: 0, F1: 32768}
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    _ = fi36
                    _ = v37
                    _ = v42
                }
            }
        } else {
            switch fi32 {
            case 506:
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            case -2147483648:
                for fi43 := 0; fi43 < 20; fi43++ {
                    if (7 != 127) {
                        var v44 int16 = 0
                        v44 += (v44 & -516)
                        _ = v44
                    } else {
                        v18 -= (v18 ^ v18)
                    }
                    _ = fi43
                }
            }
            v45 := GenS2{F0: -124, F1: 255}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v45
        }
        _ = fi32
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v46 := GenS4{F0: 2, F1: 65535, F2: 42}
        _ = v46
    }
    v47 := GenS5{F0: -188}
    switch "\n\t" {
    case "\n\t":
        v48 := GenS6{F0: 0}
        _ = v48
    case "hello":
        var v49 int64 = 832
        _ = v49
    default:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v47
    _ = v7
}
