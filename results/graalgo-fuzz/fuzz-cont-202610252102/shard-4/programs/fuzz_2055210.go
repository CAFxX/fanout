package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
type GenS5 struct { F0 string; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int32 }
func genFn0() uint64 {
    var v7 uint32 = 2147483648
    if true {
        v7 = (^(v7))
    }
    _ = v7
    return 4294967295
}
func genFn1() int32 {
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 42, F1: 128, F2: 0}
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 42
    v11 := GenS2{F0: -485, F1: 42}
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint8 = 24
    var v5 float32 = 0.0
    v6 := make(map[uint32]int)
    var v12 uint = 0
    switch ("" + (("a" + "hello") + (("hello" + ("foo bar" + "a")) + ("hello" + "\n\t")))) {
    case "foo bar":
        v13 := GenS2{F0: 0, F1: 0}
        _ = v13
    case "\n\t":
        v14 := GenS1{F0: 695, F1: (^(v4)), F2: 0}
        _ = v14
    }
    if ("hello" <= "a") {
        if (1e10 >= 1.5) {
            v12 += v12
        } else {
            v5 = v5
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint64 = 572
                _ = fi15
                _ = v16
            }
            var v17 int = -345
            _ = v17
        }
    }
    v18 := GenS2{F0: 2147483647, F1: -1}
    switch 249 {
    case 458:
        switch (("hello" + ((("foo bar" + "hello") + ("\n\t" + "hello")) + "")) + "foo bar") {
        case "a":
            v12 = v12
        }
    case 4294967295:
        if (true || (!true)) {
            v4 += (v4 << (((v12 - (v12 * v12))) & 63))
            v19 := []int{}
            _ = v19
        } else {
            if true {
                v20 := GenS1{F0: 65535, F1: (v4 & v4), F2: 0}
                v4 += (v4 + (v4 * v4))
                _ = v20
            }
            for fi21 := 0; fi21 < 8; fi21++ {
                v5 = v5
                if (18446744073709551615 == uint64(v12)) {
                    v12 = v12
                    switch "a" {
                    case "":
                        fmt.Println(fi21)
                    case "hello":
                        var v22 float64 = -2.25
                        _ = v22
                    case "\n\t":
                        for fi23 := 0; fi23 < 9; fi23++ {
                            if (!false) {
                                var v24 float32 = (((v5 + float32(3.14159)) * 0.0) - v5)
                                _ = v24
                            }
                            v25 := GenS2{F0: 1, F1: 1}
                            _ = fi23
                            _ = v25
                        }
                    default:
                        var v26 int16 = 594
                        _ = v26
                    }
                }
                _ = fi21
            }
            var v27 float32 = ((v5 + (float32(1.5) * 3.14159)) - v5)
            _ = v27
        }
    }
    v4 = (^(v4))
    if (-1 < 128) {
        switch (("\n\t" + "\n\t") + ("\n\t" + "\n\t")) {
        case "":
            if (true || true) {
                for fi28 := 0; fi28 < 2; fi28++ {
                    fmt.Println(v4)
                    for fi29 := 0; fi29 < 4; fi29++ {
                        v5 += v5
                        _ = fi29
                    }
                    if ((((3.14159 + float64(float32(1e10))) + (0.0 - 0.1)) >= -2.25) || (((v12 % ((v12) | 1)) > v12) && true)) {
                        var v30 uint64 = 2
                        _ = v30
                    } else {
                        switch fi28 {
                        case 4294967295:
                            if (false || ((343 < 914) && false)) {
                                v31 := []int{(-(fi28)), (fi28 - (fi28 ^ int(v4)))}
                                v12 = (v12 % (((^(v12))) | 1))
                                var v32 uint64 = 42
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(v5, v12, v4)
                            }
                        case 2147483647:
                            var v33 float32 = v5
                            _ = v33
                        default:
                            v34 := GenS1{F0: 1, F1: 45, F2: 42}
                            _ = v34
                        }
                        fmt.Println(v5, v5, v5)
                        var v35 float32 = (float32((0.0 + 0.0)) / -2.25)
                        _ = v35
                    }
                    _ = fi28
                }
                var v36 uint = v12
                _ = v36
            } else {
                if (182 <= 32767) {
                    var v37 uint32 = 0
                    fmt.Println(v37)
                    fmt.Println(v37)
                    _ = v37
                } else {
                    v4 += (v4 + (v4 + 1))
                    var v38 string = (("hello" + "") + ("a" + (("foo bar" + ("a" + "foo bar")) + "\n\t")))
                    _ = v38
                }
                var v39 uint = 4294967295
                var v40 int16 = 32767
                _ = v39
                _ = v40
            }
        case "hello":
            v5 = 1e10
        default:
            var v41 float32 = float32(float64((3.14159 + v5)))
            _ = v41
        }
        v42 := []int{9223372036854775807}
        _ = v42
    } else {
        var v43 uint16 = 58
        delete(v3, -572)
        _ = len(v6)
        _ = v43
    }
    v44 := GenS3{F0: 9223372036854775807, F1: 0}
    var v45 int16 = 1
    var v46 int8 = -26
    v47 := GenS4{F0: "", F1: 2147483648, F2: 0}
    var v48 int8 = v46
    if true {
        if true {
            var v49 uint = v12
            _ = v49
        }
    }
    if (2147483647 > -2147483648) {
        v50 := []int{-226, 1, -747, -898}
        if false {
            if ((true || true) || false) {
                if true {
                    v51 := GenS5{F0: "\n\t", F1: 274, F2: (^(v4))}
                    _ = v51
                }
            }
            v3[114] = 0
            _ = len(v3)
        }
        v52 := GenS6{F0: 42}
        _ = v50
        _ = v52
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
