package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
func genFn0(p0 uint64) int {
    var v7 bool = true
    var v8 uint64 = 1
    var v9 uint = 255
    v10 := []int{-9223372036854775808, 2147483647, 2147483647, 482}
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(v8, fi11)
        for fi12 := 0; fi12 < 1; fi12++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v7 = ((!true) || ((true && v7) || (-1 >= 1)))
            var v13 uint8 = 198
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 bool, p1 string, p2 bool) int32 {
    p1 = p1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{("" + "\n\t"), "a"}
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v5 int16 = 0
    v6 := make(map[uint32]int)
    var v14 float32 = ((0.1 - float32((1e10 - 3.14159))) + 0.1)
    var v15 float64 = (1e10 - 3.14159)
    fmt.Println(v14)
    v16 := GenS1{F0: -799, F1: 42}
    if ("" != "hello") {
        v5 = (v5 % ((-1) | 1))
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                var v18 uint16 = 255
                if false {
                    v19 := GenS1{F0: -589, F1: -1}
                    if false {
                        if (2 != 47) {
                            var v20 uint = 4294967295
                            _ = v20
                        } else {
                            if false {
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    if len(v2) > 0 {
                                        v2[0] = ("hello" + "")
                                    }
                                    _ = fi21
                                }
                                for fi22 := 0; fi22 < 16; fi22++ {
                                    fmt.Println(fi17)
                                    fmt.Println(v18, fi17)
                                    v23 := GenS2{F0: 947, F1: -9, F2: 2}
                                    _ = fi22
                                    _ = v23
                                }
                            } else {
                                v14 += ((3.14159 / (0.1 + v14)) - (v14 * (0.0 - (v14 / 3.14159))))
                                v14 += 1.5
                            }
                            v18 = 255
                        }
                        fmt.Println(v18, fi17)
                    }
                    _ = v19
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v18
            } else {
                fmt.Println(v14)
                var v24 int = genFn0(332)
                for fi25 := 0; fi25 < 14; fi25++ {
                    switch v24 {
                    case 906:
                        var v26 int = genFn0(0)
                        _ = v26
                    default:
                        var v27 bool = (-128 != 42)
                        _ = v27
                    }
                    _ = fi25
                }
                _ = v24
            }
            v28 := GenS2{F0: 693, F1: v5, F2: 2147483648}
            switch "foo bar" {
            case "\n\t":
                v29 := GenS3{F0: 56, F1: 56, F2: genFn1((v15 > 1e10), "\n\t", (!false))}
                _ = v29
            case "":
                var v30 float64 = v15
                _ = v30
            case "a":
                for fi31 := 0; fi31 < 2; fi31++ {
                    if true {
                        var v32 uint64 = 370
                        _ = v32
                    } else {
                        fmt.Println(v14)
                        if false {
                            var v33 int16 = 1
                            _ = v33
                        }
                        v34 := GenS2{F0: 2147483647, F1: v5, F2: 695}
                        if false {
                            delete(v3, 368)
                        } else {
                            var v35 uint = 4294967295
                            v14 += ((0.1 + 0.0) * ((float32(v15) * v14) - 3.14159))
                            _ = v35
                        }
                        _ = v34
                    }
                    v15 = 1e10
                    v15 -= v15
                    _ = fi31
                }
            default:
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 uint32 = 4294967295
                    var v38 uint32 = (^(v37))
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
            _ = fi17
            _ = v28
        }
    } else {
        v39 := GenS2{F0: 621, F1: v5, F2: 166}
        _ = v39
    }
    v2 = append(v2, "hello")
    for fi40 := 0; fi40 < 19; fi40++ {
        v2 = append(v2, "foo bar")
        v5 = (v5 ^ (v5 << ((585) & 63)))
        _ = fi40
    }
    var v41 uint8 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 32768 {
    case 127:
        if false {
            var v42 int32 = 146
            _ = v42
        } else {
            v5 = v5
            if (1 == -9223372036854775808) {
                v14 = v14
                v14 = float32(((float64(v14) + (v15 * ((v15 + v15) - 1e10))) / (0.1 / v15)))
                for fi43 := 0; fi43 < 2; fi43++ {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi44
                    }
                    var v45 bool = (v15 > 1e10)
                    v15 = v15
                    _ = fi43
                    _ = v45
                }
            } else {
                if false {
                    v46 := GenS3{F0: 4294967295, F1: -1, F2: 0}
                    _ = v46
                }
                fmt.Println(v5, v5, v5)
                v47 := GenS3{F0: 503, F1: -1, F2: genFn1((678 >= 691), "hello", true)}
                _ = v47
            }
        }
    case 637:
        v41 -= 127
    }
    v48 := GenS3{F0: 255, F1: -892, F2: 0}
    var v49 uint8 = 238
    switch 0 {
    case -1:
        if false {
            if (false || (true && false)) {
                var v50 uint8 = (v41 >> ((42) & 63))
                fmt.Println(v49, v49)
                _ = v50
            }
            var v51 float64 = 1e10
            if false {
                var v52 uint32 = 0
                v53 := GenS3{F0: 102, F1: 0, F2: 2147483647}
                _ = v52
                _ = v53
            } else {
                v5 += (v5 & 32767)
                fmt.Println(v49, v51)
            }
            _ = v51
        }
    }
    var v54 int64 = -172
    var v55 float64 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v48
    _ = v49
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
