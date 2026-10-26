package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint8, p1 uint16, p2 uint32) uint8 {
    switch "\n\t" {
    case "\n\t":
        var v7 uint8 = (p0 >> ((0) & 63))
        _ = v7
    case "foo bar":
        var v8 uint = 255
        _ = v8
    case "a":
        v9 := GenS1{F0: 205}
        _ = v9
    default:
        v10 := []int{32768, 32768, 32768, -63}
        _ = v10
    }
    v11 := GenS1{F0: (^(p0))}
    v12 := []int{128, 892, 4294967295}
    var v13 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return uint8(p2)
}
func genFn1(p0 int, p1 int16, p2 bool) int64 {
    for fi14 := 0; fi14 < 15; fi14++ {
        if false {
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := []int{p0, (^(fi15)), 32767, p0}
                switch "\n\t" {
                case "a":
                    var v17 int64 = -9223372036854775808
                    _ = v17
                case "hello":
                    var v18 int = (fi14 - p0)
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
        p0 = (p0 ^ (p0 + 705))
        _ = fi14
    }
    var v19 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 9223372036854775807
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]bool)
    var v3 uint = 18446744073709551615
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]string)
    v20 := GenS1{F0: genFn0(42, 1, 65535)}
    v21 := GenS1{F0: 14}
    var v22 int16 = -1
    v23 := GenS2{F0: 128, F1: (v22 | v22), F2: (v3 + (v3 / (((v3 + (^(v3)))) | 1)))}
    if (1 <= -9223372036854775808) {
        if false {
            v24 := v2[891]
            var v25 bool = v4
            _ = v24
            _ = v25
        } else {
            var v26 int16 = -416
            _ = v26
        }
        switch -2147483648 {
        case 1:
            switch "foo bar" {
            case "\n\t":
                for fi27 := 0; fi27 < 5; fi27++ {
                    switch "hello" {
                    case "a":
                        v28 := []int{(fi27 % ((int(v3)) | 1)), (fi27 | (^(fi27))), 4294967295, fi27}
                        _ = v28
                    default:
                        v2[898] = (false || ((v3 / (((^(v3))) | 1)) <= v3))
                    }
                    for fi29 := 0; fi29 < 17; fi29++ {
                        v30 := GenS1{F0: genFn0(0, 255, 1)}
                        v4 = true
                        var v31 float64 = 1.5
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS2{F0: 1, F1: v22, F2: (v3 & v3)}
                    _ = fi27
                    _ = v32
                }
            case "hello":
                v33 := []int{}
                _ = v33
            case "":
                var v34 uint32 = 203
                _ = v34
            default:
                var v35 float64 = 0.0
                _ = v35
            }
        case 9223372036854775807:
            var v36 bool = (!false)
            _ = v36
        default:
            var v37 float32 = -2.25
            _ = v37
        }
    }
    if v4 {
        fmt.Println(v3)
    }
    switch ("" + "\n\t") {
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        if (!(4294967295 < 699)) {
            for fi38 := 0; fi38 < 2; fi38++ {
                var v39 uint = 0
                _ = fi38
                _ = v39
            }
            var v40 int16 = int16(v1)
            _ = v40
        } else {
            v22 += (v22 + (^(v22)))
            v41 := []int{}
            _ = v41
        }
    }
    v22 += v22
    v42 := GenS3{F0: (v3 + (v3 + 0)), F1: genFn1(-474, -32768, (1 != 2))}
    for fi43 := 0; fi43 < 1; fi43++ {
        if (0.1 >= 3.14159) {
            fmt.Println(v3)
            fmt.Println(v1)
        }
        switch (^(fi43)) {
        case 32768:
            if false {
                v1 += v1
            }
        }
        _ = fi43
    }
    var v44 uint64 = 563
    if v4 {
        delete(v6, 65535)
    } else {
        fmt.Println(v44)
        v45 := GenS4{F0: (v3 ^ (v3 - v3)), F1: 32767}
        var v46 uint = v3
        _ = v45
        _ = v46
    }
    v47 := v6[2]
    var v48 int32 = 205
    fmt.Println(v3, v44)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
