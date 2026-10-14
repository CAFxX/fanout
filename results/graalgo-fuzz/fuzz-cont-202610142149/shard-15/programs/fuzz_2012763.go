package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int }
func genFn0(p0 int32, p1 int8, p2 int16) int32 {
    fmt.Println(p1, p2, p1)
    v5 := GenS1{F0: 826, F1: 829}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}
func genFn1() uint32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        if (!(9223372036854775807 <= 9223372036854775807)) {
            v7 := GenS1{F0: 76, F1: -1}
            for fi8 := 0; fi8 < 18; fi8++ {
                fmt.Println(fi6, fi8)
                v9 := []int{-1, (fi8 & 42), (fi8 ^ 2147483647)}
                _ = fi8
                _ = v9
            }
            var v10 int32 = -607
            _ = v10
            _ = v7
        } else {
            var v11 int32 = -2147483648
            _ = v11
        }
        fmt.Println(fi6)
        switch "a" {
        case "":
            v12 := GenS1{F0: 977, F1: 1}
            _ = v12
        case "foo bar":
            v13 := []int{-909, fi6}
            _ = v13
        }
        _ = fi6
    }
    var v14 int8 = 127
    fmt.Println(v14)
    _ = v14
    return 2147483648
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]string)
    var v3 uint32 = 927
    var v4 int = 770
    var v15 string = "foo bar"
    v3 += genFn1()
    v1 = (-(v1))
    v15 = "hello"
    v1 -= (v1 | (v1 + (-(v1))))
    fmt.Println(v3)
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = fi16
        _ = v17
    }
    if (!true) {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int16 = (v1 | (v1 >> ((65535) & 63)))
            if ((true || (3.14159 >= 0.0)) && false) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 bool = true
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS2{F0: 2}
                    switch -1 {
                    case -2147483648:
                        v2[fi22] = (v15 + "foo bar")
                    case 128:
                        var v24 bool = false
                        _ = v24
                    case 127:
                        var v25 float64 = float64(3.14159)
                        _ = v25
                    }
                    v15 = ((v15 + v15) + v15)
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi18
            _ = v19
        }
        v26 := GenS3{F0: 1, F1: (-(v4))}
        switch v4 {
        case 32768:
            v4 = v4
        case -2147483648:
            v15 = ((v15 + (v15 + (v15 + v15))) + v15)
        case 127:
            if (55 <= 255) {
                var v27 uint64 = 4294967295
                v28 := []int{v4}
                _ = v27
                _ = v28
            }
        default:
            var v29 float32 = 0.1
            _ = v29
        }
        _ = v26
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        if true {
            delete(v2, 1)
            v1 = (v1 + (v1 * (v1 | 342)))
            for fi31 := 0; fi31 < 4; fi31++ {
                if false {
                    if (9223372036854775807 == 9223372036854775807) {
                        v3 -= (v3 & (v3 & genFn1()))
                    } else {
                        v32 := v2[fi31]
                        switch v4 {
                        case -820:
                            var v33 int64 = 788
                            _ = v33
                        }
                        _ = v32
                    }
                }
                v34 := []int{(fi31 | fi30)}
                v35 := GenS2{F0: 934}
                _ = fi31
                _ = v34
                _ = v35
            }
        }
        var v36 int = (fi30 / (((v4 | fi30)) | 1))
        _ = fi30
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        fmt.Println(fi37, v4, fi37)
        v15 = v15
        for fi38 := 0; fi38 < 9; fi38++ {
            if true {
                v39 := GenS3{F0: 645, F1: v4}
                v1 -= v1
                fmt.Println(fi38, v3, fi37)
                _ = v39
            }
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println(v15)
    var v40 int = (v4 % (((v4 * (v4 >> ((259) & 63)))) | 1))
    var v41 int32 = genFn0(2147483647, 47, 1)
    for fi42 := 0; fi42 < 11; fi42++ {
        if ((!true) || (9223372036854775807 >= 253)) {
            var v43 string = ((v15 + "") + "\n\t")
            var v44 string = v15
            _ = v43
            _ = v44
        } else {
            switch ("\n\t" + "a") {
            case "":
                var v45 bool = true
                _ = v45
            case "hello":
                v46 := GenS2{F0: 673}
                _ = v46
            }
        }
        _ = fi42
    }
    v40 -= v40
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
