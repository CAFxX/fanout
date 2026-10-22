package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 int8) int8 {
    var v4 int64 = 0
    v4 += v4
    v5 := []int{p0, 862}
    p0 = 945
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 int64, p1 int) int8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{p1, -9223372036854775808, -9223372036854775808, 4294967295}
        for fi8 := 0; fi8 < 4; fi8++ {
            if false {
                var v9 float32 = 1.5
                p0 = p0
                _ = len(v7)
                _ = v9
            } else {
                var v10 float64 = 1e10
                _ = v10
            }
            p0 = 42
            _ = fi8
        }
        _ = len(v7)
        _ = fi6
        _ = v7
    }
    p1 -= -333
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 54
    v2 := []uint32{156, 770, 4294967295, 261}
    v3 := make(map[string]bool)
    v11 := GenS1{F0: 0, F1: (v1 % (((v1 ^ v1)) | 1)), F2: -9223372036854775808}
    v12 := GenS2{F0: 4294967295}
    v13 := v3[("\n\t" + "a")]
    if (4294967295 >= 949) {
        var v14 int = 841
        _ = v14
    }
    _ = len(v3)
    switch 32768 {
    case 1:
        if v13 {
            var v15 int8 = 100
            _ = v15
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v16 := GenS1{F0: 2, F1: 1, F2: 42}
    v17 := GenS3{F0: 423}
    if (4294967295 != 4294967295) {
        fmt.Println(v13, v13)
    }
    v1 += v1
    fmt.Println(v13, v13)
    var v18 string = ((("hello" + "foo bar") + "") + ("hello" + ""))
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            v13 = v13
            v13 = (-1 < -519)
            var v21 uint8 = 127
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v1 = (v1 | v1)
        var v23 uint = 42
        if v13 {
            v24 := GenS1{F0: 4294967295, F1: v1, F2: 0}
            var v25 int64 = 42
            _ = v24
            _ = v25
        } else {
            var v26 uint8 = 0
            var v27 float64 = 0.0
            v28 := v3["hello"]
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi22
        _ = v23
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
