package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = (630 <= 2147483647)
    var v2 string = ("foo bar" + "")
    var v3 bool = v1
    var v4 uint8 = 183
    var v5 uint8 = 0
    v6 := make(map[uint32]uint32)
    if true {
        v7 := GenS1{F0: 2}
        var v8 int8 = 0
        fmt.Println(v2, v1, v3)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint32 = 2
            if v1 {
                v1 = ((v11 << ((uint(v11)) & 63)) > (v11 * v11))
                if (1.5 <= 1.5) {
                    switch v2 {
                    case "":
                        v11 += (v11 >> ((0) & 63))
                        fmt.Println(v2, v2)
                    default:
                        v12 := GenS1{F0: (v11 * (v11 - 1))}
                        _ = v12
                    }
                    var v13 uint16 = 32767
                    _ = v13
                } else {
                    var v14 int16 = -1
                    _ = v14
                }
                v15 := GenS1{F0: v11}
                v16 := v6[v11]
                if v1 {
                    var v17 float64 = 1.5
                    if v3 {
                        if v1 {
                            var v18 uint16 = uint16(v16)
                            var v19 string = "a"
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{fi10, (fi9 * (fi9 ^ fi9)), 32768}
                        var v21 int16 = -358
                        var v22 float64 = 3.14159
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = v17
                } else {
                    if v3 {
                        if v3 {
                            fmt.Println(v3)
                            var v23 float64 = float64(3.14159)
                            _ = v23
                        }
                        var v24 int16 = 949
                        _ = v24
                    }
                    fmt.Println(v11)
                    v25 := GenS1{F0: (v16 | 519)}
                    _ = v25
                }
                v2 = ("hello" + v2)
                _ = v15
                _ = v16
            } else {
                fmt.Println(v4)
                var v26 uint = 4294967295
                if ((fi9 ^ fi10) < (fi10 << (((v26 + (v26 | 1))) & 63))) {
                    if (-999 == -32768) {
                        switch v2 {
                        case "a":
                            fmt.Println(v11)
                        default:
                            var v27 uint16 = 255
                            _ = v27
                        }
                        v28 := GenS2{F0: 42, F1: v2}
                        _ = v28
                    }
                    v29 := GenS1{F0: (v11 / (((v11 & v11)) | 1))}
                    fmt.Println(v2, v4, v5)
                    _ = v29
                }
                _ = v26
            }
            var v30 uint8 = v5
            _ = fi10
            _ = v11
            _ = v30
        }
        var v31 bool = true
        v32 := []int{(fi9 + fi9), fi9, -909}
        _ = fi9
        _ = v31
        _ = v32
    }
    var v33 string = "foo bar"
    switch 281 {
    case 127:
        v34 := GenS1{F0: 48}
        _ = v34
    case -9223372036854775808:
        if (42 < 2) {
            v35 := GenS2{F0: -746, F1: v33}
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS3{F0: 4294967295}
                var v38 uint32 = 42
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v35
        } else {
            var v39 float32 = (0.1 / 1e10)
            _ = len(v6)
            var v40 string = v2
            _ = v39
            _ = v40
        }
    }
    var v41 int32 = 0
    switch v33 {
    case "foo bar":
        var v42 int = 1
        _ = v42
    default:
        fmt.Println(v3)
    }
    var v43 int16 = -32768
    _ = len(v6)
    fmt.Println(v41, v3, v43)
    v41 += (v41 / (((v41 & (v41 ^ 579))) | 1))
    _ = len(v6)
    v44 := []int{-15, 0, 127, 2147483647}
    var v45 uint = 18446744073709551615
    v41 += 42
    v46 := GenS2{F0: (v41 / (((v41 % (((-(v41))) | 1))) | 1)), F1: ("\n\t" + "a")}
    if v1 {
        fmt.Println(v41, v45)
    } else {
        v45 -= v45
        var v47 float32 = 1.5
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
