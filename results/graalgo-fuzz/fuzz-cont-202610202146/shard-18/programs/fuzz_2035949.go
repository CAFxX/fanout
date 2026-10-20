package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int32 = 0
    var v4 int16 = 1000
    var v5 uint8 = 129
    v5 = (v5 << ((1) & 63))
    var v6 float64 = 0.1
    if (255 >= 32767) {
        var v7 uint64 = 18446744073709551615
        delete(v1, "foo bar")
        v8 := v2[53]
        _ = v7
        _ = v8
    } else {
        if (!(65535 <= 577)) {
            if (42 > (v5 << ((880) & 63))) {
                v9 := GenS1{F0: -1, F1: 220}
                _ = v9
            }
        } else {
            v10 := []int{-690}
            var v11 float32 = float32(((float64(0.0) - 1e10) * float64((float32(1.5) * 0.1))))
            _ = v10
            _ = v11
        }
        if (!false) {
            var v12 int = 507
            _ = v12
        }
    }
    var v13 int16 = v4
    fmt.Println(v3, v13)
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            v3 -= v3
            _ = fi15
        }
        if ((!false) || true) {
            fmt.Println(v6)
            for fi16 := 0; fi16 < 3; fi16++ {
                v5 -= 2
                v3 = v3
                v17 := GenS1{F0: -9223372036854775808, F1: 4294967295}
                if (0 != 759) {
                    fmt.Println(v4)
                    v1[""] = true
                } else {
                    v18 := GenS2{F0: 973, F1: 818, F2: 4294967295}
                    v5 = v5
                    _ = v18
                }
                switch ("hello" + (((("\n\t" + (((((("hello" + ("a" + "")) + "a") + ("hello" + "a")) + "hello") + "a") + ("hello" + "hello"))) + "foo bar") + ("" + "a")) + "")) {
                case "\n\t":
                    _ = len(v2)
                case "hello":
                    v6 = (v6 - v6)
                case "a":
                    v19 := GenS1{F0: 600, F1: 653}
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
    }
    delete(v2, -1)
    fmt.Println(v13)
    if (uint64(v5) < 0) {
        for fi20 := 0; fi20 < 14; fi20++ {
            switch (("" + "a") + "\n\t") {
            case "foo bar":
                if ((-(v4)) <= v13) {
                    _ = len(v2)
                    var v21 float64 = v6
                    _ = v21
                } else {
                    switch (fi20 * fi20) {
                    case 9223372036854775807:
                        for fi22 := 0; fi22 < 7; fi22++ {
                            if true {
                                fmt.Println(v5, v6, fi22)
                                if true {
                                    var v23 int64 = 1
                                    var v24 int = fi20
                                    var v25 int16 = v4
                                    v26 := GenS1{F0: v23, F1: 0}
                                    switch "foo bar" {
                                    case "hello":
                                        var v27 float64 = float64(float32(v6))
                                        _ = v27
                                    }
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                var v28 uint64 = 42
                                _ = v28
                            }
                            v29 := []int{42, fi20, fi20, int(v3)}
                            var v30 string = "\n\t"
                            _ = fi22
                            _ = v29
                            _ = v30
                        }
                    case -655:
                        fmt.Println(v5, v5)
                    case 474:
                        switch "" {
                        case "foo bar":
                            v3 = v3
                        case "hello":
                            v31 := GenS1{F0: 42, F1: 582}
                            _ = v31
                        default:
                            fmt.Println(v5, v5, fi20)
                        }
                    }
                }
            }
            v32 := GenS2{F0: 2, F1: v13, F2: 107}
            _ = fi20
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v5 = 0
        _ = fi33
    }
    v13 += 0
    for fi34 := 0; fi34 < 14; fi34++ {
        if false {
            v13 = -671
            fmt.Println(v5, v13, v4)
            if false {
                _ = len(v1)
                var v35 int8 = -128
                if true {
                    switch (^(fi34)) {
                    case -2147483648:
                        for fi36 := 0; fi36 < 18; fi36++ {
                            v37 := []int{fi34, (^(fi36)), (fi34 - (-(fi36)))}
                            switch "hello" {
                            case "\n\t":
                                var v38 int64 = 1
                                _ = v38
                            case "foo bar":
                                var v39 float32 = (-2.25 - 0.0)
                                _ = v39
                            case "a":
                                if len(v37) > 0 {
                                    v37[0] = (fi36 * (^(fi36)))
                                }
                            default:
                                if true {
                                    fmt.Println(v5, fi36)
                                    v37 = append(v37, 1)
                                }
                                if false {
                                    v35 = (v35 % ((-110) | 1))
                                    var v40 string = "a"
                                    _ = v40
                                } else {
                                    if (42 > 42) {
                                        if (0 == 18446744073709551615) {
                                            v41 := GenS2{F0: 2, F1: (v4 ^ (v13 + (v4 + v4))), F2: 614}
                                            _ = v41
                                        }
                                        var v42 int8 = (v35 - v35)
                                        v37 = append(v37, 32767)
                                        _ = v42
                                    }
                                    var v43 int16 = v4
                                    var v44 uint = 4294967295
                                    _ = v43
                                    _ = v44
                                }
                            }
                            fmt.Println(v4, fi36)
                            _ = fi36
                            _ = v37
                        }
                    case 127:
                        v45 := GenS3{F0: 590, F1: 461}
                        _ = v45
                    default:
                        v6 = float64(float32((float64(0.0) + (v6 * -2.25))))
                    }
                    v46 := GenS1{F0: 42, F1: 18446744073709551615}
                    _ = v46
                } else {
                    v47 := GenS1{F0: -9223372036854775808, F1: 42}
                    fmt.Println(v6, v4)
                    var v48 float32 = -2.25
                    _ = v47
                    _ = v48
                }
                _ = v35
            }
        }
        _ = fi34
    }
    v49 := GenS4{F0: v4, F1: "a", F2: 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
