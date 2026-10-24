package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int16; F2 int }
type GenS5 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint32; F1 int }
func genFn0(p0 int32) int16 {
    v7 := []int{32768}
    if true {
        v8 := GenS1{F0: 0, F1: 890, F2: 100}
        p0 = p0
        _ = v8
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 string) uint8 {
    var v9 int32 = -500
    var v10 uint16 = 772
    v11 := GenS2{F0: 183}
    switch p1 {
    case "hello":
        fmt.Println(p1)
    default:
        var v12 int64 = 1
        _ = v12
    }
    if (!(0 != 679)) {
        if false {
            v13 := GenS3{F0: (110 < 2147483648)}
            _ = v13
        } else {
            v14 := []int{42, -877, 0}
            if (true || (p1 >= p1)) {
                var v15 float32 = 1e10
                fmt.Println(p1, v9, v10)
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 int32 = (v9 >> ((42) & 63))
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v17 += v9
                        v19 := GenS3{F0: false}
                        var v20 int16 = p0
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 8; fi21++ {
                        switch p1 {
                        case "":
                            v14 = append(v14, (-(fi21)))
                        case "foo bar":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                        _ = fi21
                    }
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            var v22 int16 = p0
            _ = v14
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            v24 := []int{fi23, fi23, fi23}
            _ = fi23
            _ = v24
        }
        p1 = p1
    } else {
        v25 := []int{127, 2147483647, -482, 127}
        switch 341 {
        case 32768:
            if true {
                if (!false) {
                    if true {
                        v9 -= (-(v9))
                    } else {
                        v26 := GenS4{F0: 18446744073709551615, F1: p0, F2: 9223372036854775807}
                        fmt.Println(v10, v10, p1)
                        p1 = p1
                        _ = v26
                    }
                }
                v27 := GenS4{F0: 2, F1: genFn0((v9 + 298)), F2: 4294967295}
                _ = v27
            } else {
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 float64 = 1e10
                    _ = fi28
                    _ = v29
                }
                v9 += (v9 / (((v9 - (v9 ^ (v9 - v9)))) | 1))
                v10 = (v10 + (v10 % ((42) | 1)))
            }
        case 0:
            v10 = v10
        default:
            var v30 int16 = -1
            _ = v30
        }
        fmt.Println(v9, v10)
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 147
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 uint16 = 245
    var v4 int64 = -557
    var v5 int16 = 32767
    v6 := make(map[string]bool)
    v5 -= genFn0(1)
    v31 := GenS5{F0: 13, F1: -128, F2: 682}
    if true {
        switch 127 {
        case -1:
            var v32 int8 = 100
            _ = v32
        case 128:
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := GenS6{F0: 42}
                _ = fi33
                _ = v34
            }
        }
        fmt.Println(v5, v2, v3)
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        v3 = (v3 | 770)
        if v2 {
            var v36 uint32 = 2147483648
            v37 := GenS3{F0: (!((195 == 1) && (!false)))}
            v2 = (genFn1((v5 * (^(v5))), "a") != 226)
            _ = v36
            _ = v37
        }
        delete(v6, "hello")
        _ = fi35
    }
    if false {
        var v38 bool = ((245 != 0) || false)
        v39 := GenS7{F0: 42, F1: 128}
        _ = len(v6)
        _ = v38
        _ = v39
    } else {
        v40 := GenS4{F0: 0, F1: 42, F2: 386}
        _ = v40
    }
    v3 = v3
    if (259 != 42) {
        for fi41 := 0; fi41 < 12; fi41++ {
            _ = len(v1)
            for fi42 := 0; fi42 < 4; fi42++ {
                var v43 uint8 = 128
                _ = fi42
                _ = v43
            }
            v44 := GenS8{F0: 42, F1: fi41}
            _ = fi41
            _ = v44
        }
        v2 = true
        if (!v2) {
            if (32768 != 127) {
                var v45 int32 = 42
                for fi46 := 0; fi46 < 20; fi46++ {
                    fmt.Println(v2)
                    _ = fi46
                }
                _ = v45
            } else {
                for fi47 := 0; fi47 < 15; fi47++ {
                    switch ("foo bar" + ("\n\t" + "")) {
                    case "foo bar":
                        var v48 float32 = 1e10
                        _ = v48
                    }
                    _ = fi47
                }
                var v49 int64 = 42
                for fi50 := 0; fi50 < 16; fi50++ {
                    if v2 {
                        for fi51 := 0; fi51 < 7; fi51++ {
                            v4 = -1
                            _ = fi51
                        }
                        fmt.Println(v4, fi50, v3)
                    }
                    _ = fi50
                }
                _ = v49
            }
            var v52 uint8 = uint8(v3)
            _ = v52
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            var v53 uint16 = (v3 | (v3 | (v3 >> ((564) & 63))))
            _ = v53
        case "hello":
            for fi54 := 0; fi54 < 17; fi54++ {
                if v2 {
                    if (2 > 0) {
                        switch fi54 {
                        case 741:
                            v3 += (v3 * 1)
                        }
                        var v55 float64 = 3.14159
                        _ = v55
                    }
                }
                _ = fi54
            }
        case "foo bar":
            var v56 uint64 = 579
            _ = v56
        }
        for fi57 := 0; fi57 < 15; fi57++ {
            fmt.Println(fi57)
            fmt.Println(v4)
            _ = fi57
        }
    }
    if ((v2 || false) || (!v2)) {
        if v2 {
            for fi58 := 0; fi58 < 8; fi58++ {
                v59 := []int{int(v3), (fi58 << ((654) & 63)), 684, fi58}
                v3 -= 0
                _ = fi58
                _ = v59
            }
            fmt.Println(v2, v5, v4)
        }
    }
    fmt.Println(v5)
    v6[""] = (v2 && v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
