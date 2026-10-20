package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 uint64) int64 {
    fmt.Println(p1, p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6)
        var v7 uint32 = 2147483647
        for fi8 := 0; fi8 < 19; fi8++ {
            v7 = (v7 / ((v7) | 1))
            p1 = p1
            v9 := GenS1{F0: 1, F1: 0}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if false {
        for fi10 := 0; fi10 < 8; fi10++ {
            p1 += (p1 % ((p1) | 1))
            _ = fi10
        }
        switch 127 {
        case 0:
            fmt.Println(p1, p0)
        case -9223372036854775808:
            v11 := []int{9223372036854775807}
            _ = v11
        case 2147483647:
            v12 := []int{9223372036854775807, 32768, 128, -2147483648}
            _ = v12
        }
    } else {
        var v13 float32 = 1e10
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 bool = false
            if (0 <= 42) {
                switch 9223372036854775807 {
                case -582:
                    var v16 float64 = float64(0.1)
                    _ = v16
                case 712:
                    v17 := []int{}
                    _ = v17
                case 0:
                    p1 = (p1 >> ((0) & 63))
                }
            } else {
                var v18 uint64 = p1
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v19 float32 = 0.0
    _ = p0
    _ = p1
    _ = v19
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = ((-2.25 / (1.5 * ((3.14159 - (1.5 - (1e10 - 3.14159))) + ((1.5 + ((3.14159 - float32((0.0 - -2.25))) / 3.14159)) * 1.5)))) + 3.14159)
    var v3 uint64 = 42
    v4 := make(map[int]string)
    var v5 int32 = 892
    if (-74 == 42) {
        switch 945 {
        case 4294967295:
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := []int{0}
                _ = fi20
                _ = v21
            }
        case 128:
            v22 := []int{}
            _ = v22
        }
        fmt.Println(v3)
        v5 += -865
    } else {
        var v23 int32 = v5
        var v24 int64 = 9223372036854775807
        v25 := []int{32767, 32768, 32767, 32768}
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := []int{156}
    v2 -= v2
    fmt.Println(v5, v3)
    fmt.Println(v3)
    _ = len(v1)
    v3 = v3
    v27 := GenS2{F0: v3}
    v2 += 0.0
    switch "" {
    case "":
        for fi28 := 0; fi28 < 10; fi28++ {
            switch ("a" + "hello") {
            case "hello":
                for fi29 := 0; fi29 < 18; fi29++ {
                    v3 -= 42
                    _ = fi29
                }
            }
            var v30 uint8 = 255
            _ = fi28
            _ = v30
        }
    }
    if true {
        v2 += float32(-2.25)
        v5 = 1
    } else {
        var v31 uint64 = (v3 % ((449) | 1))
        v32 := GenS1{F0: (v5 << ((42) & 63)), F1: 1}
        var v33 int64 = -1
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
