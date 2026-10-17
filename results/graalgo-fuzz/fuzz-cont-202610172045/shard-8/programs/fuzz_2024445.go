package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64) int64 {
    v5 := []int{128, -9223372036854775808}
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 += 0
        _ = fi6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: -128, F1: 2147483648}
        var v9 uint8 = 37
        for fi10 := 0; fi10 < 8; fi10++ {
            if len(v5) > 0 {
                v5[0] = 32768
            }
            v11 := GenS1{F0: -1, F1: 2147483648}
            var v12 uint16 = 255
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 uint32, p1 int8) int {
    var v13 int = 189
    p0 = p0
    v14 := []int{2147483647}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (v13 - v13)
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = false
    var v3 int32 = -638
    var v4 int64 = 907
    v15 := []int{}
    v16 := GenS1{F0: 42, F1: 4294967295}
    var v17 string = (((v1 + v1) + v1) + v1)
    switch "\n\t" {
    case "hello":
        for fi18 := 0; fi18 < 4; fi18++ {
            switch (fi18 * 1) {
            case 9223372036854775807:
                if ((v4 * (v4 + v4)) != v4) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v15 = append(v15, genFn1(4294967295, -1))
                        _ = len(v15)
                        v2 = v2
                        _ = fi19
                    }
                    var v20 int32 = v3
                    _ = v20
                } else {
                    v21 := GenS1{F0: -116, F1: 65535}
                    switch v17 {
                    case "foo bar":
                        v17 = ("\n\t" + v1)
                    }
                    var v22 float32 = 3.14159
                    _ = v21
                    _ = v22
                }
            case -1:
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 float32 = 3.14159
                    _ = fi23
                    _ = v24
                }
            case 127:
                fmt.Println(v1, v4)
            }
            v1 = v1
            switch fi18 {
            case -680:
                if (v2 || v2) {
                    v25 := GenS2{F0: v4, F1: genFn0(v4), F2: v2}
                    var v26 uint64 = 42
                    _ = v25
                    _ = v26
                } else {
                    v1 = v1
                    if (0 == 2) {
                        var v27 string = v1
                        _ = v27
                    }
                }
            }
            _ = fi18
        }
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 uint8 = 127
            var v30 uint32 = 330
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    if (!(!v2)) {
        v2 = v2
        switch 42 {
        case 4294967295:
            var v31 uint64 = 348
            _ = v31
        default:
            v32 := GenS2{F0: v4, F1: v4, F2: (4294967295 <= 42)}
            _ = v32
        }
    }
    v33 := GenS3{F0: 0}
    if ((v17 + "hello") <= v1) {
        var v34 float64 = ((float64(3.14159) * 0.0) + 1.5)
        _ = v34
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
