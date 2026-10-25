package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0(p0 int32, p1 int16, p2 uint64) int64 {
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: p0, F1: (v6 - 1), F2: 1}
    p1 += 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 string) int16 {
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(p0, fi9)
            switch ("" + ("a" + p0)) {
            case "\n\t":
                v10 := []int{2147483647, fi9, -250}
                _ = v10
            case "foo bar":
                v11 := []int{fi8}
                _ = v11
            default:
                var v12 int8 = 1
                _ = v12
            }
            if (405 <= 4294967295) {
                fmt.Println(fi9, fi9, p0)
            } else {
                v13 := []int{fi8, fi9, fi8, (fi8 + -576)}
                var v14 uint = 198
                p0 = p0
                switch "\n\t" {
                case "a":
                    var v15 int16 = -1
                    _ = v15
                case "hello":
                    p0 = "a"
                default:
                    v14 = v14
                }
                var v16 uint16 = 537
                switch -9223372036854775808 {
                case 2147483647:
                    v17 := GenS1{F0: 2147483647, F1: 359, F2: -9223372036854775808}
                    _ = v17
                case 9223372036854775807:
                    var v18 bool = (42 < 2147483647)
                    _ = v18
                default:
                    fmt.Println(fi8, p0)
                    if true {
                        if false {
                            if true {
                                if false {
                                    v19 := GenS2{F0: 1, F1: fi9}
                                    _ = v19
                                }
                                p0 = "\n\t"
                                fmt.Println(fi8, fi8)
                            } else {
                                var v20 uint = (^(v14))
                                _ = v20
                            }
                            if false {
                                p0 = (("\n\t" + p0) + p0)
                                var v21 float32 = 3.14159
                                v22 := GenS3{F0: 0}
                                _ = v21
                                _ = v22
                            }
                        }
                        v16 = (v16 * 2)
                        var v23 int16 = 42
                        _ = v23
                    } else {
                        p0 = p0
                    }
                    var v24 string = p0
                    _ = v24
                }
                _ = v13
                _ = v14
                _ = v16
            }
            _ = fi9
        }
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    v25 := []int{}
    var v26 uint = 4294967295
    _ = p0
    _ = v25
    _ = v26
    return 0
}

func main() {
    var v1 string = "hello"
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 bool = (float64(1.5) <= (float64(1e10) * 0.0))
    v27 := GenS4{F0: v2, F1: 2}
    v5 = v5
    fmt.Println(v5)
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int32 = 179
        _ = fi28
        _ = v29
    }
    v5 = v5
    v2 = -1
    delete(v4, 0)
    v30 := GenS5{F0: 227, F1: -1, F2: 42}
    v1 = v1
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(v2, v1)
        switch fi31 {
        case 128:
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS2{F0: 1, F1: fi31}
                delete(v4, 2)
                var v34 int32 = -200
                fmt.Println(v5)
                if v5 {
                    var v35 string = v1
                    fmt.Println(v34, v1, v2)
                    if v5 {
                        var v36 int64 = (v2 ^ v2)
                        v5 = (428 == 42)
                        _ = v36
                    }
                    _ = v35
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
