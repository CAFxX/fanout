package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 int8) uint32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint = 736
        v8 := GenS1{F0: 0, F1: "hello", F2: 715}
        var v9 int64 = -1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 4294967295, F1: 65535}
    var v11 float64 = float64(-2.25)
    v12 := GenS3{F0: 0, F1: p0, F2: 128}
    for fi13 := 0; fi13 < 19; fi13++ {
        if ((true || (953 > 2)) && false) {
            v14 := []int{128, fi13}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 11; fi15++ {
                if false {
                    var v16 int8 = (p0 % (((p0 ^ p0)) | 1))
                    var v17 int = (fi15 * fi15)
                    _ = v16
                    _ = v17
                } else {
                    var v18 uint32 = 42
                    if (!false) {
                        var v19 int64 = 42
                        _ = v19
                    }
                    _ = v18
                }
                _ = fi15
            }
            for fi20 := 0; fi20 < 18; fi20++ {
                if true {
                    var v21 uint8 = 172
                    v22 := []int{(^(fi13))}
                    if true {
                        v11 -= ((0.1 + v11) - v11)
                        var v23 uint64 = 18446744073709551615
                        _ = v23
                    }
                    v21 = v21
                    var v24 uint32 = 186
                    var v25 float64 = (3.14159 + float64(1e10))
                    _ = v21
                    _ = v22
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint16 = 32767
                    p0 = (^(p0))
                    _ = v26
                }
                _ = fi20
            }
            p0 -= p0
        }
        var v27 uint16 = 0
        v28 := GenS4{F0: 4294967295}
        _ = fi13
        _ = v27
        _ = v28
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 42
}
func genFn1() uint {
    var v29 float64 = ((((0.1 + -2.25) * float64(0.0)) + 1e10) + float64(0.1))
    v30 := []int{}
    _ = v29
    _ = v30
    return 2
}

func main() {
    v1 := []uint32{4294967295, 553}
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 string = "hello"
    var v31 int64 = 9223372036854775807
    var v32 int32 = 0
    var v33 int8 = 0
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 int16 = -378
        v1 = append(v1, 2147483647)
        v35 += int16(v31)
        _ = fi34
        _ = v35
    }
    var v36 string = v5
    _ = len(v4)
    v31 -= (v31 << ((genFn1()) & 63))
    var v37 float64 = 0.0
    var v38 int8 = 127
    v39 := GenS4{F0: 4294967295}
    if true {
        if true {
            var v40 int8 = v33
            var v41 bool = ((float32(v37) <= 0.0) || false)
            _ = v40
            _ = v41
        } else {
            var v42 uint8 = 2
            v43 := GenS4{F0: 0}
            if (!false) {
                switch "foo bar" {
                case "foo bar":
                    v44 := GenS2{F0: 65535, F1: 42}
                    _ = v44
                }
                if ((v42 >> ((4294967295) & 63)) < 1) {
                    for fi45 := 0; fi45 < 7; fi45++ {
                        v38 = v38
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if true {
                            for fi46 := 0; fi46 < 4; fi46++ {
                                v5 = "foo bar"
                                v47 := GenS4{F0: 2}
                                var v48 uint = 578
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        } else {
                            if true {
                                _ = len(v4)
                            }
                        }
                        _ = fi45
                    }
                    var v49 uint8 = 127
                    v50 := GenS5{F0: v33, F1: 1, F2: "\n\t"}
                    _ = v49
                    _ = v50
                }
                v51 := GenS5{F0: 1, F1: 970, F2: "\n\t"}
                _ = v51
            } else {
                v52 := GenS6{F0: -1, F1: 4294967295}
                delete(v3, 42)
                delete(v2, ("" + "hello"))
                _ = v52
            }
            _ = v42
            _ = v43
        }
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 int16 = 0
            _ = fi53
            _ = v54
        }
    } else {
        var v55 float32 = 3.14159
        delete(v2, v36)
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
