package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int, p1 uint32) bool {
    v6 := []int{}
    if ((1.5 + (((0.0 - 3.14159) / 0.1) * 0.0)) == 0.0) {
        var v7 int = p0
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(p0, p0)
            p0 = (v7 >> ((160) & 63))
            _ = fi8
        }
        p1 += (p1 >> ((911) & 63))
        _ = v7
    }
    v6 = append(v6, 9223372036854775807)
    _ = p0
    _ = p1
    _ = v6
    return false
}
func genFn1(p0 int16) bool {
    p0 -= (p0 % (((p0 - p0)) | 1))
    var v9 int16 = 0
    v10 := []int{128, 127}
    fmt.Println(v9)
    _ = p0
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 int16 = 42
    v5 := make(map[string]int)
    if false {
        v1 += v1
        v1 = 1
    } else {
        var v11 float64 = 1.5
        v11 += (0.1 * v11)
        _ = v11
    }
    var v12 float64 = 1e10
    if (!false) {
        if (42 <= 4294967295) {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v12, v4)
            case "hello":
                v1 += 1
            case "\n\t":
                fmt.Println(v1, v1)
            }
            fmt.Println(v4, v4)
            v1 = v1
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 924, F1: v1}
            v1 -= 34
            _ = fi13
            _ = v14
        }
    } else {
        if (!true) {
            v4 = v4
            var v15 float64 = v12
            v16 := GenS1{F0: 896, F1: (v1 + (v1 >> ((42) & 63)))}
            _ = v15
            _ = v16
        }
        v17 := []int{4294967295, -1, 128}
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int64 = 42
        _ = fi18
        _ = v19
    }
    var v20 float64 = (-2.25 - v12)
    var v21 uint16 = 255
    if false {
        v2[937] = "a"
        v21 = 255
    }
    fmt.Println(v21, v20)
    if (349 > -677) {
        var v22 float32 = (float32(0.0) + 0.0)
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                v22 -= -2.25
                _ = fi24
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                if true {
                    switch (fi25 - (fi23 * -177)) {
                    case 4294967295:
                        if false {
                            v26 := []int{0, 128, fi23, fi25}
                            if false {
                                v27 := GenS1{F0: uint32(v1), F1: v1}
                                v28 := GenS2{F0: -2147483648, F1: 336}
                                delete(v5, "")
                                var v29 int8 = -128
                                v30 := GenS3{F0: v21}
                                v31 := v2[(fi23 ^ -734)]
                                if true {
                                    if false {
                                        v32 := GenS1{F0: 428, F1: 1}
                                        if true {
                                            v29 = v29
                                        } else {
                                            _ = len(v26)
                                            _ = len(v5)
                                            v33 := GenS2{F0: 1, F1: 42}
                                            v34 := GenS2{F0: -2147483648, F1: 65535}
                                            _ = v33
                                            _ = v34
                                        }
                                        if false {
                                            v4 = -1
                                            v12 = (v20 * v20)
                                            v21 = (^(v21))
                                            v35 := GenS2{F0: 500, F1: 845}
                                            _ = v35
                                        }
                                        v36 := v2[(fi25 - -2147483648)]
                                        _ = v32
                                        _ = v36
                                    }
                                }
                                var v37 string = (v31 + ("hello" + v31))
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v37
                            }
                            _ = v26
                        }
                    case -11:
                        var v38 bool = ("foo bar" < "foo bar")
                        _ = v38
                    default:
                        var v39 int64 = 9223372036854775807
                        _ = v39
                    }
                } else {
                    var v40 uint8 = 48
                    var v41 int16 = v4
                    _ = v40
                    _ = v41
                }
                _ = fi25
            }
            _ = fi23
        }
        _ = v22
    } else {
        v1 = v1
    }
    var v42 int64 = int64(v4)
    var v43 string = "\n\t"
    switch 128 {
    case -9223372036854775808:
        v44 := []int{15, 32767, 2147483647}
        _ = v44
    case 9223372036854775807:
        var v45 uint64 = 664
        _ = v45
    }
    var v46 uint = 878
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v5
}
