package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint8; F1 int64; F2 int }
type GenS6 struct { F0 uint16; F1 int; F2 bool }
type GenS7 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 int16) uint16 {
    v7 := GenS1{F0: (p0 * 1)}
    for fi8 := 0; fi8 < 9; fi8++ {
        if (255 <= 1) {
            p0 = (p0 << ((587) & 63))
        } else {
            p0 -= (p0 * 172)
            v9 := []int{(fi8 * -1)}
            _ = v9
        }
        p0 = 395
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int = (fi8 & fi8)
            var v12 uint64 = 2
            v13 := GenS2{F0: (v11 << ((1) & 63))}
            var v14 int32 = 1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    if true {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := []int{}
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 string, p1 uint16, p2 string) int64 {
    v17 := []int{42, 32767, 0}
    fmt.Println(p0, p2)
    var v18 uint16 = (p1 | 897)
    var v19 string = (p0 + p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 int8 = 100
    var v5 int8 = -39
    var v6 float64 = float64(-2.25)
    var v20 int8 = -35
    fmt.Println(v6, v6)
    v21 := GenS2{F0: 4294967295}
    if false {
        var v22 float64 = v6
        v5 += v20
        v23 := v3[("" + "foo bar")]
        _ = v22
        _ = v23
    } else {
        var v24 int64 = 42
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                if false {
                    v20 = (v20 ^ v4)
                }
                switch "\n\t" {
                case "foo bar":
                    v24 -= v24
                default:
                    fmt.Println(v4, fi25, v6)
                }
                fmt.Println(v4, v6, v24)
                v24 = (-(v24))
                _ = fi26
            }
            _ = fi25
        }
        _ = v24
    }
    var v27 uint16 = genFn0(-198)
    if (!true) {
        for fi28 := 0; fi28 < 6; fi28++ {
            switch fi28 {
            case 128:
                v27 += (^(v27))
            case 9223372036854775807:
                v20 = (v5 % ((v20) | 1))
            case 42:
                var v29 int16 = -51
                _ = v29
            default:
                for fi30 := 0; fi30 < 17; fi30++ {
                    if false {
                        if true {
                            fmt.Println(v27)
                            var v31 int16 = 42
                            if true {
                                v5 -= -28
                                var v32 string = "hello"
                                _ = v32
                            } else {
                                fmt.Println(v20)
                                v33 := GenS3{F0: 1, F1: 42}
                                v27 += v27
                                if true {
                                    v34 := GenS1{F0: (v31 ^ (v31 >> ((65535) & 63)))}
                                    var v35 int = (fi28 * (fi30 + (^(fi28))))
                                    if false {
                                        switch (v35 >> ((18446744073709551615) & 63)) {
                                        case 127:
                                            fmt.Println(v27, v35, v20)
                                        default:
                                            var v36 float32 = (3.14159 - -2.25)
                                            _ = v36
                                        }
                                        v4 = 127
                                    }
                                    switch 32767 {
                                    case -153:
                                        v4 += (v5 | (v4 | v5))
                                    case -9223372036854775808:
                                        v37 := v1[65535]
                                        _ = v37
                                    case 0:
                                        if (42 > -9223372036854775808) {
                                            v38 := GenS4{F0: -1, F1: 42}
                                            _ = v38
                                        }
                                    default:
                                        var v39 bool = false
                                        _ = v39
                                    }
                                    var v40 int32 = 2147483647
                                    v41 := GenS1{F0: v31}
                                    v20 += 125
                                    fmt.Println(v6, fi28)
                                    if ((v40 < v40) && false) {
                                        v31 = v31
                                        _ = len(v3)
                                        var v42 int64 = -299
                                        _ = v42
                                    } else {
                                        switch 32768 {
                                        case -1:
                                            v2[(("a" + "\n\t") + "")] = 619
                                        case -535:
                                            v43 := []int{fi30}
                                            _ = v43
                                        }
                                        var v44 uint16 = v27
                                        v45 := GenS3{F0: 42, F1: fi28}
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = v34
                                    _ = v35
                                    _ = v40
                                    _ = v41
                                }
                                v46 := GenS5{F0: 1, F1: 42, F2: (fi30 - fi30)}
                                _ = v33
                                _ = v46
                            }
                            var v47 uint64 = 1
                            _ = v31
                            _ = v47
                        }
                        v48 := []int{-1, fi30}
                        _ = v48
                    } else {
                        var v49 float64 = -2.25
                        fmt.Println(v49, v49)
                        v4 = (v5 / ((127) | 1))
                        v50 := []int{fi30}
                        _ = v49
                        _ = v50
                    }
                    v6 = (v6 * 0.1)
                    _ = fi30
                }
            }
            var v51 float64 = float64(0.0)
            _ = len(v3)
            _ = fi28
            _ = v51
        }
        fmt.Println(v20, v4)
    }
    v52 := GenS6{F0: v27, F1: 318, F2: true}
    fmt.Println(v6)
    if true {
        switch "hello" {
        case "foo bar":
            for fi53 := 0; fi53 < 20; fi53++ {
                v54 := []int{fi53}
                _ = fi53
                _ = v54
            }
        }
        v55 := GenS7{F0: 127, F1: 92, F2: ("foo bar" + "hello")}
        if (1 >= -1) {
            var v56 float32 = 1e10
            _ = v56
        }
        _ = v55
    } else {
        v57 := []int{}
        switch ("hello" + "hello") {
        case "":
            var v58 uint8 = 42
            _ = v58
        case "foo bar":
            var v59 int = 127
            _ = v59
        default:
            for fi60 := 0; fi60 < 16; fi60++ {
                switch "\n\t" {
                case "":
                    var v61 uint = 985
                    _ = v61
                }
                _ = fi60
            }
        }
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
