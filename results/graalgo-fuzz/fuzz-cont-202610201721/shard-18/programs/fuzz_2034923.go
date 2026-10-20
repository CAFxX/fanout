package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
func genFn0() int {
    if (-1 <= 0) {
        if true {
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS1{F0: 754, F1: 582, F2: true}
                fmt.Println(fi7)
                _ = fi7
                _ = v8
            }
        } else {
            v9 := GenS1{F0: 65535, F1: 2147483648, F2: (255 >= 65535)}
            var v10 int32 = 249
            _ = v10
            _ = v9
        }
        var v11 uint = 255
        _ = v11
    } else {
        v12 := GenS1{F0: 42, F1: 0, F2: true}
        var v13 int = -9223372036854775808
        _ = v12
        _ = v13
    }
    if (-1 >= 588) {
        v14 := []int{128, -9223372036854775808, 128}
        var v15 int32 = -1
        _ = v14
        _ = v15
    }
    var v16 int8 = -128
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: 817, F1: 0, F2: false}
        _ = fi17
        _ = v18
    }
    _ = v16
    return 32768
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    var v4 uint32 = 1
    var v5 uint = uint(v3)
    var v6 int64 = 9223372036854775807
    if false {
        v4 += (v4 << ((853) & 63))
    } else {
        v1 -= (v1 + v1)
        v4 = v4
    }
    v1 = v1
    v6 = (v6 / ((v6) | 1))
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS1{F0: (v5 | (^(v5))), F1: (^(v4)), F2: (1 == -30)}
        var v21 uint8 = 1
        v22 := []int{genFn0(), fi19, 127, genFn0()}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    if false {
        var v23 int8 = -1
        if (!(v3 > v3)) {
            v24 := GenS2{F0: -949}
            var v25 int64 = (v6 % (((v6 / ((-9223372036854775808) | 1))) | 1))
            _ = v24
            _ = v25
        } else {
            v2[v4] = v4
        }
        var v26 int16 = 0
        _ = v23
        _ = v26
    } else {
        var v27 int8 = 0
        for fi28 := 0; fi28 < 20; fi28++ {
            if false {
                if (v1 == v1) {
                    switch "\n\t" {
                    case "hello":
                        v6 += (v6 - 0)
                    default:
                        v29 := []int{fi28, fi28, fi28}
                        _ = v29
                    }
                    if true {
                        if false {
                            v30 := GenS2{F0: -980}
                            v5 = 18446744073709551615
                            _ = v30
                        } else {
                            var v31 int32 = 42
                            _ = v31
                        }
                        var v32 uint32 = (^(v4))
                        _ = v32
                    }
                }
                fmt.Println(v4, fi28)
            }
            if true {
                if (!true) {
                    if (false && true) {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            var v34 uint64 = 18446744073709551615
                            v3 = (^(v3))
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        for fi35 := 0; fi35 < 18; fi35++ {
                            var v36 int64 = (v6 / ((v6) | 1))
                            v37 := []int{fi28}
                            fmt.Println(fi35)
                            var v38 int = fi35
                            v6 -= 42
                            var v39 uint64 = 1
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        if ((!false) && ((v1 >> ((v5) & 63)) < v1)) {
                            var v40 uint64 = 496
                            var v41 uint16 = (^(v1))
                            _ = v40
                            _ = v41
                        } else {
                            v42 := GenS1{F0: v5, F1: 0, F2: false}
                            for fi43 := 0; fi43 < 12; fi43++ {
                                if false {
                                    var v44 uint16 = v1
                                    _ = v44
                                }
                                _ = fi43
                            }
                            for fi45 := 0; fi45 < 7; fi45++ {
                                fmt.Println(fi45, v1, v5)
                                _ = fi45
                            }
                            _ = v42
                        }
                    }
                } else {
                    var v46 uint = v5
                    if true {
                        v47 := GenS1{F0: v5, F1: (v4 / (((v4 >> ((1) & 63))) | 1)), F2: true}
                        var v48 uint64 = 0
                        v49 := GenS2{F0: (fi28 % (((fi28 % (((fi28 << (((v46 ^ 2)) & 63))) | 1))) | 1))}
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    v50 := GenS3{F0: 89}
                    _ = v46
                    _ = v50
                }
                v51 := GenS1{F0: (^(v5)), F1: v4, F2: false}
                var v52 int32 = -112
                _ = v51
                _ = v52
            }
            _ = fi28
        }
        _ = v27
    }
    v1 = (v1 ^ (v1 >> ((4294967295) & 63)))
    v53 := v2[v4]
    v3 = v3
    v54 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
