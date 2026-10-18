package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0(p0 int32, p1 int32) uint64 {
    var v5 int32 = p0
    p0 = v5
    p0 = (^(v5))
    if true {
        v5 -= p0
        switch 4294967295 {
        case 845:
            if false {
                v6 := GenS1{F0: 1}
                p1 += (-(v5))
                _ = v6
            }
        }
    } else {
        v7 := GenS2{F0: 42}
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                switch "hello" {
                case "foo bar":
                    switch (fi9 + (-(fi9))) {
                    case 0:
                        if (!false) {
                            var v10 uint32 = 42
                            v11 := GenS3{F0: 191, F1: (fi8 | (^(fi9))), F2: (v10 * 2)}
                            p1 = 39
                            _ = v10
                            _ = v11
                        } else {
                            var v12 float64 = (0.0 * 0.0)
                            var v13 uint64 = 0
                            var v14 int8 = 100
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        v15 := []int{fi8, fi8, (fi8 + (fi9 + -990)), -850}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = v15
                    default:
                        if (9223372036854775807 < 1) {
                            p1 -= 0
                            v16 := GenS3{F0: 255, F1: fi8, F2: 1}
                            v5 += v5
                            p1 = 2147483647
                            v17 := GenS4{F0: -32768, F1: 42}
                            p0 = (p1 ^ v5)
                            switch "foo bar" {
                            case "hello":
                                if (!false) {
                                    switch "hello" {
                                    case "":
                                        p0 += (p1 * -2147483648)
                                    }
                                    v18 := GenS4{F0: 115, F1: 32767}
                                    _ = v18
                                }
                            case "foo bar":
                                p0 -= -337
                            case "\n\t":
                                p1 -= (v5 * v5)
                            }
                            _ = v16
                            _ = v17
                        }
                        v19 := GenS5{F0: 843, F1: 0}
                        var v20 uint8 = 1
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi9
            }
            var v21 float64 = 1.5
            _ = fi8
            _ = v21
        }
        v22 := []int{2147483647, -2147483648, -990}
        _ = v22
        _ = v7
    }
    if true {
        switch ("a" + "") {
        case "hello":
            var v23 bool = false
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 int = 893
    var v4 int16 = -478
    fmt.Println(v3, v3)
    var v24 int64 = -9223372036854775808
    var v25 bool = false
    if (699 > 1) {
        var v26 uint = 4294967295
        if (!(!(128 < 17))) {
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS2{F0: v24}
                _ = fi27
                _ = v28
            }
        }
        _ = v26
    } else {
        switch ("hello" + "hello") {
        case "hello":
            if true {
                switch "\n\t" {
                case "\n\t":
                    v29 := GenS6{F0: "a", F1: 0}
                    _ = v29
                }
            }
        default:
            if ((65535 != 4294967295) && v25) {
                var v30 float64 = 0.1
                if (true || (-1 != -128)) {
                    v3 = (v3 & 42)
                    var v31 uint32 = 358
                    _ = v31
                } else {
                    fmt.Println(v24, v30)
                    switch v3 {
                    case -591:
                        var v32 uint8 = 28
                        _ = v32
                    case -9223372036854775808:
                        var v33 int16 = -325
                        _ = v33
                    case 4294967295:
                        v4 = -32768
                    default:
                        v30 = v30
                    }
                    v30 = 3.14159
                }
                fmt.Println(v30, v25)
                _ = v30
            } else {
                v24 = (v24 + 9223372036854775807)
                v24 = -1
            }
        }
        switch (("a" + "") + "") {
        case "a":
            v34 := v2[337]
            _ = v34
        }
    }
    v4 = v4
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 int = (v3 ^ (fi35 & 761))
        _ = fi35
        _ = v36
    }
    if (!v25) {
        v4 = (v4 * v4)
    }
    delete(v2, 298)
    _ = len(v2)
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 int64 = v24
        _ = fi37
        _ = v38
    }
    fmt.Println(v3, v3)
    fmt.Println(v25, v4, v25)
    if false {
        for fi39 := 0; fi39 < 2; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 uint32 = 65535
                _ = fi40
                _ = v41
            }
            if (1e10 == float64(3.14159)) {
                for fi42 := 0; fi42 < 12; fi42++ {
                    v4 = 804
                    _ = fi42
                }
                v43 := GenS7{F0: v24}
                _ = v43
            } else {
                switch "hello" {
                case "hello":
                    v25 = (42 != -1)
                case "foo bar":
                    var v44 uint64 = genFn0(-429, 2147483647)
                    _ = v44
                default:
                    v45 := GenS8{F0: v25, F1: 646, F2: -218}
                    _ = v45
                }
                v46 := []int{4294967295, v3, -274}
                _ = v46
            }
            _ = fi39
        }
    } else {
        var v47 uint32 = 2
        var v48 int32 = 2147483647
        _ = v47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
