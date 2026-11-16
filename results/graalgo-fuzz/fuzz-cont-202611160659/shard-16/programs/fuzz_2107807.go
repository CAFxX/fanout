package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 string; F1 uint8; F2 string }
func genFn0(p0 int32) uint {
    fmt.Println(p0)
    v7 := GenS1{F0: 1, F1: 39}
    v8 := GenS1{F0: 1, F1: 52}
    if ("\n\t" > "hello") {
        v9 := []int{767}
        v10 := GenS1{F0: 2, F1: 42}
        _ = v10
        _ = v9
    } else {
        p0 -= (p0 * (p0 * (^(p0))))
        v11 := []int{32767, -9223372036854775808, -2147483648, 32767}
        _ = v11
    }
    v12 := []int{42, 158, -1}
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := []uint32{0, 2147483648}
    var v2 int = 836
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    var v5 int16 = 32767
    v6 := make(map[int]int)
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            if false {
                var v15 uint64 = 42
                v16 := GenS2{F0: (995 != -1), F1: (9223372036854775807 == 9223372036854775807)}
                if (65535 >= 2147483648) {
                    v17 := GenS1{F0: 255, F1: 881}
                    _ = v17
                } else {
                    var v18 uint = genFn0(42)
                    _ = v18
                }
                fmt.Println(fi14, v5)
                _ = v15
                _ = v16
            }
            _ = fi14
        }
        v3[19] = 1
        _ = fi13
    }
    var v19 int8 = 62
    v3[0] = (v2 + v2)
    v5 -= v5
    if true {
        v20 := GenS1{F0: 194, F1: 42}
        v2 -= (v2 >> ((785) & 63))
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS1{F0: 128, F1: 2}
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS1{F0: 255, F1: 4294967295}
            delete(v3, 2)
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 uint8 = 1
                var v27 float32 = 0.1
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v19)
        switch 1 {
        case -205:
            v2 -= v2
        }
    }
    var v28 bool = ((3.14159 > (0.1 * -2.25)) && false)
    var v29 uint16 = 2
    fmt.Println(v28)
    v30 := GenS3{F0: "foo bar", F1: 42, F2: "foo bar"}
    switch "hello" {
    case "hello":
        var v31 int32 = -578
        _ = v31
    }
    var v32 uint16 = 32767
    var v33 uint = 0
    if v28 {
        _ = len(v3)
    }
    v34 := GenS3{F0: "", F1: 138, F2: ("a" + ("\n\t" + ""))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
