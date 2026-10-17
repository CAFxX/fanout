package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint32; F1 string; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 int32 }
func genFn0(p0 int64, p1 int32, p2 bool) uint {
    v5 := GenS1{F0: 1, F1: 183}
    fmt.Println(p1)
    fmt.Println(p0, p0)
    v6 := []int{128, 42, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int, p1 uint8) int64 {
    switch -901 {
    case -9223372036854775808:
        p1 = (^(p1))
    case -1:
        p0 += p0
    }
    v7 := []int{(^(p0)), (^(p0))}
    switch (p0 / ((p0) | 1)) {
    case 534:
        if (0.1 != float64(1e10)) {
            p1 += (p1 * 1)
            if false {
                if (!((!((p1 <= 100) && (46 > -25))) && false)) {
                    v8 := GenS2{F0: 1, F1: 755, F2: ("foo bar" + "hello")}
                    _ = v8
                }
            }
            var v9 uint = 18446744073709551615
            _ = v9
        } else {
            if (42 != 42) {
                var v10 uint32 = 2147483648
                var v11 bool = false
                _ = v10
                _ = v11
            } else {
                p1 = (p1 / ((p1) | 1))
            }
            var v12 uint8 = p1
            _ = v12
        }
    case 420:
        var v13 uint = 18446744073709551615
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint16 = 877
    v14 := []int{1, 42, 107}
    var v15 uint16 = v4
    fmt.Println(v15)
    v4 = (v4 * 418)
    v4 += v4
    v4 += (v15 + v4)
    v16 := v1[127]
    v17 := GenS3{F0: ""}
    if (!((!(42 <= 18446744073709551615)) || (0 != 423))) {
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v4)
            fmt.Println(v15)
            v16 = true
            _ = fi18
        }
        switch -949 {
        case -413:
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS4{F0: 721, F1: 0}
                _ = fi19
                _ = v20
            }
        case 127:
            var v21 uint8 = 104
            _ = v21
        case -9223372036854775808:
            var v22 uint32 = 1
            _ = v22
        default:
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS2{F0: -2147483648, F1: 32767, F2: "foo bar"}
                fmt.Println(fi23, fi23, v16)
                switch "" {
                case "foo bar":
                    if ((-2.25 - 0.1) > 3.14159) {
                        switch "foo bar" {
                        case "hello":
                            for fi25 := 0; fi25 < 8; fi25++ {
                                switch 89 {
                                case 4294967295:
                                    v26 := v3[(fi23 & fi23)]
                                    _ = v26
                                case 32768:
                                    switch 128 {
                                    case 32768:
                                        v4 -= 65535
                                    case 128:
                                        switch fi25 {
                                        case -2147483648:
                                            v4 = v15
                                        }
                                    }
                                default:
                                    var v27 int16 = -32768
                                    _ = v27
                                }
                                if (4294967295 < 65535) {
                                    if (0 != 832) {
                                        var v28 int8 = 0
                                        if (3.14159 == ((1e10 / (-2.25 - (0.1 / (-2.25 * -2.25)))) - 0.0)) {
                                            if true {
                                                fmt.Println(v4, v15)
                                                var v29 int16 = 1000
                                                _ = v29
                                            }
                                            var v30 uint32 = 2147483647
                                            _ = v30
                                        } else {
                                            var v31 float64 = (1.5 / float64((0.1 / float32(-2.25))))
                                            switch "hello" {
                                            case "\n\t":
                                                switch "a" {
                                                case "foo bar":
                                                    v32 := GenS5{F0: 42, F1: 9223372036854775807}
                                                    _ = v32
                                                case "\n\t":
                                                    if (!v16) {
                                                        v28 = v28
                                                    }
                                                }
                                            default:
                                                v16 = true
                                            }
                                            _ = v31
                                        }
                                        _ = v28
                                    }
                                    var v33 uint16 = (^(v4))
                                    _ = v33
                                }
                                _ = fi25
                            }
                        case "a":
                            v34 := GenS6{F0: 42, F1: ("hello" + ""), F2: true}
                            _ = v34
                        case "foo bar":
                            fmt.Println(v16, v15)
                        default:
                            v35 := GenS7{F0: 0}
                            _ = v35
                        }
                        v36 := v1[fi23]
                        for fi37 := 0; fi37 < 4; fi37++ {
                            v14 = append(v14, (fi23 - (fi37 - 127)))
                            var v38 uint64 = 2
                            _ = fi37
                            _ = v38
                        }
                        _ = v36
                    }
                case "hello":
                    var v39 int32 = -1
                    _ = v39
                }
                _ = fi23
                _ = v24
            }
        }
    } else {
        v40 := GenS8{F0: 128, F1: 18446744073709551615}
        var v41 int32 = -575
        _ = v40
        _ = v41
    }
    var v42 int64 = -313
    delete(v1, 128)
    v4 = v4
    for fi43 := 0; fi43 < 9; fi43++ {
        v15 -= (v15 + v4)
        _ = fi43
    }
    v44 := GenS9{F0: 1}
    v45 := GenS4{F0: 0, F1: 42}
    v46 := v2[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
}
