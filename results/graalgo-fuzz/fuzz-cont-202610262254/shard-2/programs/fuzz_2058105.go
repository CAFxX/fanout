package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 string; F1 uint8; F2 int32 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int64; F1 uint64; F2 string }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int32 }
func genFn0(p0 int64, p1 string, p2 bool) int32 {
    v5 := GenS1{F0: (("foo bar" + p1) + (((p1 + p1) + p1) + ""))}
    v6 := GenS1{F0: p1}
    v7 := GenS2{F0: 932}
    for fi8 := 0; fi8 < 12; fi8++ {
        if p2 {
            p1 = p1
        } else {
            for fi9 := 0; fi9 < 15; fi9++ {
                switch ((("\n\t" + (((p1 + p1) + ("" + (p1 + ""))) + p1)) + p1) + p1) {
                case "":
                    switch ((p1 + p1) + ("foo bar" + (p1 + p1))) {
                    case "\n\t":
                        var v10 uint = 2
                        _ = v10
                    }
                case "foo bar":
                    v11 := GenS3{F0: (p0 << ((255) & 63)), F1: 127}
                    _ = v11
                case "\n\t":
                    var v12 uint64 = 4294967295
                    _ = v12
                default:
                    if (p1 < (p1 + "a")) {
                        v13 := GenS4{F0: "\n\t", F1: 128, F2: -1}
                        _ = v13
                    }
                }
                if p2 {
                    var v14 uint32 = 74
                    _ = v14
                }
                _ = fi9
            }
            if ((p1 != p1) || (p2 && p2)) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi8, p2, p2)
                case "":
                    v15 := []int{(fi8 / (((fi8 % ((9223372036854775807) | 1))) | 1)), fi8, 9223372036854775807}
                    _ = v15
                default:
                    v16 := []int{fi8, fi8}
                    _ = v16
                }
            }
        }
        p1 = ("a" + p1)
        p2 = (false || false)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() int16 {
    var v17 string = ""
    v18 := []int{9223372036854775807}
    for fi19 := 0; fi19 < 3; fi19++ {
        v17 = "a"
        var v20 int = fi19
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v17 = "foo bar"
        _ = fi21
    }
    _ = v17
    _ = v18
    return 32767
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{2147483648, 65535, 100, 1}
    v3 := make(map[int]bool)
    var v4 uint16 = 0
    switch "\n\t" {
    case "":
        if (9223372036854775807 <= 9223372036854775807) {
            v2 = append(v2, 4294967295)
            switch (("foo bar" + ("foo bar" + "a")) + "hello") {
            case "\n\t":
                switch -2147483648 {
                case -35:
                    v22 := GenS1{F0: "foo bar"}
                    _ = v22
                case 4294967295:
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v1 = v1
                        var v24 int8 = 0
                        _ = fi23
                        _ = v24
                    }
                case 2147483647:
                    v4 += (v4 + v4)
                default:
                    v1 = (v1 << ((4294967295) & 63))
                }
            }
        } else {
            var v25 uint32 = 2147483648
            _ = v25
        }
    case "hello":
        v26 := GenS5{F0: (("a" + "hello") + ("\n\t" + "hello")), F1: ((("" + "foo bar") + "a") + ("a" + "foo bar"))}
        _ = v26
    case "a":
        switch -2147483648 {
        case 1:
            if false {
                var v27 int64 = 574
                _ = v27
            } else {
                v4 = (v4 | 507)
            }
        case -813:
            delete(v3, 4294967295)
        }
    }
    v28 := GenS6{F0: -1, F1: v1, F2: ""}
    switch 4294967295 {
    case -748:
        fmt.Println(v4)
    case 42:
        if (2 != 2) {
            if len(v2) > 0 {
                v2[0] = 42
            }
            var v29 bool = true
            switch 1 {
            case -1:
                for fi30 := 0; fi30 < 18; fi30++ {
                    v29 = (((65535 >= 65535) || (1 < 2147483648)) || (((1.5 >= (((float64(float32(3.14159)) + (0.1 * 3.14159)) / 1.5) - 1.5)) && true) || false))
                    _ = fi30
                }
            case -269:
                var v31 int32 = -2147483648
                _ = v31
            case 128:
                var v32 bool = (v1 < v1)
                _ = v32
            default:
                switch "" {
                case "":
                    v1 = v1
                case "\n\t":
                    var v33 uint64 = (v1 >> ((1) & 63))
                    _ = v33
                case "foo bar":
                    fmt.Println(v29, v29)
                default:
                    v4 += 65535
                }
            }
            _ = v29
        }
    case 128:
        if (-2147483648 >= 2147483647) {
            var v34 int8 = -128
            _ = v34
        } else {
            fmt.Println(v4, v4)
        }
    default:
        switch "foo bar" {
        case "":
            v1 -= (v1 + (v1 % ((42) | 1)))
        case "foo bar":
            switch "a" {
            case "hello":
                for fi35 := 0; fi35 < 3; fi35++ {
                    v1 += 18446744073709551615
                    v1 = v1
                    _ = fi35
                }
            case "a":
                var v36 uint8 = 208
                _ = v36
            default:
                v1 += v1
            }
        case "hello":
            delete(v3, -249)
        default:
            var v37 int = 42
            _ = v37
        }
    }
    v2 = append(v2, 42)
    v38 := GenS7{F0: 2147483647, F1: 127}
    v39 := GenS2{F0: int32(v4)}
    for fi40 := 0; fi40 < 17; fi40++ {
        v41 := GenS3{F0: 1, F1: 659}
        v42 := GenS2{F0: genFn0(-827, "\n\t", true)}
        v4 += 184
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int = 128
    var v44 int32 = 2147483647
    v43 = v43
    for fi45 := 0; fi45 < 11; fi45++ {
        var v46 string = (("hello" + ((((("hello" + "") + (("\n\t" + (("\n\t" + ((("\n\t" + "hello") + "hello") + "foo bar")) + "foo bar")) + ("hello" + "foo bar"))) + "a") + "") + ("" + ("a" + "foo bar")))) + "foo bar")
        if len(v2) > 0 {
            v2[0] = 358
        }
        _ = fi45
        _ = v46
    }
    v43 -= (-(v43))
    v4 -= 0
    v1 = (^(v1))
    for fi47 := 0; fi47 < 13; fi47++ {
        var v48 bool = true
        _ = fi47
        _ = v48
    }
    if false {
        v4 = (v4 + v4)
        v1 += (^(v1))
        v49 := GenS8{F0: int32(v43)}
        _ = v49
    } else {
        for fi50 := 0; fi50 < 6; fi50++ {
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 uint = 65535
                _ = fi51
                _ = v52
            }
            if ((-(fi50)) == -572) {
                var v53 bool = true
                v44 = (v44 / (((v44 - (-(v44)))) | 1))
                _ = v53
            }
            _ = fi50
        }
        v54 := GenS3{F0: 9223372036854775807, F1: (v43 << ((65535) & 63))}
        fmt.Println(v4, v43)
        _ = v54
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        var v55 int8 = 0
        _ = v55
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
}
