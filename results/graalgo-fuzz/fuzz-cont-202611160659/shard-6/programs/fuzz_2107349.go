package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8 }
func genFn0() int16 {
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int64 = 42
        _ = fi8
        _ = v9
    }
    var v10 bool = true
    v10 = (1 > 612)
    _ = v10
    _ = v7
    return -167
}
func genFn1(p0 string, p1 uint32) int16 {
    var v11 uint32 = (p1 / (((p1 + (p1 + p1))) | 1))
    var v12 uint64 = 147
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 bool = (0 > 1)
        var v15 int64 = -454
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return -32768
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int64 = -1
    var v5 bool = (!true)
    var v6 int64 = v4
    delete(v3, 65535)
    if true {
        var v16 int = 42
        _ = v16
    } else {
        if ("\n\t" > "foo bar") {
            var v17 int32 = -1
            v5 = (0 != 127)
            _ = v17
        }
        if (0.1 <= float64(float32(1e10))) {
            fmt.Println(v4, v6, v4)
            v18 := []int{2147483647, 4294967295, 4294967295}
            _ = v18
        } else {
            var v19 uint8 = 0
            v20 := GenS1{F0: -103}
            v5 = (!v5)
            _ = v19
            _ = v20
        }
    }
    v21 := GenS1{F0: 100}
    v4 -= v6
    v6 += (-(v4))
    if true {
        if ((v6 >> ((827) & 63)) < v4) {
            v22 := GenS2{F0: 502, F1: 42}
            _ = v22
        } else {
            v3[589] = -1
            fmt.Println(v4, v1)
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 float32 = 3.14159
                for fi25 := 0; fi25 < 13; fi25++ {
                    if v5 {
                        v5 = ((4294967295 != 0) && v5)
                        v1 -= (v1 | (v1 << ((1) & 63)))
                    } else {
                        v6 = (v4 + (v4 + 9223372036854775807))
                        switch -882 {
                        case -48:
                            v6 -= 9223372036854775807
                        case -655:
                            v4 = (v4 * (v4 | (^(v4))))
                        case -9223372036854775808:
                            v26 := v3[2]
                            _ = v26
                        }
                        switch "" {
                        case "":
                            var v27 uint32 = 695
                            fmt.Println(fi23)
                            var v28 int64 = v4
                            _ = v27
                            _ = v28
                        case "foo bar":
                            var v29 int32 = -757
                            _ = v29
                        case "hello":
                            v30 := []int{fi23}
                            _ = len(v3)
                            _ = v30
                        }
                    }
                    delete(v3, 736)
                    _ = fi25
                }
                _ = fi23
                _ = v24
            }
        }
    } else {
        switch "foo bar" {
        case "hello":
            var v31 int64 = (v6 / (((v6 - v4)) | 1))
            _ = v31
        default:
            var v32 int8 = -1
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 int64 = v4
        _ = fi33
        _ = v34
    }
    if v5 {
        var v35 uint = 4294967295
        _ = v35
    }
    var v36 int64 = 9223372036854775807
    switch ("a" + "\n\t") {
    case "foo bar":
        var v37 int8 = -120
        _ = v37
    default:
        var v38 uint = 2
        _ = v38
    }
    var v39 int32 = -794
    v40 := GenS1{F0: 1}
    v41 := GenS3{F0: 0}
    var v42 uint8 = 53
    var v43 uint8 = 204
    v36 = (v36 ^ 0)
    if v5 {
        v6 -= v4
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        fmt.Println(v4, fi44)
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
