package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 int8 }
func genFn0(p0 int8, p1 bool) int {
    v7 := GenS1{F0: -1, F1: (!true)}
    for fi8 := 0; fi8 < 8; fi8++ {
        if p1 {
            var v9 int32 = 0
            fmt.Println(p0, v9, p1)
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: (-(fi10)), F1: (!("foo bar" > ("hello" + "")))}
                if p1 {
                    v9 = (v9 << ((65535) & 63))
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            fmt.Println(p0, p1, p0)
            switch (fi8 + 1) {
            case 32767:
                switch 0 {
                case -2147483648:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 string = ("" + "")
                        if p1 {
                            v14 := []int{(fi8 % (((^(fi8))) | 1)), fi12, (-(fi12))}
                            _ = v14
                        }
                        v15 := GenS1{F0: (fi12 * fi8), F1: (1 == 2)}
                        _ = fi12
                        _ = v13
                        _ = v15
                    }
                default:
                    fmt.Println(p0, p0)
                }
            default:
                if p1 {
                    p0 += -1
                } else {
                    v16 := []int{fi8, (fi8 - 9223372036854775807)}
                    p1 = false
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint16 = 65535
                v19 := []int{}
                p1 = (42 != 1)
                fmt.Println(p1, p0)
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        v20 := GenS2{F0: false, F1: (fi8 + -9223372036854775808), F2: 1}
        fmt.Println(p0)
        _ = fi8
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v7
    return 32768
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1e10
    var v3 uint64 = 1
    var v4 uint8 = 55
    v5 := []int{43}
    var v6 int64 = -838
    if false {
        var v21 uint64 = 4294967295
        _ = v21
    } else {
        var v22 float32 = v2
        fmt.Println(v3)
        _ = v22
    }
    switch 42 {
    case -9223372036854775808:
        v23 := GenS3{F0: -93}
        _ = v23
    }
    v1[-336] = -423
    switch 99 {
    case 803:
        v24 := GenS4{F0: -408}
        _ = v24
    case 42:
        v4 += (^(v4))
    default:
        fmt.Println(v2)
    }
    var v25 int = -9223372036854775808
    v2 -= (v2 - ((v2 / v2) + v2))
    var v26 int8 = -8
    if ((-(v6)) == v6) {
        fmt.Println(v3)
    }
    if true {
        for fi27 := 0; fi27 < 13; fi27++ {
            switch "hello" {
            case "hello":
                fmt.Println(v25)
            case "foo bar":
                var v28 bool = ((115 < 2147483647) && true)
                _ = v28
            case "a":
                if false {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if (255 != 4294967295) {
                            _ = len(v5)
                            var v30 uint16 = 2
                            _ = v30
                        }
                        v3 = (v3 << ((18446744073709551615) & 63))
                        _ = fi29
                    }
                } else {
                    if false {
                        var v31 uint8 = 2
                        _ = v31
                    } else {
                        delete(v1, genFn0(v26, false))
                    }
                    v6 += v6
                }
            default:
                switch int(v4) {
                case 32768:
                    if false {
                        var v32 float64 = ((1e10 / (-2.25 + 0.0)) * 1.5)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v33 := GenS1{F0: (v25 % (((^(fi27))) | 1)), F1: false}
                        _ = v32
                        _ = v33
                    } else {
                        fmt.Println(v6, v4)
                        _ = len(v5)
                    }
                case 32767:
                    v34 := GenS2{F0: (true || true), F1: (v25 / ((v25) | 1)), F2: 128}
                    _ = v34
                default:
                    if false {
                        switch "foo bar" {
                        case "hello":
                            var v35 int = (^(fi27))
                            _ = v35
                        }
                        var v36 int16 = 1
                        fmt.Println(v3)
                        _ = v36
                    } else {
                        switch "hello" {
                        case "\n\t":
                            v37 := GenS5{F0: -1}
                            _ = v37
                        case "foo bar":
                            for fi38 := 0; fi38 < 1; fi38++ {
                                fmt.Println(fi27, v2, fi27)
                                v39 := GenS1{F0: fi38, F1: (587 < 779)}
                                var v40 uint64 = 0
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        }
                        for fi41 := 0; fi41 < 15; fi41++ {
                            var v42 int = -1
                            v25 = (v25 | (v25 * v42))
                            _ = fi41
                            _ = v42
                        }
                    }
                }
            }
            _ = fi27
        }
        if (-32768 == 0) {
            var v43 uint64 = v3
            var v44 int16 = 0
            v6 -= (^(v6))
            _ = v43
            _ = v44
        } else {
            var v45 uint16 = 364
            var v46 float32 = v2
            var v47 int16 = -1
            _ = v45
            _ = v46
            _ = v47
        }
    }
    v4 -= (v4 & uint8(v3))
    for fi48 := 0; fi48 < 10; fi48++ {
        v3 += v3
        fmt.Println(v6)
        v49 := GenS6{F0: 153, F1: -702, F2: 16}
        _ = fi48
        _ = v49
    }
    switch ("" + "\n\t") {
    case "\n\t":
        v50 := GenS7{F0: uint(v3)}
        _ = v50
    }
    switch (v25 - v25) {
    case 2147483647:
        var v51 uint64 = 841
        _ = v51
    case 0:
        v52 := GenS8{F0: v26, F1: v26}
        _ = v52
    case 128:
        v3 -= v3
    }
    fmt.Println(v2)
    fmt.Println(v4, v3)
    delete(v1, (-(v25)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
