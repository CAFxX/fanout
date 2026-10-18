package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
func genFn0() int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        if false {
            var v5 float64 = (3.14159 - -2.25)
            v5 += v5
            v6 := GenS1{F0: 1000, F1: (!true)}
            _ = v5
            _ = v6
        }
        var v7 int16 = 170
        _ = fi4
        _ = v7
    }
    if (42 <= 0) {
        var v8 int8 = -19
        _ = v8
    }
    return -65
}
func genFn1(p0 uint, p1 int16) int64 {
    if true {
        if false {
            v9 := GenS1{F0: p1, F1: (-16 > 32768)}
            var v10 uint16 = 158
            _ = v10
            _ = v9
        }
        var v11 uint16 = 65535
        v11 = (v11 >> ((p0) & 63))
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        if true {
            p1 = (p1 / ((p1) | 1))
        }
        var v13 int = fi12
        if false {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := []int{-8, v13, v13, fi14}
                if true {
                    if true {
                        v16 := GenS2{F0: 18446744073709551615}
                        v13 = v13
                        _ = v16
                    } else {
                        p1 = 0
                    }
                } else {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    switch ("foo bar" + "foo bar") {
                    case "":
                        fmt.Println(fi14)
                    }
                }
                _ = fi14
                _ = v15
            }
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int16 = -32768
                _ = fi17
                _ = v18
            }
        } else {
            p1 = (p1 >> ((255) & 63))
            v19 := []int{(v13 ^ fi12), v13, -1, v13}
            _ = v19
        }
        _ = fi12
        _ = v13
    }
    p1 += p1
    v20 := []int{42, 9223372036854775807}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v20
    return 0
}

func main() {
    v1 := []int{}
    var v2 uint = 65535
    v3 := make(map[string]int)
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int16 = -468
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: -1, F1: v2}
        _ = fi21
        _ = v24
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v25 int32 = -659
    var v26 uint32 = 346
    var v27 uint64 = 1
    if (genFn1(v2, 1) == genFn1(v2, 0)) {
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 float64 = 0.1
            v26 -= (v26 - 2147483648)
            v30 := v3["foo bar"]
            _ = fi28
            _ = v29
            _ = v30
        }
        if (869 < 32767) {
            var v31 int8 = 0
            v32 := GenS1{F0: 0, F1: false}
            v33 := GenS4{F0: (128 > 1)}
            _ = v31
            _ = v32
            _ = v33
        } else {
            v27 += v27
            delete(v3, "")
        }
    } else {
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 bool = false
            delete(v3, "\n\t")
            v36 := GenS4{F0: ((^(v27)) <= v27)}
            _ = fi34
            _ = v35
            _ = v36
        }
        var v37 uint = v2
        v38 := GenS2{F0: (^(v2))}
        _ = v37
        _ = v38
    }
    switch 2147483647 {
    case 4294967295:
        if (!false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v39 := GenS1{F0: -600, F1: (1 != 39)}
            _ = v39
        } else {
            v40 := GenS1{F0: -491, F1: ((9223372036854775807 == 42) && true)}
            var v41 int64 = 9223372036854775807
            _ = v40
            _ = v41
        }
    case 32767:
        for fi42 := 0; fi42 < 7; fi42++ {
            var v43 uint32 = (v26 >> (((v2 ^ v2)) & 63))
            for fi44 := 0; fi44 < 10; fi44++ {
                v45 := GenS5{F0: v2}
                _ = fi44
                _ = v45
            }
            _ = fi42
            _ = v43
        }
    case 588:
        var v46 bool = (!((!true) || (0 == 0)))
        _ = v46
    default:
        for fi47 := 0; fi47 < 7; fi47++ {
            v25 = v25
            if false {
                fmt.Println(fi47)
                var v48 bool = false
                _ = v48
            } else {
                v27 = (^(v27))
                if (-128 < -1) {
                    v26 += v26
                    if true {
                        var v49 int8 = 127
                        delete(v3, ("" + "hello"))
                        v27 += (^(v27))
                        _ = v49
                    }
                }
            }
            switch 4294967295 {
            case -696:
                v1 = append(v1, 2147483647)
            case -9223372036854775808:
                v3["\n\t"] = (fi47 * (fi47 % (((fi47 * (fi47 % (((fi47 >> (((v2 ^ (v2 - 4294967295))) & 63))) | 1)))) | 1)))
            default:
                for fi50 := 0; fi50 < 10; fi50++ {
                    var v51 string = ("" + "\n\t")
                    var v52 uint16 = 479
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            }
            _ = fi47
        }
    }
    for fi53 := 0; fi53 < 7; fi53++ {
        v1 = append(v1, fi53)
        for fi54 := 0; fi54 < 5; fi54++ {
            v3["\n\t"] = 32768
            _ = fi54
        }
        v25 = (v25 ^ v25)
        _ = fi53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
