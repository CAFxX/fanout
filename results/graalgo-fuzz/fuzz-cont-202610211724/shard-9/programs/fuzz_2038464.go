package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int8; F1 int32; F2 int32 }
type GenS7 struct { F0 uint8; F1 int }
func genFn0() int8 {
    var v7 float32 = (0.1 / 1e10)
    fmt.Println(v7)
    v8 := []int{-9223372036854775808, -595}
    var v9 float32 = 0.1
    v9 -= v9
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int8, p1 int32, p2 uint8) int64 {
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(p0, p2, p1)
        v11 := GenS1{F0: 0}
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 677
}

func main() {
    v1 := []string{"hello", ""}
    var v2 int16 = -870
    var v3 int16 = (-(v2))
    v4 := []string{}
    v5 := make(map[string]string)
    var v6 int16 = v3
    if (!(true && false)) {
        var v13 int8 = 100
        var v14 float32 = (0.0 + 1e10)
        var v15 uint8 = 1
        _ = v13
        _ = v14
        _ = v15
    } else {
        v5[(("foo bar" + "hello") + ("\n\t" + ""))] = "a"
        var v16 uint32 = 748
        _ = v16
    }
    var v17 int = 32767
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (545 == 255) {
        v18 := GenS2{F0: 42, F1: 2147483647, F2: 1}
        var v19 uint16 = 2
        _ = v18
        _ = v19
    }
    var v20 int64 = 42
    v6 -= (-(v3))
    v20 -= -9223372036854775808
    for fi21 := 0; fi21 < 11; fi21++ {
        v4 = append(v4, "foo bar")
        switch (("hello" + ("a" + "a")) + "a") {
        case "":
            if (!true) {
                v2 = -305
                fmt.Println(v3)
                for fi22 := 0; fi22 < 5; fi22++ {
                    v23 := GenS3{F0: 65535}
                    _ = fi22
                    _ = v23
                }
            } else {
                switch fi21 {
                case -641:
                    fmt.Println(v17, v17)
                }
                v24 := GenS4{F0: ("foo bar" + "a")}
                _ = v24
            }
        case "a":
            v25 := GenS2{F0: 255, F1: (fi21 >> ((255) & 63)), F2: 1}
            _ = v25
        case "\n\t":
            v5[""] = "foo bar"
        default:
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 int8 = 42
                if (167 != 1) {
                    if (genFn0() < 100) {
                        var v28 int16 = (-(v6))
                        var v29 uint = 42
                        var v30 int8 = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        v27 = 1
                        if ("a" == "a") {
                            v20 += (v20 * (v20 / ((v20) | 1)))
                            v31 := GenS3{F0: 0}
                            var v32 uint16 = 255
                            _ = v31
                            _ = v32
                        } else {
                            v33 := GenS3{F0: 2}
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v3)
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v33
                        }
                        fmt.Println(v2, fi26, v17)
                    }
                    v27 += (v27 & 0)
                } else {
                    _ = len(v4)
                    switch -715 {
                    case 306:
                        v34 := GenS2{F0: 127, F1: fi21, F2: -2147483648}
                        _ = v34
                    default:
                        v35 := GenS5{F0: 18446744073709551615, F1: 720}
                        _ = v35
                    }
                    v36 := GenS6{F0: genFn0(), F1: 739, F2: 39}
                    _ = v36
                }
                fmt.Println(fi26)
                v37 := GenS4{F0: "hello"}
                if ((v20 % (((v20 >> ((433) & 63))) | 1)) != v20) {
                    var v38 uint32 = 4294967295
                    if ("" == "foo bar") {
                        var v39 string = "a"
                        var v40 uint64 = 175
                        v3 -= (v3 * (v2 / ((v6) | 1)))
                        _ = v39
                        _ = v40
                    }
                    v41 := GenS6{F0: (v27 % ((114) | 1)), F1: -845, F2: 42}
                    _ = v38
                    _ = v41
                } else {
                    v17 += (v17 + (-(fi21)))
                    _ = len(v1)
                    v42 := GenS7{F0: 0, F1: (v17 + fi26)}
                    _ = v42
                }
                _ = fi26
                _ = v27
                _ = v37
            }
        }
        _ = fi21
    }
    v3 += 42
    if ((1 <= 255) || true) {
        var v43 int = (v17 % ((v17) | 1))
        _ = v43
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (32767 != 1) {
        fmt.Println(v17, v6)
    } else {
        v2 -= (v3 % ((v3) | 1))
        if len(v4) > 0 {
            v4[0] = ("foo bar" + "hello")
        }
    }
    fmt.Println(v17)
    var v44 int = v17
    var v45 uint8 = 12
    var v46 int16 = -1
    switch (v17 % ((v17) | 1)) {
    case 2147483647:
        for fi47 := 0; fi47 < 19; fi47++ {
            v46 = (v2 + v2)
            _ = fi47
        }
    case -49:
        v45 += 127
    default:
        v48 := GenS3{F0: 682}
        _ = v48
    }
    v45 = (v45 & (^(v45)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
