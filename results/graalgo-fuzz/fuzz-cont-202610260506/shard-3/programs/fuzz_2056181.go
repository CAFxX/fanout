package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint32 }
func genFn0() int16 {
    if false {
        var v7 uint = 2
        _ = v7
    }
    var v8 uint8 = 42
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{}
        v11 := GenS1{F0: fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint8 = (v8 + (v8 | (^(v8))))
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            if (-2.25 >= 1.5) {
                if true {
                    if (!false) {
                        if (1 >= 1) {
                            v8 -= v12
                            v15 := []int{}
                            _ = v15
                        } else {
                            if (true && false) {
                                var v16 float32 = 0.0
                                var v17 bool = (1 > -32768)
                                var v18 uint8 = (v8 << ((399) & 63))
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            v12 += (v8 | 1)
                            var v19 int16 = 1
                            _ = v19
                        }
                    }
                    var v20 int8 = 0
                    switch 127 {
                    case 42:
                        switch "a" {
                        case "hello":
                            var v21 uint = 814
                            _ = v21
                        case "a":
                            v22 := []int{}
                            _ = v22
                        default:
                            if (!(!(true || true))) {
                                v8 = (v8 ^ (^(v12)))
                                v23 := []int{fi14, (-(fi13)), fi14, fi13}
                                _ = v23
                            } else {
                                if (65535 <= 42) {
                                    v24 := []int{2147483647}
                                    v25 := GenS2{F0: 0, F1: -277}
                                    fmt.Println(v20)
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint = 18446744073709551615
                                    v27 := []int{-159, fi14, fi14, fi13}
                                    switch fi14 {
                                    case 4294967295:
                                        if len(v27) > 0 {
                                            v27[0] = (-(fi14))
                                        }
                                    }
                                    var v28 uint8 = (v12 ^ (v8 * v8))
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                if true {
                                    if false {
                                        var v29 uint16 = 2
                                        v30 := []int{(fi14 << ((0) & 63)), fi13, fi13, 127}
                                        var v31 int32 = 2147483647
                                        v8 = (v12 << ((327) & 63))
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    v12 = v12
                                    if (612 >= 0) {
                                        switch ((("\n\t" + "a") + "foo bar") + ("foo bar" + "")) {
                                        case "hello":
                                            var v32 bool = ("a" <= "\n\t")
                                            _ = v32
                                        }
                                        v12 = v8
                                    } else {
                                        v33 := GenS2{F0: 42, F1: 42}
                                        if (false || (!false)) {
                                            var v34 int16 = 1000
                                            _ = v34
                                        } else {
                                            if ((^(v8)) < v8) {
                                                switch (fi14 - 4294967295) {
                                                case -725:
                                                    var v35 uint = 0
                                                    _ = v35
                                                }
                                                v8 += v12
                                            }
                                            v36 := []int{}
                                            var v37 uint32 = 2147483648
                                            _ = v36
                                            _ = v37
                                        }
                                        fmt.Println(v12)
                                        _ = v33
                                    }
                                } else {
                                    v38 := GenS1{F0: fi14}
                                    _ = v38
                                }
                            }
                        }
                    case -200:
                        fmt.Println(v12, v12, v8)
                    case -839:
                        switch (fi13 * fi14) {
                        case 2147483647:
                            v39 := []int{}
                            _ = v39
                        case 635:
                            v20 = (^(v20))
                            v12 = v8
                        case -164:
                            if true {
                                v20 = (v20 ^ 8)
                            }
                        default:
                            if false {
                                if (false && true) {
                                    if (42 > 2) {
                                        v40 := []int{696, (^(fi13))}
                                        v40 = append(v40, fi13)
                                        var v41 uint = 42
                                        v42 := GenS3{F0: v20, F1: 588, F2: true}
                                        var v43 string = (("a" + "a") + "")
                                        var v44 uint32 = 42
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        switch (("hello" + ("" + "\n\t")) + ("hello" + ("foo bar" + "foo bar"))) {
                                        case "\n\t":
                                            var v45 uint32 = 65535
                                            _ = v45
                                        case "foo bar":
                                            v46 := []int{(fi13 >> ((uint(v12)) & 63)), fi14, (fi14 >> ((65535) & 63))}
                                            _ = v46
                                        }
                                    }
                                }
                                if (214 <= 903) {
                                    var v47 string = ("foo bar" + "")
                                    _ = v47
                                }
                            }
                        }
                    default:
                        v8 += 255
                    }
                    v48 := []int{-191, fi14}
                    var v49 string = ""
                    _ = v20
                    _ = v48
                    _ = v49
                } else {
                    var v50 uint64 = 42
                    _ = v50
                }
            }
            _ = fi14
        }
        _ = fi13
    }
    _ = v12
    _ = v8
    return 0
}

func main() {
    var v1 bool = true
    var v2 uint8 = 83
    var v3 float32 = (0.0 + float32(1.5))
    var v4 int8 = -1
    v5 := make(map[string]uint32)
    var v6 int64 = 42
    var v51 float32 = (v3 * v3)
    var v52 uint32 = 310
    var v53 string = "a"
    v54 := GenS2{F0: 1, F1: -2147483648}
    var v55 int64 = v6
    if v1 {
        v56 := []int{1, 9223372036854775807, 42, -134}
        _ = v56
    }
    if v1 {
        fmt.Println(v3, v1)
    } else {
        var v57 bool = (771 < 478)
        _ = v57
    }
    var v58 int = -9223372036854775808
    if v1 {
        v1 = v1
        v59 := GenS4{F0: -2147483648, F1: -430, F2: v52}
        _ = v59
    } else {
        v60 := GenS1{F0: 42}
        _ = v60
    }
    v4 += (v4 & v4)
    for fi61 := 0; fi61 < 12; fi61++ {
        fmt.Println(v6)
        _ = fi61
    }
    v62 := GenS5{F0: 1000, F1: 1, F2: v52}
    if v1 {
        v2 += v2
        v63 := GenS2{F0: 255, F1: 42}
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v58
    _ = v6
    _ = v62
}
