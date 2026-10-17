package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 uint64 }
func genFn0(p0 int64, p1 uint32) uint64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(p0)
        if false {
            if (-128 > 0) {
                p1 += p1
            }
            p1 = p1
        }
        _ = fi5
    }
    if ((((0.1 / 3.14159) - (1e10 + -2.25)) != -2.25) || true) {
        p0 += (p0 >> ((843) & 63))
        var v6 int = -604
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 uint16 = 1
                p1 = 65535
                v10 := []int{}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    }
    var v11 uint32 = 743
    v12 := GenS1{F0: 9223372036854775807}
    v13 := GenS2{F0: (p0 + p0)}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 2
}
func genFn1(p0 uint64, p1 int) int32 {
    v14 := GenS3{F0: 0, F1: 42}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v14
    return -966
}

func main() {
    v1 := []uint32{}
    var v2 int32 = 0
    v3 := []string{(("a" + "hello") + ("hello" + "")), "hello", "", "foo bar"}
    v4 := []string{}
    var v15 bool = true
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint16 = 255
            switch ((("foo bar" + "hello") + "") + ("hello" + "\n\t")) {
            case "\n\t":
                v15 = false
            case "":
                v15 = false
                fmt.Println(v2, v15)
                var v19 string = (("hello" + "foo bar") + "")
                _ = v19
            case "hello":
                if len(v4) > 0 {
                    v4[0] = (("hello" + "") + "hello")
                }
                switch fi17 {
                case 703:
                    fmt.Println(v2)
                case -2147483648:
                    v15 = (false || false)
                case 2147483647:
                    v20 := GenS2{F0: -1}
                    _ = v20
                }
                v21 := GenS1{F0: fi17}
                v22 := GenS4{F0: 100, F1: 9223372036854775807}
                _ = v21
                _ = v22
            }
            var v23 uint64 = genFn0(-9223372036854775808, 392)
            _ = fi17
            _ = v18
            _ = v23
        }
        _ = fi16
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v2 -= -738
        _ = fi24
    }
    v25 := GenS1{F0: 127}
    v26 := GenS4{F0: -128, F1: -154}
    v27 := GenS2{F0: 1}
    v15 = (1 >= 0)
    var v28 uint8 = 228
    var v29 uint8 = v28
    v28 += v28
    v28 += 204
    for fi30 := 0; fi30 < 2; fi30++ {
        _ = len(v1)
        _ = fi30
    }
    if (42 > 1000) {
        switch ("\n\t" + "a") {
        case "hello":
            var v31 uint32 = 2147483647
            _ = v31
        default:
            var v32 string = ((("\n\t" + "a") + "") + ("" + "hello"))
            _ = v32
        }
    } else {
        fmt.Println(v29, v29, v15)
    }
    v33 := GenS5{F0: 18446744073709551615}
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v29, v28, v29)
        switch fi34 {
        case -302:
            _ = len(v4)
        case 9223372036854775807:
            for fi35 := 0; fi35 < 6; fi35++ {
                if len(v4) > 0 {
                    v4[0] = (((("" + "a") + "\n\t") + (("" + ("a" + ("hello" + "a"))) + ("" + "hello"))) + "hello")
                }
                _ = fi35
            }
        case -522:
            v15 = v15
        default:
            switch "hello" {
            case "foo bar":
                var v36 int64 = -172
                _ = v36
            case "\n\t":
                if (v15 && (2147483647 <= (v2 << ((1) & 63)))) {
                    fmt.Println(v15, v29, v28)
                } else {
                    v37 := GenS2{F0: -1}
                    _ = v37
                }
            case "":
                for fi38 := 0; fi38 < 6; fi38++ {
                    var v39 float32 = (-2.25 - -2.25)
                    v28 -= (v28 >> ((726) & 63))
                    if ((fi34 - (-(fi34))) != (fi38 ^ -2147483648)) {
                        var v40 uint64 = 101
                        fmt.Println(v40, v2, v29)
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "foo bar")
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v40
                    } else {
                        var v41 float64 = ((1e10 + float64((v39 * 1e10))) * (-2.25 - ((1.5 + float64(3.14159)) - (0.0 + -2.25))))
                        v39 += 0.0
                        _ = v41
                    }
                    v42 := GenS6{F0: -152, F1: 478}
                    v15 = true
                    _ = fi38
                    _ = v39
                    _ = v42
                }
            default:
                if v15 {
                    var v43 int64 = 9223372036854775807
                    fmt.Println(v28, v29, fi34)
                    v4 = append(v4, "foo bar")
                    _ = v43
                }
            }
        }
        for fi44 := 0; fi44 < 5; fi44++ {
            switch fi44 {
            case -197:
                var v45 uint8 = v28
                _ = v45
            }
            var v46 int16 = 0
            _ = fi44
            _ = v46
        }
        _ = fi34
    }
    var v47 int = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v47
}
