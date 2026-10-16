package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 uint8, p2 uint8) int8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(p1)
        v7 := []int{274, fi6, 32768, (fi6 ^ 32767)}
        _ = fi6
        _ = v7
    }
    p1 = 1
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    var v1 uint8 = 217
    v2 := make(map[uint32]string)
    var v3 uint16 = 1
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    var v8 int8 = -1
    var v9 uint8 = v1
    v10 := []int{9223372036854775807}
    var v11 uint32 = 2
    switch "a" {
    case "\n\t":
        var v12 uint64 = 363
        _ = v12
    case "a":
        var v13 float64 = 0.0
        _ = v13
    case "hello":
        if false {
            if ((int32(v3) >= int32(v8)) || (!false)) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS1{F0: -1, F1: v11, F2: fi14}
                    v9 -= v9
                    v9 = (^(v9))
                    _ = fi14
                    _ = v15
                }
                v16 := GenS2{F0: 1, F1: 1}
                _ = v16
            } else {
                for fi17 := 0; fi17 < 4; fi17++ {
                    fmt.Println(v1)
                    _ = fi17
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 uint32 = v11
                    _ = fi18
                    _ = v19
                }
            }
            fmt.Println(v8)
        }
    }
    v8 = -36
    var v20 uint = uint(v3)
    var v21 uint = v20
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint64 = 42
        v24 := v5[("" + ("" + ""))]
        v25 := GenS3{F0: (fi22 >> (((v20 / ((42) | 1))) & 63))}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 bool = false
    if (("" + "\n\t") < ("\n\t" + "\n\t")) {
        v27 := GenS2{F0: -1, F1: 596}
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := GenS3{F0: (fi28 ^ fi28)}
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        var v30 float64 = ((0.0 / (3.14159 - (0.0 + float64(0.1)))) / (1.5 - float64(-2.25)))
        _ = v30
    }
    var v31 int = 117
    if ((-(v31)) == (v31 * v31)) {
        v10 = append(v10, -566)
        if (!(-32768 > 42)) {
            v26 = (!(v26 && v26))
            var v32 uint8 = 85
            _ = len(v10)
            _ = v32
        }
    } else {
        switch 1 {
        case 42:
            var v33 float64 = 0.1
            _ = v33
        case 9223372036854775807:
            switch (v31 / ((42) | 1)) {
            case 2147483647:
                v20 -= v21
            }
        }
        switch (v31 + v31) {
        case 42:
            var v34 uint8 = v1
            _ = v34
        case 989:
            v31 = v31
        case 127:
            v20 += (v20 + 4294967295)
        }
    }
    fmt.Println(v31)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
