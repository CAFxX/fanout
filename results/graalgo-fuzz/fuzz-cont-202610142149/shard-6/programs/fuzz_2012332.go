package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 uint, p2 uint32) uint64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{fi6, (fi6 & fi6), (^(fi6))}
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: -1}
            var v10 int64 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v11 := []int{-291}
    v11 = append(v11, -2147483648)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 bool = (-2.25 != 3.14159)
    v3 := make(map[uint32]bool)
    v4 := []int{2147483647, 0, 9223372036854775807, 4294967295}
    var v5 int16 = 862
    var v12 bool = (-315 >= 1)
    var v13 int16 = -408
    if (v5 >= -1) {
        var v14 uint32 = v1
        var v15 uint64 = 4294967295
        var v16 bool = v2
        _ = v14
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: -128}
        _ = v17
    }
    _ = len(v3)
    v5 -= (^(v13))
    if true {
        v12 = (42 > 42)
    } else {
        if v12 {
            v1 = 2147483648
            if v12 {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS1{F0: -128}
                    if v12 {
                        var v20 int32 = 1
                        _ = v20
                    }
                    var v21 uint8 = 128
                    _ = fi18
                    _ = v19
                    _ = v21
                }
            }
        } else {
            if len(v4) > 0 {
                v4[0] = 32768
            }
        }
    }
    var v22 uint = 4294967295
    if v2 {
        v22 += (^(v22))
        switch 4294967295 {
        case 2147483647:
            if (2 < 1) {
                switch ((("" + "foo bar") + (("hello" + "\n\t") + ((("foo bar" + "") + "a") + "\n\t"))) + "") {
                case "foo bar":
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if (!v12) {
                            var v24 int = 9223372036854775807
                            _ = v24
                        } else {
                            var v25 string = ""
                            v26 := GenS2{F0: -1}
                            if (255 <= 128) {
                                v2 = ((!v12) && (839 < (^(fi23))))
                                v22 = v22
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = v25
                            _ = v26
                        }
                        _ = fi23
                    }
                case "hello":
                    if (!(v12 && ((1e10 * 0.0) != (-2.25 + float64(0.0))))) {
                        var v27 float32 = 1e10
                        var v28 int = 1
                        _ = v27
                        _ = v28
                    }
                case "":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 bool = (65535 <= 0)
                        _ = fi29
                        _ = v30
                    }
                default:
                    v31 := GenS3{F0: -222, F1: 9223372036854775807, F2: 847}
                    _ = v31
                }
                v32 := GenS4{F0: v5}
                if v2 {
                    v33 := GenS2{F0: (v13 | (v5 - v13))}
                    v1 = v1
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 uint16 = 2
                        v36 := GenS2{F0: 32767}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    if v12 {
                        if false {
                            delete(v3, (v1 << ((v22) & 63)))
                        }
                    } else {
                        for fi37 := 0; fi37 < 2; fi37++ {
                            v4 = append(v4, (fi37 * fi37))
                            delete(v3, 2147483647)
                            _ = fi37
                        }
                        if v2 {
                            if (genFn0("", (v22 % (((v22 * 4294967295)) | 1)), 2) >= 1) {
                                var v38 int = 9223372036854775807
                                _ = v38
                            } else {
                                v13 -= (v13 << (((v22 & v22)) & 63))
                                v22 = 581
                            }
                        } else {
                            v22 -= v22
                            _ = len(v3)
                        }
                    }
                    v5 -= v13
                }
                _ = v32
            }
        case 29:
            v3[4294967295] = true
        }
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
