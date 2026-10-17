package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 uint) bool {
    p0 = (p0 + p0)
    var v7 uint64 = 1
    v8 := GenS1{F0: 255, F1: 42}
    v9 := GenS1{F0: 1, F1: 865}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 uint16) int32 {
    v10 := GenS2{F0: 100, F1: 1}
    v11 := GenS2{F0: 109, F1: 668}
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint = 791
        var v14 int16 = 32767
        if false {
            v15 := GenS2{F0: 127, F1: v13}
            v16 := GenS3{F0: 65, F1: uint16(v13), F2: false}
            var v17 int64 = -783
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v11
    return -509
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint8 = 0
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v6 := make(map[string]int)
    v18 := GenS3{F0: 58, F1: 0, F2: (true || false)}
    if (((!(!((24 >= 1) && false))) || (767 != 1)) || true) {
        var v19 int8 = -80
        _ = v19
    } else {
        v3 = (v3 * v3)
        var v20 string = ""
        v21 := GenS3{F0: 255, F1: 0, F2: true}
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 42, F1: 1}
    if false {
        v3 = (v3 | (v3 | 0))
    }
    if (true && false) {
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := []int{-1, fi23}
            _ = fi23
            _ = v24
        }
    } else {
        var v25 int = -565
        switch v25 {
        case 575:
            var v26 int64 = -9223372036854775808
            _ = v26
        case 15:
            var v27 float32 = 1e10
            _ = v27
        case -9223372036854775808:
            fmt.Println(v3)
        default:
            v28 := []int{v25, 127, v25}
            _ = v28
        }
        if true {
            var v29 uint32 = 2147483648
            _ = v29
        }
        _ = v25
    }
    v3 = v3
    if true {
        for fi30 := 0; fi30 < 3; fi30++ {
            v3 -= (v3 - v3)
            v3 = v3
            _ = fi30
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            switch "a" {
            case "":
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 bool = false
                    var v34 uint32 = 0
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            case "a":
                var v35 int32 = 0
                _ = v35
            }
            _ = fi31
        }
        v3 = v3
    }
    var v36 uint16 = 32767
    v37 := GenS4{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
