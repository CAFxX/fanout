package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int8 }
type GenS4 struct { F0 int16; F1 int; F2 uint }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 bool) int {
    var v7 uint16 = 2
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 16; fi9++ {
        v8 = 42
        v8 += v8
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return -559
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    v3 := []uint32{2147483647, 2147483647, 4294967295, 2147483648}
    var v4 uint = 974
    v5 := make(map[uint32]string)
    v6 := []int{-861, -1}
    delete(v1, 1)
    v4 += v4
    if len(v3) > 0 {
        v3[0] = 500
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int = fi10
        if true {
            v12 := GenS1{F0: 100}
            for fi13 := 0; fi13 < 18; fi13++ {
                v4 -= (v4 - (^(v4)))
                var v14 uint64 = 0
                var v15 int8 = 99
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        } else {
            v11 -= fi10
            v16 := GenS2{F0: "a", F1: 1}
            var v17 uint64 = 42
            _ = v16
            _ = v17
        }
        _ = fi10
        _ = v11
    }
    var v18 bool = ((float64(1e10) + 0.1) > 0.0)
    for fi19 := 0; fi19 < 10; fi19++ {
        v4 = (v4 * 1)
        _ = fi19
    }
    if v18 {
        var v20 int = -765
        if (-9223372036854775808 == -680) {
            switch genFn0(false) {
            case 790:
                var v21 uint64 = 2
                _ = v21
            case 2147483647:
                fmt.Println(v2, v4, v2)
            case 0:
                var v22 int8 = 127
                _ = v22
            }
            if (750 == -1) {
                v23 := GenS2{F0: "foo bar", F1: 2147483648}
                var v24 string = "foo bar"
                for fi25 := 0; fi25 < 16; fi25++ {
                    v26 := GenS1{F0: 100}
                    for fi27 := 0; fi27 < 12; fi27++ {
                        var v28 float32 = (1.5 - 0.1)
                        switch genFn0((!v18)) {
                        case 0:
                            var v29 bool = (4294967295 != fi25)
                            _ = v29
                        default:
                            if v18 {
                                v30 := GenS3{F0: v2, F1: 2147483647, F2: 0}
                                _ = v30
                            } else {
                                v4 -= (v4 & v4)
                            }
                            _ = len(v6)
                            v20 = (fi27 & (fi25 / (((fi25 & genFn0(false))) | 1)))
                            v2 = v2
                        }
                        _ = fi27
                        _ = v28
                    }
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v18 = (!v18)
                        v32 := GenS4{F0: 0, F1: 1, F2: (v4 / (((^(v4))) | 1))}
                        _ = fi31
                        _ = v32
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v23
                _ = v24
            } else {
                for fi33 := 0; fi33 < 20; fi33++ {
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v35 := GenS4{F0: 42, F1: fi33, F2: 0}
                        v20 -= (^(v20))
                        v36 := GenS4{F0: -460, F1: fi33, F2: (v4 + 42)}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    if v18 {
                        var v37 float32 = 3.14159
                        _ = v37
                    } else {
                        v18 = true
                    }
                    for fi38 := 0; fi38 < 17; fi38++ {
                        var v39 uint16 = 255
                        v2 += (v2 | v2)
                        _ = fi38
                        _ = v39
                    }
                    _ = fi33
                }
                _ = len(v5)
                for fi40 := 0; fi40 < 1; fi40++ {
                    var v41 uint16 = 449
                    v42 := GenS3{F0: 255, F1: 42, F2: 127}
                    if v18 {
                        _ = len(v5)
                        if v18 {
                            v43 := v1[2]
                            v20 = (^(v20))
                            v20 += (-(fi40))
                            _ = v43
                        } else {
                            var v44 float64 = (1e10 + (float64(float32(0.0)) * float64(0.1)))
                            var v45 float32 = float32((v44 + (v44 - float64(-2.25))))
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            v46 := GenS5{F0: v4, F1: int32(v20)}
            _ = v46
        } else {
            for fi47 := 0; fi47 < 9; fi47++ {
                var v48 string = "a"
                _ = fi47
                _ = v48
            }
            var v49 uint = v4
            _ = v49
        }
        _ = v20
    } else {
        var v50 int = 42
        if (!true) {
            v51 := GenS2{F0: "foo bar", F1: 225}
            v52 := GenS1{F0: 127}
            _ = v51
            _ = v52
        }
        fmt.Println(v50, v18, v2)
        _ = v50
    }
    for fi53 := 0; fi53 < 15; fi53++ {
        v54 := GenS2{F0: "foo bar", F1: 2147483647}
        _ = len(v6)
        _ = fi53
        _ = v54
    }
    var v55 int32 = 0
    var v56 string = "foo bar"
    var v57 int32 = v55
    fmt.Println(v4, v2)
    v4 += v4
    v3 = append(v3, 2147483647)
    v2 += v2
    fmt.Println(v55, v56, v18)
    if (v2 != (^(v2))) {
        v58 := v1[2147483647]
        switch "\n\t" {
        case "hello":
            if true {
                _ = len(v3)
            }
        }
        for fi59 := 0; fi59 < 14; fi59++ {
            switch (fi59 << (((v4 % ((v4) | 1))) & 63)) {
            case -870:
                v60 := GenS2{F0: v58, F1: 65535}
                _ = v60
            default:
                if len(v6) > 0 {
                    v6[0] = (fi59 % ((9223372036854775807) | 1))
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi59
        }
        _ = v58
    } else {
        v61 := GenS1{F0: -128}
        fmt.Println(v4)
        _ = v61
    }
    v62 := GenS6{F0: 32767, F1: (((1e10 + 0.1) >= 1e10) || v18)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v62
}
