package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0() uint8 {
    v6 := []int{4294967295, -2147483648, 564}
    var v7 uint32 = 281
    fmt.Println(v7)
    if len(v6) > 0 {
        v6[0] = 330
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 uint16) uint32 {
    fmt.Println(p0, p0)
    if ((p0 + uint64(p1)) != p0) {
        v8 := GenS1{F0: 0}
        fmt.Println(p0, p0)
        _ = v8
    }
    switch "" {
    case "a":
        switch (((("a" + ("foo bar" + ((("a" + "foo bar") + "foo bar") + (("hello" + "hello") + "a")))) + "hello") + ("" + ("hello" + ("hello" + ("\n\t" + ""))))) + ((("" + "\n\t") + "a") + (("a" + ("hello" + "")) + "hello"))) {
        case "hello":
            if (178 >= 2) {
                for fi9 := 0; fi9 < 12; fi9++ {
                    fmt.Println(p0, p0, fi9)
                    _ = fi9
                }
                var v10 uint = 1
                _ = v10
            } else {
                var v11 uint8 = 128
                for fi12 := 0; fi12 < 13; fi12++ {
                    if ((^(p0)) <= 354) {
                        v13 := []int{(fi12 << ((1) & 63))}
                        _ = v13
                    } else {
                        var v14 int16 = 0
                        for fi15 := 0; fi15 < 9; fi15++ {
                            v11 -= (v11 | v11)
                            _ = fi15
                        }
                        _ = v14
                    }
                    _ = fi12
                }
                if true {
                    var v16 int32 = -1
                    switch "hello" {
                    case "hello":
                        if (true || (v16 != v16)) {
                            p0 -= 2
                        }
                    case "a":
                        var v17 uint64 = p0
                        _ = v17
                    case "":
                        v18 := []int{-944, 0, int(v16), 32767}
                        _ = v18
                    default:
                        var v19 uint8 = 42
                        _ = v19
                    }
                    _ = v16
                } else {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        fmt.Println(fi20, fi20)
                        var v21 uint32 = 2147483647
                        var v22 int = -416
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    switch -2147483648 {
                    case 1:
                        for fi23 := 0; fi23 < 16; fi23++ {
                            if false {
                                v24 := []int{}
                                _ = v24
                            }
                            _ = fi23
                        }
                    }
                    fmt.Println(v11, p0)
                }
                _ = v11
            }
        default:
            p1 = (p1 << ((65535) & 63))
        }
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        v25 := []int{}
        _ = v25
    }
    if true {
        p1 += (p1 * (p1 | p1))
    } else {
        v26 := []int{-9223372036854775808, 2147483647}
        if false {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            p1 += (p1 & p1)
            var v27 int = -2147483648
            _ = v27
        }
        if true {
            var v28 int16 = -677
            for fi29 := 0; fi29 < 20; fi29++ {
                v26 = append(v26, 127)
                var v30 int8 = 88
                v28 = (v28 | 42)
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        _ = v26
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int32 = -438
    v2 := make(map[uint32]bool)
    var v3 uint16 = 92
    var v4 int64 = 0
    var v5 int32 = (v1 - v1)
    if (v1 != (v5 + (v5 % ((v1) | 1)))) {
        var v31 int64 = (v4 << ((4294967295) & 63))
        v1 = v1
        _ = len(v2)
        _ = v31
    } else {
        var v32 int8 = -128
        _ = v32
    }
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "":
        for fi33 := 0; fi33 < 17; fi33++ {
            v3 = v3
            if (v3 == 0) {
                for fi34 := 0; fi34 < 12; fi34++ {
                    switch "foo bar" {
                    case "foo bar":
                        v35 := GenS2{F0: 161}
                        var v36 float32 = 0.1
                        var v37 float64 = ((0.1 / 3.14159) - 1e10)
                        if (true || ((!true) || true)) {
                            v4 += (v4 << ((469) & 63))
                            var v38 uint32 = 894
                            v2[v38] = false
                            _ = v38
                        } else {
                            if (false || true) {
                                var v39 int = (fi34 ^ (fi33 & 553))
                                v4 -= v4
                                _ = v39
                            } else {
                                v40 := GenS1{F0: 1}
                                fmt.Println(v37, fi33)
                                _ = v40
                            }
                            v41 := GenS1{F0: 127}
                            _ = v41
                        }
                        _ = v35
                        _ = v36
                        _ = v37
                    case "hello":
                        v42 := GenS3{F0: v1, F1: (!(v5 < (-(v5))))}
                        _ = v42
                    default:
                        v3 += 1
                    }
                    var v43 uint = 383
                    v2[uint32(v43)] = (447 != 4294967295)
                    _ = fi34
                    _ = v43
                }
            }
            _ = fi33
        }
    case "hello":
        v4 -= v4
    }
    if true {
        var v44 string = ""
        v45 := GenS4{F0: v3, F1: 127, F2: 127}
        _ = v44
        _ = v45
    } else {
        v5 = v1
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        for fi47 := 0; fi47 < 10; fi47++ {
            v48 := []int{307, int(v5)}
            v3 = (v3 ^ v3)
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    switch "foo bar" {
    case "a":
        var v49 uint = 42
        _ = v49
    case "":
        for fi50 := 0; fi50 < 20; fi50++ {
            var v51 uint = 0
            var v52 float64 = 0.1
            v53 := []int{fi50}
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
        }
    }
    v5 -= -459
    if ((v4 / ((0) | 1)) >= -868) {
        fmt.Println(v4, v4)
        if ((v4 * int64(v3)) == 679) {
            var v54 uint16 = (v3 | 730)
            var v55 uint = 918
            _ = v54
            _ = v55
        } else {
            if (true && ("\n\t" < "hello")) {
                v56 := []int{9223372036854775807, -464, 42}
                var v57 int16 = 1
                switch "hello" {
                case "":
                    v5 = (v5 + (v1 ^ -2147483648))
                }
                _ = v56
                _ = v57
            }
            var v58 int8 = -96
            _ = v58
        }
        var v59 uint8 = 2
        _ = v59
    } else {
        v5 = v5
    }
    v5 = v5
    v60 := GenS3{F0: v5, F1: ((!(true || false)) || true)}
    var v61 float32 = 3.14159
    v62 := []int{-961, -401, 32767, 9223372036854775807}
    if len(v62) > 0 {
        v62[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
}
