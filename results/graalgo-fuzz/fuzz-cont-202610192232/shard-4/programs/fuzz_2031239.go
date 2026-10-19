package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int64; F1 int; F2 uint }
func genFn0(p0 int64, p1 int64, p2 uint) int16 {
    var v7 int16 = 1
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            if true {
                v10 := []int{fi8, (fi9 >> ((p2) & 63)), fi8}
                var v11 uint8 = 255
                if false {
                    if false {
                        p0 -= p1
                    } else {
                        var v12 int64 = p0
                        fmt.Println(fi8)
                        v12 = (v12 & p0)
                        _ = v12
                    }
                    if false {
                        fmt.Println(fi9, fi9, p2)
                        var v13 int64 = 42
                        v11 += v11
                        _ = v13
                    } else {
                        p1 = (p1 ^ (p1 >> ((p2) & 63)))
                        var v14 bool = ((("foo bar" + "hello") + "") >= ("" + "a"))
                        _ = v14
                    }
                    p2 += p2
                } else {
                    v15 := GenS1{F0: v11}
                    p0 = p1
                    if (!true) {
                        fmt.Println(v11)
                        v16 := GenS2{F0: v7}
                        if false {
                            v7 = int16(v11)
                            var v17 uint8 = v11
                            _ = v17
                        } else {
                            v18 := GenS1{F0: v11}
                            _ = v18
                        }
                        _ = v16
                    }
                    _ = v15
                }
                switch (fi9 & 2147483647) {
                case 608:
                    switch "foo bar" {
                    case "hello":
                        _ = len(v10)
                        if ((-2.25 <= 1e10) && true) {
                            var v19 uint8 = v11
                            _ = v19
                        } else {
                            if (1 != 14) {
                                v7 -= (-(v7))
                                fmt.Println(v11)
                                if false {
                                    var v20 int8 = -1
                                    fmt.Println(fi9)
                                    switch fi9 {
                                    case 1:
                                        p0 = (-(p1))
                                    case 127:
                                        p2 = (p2 >> ((1) & 63))
                                    case 219:
                                        v21 := GenS3{F0: 823, F1: "", F2: ""}
                                        _ = v21
                                    default:
                                        var v22 int32 = 197
                                        _ = v22
                                    }
                                    var v23 int16 = v7
                                    _ = v20
                                    _ = v23
                                } else {
                                    var v24 int16 = (-(v7))
                                    if false {
                                        var v25 int16 = (v24 * -99)
                                        _ = v25
                                    }
                                    _ = v24
                                }
                            } else {
                                v26 := GenS4{F0: 406}
                                v11 = (^(v11))
                                v27 := GenS5{F0: p2, F1: 93}
                                if (true && false) {
                                    var v28 int8 = int8(p2)
                                    _ = v28
                                } else {
                                    p0 = int64(v7)
                                    p1 = (p1 ^ 9223372036854775807)
                                }
                                _ = v26
                                _ = v27
                            }
                        }
                    case "\n\t":
                        fmt.Println(p2)
                    }
                default:
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                _ = v10
                _ = v11
            } else {
                p1 -= p1
                var v29 string = "a"
                _ = v29
            }
            _ = fi9
        }
        _ = fi8
    }
    v7 = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := []int{}
    var v2 float32 = 1.5
    v3 := make(map[uint32]int)
    var v4 int32 = -1
    var v5 int8 = 127
    var v6 int8 = v5
    v3[171] = 32767
    v4 = -2147483648
    if (-2147483648 <= (v4 >> ((42) & 63))) {
        var v30 int16 = 931
        var v31 uint = 4294967295
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 bool = false
            v34 := GenS6{F0: 2, F1: 309}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v30
        _ = v31
    } else {
        var v35 uint = 607
        v6 = 127
        v36 := GenS6{F0: 0, F1: 515}
        _ = v35
        _ = v36
    }
    var v37 float32 = float32(1.5)
    var v38 float64 = (3.14159 - 1e10)
    var v39 uint32 = 2147483647
    var v40 uint = 884
    if true {
        var v41 int32 = -212
        _ = v41
    }
    if true {
        if ((false || false) && (v4 == v4)) {
            if true {
                var v42 int16 = -1
                var v43 float32 = v2
                v42 = (v42 >> ((uint(v39)) & 63))
                _ = v42
                _ = v43
            } else {
                _ = len(v3)
                v2 = (v2 / v2)
            }
        }
        switch -365 {
        case -942:
            var v44 string = ("a" + "foo bar")
            _ = v44
        default:
            for fi45 := 0; fi45 < 10; fi45++ {
                v37 -= float32((v38 / (v38 / (v38 * (((float64(v37) * (v38 + 1.5)) * ((v38 / -2.25) - v38)) / v38)))))
                v46 := GenS3{F0: v4, F1: "\n\t", F2: "foo bar"}
                v4 = 1
                _ = fi45
                _ = v46
            }
        }
        if false {
            var v47 uint = 18446744073709551615
            _ = v47
        } else {
            var v48 int64 = 9223372036854775807
            for fi49 := 0; fi49 < 6; fi49++ {
                v48 = v48
                _ = fi49
            }
            v50 := GenS7{F0: -9223372036854775808, F1: 32767, F2: v40}
            _ = v48
            _ = v50
        }
    }
    var v51 int8 = -63
    for fi52 := 0; fi52 < 16; fi52++ {
        for fi53 := 0; fi53 < 12; fi53++ {
            var v54 float32 = 0.0
            _ = fi53
            _ = v54
        }
        _ = fi52
    }
    var v55 string = ("\n\t" + "foo bar")
    v55 = v55
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v51
    _ = v55
    _ = v6
}
