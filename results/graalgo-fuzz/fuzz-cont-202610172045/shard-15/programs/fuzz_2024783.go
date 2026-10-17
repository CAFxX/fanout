package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint16 }

func main() {
    var v1 string = "hello"
    var v2 int64 = 9223372036854775807
    var v3 uint64 = 436
    var v4 string = v1
    var v5 int = 284
    if true {
        v6 := GenS1{F0: 0, F1: (v5 - v5), F2: 42}
        _ = v6
    } else {
        var v7 int16 = int16(v5)
        _ = v7
    }
    var v8 uint8 = 128
    var v9 int32 = 0
    var v10 int32 = v9
    if false {
        v11 := []int{(v5 << ((18446744073709551615) & 63)), (v5 ^ v5), (-(v5)), (v5 % ((v5) | 1))}
        _ = v11
    } else {
        switch (^(v5)) {
        case 321:
            for fi12 := 0; fi12 < 6; fi12++ {
                v5 = (fi12 % ((-9223372036854775808) | 1))
                fmt.Println(v4, v3, v3)
                _ = fi12
            }
        case 32768:
            for fi13 := 0; fi13 < 13; fi13++ {
                if ((true || true) && (true || ((v4 + v1) < ("" + ((("\n\t" + (v1 + "\n\t")) + "a") + ((v1 + "foo bar") + "hello")))))) {
                    v9 = v10
                    var v14 float64 = 0.0
                    fmt.Println(v5, v8, v5)
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        var v16 int64 = 42
                        var v17 uint8 = v8
                        v5 -= (-(fi13))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 20; fi18++ {
                        v19 := []int{fi18, 282, 42}
                        if false {
                            _ = len(v19)
                            var v20 int16 = 32767
                            if false {
                                var v21 int = fi18
                                var v22 string = "hello"
                                _ = v21
                                _ = v22
                            }
                            _ = v20
                        }
                        _ = fi18
                        _ = v19
                    }
                    if false {
                        v23 := []int{850, v5, 32767, (^(fi13))}
                        switch (v1 + v4) {
                        case "a":
                            fmt.Println(v4)
                        }
                        _ = v23
                    } else {
                        v2 -= (v2 << ((2) & 63))
                        var v24 uint16 = 65535
                        _ = v24
                    }
                }
                _ = fi13
            }
        default:
            if (!false) {
                fmt.Println(v8, v5)
                for fi25 := 0; fi25 < 18; fi25++ {
                    var v26 uint = 4294967295
                    _ = fi25
                    _ = v26
                }
                if ((true && true) && (true && (false || true))) {
                    if false {
                        if (((-(v5)) >= (v5 & (-(v5)))) || false) {
                            fmt.Println(v9, v3)
                        } else {
                            v8 = (v8 ^ (^(v8)))
                            switch v1 {
                            case "hello":
                                v27 := GenS1{F0: 32767, F1: (v5 >> ((0) & 63)), F2: (^(v8))}
                                _ = v27
                            case "\n\t":
                                v28 := []int{}
                                _ = v28
                            default:
                                v4 = ("" + v1)
                            }
                            var v29 uint = uint(v3)
                            _ = v29
                        }
                    } else {
                        if true {
                            if false {
                                v30 := []int{-9223372036854775808, (v5 ^ (v5 % ((2147483647) | 1)))}
                                _ = v30
                            } else {
                                if (902 == 42) {
                                    v31 := GenS2{F0: v8}
                                    v32 := []int{}
                                    _ = v31
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 9; fi33++ {
                                    var v34 int = (v5 + (v5 ^ (v5 + (v5 - (v5 % ((-544) | 1))))))
                                    for fi35 := 0; fi35 < 6; fi35++ {
                                        var v36 uint32 = 670
                                        _ = fi35
                                        _ = v36
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                                v8 = v8
                            }
                            for fi37 := 0; fi37 < 2; fi37++ {
                                var v38 uint = 42
                                switch ((v1 + v4) + (v1 + ("hello" + "foo bar"))) {
                                case "":
                                    v39 := GenS3{F0: 65535, F1: 18446744073709551615}
                                    _ = v39
                                }
                                _ = fi37
                                _ = v38
                            }
                            v1 = v1
                        }
                    }
                } else {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 float32 = (-2.25 / 1e10)
                        if false {
                            for fi42 := 0; fi42 < 20; fi42++ {
                                switch fi42 {
                                case 2147483647:
                                    v3 -= (v3 / ((v3) | 1))
                                case 880:
                                    v43 := []int{int(v10), (fi40 >> ((140) & 63)), 128}
                                    _ = v43
                                case -618:
                                    v44 := GenS3{F0: 2147483648, F1: v3}
                                    _ = v44
                                }
                                v45 := []int{9223372036854775807, 0}
                                _ = len(v45)
                                _ = fi42
                                _ = v45
                            }
                            var v46 int = v5
                            _ = v46
                        }
                        var v47 int16 = 32767
                        _ = fi40
                        _ = v41
                        _ = v47
                    }
                }
            } else {
                if false {
                    v48 := GenS4{F0: 42}
                    _ = v48
                } else {
                    if true {
                        v8 = (v8 ^ v8)
                        for fi49 := 0; fi49 < 5; fi49++ {
                            v5 = int(v9)
                            _ = fi49
                        }
                    } else {
                        v1 = ((v1 + v1) + ((v1 + (v4 + ("hello" + v1))) + v1))
                        for fi50 := 0; fi50 < 10; fi50++ {
                            v1 = v1
                            var v51 int8 = 127
                            _ = fi50
                            _ = v51
                        }
                    }
                    if false {
                        var v52 string = v4
                        var v53 bool = true
                        _ = v52
                        _ = v53
                    }
                    v2 -= v2
                }
                if (1.5 <= float64(0.1)) {
                    v10 = -2147483648
                    fmt.Println(v9)
                    v10 = (^(v9))
                }
                v54 := []int{(v5 / ((v5) | 1))}
                _ = v54
            }
        }
        var v55 int32 = 415
        _ = v55
    }
    v56 := GenS5{F0: 127}
    if ((^(v10)) == v10) {
        v1 = v1
        v5 += (v5 * v5)
    } else {
        for fi57 := 0; fi57 < 2; fi57++ {
            v58 := GenS6{F0: (!false), F1: 255}
            _ = fi57
            _ = v58
        }
    }
    v4 = "a"
    v2 = v2
    v2 = (v2 & (-(v2)))
    switch v1 {
    case "hello":
        v59 := GenS6{F0: (-1 == 0), F1: 264}
        _ = v59
    }
    var v60 int64 = (-(v2))
    var v61 int = (v5 * v5)
    v62 := []int{32767, (-(v61)), v5, v61}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v60
    _ = v61
    _ = v62
    _ = v8
    _ = v9
}
