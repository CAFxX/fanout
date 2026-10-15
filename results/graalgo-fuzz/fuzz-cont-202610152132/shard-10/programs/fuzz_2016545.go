package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 int8 }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 int32 }
type GenS5 struct { F0 bool; F1 int; F2 int8 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 uint16, p1 int32, p2 int16) uint32 {
    v7 := []int{-924, 678}
    p0 += p0
    for fi8 := 0; fi8 < 17; fi8++ {
        p1 -= (p1 * -811)
        _ = fi8
    }
    switch "hello" {
    case "foo bar":
        v9 := GenS1{F0: (p1 >> ((1) & 63)), F1: (p0 % ((p0) | 1)), F2: 0}
        _ = v9
    case "hello":
        var v10 bool = true
        _ = v10
    case "\n\t":
        v11 := GenS2{F0: 128, F1: 0, F2: 1}
        _ = v11
    default:
        var v12 uint64 = 4294967295
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}
func genFn1(p0 uint, p1 int32, p2 uint8) int64 {
    var v13 bool = true
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS1{F0: p1, F1: 512, F2: 2147483648}
        _ = fi14
        _ = v15
    }
    v16 := []int{-801, 150}
    if false {
        v16 = append(v16, -329)
        _ = len(v16)
    } else {
        var v17 bool = v13
        fmt.Println(p1, v17, v13)
        if (42 == 898) {
            var v18 string = ("" + ("\n\t" + (((("" + "hello") + "a") + "hello") + "hello")))
            if v17 {
                switch v18 {
                case "\n\t":
                    if (!(9223372036854775807 != 17)) {
                        switch -1 {
                        case 830:
                            var v19 string = (((v18 + "") + (v18 + v18)) + v18)
                            _ = v19
                        case 143:
                            for fi20 := 0; fi20 < 3; fi20++ {
                                var v21 uint8 = 0
                                v18 = (v18 + "a")
                                _ = fi20
                                _ = v21
                            }
                        }
                        switch -903 {
                        case 1:
                            for fi22 := 0; fi22 < 20; fi22++ {
                                var v23 string = v18
                                v24 := GenS3{F0: 100, F1: p0, F2: 65535}
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    fmt.Println(v13)
                                    v23 = "a"
                                    _ = fi25
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case -2147483648:
                            v26 := GenS4{F0: true, F1: 0, F2: p1}
                            _ = v26
                        case 32767:
                            _ = len(v16)
                        }
                        switch "hello" {
                        case "foo bar":
                            var v27 uint64 = 42
                            _ = v27
                        case "hello":
                            var v28 string = ((v18 + "hello") + v18)
                            _ = v28
                        }
                    }
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    fmt.Println(p0, fi29, p0)
                    v30 := GenS2{F0: (fi29 % ((fi29) | 1)), F1: (p0 << (((p0 % (((p0 + (p0 * 42))) | 1))) & 63)), F2: -39}
                    _ = fi29
                    _ = v30
                }
            }
            _ = v18
        } else {
            p2 += (p2 - 0)
        }
        _ = v17
    }
    var v31 uint = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v16
    _ = v31
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 4294967295
    var v3 int16 = 306
    var v4 float32 = (0.0 - ((3.14159 - 3.14159) - -2.25))
    v5 := []string{"hello", ("a" + (("a" + "hello") + "")), ("foo bar" + ""), ""}
    var v6 int = v2
    switch v6 {
    case 0:
        if (0 <= 255) {
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 uint = 0
                for fi34 := 0; fi34 < 9; fi34++ {
                    v33 = v33
                    _ = fi34
                }
                _ = fi32
                _ = v33
            }
            v2 = (v6 | v6)
            if ((-(v2)) < (v6 / ((32767) | 1))) {
                var v35 int32 = 704
                _ = v35
            } else {
                fmt.Println(v3, v4)
            }
        } else {
            v3 = v3
            var v36 int8 = 0
            _ = v36
        }
    case 1:
        switch (-(v2)) {
        case -609:
            var v37 bool = (-90 != -69)
            _ = v37
        case 9223372036854775807:
            v38 := GenS5{F0: (true && (0 == -263)), F1: (v6 << ((2) & 63)), F2: -125}
            _ = v38
        }
    }
    v3 = (v3 + 42)
    v39 := GenS4{F0: (42 >= 42), F1: 2, F2: -2147483648}
    if len(v5) > 0 {
        v5[0] = ("hello" + "")
    }
    v5 = append(v5, "foo bar")
    if (v4 >= v4) {
        v40 := GenS6{F0: "hello", F1: 958}
        var v41 int = v2
        v3 = (v3 ^ 0)
        _ = v40
        _ = v41
    } else {
        for fi42 := 0; fi42 < 18; fi42++ {
            for fi43 := 0; fi43 < 10; fi43++ {
                v4 -= v4
                var v44 uint64 = 1
                _ = fi43
                _ = v44
            }
            v2 += (v6 >> ((255) & 63))
            for fi45 := 0; fi45 < 6; fi45++ {
                var v46 float64 = 0.0
                fmt.Println(v4)
                _ = fi45
                _ = v46
            }
            _ = fi42
        }
        var v47 int32 = -2147483648
        if true {
            var v48 uint16 = 26
            var v49 int32 = v47
            switch ("foo bar" + "hello") {
            case "hello":
                var v50 uint = uint(v48)
                _ = v50
            case "\n\t":
                v51 := GenS6{F0: "\n\t", F1: 131}
                _ = v51
            }
            _ = v48
            _ = v49
        }
        _ = v47
    }
    v4 -= (float32(1e10) / -2.25)
    var v52 int16 = 1
    v5 = append(v5, "")
    v53 := GenS7{F0: -2147483648, F1: -2147483648, F2: false}
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
