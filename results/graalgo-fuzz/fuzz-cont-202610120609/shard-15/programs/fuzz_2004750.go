package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint32; F1 bool }
func genFn0() int32 {
    var v5 int64 = 42
    if false {
        var v6 uint = 0
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                if (true || false) {
                    fmt.Println(fi8, fi8, fi8)
                }
                _ = fi8
            }
            _ = fi7
        }
    }
    var v9 float32 = -2.25
    if false {
        v10 := []int{814, 1}
        _ = v10
    } else {
        fmt.Println(v5)
        if true {
            var v11 uint16 = 1
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: -9, F1: 62, F2: (^(v11))}
                switch "hello" {
                case "":
                    v14 := GenS1{F0: 0, F1: 1, F2: 65535}
                    _ = v14
                case "a":
                    v9 -= float32(((3.14159 / 0.1) * 1.5))
                case "\n\t":
                    if (fi12 != 4294967295) {
                        if (true || (42 > 0)) {
                            v15 := GenS2{F0: 255, F1: 1, F2: v5}
                            _ = v15
                        }
                        if false {
                            if false {
                                fmt.Println(v5, v9, v11)
                            } else {
                                if false {
                                    fmt.Println(fi12)
                                    var v16 float64 = 3.14159
                                    if ((v16 + v16) > 3.14159) {
                                        v5 -= -298
                                        fmt.Println(fi12, v11, v5)
                                        fmt.Println(v16, v11, fi12)
                                    } else {
                                        for fi17 := 0; fi17 < 5; fi17++ {
                                            v18 := []int{fi17, (^(fi12)), fi12, fi17}
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            _ = fi17
                                            _ = v18
                                        }
                                        var v19 uint32 = 42
                                        var v20 uint16 = v11
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v16
                                }
                                v9 = v9
                            }
                        }
                        v5 += v5
                    } else {
                        if true {
                            var v21 int32 = -643
                            for fi22 := 0; fi22 < 15; fi22++ {
                                v23 := []int{}
                                _ = fi22
                                _ = v23
                            }
                            v24 := []int{(fi12 | (fi12 / ((fi12) | 1))), fi12, fi12, fi12}
                            _ = v21
                            _ = v24
                        } else {
                            var v25 uint = 18446744073709551615
                            _ = v25
                        }
                    }
                default:
                    if (128 <= 2) {
                        v26 := []int{}
                        for fi27 := 0; fi27 < 5; fi27++ {
                            fmt.Println(v5, v9)
                            _ = len(v26)
                            _ = fi27
                        }
                        v28 := GenS3{F0: 177}
                        _ = v26
                        _ = v28
                    } else {
                        if true {
                            for fi29 := 0; fi29 < 1; fi29++ {
                                if true {
                                    v9 -= float32((3.14159 * 1.5))
                                    if (255 < 2) {
                                        v30 := GenS2{F0: 23, F1: 127, F2: v5}
                                        v31 := GenS3{F0: 127}
                                        if false {
                                            if true {
                                                fmt.Println(v11, v5)
                                                var v32 float32 = (v9 - (v9 + float32(((float64(1.5) - 3.14159) + 1.5))))
                                                _ = v32
                                            }
                                            v33 := []int{9223372036854775807, (-(fi29)), (fi12 ^ (^(fi12)))}
                                            fmt.Println(v5, v5, v5)
                                            _ = v33
                                        } else {
                                            var v34 uint32 = 848
                                            var v35 int64 = (v5 / (((v5 * int64(v11))) | 1))
                                            var v36 int16 = 0
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                    var v37 uint16 = 843
                                    _ = v37
                                } else {
                                    var v38 int64 = (^(v5))
                                    _ = v38
                                }
                                _ = fi29
                            }
                            v39 := []int{(-(fi12)), fi12, fi12}
                            v9 = (((1.5 / 1e10) / v9) + (((((v9 * v9) - v9) / (float32(3.14159) - v9)) + v9) - 0.0))
                            _ = v39
                        }
                    }
                }
                _ = fi12
                _ = v13
            }
            var v40 int64 = v5
            _ = v11
            _ = v40
        }
    }
    _ = v5
    _ = v9
    return -993
}
func genFn1(p0 uint32) bool {
    if true {
        v41 := []int{}
        _ = v41
    }
    var v42 uint8 = 127
    _ = p0
    _ = v42
    return false
}

func main() {
    v1 := make(map[string]int)
    var v2 string = "foo bar"
    var v3 uint8 = 127
    v4 := make(map[uint32]string)
    for fi43 := 0; fi43 < 9; fi43++ {
        var v44 uint64 = 0
        for fi45 := 0; fi45 < 3; fi45++ {
            var v46 int64 = -1
            _ = fi45
            _ = v46
        }
        _ = fi43
        _ = v44
    }
    v47 := GenS4{F0: 255, F1: 0}
    var v48 uint64 = 42
    if true {
        var v49 float32 = -2.25
        _ = v49
    } else {
        var v50 string = "a"
        for fi51 := 0; fi51 < 9; fi51++ {
            var v52 int16 = 32767
            var v53 float64 = 1e10
            fmt.Println(v50, v53, v52)
            _ = fi51
            _ = v52
            _ = v53
        }
        var v54 int32 = 0
        _ = v50
        _ = v54
    }
    if (32767 < -708) {
        var v55 int64 = 9223372036854775807
        v56 := GenS5{F0: v2}
        var v57 int16 = -1
        _ = v55
        _ = v56
        _ = v57
    }
    _ = len(v4)
    var v58 string = "\n\t"
    if true {
        if false {
            switch 127 {
            case 128:
                if (!false) {
                    v3 += (^(v3))
                    switch (v58 + v58) {
                    case "":
                        var v59 int32 = 2147483647
                        _ = v59
                    case "hello":
                        var v60 float32 = 0.1
                        _ = v60
                    }
                } else {
                    fmt.Println(v48, v2, v58)
                    v3 += (v3 | (v3 << ((42) & 63)))
                }
            case 42:
                v48 += 4294967295
            }
            if ((631 == v48) || true) {
                for fi61 := 0; fi61 < 6; fi61++ {
                    v58 = "hello"
                    _ = fi61
                }
            } else {
                v1["\n\t"] = 32767
                for fi62 := 0; fi62 < 18; fi62++ {
                    fmt.Println(fi62)
                    fmt.Println(v48)
                    _ = fi62
                }
            }
        }
        if (-528 >= 147) {
            for fi63 := 0; fi63 < 10; fi63++ {
                if (!(1 == 861)) {
                    var v64 float32 = 1.5
                    if ((("\n\t" >= v58) || true) || (65535 <= 177)) {
                        v65 := GenS6{F0: false, F1: -1}
                        _ = v65
                    } else {
                        v66 := v1[v2]
                        _ = v66
                    }
                    v67 := []int{fi63}
                    _ = v64
                    _ = v67
                }
                v2 = (v58 + "\n\t")
                _ = fi63
            }
        } else {
            v68 := GenS3{F0: (v3 >> ((1) & 63))}
            fmt.Println(v2, v2, v48)
            _ = v68
        }
        var v69 uint8 = v3
        _ = v69
    } else {
        v70 := GenS7{F0: 42, F1: (42 >= 2)}
        _ = v70
    }
    v2 = ((("hello" + "") + ((v2 + v58) + v2)) + (v2 + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v58
}
