package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int; F1 uint32; F2 bool }
type GenS7 struct { F0 int32; F1 int8 }
func genFn0(p0 int32) bool {
    var v5 float32 = 1.5
    var v6 float32 = float32((float64(3.14159) + 1e10))
    _ = p0
    _ = v5
    _ = v6
    return (!true)
}
func genFn1(p0 uint16, p1 int64) uint {
    v7 := []int{4294967295, 634, 32767}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v7
    return 65535
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 752
    var v3 uint8 = 176
    v4 := make(map[int]string)
    if ((true && (2147483648 == 2147483647)) || ("\n\t" != ("hello" + ("hello" + ("" + "foo bar"))))) {
        v8 := []int{}
        _ = v8
    } else {
        v9 := GenS1{F0: -107, F1: ((!(v1 < (v1 >> ((42) & 63)))) || (42 == -2147483648)), F2: 317}
        var v10 bool = (v1 <= (v1 | (v1 >> ((968) & 63))))
        if ((!true) && (1.5 < (1.5 + 0.1))) {
            var v11 int16 = 920
            if v10 {
                v2 = v2
                v12 := GenS1{F0: 1, F1: false, F2: v2}
                _ = v12
            }
            _ = v11
        } else {
            var v13 uint8 = v3
            for fi14 := 0; fi14 < 16; fi14++ {
                v1 = 42
                if v10 {
                    var v15 int16 = 42
                    var v16 int16 = (v15 - v15)
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
            if true {
                for fi17 := 0; fi17 < 5; fi17++ {
                    v1 = v1
                    _ = fi17
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    switch int(v3) {
                    case 0:
                        var v19 uint = 0
                        _ = v19
                    case 85:
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v21 := []int{-2147483648}
                            v22 := GenS1{F0: -1, F1: (!(v10 || (!v10))), F2: int64(fi20)}
                            v10 = (false || ((!v10) && v10))
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    v23 := []int{fi18, fi18, fi18}
                    _ = fi18
                    _ = v23
                }
            } else {
                switch "\n\t" {
                case "hello":
                    v24 := GenS1{F0: 42, F1: false, F2: (v2 ^ v2)}
                    _ = v24
                }
                v10 = v10
                v3 -= v13
            }
            _ = v13
        }
        _ = v10
        _ = v9
    }
    var v25 int32 = 0
    v26 := v4[-1]
    v1 = (v1 >> ((0) & 63))
    if (v25 >= (v25 ^ (v25 + v25))) {
        v2 -= 443
    }
    fmt.Println(v25)
    if false {
        v27 := GenS1{F0: -1, F1: (v26 == "foo bar"), F2: (-(v2))}
        _ = v27
    } else {
        delete(v4, int(v2))
        var v28 float32 = ((-2.25 + -2.25) + 0.1)
        if ((v25 * (-(v25))) == (v25 & (v25 / ((v25) | 1)))) {
            v25 = (v25 << ((2) & 63))
        } else {
            if (v28 <= 0.0) {
                for fi29 := 0; fi29 < 10; fi29++ {
                    v30 := GenS2{F0: 65535}
                    _ = fi29
                    _ = v30
                }
            } else {
                if ((^(v25)) != v25) {
                    if (-445 > v2) {
                        fmt.Println(v2)
                    } else {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            for fi32 := 0; fi32 < 6; fi32++ {
                                v26 = v26
                                v2 = 0
                                var v33 int = -1
                                _ = fi32
                                _ = v33
                            }
                            if (false && (false || false)) {
                                var v34 int64 = (v2 >> ((18446744073709551615) & 63))
                                _ = v34
                            } else {
                                v35 := GenS3{F0: -128, F1: 1000}
                                var v36 float64 = 0.1
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS4{F0: 42, F1: 1}
                            _ = fi31
                            _ = v37
                        }
                    }
                    switch (v26 + "hello") {
                    case "\n\t":
                        fmt.Println(v26, v1)
                    }
                    v2 = (^(v2))
                }
                fmt.Println(v3, v1, v2)
                if ("\n\t" > (v26 + v26)) {
                    v2 += (v2 | v2)
                    fmt.Println(v3, v26)
                }
            }
        }
        _ = v28
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := GenS5{F0: 148, F1: 18446744073709551615}
        if ((0 <= 72) || (fi38 >= (fi38 / (((-(fi38))) | 1)))) {
            v40 := GenS4{F0: 445, F1: 102}
            var v41 int = (-(fi38))
            _ = v40
            _ = v41
        }
        v42 := GenS6{F0: fi38, F1: 2147483648, F2: (!true)}
        _ = fi38
        _ = v39
        _ = v42
    }
    if false {
        var v43 string = v26
        v3 = (v3 >> ((4294967295) & 63))
        _ = v43
    }
    fmt.Println(v3, v1, v2)
    var v44 uint16 = 131
    var v45 uint16 = (v44 / ((v44) | 1))
    v2 += 1
    var v46 bool = false
    v26 = "a"
    v44 = (^(v45))
    v26 = v26
    if true {
        fmt.Println(v44, v1)
        v47 := GenS7{F0: -109, F1: 127}
        v46 = false
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
