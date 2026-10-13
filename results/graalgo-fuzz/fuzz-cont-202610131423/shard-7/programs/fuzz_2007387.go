package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int8; F1 int; F2 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint) uint8 {
    v7 := []int{-9223372036854775808}
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 14; fi8++ {
            if false {
                fmt.Println(p0)
                var v9 int64 = 42
                _ = v9
            } else {
                v7 = append(v7, fi8)
            }
            _ = fi8
        }
    case "foo bar":
        if (!(!(0 <= 0))) {
            var v10 int32 = 378
            _ = v10
        } else {
            if (42 <= 806) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := GenS1{F0: 0, F1: -136, F2: 18446744073709551615}
                    _ = len(v7)
                    _ = fi11
                    _ = v12
                }
                v7 = append(v7, -381)
            }
            var v13 int16 = -880
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 int32 = 264
                v16 := GenS2{F0: 2, F1: 2147483648}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        }
    case "hello":
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int = 0
            p0 += (^(p0))
            _ = fi17
            _ = v18
        }
    }
    var v19 uint8 = 2
    switch -684 {
    case -148:
        if (true || ("hello" <= "\n\t")) {
            switch 1 {
            case -150:
                fmt.Println(p0, p0)
            case 32768:
                if false {
                    var v20 string = "a"
                    _ = v20
                }
            case 127:
                switch -987 {
                case -787:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        p0 = p0
                        switch (("hello" + "a") + "foo bar") {
                        case "foo bar":
                            if true {
                                var v22 int64 = -1
                                _ = v22
                            } else {
                                p0 -= (p0 & p0)
                                fmt.Println(v19, fi21)
                                v19 = (v19 >> (((^(p0))) & 63))
                            }
                        case "a":
                            var v23 bool = true
                            _ = v23
                        default:
                            var v24 int32 = -1
                            _ = v24
                        }
                        v19 = 128
                        _ = fi21
                    }
                case 32768:
                    p0 = (p0 << (((p0 ^ (p0 % ((p0) | 1)))) & 63))
                }
            default:
                if true {
                    v25 := GenS1{F0: (v19 << ((p0) & 63)), F1: -742, F2: 2}
                    v26 := GenS1{F0: v19, F1: -116, F2: 787}
                    _ = v25
                    _ = v26
                } else {
                    fmt.Println(v19, p0)
                    switch "a" {
                    case "":
                        var v27 int32 = 0
                        _ = v27
                    case "a":
                        var v28 float64 = 0.1
                        _ = v28
                    }
                    fmt.Println(p0, p0)
                }
            }
        }
    case 127:
        var v29 int16 = -32768
        _ = v29
    default:
        v30 := GenS3{F0: 58, F1: 0, F2: (float32((3.14159 - 0.1)) >= ((1.5 * 1e10) * 0.0))}
        _ = v30
    }
    _ = p0
    _ = v19
    _ = v7
    return (v19 & 2)
}

func main() {
    var v1 int64 = 42
    var v2 int16 = 32767
    v3 := make(map[int]int)
    var v4 bool = (false && true)
    var v5 int32 = 42
    var v6 int = 4294967295
    v1 = v1
    fmt.Println(v1, v1)
    var v31 bool = v4
    v31 = true
    fmt.Println(v5, v1)
    v32 := []int{v6, 2147483647, v6, v6}
    if ("hello" <= "\n\t") {
        v3[(v6 * 32768)] = 915
        if v31 {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v31, v6)
            case "\n\t":
                switch "a" {
                case "":
                    switch (^(v6)) {
                    case 1:
                        _ = len(v3)
                    default:
                        v32 = append(v32, (v6 ^ (-(v6))))
                    }
                }
            default:
                fmt.Println(v5, v2, v4)
            }
            if v4 {
                v32 = append(v32, v6)
                switch 32767 {
                case 0:
                    v33 := GenS1{F0: 139, F1: v5, F2: 18446744073709551615}
                    _ = v33
                }
                _ = len(v3)
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                var v35 float32 = 0.1
                _ = fi34
                _ = v35
            }
        }
        var v36 string = "\n\t"
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        v5 += (v5 ^ (v5 & -1))
        var v38 int64 = v1
        _ = fi37
        _ = v38
    }
    v39 := GenS4{F0: 4294967295}
    fmt.Println(v5, v2, v4)
    v5 += 965
    switch -106 {
    case 32767:
        if (("a" + "") > (("" + "foo bar") + (("hello" + ("hello" + ("foo bar" + "hello"))) + "foo bar"))) {
            for fi40 := 0; fi40 < 10; fi40++ {
                v31 = (v31 || v4)
                _ = fi40
            }
            var v41 bool = (v31 && v4)
            _ = v41
        }
    case 42:
        v42 := GenS1{F0: 1, F1: (v5 / (((v5 * v5)) | 1)), F2: 4294967295}
        _ = v42
    case 18:
        fmt.Println(v6, v5)
    default:
        var v43 uint16 = 930
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
