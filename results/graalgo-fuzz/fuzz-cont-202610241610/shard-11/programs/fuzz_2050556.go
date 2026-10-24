package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int32 = -2147483648
    if false {
        if true {
            if (false || true) {
                v3 = v3
                v4 := v1[2147483648]
                var v5 int = 456
                _ = v4
                _ = v5
            } else {
                v3 = -2147483648
                var v6 uint32 = 2
                var v7 int64 = 1
                _ = v6
                _ = v7
            }
        } else {
            if (true && ((v3 ^ v3) < (^(v3)))) {
                v8 := []int{-2147483648, 1}
                v8 = append(v8, 1)
                _ = v8
            }
            v9 := []int{-435, -844, 4294967295}
            v3 += (v3 + v3)
            _ = v9
        }
        var v10 string = (((("a" + "hello") + "\n\t") + "hello") + "")
        _ = v10
    } else {
        v11 := GenS1{F0: 277}
        _ = v11
    }
    if false {
        v12 := GenS2{F0: "a"}
        _ = len(v2)
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: 1}
        v3 -= (v3 - (v3 ^ 430))
        switch fi13 {
        case -2147483648:
            fmt.Println(fi13)
        default:
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    fmt.Println(v3, fi13)
                    v16 := []int{fi15}
                    _ = fi15
                    _ = v16
                }
                var v17 float64 = 3.14159
                _ = v17
            }
        }
        _ = fi13
        _ = v14
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int32 = (v3 + -2147483648)
        _ = fi18
        _ = v19
    }
    fmt.Println(v3)
    if true {
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                v3 = (v3 >> ((4294967295) & 63))
                v22 := []int{}
                v3 = v3
                _ = fi21
                _ = v22
            }
            v23 := []int{2147483647, fi20, fi20}
            var v24 float32 = (float32(float64((float32(3.14159) * 1.5))) / 1e10)
            _ = fi20
            _ = v23
            _ = v24
        }
        fmt.Println(v3)
    }
    fmt.Println(v3)
    var v25 int8 = 1
    v3 = -926
    var v26 int64 = int64(v25)
    switch -2147483648 {
    case -42:
        v2[2147483647] = (1 != 606)
    case -951:
        v27 := GenS2{F0: "\n\t"}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
