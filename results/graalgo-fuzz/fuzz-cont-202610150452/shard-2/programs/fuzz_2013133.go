package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 uint64) int {
    fmt.Println(p0)
    var v7 int = 83
    v8 := GenS1{F0: 42}
    var v9 uint16 = 65535
    if true {
        v7 += (v7 >> ((651) & 63))
        v10 := []int{1, -2147483648}
        _ = v10
    } else {
        v11 := GenS1{F0: 2}
        if false {
            switch v7 {
            case 4294967295:
                if (!(("hello" > "\n\t") && false)) {
                    v9 += (v9 / ((v9) | 1))
                    switch (("" + "a") + "") {
                    case "foo bar":
                        v7 -= (^(v7))
                    case "hello":
                        fmt.Println(v7)
                    default:
                        switch v7 {
                        case -153:
                            switch ("" + "a") {
                            case "hello":
                                var v12 uint16 = (v9 - 858)
                                _ = v12
                            case "":
                                if (false && (!(!true))) {
                                    var v13 int32 = -1
                                    fmt.Println(v7, v13, v7)
                                    var v14 float32 = ((0.1 - 3.14159) / 1.5)
                                    _ = v13
                                    _ = v14
                                }
                            case "foo bar":
                                v15 := GenS1{F0: 9}
                                _ = v15
                            default:
                                v16 := GenS1{F0: 236}
                                _ = v16
                            }
                        case -137:
                            var v17 uint8 = 42
                            _ = v17
                        default:
                            v18 := []int{v7, (v7 * (v7 * (v7 << ((42) & 63)))), -1}
                            _ = v18
                        }
                    }
                } else {
                    v19 := []int{}
                    p0 += p0
                    v7 = 4294967295
                    _ = v19
                }
            case -171:
                var v20 uint8 = 110
                _ = v20
            case 32767:
                if (false || (-736 == 435)) {
                    v7 = (v7 & v7)
                    v21 := []int{}
                    switch "a" {
                    case "hello":
                        fmt.Println(v9, p0, v9)
                    case "\n\t":
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    case "a":
                        fmt.Println(v7, v9)
                    }
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 int8 = 2
                        v9 = v9
                        if false {
                            var v24 int8 = (v23 - (v23 * 27))
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 uint32 = 42
                                fmt.Println(v9)
                                _ = fi25
                                _ = v26
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                    v27 := GenS1{F0: 150}
                    _ = v27
                }
            default:
                switch v7 {
                case -6:
                    v7 = (v7 / ((v7) | 1))
                case -9223372036854775808:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v9 += 924
                        _ = fi28
                    }
                default:
                    var v29 bool = (9223372036854775807 > 42)
                    _ = v29
                }
            }
            v30 := []int{32768, -1, v7}
            p0 -= (p0 % ((p0) | 1))
            _ = v30
        } else {
            for fi31 := 0; fi31 < 10; fi31++ {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        v32 := GenS1{F0: 127}
                        _ = v32
                    }
                    v33 := []int{(v7 % ((-2147483648) | 1))}
                    p0 = p0
                    _ = v33
                } else {
                    v7 += v7
                    p0 = (p0 % ((p0) | 1))
                }
                v34 := []int{fi31}
                _ = fi31
                _ = v34
            }
            for fi35 := 0; fi35 < 16; fi35++ {
                p0 = (p0 << ((uint(p0)) & 63))
                var v36 int64 = -1
                _ = fi35
                _ = v36
            }
            v37 := GenS1{F0: 128}
            _ = v37
        }
        var v38 float64 = ((-2.25 - (float64(1e10) / 1.0)) - float64(((0.1 / -2.25) / 1.0)))
        _ = v11
        _ = v38
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 int8 = 42
    var v4 int16 = -933
    v5 := make(map[int]bool)
    var v6 int64 = -9223372036854775808
    v39 := GenS2{F0: -2147483648}
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v40 int16 = v4
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(v6)
        _ = fi41
    }
    v4 -= v4
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 uint64 = 1
        var v44 int32 = -1
        v45 := GenS3{F0: 1, F1: 0}
        _ = fi42
        _ = v43
        _ = v44
        _ = v45
    }
    v2 = append(v2, "a")
    v6 -= v6
    fmt.Println(v4, v40, v40)
    fmt.Println(v6)
    if (99 != 42) {
        v46 := GenS4{F0: "foo bar"}
        var v47 string = ("" + "a")
        _ = v46
        _ = v47
    }
    if true {
        v48 := GenS5{F0: 4294967295, F1: 42}
        for fi49 := 0; fi49 < 8; fi49++ {
            v50 := GenS4{F0: ""}
            v51 := GenS2{F0: (fi49 << ((18446744073709551615) & 63))}
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = v48
    }
    v3 = (v3 >> ((18446744073709551615) & 63))
    for fi52 := 0; fi52 < 16; fi52++ {
        if true {
            v40 += 0
        }
        _ = fi52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
