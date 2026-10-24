package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }
type GenS3 struct { F0 bool }
func genFn0() int {
    switch 856 {
    case 32767:
        v7 := GenS1{F0: 2147483647, F1: 4294967295}
        _ = v7
    case 42:
        var v8 uint8 = 2
        _ = v8
    case 9223372036854775807:
        var v9 int32 = 42
        _ = v9
    default:
        switch -2147483648 {
        case 2147483647:
            var v10 string = "hello"
            _ = v10
        case 32768:
            if (!true) {
                var v11 int8 = 1
                var v12 float32 = 1e10
                v12 = v12
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := []int{}
                    var v15 bool = (151 > 42)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                if false {
                    v16 := GenS1{F0: 1, F1: 799}
                    var v17 uint64 = 18446744073709551615
                    _ = v16
                    _ = v17
                } else {
                    var v18 int = 32767
                    v18 += (v18 % ((424) | 1))
                    var v19 string = "foo bar"
                    _ = v18
                    _ = v19
                }
                var v20 uint = 1
                _ = v20
            }
        }
    }
    v21 := []int{-735, -2147483648, -9223372036854775808}
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint8 = 2
        v23 = v23
        _ = fi22
        _ = v23
    }
    _ = v21
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483648, 4294967295, 2147483648, 42}
    v4 := []int{}
    v5 := make(map[int]int)
    var v6 uint32 = 0
    fmt.Println(v6)
    switch 129 {
    case -282:
        v4 = append(v4, 0)
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: (v6 << ((960) & 63)), F1: v6}
        v26 := GenS1{F0: (v6 - v6), F1: v6}
        v27 := GenS2{F0: true, F1: 42, F2: 42}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: (2 < 268), F1: 9223372036854775807, F2: 104}
    if true {
        var v29 int8 = -1
        _ = v29
    } else {
        var v30 uint64 = 42
        _ = v30
    }
    if (float64(0.1) <= (0.1 + (((3.14159 - (float64(1.5) / float64((-2.25 * 1.5)))) - 3.14159) + -2.25))) {
        var v31 int32 = -2147483648
        v6 = v6
        _ = v31
    }
    v4 = append(v4, genFn0())
    var v32 int8 = -59
    var v33 int64 = 9223372036854775807
    v34 := GenS3{F0: (-230 < 4294967295)}
    v33 = v33
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 uint64 = 0
        for fi37 := 0; fi37 < 16; fi37++ {
            switch (-(fi35)) {
            case 9223372036854775807:
                fmt.Println(v36, v33)
            default:
                v32 += (v32 - v32)
            }
            v32 = v32
            v3 = append(v3, (^(v6)))
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
