package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint }
func genFn0() uint {
    var v7 uint8 = 156
    if (!false) {
        v7 += (^(v7))
        if (!((740 == 0) && false)) {
            var v8 uint64 = 2
            v9 := []int{32767, -864}
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: 2147483648, F1: 127}
                if (1 != -128) {
                    v7 = 42
                    fmt.Println(v8)
                    if true {
                        if len(v9) > 0 {
                            v9[0] = fi10
                        }
                        fmt.Println(fi10)
                        if (false && (fi10 > fi10)) {
                            fmt.Println(v8, fi10)
                            v8 = (^(v8))
                            var v12 float64 = 0.1
                            _ = v12
                        }
                    }
                }
                _ = len(v9)
                _ = fi10
                _ = v11
            }
            _ = v8
            _ = v9
        } else {
            for fi13 := 0; fi13 < 3; fi13++ {
                switch fi13 {
                case -839:
                    if true {
                        if true {
                            fmt.Println(fi13)
                            v14 := []int{4294967295, 2147483647, (-(fi13))}
                            _ = v14
                        }
                    }
                default:
                    v7 += v7
                }
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int16 = -32768
                    switch fi13 {
                    case 42:
                        v17 := []int{fi13, fi15, 2147483647, (-(fi15))}
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                switch 4294967295 {
                case 1:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v19 := GenS2{F0: 837, F1: 0}
                        var v20 int = fi13
                        if false {
                            if ("" <= "foo bar") {
                                var v21 uint16 = 788
                                switch (fi13 & 4294967295) {
                                case -1:
                                    v21 -= v21
                                case 0:
                                    var v22 int16 = 358
                                    _ = v22
                                case 32768:
                                    if true {
                                        var v23 int8 = 1
                                        fmt.Println(v23, v21, v21)
                                        switch "a" {
                                        case "a":
                                            var v24 int8 = v23
                                            _ = v24
                                        }
                                        _ = v23
                                    }
                                }
                                v25 := []int{}
                                _ = v21
                                _ = v25
                            } else {
                                v26 := []int{v20}
                                _ = v26
                            }
                            var v27 uint16 = 240
                            fmt.Println(fi13, v27)
                            var v28 int = (fi18 ^ fi18)
                            var v29 int64 = 1
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS2{F0: 2147483647, F1: 0}
                            v31 := []int{0, (fi13 << ((1) & 63))}
                            if false {
                                v7 += 42
                            }
                            _ = v30
                            _ = v31
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                default:
                    var v32 uint64 = 276
                    _ = v32
                }
                _ = fi13
            }
        }
    }
    v33 := []int{127, 42, -2147483648, int(v7)}
    v34 := GenS3{F0: 4294967295}
    _ = v33
    _ = v34
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    v3 := []uint32{2147483648, 1, 0}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        if true {
            var v36 int = -584
            if len(v3) > 0 {
                v3[0] = 0
            }
            fmt.Println(v36)
            _ = v36
        }
        _ = fi35
    }
    v5[65535] = -9223372036854775808
    for fi37 := 0; fi37 < 3; fi37++ {
        v38 := GenS1{F0: 738, F1: 42}
        _ = fi37
        _ = v38
    }
    if true {
        delete(v6, "")
        var v39 uint64 = 1
        _ = v39
    } else {
        var v40 int64 = 1
        _ = v40
    }
    var v41 float32 = 0.0
    fmt.Println(v41)
    fmt.Println(v41)
    var v42 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
