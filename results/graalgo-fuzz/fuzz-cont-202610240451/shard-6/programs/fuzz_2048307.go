package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int32; F1 string; F2 uint32 }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 uint16) int8 {
    var v6 uint16 = p0
    var v7 int8 = -1
    fmt.Println(v6)
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 32767}
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v7)
    _ = p0
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 uint, p1 int8, p2 uint64) uint16 {
    if false {
        p1 = (^(p1))
        fmt.Println(p0)
    }
    v10 := []int{9223372036854775807, 272, 127, 4294967295}
    fmt.Println(p1, p1, p0)
    switch 42 {
    case 9223372036854775807:
        fmt.Println(p1, p0)
    case 641:
        if true {
            var v11 int8 = 73
            p1 = 0
            switch 9223372036854775807 {
            case 42:
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := GenS2{F0: 9223372036854775807, F1: true, F2: 42}
                    _ = fi12
                    _ = v13
                }
            default:
                var v14 string = ""
                _ = v14
            }
            _ = v11
        } else {
            var v15 uint64 = p2
            _ = v15
        }
    case 0:
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(p2, p1)
            _ = fi16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 255
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 bool = ((v1 != v1) && (420 == 477))
    var v4 int32 = -1
    var v5 float64 = 0.1
    if ((0.1 * -2.25) < (0.0 + (-2.25 / 1e10))) {
        v17 := GenS2{F0: -1, F1: v3, F2: -32768}
        _ = v17
    }
    var v18 string = ("hello" + ("hello" + "hello"))
    var v19 uint16 = 65535
    var v20 uint16 = (v19 & genFn1(255, genFn0((v19 << ((1000) & 63))), 42))
    if (164 < 42) {
        if (v3 || (true || (v20 != (^(v20))))) {
            v21 := GenS1{F0: 1}
            if false {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v3 = v3
                    v23 := []int{(fi22 / ((fi22) | 1)), (fi22 - 1), fi22, 32768}
                    _ = fi22
                    _ = v23
                }
                if (38 >= 4294967295) {
                    v24 := GenS2{F0: 42, F1: (false && true), F2: 1000}
                    switch v18 {
                    case "hello":
                        v1 = (v1 - 122)
                    case "a":
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v26 := []int{(fi25 >> ((0) & 63)), (fi25 * (fi25 % ((fi25) | 1))), -940}
                            _ = fi25
                            _ = v26
                        }
                    case "foo bar":
                        fmt.Println(v3)
                    }
                    switch -323 {
                    case 876:
                        v19 = 255
                    case 9223372036854775807:
                        delete(v2, (("hello" + v18) + ""))
                    case 858:
                        var v27 int16 = -32768
                        _ = v27
                    }
                    _ = v24
                } else {
                    v28 := v2[v18]
                    _ = v28
                }
                v19 = genFn1(2, 53, 2)
            }
            if (v3 || false) {
                switch 128 {
                case 128:
                    switch 2147483647 {
                    case 263:
                        v1 -= (v1 - (v1 * (v1 / ((v1) | 1))))
                    case -58:
                        v29 := GenS3{F0: v1}
                        _ = v29
                    }
                case 42:
                    if (-1 == -195) {
                        v30 := GenS4{F0: 97, F1: v3}
                        _ = v30
                    } else {
                        var v31 uint32 = 0
                        v32 := []int{797, -9223372036854775808, 42, 0}
                        _ = v31
                        _ = v32
                    }
                case 0:
                    if v3 {
                        v33 := GenS5{F0: 4294967295}
                        _ = v33
                    } else {
                        switch (v18 + v18) {
                        case "hello":
                            v5 += ((v5 + 0.0) * v5)
                        case "foo bar":
                            var v34 float64 = (((v5 - 1e10) + 0.1) / v5)
                            _ = v34
                        case "":
                            var v35 uint32 = 907
                            _ = v35
                        }
                        v5 += (((0.0 / 1.0) - 3.14159) - v5)
                    }
                default:
                    v3 = v3
                }
                for fi36 := 0; fi36 < 12; fi36++ {
                    for fi37 := 0; fi37 < 7; fi37++ {
                        fmt.Println(v19, v19, v1)
                        _ = len(v2)
                        v38 := GenS6{F0: v20, F1: v18}
                        _ = fi37
                        _ = v38
                    }
                    _ = fi36
                }
                v39 := GenS7{F0: -1, F1: "\n\t", F2: 211}
                _ = v39
            }
            _ = v21
        } else {
            if v3 {
                var v40 int8 = 119
                _ = v40
            }
        }
        fmt.Println(v20, v20, v1)
        var v41 string = v18
        _ = v41
    } else {
        if (true || v3) {
            fmt.Println(v4, v5, v18)
            v3 = false
        }
    }
    v19 = v20
    fmt.Println(v1, v5)
    var v42 int8 = 83
    var v43 uint = 255
    var v44 uint64 = uint64(v43)
    fmt.Println(v5)
    switch (("" + "\n\t") + v18) {
    case "a":
        fmt.Println(v42, v3, v1)
    case "\n\t":
        switch "hello" {
        case "hello":
            v42 -= 100
        case "foo bar":
            var v45 bool = (2 < v44)
            _ = v45
        default:
            v2[""] = v3
        }
    }
    v18 = v18
    for fi46 := 0; fi46 < 1; fi46++ {
        switch v18 {
        case "":
            if true {
                if (32767 >= 32767) {
                    v47 := GenS8{F0: v44, F1: 248, F2: (v43 & v43)}
                    _ = v47
                } else {
                    var v48 float64 = 1e10
                    _ = v48
                }
                v49 := []int{fi46}
                _ = v49
            }
        case "hello":
            for fi50 := 0; fi50 < 8; fi50++ {
                v42 -= v42
                if ((^(v1)) <= 2) {
                    var v51 int32 = (-(v4))
                    _ = v51
                }
                var v52 float32 = 0.1
                var v53 int = fi46
                switch (fi46 & 2147483647) {
                case 406:
                    var v54 uint8 = (v1 - v1)
                    _ = v54
                case -1:
                    var v55 int16 = -1
                    _ = v55
                }
                v2[((("hello" + v18) + v18) + "")] = (v43 <= 2)
                fmt.Println(v5, v42)
                _ = fi50
                _ = v52
                _ = v53
            }
        }
        _ = fi46
    }
    if (v19 == (^(v20))) {
        v4 += 2147483647
    }
    var v56 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v56
}
