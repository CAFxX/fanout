package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int32; F2 int8 }
func genFn0(p0 uint16) int8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6, fi6)
        p0 += p0
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 bool = (42 > 32767)
    var v8 float32 = 1.5
    switch "\n\t" {
    case "\n\t":
        v9 := []int{32768, 32768, -610, 4294967295}
        _ = v9
    }
    if ((42 == 42) || false) {
        for fi10 := 0; fi10 < 7; fi10++ {
            if v7 {
                v11 := GenS1{F0: 0, F1: "foo bar", F2: -1}
                v8 += (v8 * (v8 * v8))
                switch (fi10 << ((528) & 63)) {
                case 128:
                    switch ("hello" + "a") {
                    case "hello":
                        var v12 int8 = 100
                        _ = v12
                    case "foo bar":
                        v8 = float32(((1.5 - -2.25) / (0.1 + (((1.5 / 0.1) / 1e10) * float64(v8)))))
                    }
                case -180:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v8 -= ((v8 + v8) - (v8 + float32((0.0 + 0.0))))
                        _ = fi13
                    }
                case 0:
                    switch ("foo bar" + "a") {
                    case "a":
                        for fi14 := 0; fi14 < 1; fi14++ {
                            if (1 == 2) {
                                v7 = (0 >= 128)
                                if ((!v7) || v7) {
                                    if v7 {
                                        if (-9223372036854775808 < 790) {
                                            v15 := GenS1{F0: 0, F1: "\n\t", F2: 322}
                                            v8 += v8
                                            v7 = (fi14 == fi14)
                                            _ = v15
                                        }
                                        if ((false && (111 < int32(fi14))) || v7) {
                                            if false {
                                                var v16 string = "a"
                                                v16 = ((("foo bar" + ((v16 + ((v16 + v16) + v16)) + "\n\t")) + v16) + "hello")
                                                v17 := []int{}
                                                _ = v16
                                                _ = v17
                                            }
                                            v7 = v7
                                            var v18 uint8 = 128
                                            _ = v18
                                        }
                                    } else {
                                        var v19 uint32 = 65535
                                        _ = v19
                                    }
                                    switch (("" + "a") + (("hello" + (("a" + "a") + "a")) + ("\n\t" + "hello"))) {
                                    case "\n\t":
                                        v8 = v8
                                    case "a":
                                        var v20 string = ("\n\t" + (("a" + "foo bar") + "a"))
                                        _ = v20
                                    case "foo bar":
                                        fmt.Println(fi10)
                                    default:
                                        fmt.Println(v8, fi14, fi14)
                                    }
                                    var v21 float32 = ((v8 - v8) - 0.0)
                                    v22 := GenS2{F0: -383, F1: 256}
                                    if (178 == 42) {
                                        var v23 uint = 1
                                        _ = v23
                                    } else {
                                        v24 := GenS3{F0: -648, F1: ("foo bar" != "a"), F2: 1}
                                        var v25 int32 = -1
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = v21
                                    _ = v22
                                }
                                v7 = v7
                                if v7 {
                                    v26 := GenS3{F0: -1, F1: false, F2: 42}
                                    _ = v26
                                }
                            }
                            fmt.Println(p0, fi14)
                            _ = fi14
                        }
                    default:
                        v8 = float32(1e10)
                    }
                }
                _ = v11
            }
            var v27 uint64 = 92
            _ = fi10
            _ = v27
        }
        v28 := GenS4{F0: ((p0 ^ (p0 * (p0 ^ p0))) > (p0 << ((4294967295) & 63)))}
        _ = v28
    }
    _ = p0
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int8, p1 uint16, p2 bool) int16 {
    v29 := GenS4{F0: (p2 || true)}
    v30 := []int{}
    if false {
        v31 := GenS3{F0: 1, F1: true, F2: 2147483648}
        var v32 uint8 = 127
        if (599 == 2147483647) {
            v30 = append(v30, 4294967295)
        } else {
            if (0.1 <= (0.1 - 1.5)) {
                switch "" {
                case "\n\t":
                    var v33 uint8 = (v32 * v32)
                    _ = v33
                case "a":
                    var v34 int32 = 2147483647
                    _ = v34
                case "":
                    fmt.Println(p0, v32)
                default:
                    v32 = v32
                }
                for fi35 := 0; fi35 < 8; fi35++ {
                    if false {
                        var v36 int32 = -1
                        var v37 uint = 42
                        _ = v36
                        _ = v37
                    }
                    _ = fi35
                }
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
            }
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
        }
        _ = v31
        _ = v32
    }
    p1 -= p1
    v38 := GenS4{F0: (!(1 == 32767))}
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    _ = v38
    return -227
}

func main() {
    v1 := []uint32{65535, 65535}
    v2 := make(map[uint32]int)
    var v3 float32 = -2.25
    v4 := make(map[string]bool)
    var v5 int8 = -72
    _ = len(v4)
    if ((v3 / v3) != v3) {
        var v39 uint32 = 42
        if true {
            v5 = 1
        }
        _ = v39
    }
    var v40 uint = 2
    fmt.Println(v40, v40)
    for fi41 := 0; fi41 < 6; fi41++ {
        v3 = (-2.25 - 3.14159)
        var v42 float64 = 1.5
        var v43 uint32 = 1
        _ = fi41
        _ = v42
        _ = v43
    }
    v44 := GenS3{F0: 943, F1: (false && true), F2: 65535}
    var v45 bool = true
    if v45 {
        v5 = (v5 * (v5 & v5))
        v5 = v5
    } else {
        _ = len(v1)
        fmt.Println(v40, v3)
    }
    v45 = true
    v46 := GenS5{F0: 1, F1: 42, F2: 57}
    var v47 int64 = 0
    _ = len(v4)
    var v48 int64 = v47
    _ = len(v4)
    var v49 uint64 = 18446744073709551615
    var v50 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
