package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 string; F1 bool; F2 int8 }
func genFn0() uint64 {
    v4 := GenS1{F0: -1}
    var v5 uint8 = 123
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS2{F0: 1}
        for fi8 := 0; fi8 < 6; fi8++ {
            v5 -= (v5 ^ v5)
            var v9 uint16 = 164
            _ = fi8
            _ = v9
        }
        var v10 uint16 = 42
        _ = fi6
        _ = v10
        _ = v7
    }
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int, p1 string) uint32 {
    v11 := GenS1{F0: -1}
    switch "\n\t" {
    case "a":
        p0 = p0
    default:
        var v12 uint64 = genFn0()
        _ = v12
    }
    if (0 < -119) {
        p1 = p1
        v13 := []int{p0, (p0 << ((737) & 63)), (^(p0)), 32767}
        _ = v13
    }
    if ((false && true) && (p1 == p1)) {
        if false {
            switch "a" {
            case "\n\t":
                p0 += 32767
            case "a":
                var v14 int16 = -452
                _ = v14
            case "foo bar":
                p1 = ((p1 + ("a" + (p1 + p1))) + (p1 + p1))
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                p1 = ((p1 + "\n\t") + "")
                _ = fi15
            }
            for fi16 := 0; fi16 < 16; fi16++ {
                p1 = (("\n\t" + p1) + "")
                _ = fi16
            }
        } else {
            switch "" {
            case "":
                if (0 >= 71) {
                    fmt.Println(p1, p0)
                    v17 := GenS2{F0: 1}
                    v18 := GenS3{F0: 42, F1: ("a" + ("hello" + "\n\t")), F2: 0}
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint8 = 0
                    _ = v19
                }
            default:
                p1 = (((p1 + ("a" + p1)) + ("" + "a")) + (((p1 + (((p1 + "foo bar") + p1) + p1)) + p1) + "foo bar"))
            }
            p1 = p1
            switch (p0 | -889) {
            case 128:
                for fi20 := 0; fi20 < 20; fi20++ {
                    v21 := []int{}
                    _ = fi20
                    _ = v21
                }
            default:
                p0 = p0
            }
        }
        p1 = "foo bar"
        v22 := GenS4{F0: p1, F1: (2 <= 0), F2: 42}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 8; fi23++ {
            fmt.Println(p1, fi23)
            if (685 == -9223372036854775808) {
                p1 = p1
                v24 := []int{(fi23 ^ p0)}
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 uint32 = 657
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 uint32 = 2147483648
                var v29 uint64 = 0
                var v30 uint16 = 738
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi23
        }
        p0 = 216
        for fi31 := 0; fi31 < 17; fi31++ {
            fmt.Println(p0)
            _ = fi31
        }
    }
    var v32 uint32 = 528
    _ = p0
    _ = p1
    _ = v11
    _ = v32
    return v32
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 1}
    v2 := make(map[string]int)
    var v3 int = 42
    if false {
        v3 += v3
        v3 -= (-(v3))
    } else {
        v3 += v3
        var v33 uint = 18446744073709551615
        var v34 bool = true
        _ = v33
        _ = v34
    }
    v35 := GenS3{F0: 206, F1: "a", F2: 766}
    for fi36 := 0; fi36 < 15; fi36++ {
        if (-2.25 >= 0.1) {
            var v37 int = fi36
            _ = v37
        }
        _ = fi36
    }
    if (-76 <= 43) {
        v3 = 2147483647
        fmt.Println(v3)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = v3
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 float32 = (float32(0.0) - 1.5)
        for fi40 := 0; fi40 < 19; fi40++ {
            fmt.Println(v39)
            _ = fi40
        }
        v2[("\n\t" + "")] = (-(v3))
        _ = fi38
        _ = v39
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        if ((!false) && true) {
            switch (fi41 | -413) {
            case 2147483647:
                for fi42 := 0; fi42 < 15; fi42++ {
                    var v43 int16 = 0
                    v3 += -437
                    switch "" {
                    case "\n\t":
                        v44 := GenS4{F0: "hello", F1: (-17 <= 104), F2: -1}
                        _ = v44
                    case "hello":
                        var v45 int = fi42
                        _ = v45
                    }
                    _ = fi42
                    _ = v43
                }
            case 32767:
                if (false && true) {
                    var v46 int16 = 0
                    if (!true) {
                        var v47 uint32 = 730
                        _ = v47
                    } else {
                        fmt.Println(v46)
                        v48 := GenS1{F0: v46}
                        var v49 float64 = 3.14159
                        _ = v48
                        _ = v49
                    }
                    _ = v46
                } else {
                    var v50 uint8 = 127
                    var v51 int64 = -9223372036854775808
                    fmt.Println(fi41, v3)
                    _ = v50
                    _ = v51
                }
            case 4294967295:
                v3 += 42
            default:
                for fi52 := 0; fi52 < 17; fi52++ {
                    var v53 bool = (-2147483648 >= 42)
                    v53 = v53
                    _ = fi52
                    _ = v53
                }
            }
        } else {
            if true {
                if true {
                    if false {
                        fmt.Println(fi41)
                    }
                    v3 = (fi41 / ((v3) | 1))
                    v3 -= fi41
                } else {
                    v54 := GenS4{F0: ("" + ("\n\t" + ("" + ""))), F1: true, F2: -128}
                    switch ("a" + "hello") {
                    case "\n\t":
                        var v55 uint32 = 2147483647
                        _ = v55
                    }
                    v3 += v3
                    _ = v54
                }
                v2["a"] = (fi41 | (fi41 + (fi41 / ((-741) | 1))))
            } else {
                _ = len(v1)
                v3 -= v3
            }
            fmt.Println(v3, fi41)
            v3 = (fi41 / (((v3 & 32767)) | 1))
        }
        _ = fi41
    }
    if (!((2 < 560) && true)) {
        var v56 int64 = 0
        v3 = (v3 | v3)
        _ = v56
    } else {
        v3 -= (v3 - v3)
        v57 := GenS3{F0: 42, F1: "a", F2: 65535}
        _ = v57
    }
    var v58 uint = 87
    v3 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v58
}
