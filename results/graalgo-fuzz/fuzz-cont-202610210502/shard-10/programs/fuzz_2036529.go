package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 uint }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 int16, p1 int16) uint64 {
    v7 := []int{}
    switch ("" + "hello") {
    case "\n\t":
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 75, F1: 447, F2: 622}
            _ = fi8
            _ = v9
        }
    }
    p0 = 0
    switch "a" {
    case "":
        p0 = -1
    case "foo bar":
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int, p1 bool) int {
    p0 += p0
    p1 = (((3.14159 / (3.14159 + (1e10 - 1e10))) != 0.1) && (p1 && p1))
    p1 = (32767 != 2)
    p0 += (^(p0))
    switch "\n\t" {
    case "":
        fmt.Println(p0, p0)
    case "foo bar":
        fmt.Println(p1)
    default:
        if (693 < -438) {
            switch p0 {
            case 32768:
                var v10 uint16 = 0
                _ = v10
            }
            switch 398 {
            case -9223372036854775808:
                v11 := GenS2{F0: true, F1: -581, F2: 65535}
                _ = v11
            default:
                if true {
                    p0 = (p0 - p0)
                    fmt.Println(p0)
                }
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    p1 = (p1 && (!p1))
                    if (true || (100 > -112)) {
                        switch "" {
                        case "\n\t":
                            switch "hello" {
                            case "hello":
                                v14 := GenS1{F0: -29, F1: 4294967295, F2: 52}
                                switch "a" {
                                case "foo bar":
                                    p1 = p1
                                default:
                                    v15 := []int{(p0 ^ fi13), 42, (p0 | 9223372036854775807), -615}
                                    _ = v15
                                }
                                var v16 int32 = -1
                                _ = v14
                                _ = v16
                            case "foo bar":
                                v17 := []int{(p0 - (fi13 + (p0 & (fi13 << ((656) & 63))))), (fi12 << ((960) & 63)), fi12, 0}
                                _ = v17
                            case "":
                                v18 := GenS1{F0: 42, F1: 787, F2: 2}
                                _ = v18
                            }
                        case "a":
                            v19 := GenS3{F0: (("a" + "foo bar") + ("foo bar" + ("\n\t" + (("foo bar" + "foo bar") + "foo bar")))), F1: (p1 || ("a" == "")), F2: 1}
                            _ = v19
                        case "hello":
                            v20 := GenS4{F0: 209, F1: ((385 >= 42) && true)}
                            _ = v20
                        default:
                            fmt.Println(fi13, p0)
                        }
                        p0 += p0
                        p0 = p0
                    }
                    v21 := GenS4{F0: 0, F1: ((4294967295 >= 18446744073709551615) || (33 < 2))}
                    _ = fi13
                    _ = v21
                }
                if false {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        if ((fi22 / (((p0 / (((-(fi12))) | 1))) | 1)) > (^(fi22))) {
                            fmt.Println(p1, fi12, fi22)
                            var v23 uint64 = 612
                            _ = v23
                        }
                        v24 := []int{p0, -512, p0}
                        var v25 uint8 = 2
                        var v26 int = p0
                        fmt.Println(fi22, fi12, fi22)
                        v27 := GenS1{F0: 100, F1: 4294967295, F2: 306}
                        _ = fi22
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                } else {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 uint = 859
                        fmt.Println(p0, fi28)
                        fmt.Println(v29)
                        _ = fi28
                        _ = v29
                    }
                    var v30 int64 = -9223372036854775808
                    _ = v30
                }
                _ = fi12
            }
        } else {
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 uint64 = genFn0(-32768, 0)
                if ((!(!p1)) && true) {
                    var v33 int16 = 42
                    fmt.Println(v33, v32)
                    _ = v33
                } else {
                    v34 := []int{fi31, (^(fi31)), 32767}
                    v32 = genFn0(42, -32768)
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
            v35 := GenS2{F0: (65535 == 831), F1: 32767, F2: 4294967295}
            if (!p1) {
                v36 := []int{-2147483648, (p0 - p0)}
                v37 := GenS4{F0: 42, F1: (!p1)}
                _ = v36
                _ = v37
            } else {
                for fi38 := 0; fi38 < 5; fi38++ {
                    v39 := []int{(fi38 ^ (p0 & p0)), fi38, (^(p0))}
                    _ = fi38
                    _ = v39
                }
                for fi40 := 0; fi40 < 11; fi40++ {
                    p0 = (p0 - (^(fi40)))
                    var v41 uint64 = genFn0(1, 364)
                    _ = fi40
                    _ = v41
                }
            }
            _ = v35
        }
    }
    _ = p0
    _ = p1
    return (p0 + (p0 - (p0 << ((255) & 63))))
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]bool)
    v3 := []string{"hello", "hello", ("a" + ("" + ("hello" + "a"))), ((((("foo bar" + "\n\t") + "") + "hello") + "a") + "")}
    var v4 int64 = 597
    v5 := []uint32{(v1 << ((42) & 63)), v1}
    v6 := make(map[string]int)
    if true {
        switch (("" + "\n\t") + ((("foo bar" + "hello") + "a") + ("\n\t" + ("\n\t" + "hello")))) {
        case "\n\t":
            v42 := GenS1{F0: 42, F1: 102, F2: 1}
            _ = v42
        default:
            var v43 uint16 = 1
            _ = v43
        }
        fmt.Println(v1, v1)
        var v44 float64 = -2.25
        _ = v44
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        v4 += v4
        var v46 int16 = -404
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        var v48 int64 = 368
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 += v1
        _ = fi47
        _ = v48
    }
    v6[""] = 58
    var v49 bool = false
    var v50 uint64 = 319
    if (!((genFn0(447, 32767) > (^(v50))) || ((v4 / (((-(v4))) | 1)) > v4))) {
        var v51 int8 = -1
        v52 := GenS2{F0: v49, F1: 1000, F2: uint32(v50)}
        _ = v51
        _ = v52
    }
    v53 := GenS3{F0: "hello", F1: ((("hello" + "a") + ("hello" + (("\n\t" + "hello") + ("" + "foo bar")))) > ""), F2: 849}
    v4 += (v4 >> ((251) & 63))
    fmt.Println(v1, v50, v49)
    for fi54 := 0; fi54 < 7; fi54++ {
        if len(v5) > 0 {
            v5[0] = 42
        }
        _ = fi54
    }
    v49 = false
    if (464 <= -1) {
        for fi55 := 0; fi55 < 20; fi55++ {
            v56 := v6[("\n\t" + ("" + "foo bar"))]
            _ = fi55
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v6
}
