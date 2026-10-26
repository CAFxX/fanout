package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 int }
func genFn0(p0 int32) int64 {
    p0 = (p0 - (^(p0)))
    fmt.Println(p0)
    if (0 != 32767) {
        v4 := []int{-2147483648, 2147483647, 127}
        _ = v4
    } else {
        if false {
            var v5 bool = true
            v6 := []int{}
            _ = v5
            _ = v6
        }
        fmt.Println(p0)
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: -1}
        var v9 float64 = 0.0
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    return -808
}

func main() {
    var v1 int8 = -102
    var v2 bool = (true || false)
    var v3 int16 = 1000
    v10 := []int{-409, 794, -95}
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 float32 = (3.14159 - 1e10)
        _ = fi11
        _ = v12
    }
    v1 -= v1
    v1 = v1
    if v2 {
        if len(v10) > 0 {
            v10[0] = 345
        }
        _ = len(v10)
        v13 := GenS2{F0: -125, F1: 4294967295}
        _ = v13
    } else {
        switch -1 {
        case 1:
            if ((32767 != 255) && v2) {
                if len(v10) > 0 {
                    v10[0] = 42
                }
                var v14 uint32 = 773
                for fi15 := 0; fi15 < 5; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v3 += (v3 + v3)
                        v10 = append(v10, (fi15 ^ 9223372036854775807))
                        _ = fi16
                    }
                    var v17 uint8 = 128
                    var v18 int = 1
                    _ = fi15
                    _ = v17
                    _ = v18
                }
                _ = v14
            } else {
                v3 -= 42
            }
        case 4294967295:
            if v2 {
                v19 := GenS2{F0: 100, F1: 1}
                _ = v19
            }
        case 0:
            var v20 uint16 = 0
            _ = v20
        default:
            if ((1 == 2) || true) {
                if (!(v2 && v2)) {
                    _ = len(v10)
                    _ = len(v10)
                }
            } else {
                v21 := GenS3{F0: (^(v3)), F1: 195}
                _ = v21
            }
        }
        v22 := GenS4{F0: 128, F1: (v1 | (v1 & v1)), F2: 32767}
        if (1.5 != -2.25) {
            var v23 int32 = 2147483647
            v24 := GenS4{F0: 0, F1: v1, F2: 128}
            _ = v23
            _ = v24
        } else {
            if (2 == 2147483647) {
                v25 := GenS1{F0: (v1 * 78)}
                var v26 int64 = int64(v3)
                _ = v25
                _ = v26
            } else {
                fmt.Println(v2, v2, v3)
                for fi27 := 0; fi27 < 7; fi27++ {
                    switch -300 {
                    case 42:
                        switch "foo bar" {
                        case "a":
                            for fi28 := 0; fi28 < 9; fi28++ {
                                var v29 uint = 2
                                v30 := GenS4{F0: 127, F1: v1, F2: 963}
                                v1 += v1
                                if (v2 || false) {
                                    v2 = true
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    if (true || (42 >= 2)) {
                                        _ = len(v10)
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                } else {
                                    _ = len(v10)
                                    v31 := GenS5{F0: 32767}
                                    _ = v31
                                }
                                v2 = (255 > 1)
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                        case "foo bar":
                            v1 -= (-(v1))
                        case "\n\t":
                            v3 = (v3 >> ((2) & 63))
                        }
                    default:
                        v32 := GenS6{F0: 749}
                        _ = v32
                    }
                    _ = fi27
                }
                if len(v10) > 0 {
                    v10[0] = -9223372036854775808
                }
            }
        }
        _ = v22
    }
    v2 = false
    switch 32768 {
    case 127:
        fmt.Println(v3, v2, v3)
    default:
        if (true && v2) {
            fmt.Println(v1)
            var v33 int8 = -118
            _ = v33
        } else {
            v34 := GenS7{F0: (!(437 > 0))}
            var v35 float64 = 0.0
            _ = v34
            _ = v35
        }
    }
    if false {
        fmt.Println(v2)
        if (0.1 == float32(1.5)) {
            v36 := GenS8{F0: (v1 ^ 42), F1: 1, F2: 82}
            v37 := GenS6{F0: 0}
            for fi38 := 0; fi38 < 11; fi38++ {
                fmt.Println(v2, v2)
                v39 := GenS9{F0: 18446744073709551615, F1: 65535}
                fmt.Println(fi38, v3)
                _ = fi38
                _ = v39
            }
            _ = v36
            _ = v37
        } else {
            switch "hello" {
            case "hello":
                fmt.Println(v1, v1)
            }
            v1 = 127
            if ((!(42 >= 2)) && (!v2)) {
                v2 = ("" <= "\n\t")
                if (4294967295 < 904) {
                    v1 += 42
                    fmt.Println(v1, v2)
                } else {
                    _ = len(v10)
                    v3 -= -367
                }
            } else {
                v40 := GenS10{F0: 997}
                var v41 uint8 = 25
                var v42 int16 = v3
                _ = v40
                _ = v41
                _ = v42
            }
        }
    } else {
        v2 = (false && ((("foo bar" + "\n\t") + "\n\t") <= ((("foo bar" + "hello") + "\n\t") + "foo bar")))
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
