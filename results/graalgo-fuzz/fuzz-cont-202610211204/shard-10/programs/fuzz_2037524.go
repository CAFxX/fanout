package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 uint8 }
func genFn0() int32 {
    var v5 uint = 851
    for fi6 := 0; fi6 < 9; fi6++ {
        v5 = v5
        fmt.Println(fi6, v5)
        _ = fi6
    }
    fmt.Println(v5)
    _ = v5
    return -142
}
func genFn1() uint16 {
    var v7 int64 = 0
    if (-1 <= -1) {
        v8 := GenS1{F0: v7, F1: (255 > 942)}
        var v9 int64 = v7
        _ = v8
        _ = v9
    }
    if true {
        v7 -= (v7 * v7)
        if false {
            v10 := GenS1{F0: v7, F1: (false || true)}
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint64 = 4294967295
            v13 := []int{fi11, -1, fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        v14 := []int{80, 1}
        _ = v14
    }
    v7 = 42
    var v15 uint32 = 0
    _ = v15
    _ = v7
    return 65535
}

func main() {
    var v1 uint32 = 2
    var v2 string = ""
    var v3 string = "foo bar"
    var v4 int8 = 42
    var v16 int8 = -51
    v4 = (v16 ^ v4)
    for fi17 := 0; fi17 < 14; fi17++ {
        if (181 == 938) {
            if ((!false) && false) {
                v18 := GenS2{F0: 0, F1: 941, F2: -1}
                v2 = (v2 + (("\n\t" + "a") + ""))
                v4 = 100
                _ = v18
            } else {
                var v19 int64 = 42
                switch (v2 + "hello") {
                case "a":
                    v3 = v3
                case "hello":
                    v16 = v16
                case "":
                    var v20 bool = (!true)
                    _ = v20
                }
                v2 = ""
                _ = v19
            }
        }
        var v21 int8 = -8
        _ = fi17
        _ = v21
    }
    var v22 int8 = v4
    v22 = (-(v22))
    var v23 int64 = -9223372036854775808
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := []int{-2147483648}
        for fi26 := 0; fi26 < 9; fi26++ {
            v23 += v23
            if false {
                switch 42 {
                case 0:
                    if false {
                        var v27 int = fi26
                        _ = v27
                    }
                case 127:
                    v2 = v3
                default:
                    v3 = "a"
                }
                switch (v2 + (v3 + v2)) {
                case "":
                    _ = len(v25)
                case "a":
                    v28 := GenS3{F0: v4, F1: 127}
                    v1 = 2147483648
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    _ = v28
                case "hello":
                    v2 = (("hello" + ((v3 + (("" + (v3 + ("" + ("a" + "foo bar")))) + v2)) + v3)) + "\n\t")
                default:
                    var v29 uint64 = 1
                    _ = v29
                }
            } else {
                v22 -= -128
                if (-135 != -432) {
                    if false {
                        _ = len(v25)
                        if false {
                            v30 := GenS1{F0: -259, F1: (!true)}
                            var v31 int32 = 0
                            v32 := GenS4{F0: fi26, F1: 591}
                            v31 = genFn0()
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        if (true && false) {
                            if false {
                                switch (^(fi24)) {
                                case 883:
                                    v33 := GenS4{F0: 363, F1: 42}
                                    _ = v33
                                }
                                if (false || false) {
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    v4 += (-(v22))
                                }
                                v34 := GenS2{F0: 707, F1: 65535, F2: 0}
                                _ = v34
                            }
                            var v35 uint16 = genFn1()
                            v25 = append(v25, (fi24 * (fi24 / (((fi24 | (fi26 >> ((42) & 63)))) | 1))))
                            switch 2147483647 {
                            case -1:
                                if false {
                                    v3 = v2
                                } else {
                                    var v36 string = v3
                                    _ = len(v25)
                                    fmt.Println(fi24)
                                    v25 = append(v25, 590)
                                    _ = v36
                                }
                            case 128:
                                if true {
                                    if (-1 != 789) {
                                        switch ("hello" + "hello") {
                                        case "":
                                            v3 = ("\n\t" + v3)
                                        case "hello":
                                            fmt.Println(fi24, v4)
                                        case "a":
                                            if (1.5 != (-2.25 + 0.0)) {
                                                if true {
                                                    v35 += (v35 ^ genFn1())
                                                } else {
                                                    v35 = v35
                                                }
                                            }
                                        default:
                                            v2 = "foo bar"
                                        }
                                        if len(v25) > 0 {
                                            v25[0] = fi26
                                        }
                                    }
                                    v37 := GenS3{F0: -29, F1: 127}
                                    v35 += 65535
                                    if true {
                                        var v38 uint32 = (^(v1))
                                        _ = v38
                                    }
                                    if (((v23 >> ((847) & 63)) < 0) && false) {
                                        if ((fi24 * (fi26 * -1)) == (fi26 | (^(fi24)))) {
                                            v39 := GenS5{F0: 32767}
                                            v40 := GenS6{F0: int8(v23), F1: (v1 | v1)}
                                            v4 += v16
                                            fmt.Println(fi24, v22, v3)
                                            if false {
                                                switch v2 {
                                                case "foo bar":
                                                    var v41 uint32 = (v1 - v1)
                                                    _ = len(v25)
                                                    v42 := GenS7{F0: 69}
                                                    _ = v41
                                                    _ = v42
                                                case "":
                                                    var v43 uint8 = 10
                                                    _ = v43
                                                case "hello":
                                                    if true {
                                                        var v44 int = (fi24 / (((fi24 / ((fi26) | 1))) | 1))
                                                        v4 = v22
                                                        switch (fi24 ^ v44) {
                                                        case 127:
                                                            _ = len(v25)
                                                        }
                                                        _ = v44
                                                    } else {
                                                        var v45 int = fi26
                                                        _ = v45
                                                    }
                                                    fmt.Println(fi24, v2)
                                                }
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                    _ = v37
                                }
                            }
                            v1 -= v1
                            fmt.Println(fi24, v3)
                            _ = len(v25)
                            if (int16(v23) == 32767) {
                                v46 := GenS8{F0: 32767}
                                _ = v46
                            } else {
                                v35 -= (v35 & 1)
                                var v47 float64 = (1.5 + 0.1)
                                if len(v25) > 0 {
                                    v25[0] = fi26
                                }
                                _ = v47
                            }
                            _ = v35
                        }
                    }
                    v4 = (v22 << ((42) & 63))
                }
            }
            fmt.Println(v23, v1, v4)
            _ = fi26
        }
        if true {
            v1 = (v1 << ((42) & 63))
            switch (fi24 / (((^(fi24))) | 1)) {
            case 2147483647:
                for fi48 := 0; fi48 < 18; fi48++ {
                    if true {
                        var v49 int64 = v23
                        _ = v49
                    }
                    var v50 float32 = 0.0
                    _ = fi48
                    _ = v50
                }
            case -1:
                for fi51 := 0; fi51 < 20; fi51++ {
                    v52 := GenS9{F0: -1, F1: 128}
                    fmt.Println(fi24, v22, fi24)
                    var v53 float32 = 0.1
                    switch ((v2 + (v3 + "hello")) + (v3 + v3)) {
                    case "hello":
                        v3 = v2
                    default:
                        var v54 int = (fi24 << ((2) & 63))
                        _ = v54
                    }
                    v16 -= v16
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            }
            v23 -= (v23 | v23)
        }
        _ = fi24
        _ = v25
    }
    v2 = v3
    var v55 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v55
}
