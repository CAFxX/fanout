package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int, p1 uint16, p2 bool) int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "" {
        case "":
            for fi8 := 0; fi8 < 18; fi8++ {
                p0 += fi8
                var v9 int16 = -1
                v10 := []int{fi8, p0}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        v11 := []int{(fi7 - p0), (fi7 << ((18446744073709551615) & 63))}
        switch p0 {
        case 4294967295:
            v12 := GenS1{F0: (p2 || (!p2)), F1: 300}
            _ = v12
        case 2147483647:
            if true {
                var v13 uint8 = 42
                if len(v11) > 0 {
                    v11[0] = (-(fi7))
                }
                var v14 int32 = -256
                _ = v13
                _ = v14
            } else {
                fmt.Println(p1, p1)
            }
        }
        _ = fi7
        _ = v11
    }
    var v15 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 498
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint = 42
    v3 := make(map[string]int)
    var v4 uint8 = 42
    var v5 float64 = 3.14159
    var v6 string = "hello"
    fmt.Println(v5, v4)
    v16 := GenS1{F0: (81 >= -128), F1: 987}
    v2 += (^(v2))
    v5 = 0.0
    v2 += 0
    switch (v1 & (v1 / (((v1 + 4294967295)) | 1))) {
    case 483:
        v1 = (v1 | 32767)
    case 9223372036854775807:
        var v17 float64 = 0.1
        _ = v17
    case 790:
        var v18 uint8 = v4
        _ = v18
    default:
        if false {
            v2 = (v2 - v2)
        } else {
            if (true && (true && (0 != 32767))) {
                v3[""] = -9223372036854775808
                for fi19 := 0; fi19 < 14; fi19++ {
                    v20 := GenS2{F0: 255}
                    var v21 bool = true
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                if false {
                    var v22 int = (v1 | v1)
                    v5 = v5
                    _ = v22
                } else {
                    if (100 <= -128) {
                        v1 -= v1
                        var v23 int16 = 674
                        _ = v23
                    }
                }
            }
            v24 := GenS1{F0: true, F1: 1}
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 uint8 = 42
                for fi27 := 0; fi27 < 1; fi27++ {
                    _ = len(v3)
                    v26 += v4
                    _ = fi27
                }
                fmt.Println(fi25)
                _ = fi25
                _ = v26
            }
            _ = v24
        }
    }
    fmt.Println(v6, v2)
    v6 = "a"
    if false {
        v28 := GenS3{F0: 18446744073709551615, F1: -1, F2: 0}
        _ = v28
    } else {
        var v29 int8 = 127
        _ = v29
    }
    var v30 string = "foo bar"
    if true {
        _ = len(v3)
        fmt.Println(v6)
    }
    v31 := []int{-96, (v1 % ((32768) | 1))}
    v32 := GenS4{F0: v2}
    fmt.Println(v4, v2, v5)
    for fi33 := 0; fi33 < 9; fi33++ {
        v31 = append(v31, fi33)
        v1 = (-(fi33))
        v30 = (v30 + (v30 + v30))
        _ = fi33
    }
    var v34 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
