package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 int8; F1 string; F2 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint8) string {
    v5 := GenS1{F0: 1000, F1: 0, F2: (true && p0)}
    if (!(128 != (^(p1)))) {
        fmt.Println(p0, p0)
        p0 = true
    }
    _ = p0
    _ = p1
    _ = v5
    return ""
}
func genFn1(p0 int32, p1 uint64) uint {
    switch ("hello" + "a") {
    case "\n\t":
        var v6 float64 = (3.14159 / (float64(0.0) + (1e10 + -2.25)))
        _ = v6
    case "a":
        p0 = -2147483648
    }
    var v7 int32 = (p0 | (p0 >> ((820) & 63)))
    if (-747 <= 218) {
        v7 += 42
        p1 = p1
    } else {
        var v8 float64 = (0.0 - 1e10)
        v9 := []int{4294967295, -2147483648}
        _ = v8
        _ = v9
    }
    v7 = v7
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 32767
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    v2 = (v2 | v2)
    _ = len(v4)
    if true {
        v2 += -1
        if (1 == 42) {
            v10 := GenS1{F0: v2, F1: 42, F2: (-1 == 42)}
            _ = v10
        } else {
            v2 -= v2
        }
        if false {
            v2 += (v2 ^ v2)
        } else {
            v4["foo bar"] = 42
            var v11 uint32 = 264
            var v12 uint64 = 1
            _ = v11
            _ = v12
        }
    } else {
        if false {
            var v13 uint8 = 2
            v14 := GenS1{F0: v2, F1: 255, F2: true}
            v13 += 1
            _ = v13
            _ = v14
        }
    }
    var v15 int32 = -1
    fmt.Println(v2)
    var v16 float64 = (float64((-2.25 / (1e10 / 3.14159))) * 3.14159)
    v17 := GenS1{F0: v2, F1: 2, F2: true}
    switch -9223372036854775808 {
    case -2147483648:
        fmt.Println(v16, v2, v2)
    case 127:
        v2 -= 1000
    case 595:
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 int64 = 42
            var v20 int32 = v15
            for fi21 := 0; fi21 < 16; fi21++ {
                fmt.Println(v19, v15, v20)
                var v22 uint16 = 255
                fmt.Println(fi21)
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        v15 -= v15
    }
    v2 = (v2 * (v2 ^ v2))
    v23 := GenS1{F0: v2, F1: 155, F2: (!(1 >= -45))}
    var v24 uint32 = 2
    switch -9223372036854775808 {
    case -72:
        if false {
            var v25 int8 = 78
            if false {
                if false {
                    if (819 < 471) {
                        if (128 <= 1) {
                            for fi26 := 0; fi26 < 2; fi26++ {
                                fmt.Println(v24, v16)
                                switch 128 {
                                case 9223372036854775807:
                                    v1["\n\t"] = ("\n\t" + "a")
                                case 4294967295:
                                    v27 := GenS1{F0: (v2 & (v2 - v2)), F1: 872, F2: ((v24 - 65535) > 2147483647)}
                                    _ = v27
                                case 128:
                                    if (1 <= -1) {
                                        v28 := GenS1{F0: (-(v2)), F1: 32767, F2: (4294967295 <= 65535)}
                                        for fi29 := 0; fi29 < 1; fi29++ {
                                            v15 = int32(fi26)
                                            var v30 int64 = -1
                                            _ = fi29
                                            _ = v30
                                        }
                                        if false {
                                            for fi31 := 0; fi31 < 14; fi31++ {
                                                v2 += v2
                                                _ = fi31
                                            }
                                        } else {
                                            var v32 uint32 = (v24 >> ((1) & 63))
                                            _ = v32
                                        }
                                        _ = v28
                                    }
                                }
                                _ = fi26
                            }
                        } else {
                            if true {
                                v24 = (v24 / ((v24) | 1))
                            }
                            var v33 uint = 65535
                            _ = v33
                        }
                    } else {
                        v34 := GenS2{F0: (v25 + 127), F1: "", F2: v24}
                        _ = v34
                    }
                }
            }
            _ = v25
        }
    default:
        var v35 uint = uint(v24)
        _ = v35
    }
    var v36 float64 = (v16 / v16)
    var v37 uint64 = 0
    v38 := GenS3{F0: (v37 << ((42) & 63))}
    var v39 bool = (-9223372036854775808 != 0)
    var v40 uint64 = v37
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
