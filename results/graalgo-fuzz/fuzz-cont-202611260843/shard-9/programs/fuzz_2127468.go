package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 string }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 uint8) uint {
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 -= (p0 % ((p0) | 1))
        _ = fi5
    }
    p0 -= (p0 - p0)
    v6 := []int{-9223372036854775808, 127, 127}
    var v7 int = 1
    _ = p0
    _ = v6
    _ = v7
    return 586
}

func main() {
    var v1 float64 = 1.5
    v2 := []string{"", ((("\n\t" + ("foo bar" + "a")) + "") + ((("a" + "hello") + ("\n\t" + "\n\t")) + (("hello" + "foo bar") + ("" + (("" + (("hello" + (("" + "foo bar") + ("a" + ""))) + ((("" + "hello") + "a") + ((("\n\t" + (("hello" + (("hello" + "hello") + ("foo bar" + "\n\t"))) + "hello")) + "\n\t") + ("hello" + "a"))))) + "")))))}
    var v3 uint32 = 2147483648
    v4 := []uint32{v3, v3, v3}
    v8 := GenS1{F0: (("foo bar" + ("\n\t" + "a")) + ""), F1: 42, F2: 18446744073709551615}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: -1, F2: 1}
        fmt.Println(v1)
        for fi11 := 0; fi11 < 11; fi11++ {
            v3 += (v3 ^ (v3 >> ((2) & 63)))
            if true {
                var v12 int = (fi9 * (-(fi11)))
                _ = v12
            } else {
                var v13 string = ("" + ("" + "\n\t"))
                _ = v13
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v3 -= v3
    v14 := GenS2{F0: 1, F1: -243, F2: ""}
    fmt.Println(v3, v1)
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(v3)
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int32 = 1
            if true {
                v1 -= (0.1 / (v1 * v1))
                if (true || false) {
                    v18 := GenS2{F0: 0, F1: (fi16 ^ (fi15 >> ((42) & 63))), F2: "foo bar"}
                    _ = v18
                }
                if false {
                    v3 += 65535
                    v19 := GenS2{F0: 665, F1: fi16, F2: ("hello" + (("\n\t" + "hello") + (("foo bar" + "a") + "foo bar")))}
                    _ = v19
                } else {
                    switch "hello" {
                    case "foo bar":
                        v20 := GenS1{F0: "hello", F1: 506, F2: 18446744073709551615}
                        _ = v20
                    case "\n\t":
                        switch fi16 {
                        case 32:
                            v17 -= 42
                        case 127:
                            _ = len(v2)
                        default:
                            var v21 float32 = (((float32((0.0 + -2.25)) / 1e10) * (0.1 + (float32(v1) / float32(3.14159)))) - 1.5)
                            _ = v21
                        }
                    case "":
                        v3 = v3
                    default:
                        v1 += (v1 - v1)
                    }
                    v22 := GenS3{F0: fi15, F1: 42}
                    var v23 string = "foo bar"
                    v24 := GenS3{F0: fi15, F1: 42}
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            fmt.Println(v3, v1, v3)
            _ = fi16
            _ = v17
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            if (-32768 < 32767) {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                v3 = (^(v3))
                var v26 int16 = -1
                var v27 float32 = 0.0
                _ = v26
                _ = v27
            } else {
                v3 = v3
            }
            var v28 string = ("foo bar" + ("foo bar" + "a"))
            _ = fi25
            _ = v28
        }
        _ = fi15
    }
    if false {
        v1 -= v1
    } else {
        var v29 bool = (false && true)
        fmt.Println(v29)
        _ = v29
    }
    v1 -= v1
    var v30 bool = true
    fmt.Println(v3, v1, v1)
    if (127 < -6) {
        v2 = append(v2, "hello")
        switch "\n\t" {
        case "a":
            var v31 int8 = -128
            _ = v31
        default:
            v1 -= v1
        }
    } else {
        v3 -= (v3 ^ (^(v3)))
    }
    v3 += (v3 << ((367) & 63))
    switch "foo bar" {
    case "foo bar":
        var v32 float32 = float32(((v1 + v1) / v1))
        _ = v32
    case "hello":
        var v33 uint8 = 128
        _ = v33
    default:
        for fi34 := 0; fi34 < 11; fi34++ {
            v35 := GenS3{F0: (fi34 >> ((255) & 63)), F1: 1}
            var v36 uint16 = 2
            if v30 {
                for fi37 := 0; fi37 < 14; fi37++ {
                    if v30 {
                        var v38 uint = 255
                        fmt.Println(fi34, v30, v30)
                        _ = v38
                    } else {
                        v30 = (true || true)
                        v39 := GenS3{F0: (fi34 - (fi34 & fi37)), F1: (^(v36))}
                        fmt.Println(v1, fi37)
                        _ = v39
                    }
                    var v40 uint = 976
                    _ = fi37
                    _ = v40
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                switch fi34 {
                case -466:
                    var v41 uint64 = 936
                    _ = v41
                }
            } else {
                _ = len(v2)
                var v42 int8 = 42
                _ = v42
            }
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v1, v1, v30)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
