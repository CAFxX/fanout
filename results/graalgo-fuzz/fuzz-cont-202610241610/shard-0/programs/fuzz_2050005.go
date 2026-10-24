package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v4 bool = v2
    var v5 int32 = 1
    if (801 < 128) {
        v6 := GenS1{F0: (v5 ^ v5), F1: (^(v5)), F2: -850}
        v5 += v5
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(v2, v5)
            switch "foo bar" {
            case "\n\t":
                v1[""] = (v2 && v2)
            case "a":
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 string = ("a" + "")
                    if (1 > -32768) {
                        switch (^(fi8)) {
                        case 1:
                            v2 = v4
                        case -9223372036854775808:
                            v5 = (v5 + (v5 << ((503) & 63)))
                            v10 := GenS2{F0: -407, F1: 42, F2: 2}
                            _ = v10
                        }
                        fmt.Println(v4)
                        fmt.Println(fi8)
                    } else {
                        v11 := GenS2{F0: -1, F1: 65535, F2: 46}
                        var v12 int8 = 81
                        _ = v11
                        _ = v12
                    }
                    v13 := v1[v9]
                    _ = fi8
                    _ = v13
                    _ = v9
                }
            }
            v14 := []int{}
            _ = fi7
            _ = v14
        }
        _ = v6
    } else {
        v15 := GenS1{F0: -1, F1: v5, F2: -685}
        fmt.Println(v2, v2, v5)
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        if ((v5 << ((833) & 63)) <= 0) {
            v17 := GenS1{F0: v5, F1: 2147483647, F2: -1}
            v18 := []int{fi16, fi16, 0, (fi16 | 4294967295)}
            v5 = 408
            _ = v17
            _ = v18
        } else {
            var v19 float32 = float32((0.0 + (1e10 + 1e10)))
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := []int{(fi16 % ((fi20) | 1)), (fi20 % ((-261) | 1)), fi16}
                v5 = v5
                fmt.Println(v2, v2, fi20)
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                if true {
                    switch "" {
                    case "foo bar":
                        var v23 float32 = v19
                        _ = v23
                    case "hello":
                        v5 = v5
                    case "a":
                        v19 += v19
                    default:
                        var v24 int16 = 1000
                        v25 := v3[493]
                        _ = v24
                        _ = v25
                    }
                } else {
                    v2 = ((v2 && v4) && ((v5 <= (v5 ^ (-(v5)))) && ((v4 || v2) || v4)))
                }
                var v26 int32 = v5
                v19 += (v19 / 1e10)
                _ = fi22
                _ = v26
            }
            _ = v19
        }
        _ = len(v3)
        _ = fi16
    }
    v27 := []int{4294967295}
    var v28 string = "hello"
    v27 = append(v27, 32767)
    if v2 {
        _ = len(v3)
        v4 = (128 <= 2)
    } else {
        v29 := GenS2{F0: -1, F1: 255, F2: 2}
        _ = v29
    }
    var v30 int32 = (v5 | v5)
    var v31 int8 = -101
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
