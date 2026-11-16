package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS5 struct { F0 bool; F1 int8; F2 bool }
func genFn0(p0 uint) uint32 {
    var v5 int8 = -128
    v5 = -1
    p0 = (p0 - (p0 | 18446744073709551615))
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 int = 128
    var v2 float64 = ((1.5 - 0.0) / ((1e10 - 1.5) - float64(0.0)))
    var v3 uint64 = 42
    v4 := make(map[string]bool)
    switch (v1 | v1) {
    case 9223372036854775807:
        v6 := GenS1{F0: 18446744073709551615, F1: 508, F2: (v1 >> ((775) & 63))}
        _ = v6
    default:
        switch "foo bar" {
        case "":
            var v7 int = (^(v1))
            _ = v7
        case "foo bar":
            v8 := GenS1{F0: 1, F1: 18446744073709551615, F2: v1}
            _ = v8
        default:
            delete(v4, "foo bar")
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 += (fi9 | v1)
        if (false || (!true)) {
            v10 := v4[""]
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int = fi11
                v1 -= (fi11 * (v12 % (((^(v12))) | 1)))
                switch (v12 >> ((255) & 63)) {
                case 2147483647:
                    var v13 string = "hello"
                    _ = v13
                case -2147483648:
                    v10 = true
                    var v14 int64 = 9223372036854775807
                    v2 -= (float64(((0.1 * (0.0 * -2.25)) - ((0.0 - -2.25) - 1.5))) + float64(((0.1 * ((((0.1 + ((-2.25 - float32((v2 / 1.5))) / 3.14159)) / float32(v2)) / 1.5) / 1e10)) + 1e10)))
                    _ = v14
                case 42:
                    var v15 uint32 = 287
                    _ = v15
                default:
                    if v10 {
                        fmt.Println(v2, v12)
                        if (!true) {
                            v12 += -460
                        } else {
                            v12 -= (fi11 << ((18446744073709551615) & 63))
                            var v16 int32 = -1
                            _ = v16
                        }
                    }
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            v17 := GenS1{F0: 1, F1: 444, F2: v1}
            v4[""] = (!(!false))
            var v18 uint32 = 432
            _ = v17
            _ = v18
        }
        _ = fi9
    }
    v3 = v3
    if (!true) {
        switch "a" {
        case "\n\t":
            var v19 uint32 = 65535
            _ = v19
        case "a":
            var v20 uint64 = v3
            _ = v20
        }
    } else {
        var v21 float64 = -2.25
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint8 = 148
            v23 = 165
            _ = fi22
            _ = v23
        }
        v2 = 3.14159
        _ = v21
    }
    v24 := GenS1{F0: 42, F1: 1, F2: (v1 * (v1 * v1))}
    var v25 uint8 = 124
    v1 -= v1
    var v26 int32 = 999
    v27 := GenS2{F0: 255}
    v28 := GenS3{F0: (false && true), F1: 32767, F2: (v25 / ((255) | 1))}
    v29 := GenS1{F0: 255, F1: 2, F2: v1}
    v1 += (^(v1))
    var v30 int64 = 0
    if false {
        if (-32768 < 0) {
            v25 = (v25 + 2)
            fmt.Println(v26)
        } else {
            var v31 int16 = 0
            var v32 int64 = (v30 - (v30 & v30))
            _ = v31
            _ = v32
        }
        var v33 int32 = (v26 + v26)
        var v34 int16 = 32767
        _ = v33
        _ = v34
    } else {
        v1 -= (v1 ^ 1)
        v1 += v1
    }
    v35 := []int{-1, (v1 ^ (v1 + v1)), (v1 & 0)}
    var v36 uint16 = 392
    if (945 > 989) {
        v25 = v25
        for fi37 := 0; fi37 < 18; fi37++ {
            v2 = v2
            var v38 string = ("foo bar" + (("a" + (("\n\t" + "") + "\n\t")) + ("" + "\n\t")))
            _ = fi37
            _ = v38
        }
    } else {
        switch ("hello" + "foo bar") {
        case "foo bar":
            if true {
                fmt.Println(v3, v3)
                switch (v1 | 4294967295) {
                case -122:
                    delete(v4, ("foo bar" + "\n\t"))
                case 9223372036854775807:
                    switch "hello" {
                    case "":
                        fmt.Println(v2, v1)
                    case "hello":
                        v39 := GenS4{F0: (v25 | 2), F1: genFn0(4294967295), F2: (v26 * v26)}
                        _ = v39
                    case "\n\t":
                        switch v1 {
                        case 127:
                            v40 := GenS5{F0: true, F1: -128, F2: ((^(v26)) < (v26 * 155))}
                            _ = v40
                        case 726:
                            fmt.Println(v2, v30, v3)
                        }
                    }
                default:
                    var v41 uint8 = (^(v25))
                    _ = v41
                }
                switch 32768 {
                case -15:
                    v35 = append(v35, (v1 - v1))
                }
            }
        default:
            for fi42 := 0; fi42 < 6; fi42++ {
                v2 = -2.25
                var v43 float32 = -2.25
                var v44 bool = (fi42 > (v1 << ((65535) & 63)))
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        v3 += (v3 * (v3 * (v3 >> ((478) & 63))))
    }
    var v45 uint16 = (v36 >> ((94) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v45
}
