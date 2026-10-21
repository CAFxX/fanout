package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
func genFn0(p0 int32, p1 int64) int32 {
    v5 := GenS1{F0: 65535, F1: 0}
    fmt.Println(p0)
    fmt.Println(p0)
    var v6 int16 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 ^ p0)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{906, 65535, 186, 4294967295}
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v7 int = 1
    var v8 uint = 689
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            v9 := GenS1{F0: v8, F1: 42}
            _ = v9
        default:
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: 4294967295, F1: -939}
                fmt.Println(v8)
                _ = fi10
                _ = v11
            }
        }
    case "":
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(v7)
            var v13 float64 = (1e10 / 1e10)
            v13 -= v13
            _ = fi12
            _ = v13
        }
    case "hello":
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    v15 := GenS1{F0: (v8 ^ v8), F1: 9223372036854775807}
    _ = len(v3)
    var v16 uint8 = 75
    v17 := v3["a"]
    fmt.Println(v7, v17, v8)
    v2 = append(v2, (v17 - v17))
    fmt.Println(v8, v8, v17)
    var v18 uint8 = (v16 + v16)
    var v19 int64 = -1
    if (!(v19 <= (v19 ^ (-(v19))))) {
        for fi20 := 0; fi20 < 14; fi20++ {
            if (0 != -737) {
                var v21 bool = true
                v18 -= (v18 & (v18 + v18))
                var v22 uint32 = v17
                _ = v21
                _ = v22
            } else {
                switch v7 {
                case -9223372036854775808:
                    _ = len(v1)
                case -489:
                    var v23 uint16 = 2
                    _ = v23
                default:
                    _ = len(v2)
                }
                v7 += (^(v7))
            }
            if true {
                for fi24 := 0; fi24 < 16; fi24++ {
                    var v25 bool = false
                    if false {
                        v26 := v1[2147483648]
                        v27 := GenS1{F0: v8, F1: (v19 / (((v19 / (((v19 - -1)) | 1))) | 1))}
                        var v28 uint64 = 18446744073709551615
                        switch v7 {
                        case -541:
                            v29 := GenS1{F0: (^(v8)), F1: (v19 >> ((v8) & 63))}
                            _ = v29
                        case 860:
                            var v30 uint64 = v28
                            _ = v30
                        }
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v16, v8, v7)
                        var v31 uint32 = v17
                        _ = v31
                    }
                    _ = fi24
                    _ = v25
                }
                switch "" {
                case "foo bar":
                    fmt.Println(fi20, v8)
                case "":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 uint16 = 910
                        _ = len(v4)
                        _ = fi32
                        _ = v33
                    }
                default:
                    delete(v3, ((("foo bar" + "") + (("\n\t" + "") + "foo bar")) + ("" + "hello")))
                }
                var v34 float32 = 3.14159
                _ = v34
            }
            _ = fi20
        }
        _ = len(v2)
        v19 += v19
    } else {
        v35 := GenS2{F0: 42, F1: 2, F2: 796}
        var v36 int = v7
        fmt.Println(v36, v19)
        _ = v35
        _ = v36
    }
    switch (v7 / ((v7) | 1)) {
    case 155:
        v19 = -1
    case -9223372036854775808:
        var v37 int = (v7 & -659)
        _ = v37
    case -876:
        var v38 string = (("a" + "hello") + ("" + ""))
        _ = v38
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v39 uint16 = 1
        _ = v39
    case "\n\t":
        if true {
            for fi40 := 0; fi40 < 2; fi40++ {
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 int8 = 127
                    _ = fi41
                    _ = v42
                }
                _ = len(v2)
                fmt.Println(v8, v17)
                _ = fi40
            }
            fmt.Println(v16, v8)
        }
    case "a":
        v7 += (v7 / (((v7 >> ((uint(v17)) & 63))) | 1))
    }
    var v43 int64 = v19
    v16 -= v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v7
    _ = v8
}
