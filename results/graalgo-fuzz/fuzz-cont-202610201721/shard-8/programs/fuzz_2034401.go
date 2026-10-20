package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int64; F1 int32; F2 int64 }
func genFn0(p0 int16, p1 int16, p2 int) string {
    var v7 uint8 = 2
    switch (p2 & 4294967295) {
    case 2147483647:
        v8 := GenS1{F0: 426, F1: 1, F2: 42}
        _ = v8
    }
    v9 := GenS2{F0: (^(p2))}
    switch ("a" + "a") {
    case "foo bar":
        fmt.Println(p0, p1)
    case "hello":
        if (p1 < -32768) {
            v7 = v7
            switch p2 {
            case 0:
                switch ("foo bar" + "hello") {
                case "\n\t":
                    fmt.Println(p1, v7)
                }
            case 889:
                var v10 uint16 = 548
                _ = v10
            default:
                if true {
                    if false {
                        v11 := []int{}
                        _ = v11
                    }
                    v12 := []int{}
                    switch "foo bar" {
                    case "hello":
                        if ((^(p2)) > p2) {
                            p1 += (p0 / (((p1 - p1)) | 1))
                        } else {
                            fmt.Println(p0, p2)
                            if true {
                                p2 = (^(p2))
                                v13 := GenS2{F0: int(p0)}
                                _ = v13
                            } else {
                                var v14 float64 = (0.1 - 0.0)
                                var v15 uint32 = 1
                                for fi16 := 0; fi16 < 9; fi16++ {
                                    for fi17 := 0; fi17 < 20; fi17++ {
                                        fmt.Println(v14)
                                        var v18 float64 = (0.1 / v14)
                                        _ = fi17
                                        _ = v18
                                    }
                                    _ = fi16
                                }
                                _ = v14
                                _ = v15
                            }
                            var v19 bool = ((float64(((float32(-2.25) + 0.0) + 3.14159)) != (0.1 + 1.5)) && (!true))
                            _ = v19
                        }
                    }
                    _ = v12
                } else {
                    var v20 int16 = p0
                    _ = v20
                }
            }
        }
    default:
        v21 := GenS3{F0: 42}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return "hello"
}
func genFn1(p0 uint16, p1 uint8, p2 uint16) int8 {
    p0 = p0
    if (false || true) {
        switch 128 {
        case 0:
            p2 += (^(p0))
        case 9223372036854775807:
            var v22 uint8 = p1
            _ = v22
        }
    } else {
        p2 += (p0 | (p0 & 255))
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 string = ((("foo bar" + "") + "a") + "")
    v2 := make(map[int]int)
    var v3 int8 = 42
    var v4 bool = ((v1 + "a") <= (((v1 + ("foo bar" + "")) + v1) + "\n\t"))
    var v5 uint64 = 4294967295
    v6 := make(map[uint32]string)
    for fi23 := 0; fi23 < 2; fi23++ {
        switch v1 {
        case "":
            v1 = (v1 + "hello")
        default:
            var v24 float64 = 1.5
            _ = v24
        }
        _ = fi23
    }
    v25 := []int{9223372036854775807, 0, -1}
    var v26 uint32 = 65535
    v5 -= (v5 * v5)
    if v4 {
        v27 := GenS4{F0: 255}
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        if v4 {
            if (v26 >= 589) {
                for fi29 := 0; fi29 < 13; fi29++ {
                    if v4 {
                        _ = len(v25)
                        v30 := GenS1{F0: 255, F1: 2, F2: v26}
                        _ = v30
                    }
                    fmt.Println(v3)
                    _ = fi29
                }
                v31 := GenS3{F0: 2}
                _ = v31
            } else {
                v2[fi28] = (fi28 * 9223372036854775807)
            }
            switch 4294967295 {
            case 9223372036854775807:
                fmt.Println(v1, v4)
            }
            switch (-(fi28)) {
            case -9223372036854775808:
                for fi32 := 0; fi32 < 20; fi32++ {
                    switch v1 {
                    case "a":
                        fmt.Println(v3, v5, v4)
                    case "foo bar":
                        _ = len(v6)
                    }
                    _ = fi32
                }
            case -107:
                if true {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        fmt.Println(v26)
                        v34 := GenS5{F0: 1, F1: 88}
                        _ = fi33
                        _ = v34
                    }
                    if v4 {
                        v4 = (!true)
                        v1 = v1
                    } else {
                        v3 = (-(v3))
                        v35 := GenS1{F0: 42, F1: 1, F2: 292}
                        v26 = (v26 + (v26 * 2147483647))
                        _ = v35
                    }
                    v4 = v4
                } else {
                    v36 := GenS6{F0: -9223372036854775808, F1: 1, F2: 0}
                    _ = v36
                }
            case 32767:
                if v4 {
                    var v37 string = v1
                    _ = v37
                } else {
                    v38 := v6[v26]
                    v1 = ("\n\t" + ("a" + v38))
                    _ = v38
                }
            }
        }
        _ = fi28
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS2{F0: -676}
        if len(v25) > 0 {
            v25[0] = (^(fi39))
        }
        v41 := GenS3{F0: 4294967295}
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            switch (("foo bar" + "\n\t") + v1) {
            case "hello":
                var v44 uint32 = v26
                _ = v44
            default:
                fmt.Println(fi42)
            }
            var v45 float64 = -2.25
            _ = fi43
            _ = v45
        }
        v4 = ((18446744073709551615 <= (v5 << ((0) & 63))) || ((fi42 > (fi42 + fi42)) || (-2.25 != 1e10)))
        if false {
            v4 = (((v1 + "a") <= (("foo bar" + "\n\t") + "a")) || true)
            v5 += 42
            var v46 uint = 4294967295
            _ = v46
        }
        _ = fi42
    }
    if ((v4 || v4) && v4) {
        v47 := GenS3{F0: v5}
        v5 -= (v5 * v5)
        _ = v47
    }
    v3 += genFn1(925, 2, 32767)
    if len(v25) > 0 {
        v25[0] = 32768
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
