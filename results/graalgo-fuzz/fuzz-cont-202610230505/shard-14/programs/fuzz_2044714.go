package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
func genFn0() int32 {
    switch -1 {
    case 128:
        for fi5 := 0; fi5 < 17; fi5++ {
            fmt.Println(fi5)
            fmt.Println(fi5)
            _ = fi5
        }
    default:
        var v6 uint64 = 42
        _ = v6
    }
    if true {
        var v7 int64 = 1
        for fi8 := 0; fi8 < 8; fi8++ {
            v7 = (v7 | (v7 ^ (v7 / (((v7 * (v7 % ((v7) | 1)))) | 1))))
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 uint32 = 2147483648
                switch fi9 {
                case -893:
                    v11 := GenS1{F0: -32768}
                    _ = v11
                case 0:
                    var v12 uint16 = 434
                    switch (fi8 | (fi8 % ((fi9) | 1))) {
                    case 32767:
                        v13 := GenS1{F0: -251}
                        _ = v13
                    case -398:
                        v10 -= 4294967295
                    case 32768:
                        v7 = -947
                    }
                    switch (fi9 & (fi8 & (fi8 << ((1) & 63)))) {
                    case 58:
                        var v14 float64 = (3.14159 / float64((3.14159 - 1.5)))
                        _ = v14
                    default:
                        var v15 float64 = -2.25
                        _ = v15
                    }
                    _ = v12
                default:
                    var v16 uint8 = 110
                    _ = v16
                }
                v17 := []int{4294967295, -1}
                fmt.Println(v10, v10)
                v18 := GenS1{F0: 42}
                if true {
                    fmt.Println(fi9)
                }
                _ = fi9
                _ = v10
                _ = v17
                _ = v18
            }
            fmt.Println(v7, fi8)
            _ = fi8
        }
        _ = v7
    } else {
        var v19 uint8 = 141
        var v20 int32 = 0
        v21 := []int{32768, -347, -143, -702}
        _ = v19
        _ = v20
        _ = v21
    }
    return 1
}
func genFn1(p0 bool) int8 {
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            p0 = false
            switch (("a" + "\n\t") + "a") {
            case "\n\t":
                p0 = (p0 && false)
            }
            _ = fi23
        }
        v24 := []int{(fi22 + fi22), fi22}
        fmt.Println(fi22, fi22)
        _ = fi22
        _ = v24
    }
    var v25 uint32 = 208
    fmt.Println(p0)
    _ = p0
    _ = v25
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 997
    var v3 int8 = 1
    var v4 int16 = 32767
    v2 += v2
    switch "a" {
    case "\n\t":
        v26 := GenS2{F0: false, F1: int32(v4), F2: 18446744073709551615}
        _ = v26
    case "":
        v27 := GenS1{F0: (v4 ^ v4)}
        _ = v27
    }
    v28 := GenS3{F0: 65535}
    if false {
        v1[2147483647] = -2147483648
        var v29 float64 = (0.1 * 1e10)
        _ = v29
    } else {
        var v30 uint8 = 1
        v31 := GenS4{F0: false}
        var v32 bool = false
        _ = v30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        if false {
            fmt.Println(v3, v4)
            v34 := GenS5{F0: 305}
            _ = v34
        } else {
            var v35 uint16 = 65535
            if false {
                fmt.Println(v35)
            }
            v4 += v4
            _ = v35
        }
        var v36 float32 = 0.1
        _ = fi33
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        v4 = v4
        if (42 >= 2147483647) {
            _ = len(v1)
            var v38 uint16 = 72
            _ = v38
        }
        _ = fi37
    }
    var v39 uint16 = 0
    if (2 < 42) {
        fmt.Println(v39, v39)
        fmt.Println(v3)
    }
    v40 := GenS6{F0: -1}
    var v41 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
