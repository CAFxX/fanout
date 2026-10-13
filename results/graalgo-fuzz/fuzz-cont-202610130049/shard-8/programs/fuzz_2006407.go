package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 int32 }
func genFn0(p0 bool) int8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            switch "a" {
            case "":
                var v6 uint8 = 255
                _ = v6
            case "hello":
                switch fi4 {
                case 32768:
                    p0 = (!p0)
                }
            default:
                p0 = (4294967295 == 2)
            }
            var v7 uint = 826
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 string = ""
        _ = fi8
        _ = v9
    }
    if (p0 && true) {
        var v10 int16 = 1000
        v10 = v10
        v10 -= (v10 & 1)
        _ = v10
    } else {
        fmt.Println(p0)
        if (42 == -9223372036854775808) {
            var v11 int8 = -128
            _ = v11
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint16 = 1
                v14 := GenS1{F0: -322}
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: 1}
            v16 := []int{32768}
            _ = v15
            _ = v16
        }
        p0 = (2147483647 != 128)
    }
    v17 := GenS2{F0: 42, F1: false, F2: 976}
    _ = p0
    _ = v17
    return 42
}
func genFn1() uint16 {
    var v18 uint = 971
    v18 -= (v18 - (v18 & v18))
    v18 -= v18
    v19 := []int{1, 622, -306, -73}
    _ = v18
    _ = v19
    return 42
}

func main() {
    v1 := []uint32{4294967295, 636}
    var v2 int8 = 127
    v3 := make(map[int]int)
    var v20 string = "foo bar"
    v21 := GenS2{F0: 311, F1: (-395 >= 1000), F2: 2}
    for fi22 := 0; fi22 < 11; fi22++ {
        v2 = v2
        v23 := v3[-391]
        v2 = 1
        _ = fi22
        _ = v23
    }
    if false {
        v20 = ""
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 uint = 4294967295
            var v26 float32 = 0.0
            v2 += 42
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 string = v20
        _ = v27
    }
    v2 += -7
    var v28 int8 = -1
    var v29 uint16 = 65535
    v30 := GenS3{F0: 2, F1: 1, F2: -1}
    v2 = (v2 | (v2 * v2))
    v1 = append(v1, 42)
    for fi31 := 0; fi31 < 3; fi31++ {
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint64 = 887
            v28 -= (v28 | 0)
            var v34 uint16 = 42
            var v35 bool = false
            fmt.Println(v2)
            _ = len(v1)
            var v36 float64 = (-2.25 * 0.0)
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
            _ = v36
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v20, v2)
        _ = fi31
    }
    var v37 float64 = 1e10
    var v38 uint = 0
    _ = len(v1)
    fmt.Println(v28, v20, v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
}
