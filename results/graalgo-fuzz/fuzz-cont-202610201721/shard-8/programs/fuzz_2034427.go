package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 int64; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 string = "foo bar"
    v5 := []int{4294967295}
    v4 = ("hello" + "")
    var v6 float64 = -2.25
    v7 := v3[452]
    v8 := GenS1{F0: 134, F1: (-343 != 1000)}
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 2, F1: ((!true) && true)}
        _ = fi9
        _ = v10
    }
    v4 = v4
    if (!(3.14159 < 0.0)) {
        v5 = append(v5, v7)
    } else {
        switch v7 {
        case 1:
            v7 += v7
        case 0:
            var v11 int32 = -2147483648
            _ = v11
        default:
            v4 = ((v4 + "hello") + v4)
        }
        if (!false) {
            var v12 uint64 = 42
            if true {
                v13 := GenS2{F0: 0, F1: -551, F2: 782}
                _ = v13
            } else {
                if (((("foo bar" + v4) + (("" + v4) + ("foo bar" + (("foo bar" + "") + v4)))) + (("a" + v4) + (v4 + ""))) >= v4) {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v12 += (v12 * v12)
                        var v15 bool = false
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(v4, v7, v6)
                    v12 = 18446744073709551615
                } else {
                    v16 := GenS3{F0: 1}
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v18 := GenS4{F0: -1, F1: 295}
                        if (true && false) {
                            v19 := GenS2{F0: 42, F1: -1, F2: 498}
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                if (!((!(int(v12) < v7)) || (!(true || true)))) {
                    switch "" {
                    case "":
                        v20 := GenS3{F0: 0}
                        _ = v20
                    case "\n\t":
                        fmt.Println(v4, v4)
                    default:
                        var v21 uint8 = 0
                        _ = v21
                    }
                } else {
                    fmt.Println(v7)
                }
                v7 += v7
            }
            _ = v12
        } else {
            v7 += v7
        }
    }
    if false {
        if false {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 8; fi23++ {
                    switch fi22 {
                    case 127:
                        v24 := v3[v7]
                        _ = v24
                    default:
                        if false {
                            v6 -= v6
                        }
                    }
                    _ = fi23
                }
                for fi25 := 0; fi25 < 20; fi25++ {
                    v26 := GenS1{F0: 2, F1: false}
                    v4 = "foo bar"
                    _ = len(v5)
                    _ = fi25
                    _ = v26
                }
                v7 = (v7 & v7)
                _ = fi22
            }
            var v27 string = v4
            _ = v27
        } else {
            if len(v5) > 0 {
                v5[0] = v7
            }
            v6 -= v6
        }
        v28 := GenS5{F0: 4294967295}
        fmt.Println(v6)
        _ = v28
    } else {
        var v29 uint32 = 42
        _ = v29
    }
    var v30 uint64 = 4294967295
    var v31 int64 = 42
    v32 := GenS2{F0: -128, F1: v31, F2: 1}
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(v6)
        _ = fi33
    }
    v31 += (v31 ^ (v31 >> ((548) & 63)))
    var v34 int64 = 897
    switch v7 {
    case 956:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case -9223372036854775808:
        var v35 uint16 = 42
        _ = v35
    }
    if (2147483647 <= 42) {
        for fi36 := 0; fi36 < 17; fi36++ {
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 string = v4
                v39 := GenS2{F0: 0, F1: int64(v30), F2: 17}
                _ = fi37
                _ = v38
                _ = v39
            }
            if false {
                var v40 int8 = 101
                if len(v5) > 0 {
                    v5[0] = (v7 + -1)
                }
                v41 := GenS5{F0: 2147483647}
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 4; fi42++ {
                v5 = append(v5, fi42)
                _ = len(v1)
                v43 := GenS6{F0: v7, F1: (int64(v30) <= 9223372036854775807)}
                _ = fi42
                _ = v43
            }
            _ = fi36
        }
        if false {
            _ = len(v3)
            var v44 uint32 = 359
            v45 := GenS1{F0: 127, F1: false}
            _ = v44
            _ = v45
        }
        if true {
            var v46 uint64 = v30
            for fi47 := 0; fi47 < 10; fi47++ {
                v31 -= 0
                for fi48 := 0; fi48 < 4; fi48++ {
                    v49 := GenS5{F0: 0}
                    v30 = v30
                    _ = fi48
                    _ = v49
                }
                fmt.Println(v4, v34)
                _ = fi47
            }
            if len(v5) > 0 {
                v5[0] = (v7 % (((v7 >> ((uint(v30)) & 63))) | 1))
            }
            _ = v46
        } else {
            var v50 float64 = v6
            var v51 int = (-(v7))
            _ = v50
            _ = v51
        }
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
