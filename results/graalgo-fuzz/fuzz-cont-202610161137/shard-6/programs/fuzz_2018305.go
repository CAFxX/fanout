package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 string; F1 bool; F2 int64 }
func genFn0() uint {
    v6 := GenS1{F0: 0}
    if false {
        if true {
            var v7 uint64 = 4294967295
            _ = v7
        } else {
            if true {
                v8 := GenS2{F0: 40, F1: 0, F2: 0}
                _ = v8
            }
            var v9 uint32 = 869
            _ = v9
        }
        v10 := GenS3{F0: 9223372036854775807, F1: 127}
        _ = v10
    }
    var v11 uint64 = 879
    v12 := GenS4{F0: ("\n\t" + "a"), F1: 743}
    v11 += 681
    _ = v11
    _ = v12
    _ = v6
    return uint(v11)
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]uint32)
    var v3 bool = true
    v4 := []int{0}
    v5 := []uint32{}
    v13 := GenS5{F0: 1, F1: 42}
    fmt.Println(v3)
    v14 := GenS6{F0: 130}
    if (821 != 2147483647) {
        if false {
            v1 = v1
            switch 124 {
            case -1:
                if true {
                    if (!v3) {
                        _ = len(v2)
                        fmt.Println(v1, v3)
                        if v3 {
                            v15 := GenS6{F0: 30}
                            _ = v15
                        } else {
                            delete(v2, 0)
                            if true {
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    v17 := GenS3{F0: fi16, F1: (fi16 - (fi16 - fi16))}
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            v3 = false
                        }
                    } else {
                        delete(v2, 1)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v18 int8 = -108
                        _ = v18
                    }
                    fmt.Println(v3, v1)
                } else {
                    _ = len(v2)
                }
            case -375:
                switch ((("" + "foo bar") + "") + ("hello" + "\n\t")) {
                case "foo bar":
                    if (0 > 453) {
                        v3 = (v3 && v3)
                    }
                case "\n\t":
                    v5 = append(v5, 2147483647)
                default:
                    delete(v2, 42)
                }
            case 32768:
                var v19 float64 = -2.25
                _ = v19
            }
        } else {
            if (false && true) {
                fmt.Println(v1, v1)
                v1 -= (v1 ^ 42)
                var v20 int32 = 0
                _ = v20
            }
            v3 = v3
        }
    } else {
        v21 := GenS7{F0: 255, F1: -126}
        _ = v21
    }
    var v22 int8 = 127
    var v23 int8 = 42
    var v24 string = "foo bar"
    switch 0 {
    case 4294967295:
        if v3 {
            delete(v2, 2147483647)
            v25 := GenS8{F0: v24, F1: v24}
            _ = v25
        } else {
            _ = len(v2)
            switch (v24 + v24) {
            case "":
                var v26 int16 = (v1 ^ -32768)
                _ = v26
            case "foo bar":
                if true {
                    var v27 string = v24
                    _ = v27
                } else {
                    v22 = v23
                }
            case "a":
                if v3 {
                    if true {
                        var v28 int = -1
                        if true {
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                        _ = v28
                    }
                } else {
                    fmt.Println(v23, v1, v23)
                }
            default:
                fmt.Println(v22, v3)
            }
        }
    case -9223372036854775808:
        if ((!((v1 & -339) < v1)) && (((((3.14159 * -2.25) / 1.0) / 1e10) - ((-2.25 + (1e10 / ((0.1 - 1.5) / 0.1))) / 1.0)) > 1e10)) {
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 int32 = -658
                var v31 int64 = -29
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS3{F0: (fi29 >> ((genFn0()) & 63)), F1: (fi29 & fi29)}
                    var v34 int32 = (v30 >> ((1) & 63))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = fi29
                _ = v30
                _ = v31
            }
            v35 := GenS9{F0: "foo bar", F1: ((v23 % ((42) | 1)) <= 37), F2: 260}
            _ = v35
        }
    case 127:
        _ = len(v2)
    default:
        var v36 uint = 367
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
