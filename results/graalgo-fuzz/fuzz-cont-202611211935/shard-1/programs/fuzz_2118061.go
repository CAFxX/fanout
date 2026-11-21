package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
func genFn0(p0 int8) int16 {
    var v6 string = "a"
    p0 = p0
    if (0 != 1) {
        v7 := []int{}
        fmt.Println(p0, p0)
        v7 = append(v7, 2147483647)
        _ = v7
    }
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 int64 = -745
    var v2 string = "\n\t"
    v3 := make(map[int]int)
    var v4 float32 = (3.14159 - 0.1)
    var v5 uint8 = 0
    v8 := []int{-9223372036854775808, 607}
    switch 277 {
    case -827:
        switch -490 {
        case -482:
            v5 = 183
        default:
            var v9 float32 = (v4 / v4)
            _ = v9
        }
    }
    switch (v2 + (("hello" + "a") + v2)) {
    case "":
        delete(v3, 90)
    case "\n\t":
        if (-2147483648 == -1) {
            switch 0 {
            case 4294967295:
                var v10 int16 = 32767
                _ = v10
            case 128:
                if false {
                    if (!false) {
                        v3[-758] = -2147483648
                        var v11 string = "foo bar"
                        v11 = v2
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v5 += (v5 >> ((1) & 63))
                        _ = fi12
                    }
                    switch -2147483648 {
                    case -94:
                        if false {
                            v13 := GenS1{F0: (^(v5))}
                            fmt.Println(v5)
                            _ = v13
                        } else {
                            v14 := GenS1{F0: 127}
                            v15 := v3[127]
                            _ = v14
                            _ = v15
                        }
                    case 0:
                        v16 := v3[9223372036854775807]
                        _ = v16
                    }
                }
            default:
                if true {
                    fmt.Println(v1)
                }
            }
            switch "\n\t" {
            case "hello":
                if (0 <= -748) {
                    var v17 int32 = 1
                    _ = v17
                }
            }
            var v18 uint64 = 80
            _ = v18
        }
    default:
        var v19 uint64 = 1
        _ = v19
    }
    var v20 uint = 91
    fmt.Println(v1)
    v21 := GenS1{F0: 1}
    var v22 int16 = 1000
    if true {
        switch v2 {
        case "":
            v4 -= float32(((float64(v4) - float64((v4 - 0.1))) * 3.14159))
        case "a":
            v1 = -1
        case "hello":
            v23 := GenS2{F0: v1}
            _ = v23
        default:
            fmt.Println(v20, v4)
        }
    } else {
        fmt.Println(v2, v4)
    }
    v22 = genFn0(0)
    if false {
        switch -9223372036854775808 {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 2; fi24++ {
                v1 += 1
                _ = fi24
            }
        case 889:
            fmt.Println(v2, v1, v20)
        default:
            switch "\n\t" {
            case "hello":
                var v25 int64 = 1
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            v5 = 0
            v27 := GenS3{F0: 127}
            switch fi26 {
            case 1:
                switch (fi26 ^ fi26) {
                case 42:
                    v28 := v3[fi26]
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v30 := GenS4{F0: 42, F1: 1, F2: 32767}
                        fmt.Println(fi29, v1, v1)
                        if true {
                            v31 := v3[-668]
                            fmt.Println(fi29)
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                }
            }
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
