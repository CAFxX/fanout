package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 int16; F1 uint }
func genFn0(p0 string, p1 string) uint8 {
    fmt.Println(p0, p0)
    p1 = ""
    var v5 int8 = -1
    v6 := []int{}
    p0 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 129
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 1
    v3 := make(map[uint32]int)
    var v4 int16 = 832
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 bool = (fi7 >= 2147483647)
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint32 = 4294967295
            v9 = append(v9, fi11)
            if false {
                switch "a" {
                case "hello":
                    var v13 uint64 = 918
                    _ = v13
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        v14 := GenS1{F0: 1}
                        _ = v14
                    }
                case "a":
                    var v15 int = 2147483647
                    _ = v15
                default:
                    var v16 uint16 = 42
                    _ = v16
                }
            }
            if (int8(v4) < 127) {
                fmt.Println(fi10, v4, v2)
                if (18 >= v4) {
                    _ = len(v9)
                }
                v12 = 42
                fmt.Println(fi10, v4, fi10)
                v12 += v12
            } else {
                if ((true || false) || (int8(v2) == 67)) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                if true {
                    fmt.Println(v12, fi10)
                    if (false || (false && false)) {
                        v12 -= v12
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    v2 += (v2 + (v2 | (^(v2))))
                }
                switch ("hello" + ("hello" + "foo bar")) {
                case "hello":
                    v12 = 65535
                }
            }
            _ = fi11
            _ = v12
        }
        v4 += (^(v4))
        _ = fi10
    }
    var v17 int = 409
    if (2147483647 < 94) {
        v18 := v3[420]
        if false {
            if (true && (!(1 > 42))) {
                switch (-(v18)) {
                case -1:
                    var v19 float32 = float32(3.14159)
                    _ = v19
                }
                fmt.Println(v17, v4, v2)
            } else {
                var v20 int64 = int64(v4)
                var v21 bool = (-2147483648 <= 128)
                v21 = true
                _ = v20
                _ = v21
            }
            switch (v17 + (v17 & (v17 / (((v17 / ((1) | 1))) | 1)))) {
            case -2147483648:
                var v22 int64 = 1
                _ = v22
            case 157:
                if (!(127 <= -2)) {
                    v23 := GenS1{F0: 2}
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v3[42] = v17
                        _ = fi24
                    }
                    v2 += v2
                    _ = v23
                }
            case 4294967295:
                v17 += v17
            default:
                var v25 int64 = 191
                _ = v25
            }
            switch (v17 ^ 4294967295) {
            case 32767:
                if (!(false || (false || (!false)))) {
                    v4 -= 42
                    if ((0.0 / (1.5 + 0.0)) < -2.25) {
                        v18 -= (v18 / ((v18) | 1))
                    }
                } else {
                    switch "a" {
                    case "hello":
                        switch "hello" {
                        case "a":
                            if true {
                                v18 = (v17 / ((v17) | 1))
                            }
                        case "foo bar":
                            if (128 == 42) {
                                v18 = -279
                            } else {
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v4, v17, v2)
                                case "hello":
                                    v4 -= 0
                                case "":
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        fmt.Println(v4)
                                        if false {
                                            v27 := GenS1{F0: genFn0("foo bar", ("foo bar" + "foo bar"))}
                                            v28 := GenS1{F0: genFn0("foo bar", (((((("\n\t" + ("foo bar" + "hello")) + "a") + ("\n\t" + "foo bar")) + "") + "hello") + ("foo bar" + "")))}
                                            v1[(v17 | fi26)] = ("a" + ("a" + ""))
                                            _ = v27
                                            _ = v28
                                        }
                                        _ = len(v3)
                                        _ = fi26
                                    }
                                }
                                switch (^(v18)) {
                                case -2147483648:
                                    v4 -= int16(v2)
                                default:
                                    v29 := GenS2{F0: 0, F1: v4, F2: 792}
                                    _ = v29
                                }
                            }
                        case "\n\t":
                            v17 = (v18 - (v17 ^ (v17 - v18)))
                        }
                    default:
                        var v30 int16 = v4
                        _ = v30
                    }
                }
            case -156:
                for fi31 := 0; fi31 < 2; fi31++ {
                    v9 = append(v9, int(v4))
                    fmt.Println(v4, v18, v2)
                    v1[(v17 << (((v2 ^ v2)) & 63))] = "\n\t"
                    _ = fi31
                }
            case 4294967295:
                var v32 int8 = 42
                _ = v32
            default:
                v33 := GenS3{F0: 0, F1: v4, F2: 9223372036854775807}
                _ = v33
            }
        } else {
            v2 = v2
        }
        var v34 string = "\n\t"
        _ = v18
        _ = v34
    } else {
        if false {
            v35 := GenS3{F0: 0, F1: -32768, F2: v17}
            fmt.Println(v4, v17)
            var v36 bool = false
            _ = v35
            _ = v36
        } else {
            var v37 int16 = (v4 ^ 905)
            switch (v17 | (v17 % (((v17 / (((v17 ^ (v17 ^ (v17 | v17)))) | 1))) | 1))) {
            case 9223372036854775807:
                if (((v4 | v4) > -1) || false) {
                    v17 = v17
                }
            case 32768:
                v38 := GenS4{F0: 1, F1: v2}
                _ = v38
            case 459:
                v37 += -932
            }
            switch 342 {
            case -623:
                if ("foo bar" < "hello") {
                    v37 -= v4
                } else {
                    var v39 int16 = int16(v2)
                    _ = v39
                }
            case 1:
                var v40 int32 = 42
                _ = v40
            }
            _ = v37
        }
    }
    switch ((("\n\t" + ("a" + "a")) + "\n\t") + (("" + (("a" + "") + "a")) + "")) {
    case "hello":
        if false {
            v41 := GenS3{F0: 0, F1: (v4 & (^(v4))), F2: v17}
            _ = v41
        } else {
            v17 -= (v17 + (v17 - 4294967295))
            var v42 bool = (-2147483648 < -2147483648)
            _ = v42
        }
    default:
        if false {
            var v43 int64 = -9223372036854775808
            switch "hello" {
            case "hello":
                v43 = v43
            case "foo bar":
                for fi44 := 0; fi44 < 2; fi44++ {
                    var v45 int32 = 1
                    v46 := GenS3{F0: (v45 * 0), F1: 64, F2: (-(fi44))}
                    fmt.Println(v17, v17, v4)
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            for fi47 := 0; fi47 < 5; fi47++ {
                v2 = v2
                if false {
                    var v48 int32 = -1
                    _ = v48
                } else {
                    var v49 uint32 = 252
                    v9 = append(v9, (v17 % ((v17) | 1)))
                    _ = v49
                }
                _ = fi47
            }
            _ = v43
        }
    }
    fmt.Println(v4, v4, v4)
    switch "hello" {
    case "":
        var v50 int8 = 42
        _ = v50
    case "hello":
        v2 = (v2 + v2)
    }
    var v51 int16 = (v4 ^ (v4 ^ v4))
    var v52 float64 = float64(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v9
}
