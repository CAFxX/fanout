package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 int16 }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 int64; F2 int }
type GenS6 struct { F0 bool; F1 int }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 180, F1: 439, F2: 0}
        var v8 int8 = 100
        v9 := GenS1{F0: 755, F1: 4294967295, F2: (^(v8))}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    var v11 bool = (!((((0.1 + float64(1e10)) / 1.0) * ((-2.25 / 1.0) * 0.1)) > 1e10))
    v12 := []int{}
    fmt.Println(v11)
    _ = v11
    _ = v12
    return 2
}
func genFn1(p0 int64, p1 string, p2 string) uint {
    v13 := GenS1{F0: 2, F1: 65535, F2: 110}
    p1 = p1
    fmt.Println(p1, p2)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    var v1 bool = true
    var v2 uint8 = 132
    v3 := make(map[uint32]bool)
    var v4 int16 = 1000
    var v5 uint = 18446744073709551615
    for fi14 := 0; fi14 < 13; fi14++ {
        v1 = true
        v15 := []int{2147483647}
        v16 := GenS1{F0: 1, F1: 255, F2: -79}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{458, 127}
    v4 = 0
    for fi18 := 0; fi18 < 19; fi18++ {
        if v1 {
            fmt.Println(v5)
        }
        var v19 uint64 = 42
        if (((v1 && v1) && (!true)) && (!(true && v1))) {
            if (4294967295 == 2147483647) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 bool = (v1 && (128 != v2))
                    if (((((1.5 + (float64(float32(0.1)) - -2.25)) / ((0.1 - (((3.14159 + 0.0) / (1e10 * 1.5)) + float64(-2.25))) * (((3.14159 - ((-2.25 - 1.5) / -2.25)) - float64(3.14159)) + 3.14159))) / 1.0) * 1e10) <= -2.25) {
                        v1 = ((-2.25 > 1.5) || v21)
                        switch (fi18 & fi18) {
                        case -152:
                            var v22 uint64 = (v19 & uint64(v5))
                            _ = v22
                        case -9223372036854775808:
                            var v23 uint32 = uint32(v19)
                            _ = v23
                        default:
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        }
                    } else {
                        v17 = append(v17, 2147483647)
                        v4 = (v4 % ((-1) | 1))
                        var v25 float64 = 0.0
                        _ = v25
                    }
                    _ = fi20
                    _ = v21
                }
                var v26 bool = (!(-874 > 42))
                _ = v26
            } else {
                v27 := GenS2{F0: 2, F1: 2}
                _ = v27
            }
            switch ("foo bar" + "foo bar") {
            case "hello":
                delete(v3, 2147483647)
            }
        }
        _ = fi18
        _ = v19
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        fmt.Println(fi28)
        v2 = (v2 ^ v2)
        _ = fi28
    }
    if v1 {
        v29 := GenS3{F0: 843, F1: 0, F2: (v4 << (((^(v5))) & 63))}
        if true {
            v30 := GenS4{F0: 100, F1: v2, F2: -1}
            var v31 int64 = -55
            _ = v30
            _ = v31
        } else {
            v5 += v5
            var v32 uint16 = uint16(v2)
            switch 924 {
            case -1:
                _ = len(v3)
            case 32768:
                _ = len(v3)
            default:
                if true {
                    fmt.Println(v5, v5, v1)
                } else {
                    var v33 int16 = (v4 / (((v4 << ((633) & 63))) | 1))
                    v32 += (v32 << ((v5) & 63))
                    _ = v33
                }
            }
            _ = v32
        }
        _ = v29
    }
    var v34 int64 = 554
    v35 := GenS5{F0: 4294967295, F1: (v34 - v34), F2: 32768}
    var v36 uint16 = uint16(v5)
    for fi37 := 0; fi37 < 1; fi37++ {
        fmt.Println(v36, v2)
        v1 = (!v1)
        _ = fi37
    }
    fmt.Println(v36, v34, v2)
    if ((v4 % (((v4 - v4)) | 1)) < v4) {
        switch "foo bar" {
        case "\n\t":
            v38 := GenS6{F0: false, F1: 757}
            _ = v38
        default:
            fmt.Println(v2, v36, v1)
        }
    }
    fmt.Println(v4, v5, v34)
    var v39 int32 = 2147483647
    v36 = genFn0()
    var v40 float32 = 3.14159
    var v41 uint = (v5 / ((v5) | 1))
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
