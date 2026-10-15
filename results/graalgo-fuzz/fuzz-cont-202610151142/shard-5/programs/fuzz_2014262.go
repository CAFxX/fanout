package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int64; F1 uint8 }
func genFn0() uint {
    v6 := []int{32767, 128, -2147483648}
    v7 := GenS1{F0: 42}
    var v8 string = "\n\t"
    _ = len(v6)
    if false {
        var v9 float64 = -2.25
        if true {
            v9 -= 3.14159
            v10 := GenS1{F0: 127}
            _ = v10
        } else {
            var v11 uint16 = 2
            _ = v11
        }
        var v12 uint64 = 0
        _ = v12
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 625
}
func genFn1(p0 bool, p1 uint8, p2 int) string {
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := []int{p2, p2}
        v15 := GenS1{F0: 127}
        for fi16 := 0; fi16 < 14; fi16++ {
            if p0 {
                fmt.Println(fi13, p2, fi13)
                if p0 {
                    v14 = append(v14, 42)
                } else {
                    if len(v14) > 0 {
                        v14[0] = fi16
                    }
                    var v17 int64 = -1
                    _ = v17
                }
                fmt.Println(p1, fi16)
                if len(v14) > 0 {
                    v14[0] = (-(fi13))
                }
            }
            _ = fi16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    if true {
        var v18 int16 = 42
        _ = v18
    } else {
        p0 = (1000 <= -1)
        v19 := []int{1, 1, 909}
        _ = v19
    }
    var v20 int64 = -1
    v21 := GenS2{F0: ""}
    for fi22 := 0; fi22 < 5; fi22++ {
        switch "foo bar" {
        case "a":
            for fi23 := 0; fi23 < 12; fi23++ {
                if p0 {
                    v24 := []int{(fi23 + fi23), 658, 0, p2}
                    var v25 uint32 = 42
                    v25 += (v25 << ((749) & 63))
                    fmt.Println(p0, fi23)
                    fmt.Println(p2, v20)
                    _ = v24
                    _ = v25
                } else {
                    fmt.Println(p0, p2)
                    v20 = 9223372036854775807
                    v26 := GenS2{F0: "hello"}
                    _ = v26
                }
                v27 := GenS3{F0: 100, F1: -29}
                _ = fi23
                _ = v27
            }
        default:
            p1 = p1
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return (("a" + "a") + "a")
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]uint32)
    v3 := []uint32{1, 65535, 466}
    var v4 int = 4294967295
    var v5 uint16 = 317
    fmt.Println(v4, v5)
    if false {
        v28 := GenS4{F0: 42}
        _ = v28
    } else {
        if (true && true) {
            if (!(!((!false) && (127 > (v1 / ((0) | 1)))))) {
                v5 = v5
                for fi29 := 0; fi29 < 13; fi29++ {
                    v4 += (^(fi29))
                    v5 = (v5 << ((4294967295) & 63))
                    _ = fi29
                }
            } else {
                var v30 int64 = 9223372036854775807
                _ = v30
            }
            fmt.Println(v5, v4)
        }
        v31 := GenS2{F0: "hello"}
        _ = v31
    }
    var v32 uint64 = 745
    var v33 float32 = (3.14159 + (float32(3.14159) - 3.14159))
    fmt.Println(v32, v4)
    if false {
        v3 = append(v3, 2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (!false) {
            if false {
                v34 := GenS5{F0: true}
                var v35 int64 = -1
                v3 = append(v3, 65535)
                _ = v34
                _ = v35
            } else {
                v36 := GenS6{F0: 2147483647, F1: ((v32 << ((48) & 63)) < v32)}
                switch "\n\t" {
                case "a":
                    if (true || true) {
                        v5 = v5
                        if (true || (0.1 == (((0.1 * 0.0) + ((0.1 + (1.5 - -2.25)) / 1.0)) + 3.14159))) {
                            if (!false) {
                                v4 = (-(v4))
                                if (!(!true)) {
                                    v37 := GenS6{F0: 0, F1: (false && false)}
                                    _ = v37
                                } else {
                                    delete(v2, v4)
                                    v33 = 3.14159
                                    _ = len(v3)
                                }
                            }
                        }
                    } else {
                        for fi38 := 0; fi38 < 7; fi38++ {
                            var v39 int64 = 9223372036854775807
                            _ = fi38
                            _ = v39
                        }
                        var v40 int8 = -67
                        _ = v40
                    }
                case "\n\t":
                    fmt.Println(v32)
                }
                v4 = v4
                _ = v36
            }
            v1 -= (v1 % (((v1 * 2)) | 1))
        }
    }
    var v41 uint64 = 4294967295
    fmt.Println(v4, v5)
    if false {
        switch "hello" {
        case "a":
            for fi42 := 0; fi42 < 8; fi42++ {
                if ((!(1 == -1)) || false) {
                    var v43 int = (fi42 << ((116) & 63))
                    _ = v43
                } else {
                    switch fi42 {
                    case -297:
                        v41 = 682
                    case 32767:
                        v5 -= (v5 / (((v5 % ((192) | 1))) | 1))
                    case -433:
                        v44 := GenS3{F0: 42, F1: 100}
                        _ = v44
                    default:
                        v45 := GenS7{F0: 205, F1: v1}
                        _ = v45
                    }
                    fmt.Println(v33, v1, v32)
                }
                v46 := GenS3{F0: 60, F1: 100}
                for fi47 := 0; fi47 < 13; fi47++ {
                    switch (v4 - fi42) {
                    case 1:
                        var v48 int = fi42
                        _ = v48
                    default:
                        switch (fi47 ^ 396) {
                        case 4294967295:
                            var v49 uint64 = (v41 - v41)
                            _ = v49
                        }
                    }
                    _ = fi47
                }
                _ = fi42
                _ = v46
            }
        case "foo bar":
            var v50 int16 = -1
            _ = v50
        }
        v4 = 0
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        fmt.Println(v41, v41, v1)
        v5 -= (^(v5))
        _ = fi51
    }
    var v52 uint16 = uint16(v1)
    var v53 float32 = v33
    var v54 bool = false
    if (float32(-2.25) == float32(1e10)) {
        switch "\n\t" {
        case "":
            fmt.Println(v53, v54)
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 2
            }
        }
        v1 = v1
        for fi55 := 0; fi55 < 14; fi55++ {
            for fi56 := 0; fi56 < 17; fi56++ {
                if true {
                    var v57 int8 = int8(fi56)
                    _ = v57
                }
                var v58 uint32 = 65535
                _ = fi56
                _ = v58
            }
            v52 = (v52 >> ((4294967295) & 63))
            _ = fi55
        }
    } else {
        switch ("" + "\n\t") {
        case "a":
            v59 := v2[v4]
            _ = v59
        default:
            v54 = v54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
