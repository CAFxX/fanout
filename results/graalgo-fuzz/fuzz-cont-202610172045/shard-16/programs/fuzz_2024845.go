package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint64 }
func genFn0(p0 uint8) uint16 {
    var v6 uint32 = 4294967295
    p0 = 255
    fmt.Println(p0, p0)
    _ = p0
    _ = v6
    return 2
}
func genFn1() uint32 {
    var v7 int8 = 100
    v8 := []int{-2147483648, -1, -2147483648, -9223372036854775808}
    switch -46 {
    case 9223372036854775807:
        v7 -= (v7 | 112)
    case -9223372036854775808:
        for fi9 := 0; fi9 < 14; fi9++ {
            v7 = (v7 % ((v7) | 1))
            _ = fi9
        }
    default:
        var v10 float64 = 1e10
        _ = v10
    }
    _ = v7
    _ = v8
    return 563
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), ((("" + "a") + "foo bar") + "a")}
    v2 := []int{252, -621, 485, -193}
    var v3 bool = (850 > 2)
    v4 := []uint32{4294967295}
    var v5 uint = 42
    var v11 uint16 = 0
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint32 = genFn1()
        _ = fi12
        _ = v13
    }
    var v14 bool = v3
    var v15 int64 = 42
    if len(v1) > 0 {
        v1[0] = (("hello" + "hello") + "")
    }
    var v16 int8 = 1
    _ = len(v4)
    var v17 int16 = 42
    v17 -= -32768
    v14 = false
    var v18 int = 44
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            if true {
                var v21 float32 = -2.25
                _ = v21
            } else {
                if (244 <= 1) {
                    var v22 int8 = (v16 - v16)
                    v23 := GenS1{F0: v5}
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS2{F0: 2}
                    switch "a" {
                    case "\n\t":
                        switch 801 {
                        case 813:
                            switch v18 {
                            case -322:
                                var v25 float32 = 0.1
                                switch int(v11) {
                                case 0:
                                    v26 := GenS3{F0: -9223372036854775808}
                                    v27 := GenS3{F0: 1}
                                    _ = v26
                                    _ = v27
                                case 806:
                                    v28 := GenS1{F0: (^(v5))}
                                    fmt.Println(v3, fi19)
                                    _ = v28
                                }
                                _ = v25
                            case 402:
                                v14 = (true || v3)
                            }
                        case 938:
                            var v29 int = v18
                            _ = v29
                        default:
                            v30 := GenS1{F0: (v5 & (v5 << ((v5) & 63)))}
                            _ = v30
                        }
                    default:
                        if v3 {
                            v31 := GenS2{F0: 0}
                            if v14 {
                                var v32 uint = v5
                                _ = v32
                            }
                            _ = v31
                        }
                    }
                    v11 = (v11 * (v11 % (((v11 % ((106) | 1))) | 1)))
                    _ = v24
                }
                v5 -= v5
                switch ("" + "\n\t") {
                case "\n\t":
                    if (!(356 == (v5 >> ((513) & 63)))) {
                        var v33 uint16 = v11
                        _ = v33
                    } else {
                        fmt.Println(fi20, v17, v18)
                    }
                }
            }
            var v34 bool = (!(!(230 < 128)))
            var v35 int64 = v15
            var v36 uint16 = 1
            _ = fi20
            _ = v34
            _ = v35
            _ = v36
        }
        _ = fi19
    }
    v2 = append(v2, v18)
    fmt.Println(v3)
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 int32 = 42
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 uint16 = 176
            if v14 {
                switch ("hello" + "") {
                case "foo bar":
                    v15 = 1
                }
                var v41 uint64 = 573
                _ = v41
            } else {
                var v42 uint32 = 632
                v2 = append(v2, (^(fi37)))
                _ = v42
            }
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = fi39
            _ = v40
        }
        for fi43 := 0; fi43 < 8; fi43++ {
            v44 := GenS4{F0: 42, F1: 119}
            fmt.Println(v18, fi43)
            fmt.Println(fi37)
            _ = fi43
            _ = v44
        }
        _ = fi37
        _ = v38
    }
    var v45 uint16 = genFn0(120)
    v46 := GenS5{F0: 240, F1: v16}
    for fi47 := 0; fi47 < 14; fi47++ {
        v48 := GenS6{F0: int16(fi47)}
        fmt.Println(v18, v45)
        for fi49 := 0; fi49 < 4; fi49++ {
            v50 := GenS7{F0: 42, F1: v15, F2: 18446744073709551615}
            _ = fi49
            _ = v50
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
