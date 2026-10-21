package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int; F2 int64 }
type GenS7 struct { F0 string; F1 int32; F2 int8 }
func genFn0(p0 uint64, p1 uint, p2 int32) uint8 {
    p0 -= p0
    switch "" {
    case "":
        for fi4 := 0; fi4 < 13; fi4++ {
            v5 := []int{fi4, (fi4 & (fi4 >> (((p1 >> ((p1) & 63))) & 63)))}
            if false {
                for fi6 := 0; fi6 < 1; fi6++ {
                    var v7 bool = true
                    if v7 {
                        v8 := GenS1{F0: ((-(p2)) < p2), F1: p2, F2: -289}
                        _ = v8
                    } else {
                        v9 := GenS1{F0: v7, F1: -1, F2: -9223372036854775808}
                        var v10 float64 = 3.14159
                        if ("\n\t" < (("" + "a") + (("hello" + "foo bar") + ("a" + "foo bar")))) {
                            var v11 uint64 = p0
                            var v12 float32 = (-2.25 + 0.1)
                            var v13 uint16 = 607
                            if v7 {
                                v5 = append(v5, fi6)
                                v10 -= (v10 - (v10 * v10))
                            }
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        _ = v10
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
                var v14 uint32 = 0
                _ = v14
            }
            v15 := GenS2{F0: -869}
            _ = fi4
            _ = v15
            _ = v5
        }
    default:
        var v16 int = -1
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint16, p1 int32, p2 uint) uint64 {
    var v17 uint32 = 1
    v18 := GenS3{F0: (p1 - p1), F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 4294967295
}

func main() {
    v1 := []uint32{4294967295, 659, 381}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    v19 := GenS2{F0: -2147483648}
    var v20 uint64 = 598
    v20 -= (v20 << ((65535) & 63))
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS2{F0: 2147483647}
        v23 := GenS4{F0: 0}
        delete(v2, 4294967295)
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS5{F0: -1}
    v25 := GenS2{F0: -1}
    if false {
        var v26 uint = 4294967295
        v20 += 2
        _ = v26
    } else {
        for fi27 := 0; fi27 < 13; fi27++ {
            _ = len(v1)
            _ = fi27
        }
        v28 := GenS1{F0: (196 >= 0), F1: 1, F2: -57}
        if (false && (42 <= 32767)) {
            for fi29 := 0; fi29 < 19; fi29++ {
                if false {
                    for fi30 := 0; fi30 < 4; fi30++ {
                        v20 = v20
                        switch "" {
                        case "foo bar":
                            var v31 int16 = 1
                            var v32 int64 = 0
                            if (!true) {
                                v31 = 32767
                                var v33 int16 = -401
                                _ = v33
                            }
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi30 | fi29) {
                                case -9223372036854775808:
                                    switch fi29 {
                                    case 42:
                                        delete(v2, 4294967295)
                                    case 127:
                                        var v34 int16 = v31
                                        _ = v34
                                    }
                                    var v35 float64 = ((-2.25 + float64(3.14159)) / (0.1 + 1e10))
                                    switch (("a" + ((("hello" + (("foo bar" + (("" + "hello") + "a")) + "a")) + "hello") + ("foo bar" + "hello"))) + "foo bar") {
                                    case "":
                                        fmt.Println(v20, fi29)
                                    default:
                                        var v36 int8 = 42
                                        _ = v36
                                    }
                                    var v37 float32 = ((((1.5 / (0.1 / 1.5)) / 0.1) / 1.5) * 3.14159)
                                    var v38 int = -1
                                    _ = v35
                                    _ = v37
                                    _ = v38
                                case -500:
                                    var v39 float64 = ((0.0 * 0.1) / 1.5)
                                    _ = v39
                                case 42:
                                    v32 = v32
                                default:
                                    switch fi30 {
                                    case 42:
                                        var v40 string = "hello"
                                        _ = v40
                                    default:
                                        if true {
                                            v31 = (v31 % ((32767) | 1))
                                        }
                                    }
                                }
                            }
                            _ = v31
                            _ = v32
                        case "":
                            fmt.Println(fi30)
                        case "\n\t":
                            if ((true && (62 <= 182)) && true) {
                                v41 := GenS6{F0: 32767, F1: 1, F2: -9223372036854775808}
                                if true {
                                    v42 := v3[(("" + "\n\t") + "foo bar")]
                                    _ = v42
                                }
                                v43 := GenS3{F0: -1, F1: 42}
                                _ = v41
                                _ = v43
                            }
                        }
                        v44 := GenS4{F0: 829}
                        _ = fi30
                        _ = v44
                    }
                    fmt.Println(v20)
                }
                _ = fi29
            }
        } else {
            v45 := GenS7{F0: ("\n\t" + ("\n\t" + (("\n\t" + "") + ""))), F1: 42, F2: -1}
            _ = v45
        }
        _ = v28
    }
    var v46 uint64 = 503
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v46
}
