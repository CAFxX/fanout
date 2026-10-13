package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 0
    var v3 string = "hello"
    if false {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 int = (-(fi4))
            v5 = (fi4 - v5)
            _ = fi4
            _ = v5
        }
        if ("foo bar" > v3) {
            for fi6 := 0; fi6 < 4; fi6++ {
                fmt.Println(v2)
                v7 := GenS1{F0: (fi6 ^ 32768), F1: 65535, F2: (v2 * v2)}
                _ = fi6
                _ = v7
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                if (2 > 346) {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            v9 := GenS1{F0: 9223372036854775807, F1: 65535, F2: v2}
                            _ = v9
                        }
                        var v10 uint32 = 1
                        _ = v10
                    }
                    v3 = v3
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v3 = "foo bar"
                        switch fi8 {
                        case 127:
                            v12 := []int{fi8, fi8}
                            _ = v12
                        case -128:
                            v13 := GenS2{F0: 199, F1: 127, F2: 42}
                            _ = v13
                        case 9223372036854775807:
                            var v14 uint64 = 0
                            _ = v14
                        default:
                            fmt.Println(fi11, fi8, fi8)
                        }
                        fmt.Println(fi11, fi8, v3)
                        _ = fi11
                    }
                }
                if (false || true) {
                    var v15 int = 916
                    for fi16 := 0; fi16 < 16; fi16++ {
                        switch v3 {
                        case "a":
                            fmt.Println(fi8, v3)
                        case "":
                            if (1 < 4294967295) {
                                v3 = "hello"
                            }
                        case "\n\t":
                            var v17 int32 = 468
                            _ = v17
                        default:
                            v3 = "hello"
                        }
                        v18 := []int{}
                        _ = fi16
                        _ = v18
                    }
                    v19 := GenS1{F0: (-(v15)), F1: 1, F2: v2}
                    _ = v15
                    _ = v19
                } else {
                    switch 32768 {
                    case 50:
                        v2 = (v2 << ((450) & 63))
                    case 4294967295:
                        if true {
                            var v20 uint32 = 529
                            v21 := []int{128, 331, fi8}
                            _ = v20
                            _ = v21
                        }
                    }
                    v3 = "\n\t"
                }
                var v22 float32 = -2.25
                _ = fi8
                _ = v22
            }
            fmt.Println(v3)
        }
    } else {
        if true {
            v23 := GenS2{F0: 2, F1: -34, F2: -32768}
            _ = v23
        }
        delete(v1, ((v3 + "") + "a"))
        switch 127 {
        case 919:
            v3 = (("hello" + "") + (v3 + v3))
        }
    }
    v24 := GenS3{F0: 9223372036854775807}
    var v25 uint64 = 2
    v26 := GenS4{F0: 1}
    var v27 uint32 = 2
    var v28 float64 = 1e10
    v2 = (v2 - v2)
    v29 := []int{38, -9223372036854775808, -186}
    v2 = (v2 / ((v2) | 1))
    v25 -= v25
    var v30 uint64 = v25
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 int32 = v2
        if true {
            v33 := v1[v3]
            _ = v33
        }
        v3 = (((v3 + ("\n\t" + "\n\t")) + v3) + v3)
        _ = fi31
        _ = v32
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        if true {
            v25 = 18446744073709551615
            var v35 int8 = -1
            _ = v35
        }
        v30 += v25
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
