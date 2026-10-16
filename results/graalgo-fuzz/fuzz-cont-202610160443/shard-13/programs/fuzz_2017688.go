package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 int; F2 int64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 int64 }
func genFn0() uint8 {
    v4 := GenS1{F0: 988, F1: 128}
    var v5 uint = 18446744073709551615
    v5 = v5
    v5 += v5
    _ = v4
    _ = v5
    return 42
}
func genFn1() int {
    var v6 int64 = -1
    var v7 float32 = 0.1
    v8 := []int{32767, 741, -2147483648, 662}
    if (!false) {
        var v9 int64 = v6
        for fi10 := 0; fi10 < 19; fi10++ {
            v7 = v7
            switch "" {
            case "\n\t":
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS2{F0: (-(v6))}
                    _ = fi11
                    _ = v12
                }
            }
            switch ("" + "foo bar") {
            case "a":
                v13 := GenS3{F0: 618, F1: 255}
                _ = v13
            }
            _ = fi10
        }
        _ = v9
    } else {
        if (1.5 <= ((v7 / (v7 * v7)) / (1.5 + 3.14159))) {
            if true {
                var v14 int32 = -181
                v14 -= (v14 << ((18446744073709551615) & 63))
                if false {
                    switch 127 {
                    case -23:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    default:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    fmt.Println(v6)
                    var v15 uint16 = 42
                    _ = v15
                }
                _ = v14
            }
            var v16 int64 = v6
            _ = len(v8)
            _ = v16
        }
    }
    _ = v6
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    v3 := make(map[uint32]bool)
    var v17 float32 = -2.25
    _ = len(v1)
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint32 = 259
        switch (fi18 & (fi18 - genFn1())) {
        case 2147483647:
            var v20 bool = v2
            _ = v20
        default:
            var v21 int16 = -32768
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v2 = (v17 < v17)
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := []int{(-(fi22)), (fi22 % ((fi22) | 1)), fi22}
        var v24 uint8 = genFn0()
        _ = fi22
        _ = v23
        _ = v24
    }
    if (!(false && v2)) {
        v25 := GenS4{F0: -1}
        _ = v25
    } else {
        var v26 uint8 = genFn0()
        _ = v26
    }
    fmt.Println(v17, v2)
    v27 := GenS1{F0: 42, F1: 128}
    fmt.Println(v17)
    var v28 int32 = -280
    var v29 int8 = -1
    v30 := GenS3{F0: 435, F1: 0}
    if ("a" > (("hello" + "\n\t") + "")) {
        fmt.Println(v29, v29, v28)
    }
    var v31 uint = 800
    if v2 {
        var v32 int = 485
        v29 += (v29 / ((-128) | 1))
        if (true && v2) {
            if ("hello" == (("" + ("" + "")) + "hello")) {
                for fi33 := 0; fi33 < 14; fi33++ {
                    v2 = v2
                    if v2 {
                        for fi34 := 0; fi34 < 10; fi34++ {
                            v35 := GenS5{F0: 127, F1: 1}
                            if true {
                                var v36 uint64 = 813
                                fmt.Println(v29)
                                var v37 uint32 = 2
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS1{F0: 65535, F1: v32}
                                v39 := GenS2{F0: -9223372036854775808}
                                _ = v38
                                _ = v39
                            }
                            if true {
                                v31 -= (v31 << (((v31 % ((0) | 1))) & 63))
                            } else {
                                v40 := GenS1{F0: 32767, F1: -1}
                                _ = v40
                            }
                            _ = len(v3)
                            _ = fi34
                            _ = v35
                        }
                    } else {
                        switch v32 {
                        case -9223372036854775808:
                            v41 := GenS3{F0: 32767, F1: 1}
                            _ = v41
                        default:
                            v17 = 3.14159
                        }
                        if (302 < -795) {
                            delete(v3, 1)
                        } else {
                            v42 := []int{fi33, -160, v32, v32}
                            v43 := v3[0]
                            _ = v42
                            _ = v43
                        }
                    }
                    v44 := GenS4{F0: 0}
                    _ = fi33
                    _ = v44
                }
                v45 := GenS2{F0: -867}
                _ = v45
            } else {
                var v46 uint16 = 255
                for fi47 := 0; fi47 < 10; fi47++ {
                    _ = len(v1)
                    v48 := GenS5{F0: 255, F1: v46}
                    _ = fi47
                    _ = v48
                }
                _ = v46
            }
            v49 := GenS1{F0: uint16(v31), F1: (v32 + v32)}
            _ = v49
        } else {
            v28 += 643
            var v50 int = 0
            _ = v50
        }
        _ = v32
    } else {
        switch "" {
        case "a":
            v51 := GenS1{F0: 866, F1: 0}
            _ = v51
        }
        v29 = 127
    }
    for fi52 := 0; fi52 < 15; fi52++ {
        switch fi52 {
        case -1:
            for fi53 := 0; fi53 < 10; fi53++ {
                var v54 int8 = (v29 ^ v29)
                v29 -= 0
                v2 = false
                _ = fi53
                _ = v54
            }
        }
        var v55 bool = (false && v2)
        _ = fi52
        _ = v55
    }
    if true {
        switch "\n\t" {
        case "":
            var v56 int8 = (v29 | -128)
            _ = v56
        }
    } else {
        for fi57 := 0; fi57 < 16; fi57++ {
            for fi58 := 0; fi58 < 8; fi58++ {
                v31 = (v31 << ((v31) & 63))
                _ = fi58
            }
            v59 := GenS2{F0: 0}
            _ = fi57
            _ = v59
        }
        for fi60 := 0; fi60 < 17; fi60++ {
            if v2 {
                for fi61 := 0; fi61 < 16; fi61++ {
                    if (true && v2) {
                        if false {
                            var v62 string = ""
                            _ = v62
                        } else {
                            v1[228] = 2
                            var v63 float32 = ((1e10 - 3.14159) + (1e10 / float32((0.1 * -2.25))))
                            _ = v63
                        }
                        if (!(v2 || ((fi61 >> ((0) & 63)) < 865))) {
                            var v64 string = "hello"
                            fmt.Println(v31, v31)
                            v65 := GenS6{F0: ("hello" + ((v64 + ("" + v64)) + v64))}
                            _ = v64
                            _ = v65
                        } else {
                            var v66 bool = true
                            if v2 {
                                v1[806] = 2
                            } else {
                                var v67 uint32 = 0
                                if (v29 == (^(v29))) {
                                    v66 = v2
                                    v68 := GenS7{F0: fi60, F1: fi61, F2: 42}
                                    _ = v68
                                }
                                var v69 string = ""
                                _ = v67
                                _ = v69
                            }
                            _ = v66
                        }
                    }
                    var v70 uint8 = genFn0()
                    v29 = (v29 - 0)
                    _ = fi61
                    _ = v70
                }
            }
            v71 := GenS8{F0: 127}
            var v72 uint64 = 196
            _ = fi60
            _ = v71
            _ = v72
        }
        v73 := GenS9{F0: -2147483648, F1: 42}
        _ = v73
    }
    v74 := []int{42, -2147483648, 326}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v74
}
