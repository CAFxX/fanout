package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 65535
    var v3 int16 = 1
    var v4 uint32 = v2
    var v5 bool = true
    var v6 uint = 908
    v5 = (!v5)
    v7 := GenS1{F0: -2147483648}
    if true {
        if v5 {
            v8 := GenS2{F0: false, F1: 1}
            v9 := []int{2147483647, -342, -2147483648}
            _ = v8
            _ = v9
        }
    } else {
        var v10 float32 = 1.5
        v6 = v6
        _ = v10
    }
    v4 = (^(v4))
    var v11 int = 499
    if v5 {
        fmt.Println(v6, v4)
        v3 += v3
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: (v6 < 537), F1: 42}
        fmt.Println(v3, v11)
        v4 = (v4 + (v4 | (v4 % ((2147483648) | 1))))
        _ = fi12
        _ = v13
    }
    switch "foo bar" {
    case "hello":
        v14 := []int{-943}
        _ = v14
    default:
        var v15 int64 = -544
        _ = v15
    }
    var v16 uint = v6
    if (false || false) {
        var v17 int8 = -1
        _ = v17
    } else {
        var v18 int8 = -45
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 uint64 = 0
            _ = fi19
            _ = v20
        }
        var v21 bool = v5
        _ = v18
        _ = v21
    }
    v22 := GenS1{F0: -159}
    var v23 int = v11
    if (v2 == (v4 >> ((598) & 63))) {
        if v5 {
            v24 := []int{}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                if (!v5) {
                    if true {
                        fmt.Println(v11, v3, v5)
                        v26 := GenS3{F0: 776, F1: v23, F2: 2}
                        _ = v26
                    } else {
                        if false {
                            if ((^(v6)) != (^(v16))) {
                                switch "a" {
                                case "":
                                    v27 := GenS3{F0: int16(v4), F1: fi25, F2: 2}
                                    _ = v27
                                default:
                                    for fi28 := 0; fi28 < 13; fi28++ {
                                        var v29 uint16 = 2
                                        fmt.Println(v23)
                                        v23 -= (v11 >> (((^(v6))) & 63))
                                        v16 = 0
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                                if false {
                                    switch v23 {
                                    case -730:
                                        v30 := []int{(fi25 - (-(v23))), (v23 - v23)}
                                        _ = v30
                                    }
                                    fmt.Println(v2)
                                    var v31 bool = v5
                                    _ = v31
                                } else {
                                    var v32 int16 = (v3 << ((v16) & 63))
                                    if (!((238 >= 767) && v5)) {
                                        if (-2147483648 < -945) {
                                            for fi33 := 0; fi33 < 1; fi33++ {
                                                if (42 < 127) {
                                                    v4 += (v4 >> ((v6) & 63))
                                                    if (!v5) {
                                                        var v34 int8 = 0
                                                        if (!(!((false || false) && v5))) {
                                                            v16 = (v6 + (v6 / ((v6) | 1)))
                                                            v23 += 216
                                                            var v35 float32 = 1.5
                                                            _ = v35
                                                        }
                                                        _ = v34
                                                    }
                                                } else {
                                                    v32 = (v32 + (v32 / ((v32) | 1)))
                                                }
                                                _ = fi33
                                            }
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                v3 -= int16(fi25)
                                            case "":
                                                v5 = v5
                                            }
                                            var v36 uint16 = 384
                                            _ = v36
                                        }
                                    }
                                    _ = v32
                                }
                            } else {
                                for fi37 := 0; fi37 < 19; fi37++ {
                                    v11 = (fi37 % ((fi37) | 1))
                                    _ = fi37
                                }
                                v16 = v16
                                v11 += (-(fi25))
                            }
                            v11 = int(v3)
                            if false {
                                for fi38 := 0; fi38 < 15; fi38++ {
                                    v23 -= v23
                                    _ = fi38
                                }
                            } else {
                                v11 += v11
                            }
                        }
                        v39 := []int{}
                        _ = v39
                    }
                    var v40 uint64 = 1
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 5; fi41++ {
                        var v42 int = 32767
                        switch fi41 {
                        case 4294967295:
                            if ((!true) || v5) {
                                v1[""] = fi41
                            } else {
                                v6 += (v16 * (v6 << ((v16) & 63)))
                                v3 += (v3 & int16(v4))
                                _ = len(v1)
                                switch "\n\t" {
                                case "a":
                                    if true {
                                        v16 = (v6 >> ((18446744073709551615) & 63))
                                        v4 = 42
                                    }
                                default:
                                    if (v5 || (!v5)) {
                                        _ = len(v1)
                                        fmt.Println(fi25, v3)
                                        fmt.Println(fi41)
                                    } else {
                                        if v5 {
                                            var v43 int16 = v3
                                            var v44 int8 = 127
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi41
                        _ = v42
                    }
                    v45 := GenS3{F0: (v3 & (v3 % (((v3 << ((v6) & 63))) | 1))), F1: (v11 * v11), F2: 2}
                    _ = v45
                }
                _ = fi25
            }
            switch "" {
            case "a":
                fmt.Println(v4, v6)
            case "hello":
                fmt.Println(v23, v4, v16)
            case "\n\t":
                v46 := GenS4{F0: v5}
                _ = v46
            }
            if (v5 || v5) {
                v47 := []int{0, (v11 & 42)}
                var v48 int64 = -9223372036854775808
                _ = v47
                _ = v48
            }
        }
        v1[""] = v23
        for fi49 := 0; fi49 < 5; fi49++ {
            var v50 int8 = 100
            v5 = v5
            v51 := GenS1{F0: v23}
            _ = fi49
            _ = v50
            _ = v51
        }
    } else {
        v52 := GenS5{F0: -1}
        v53 := []int{v11, (v11 + (^(v11))), v11}
        fmt.Println(v3)
        _ = v52
        _ = v53
    }
    if true {
        var v54 float64 = (1e10 - 0.0)
        switch 0 {
        case 42:
            var v55 uint = 833
            _ = v55
        default:
            v2 = (v2 >> ((v16) & 63))
        }
        v54 -= v54
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
