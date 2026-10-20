package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 int; F2 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16) bool {
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 += 42
        var v7 uint32 = uint32(p0)
        _ = fi6
        _ = v7
    }
    v8 := []int{42, 1}
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            if len(v8) > 0 {
                v8[0] = fi9
            }
        }
        var v10 float64 = float64(0.1)
        fmt.Println(p0, fi9, fi9)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        p0 += 0
        _ = fi11
    }
    var v12 uint32 = 1
    _ = p0
    _ = v12
    _ = v8
    return false
}
func genFn1(p0 uint, p1 uint8, p2 int) int8 {
    var v13 string = "a"
    v13 = (v13 + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -54
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{32768}
    var v3 uint16 = 1
    var v4 int32 = 120
    var v5 uint8 = 0
    v5 += 127
    switch 387 {
    case 549:
        v3 += v3
    }
    var v14 uint16 = v3
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: 0, F1: (v4 | -2147483648), F2: (-(v4))}
        if false {
            if false {
                var v17 float32 = 3.14159
                v2 = append(v2, fi15)
                _ = v17
            }
            if false {
                var v18 uint16 = v3
                v19 := GenS1{F0: 0, F1: v4, F2: v4}
                v14 = v3
                _ = v18
                _ = v19
            } else {
                _ = len(v1)
            }
            v20 := GenS1{F0: 2, F1: (v4 % ((v4) | 1)), F2: v4}
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    if false {
        for fi21 := 0; fi21 < 7; fi21++ {
            if (-32768 > 0) {
                switch (("hello" + "foo bar") + (("\n\t" + "hello") + "")) {
                case "\n\t":
                    var v22 uint64 = 524
                    _ = v22
                default:
                    fmt.Println(v4)
                }
            } else {
                for fi23 := 0; fi23 < 18; fi23++ {
                    if len(v2) > 0 {
                        v2[0] = fi23
                    }
                    fmt.Println(fi21, v4, v3)
                    _ = fi23
                }
                v3 = v3
                var v24 int16 = 0
                _ = v24
            }
            _ = fi21
        }
    } else {
        if ((v4 / (((v4 + v4)) | 1)) > v4) {
            fmt.Println(v4)
        }
        v25 := GenS2{F0: -42, F1: 257, F2: -1}
        var v26 uint8 = v5
        _ = v25
        _ = v26
    }
    _ = len(v2)
    v27 := GenS1{F0: 670, F1: v4, F2: v4}
    fmt.Println(v14)
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 uint32 = 42
        if false {
            var v30 string = "hello"
            _ = v30
        } else {
            if true {
                switch "a" {
                case "a":
                    var v31 int = (fi28 ^ fi28)
                    _ = v31
                case "foo bar":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        switch fi28 {
                        case 1:
                            v4 = (v4 >> ((18446744073709551615) & 63))
                        }
                        _ = fi32
                    }
                case "":
                    var v33 int64 = 42
                    _ = v33
                }
                var v34 int32 = (v4 % (((v4 | -2147483648)) | 1))
                fmt.Println(v5, v29)
                _ = v34
            }
            v35 := GenS1{F0: 18446744073709551615, F1: v4, F2: v4}
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 int64 = 935
                v4 = (v4 >> ((18446744073709551615) & 63))
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        _ = fi28
        _ = v29
    }
    v14 = 2
    var v38 float64 = (0.0 / (-2.25 + (-2.25 + 0.0)))
    if ((!false) || false) {
        for fi39 := 0; fi39 < 12; fi39++ {
            fmt.Println(fi39)
            _ = fi39
        }
        v14 = 2
    }
    switch 566 {
    case 9223372036854775807:
        if (!false) {
            delete(v1, 4294967295)
            var v40 int = 573
            _ = v40
        }
    case 32768:
        var v41 uint = 18446744073709551615
        _ = v41
    }
    v42 := GenS3{F0: 2147483647}
    v2 = append(v2, 4294967295)
    fmt.Println(v5)
    if false {
        var v43 int8 = 127
        switch 42 {
        case 128:
            var v44 float32 = 1.5
            _ = v44
        case -9223372036854775808:
            v14 -= uint16(v5)
        }
        _ = v43
    } else {
        v38 += v38
        var v45 string = ("\n\t" + "")
        _ = v45
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        v2 = append(v2, 624)
        if (v14 == v3) {
            v14 -= (v3 & v3)
            v14 += 65535
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
