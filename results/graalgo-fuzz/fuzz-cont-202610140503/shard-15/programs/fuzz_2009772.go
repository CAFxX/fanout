package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 uint8) int {
    fmt.Println(p1)
    v6 := GenS1{F0: 128}
    switch -2147483648 {
    case 42:
        switch "foo bar" {
        case "foo bar":
            var v7 float64 = (0.0 + (3.14159 + (0.0 * 0.0)))
            _ = v7
        case "hello":
            v8 := []int{9223372036854775807, 127}
            _ = v8
        case "":
            v9 := []int{42, -2147483648}
            _ = v9
        }
    case 4294967295:
        var v10 uint16 = uint16(p0)
        _ = v10
    }
    if true {
        p1 += (^(p1))
        var v11 uint64 = 42
        _ = v11
    }
    var v12 int16 = -1
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return -2147483648
}
func genFn1(p0 int, p1 uint, p2 uint32) int32 {
    v13 := []int{}
    p1 = p1
    if len(v13) > 0 {
        v13[0] = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    var v1 int64 = 387
    var v2 int8 = 57
    var v3 uint = 537
    var v4 int32 = -133
    var v5 float32 = (1.5 - (0.1 + float32(-2.25)))
    v14 := GenS1{F0: int(v1)}
    if ((v5 + (1e10 + v5)) <= v5) {
        if (358 != 2) {
            if false {
                fmt.Println(v4, v4, v2)
                v15 := GenS1{F0: 9223372036854775807}
                _ = v15
            }
        } else {
            v16 := []int{}
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint32 = 714
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        fmt.Println(v5)
    } else {
        v3 -= (v3 % (((v3 | 65535)) | 1))
    }
    v19 := []int{615, -1, -2147483648}
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int16 = 42
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v23 := GenS1{F0: -358}
    v2 += (v2 ^ (^(v2)))
    var v24 int32 = v4
    v5 = v5
    v25 := GenS1{F0: 42}
    var v26 float64 = -2.25
    v27 := GenS1{F0: 374}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
