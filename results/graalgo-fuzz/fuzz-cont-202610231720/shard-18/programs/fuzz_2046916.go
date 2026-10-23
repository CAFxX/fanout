package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
func genFn0(p0 int32) int {
    var v7 int = 32767
    v7 += 32767
    switch "a" {
    case "hello":
        if (true || false) {
            var v8 float32 = ((float32(3.14159) * ((3.14159 + 0.0) + (0.1 / -2.25))) / 1e10)
            var v9 uint = 42
            _ = v8
            _ = v9
        } else {
            switch "a" {
            case "hello":
                v10 := GenS1{F0: 2147483647}
                _ = v10
            }
        }
    default:
        v11 := []int{(v7 / ((960) | 1))}
        _ = v11
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    return (^(v7))
}
func genFn1() string {
    switch genFn0(-1) {
    case 4294967295:
        if false {
            var v12 uint32 = 1
            _ = v12
        }
    case 2147483647:
        v13 := []int{}
        _ = v13
    case -928:
        var v14 uint64 = 0
        _ = v14
    default:
        v15 := []int{921, 0, -2147483648, 4294967295}
        _ = v15
    }
    var v16 float64 = 0.1
    var v17 uint32 = 4294967295
    _ = v16
    _ = v17
    return "\n\t"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    v3 := []uint32{}
    var v4 int64 = 42
    v5 := make(map[int]int)
    var v6 uint8 = 128
    if false {
        var v18 int = 4294967295
        if (("" + "hello") > "\n\t") {
            var v19 string = "foo bar"
            var v20 uint32 = 65535
            _ = v19
            _ = v20
        }
        delete(v5, -9223372036854775808)
        _ = v18
    }
    if (2 <= 2) {
        if (1 == -507) {
            v21 := GenS1{F0: 205}
            if (false && true) {
                for fi22 := 0; fi22 < 20; fi22++ {
                    if true {
                        switch fi22 {
                        case 2147483647:
                            var v23 int = 801
                            _ = v23
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v4 = v4
                        var v25 bool = false
                        _ = fi24
                        _ = v25
                    }
                    _ = fi22
                }
            }
            _ = v21
        }
        v4 += (v4 & v4)
        fmt.Println(v4, v6)
    } else {
        if false {
            switch 663 {
            case 4294967295:
                v4 -= v4
            case 32767:
                v26 := GenS1{F0: 2147483647}
                _ = v26
            default:
                fmt.Println(v4)
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                if (fi27 <= -2147483648) {
                    if true {
                        var v28 int64 = v4
                        _ = v28
                    } else {
                        var v29 float32 = 0.0
                        v6 -= 38
                        v6 = v6
                        _ = v29
                    }
                    v30 := GenS2{F0: 0, F1: 42, F2: v4}
                    v6 = v6
                    _ = v30
                }
                var v31 int8 = 0
                _ = fi27
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 5; fi32++ {
                if true {
                    v6 += v6
                }
                fmt.Println(v4)
                v4 = 9223372036854775807
                _ = fi32
            }
        }
    }
    if true {
        v6 = (v6 ^ v6)
    }
    v4 = v4
    v6 = (v6 - v6)
    v33 := GenS1{F0: -709}
    fmt.Println(v6, v6)
    v4 = (v4 | v4)
    var v34 uint8 = v6
    if true {
        var v35 int64 = (v4 % (((v4 + 9223372036854775807)) | 1))
        _ = v35
    }
    var v36 uint8 = 0
    switch genFn0(0) {
    case -1:
        if false {
            v36 -= v34
            if ("a" == "hello") {
                fmt.Println(v36, v34, v6)
                var v37 int8 = -128
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                _ = v37
            }
            var v38 int = 4294967295
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 int8 = -128
            var v41 uint8 = v36
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
