package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int16) uint64 {
    v6 := []int{-500}
    if len(v6) > 0 {
        v6[0] = -1
    }
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1(p0 uint8, p1 int32, p2 uint8) uint64 {
    p2 = p0
    fmt.Println(p1)
    fmt.Println(p2)
    switch (("\n\t" + "\n\t") + "a") {
    case "hello":
        fmt.Println(p1)
    case "":
        for fi7 := 0; fi7 < 6; fi7++ {
            p1 = p1
            _ = fi7
        }
    case "\n\t":
        fmt.Println(p0, p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return genFn0(32767)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    v4 := []uint32{2147483648, 0}
    var v5 uint16 = 255
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := v3[(^(fi9))]
            var v11 uint = 65535
            v11 += 355
            v10 = v10
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: true, F1: -505}
        _ = fi8
        _ = v12
    }
    v5 += (^(v5))
    for fi13 := 0; fi13 < 15; fi13++ {
        switch fi13 {
        case 9223372036854775807:
            var v14 bool = ((!true) && true)
            _ = v14
        case -369:
            var v15 uint8 = 17
            _ = v15
        default:
            var v16 int = (fi13 << ((0) & 63))
            _ = v16
        }
        _ = fi13
    }
    delete(v1, 580)
    switch ("hello" + "\n\t") {
    case "\n\t":
        if true {
            delete(v3, -9223372036854775808)
        } else {
            v17 := GenS2{F0: 0}
            v5 = (v5 - v5)
            v18 := GenS2{F0: 1}
            _ = v17
            _ = v18
        }
    case "hello":
        fmt.Println(v5)
    case "a":
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 float32 = (1e10 / 0.1)
                v5 = v5
                var v22 int16 = 1
                _ = fi20
                _ = v21
                _ = v22
            }
            if false {
                fmt.Println(fi19, v5)
            } else {
                v5 = 547
                var v23 string = ""
                _ = v23
            }
            _ = fi19
        }
    default:
        v5 = 106
    }
    var v24 int64 = 671
    var v25 float64 = 0.1
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 bool = true
        switch fi26 {
        case 1:
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 int16 = -1
                _ = fi28
                _ = v29
            }
        case -353:
            fmt.Println(v25, v25, v25)
        case 32767:
            var v30 int8 = -76
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    v5 = (v5 % (((v5 - v5)) | 1))
    v5 = 874
    v31 := GenS3{F0: (v24 << ((4294967295) & 63))}
    var v32 uint = 1
    v25 = (v25 + 1.5)
    if (-522 != -818) {
        if (!((2 > v32) || (!(-1 == -32768)))) {
            var v33 uint16 = (^(v5))
            _ = v33
        } else {
            v34 := GenS2{F0: 128}
            v24 -= (v24 >> ((v32) & 63))
            _ = v34
        }
        fmt.Println(v5, v25, v32)
    } else {
        v1[918] = 65535
        if (true && true) {
            var v35 uint16 = v5
            var v36 uint16 = (v5 & v5)
            v25 = ((v25 + v25) * ((3.14159 / (1.5 * v25)) - v25))
            _ = v35
            _ = v36
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                v32 -= v32
                for fi38 := 0; fi38 < 18; fi38++ {
                    if (1 != 128) {
                        fmt.Println(fi38, v24)
                        v39 := GenS3{F0: (v24 << ((0) & 63))}
                        _ = v39
                    } else {
                        fmt.Println(v5, v5)
                        v5 = (v5 ^ 245)
                    }
                    _ = fi38
                }
                _ = fi37
            }
        }
    }
    v40 := GenS4{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
}
