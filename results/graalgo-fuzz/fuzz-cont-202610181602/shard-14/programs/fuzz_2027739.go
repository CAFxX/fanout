package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 10; fi4++ {
        switch "" {
        case "foo bar":
            fmt.Println(fi4, fi4)
        case "\n\t":
            switch (fi4 - fi4) {
            case 4294967295:
                fmt.Println(fi4, v1)
            case 42:
                if ((176 == -1) && true) {
                    var v5 float64 = (0.1 - 3.14159)
                    for fi6 := 0; fi6 < 15; fi6++ {
                        v5 = (1e10 - (v5 * v5))
                        if (-812 < 1000) {
                            v1 = v1
                            v7 := GenS1{F0: v1, F1: 78, F2: (fi4 - 949)}
                            v8 := []int{(fi6 / ((-939) | 1)), (fi4 * (fi6 % (((fi6 ^ 32768)) | 1))), int(v1)}
                            _ = v7
                            _ = v8
                        }
                        _ = fi6
                    }
                    var v9 uint64 = 1
                    _ = v5
                    _ = v9
                }
            default:
                var v10 int8 = 1
                _ = v10
            }
        case "a":
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS2{F0: 1, F1: 42}
                var v13 int = fi11
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi4
    }
    v1 = v1
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int64 = -489
        v16 := GenS3{F0: (^(v1)), F1: 32767}
        _ = len(v2)
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int8 = 100
    var v18 uint64 = 42
    fmt.Println(v17, v18)
    v19 := GenS1{F0: v1, F1: 12, F2: 46}
    var v20 int16 = -1
    v21 := []int{9223372036854775807, 4294967295}
    v18 = uint64(v1)
    fmt.Println(v17, v20, v20)
    switch 428 {
    case -2147483648:
        var v22 bool = false
        _ = v22
    case 709:
        switch "foo bar" {
        case "hello":
            fmt.Println(v20, v18)
        case "\n\t":
            v17 -= (v17 + 107)
        default:
            _ = len(v21)
        }
    default:
        var v23 uint16 = 32767
        _ = v23
    }
    var v24 float32 = 3.14159
    if true {
        if true {
            v25 := GenS3{F0: 32767, F1: v20}
            _ = v25
        }
        switch -6 {
        case 1:
            for fi26 := 0; fi26 < 13; fi26++ {
                if true {
                    fmt.Println(v17, v18, v1)
                }
                var v27 uint = 208
                var v28 uint32 = 981
                _ = fi26
                _ = v27
                _ = v28
            }
        case -2147483648:
            v21 = append(v21, 77)
        }
        delete(v3, ((((("" + "foo bar") + (("" + "\n\t") + "a")) + (("a" + "foo bar") + "\n\t")) + "foo bar") + "a"))
    } else {
        v29 := GenS2{F0: (-(v17)), F1: 229}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
}
