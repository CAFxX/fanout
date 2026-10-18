package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 bool; F1 int32; F2 string }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32, p1 int16, p2 uint8) int {
    if ((255 < (p2 ^ p2)) || (false && true)) {
        var v5 uint = 42
        var v6 bool = true
        switch -2147483648 {
        case 32768:
            var v7 uint8 = p2
            _ = v7
        case 896:
            p1 = p1
        }
        _ = v5
        _ = v6
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{2147483647, fi8}
            p0 -= p0
            switch fi8 {
            case 315:
                p2 += p2
            case 0:
                v10 := GenS1{F0: (p1 + p1), F1: "hello", F2: 2}
                _ = v10
            default:
                switch (fi8 ^ fi8) {
                case 1:
                    v11 := GenS2{F0: "foo bar"}
                    _ = v11
                }
            }
            _ = fi8
            _ = v9
        }
        v12 := GenS3{F0: 0}
        for fi13 := 0; fi13 < 13; fi13++ {
            switch "foo bar" {
            case "a":
                for fi14 := 0; fi14 < 20; fi14++ {
                    if true {
                        fmt.Println(p2, fi13)
                        v15 := []int{(fi13 / (((fi13 % ((fi13) | 1))) | 1)), -9223372036854775808, 127, fi14}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = v15
                    } else {
                        fmt.Println(p0)
                        v16 := []int{fi13, (fi13 % (((^(fi13))) | 1)), int(p0), 32768}
                        _ = v16
                    }
                    fmt.Println(fi14, p0, p1)
                    v17 := []int{fi14, (fi14 % ((fi14) | 1)), (-(fi14))}
                    if ((!true) || true) {
                        if false {
                            var v18 float32 = 3.14159
                            fmt.Println(fi13, fi14)
                            var v19 string = "a"
                            var v20 int8 = -19
                            var v21 uint8 = 2
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        p0 += p0
                        v22 := GenS2{F0: "a"}
                        p1 -= (-(p1))
                        _ = v22
                    } else {
                        var v23 uint16 = 42
                        fmt.Println(fi13)
                        _ = v23
                    }
                    _ = fi14
                    _ = v17
                }
            }
            v24 := GenS1{F0: p1, F1: (("a" + "hello") + ("foo bar" + "foo bar")), F2: 32767}
            if false {
                if false {
                    v25 := GenS4{F0: ("foo bar" + "foo bar"), F1: (p1 & -59)}
                    v26 := []int{fi13, (fi13 + fi13), fi13}
                    _ = v25
                    _ = v26
                } else {
                    if (!false) {
                        v27 := []int{-1}
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                        _ = v27
                    }
                }
                var v28 uint64 = 332
                switch "\n\t" {
                case "a":
                    var v29 uint16 = 1
                    _ = v29
                case "\n\t":
                    v30 := GenS3{F0: -9}
                    _ = v30
                }
                _ = v28
            }
            _ = fi13
            _ = v24
        }
        _ = v12
    }
    v31 := []int{242, 127, 32767, 4294967295}
    var v32 uint64 = 522
    if true {
        v33 := GenS5{F0: (42 < 1), F1: (p0 * 42), F2: (("" + "\n\t") + "")}
        _ = v33
    } else {
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
        var v34 uint16 = 1
        _ = v34
    }
    if (!false) {
        v32 = (^(v32))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    return -9223372036854775808
}
func genFn1(p0 uint8) int16 {
    p0 = (^(p0))
    v35 := []int{32768, 42}
    _ = p0
    _ = v35
    return -1
}

func main() {
    var v1 uint8 = 51
    v2 := make(map[uint32]bool)
    var v3 bool = false
    var v4 uint64 = 4294967295
    for fi36 := 0; fi36 < 6; fi36++ {
        var v37 bool = (32767 <= 265)
        for fi38 := 0; fi38 < 17; fi38++ {
            if v37 {
                var v39 string = "hello"
                v40 := GenS1{F0: genFn1(v1), F1: v39, F2: 65535}
                _ = v39
                _ = v40
            }
            v1 = 117
            _ = fi38
        }
        v41 := []int{fi36, -2147483648}
        _ = fi36
        _ = v37
        _ = v41
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        v43 := GenS1{F0: -21, F1: "", F2: 812}
        _ = fi42
        _ = v43
    }
    var v44 uint32 = 1
    v45 := GenS6{F0: genFn0(-1, genFn1((v1 / (((v1 % ((127) | 1))) | 1))), v1), F1: 0}
    if v3 {
        delete(v2, v44)
    }
    var v46 uint16 = 547
    v47 := GenS3{F0: 26}
    v2[uint32(v4)] = (!v3)
    v48 := GenS7{F0: -1, F1: (v46 % ((271) | 1))}
    v44 = 744
    v49 := GenS2{F0: "hello"}
    v1 = (v1 << ((42) & 63))
    var v50 uint = 337
    switch (("\n\t" + (("" + "") + (("\n\t" + "a") + ("foo bar" + "a")))) + ("\n\t" + "foo bar")) {
    case "":
        fmt.Println(v50)
    case "\n\t":
        v44 += (^(v44))
    case "a":
        var v51 float64 = 1.5
        _ = v51
    default:
        var v52 int = 127
        _ = v52
    }
    v53 := []int{-308, -2147483648, -1}
    v3 = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v53
}
