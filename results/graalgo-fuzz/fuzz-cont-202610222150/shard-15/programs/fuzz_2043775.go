package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 int; F1 uint; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
func genFn0(p0 int16, p1 uint32) uint8 {
    var v6 float32 = 1e10
    var v7 uint16 = 0
    fmt.Println(p1, p0, v6)
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: -1, F1: (!false)}
        switch "\n\t" {
        case "a":
            var v10 uint8 = 128
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            if true {
                v12 := GenS2{F0: 127, F1: (!true), F2: 42}
                if false {
                    var v13 uint = 0
                    _ = v13
                } else {
                    fmt.Println(fi11, fi8)
                    var v14 int16 = (p0 >> ((2) & 63))
                    var v15 int = fi11
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                switch "hello" {
                case "\n\t":
                    if ((!false) || (0 < 127)) {
                        v16 := []int{}
                        switch "hello" {
                        case "hello":
                            var v17 uint64 = 4294967295
                            _ = v17
                        }
                        _ = v16
                    } else {
                        v6 -= float32(0.0)
                    }
                case "foo bar":
                    v18 := []int{fi11, 246, fi8}
                    var v19 float32 = (float32((0.0 * 0.0)) + (((((v6 * v6) + 1.5) * v6) - (-2.25 / 1.0)) + v6))
                    fmt.Println(v19, fi8, v7)
                    _ = v18
                    _ = v19
                default:
                    var v20 uint64 = 2
                    _ = v20
                }
                v21 := []int{(fi11 & fi8), fi8}
                var v22 uint8 = 128
                _ = v21
                _ = v22
            }
            if ((^(p1)) != (^(p1))) {
                v23 := []int{}
                _ = v23
            }
            _ = fi11
        }
        _ = fi8
        _ = v9
    }
    var v24 int16 = -32768
    _ = p0
    _ = p1
    _ = v24
    _ = v6
    _ = v7
    return 152
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := []uint32{4294967295, 482, 2, 42}
    var v5 int32 = 1
    delete(v3, "a")
    _ = len(v4)
    var v25 int32 = -1
    switch ("" + "foo bar") {
    case "foo bar":
        v5 -= v5
    case "\n\t":
        v26 := GenS2{F0: 42, F1: (!(42 >= 0)), F2: 1}
        _ = v26
    case "hello":
        for fi27 := 0; fi27 < 6; fi27++ {
            if false {
                v25 += v25
                var v28 int64 = -1
                var v29 uint32 = 65535
                _ = v28
                _ = v29
            }
            v30 := GenS3{F0: 1, F1: -9223372036854775808, F2: (128 > 127)}
            _ = fi27
            _ = v30
        }
    }
    v31 := GenS3{F0: 2, F1: 1, F2: true}
    v5 = v5
    v5 += (v5 * v5)
    if (0.1 >= float64((-2.25 - float32(-2.25)))) {
        fmt.Println(v5)
        for fi32 := 0; fi32 < 4; fi32++ {
            v3["a"] = (-(fi32))
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if ((v5 < v25) || true) {
                    v25 = (v5 ^ 42)
                    var v33 float32 = (0.1 + 0.0)
                    _ = v33
                } else {
                    v25 -= (v25 % (((v5 % ((538) | 1))) | 1))
                }
            } else {
                fmt.Println(v5, v25, v25)
                v34 := v3[(("a" + "\n\t") + "hello")]
                fmt.Println(v25, v34)
                _ = v34
            }
            _ = fi32
        }
        _ = len(v2)
    }
    v3["\n\t"] = -712
    v25 += (v5 / (((v5 - (^(v25)))) | 1))
    v5 -= v5
    for fi35 := 0; fi35 < 10; fi35++ {
        if (!true) {
            for fi36 := 0; fi36 < 15; fi36++ {
                _ = len(v4)
                switch (fi35 + fi35) {
                case -9223372036854775808:
                    v37 := GenS4{F0: 127, F1: 18446744073709551615, F2: ("a" + ((("hello" + ("hello" + "\n\t")) + "\n\t") + ""))}
                    _ = v37
                default:
                    v5 = (v25 | v5)
                }
                _ = fi36
            }
            for fi38 := 0; fi38 < 3; fi38++ {
                v25 -= v25
                fmt.Println(v5)
                var v39 int64 = 1
                _ = fi38
                _ = v39
            }
            switch fi35 {
            case 128:
                v3[((("" + ("\n\t" + ("foo bar" + "foo bar"))) + "a") + ("hello" + ""))] = fi35
            case 42:
                if ((false && (!true)) || (true && (true && false))) {
                    v25 += (v5 * (v25 + (v25 >> ((18446744073709551615) & 63))))
                }
            default:
                for fi40 := 0; fi40 < 2; fi40++ {
                    fmt.Println(v5)
                    v41 := GenS5{F0: 0}
                    _ = fi40
                    _ = v41
                }
            }
        } else {
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 int8 = -77
                _ = fi42
                _ = v43
            }
            switch "" {
            case "foo bar":
                v25 = (v25 << ((18446744073709551615) & 63))
            }
            v44 := GenS6{F0: fi35}
            _ = v44
        }
        var v45 string = "\n\t"
        _ = fi35
        _ = v45
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        fmt.Println(v5)
        v2[201] = true
        _ = fi46
    }
    fmt.Println(v5)
    fmt.Println(v5)
    switch 1 {
    case 1:
        v47 := GenS5{F0: 1}
        _ = v47
    case -9223372036854775808:
        if true {
            var v48 float64 = 1e10
            _ = v48
        } else {
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            default:
                v25 = 42
            }
            v49 := v2[1]
            v50 := GenS7{F0: 255}
            _ = v49
            _ = v50
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
