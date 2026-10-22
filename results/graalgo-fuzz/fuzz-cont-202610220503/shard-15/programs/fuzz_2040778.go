package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int; F2 string }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 bool; F1 uint64 }
func genFn0(p0 bool) int64 {
    fmt.Println(p0)
    var v4 int16 = 1000
    v5 := GenS1{F0: 677, F1: 200, F2: 2}
    fmt.Println(p0)
    var v6 bool = p0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint) string {
    if ((58 >= 42) || true) {
        if (-2.25 > 1e10) {
            if false {
                var v7 uint64 = 1
                _ = v7
            } else {
                switch -2147483648 {
                case 353:
                    var v8 float32 = 3.14159
                    _ = v8
                }
                for fi9 := 0; fi9 < 11; fi9++ {
                    if (genFn0(true) != int64(fi9)) {
                        switch (fi9 + (^(fi9))) {
                        case -9223372036854775808:
                            fmt.Println(fi9, p0)
                        case 732:
                            fmt.Println(p0, fi9)
                        default:
                            v10 := []int{(fi9 * 2147483647), -9223372036854775808, -9223372036854775808}
                            _ = v10
                        }
                        if (!false) {
                            var v11 bool = true
                            var v12 string = "\n\t"
                            switch (v12 + v12) {
                            case "hello":
                                if (0 <= 32767) {
                                    var v13 string = (v12 + ("hello" + "\n\t"))
                                    _ = v13
                                }
                            case "\n\t":
                                var v14 int64 = -9223372036854775808
                                _ = v14
                            default:
                                var v15 uint32 = 0
                                _ = v15
                            }
                            _ = v11
                            _ = v12
                        } else {
                            p0 = (p0 | (^(p0)))
                            for fi16 := 0; fi16 < 4; fi16++ {
                                v17 := GenS1{F0: 4294967295, F1: 144, F2: 1}
                                v18 := GenS2{F0: 0}
                                var v19 uint16 = 255
                                _ = fi16
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                            v20 := []int{-1, fi9, (fi9 + 0)}
                            _ = v20
                        }
                    }
                    p0 += (p0 - (p0 | p0))
                    v21 := []int{}
                    _ = fi9
                    _ = v21
                }
                if false {
                    v22 := GenS3{F0: 9223372036854775807, F1: 0, F2: ("" + ("foo bar" + "\n\t"))}
                    _ = v22
                }
            }
        } else {
            var v23 uint32 = 943
            v24 := GenS2{F0: 0}
            _ = v23
            _ = v24
        }
        var v25 bool = true
        _ = v25
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    var v26 float32 = float32(-2.25)
    v27 := []int{0}
    for fi28 := 0; fi28 < 1; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            p0 = p0
            v26 = (0.0 - float32(1e10))
            fmt.Println(v26, fi28)
            _ = fi29
        }
        _ = len(v27)
        var v30 uint = p0
        _ = fi28
        _ = v30
    }
    _ = p0
    _ = v26
    _ = v27
    return ("a" + ((("foo bar" + ("\n\t" + "")) + ((("" + ("foo bar" + "\n\t")) + "foo bar") + "\n\t")) + "\n\t"))
}

func main() {
    v1 := []string{"", ("foo bar" + "a")}
    var v2 int = 32768
    var v3 float64 = (((1.5 + (1e10 + 3.14159)) / (-2.25 - 1.5)) + float64(float32(((0.1 + 3.14159) + 0.1))))
    v2 -= (v2 - (v2 + (-(v2))))
    for fi31 := 0; fi31 < 17; fi31++ {
        fmt.Println(fi31, v3)
        _ = fi31
    }
    switch (v2 | v2) {
    case -425:
        var v32 uint64 = 1
        _ = v32
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "\n\t":
        for fi33 := 0; fi33 < 7; fi33++ {
            if true {
                switch "hello" {
                case "\n\t":
                    v3 = 1.5
                default:
                    v3 += v3
                }
                v3 = v3
                fmt.Println(fi33)
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 int8 = 112
                    v2 = (fi34 | (^(fi33)))
                    _ = fi34
                    _ = v35
                }
            }
            var v36 int32 = 1
            v37 := GenS4{F0: v36, F1: "\n\t"}
            _ = fi33
            _ = v36
            _ = v37
        }
    case "hello":
        for fi38 := 0; fi38 < 5; fi38++ {
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 uint32 = 648
                fmt.Println(v2, v3, fi38)
                _ = fi39
                _ = v40
            }
            var v41 bool = true
            if v41 {
                v42 := GenS5{F0: v41, F1: 0}
                _ = v42
            }
            _ = fi38
            _ = v41
        }
    }
    var v43 int = (v2 ^ -1)
    switch "hello" {
    case "":
        if false {
            var v44 uint8 = 0
            v43 = v2
            _ = v44
        }
    }
    v45 := GenS1{F0: 823, F1: 127, F2: 238}
    for fi46 := 0; fi46 < 10; fi46++ {
        fmt.Println(fi46)
        _ = fi46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v45
}
