package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 string, p1 uint32) bool {
    var v7 int64 = 42
    switch 127 {
    case 0:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{}
            v10 := GenS1{F0: 255, F1: 255, F2: false}
            _ = fi8
            _ = v10
            _ = v9
        }
    case -1:
        var v11 int32 = 42
        _ = v11
    case 127:
        switch "" {
        case "":
            v7 = v7
        case "a":
            p1 = p1
        default:
            v12 := []int{128, 128}
            _ = v12
        }
    default:
        p0 = "foo bar"
    }
    _ = p0
    _ = p1
    _ = v7
    return true
}

func main() {
    var v1 int8 = 42
    var v2 int16 = -355
    var v3 string = "foo bar"
    v4 := []int{773, 128, 0}
    var v5 uint16 = 42
    v6 := make(map[uint32]uint32)
    v2 += (v2 | v2)
    fmt.Println(v2)
    var v13 bool = true
    for fi14 := 0; fi14 < 12; fi14++ {
        v2 -= v2
        if (3.14159 >= 3.14159) {
            if (9223372036854775807 == 1) {
                v5 = 288
                if ((!(int(v5) <= fi14)) || v13) {
                    var v15 float32 = ((3.14159 / 1.0) * 1.5)
                    if false {
                        v5 = (v5 + v5)
                        var v16 float64 = (3.14159 + 1e10)
                        delete(v6, 4294967295)
                        _ = v16
                    }
                    var v17 uint32 = 2
                    _ = v15
                    _ = v17
                }
            } else {
                v18 := v6[2]
                var v19 bool = false
                fmt.Println(fi14)
                _ = v18
                _ = v19
            }
            var v20 int8 = (v1 ^ 0)
            _ = v20
        }
        if (v13 && false) {
            _ = len(v4)
        }
        _ = fi14
    }
    if true {
        if v13 {
            for fi21 := 0; fi21 < 1; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    if v13 {
                        if (-763 < 871) {
                            if (139 >= 2) {
                                fmt.Println(v3, fi22, fi22)
                                v23 := GenS1{F0: (v5 & v5), F1: 255, F2: v13}
                                if false {
                                    delete(v6, 798)
                                }
                                _ = v23
                            } else {
                                var v24 int = (fi21 * -2147483648)
                                v4 = append(v4, (v24 / ((fi21) | 1)))
                                v1 += v1
                                fmt.Println(v13, fi22, v5)
                                v2 -= -750
                                v24 -= fi21
                                if false {
                                    v25 := GenS2{F0: (v5 * 628)}
                                    v24 += 1
                                    fmt.Println(v3, fi21, v5)
                                    _ = v25
                                } else {
                                    v2 = 539
                                    v26 := GenS3{F0: 317, F1: (v2 & v2), F2: 128}
                                    _ = v26
                                }
                                _ = v24
                            }
                            var v27 bool = false
                            _ = v27
                        }
                        var v28 float32 = 0.1
                        v29 := v6[2147483648]
                        _ = v28
                        _ = v29
                    }
                    v13 = false
                    _ = fi22
                }
                v2 -= v2
                _ = fi21
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 uint8 = 255
                v6[4294967295] = uint32(v31)
                _ = fi30
                _ = v31
            }
        }
        v32 := GenS3{F0: 1, F1: 0, F2: 54}
        var v33 int16 = v2
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 10; fi34++ {
            if len(v4) > 0 {
                v4[0] = fi34
            }
            v35 := GenS3{F0: 18446744073709551615, F1: (^(v2)), F2: 164}
            delete(v6, 2147483647)
            _ = fi34
            _ = v35
        }
        if false {
            switch ((("foo bar" + ((((v3 + (((("a" + (v3 + ((v3 + v3) + ("a" + (v3 + (((v3 + v3) + (v3 + "hello")) + (((v3 + (v3 + v3)) + v3) + "hello"))))))) + "hello") + v3) + "")) + v3) + "foo bar") + "foo bar")) + (("" + (v3 + v3)) + v3)) + ("hello" + v3)) {
            case "a":
                var v36 int16 = (v2 << ((18446744073709551615) & 63))
                _ = v36
            case "hello":
                var v37 int32 = -236
                _ = v37
            case "\n\t":
                for fi38 := 0; fi38 < 20; fi38++ {
                    switch fi38 {
                    case 128:
                        if v13 {
                            for fi39 := 0; fi39 < 3; fi39++ {
                                v40 := GenS1{F0: (^(v5)), F1: v5, F2: true}
                                switch fi39 {
                                case -975:
                                    switch (v3 + v3) {
                                    case "hello":
                                        var v41 int8 = (v1 << ((42) & 63))
                                        _ = v41
                                    default:
                                        v3 = ((v3 + (v3 + "hello")) + "")
                                    }
                                }
                                _ = fi39
                                _ = v40
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case 1:
                        v42 := GenS1{F0: v5, F1: (v5 << ((42) & 63)), F2: v13}
                        _ = v42
                    case -1:
                        fmt.Println(fi38, v1)
                    default:
                        var v43 uint16 = 984
                        _ = v43
                    }
                    v44 := GenS1{F0: 65535, F1: v5, F2: (false || (true || false))}
                    _ = fi38
                    _ = v44
                }
            default:
                for fi45 := 0; fi45 < 5; fi45++ {
                    var v46 uint8 = 157
                    if (42 >= 1) {
                        fmt.Println(v13, v3, v2)
                    }
                    v5 -= (v5 % (((v5 & v5)) | 1))
                    _ = fi45
                    _ = v46
                }
            }
            _ = len(v6)
        }
    }
    var v47 int32 = int32(v1)
    v2 -= (v2 >> ((80) & 63))
    for fi48 := 0; fi48 < 11; fi48++ {
        var v49 int64 = 9223372036854775807
        if true {
            switch 128 {
            case 32768:
                if (v13 || (-5 >= fi48)) {
                    _ = len(v6)
                    v2 = (v2 * -521)
                }
            case 1:
                fmt.Println(v47, fi48)
            case 42:
                _ = len(v6)
            default:
                var v50 int16 = 1
                _ = v50
            }
        }
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
