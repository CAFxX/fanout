package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int }
type GenS4 struct { F0 int; F1 uint; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
func genFn0(p0 int64, p1 int64, p2 int16) uint {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int32 = -596
        v7 := []int{fi5, 32767}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: int16(p1), F1: 18446744073709551615, F2: ("a" + ("\n\t" + "hello"))}
    switch 42 {
    case 42:
        var v9 int16 = p2
        _ = v9
    case -692:
        if ("a" != "hello") {
            var v10 uint64 = 18446744073709551615
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(v10, v10)
                fmt.Println(fi11)
                _ = fi11
            }
            _ = v10
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint8 = 0
                p0 += p1
                p2 -= 392
                _ = fi12
                _ = v13
            }
            v14 := GenS2{F0: 2}
            v15 := GenS1{F0: -1, F1: 1, F2: "a"}
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}
func genFn1(p0 int16, p1 bool) int16 {
    switch "foo bar" {
    case "foo bar":
        p1 = ((1 >= 1) || (1 != -91))
    }
    fmt.Println(p1, p0)
    p1 = p1
    v16 := GenS3{F0: 42, F1: -386, F2: 32768}
    var v17 int16 = (p0 - p0)
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return v17
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -201
    var v3 uint16 = 0
    var v4 uint64 = 123
    var v18 float64 = 1e10
    v19 := []int{v2}
    if true {
        for fi20 := 0; fi20 < 17; fi20++ {
            switch (v2 * fi20) {
            case -9223372036854775808:
                v21 := GenS4{F0: (^(v2)), F1: 65535, F2: false}
                _ = v21
            default:
                if (!true) {
                    v18 += v18
                }
            }
            var v22 int16 = 520
            _ = fi20
            _ = v22
        }
        switch 2147483647 {
        case -986:
            var v23 bool = false
            _ = v23
        case -9223372036854775808:
            v4 = v4
        }
        switch (v2 % (((v2 + v2)) | 1)) {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 uint8 = 127
                _ = fi24
                _ = v25
            }
        default:
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
    } else {
        fmt.Println(v3, v3, v2)
        v4 += (v4 + (v4 + v4))
    }
    v26 := v1[v2]
    switch int(v4) {
    case 4294967295:
        v18 += (float64(-2.25) + 3.14159)
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        if false {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 float32 = 0.1
                var v30 int32 = -992
                if true {
                    var v31 int = (^(fi28))
                    var v32 int8 = 100
                    if false {
                        if (false || (!true)) {
                            v33 := GenS1{F0: 0, F1: 399, F2: ""}
                            if true {
                                fmt.Println(v30, v31, v4)
                            }
                            var v34 string = "foo bar"
                            _ = v33
                            _ = v34
                        }
                        var v35 int8 = (v32 + (^(v32)))
                        var v36 uint = 42
                        _ = v35
                        _ = v36
                    } else {
                        v4 = (^(v4))
                        fmt.Println(fi28, v30)
                        _ = len(v19)
                    }
                    _ = v31
                    _ = v32
                } else {
                    var v37 int64 = -766
                    v18 = (1e10 * 0.0)
                    _ = v37
                }
                if true {
                    var v38 int8 = 0
                    _ = v38
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = fi27
    }
    var v39 int16 = 209
    var v40 float32 = 1.5
    if true {
        var v41 int8 = 1
        if (!(v39 != (v39 & v39))) {
            if true {
                v2 = v26
                for fi42 := 0; fi42 < 18; fi42++ {
                    v43 := GenS1{F0: 32767, F1: (v4 + v4), F2: "hello"}
                    var v44 float32 = v40
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                v1[(v26 | v2)] = (-(v26))
            } else {
                if true {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        var v46 uint32 = 1
                        fmt.Println(fi45)
                        _ = fi45
                        _ = v46
                    }
                    switch "hello" {
                    case "a":
                        v47 := GenS5{F0: 570}
                        _ = v47
                    case "\n\t":
                        v2 += (v26 + (v26 % ((v26) | 1)))
                    case "foo bar":
                        var v48 uint16 = 2
                        _ = v48
                    default:
                        v41 -= (v41 * (v41 % ((v41) | 1)))
                    }
                }
                for fi49 := 0; fi49 < 14; fi49++ {
                    v50 := GenS6{F0: "a"}
                    for fi51 := 0; fi51 < 15; fi51++ {
                        _ = len(v1)
                        if (128 != 0) {
                            v18 -= 0.1
                            v41 -= (v41 % ((-68) | 1))
                        }
                        v18 -= v18
                        _ = fi51
                    }
                    _ = fi49
                    _ = v50
                }
                v41 -= (v41 & v41)
            }
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 uint8 = 127
                fmt.Println(v39, v18)
                v18 = v18
                _ = fi52
                _ = v53
            }
            if true {
                switch ("a" + "hello") {
                case "":
                    switch ("" + "foo bar") {
                    case "hello":
                        var v54 float64 = (v18 + ((v18 - 1.5) * ((1.5 * ((1.5 / v18) + 3.14159)) - v18)))
                        _ = v54
                    case "\n\t":
                        var v55 int32 = -1
                        _ = v55
                    case "a":
                        var v56 int64 = 271
                        _ = v56
                    }
                case "a":
                    for fi57 := 0; fi57 < 4; fi57++ {
                        var v58 uint16 = 523
                        v59 := GenS7{F0: 42}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                }
                switch "" {
                case "foo bar":
                    delete(v1, 459)
                case "":
                    v18 = ((v18 + -2.25) / 1.0)
                }
            }
        }
        _ = v41
    } else {
        v60 := GenS8{F0: (v2 & (v2 * (v2 << ((255) & 63))))}
        switch "hello" {
        case "foo bar":
            for fi61 := 0; fi61 < 7; fi61++ {
                v2 = (v2 | 9223372036854775807)
                _ = fi61
            }
        case "hello":
            if (v4 >= v4) {
                v4 = v4
            }
        }
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
