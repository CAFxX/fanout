package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint }
func genFn0(p0 int32) int8 {
    v7 := GenS1{F0: 135, F1: 93, F2: 686}
    if true {
        for fi8 := 0; fi8 < 5; fi8++ {
            p0 += (p0 & -917)
            fmt.Println(fi8, fi8)
            _ = fi8
        }
        v9 := []int{4294967295, 42}
        _ = v9
    }
    if ((0.1 != (1.5 * (1.5 * (1.5 * 1e10)))) || (((!(!(((p0 / ((p0) | 1)) != 42) && (-119 > 0)))) && false) || true)) {
        p0 = p0
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(p0, fi10)
        p0 = p0
        fmt.Println(p0, fi10)
        _ = fi10
    }
    switch ("\n\t" + "") {
    case "hello":
        p0 = 1
    case "a":
        if false {
            var v11 bool = false
            v12 := []int{42, 32768, 2147483647, 42}
            fmt.Println(v11, v11)
            _ = v11
            _ = v12
        }
    case "":
        if true {
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 string = "\n\t"
                    fmt.Println(fi14)
                    _ = fi14
                    _ = v15
                }
                var v16 int32 = p0
                p0 = v16
                _ = fi13
                _ = v16
            }
            fmt.Println(p0)
            p0 += p0
        }
    }
    _ = p0
    _ = v7
    return -1
}
func genFn1() int {
    if false {
        if (322 == -32768) {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int = fi17
                _ = fi17
                _ = v18
            }
            var v19 uint64 = 1
            _ = v19
        } else {
            v20 := []int{-2147483648}
            v21 := GenS1{F0: 65535, F1: -108, F2: -576}
            _ = v20
            _ = v21
        }
        var v22 bool = false
        v23 := GenS2{F0: 1, F1: 255}
        _ = v22
        _ = v23
    }
    var v24 uint8 = 255
    v25 := GenS2{F0: 65535, F1: 65535}
    _ = v24
    _ = v25
    return -429
}

func main() {
    v1 := []uint32{42, 1, 2147483648}
    var v2 float64 = (float64((3.14159 * 3.14159)) / (0.0 - 3.14159))
    var v3 float32 = (1.5 * (0.0 / 1.0))
    var v4 uint16 = 0
    v5 := []string{"a", ("" + ""), "foo bar"}
    var v6 int32 = 0
    var v26 int8 = 20
    var v27 int32 = v6
    v26 = (v26 - -79)
    v28 := GenS3{F0: 486, F1: "hello"}
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 string = ("" + (((("hello" + "hello") + "") + "foo bar") + "hello"))
        if false {
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(v6, v30)
                v26 += v26
                _ = fi31
            }
        } else {
            if false {
                var v32 bool = (2147483648 < 2147483647)
                if (3.14159 == 1e10) {
                    var v33 uint8 = 127
                    v34 := GenS4{F0: (v4 & uint16(v33))}
                    _ = v33
                    _ = v34
                }
                for fi35 := 0; fi35 < 13; fi35++ {
                    v1 = append(v1, 491)
                    var v36 int16 = 42
                    switch "foo bar" {
                    case "a":
                        var v37 int = fi29
                        _ = v37
                    case "":
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                    case "foo bar":
                        v38 := GenS1{F0: 0, F1: genFn0((v27 & 0)), F2: (fi35 / ((fi35) | 1))}
                        _ = v38
                    default:
                        switch "a" {
                        case "hello":
                            v26 += v26
                        default:
                            v4 = v4
                        }
                    }
                    _ = fi35
                    _ = v36
                }
                _ = v32
            }
        }
        _ = fi29
        _ = v30
    }
    switch genFn1() {
    case 336:
        fmt.Println(v27, v4, v3)
    default:
        for fi39 := 0; fi39 < 19; fi39++ {
            v27 = v6
            var v40 uint16 = 42
            var v41 bool = (false || true)
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := GenS5{F0: (fi43 | (fi42 * fi42)), F1: 0}
            if (952 == 850) {
                var v45 uint64 = 131
                _ = v45
            }
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    switch "foo bar" {
    case "":
        fmt.Println(v3, v6)
    }
    if ((("hello" + ((("" + (("foo bar" + "\n\t") + "")) + "foo bar") + (("\n\t" + "hello") + "\n\t"))) >= "foo bar") && true) {
        fmt.Println(v26, v4, v4)
        var v46 uint = 4294967295
        v6 = 746
        _ = v46
    } else {
        for fi47 := 0; fi47 < 3; fi47++ {
            fmt.Println(v27, v26, v3)
            switch "" {
            case "foo bar":
                var v48 float32 = float32((0.1 + 3.14159))
                _ = v48
            default:
                v26 = v26
            }
            switch genFn1() {
            case 128:
                for fi49 := 0; fi49 < 3; fi49++ {
                    var v50 float64 = v2
                    _ = fi49
                    _ = v50
                }
            }
            _ = fi47
        }
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        for fi52 := 0; fi52 < 2; fi52++ {
            switch "hello" {
            case "foo bar":
                var v53 uint64 = 18446744073709551615
                _ = v53
            case "a":
                v4 -= (^(v4))
            }
            fmt.Println(v6, v27)
            _ = fi52
        }
        v54 := GenS1{F0: 208, F1: v26, F2: fi51}
        _ = fi51
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
