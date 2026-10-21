package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
func genFn0(p0 string, p1 string, p2 int64) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := GenS1{F0: false, F1: fi4}
            var v7 uint8 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
        if true {
            v8 := GenS2{F0: 1, F1: -120}
            v9 := []int{(fi4 + fi4), fi4, (-(fi4))}
            if false {
                v10 := GenS1{F0: (4294967295 == 735), F1: fi4}
                p2 -= (p2 | -9223372036854775808)
                _ = v10
            }
            _ = v8
            _ = v9
        }
        _ = fi4
    }
    var v11 int = -9223372036854775808
    var v12 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return v12
}
func genFn1(p0 int8, p1 int8, p2 int) uint {
    var v13 bool = (false || (19 >= 127))
    var v14 int32 = 42
    v15 := GenS3{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 int8 = 0
    fmt.Println(v1, v1)
    var v16 int = 1
    v3 -= -92
    var v17 uint32 = 65535
    for fi18 := 0; fi18 < 9; fi18++ {
        if (uint8(v17) == 146) {
            switch -971 {
            case 9223372036854775807:
                v19 := GenS4{F0: (288 == 232)}
                _ = v19
            case 4294967295:
                var v20 int8 = -1
                _ = v20
            case 128:
                v21 := []int{v16, (fi18 + -7), -2147483648}
                _ = v21
            }
        }
        _ = fi18
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v16, v17)
    case "hello":
        for fi22 := 0; fi22 < 3; fi22++ {
            v1 = v1
            for fi23 := 0; fi23 < 18; fi23++ {
                if (v17 >= (^(v17))) {
                    v17 += 1
                    v24 := GenS5{F0: 345}
                    v25 := []int{v16, 32767, v16, (fi23 | (^(fi22)))}
                    _ = v24
                    _ = v25
                }
                _ = fi23
            }
            var v26 uint32 = (v17 + v17)
            _ = fi22
            _ = v26
        }
    }
    if (float64(0.1) >= 1.5) {
        for fi27 := 0; fi27 < 19; fi27++ {
            switch 4294967295 {
            case 4294967295:
                var v28 uint = 255
                _ = v28
            case 2147483647:
                v16 += -1
            }
            v16 = 42
            var v29 int8 = (v3 / (((v3 | v3)) | 1))
            _ = fi27
            _ = v29
        }
        var v30 uint16 = 161
        _ = v30
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 int8 = -128
            var v33 uint16 = 1
            var v34 uint = genFn1(0, v32, (v16 * fi31))
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
    }
    v16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
