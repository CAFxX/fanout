package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 int16 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }
func genFn0(p0 int8, p1 uint8, p2 uint) uint16 {
    v6 := GenS1{F0: p0}
    fmt.Println(p2, p1, p2)
    switch "foo bar" {
    case "":
        fmt.Println(p2, p1)
    case "\n\t":
        p0 = (^(p0))
    case "hello":
        var v7 uint32 = 65535
        _ = v7
    }
    var v8 uint64 = 42
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    _ = v9
    return 854
}
func genFn1(p0 uint16) uint64 {
    switch 42 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint8 = 232
            _ = fi10
            _ = v11
        }
    case -800:
        switch "" {
        case "foo bar":
            if (("foo bar" + (("hello" + "hello") + ("a" + (("foo bar" + "foo bar") + "")))) <= "a") {
                v12 := GenS1{F0: -66}
                _ = v12
            } else {
                var v13 string = ""
                v14 := GenS1{F0: -128}
                _ = v13
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS2{F0: "a", F1: 0, F2: -1}
                switch (fi15 / (((fi15 >> ((7) & 63))) | 1)) {
                case 42:
                    if true {
                        v17 := GenS1{F0: 1}
                        fmt.Println(fi15)
                        _ = v17
                    } else {
                        v18 := []int{(fi15 & (fi15 * fi15)), (fi15 ^ (fi15 * fi15)), -1, fi15}
                        _ = v18
                    }
                case -9223372036854775808:
                    for fi19 := 0; fi19 < 8; fi19++ {
                        p0 += 918
                        _ = fi19
                    }
                case 128:
                    p0 = p0
                }
                _ = fi15
                _ = v16
            }
        }
    case 32768:
        p0 += p0
    }
    p0 += genFn0(-1, 0, 0)
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 1
    var v3 uint = 4294967295
    var v4 uint = v3
    var v5 int8 = -62
    var v20 uint = (v3 % ((4294967295) | 1))
    v21 := GenS3{F0: -32768, F1: 1}
    switch 127 {
    case -9223372036854775808:
        v22 := GenS1{F0: v5}
        _ = v22
    default:
        switch 128 {
        case 32768:
            var v23 uint32 = 65535
            _ = v23
        case 42:
            var v24 float64 = 0.1
            _ = v24
        case -534:
            var v25 float64 = 3.14159
            _ = v25
        default:
            var v26 bool = (0.0 < (1.5 / ((float32((0.1 * 1e10)) / (float32(1.5) / 0.1)) + (1e10 * float32(0.1)))))
            _ = v26
        }
    }
    delete(v1, 4294967295)
    var v27 int16 = -32768
    if (false && (-2.25 > -2.25)) {
        if (v3 > v4) {
            for fi28 := 0; fi28 < 12; fi28++ {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v29 := []int{(-(fi28)), (fi28 << (((v4 + 0)) & 63)), (fi28 * (-(fi28))), 9223372036854775807}
                    _ = v29
                default:
                    fmt.Println(v3, fi28, v27)
                }
                v5 += (v5 % (((v5 ^ v5)) | 1))
                var v30 float32 = float32(0.0)
                _ = fi28
                _ = v30
            }
            v27 += v27
            if (false && (float64(0.0) != 0.1)) {
                switch "a" {
                case "hello":
                    var v31 float64 = (float64((3.14159 / 0.1)) - 3.14159)
                    _ = v31
                case "foo bar":
                    v32 := GenS4{F0: 0}
                    _ = v32
                case "":
                    v33 := GenS3{F0: (v27 | (v27 >> (((v4 ^ (v4 >> (((v3 >> ((v3) & 63))) & 63)))) & 63))), F1: 0}
                    _ = v33
                default:
                    v5 = v5
                }
            } else {
                for fi34 := 0; fi34 < 9; fi34++ {
                    var v35 int32 = -1
                    _ = fi34
                    _ = v35
                }
                v4 = v20
            }
        } else {
            fmt.Println(v4)
            var v36 string = (("hello" + "a") + ("\n\t" + ("" + "foo bar")))
            _ = v36
        }
        if false {
            var v37 float32 = -2.25
            _ = v37
        }
    }
    switch ("\n\t" + "") {
    case "\n\t":
        v38 := []int{128}
        _ = v38
    case "hello":
        v39 := GenS1{F0: v5}
        _ = v39
    default:
        for fi40 := 0; fi40 < 7; fi40++ {
            if (!(!(1 <= 2147483648))) {
                v4 = (v4 >> (((v20 + (v3 * 2))) & 63))
                var v41 float64 = 3.14159
                fmt.Println(v27)
                _ = v41
            } else {
                if (-453 <= -1) {
                    if (v5 < v5) {
                        switch fi40 {
                        case 32768:
                            if (("foo bar" == (("" + "hello") + "\n\t")) && false) {
                                var v42 int = fi40
                                switch "hello" {
                                case "hello":
                                    var v43 int32 = 42
                                    _ = v43
                                case "":
                                    var v44 uint16 = 65535
                                    _ = v44
                                }
                                _ = v42
                            }
                        default:
                            var v45 int = (fi40 << (((v20 & v3)) & 63))
                            _ = v45
                        }
                    }
                }
            }
            if true {
                for fi46 := 0; fi46 < 1; fi46++ {
                    v47 := []int{(fi40 | (fi40 | (fi46 << ((65535) & 63)))), 0}
                    fmt.Println(fi46, v20)
                    v20 = v3
                    _ = fi46
                    _ = v47
                }
                v3 = v20
                for fi48 := 0; fi48 < 8; fi48++ {
                    v4 = (v3 << (((^(v20))) & 63))
                    _ = fi48
                }
            } else {
                var v49 uint64 = v2
                _ = v49
            }
            _ = fi40
        }
    }
    v50 := GenS5{F0: genFn0((v5 - (v5 << (((v4 << (((v4 % ((v20) | 1))) & 63))) & 63))), 183, v20)}
    var v51 uint64 = v2
    v52 := GenS4{F0: v20}
    v53 := GenS6{F0: "a"}
    for fi54 := 0; fi54 < 3; fi54++ {
        for fi55 := 0; fi55 < 16; fi55++ {
            fmt.Println(fi54, v20)
            _ = fi55
        }
        var v56 float32 = -2.25
        for fi57 := 0; fi57 < 19; fi57++ {
            v58 := GenS7{F0: "foo bar"}
            fmt.Println(v20, fi57, v2)
            _ = fi57
            _ = v58
        }
        _ = fi54
        _ = v56
    }
    v59 := GenS3{F0: (v27 ^ 42), F1: -681}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v59
}
