package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 int64; F1 int16; F2 int8 }
type GenS6 struct { F0 int; F1 uint8 }
func genFn0(p0 string, p1 uint16) int {
    for fi4 := 0; fi4 < 2; fi4++ {
        if (((4294967295 <= 800) && true) || (!(676 >= 18446744073709551615))) {
            for fi5 := 0; fi5 < 4; fi5++ {
                v6 := GenS1{F0: 1000, F1: (p1 / ((42) | 1)), F2: -32768}
                var v7 int = (fi4 / ((fi5) | 1))
                _ = fi5
                _ = v6
                _ = v7
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int64 = -9223372036854775808
                p1 = (p1 >> ((2) & 63))
                v10 := GenS1{F0: -260, F1: p1, F2: -32768}
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(p0, p1)
        } else {
            p0 = (((("\n\t" + (p0 + (p0 + p0))) + "\n\t") + (p0 + "foo bar")) + "hello")
            if ((!(-32768 < -32768)) && false) {
                for fi11 := 0; fi11 < 8; fi11++ {
                    fmt.Println(p1, p0, p1)
                    v12 := []int{-686}
                    _ = fi11
                    _ = v12
                }
                v13 := []int{(fi4 | 4294967295), (fi4 & (-(fi4))), fi4, 42}
                _ = v13
            } else {
                p1 -= p1
            }
            var v14 uint64 = 4294967295
            _ = v14
        }
        if (-128 >= -107) {
            if (128 != 96) {
                switch p0 {
                case "":
                    fmt.Println(p0)
                case "a":
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v16 := []int{(fi15 * fi15), fi4}
                        _ = fi15
                        _ = v16
                    }
                default:
                    v17 := GenS1{F0: 32767, F1: (p1 >> ((856) & 63)), F2: 1}
                    _ = v17
                }
                v18 := GenS1{F0: 1000, F1: p1, F2: -467}
                _ = v18
            } else {
                var v19 bool = true
                _ = v19
            }
            p1 = (p1 | (^(p1)))
            if false {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := []int{fi20}
                    _ = fi20
                    _ = v21
                }
                p0 = "foo bar"
                if true {
                    if true {
                        var v22 int8 = 0
                        p1 = p1
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := []int{(fi4 | int(p1)), (fi4 / ((fi4) | 1)), (fi4 + (-(fi4)))}
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v26 := GenS1{F0: 32767, F1: (p1 / (((^(p1))) | 1)), F2: 118}
                        _ = fi25
                        _ = v26
                    }
                }
            } else {
                var v27 uint8 = 128
                var v28 bool = (1 <= 322)
                switch fi4 {
                case -1:
                    for fi29 := 0; fi29 < 15; fi29++ {
                        v30 := GenS2{F0: 2147483648, F1: 24}
                        _ = fi29
                        _ = v30
                    }
                case 32768:
                    fmt.Println(v27, v28)
                case -9223372036854775808:
                    var v31 uint32 = 255
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 1; fi32++ {
                        p0 = ((p0 + "foo bar") + p0)
                        if v28 {
                            var v33 uint = 255
                            var v34 string = "foo bar"
                            _ = v33
                            _ = v34
                        } else {
                            if ((p1 / ((p1) | 1)) == 0) {
                                fmt.Println(p1, fi32)
                                v35 := []int{fi32, fi4}
                                var v36 int64 = 1
                                var v37 uint16 = (p1 % ((p1) | 1))
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                p0 = "hello"
                            }
                            v38 := GenS2{F0: 2147483648, F1: 0}
                            _ = v38
                        }
                        _ = fi32
                    }
                }
                _ = v27
                _ = v28
            }
        } else {
            v39 := []int{fi4}
            fmt.Println(p0, fi4, p1)
            if len(v39) >= 2 {
                v39 = v39[1:len(v39)]
            }
            _ = v39
        }
        switch "hello" {
        case "foo bar":
            v40 := GenS3{F0: 9223372036854775807, F1: 99}
            _ = v40
        }
        _ = fi4
    }
    switch 127 {
    case -1:
        for fi41 := 0; fi41 < 5; fi41++ {
            v42 := []int{4294967295, fi41, (fi41 & 4294967295)}
            var v43 uint64 = 1
            _ = fi41
            _ = v42
            _ = v43
        }
    case -9223372036854775808:
        v44 := GenS3{F0: -2147483648, F1: 0}
        _ = v44
    default:
        var v45 int = 1
        _ = v45
    }
    switch int(p1) {
    case 32768:
        p1 = p1
    case -986:
        var v46 float32 = 1.5
        _ = v46
    case 128:
        p1 -= p1
    }
    p1 = 317
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 2147483647
    var v3 int64 = 1
    v47 := GenS3{F0: genFn0("", 815), F1: 100}
    if true {
        v48 := GenS3{F0: -771, F1: -1}
        for fi49 := 0; fi49 < 19; fi49++ {
            var v50 int64 = v3
            _ = fi49
            _ = v50
        }
        _ = v48
    } else {
        if true {
            v3 = (v3 - (-(v3)))
            v51 := []int{}
            _ = v51
        } else {
            var v52 float32 = ((0.0 * 1e10) - (0.0 * float32(3.14159)))
            _ = v52
        }
        fmt.Println(v2, v1)
    }
    if false {
        v2 -= (v2 - v2)
    } else {
        if false {
            v53 := GenS4{F0: 329, F1: v3, F2: (v3 / (((v3 >> ((265) & 63))) | 1))}
            v2 = v2
            v54 := GenS3{F0: -1, F1: 56}
            _ = v53
            _ = v54
        }
    }
    v3 = v3
    for fi55 := 0; fi55 < 17; fi55++ {
        v56 := GenS5{F0: v3, F1: 42, F2: 42}
        v57 := GenS6{F0: (fi55 & 58), F1: 155}
        v58 := []int{}
        _ = fi55
        _ = v56
        _ = v57
        _ = v58
    }
    v1 += -2.25
    v1 -= float64(1e10)
    v59 := GenS5{F0: v3, F1: -32768, F2: -1}
    var v60 int8 = -1
    v61 := []int{0, 1}
    fmt.Println(v60, v1, v60)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v59
    _ = v60
    _ = v61
}
