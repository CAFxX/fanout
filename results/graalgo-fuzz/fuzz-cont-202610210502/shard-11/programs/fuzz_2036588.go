package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 string; F2 uint8 }
func genFn0(p0 int64, p1 string, p2 bool) string {
    fmt.Println(p0, p2)
    p0 = p0
    v7 := GenS1{F0: 128, F1: -9223372036854775808}
    v8 := GenS2{F0: (p2 && p2), F1: p1, F2: 2}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 + p1)
}
func genFn1(p0 int8, p1 int16) int8 {
    var v9 int32 = 2147483647
    if true {
        var v10 float64 = ((0.1 * float64(1.5)) / float64(((((1.5 + -2.25) / float32(3.14159)) / 1.0) - 0.0)))
        _ = v10
    } else {
        v11 := GenS2{F0: true, F1: "hello", F2: 18446744073709551615}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 100
}

func main() {
    var v1 int64 = 42
    v2 := []string{"\n\t", "\n\t"}
    var v3 int64 = (v1 + v1)
    var v4 int8 = 42
    var v5 bool = false
    var v6 float64 = 0.1
    v5 = v5
    v4 -= 100
    v12 := GenS3{F0: 1}
    if (1 > 42) {
        v13 := GenS3{F0: 2}
        var v14 float32 = (float32(v6) - 0.0)
        _ = v13
        _ = v14
    }
    if ((v6 - (v6 * v6)) != (v6 - ((float64(1.5) / (((v6 - 0.1) / v6) * ((0.1 + ((1e10 + (v6 + v6)) * v6)) + v6))) - (v6 * (3.14159 / float64(-2.25)))))) {
        v15 := GenS1{F0: 213, F1: 128}
        v1 = (-(v1))
        _ = v15
    } else {
        v16 := GenS2{F0: v5, F1: (((("a" + "\n\t") + "foo bar") + ("hello" + (((("a" + "\n\t") + ("hello" + (("a" + "\n\t") + ""))) + ("\n\t" + "\n\t")) + ("hello" + "\n\t")))) + "foo bar"), F2: 207}
        var v17 string = "a"
        var v18 int8 = v4
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println(v6, v4)
    for fi19 := 0; fi19 < 17; fi19++ {
        if ((fi19 << ((4294967295) & 63)) != fi19) {
            if v5 {
                _ = len(v2)
                v20 := GenS4{F0: "hello"}
                for fi21 := 0; fi21 < 15; fi21++ {
                    fmt.Println(v3, v1)
                    _ = fi21
                }
                _ = v20
            } else {
                var v22 uint16 = 214
                _ = v22
            }
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 string = "a"
                v24 = v24
                switch 1 {
                case 545:
                    _ = len(v2)
                case 32768:
                    switch ((((v24 + v24) + "") + "hello") + (v24 + "")) {
                    case "hello":
                        fmt.Println(v4, v1, fi19)
                    case "\n\t":
                        if v5 {
                            var v25 string = (("foo bar" + (v24 + v24)) + "")
                            switch (fi23 / (((-(fi19))) | 1)) {
                            case 4294967295:
                                var v26 int = 42
                                _ = v26
                            case -9223372036854775808:
                                var v27 string = v25
                                _ = v27
                            case 32767:
                                if (v25 > "a") {
                                    if (!((v6 * 1e10) == (1e10 * float64((0.1 + (1.5 / -2.25)))))) {
                                        var v28 int64 = v3
                                        _ = v28
                                    }
                                }
                            }
                            v4 = v4
                            _ = v25
                        } else {
                            if (42 <= 0) {
                                var v29 uint = 2
                                _ = v29
                            }
                        }
                        var v30 int16 = 42
                        _ = v30
                    case "a":
                        v1 += (v1 << ((4294967295) & 63))
                    }
                default:
                    _ = len(v2)
                }
                _ = fi23
                _ = v24
            }
            v31 := GenS5{F0: 128, F1: 288}
            _ = v31
        }
        _ = fi19
    }
    var v32 int = 860
    if (0 <= 32767) {
        var v33 bool = v5
        for fi34 := 0; fi34 < 9; fi34++ {
            for fi35 := 0; fi35 < 13; fi35++ {
                if (209 > 2147483647) {
                    if v5 {
                        v36 := GenS5{F0: fi35, F1: 2}
                        _ = len(v2)
                        _ = v36
                    } else {
                        v37 := GenS6{F0: 1}
                        switch (v32 - v32) {
                        case -222:
                            v38 := GenS7{F0: 0}
                            _ = v38
                        case 924:
                            fmt.Println(v5)
                        case -2147483648:
                            if (v33 || (v6 >= v6)) {
                                if (!(-1 >= -848)) {
                                    var v39 uint = 42
                                    _ = v39
                                } else {
                                    v3 = v1
                                    v32 = (v32 * -550)
                                }
                            } else {
                                v1 = -1
                                if len(v2) > 0 {
                                    v2[0] = ("hello" + "a")
                                }
                                v2 = append(v2, ("\n\t" + "\n\t"))
                            }
                        }
                        _ = v37
                    }
                    if false {
                        _ = len(v2)
                    }
                    v40 := GenS8{F0: 42, F1: "hello", F2: 237}
                    _ = v40
                }
                var v41 int32 = -2147483648
                _ = fi35
                _ = v41
            }
            _ = fi34
        }
        _ = v33
    } else {
        v32 += -475
        switch "hello" {
        case "foo bar":
            if true {
                _ = len(v2)
                for fi42 := 0; fi42 < 18; fi42++ {
                    v4 = (-(v4))
                    _ = fi42
                }
                var v43 int8 = (v4 << ((255) & 63))
                _ = v43
            } else {
                var v44 int32 = 0
                _ = v44
            }
        case "":
            switch (^(v32)) {
            case 900:
                v1 += -448
            }
        case "hello":
            v6 += v6
        }
    }
    var v45 uint8 = 0
    var v46 uint16 = 0
    fmt.Println(v46)
    fmt.Println(v32, v6)
    var v47 int64 = (v3 & (v3 << ((11) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
