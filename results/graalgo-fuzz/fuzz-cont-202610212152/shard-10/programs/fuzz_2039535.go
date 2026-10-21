package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint; F1 int; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint64, p2 uint64) uint32 {
    v4 := []int{-352}
    fmt.Println(p1, p2, p1)
    fmt.Println(p1, p1)
    var v5 float64 = (((3.14159 + 3.14159) / 1.5) + (float64(3.14159) + ((-2.25 * 0.1) - 0.0)))
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 65535
}
func genFn1() int16 {
    var v6 uint = 18446744073709551615
    v6 = (v6 << ((v6) & 63))
    v6 = 65535
    if true {
        v6 = v6
        var v7 int16 = 42
        _ = v7
    } else {
        if (-17 >= 491) {
            v6 += (v6 << ((v6) & 63))
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{(fi8 >> ((v6) & 63)), fi8, (^(fi8)), 1}
                v10 := GenS1{F0: 609, F1: 42, F2: 679}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v6 -= v6
        }
        v11 := GenS1{F0: v6, F1: 0, F2: 42}
        if true {
            v6 -= (v6 >> ((v6) & 63))
        }
        _ = v11
    }
    v12 := []int{1, 32768, 741}
    _ = v12
    _ = v6
    return 1000
}

func main() {
    var v1 uint16 = 1
    var v2 float64 = (float64((0.0 * 0.0)) * 0.0)
    var v3 bool = false
    v13 := []int{}
    v14 := GenS2{F0: v3, F1: 2}
    v15 := GenS3{F0: 0, F1: ("hello" == "foo bar")}
    if v3 {
        _ = len(v13)
        var v16 int = 0
        _ = v16
    } else {
        _ = len(v13)
        _ = len(v13)
        v13 = append(v13, 42)
    }
    var v17 int8 = 0
    if v3 {
        v2 = v2
    } else {
        var v18 int64 = -9223372036854775808
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: 9223372036854775807, F2: v1}
    var v20 int32 = -2147483648
    v21 := GenS1{F0: 255, F1: 1, F2: v1}
    var v22 float32 = (0.0 / (1e10 + ((0.0 + -2.25) / -2.25)))
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(fi23, v1)
        v2 -= 0.0
        _ = fi23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v2 -= (v2 - (0.0 + v2))
        var v25 int16 = 1000
        var v26 uint8 = 238
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        fmt.Println(v3, v3, v3)
        v22 = 0.0
        if v3 {
            v1 = 255
        } else {
            var v28 float64 = v2
            v22 = v22
            for fi29 := 0; fi29 < 18; fi29++ {
                if (!v3) {
                    v2 = v2
                    _ = len(v13)
                }
                var v30 float32 = v22
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        _ = fi27
    }
    v17 += (v17 << ((92) & 63))
    if (42 < (v17 ^ 100)) {
        switch (("foo bar" + "a") + "") {
        case "hello":
            v31 := GenS4{F0: 18446744073709551615, F1: 630, F2: 713}
            _ = v31
        case "\n\t":
            v17 -= 42
        }
    } else {
        v3 = (0 > -9223372036854775808)
        v22 = v22
        for fi32 := 0; fi32 < 9; fi32++ {
            fmt.Println(fi32, v17)
            if v3 {
                var v33 int16 = 80
                var v34 int32 = -480
                _ = v33
                _ = v34
            } else {
                if false {
                    v22 -= (0.0 / 3.14159)
                    if (-1 == v20) {
                        v35 := GenS3{F0: 4294967295, F1: (((v22 * 3.14159) > v22) || (true && (v3 || false)))}
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 11; fi36++ {
                        v2 -= v2
                        fmt.Println(v20)
                        _ = fi36
                    }
                } else {
                    for fi37 := 0; fi37 < 18; fi37++ {
                        if v3 {
                            v17 = v17
                            v38 := GenS3{F0: genFn0(1, 758, 0), F1: (v3 && (v3 && true))}
                            var v39 int32 = (v20 >> ((65535) & 63))
                            _ = v38
                            _ = v39
                        }
                        if len(v13) > 0 {
                            v13[0] = (fi32 + (fi32 * (fi37 << ((551) & 63))))
                        }
                        switch fi32 {
                        case 42:
                            var v40 int8 = int8(fi32)
                            _ = v40
                        case 9223372036854775807:
                            var v41 int64 = 563
                            _ = v41
                        case 127:
                            v42 := GenS5{F0: ("a" + "hello")}
                            _ = v42
                        default:
                            v3 = false
                        }
                        _ = fi37
                    }
                    switch (fi32 >> ((667) & 63)) {
                    case 482:
                        var v43 uint = 0
                        _ = v43
                    case -9223372036854775808:
                        switch (fi32 + (fi32 ^ fi32)) {
                        case -906:
                            v17 = (-(v17))
                        case -317:
                            v22 -= (float32(v2) * (v22 - 0.0))
                        default:
                            if (v2 > (v2 + v2)) {
                                var v44 string = ""
                                _ = v44
                            }
                        }
                    case 516:
                        v2 = 1.5
                    default:
                        v45 := GenS6{F0: 2, F1: v3, F2: 2}
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 7; fi46++ {
                        v22 -= (3.14159 - 3.14159)
                        _ = fi46
                    }
                }
                v3 = v3
            }
            var v47 int8 = v17
            _ = fi32
            _ = v47
        }
    }
    v48 := GenS7{F0: 42}
    if (2 != 42) {
        for fi49 := 0; fi49 < 4; fi49++ {
            fmt.Println(fi49)
            var v50 float32 = float32(float64(float32((v2 * (v2 / 0.1)))))
            _ = fi49
            _ = v50
        }
        var v51 int = 32767
        if (!(127 > 0)) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v51 = v51
        } else {
            v20 -= v20
            var v52 float64 = float64(v22)
            _ = v52
        }
        _ = v51
    }
    var v53 uint32 = 711
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v48
    _ = v53
}
