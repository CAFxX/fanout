package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint8, p1 uint16, p2 uint16) uint32 {
    var v6 uint32 = 1
    v6 = 780
    p0 -= p0
    if (!false) {
        var v7 uint64 = 706
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint32 = v6
            if true {
                v10 := []int{fi8, fi8, fi8, fi8}
                if false {
                    v10 = append(v10, (fi8 << ((65535) & 63)))
                } else {
                    switch (fi8 / (((-(fi8))) | 1)) {
                    case 2147483647:
                        v7 = (^(v7))
                    case 9223372036854775807:
                        switch (("foo bar" + "\n\t") + "") {
                        case "\n\t":
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        case "hello":
                            v10 = append(v10, (fi8 * (-(fi8))))
                        }
                    }
                    v11 := GenS1{F0: -1}
                    switch fi8 {
                    case -2147483648:
                        var v12 int32 = -2147483648
                        _ = v12
                    case 147:
                        var v13 int16 = 1000
                        _ = v13
                    case 128:
                        if (238 < fi8) {
                            v7 += v7
                        } else {
                            fmt.Println(p1, v6, v7)
                        }
                    }
                    _ = v11
                }
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        if ((3.14159 - 0.0) == 0.0) {
            var v14 bool = true
            _ = v14
        } else {
            p1 = 161
        }
        _ = v7
    }
    switch 128 {
    case 19:
        var v15 int8 = 42
        _ = v15
    default:
        var v16 int32 = 524
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (^(v6))
}
func genFn1() uint16 {
    if false {
        v17 := GenS2{F0: -501}
        var v18 uint64 = 2
        if true {
            v19 := GenS1{F0: 1000}
            if true {
                if (!(2 <= 163)) {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v18 += v18
                        v21 := []int{fi20}
                        _ = fi20
                        _ = v21
                    }
                } else {
                    v22 := []int{}
                    _ = v22
                }
            } else {
                var v23 float64 = 1e10
                v24 := GenS3{F0: 2}
                v25 := []int{32768, -9223372036854775808}
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v19
        } else {
            v18 = (v18 ^ v18)
            v18 = (v18 - v18)
            var v26 float64 = (3.14159 + 1e10)
            _ = v26
        }
        _ = v17
        _ = v18
    }
    if false {
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint32 = 4294967295
            _ = fi27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int = fi29
            v32 := []int{}
            if len(v32) > 0 {
                v32[0] = fi29
            }
            switch fi29 {
            case 0:
                if true {
                    _ = len(v32)
                    v31 -= 2147483647
                    if ((!false) && false) {
                        switch "foo bar" {
                        case "":
                            if false {
                                var v33 int = (fi29 | (fi30 / ((v31) | 1)))
                                fmt.Println(fi30, v33, fi30)
                                _ = v33
                            }
                        case "\n\t":
                            fmt.Println(fi29, fi29)
                        default:
                            var v34 bool = (65535 == 32767)
                            _ = v34
                        }
                    }
                    if len(v32) > 0 {
                        v32[0] = -430
                    }
                    if (0.0 == -2.25) {
                        if false {
                            var v35 int64 = 112
                            _ = v35
                        } else {
                            v31 = v31
                        }
                        if (32 >= -128) {
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        }
                        v31 = 2147483647
                        fmt.Println(fi29)
                    }
                } else {
                    v31 = 127
                    if false {
                        fmt.Println(fi29, fi30)
                        var v36 uint = 432
                        v36 += v36
                        _ = v36
                    }
                    var v37 int64 = -1
                    _ = v37
                }
            case 42:
                _ = len(v32)
            case 127:
                v31 += 42
            default:
                if (!false) {
                    v38 := GenS4{F0: 1}
                    var v39 int = -2147483648
                    v40 := GenS5{F0: 1, F1: 1}
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        var v41 int32 = 1
        _ = fi29
        _ = v41
    }
    return 553
}

func main() {
    var v1 int32 = 673
    v2 := make(map[uint32]int)
    var v3 int16 = 1000
    var v4 uint32 = 2
    v5 := []uint32{}
    var v42 string = ("hello" + "")
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 uint16 = 1
        if len(v5) > 0 {
            v5[0] = genFn0(2, 42, 178)
        }
        fmt.Println(v44, v1, v3)
        _ = fi43
        _ = v44
    }
    if ((2 == (^(v4))) && false) {
        for fi45 := 0; fi45 < 20; fi45++ {
            if true {
                v2[(v4 / ((v4) | 1))] = 1
                fmt.Println(fi45, v3)
                var v46 uint8 = 39
                _ = v46
            }
            switch (fi45 ^ fi45) {
            case 9223372036854775807:
                for fi47 := 0; fi47 < 9; fi47++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v1)
                    fmt.Println(v3, v3, fi45)
                    _ = fi47
                }
            }
            v1 = v1
            _ = fi45
        }
        v48 := GenS3{F0: 0}
        if true {
            _ = len(v5)
        } else {
            v1 -= -666
            fmt.Println(v42, v1, v1)
            v1 += int32(v3)
        }
        _ = v48
    }
    fmt.Println(v1)
    fmt.Println(v4)
    var v49 bool = (true && true)
    v1 += (v1 ^ v1)
    fmt.Println(v3, v4)
    for fi50 := 0; fi50 < 7; fi50++ {
        _ = len(v2)
        _ = fi50
    }
    var v51 int16 = (v3 | v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v51
}
