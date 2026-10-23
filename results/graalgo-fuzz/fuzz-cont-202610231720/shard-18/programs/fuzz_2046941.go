package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint; F2 bool }
func genFn0() string {
    if (1000 != 42) {
        v4 := GenS1{F0: 0, F1: 1}
        var v5 int = 32768
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        if (!false) {
            v7 := GenS1{F0: 255, F1: -1}
            _ = v7
        }
        switch fi6 {
        case 878:
            v8 := GenS1{F0: 4294967295, F1: 300}
            _ = v8
        case -289:
            for fi9 := 0; fi9 < 6; fi9++ {
                if false {
                    var v10 uint16 = 2
                    switch ((("hello" + "a") + "\n\t") + "") {
                    case "a":
                        v11 := GenS1{F0: 65535, F1: -2147483648}
                        _ = v11
                    }
                    _ = v10
                }
                fmt.Println(fi6, fi6)
                if (18446744073709551615 >= 42) {
                    var v12 int8 = 42
                    var v13 float32 = 1.5
                    _ = v12
                    _ = v13
                } else {
                    if false {
                        var v14 string = ""
                        v14 = "\n\t"
                        v14 = (v14 + v14)
                        var v15 int32 = -2147483648
                        _ = v14
                        _ = v15
                    }
                }
                _ = fi9
            }
        case 128:
            var v16 int16 = 1000
            _ = v16
        }
        _ = fi6
    }
    switch 128 {
    case -725:
        v17 := GenS2{F0: 2}
        _ = v17
    case -9223372036854775808:
        v18 := GenS3{F0: "hello", F1: 65535, F2: true}
        _ = v18
    case -1:
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := []int{127}
            var v21 uint32 = 4294967295
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS3{F0: "\n\t", F1: 176, F2: true}
    var v23 string = ""
    _ = v22
    _ = v23
    return v23
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = -110
    var v3 uint32 = 92
    v3 = 1
    v1 += v1
    if true {
        for fi24 := 0; fi24 < 4; fi24++ {
            fmt.Println(fi24, v1)
            switch fi24 {
            case 1:
                v1 += (v1 % (((^(v1))) | 1))
            }
            v25 := []int{(fi24 / ((fi24) | 1)), -2147483648, (fi24 % ((-899) | 1)), -991}
            _ = fi24
            _ = v25
        }
    }
    var v26 int = 9223372036854775807
    var v27 uint8 = 93
    var v28 uint8 = v27
    if (987 >= 9223372036854775807) {
        v29 := GenS1{F0: 42, F1: 418}
        for fi30 := 0; fi30 < 3; fi30++ {
            fmt.Println(v27, v3, v26)
            v28 += (v28 - (v27 | v28))
            _ = fi30
        }
        _ = v29
    }
    v31 := []int{v26, (v26 + 2147483647)}
    var v32 int64 = -413
    for fi33 := 0; fi33 < 12; fi33++ {
        v1 += v1
        var v34 uint16 = v1
        _ = fi33
        _ = v34
    }
    if (!true) {
        _ = len(v31)
        for fi35 := 0; fi35 < 12; fi35++ {
            v31 = append(v31, -9223372036854775808)
            v1 -= (v1 >> ((35) & 63))
            v26 = (fi35 * fi35)
            _ = fi35
        }
        v27 -= 131
    } else {
        if (true || true) {
            var v36 uint64 = 18446744073709551615
            _ = v36
        }
    }
    var v37 uint16 = 255
    if true {
        v38 := GenS2{F0: 4294967295}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v37
}
