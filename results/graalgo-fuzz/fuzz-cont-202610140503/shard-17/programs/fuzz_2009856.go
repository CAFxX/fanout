package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 string, p1 uint32) uint8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            p1 = p1
            p1 -= p1
        }
        fmt.Println(p1)
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: (p1 << ((609) & 63))}
            var v10 uint16 = 32767
            switch (fi8 >> ((4294967295) & 63)) {
            case -9223372036854775808:
                switch (p0 + "hello") {
                case "":
                    p0 = ""
                default:
                    v11 := []int{-439, 32768}
                    _ = v11
                }
            case 128:
                p0 = (p0 + p0)
            case 0:
                v12 := GenS1{F0: 2}
                _ = v12
            default:
                if (!false) {
                    v10 += (v10 % (((v10 * (^(v10)))) | 1))
                    var v13 uint8 = 56
                    _ = v13
                } else {
                    fmt.Println(fi8)
                    var v14 uint32 = p1
                    if true {
                        fmt.Println(v10, v14)
                        var v15 int16 = 1000
                        _ = v15
                    }
                    _ = v14
                }
                v16 := []int{}
                _ = v16
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v17 uint16 = 42
    switch ("" + "hello") {
    case "a":
        var v18 uint = 2
        _ = v18
    default:
        v19 := GenS1{F0: (^(p1))}
        _ = v19
    }
    var v20 int = 18
    _ = p0
    _ = p1
    _ = v17
    _ = v20
    return 2
}
func genFn1(p0 uint64) int64 {
    var v21 string = "\n\t"
    v21 = v21
    var v22 uint64 = (p0 % (((^(p0))) | 1))
    var v23 bool = true
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return 9223372036854775807
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    v5 := []int{4294967295, -43}
    v6 := make(map[int]uint32)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1)
    switch 563 {
    case -885:
        var v24 float32 = (3.14159 + 1.5)
        _ = v24
    case 2147483647:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 32767:
        var v25 string = "a"
        _ = v25
    default:
        if false {
            var v26 uint64 = 916
            v5 = append(v5, 1)
            _ = v26
        }
    }
    var v27 uint = 2
    fmt.Println(v1, v1)
    fmt.Println(v1, v27)
    fmt.Println(v1, v1)
    var v28 int8 = -12
    fmt.Println(v28, v27, v27)
    v27 = v27
    if false {
        v1 -= v1
    } else {
        if false {
            _ = len(v5)
            switch ("\n\t" + "hello") {
            case "a":
                v29 := GenS1{F0: 4294967295}
                _ = v29
            case "\n\t":
                _ = len(v3)
            default:
                fmt.Println(v1, v1)
            }
        } else {
            _ = len(v5)
        }
        if false {
            if (!false) {
                v27 += 1
            }
            fmt.Println(v28)
        }
    }
    v1 += (v1 + v1)
    fmt.Println(v1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
