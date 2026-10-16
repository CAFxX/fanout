package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 bool, p1 string, p2 int) uint {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(p1, fi7)
            var v8 int = 1
            var v9 bool = p0
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int32 = 579
            v11 = (v11 % (((v11 & v11)) | 1))
            p0 = false
            p2 += (fi10 * (fi10 >> ((42) & 63)))
            v12 := []int{fi6, (-(p2)), fi10, (p2 | 127)}
            _ = fi10
            _ = v11
            _ = v12
        }
        if true {
            if true {
                v13 := GenS1{F0: 180, F1: 0, F2: 42}
                v14 := GenS2{F0: 2}
                p0 = p0
                _ = v13
                _ = v14
            } else {
                fmt.Println(p1, p0, fi6)
            }
        } else {
            fmt.Println(p1)
            p1 = p1
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 477
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{}
    v3 := make(map[int]bool)
    var v4 uint = 255
    var v5 uint64 = 4294967295
    if ((false && (false && (1 != 101))) && (v5 != 1)) {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 string = "\n\t"
            _ = len(v2)
            _ = fi15
            _ = v16
        }
    } else {
        var v17 int8 = -1
        switch -913 {
        case 368:
            fmt.Println(v1, v5, v4)
        case 794:
            if (0 >= 127) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (false || false) {
                    fmt.Println(v17)
                } else {
                    switch 442 {
                    case 128:
                        for fi18 := 0; fi18 < 4; fi18++ {
                            fmt.Println(v4, v4)
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            var v19 string = "a"
                            _ = fi18
                            _ = v19
                        }
                    case -205:
                        v4 -= v4
                    default:
                        var v20 int64 = (v1 << ((v4) & 63))
                        _ = v20
                    }
                    var v21 uint = genFn0(false, "\n\t", 128)
                    _ = v21
                }
            }
        default:
            fmt.Println(v5, v17)
        }
        _ = v17
    }
    fmt.Println(v4, v1, v1)
    fmt.Println(v5)
    v22 := GenS3{F0: 42, F1: 347, F2: v5}
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 uint64 = v5
        if (((float32(1.5) != 3.14159) || false) && false) {
            v5 = v24
            if true {
                var v25 float32 = (float32(float64(-2.25)) * 1e10)
                _ = v25
            } else {
                if (((3.14159 - 3.14159) + 1e10) != (1e10 / 0.1)) {
                    var v26 string = ""
                    v27 := GenS3{F0: 664, F1: 2, F2: (v5 + 714)}
                    for fi28 := 0; fi28 < 17; fi28++ {
                        switch fi23 {
                        case 2147483647:
                            v29 := GenS3{F0: 0, F1: 0, F2: v24}
                            _ = v29
                        case 1:
                            _ = len(v3)
                        case 9223372036854775807:
                            v4 = 4294967295
                        default:
                            v1 += int64(fi23)
                        }
                        _ = fi28
                    }
                    _ = v26
                    _ = v27
                } else {
                    if true {
                        switch 42 {
                        case -2147483648:
                            if len(v2) > 0 {
                                v2[0] = ("" + "foo bar")
                            }
                        default:
                            v1 += 377
                        }
                        v1 = v1
                        for fi30 := 0; fi30 < 19; fi30++ {
                            v31 := GenS3{F0: 2, F1: 660, F2: (v24 ^ (v24 + v24))}
                            v24 -= v5
                            _ = fi30
                            _ = v31
                        }
                    }
                    var v32 float64 = 0.0
                    _ = v32
                }
                v1 = (v1 & -1)
            }
            var v33 int64 = v1
            _ = v33
        }
        _ = fi23
        _ = v24
    }
    var v34 int64 = v1
    var v35 uint64 = v5
    _ = len(v2)
    switch 438 {
    case 1:
        var v36 int16 = -32768
        _ = v36
    case 32767:
        var v37 float32 = -2.25
        _ = v37
    case 32768:
        var v38 uint = 1
        _ = v38
    default:
        var v39 int32 = 42
        _ = v39
    }
    var v40 int = -543
    fmt.Println(v1)
    v40 = (^(v40))
    v41 := GenS4{F0: 255, F1: (v40 ^ v40)}
    v3[32768] = (true || true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
