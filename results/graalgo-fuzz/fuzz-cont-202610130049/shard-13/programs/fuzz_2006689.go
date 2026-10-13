package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 string }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32, p1 int32, p2 int64) uint8 {
    if (!false) {
        fmt.Println(p1, p0)
        var v4 int = 127
        _ = v4
    } else {
        v5 := []int{4294967295, 258, int(p1), 4294967295}
        v6 := GenS1{F0: 42, F1: -2147483648, F2: 1}
        _ = v5
        _ = v6
    }
    var v7 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (v7 | uint8(p0))
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    switch 42 {
    case 0:
        var v8 uint64 = 42
        _ = v8
    case 9223372036854775807:
        v1 = "foo bar"
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 = v1
        _ = fi9
    }
    var v10 float64 = ((1.5 * 3.14159) + (0.0 / (0.1 / (1e10 - 3.14159))))
    switch 0 {
    case 584:
        switch 42 {
        case 0:
            if (("a" > "\n\t") && ((-193 < 866) && (!((true || true) && false)))) {
                var v11 float32 = 0.0
                v10 = float64((-2.25 * (v11 - v11)))
                v11 = 3.14159
                _ = v11
            }
        case -9223372036854775808:
            if false {
                if false {
                    v12 := GenS1{F0: 454, F1: -1, F2: 91}
                    _ = v12
                }
            } else {
                delete(v3, 2)
                v13 := GenS2{F0: -128, F1: 1}
                var v14 int16 = -1
                _ = v13
                _ = v14
            }
        case 4294967295:
            if true {
                fmt.Println(v10)
            } else {
                var v15 string = v1
                _ = v15
            }
        default:
            var v16 uint64 = 420
            _ = v16
        }
    case 1:
        v17 := []int{}
        _ = v17
    }
    var v18 int16 = 1000
    v1 = ("a" + v1)
    var v19 uint = 0
    if (1 > 809) {
        v10 = -2.25
        if true {
            var v20 uint8 = 127
            _ = v20
        } else {
            var v21 int64 = -9223372036854775808
            v18 -= (-(v18))
            v22 := GenS3{F0: (^(v19)), F1: 55, F2: 2}
            _ = v21
            _ = v22
        }
        var v23 uint = v19
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int32 = int32(fi24)
            var v26 int64 = 484
            _ = fi24
            _ = v25
            _ = v26
        }
        v18 += (v18 * v18)
    }
    v27 := GenS1{F0: 82, F1: 42, F2: 4294967295}
    v2[v1] = (true || (!true))
    _ = len(v3)
    for fi28 := 0; fi28 < 10; fi28++ {
        if false {
            if false {
                var v29 uint16 = 65535
                switch v1 {
                case "":
                    v30 := []int{fi28, fi28, (fi28 >> ((v19) & 63))}
                    _ = v30
                default:
                    if (!(((!(((fi28 * fi28) > (fi28 - (-(fi28)))) && (!false))) && true) || true)) {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            var v32 uint8 = 1
                            var v33 int = fi31
                            var v34 float32 = (0.0 + 1.5)
                            if (!(v34 == (float32(float64(((v34 / (3.14159 * 3.14159)) * 0.1))) * v34))) {
                                var v35 float64 = 1e10
                                _ = v35
                            } else {
                                v36 := GenS2{F0: -114, F1: 264}
                                var v37 int64 = 1
                                var v38 uint64 = 4294967295
                                v39 := GenS3{F0: (v19 >> ((v19) & 63)), F1: v29, F2: 42}
                                var v40 uint64 = v38
                                v41 := v2["a"]
                                var v42 string = ((v1 + "\n\t") + v1)
                                _ = v36
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            v43 := GenS2{F0: 20, F1: 42}
                            v34 += v34
                            var v44 uint32 = 204
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v43
                            _ = v44
                        }
                        v29 = v29
                        v45 := GenS2{F0: 87, F1: 42}
                        _ = v45
                    } else {
                        v46 := GenS4{F0: 100, F1: 0, F2: ""}
                        v10 -= float64(3.14159)
                        v10 += 0.0
                        _ = v46
                    }
                }
                _ = v29
            } else {
                if (((fi28 << ((v19) & 63)) != (fi28 << ((v19) & 63))) || (genFn0(42, -1, -1) > 255)) {
                    if ((((true || ((true && (fi28 > (fi28 >> ((421) & 63)))) && true)) && (2147483647 <= -2147483648)) || (5 != 16)) && true) {
                        var v47 float32 = 3.14159
                        v48 := GenS3{F0: v19, F1: 42, F2: 973}
                        fmt.Println(v1)
                        _ = v47
                        _ = v48
                    } else {
                        v19 = v19
                        _ = len(v3)
                        v1 = ("foo bar" + v1)
                    }
                    delete(v2, v1)
                    v49 := v3[2147483648]
                    _ = v49
                } else {
                    var v50 uint64 = 42
                    _ = v50
                }
            }
            var v51 int16 = (v18 - (v18 / ((v18) | 1)))
            _ = v51
        }
        if true {
            v52 := v2[(((v1 + ((v1 + v1) + "foo bar")) + v1) + (v1 + ("" + (v1 + (("a" + v1) + v1)))))]
            v53 := GenS5{F0: 220, F1: -1}
            v10 -= (float64(-2.25) + v10)
            _ = v52
            _ = v53
        }
        _ = fi28
    }
    if ((^(v19)) <= 1) {
        var v54 uint = v19
        fmt.Println(v18, v10, v18)
        fmt.Println(v1, v19, v19)
        _ = v54
    }
    var v55 int16 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v55
}
