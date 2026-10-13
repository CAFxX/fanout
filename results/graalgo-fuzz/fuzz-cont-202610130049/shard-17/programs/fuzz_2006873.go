package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint16; F2 int }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 int64 }
func genFn0() int16 {
    var v6 uint16 = 32767
    if false {
        v7 := GenS1{F0: 1, F1: 166}
        v8 := GenS1{F0: 666, F1: 0}
        switch "foo bar" {
        case "\n\t":
            v9 := GenS1{F0: 42, F1: 1}
            _ = v9
        case "hello":
            fmt.Println(v6)
        case "":
            for fi10 := 0; fi10 < 5; fi10++ {
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(v6, fi11, fi10)
                    if false {
                        v12 := []int{(fi11 >> ((0) & 63)), fi10, int(v6), (fi10 << ((65535) & 63))}
                        fmt.Println(fi10, fi10)
                        _ = v12
                    } else {
                        v13 := []int{-301, (fi11 * (fi11 ^ fi11)), fi10}
                        if ((((3.14159 - 1e10) * 3.14159) > (0.1 * -2.25)) || (!(("" + ("\n\t" + ("" + ("\n\t" + "hello")))) <= "a"))) {
                            if (("a" + "a") > ("a" + "a")) {
                                var v14 int = 32767
                                v14 = (fi11 & 32767)
                                _ = v14
                            }
                        } else {
                            v15 := GenS2{F0: "foo bar"}
                            var v16 int64 = 1
                            v13 = append(v13, fi11)
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            var v17 int32 = -2147483648
                            if true {
                                var v18 uint32 = 580
                                var v19 uint32 = (^(v18))
                                var v20 int64 = (v16 & v16)
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                if (false || true) {
                                    if len(v13) > 0 {
                                        v13[0] = (-(fi11))
                                    }
                                    if (v6 >= 2) {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        var v21 string = "a"
                                        if (v21 <= v21) {
                                            if false {
                                                var v22 int16 = 1
                                                _ = v22
                                            } else {
                                                var v23 uint8 = 127
                                                switch (v21 + "foo bar") {
                                                case "a":
                                                    switch int(v17) {
                                                    case -1:
                                                        fmt.Println(fi10, v23, v6)
                                                    case 982:
                                                        if false {
                                                            v24 := GenS2{F0: v21}
                                                            _ = v24
                                                        }
                                                    default:
                                                        v25 := GenS3{F0: (v6 >> ((738) & 63))}
                                                        _ = v25
                                                    }
                                                }
                                                var v26 int32 = (-(v17))
                                                _ = v23
                                                _ = v26
                                            }
                                        } else {
                                            fmt.Println(v17)
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            fmt.Println(v17, v17)
                                            var v27 uint64 = 710
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            _ = v27
                                        }
                                        _ = v21
                                    } else {
                                        var v28 int32 = (v17 << ((2) & 63))
                                        _ = v28
                                    }
                                }
                                var v29 bool = ((2 <= 42) || false)
                                v16 = (v16 & -9223372036854775808)
                                _ = v29
                            }
                            var v30 uint32 = 42
                            var v31 int32 = -950
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            fmt.Println(v31, fi10, v16)
                            var v32 int32 = -538
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        if (887 >= 0) {
                            v33 := GenS1{F0: 0, F1: 65}
                            fmt.Println(v6)
                            v34 := GenS4{F0: 100, F1: 736, F2: 0}
                            _ = v33
                            _ = v34
                        }
                        _ = v13
                    }
                    if ((1000 < 0) && true) {
                        fmt.Println(v6, fi11, fi11)
                        var v35 int8 = 96
                        v6 = 956
                        v36 := []int{}
                        _ = v35
                        _ = v36
                    } else {
                        var v37 int16 = 828
                        var v38 int16 = v37
                        v6 -= (v6 / ((v6) | 1))
                        _ = v37
                        _ = v38
                    }
                    _ = fi11
                }
                v39 := GenS3{F0: (^(v6))}
                _ = fi10
                _ = v39
            }
        default:
            if false {
                v40 := GenS5{F0: 338, F1: -25}
                v6 = v6
                _ = v40
            }
        }
        _ = v7
        _ = v8
    }
    v41 := GenS6{F0: 42}
    v6 = (v6 * v6)
    var v42 bool = false
    _ = v41
    _ = v42
    _ = v6
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := []int{}
    var v3 int = 32767
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v3 = (v3 >> ((2) & 63))
    v1 -= (v1 % ((v1) | 1))
    v1 = -110
    var v43 int64 = int64(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= v43
    v44 := GenS1{F0: (v43 & v43), F1: 800}
    v45 := GenS7{F0: 727, F1: v43}
    v5[v3] = "hello"
    _ = len(v5)
    v43 += 1
    var v46 uint16 = 42
    v3 += v3
    var v47 int64 = 285
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
