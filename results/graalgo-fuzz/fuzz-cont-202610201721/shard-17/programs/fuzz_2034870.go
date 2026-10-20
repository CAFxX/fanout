package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 uint16) uint16 {
    var v5 uint32 = 2147483648
    p1 = 42
    v5 = v5
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 uint32, p1 int8, p2 int) uint16 {
    if true {
        fmt.Println(p1, p2)
    }
    switch "foo bar" {
    case "foo bar":
        p2 = 348
    case "hello":
        var v6 uint32 = (p0 + (p0 >> ((1) & 63)))
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return uint16(p0)
}

func main() {
    var v1 int16 = 1
    var v2 int8 = -42
    var v3 float32 = (0.0 / 0.1)
    var v4 uint64 = 2
    v2 += v2
    fmt.Println(v2, v1)
    var v7 int = -523
    var v8 uint16 = 1
    var v9 int32 = -1
    if false {
        v3 = (float32(1e10) - float32((0.1 * 3.14159)))
        var v10 int16 = -126
        _ = v10
    } else {
        v7 = (v7 * v7)
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float32 = 3.14159
            v8 += 1
            v13 := []int{v7, (^(fi11)), (fi11 & -9223372036854775808), fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v14 := []int{v7, 2147483647}
    v8 = (v8 ^ 65535)
    var v15 bool = false
    if (!v15) {
        var v16 uint16 = (^(v8))
        switch "a" {
        case "foo bar":
            var v17 float64 = 0.1
            _ = v17
        case "hello":
            var v18 uint16 = (^(v16))
            _ = v18
        default:
            v15 = (v15 && ((255 <= 255) && v15))
        }
        _ = v16
    } else {
        var v19 int = (v7 % ((v7) | 1))
        v20 := GenS1{F0: genFn0(255, v8)}
        if v15 {
            v19 = (v19 << ((255) & 63))
        }
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        switch ("" + "a") {
        case "hello":
            fmt.Println(v7, fi21, v15)
        }
        if (genFn0(301, genFn1(0, v2, (fi21 - (-(v7))))) <= (v8 / ((genFn1(42, 42, fi21)) | 1))) {
            v9 = v9
            fmt.Println(v1, v15, v3)
            var v22 string = ("" + ((("\n\t" + "a") + "") + ("a" + "\n\t")))
            _ = v22
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint8 = 255
                var v25 uint8 = (v24 ^ 172)
                fmt.Println(v7, v8)
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        v14 = append(v14, fi21)
        _ = fi21
    }
    v26 := GenS1{F0: 255}
    switch 127 {
    case 740:
        var v27 uint64 = v4
        _ = v27
    case 869:
        fmt.Println(v7, v7)
    case 127:
        if v15 {
            var v28 string = "a"
            if (42 != 4294967295) {
                if v15 {
                    v3 = 1.5
                    fmt.Println(v2, v9, v3)
                }
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 uint = 42
                    var v31 int8 = int8(v9)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            _ = v28
        }
    }
    var v32 int64 = 1
    switch 692 {
    case -606:
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 float64 = float64(((v3 / v3) / (float32((3.14159 + 3.14159)) + (((1e10 / 3.14159) - v3) * (v3 + ((float32(0.0) * (0.0 * v3)) - v3))))))
            v35 := GenS1{F0: (v8 << ((225) & 63))}
            _ = fi33
            _ = v34
            _ = v35
        }
    case -785:
        fmt.Println(v8, v9, v32)
    default:
        if (2 < 4294967295) {
            for fi36 := 0; fi36 < 18; fi36++ {
                if v15 {
                    var v37 float32 = v3
                    for fi38 := 0; fi38 < 13; fi38++ {
                        switch fi36 {
                        case 459:
                            v39 := GenS1{F0: v8}
                            _ = v39
                        default:
                            if (0.1 > float64(-2.25)) {
                                fmt.Println(v37, v7)
                            }
                        }
                        switch "foo bar" {
                        case "hello":
                            var v40 int8 = 42
                            _ = v40
                        }
                        fmt.Println(v9, v4)
                        _ = fi38
                    }
                    _ = v37
                } else {
                    v41 := GenS1{F0: (v8 & 344)}
                    v32 += -9223372036854775808
                    var v42 uint16 = (v8 * v8)
                    _ = v41
                    _ = v42
                }
                v43 := GenS2{F0: v15}
                var v44 uint = 255
                _ = fi36
                _ = v43
                _ = v44
            }
        }
    }
    v15 = true
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
