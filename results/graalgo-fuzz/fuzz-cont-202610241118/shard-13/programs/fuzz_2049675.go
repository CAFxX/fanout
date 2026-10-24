package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 int; F1 int16 }
func genFn0() bool {
    if (42 > -85) {
        for fi5 := 0; fi5 < 14; fi5++ {
            if ((("a" + "hello") + (("\n\t" + "") + "a")) == "foo bar") {
                if false {
                    v6 := []int{34}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v7 uint = 42
                    _ = v6
                    _ = v7
                }
            }
            _ = fi5
        }
        v8 := GenS1{F0: 128, F1: -1, F2: 2147483648}
        _ = v8
    } else {
        var v9 string = "\n\t"
        v9 = "a"
        v10 := []int{-9223372036854775808, 0, 42, 32767}
        _ = v10
        _ = v9
    }
    if (!false) {
        if true {
            var v11 uint16 = 2
            _ = v11
        }
        switch "" {
        case "a":
            var v12 uint64 = 341
            _ = v12
        }
        if (!((-424 <= -9223372036854775808) && (-2147483648 != 32768))) {
            var v13 bool = (42 > 127)
            _ = v13
        }
    } else {
        var v14 int = 4294967295
        v14 = 32768
        if (!((true || ("hello" >= "a")) && false)) {
            fmt.Println(v14)
            if false {
                fmt.Println(v14)
                v15 := []int{(v14 >> ((1) & 63)), (^(v14)), (v14 - v14)}
                _ = len(v15)
                _ = v15
            }
            fmt.Println(v14)
        } else {
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 float64 = (1.5 / ((3.14159 * float64(-2.25)) + (-2.25 * (0.0 / -2.25))))
                if (!(!(1 != 642))) {
                    var v18 int64 = 9223372036854775807
                    if true {
                        v19 := []int{(fi16 - 42), v14}
                        switch v14 {
                        case 32767:
                            switch "" {
                            case "hello":
                                var v20 bool = (61 >= 0)
                                _ = v20
                            case "\n\t":
                                v18 += (^(v18))
                            default:
                                var v21 uint = 65535
                                _ = v21
                            }
                        case 693:
                            if (-9 > 1) {
                                var v22 float32 = -2.25
                                _ = v22
                            } else {
                                _ = len(v19)
                                var v23 float32 = 3.14159
                                _ = v23
                            }
                        case 407:
                            v24 := GenS1{F0: 143, F1: 1, F2: 1}
                            _ = v24
                        default:
                            for fi25 := 0; fi25 < 19; fi25++ {
                                v14 += fi25
                                _ = fi25
                            }
                        }
                        var v26 string = "hello"
                        _ = v19
                        _ = v26
                    } else {
                        v17 = (v17 - float64(0.1))
                        v27 := []int{}
                        v18 += 58
                        _ = v27
                    }
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        }
        _ = v14
    }
    v28 := GenS2{F0: 1, F1: 11, F2: 0}
    v29 := GenS3{F0: 852, F1: 4294967295, F2: false}
    _ = v28
    _ = v29
    return true
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{946, 4294967295, 2147483648, 4294967295}
    var v3 uint16 = 42
    var v4 uint16 = (v3 % (((v3 | (v3 + 1))) | 1))
    if len(v2) > 0 {
        v2[0] = 65535
    }
    if (1 != -9223372036854775808) {
        var v30 int = -9223372036854775808
        _ = v30
    } else {
        if false {
            v2 = append(v2, 42)
        } else {
            var v31 uint = 958
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            v1 += 1
            v4 -= (v4 + v3)
            _ = fi33
        }
        var v34 int64 = int64(fi32)
        v4 += v4
        _ = fi32
        _ = v34
    }
    v4 = v4
    var v35 int16 = -219
    for fi36 := 0; fi36 < 10; fi36++ {
        fmt.Println(fi36, v4, fi36)
        if false {
            v4 += v3
            for fi37 := 0; fi37 < 16; fi37++ {
                v4 += v4
                _ = fi37
            }
        } else {
            var v38 int8 = (-(v1))
            v35 += v35
            switch 32768 {
            case 555:
                var v39 bool = false
                _ = v39
            case 32767:
                switch (("" + "") + ("a" + "\n\t")) {
                case "foo bar":
                    v1 = (-(v1))
                case "":
                    v4 -= (v3 * v4)
                }
            default:
                if ((0.0 - -2.25) < 0.1) {
                    var v40 uint32 = 2147483648
                    for fi41 := 0; fi41 < 16; fi41++ {
                        var v42 int8 = v1
                        _ = fi41
                        _ = v42
                    }
                    v35 = 1000
                    _ = v40
                } else {
                    if true {
                        v35 += (v35 + (-(v35)))
                        switch fi36 {
                        case 2147483647:
                            fmt.Println(v4)
                        case 1:
                            v43 := GenS4{F0: 1, F1: 1}
                            _ = v43
                        }
                    }
                    var v44 int8 = (-(v38))
                    v2 = append(v2, 2)
                    _ = v44
                }
            }
            _ = v38
        }
        _ = fi36
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        for fi46 := 0; fi46 < 12; fi46++ {
            if (67 != v1) {
                switch "" {
                case "\n\t":
                    fmt.Println(v4, v35)
                case "":
                    fmt.Println(fi45, v4)
                case "foo bar":
                    if ((0.1 + 0.1) > (1.5 / ((1.5 * 1e10) * float64(-2.25)))) {
                        fmt.Println(fi46)
                        v47 := GenS4{F0: fi45, F1: v35}
                        _ = v47
                    }
                }
                var v48 uint16 = (v4 & (^(v3)))
                var v49 int = (fi46 | (fi45 * (fi46 << ((403) & 63))))
                fmt.Println(v49, v1, fi45)
                var v50 int = 2147483647
                var v51 uint16 = v4
                v2 = append(v2, 65535)
                _ = v48
                _ = v49
                _ = v50
                _ = v51
            }
            _ = fi46
        }
        _ = fi45
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v52 := GenS2{F0: 42, F1: (-(v35)), F2: 2}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v52
}
