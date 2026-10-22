package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint; F2 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
func genFn0(p0 int64, p1 uint, p2 uint32) uint32 {
    v5 := []int{}
    var v6 int16 = 32767
    var v7 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (^(p2))
}
func genFn1(p0 int64, p1 uint8, p2 bool) uint {
    var v8 uint = 729
    v9 := []int{-594}
    p0 = -469
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]uint32)
    var v3 string = "foo bar"
    var v4 int = 32767
    v1 = v1
    if (false && ((v1 % ((v1) | 1)) < v1)) {
        v10 := []int{42}
        var v11 int64 = 0
        _ = v10
        _ = v11
    }
    if (2 != 42) {
        if ((-2.25 - 0.0) > 3.14159) {
            if true {
                fmt.Println(v1, v1)
                fmt.Println(v4, v3, v1)
                v12 := GenS1{F0: 4294967295}
                _ = v12
            } else {
                v4 = (v4 % (((v4 / ((1) | 1))) | 1))
                v3 = "foo bar"
            }
        }
        var v13 uint = 0
        _ = v13
    } else {
        var v14 int16 = 1
        if false {
            switch (v4 ^ v4) {
            case 2147483647:
                fmt.Println(v4, v1)
            case 4294967295:
                if false {
                    v15 := GenS1{F0: 255}
                    v1 -= v1
                    v3 = v3
                    _ = v15
                } else {
                    fmt.Println(v1, v14, v14)
                    var v16 int64 = -9223372036854775808
                    if (!((42 > 2) && (true && (816 > 40)))) {
                        var v17 string = v3
                        if (v16 == v16) {
                            switch v4 {
                            case 32767:
                                fmt.Println(v3, v3, v3)
                            case 241:
                                var v18 uint16 = 620
                                _ = v18
                            default:
                                v16 += v16
                            }
                            v4 = v4
                        } else {
                            if (128 <= 223) {
                                var v19 float64 = -2.25
                                var v20 uint = 4294967295
                                _ = len(v2)
                                _ = v19
                                _ = v20
                            } else {
                                if (32767 > 835) {
                                    if (128 != 57) {
                                        var v21 string = (v3 + ("a" + "a"))
                                        _ = v21
                                    } else {
                                        fmt.Println(v3)
                                        if (v3 > ("" + ("\n\t" + v17))) {
                                            delete(v2, "hello")
                                            v22 := []int{(-(v4)), (v4 + v4), v4, v4}
                                            _ = v22
                                        } else {
                                            v4 = (-(v4))
                                        }
                                    }
                                } else {
                                    var v23 uint64 = 984
                                    var v24 string = ("" + "")
                                    for fi25 := 0; fi25 < 4; fi25++ {
                                        delete(v2, v24)
                                        if false {
                                            v26 := []int{}
                                            v3 = v17
                                            v27 := GenS2{F0: (v1 % ((v1) | 1)), F1: 419, F2: v14}
                                            _ = v26
                                            _ = v27
                                        }
                                        _ = fi25
                                    }
                                    _ = v23
                                    _ = v24
                                }
                                delete(v2, v3)
                            }
                        }
                        if (4294967295 >= 18446744073709551615) {
                            fmt.Println(v4, v17, v1)
                            v28 := []int{}
                            if len(v28) > 0 {
                                v28[0] = (-(v4))
                            }
                            _ = v28
                        }
                        _ = v17
                    } else {
                        v29 := []int{}
                        if len(v29) > 0 {
                            v29[0] = 265
                        }
                        v2[v3] = 774
                        _ = v29
                    }
                    _ = v16
                }
            case 32767:
                v30 := []int{}
                _ = v30
            default:
                if (!true) {
                    var v31 int64 = 9223372036854775807
                    for fi32 := 0; fi32 < 16; fi32++ {
                        fmt.Println(v31, v4, fi32)
                        for fi33 := 0; fi33 < 2; fi33++ {
                            _ = len(v2)
                            _ = fi33
                        }
                        _ = fi32
                    }
                    _ = v31
                }
            }
            switch "foo bar" {
            case "\n\t":
                var v34 uint8 = 255
                _ = v34
            case "":
                v1 += -86
            }
        } else {
            var v35 int8 = (v1 % (((v1 - int8(v4))) | 1))
            _ = v35
        }
        _ = v14
    }
    if true {
        if (42 != genFn1(-1, 50, true)) {
            if (73 >= 42) {
                var v36 uint32 = 0
                v3 = "a"
                _ = v36
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                fmt.Println(v1)
                _ = fi37
            }
        }
        fmt.Println(v3, v1, v4)
    }
    v38 := GenS3{F0: -1, F1: ((("hello" + "\n\t") + v3) + v3)}
    delete(v2, v3)
    for fi39 := 0; fi39 < 11; fi39++ {
        v40 := GenS4{F0: 42}
        var v41 float64 = 0.0
        _ = fi39
        _ = v40
        _ = v41
    }
    switch ((v3 + "foo bar") + (v3 + "")) {
    case "":
        for fi42 := 0; fi42 < 13; fi42++ {
            switch fi42 {
            case -2147483648:
                var v43 bool = (-2.25 > 0.1)
                _ = v43
            }
            v4 += fi42
            v44 := GenS5{F0: 18446744073709551615}
            _ = fi42
            _ = v44
        }
    case "\n\t":
        fmt.Println(v3, v1, v3)
    }
    v45 := GenS2{F0: v1, F1: 423, F2: 32767}
    if (-9223372036854775808 < 0) {
        v46 := []int{(v4 * v4), (v4 - (v4 * (v4 & (v4 >> ((528) & 63)))))}
        if true {
            if len(v46) > 0 {
                v46[0] = v4
            }
            var v47 uint8 = 128
            _ = v47
        }
        var v48 bool = (0 > 0)
        _ = v46
        _ = v48
    } else {
        if false {
            v4 = 127
        }
        if false {
            v3 = ("\n\t" + (v3 + "hello"))
        } else {
            var v49 int8 = v1
            _ = v49
        }
        var v50 uint = genFn1(-9223372036854775808, 66, (false || false))
        _ = v50
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        switch (("a" + v3) + "foo bar") {
        case "a":
            var v52 int = (fi51 % ((fi51) | 1))
            _ = v52
        }
        v53 := []int{fi51, v4}
        _ = fi51
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v45
}
