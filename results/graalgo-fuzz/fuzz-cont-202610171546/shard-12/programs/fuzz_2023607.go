package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int8; F1 bool }
func genFn0(p0 int) uint32 {
    v5 := GenS1{F0: 2, F1: "hello"}
    p0 = 32767
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(fi7, fi6)
            if (false || false) {
                if (!(!true)) {
                    if (3.14159 > float64(1.5)) {
                        switch p0 {
                        case 2147483647:
                            var v8 int32 = -2147483648
                            _ = v8
                        case -117:
                            v9 := []int{(p0 * p0), fi7}
                            _ = v9
                        case 0:
                            var v10 int32 = -734
                            _ = v10
                        default:
                            var v11 float64 = 1e10
                            _ = v11
                        }
                        fmt.Println(fi7, fi6)
                        var v12 int64 = -1
                        v13 := GenS1{F0: 825, F1: "foo bar"}
                        var v14 int8 = int8(v12)
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                if (true && (42 > 9223372036854775807)) {
                    v15 := GenS1{F0: 32767, F1: "foo bar"}
                    p0 += (fi6 >> ((0) & 63))
                    fmt.Println(p0, p0, fi6)
                    _ = v15
                }
                var v16 bool = (65535 > 32767)
                _ = v16
            } else {
                if false {
                    var v17 uint32 = 65535
                    v17 = (v17 & v17)
                    _ = v17
                }
                fmt.Println(fi7, p0)
            }
            v18 := GenS2{F0: 42}
            _ = fi7
            _ = v18
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return 501
}
func genFn1(p0 int8, p1 uint, p2 uint32) int16 {
    if (false || (802 >= 2147483647)) {
        for fi19 := 0; fi19 < 4; fi19++ {
            if false {
                if (!(-1 >= 0)) {
                    v20 := GenS1{F0: 32767, F1: "foo bar"}
                    _ = v20
                } else {
                    v21 := GenS3{F0: 255}
                    p1 = (p1 ^ p1)
                    _ = v21
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 string = ""
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi19
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            fmt.Println(p0, p1)
            fmt.Println(p2, p2, fi24)
            for fi25 := 0; fi25 < 3; fi25++ {
                v26 := []int{fi24, (fi24 % (((fi24 << ((619) & 63))) | 1)), fi25}
                var v27 uint64 = 18446744073709551615
                var v28 uint = 42
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            _ = fi24
        }
    } else {
        switch 106 {
        case -641:
            var v29 int64 = -274
            _ = v29
        case 128:
            v30 := GenS2{F0: 800}
            _ = v30
        }
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 uint32 = (p2 ^ (p2 / (((p2 ^ p2)) | 1)))
            _ = fi31
            _ = v32
        }
        if false {
            var v33 string = "hello"
            _ = v33
        } else {
            fmt.Println(p1)
        }
    }
    v34 := GenS3{F0: 128}
    v35 := []int{}
    fmt.Println(p0)
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    return -615
}

func main() {
    v1 := []uint32{}
    var v2 int = 1
    var v3 uint32 = 99
    v4 := make(map[string]string)
    var v36 uint = 613
    for fi37 := 0; fi37 < 7; fi37++ {
        fmt.Println(fi37, v2)
        fmt.Println(v3, v2, v2)
        if (false && (!true)) {
            v38 := GenS4{F0: 2147483647}
            if false {
                for fi39 := 0; fi39 < 11; fi39++ {
                    if false {
                        if (2 > 1) {
                            switch "hello" {
                            case "":
                                var v40 int = v2
                                _ = v40
                            }
                            v41 := GenS4{F0: 2147483647}
                            var v42 uint8 = 127
                            _ = v41
                            _ = v42
                        } else {
                            v43 := GenS5{F0: 1, F1: 42}
                            _ = v43
                        }
                        if false {
                            v2 = (v2 * fi39)
                            v36 += 18446744073709551615
                            fmt.Println(fi37)
                        }
                        v2 += (fi39 ^ fi39)
                    } else {
                        v2 -= v2
                        v2 -= v2
                        v3 = (^(v3))
                        var v44 uint32 = genFn0(v2)
                        v2 = 2147483647
                        fmt.Println(v2)
                        _ = v44
                    }
                    if false {
                        var v45 uint64 = 2
                        _ = v45
                    }
                    _ = fi39
                }
            } else {
                var v46 uint = 2
                if false {
                    var v47 int8 = -1
                    var v48 float64 = float64((1e10 + -2.25))
                    _ = v47
                    _ = v48
                }
                _ = v46
            }
            _ = v38
        } else {
            v49 := GenS1{F0: 32767, F1: ("\n\t" + "")}
            _ = v49
        }
        _ = fi37
    }
    v50 := GenS2{F0: 1}
    v2 += v2
    v2 -= (v2 - v2)
    var v51 bool = true
    v52 := GenS3{F0: 127}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v53 string = "\n\t"
    v54 := GenS3{F0: 127}
    if v51 {
        var v55 int8 = -1
        _ = v55
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        if ((v51 || (v51 && v51)) && (fi56 >= fi56)) {
            var v57 uint32 = (v3 ^ (v3 + (v3 << ((v36) & 63))))
            for fi58 := 0; fi58 < 9; fi58++ {
                if true {
                    _ = len(v1)
                    v59 := GenS6{F0: 127, F1: (v51 || (0.1 <= float32((0.1 * float64((((float32((-2.25 - ((0.1 - -2.25) * (float64(1.5) + 3.14159)))) / 1.5) * 0.0) - (float32((-2.25 + ((3.14159 * float64(0.0)) - -2.25))) * -2.25)))))))}
                    _ = v59
                }
                _ = fi58
            }
            _ = v57
        } else {
            if ((32767 > 653) || v51) {
                var v60 int64 = 833
                if v51 {
                    var v61 uint8 = 128
                    _ = v61
                } else {
                    var v62 uint8 = 255
                    switch "hello" {
                    case "hello":
                        v63 := GenS1{F0: 32767, F1: "\n\t"}
                        _ = v63
                    case "foo bar":
                        var v64 int16 = 0
                        _ = v64
                    }
                    _ = v62
                }
                _ = v60
            }
        }
        v3 = 4294967295
        _ = fi56
    }
    if ((0.0 + 3.14159) == -2.25) {
        fmt.Println(v3)
    } else {
        v2 -= (v2 & v2)
        delete(v4, "a")
        v51 = false
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
