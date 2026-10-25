package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint16 }
func genFn0(p0 int64, p1 uint16) uint32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 uint32 = 1
        if ((p0 ^ p0) == (^(p0))) {
            v6 := []int{}
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v7 := GenS1{F0: 848, F1: -760, F2: 677}
    p0 = (p0 / ((9223372036854775807) | 1))
    var v8 int64 = -1
    var v9 uint64 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 uint8 = 2
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 float32 = 1e10
        v12 := GenS2{F0: ""}
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 = 0
    var v13 int16 = 1
    if false {
        v3 += v3
        fmt.Println(v3, v3)
        var v14 uint32 = 0
        _ = v14
    } else {
        v15 := GenS3{F0: -2147483648}
        for fi16 := 0; fi16 < 15; fi16++ {
            if (-1 > 42) {
                for fi17 := 0; fi17 < 7; fi17++ {
                    fmt.Println(v3, v3, v3)
                    switch fi17 {
                    case 823:
                        if (1 == 953) {
                            _ = len(v2)
                            v18 := GenS4{F0: (^(v3)), F1: 1, F2: 760}
                            _ = v18
                        }
                    case -2147483648:
                        v19 := GenS2{F0: "a"}
                        _ = v19
                    case 9223372036854775807:
                        v3 -= 42
                    }
                    switch "a" {
                    case "\n\t":
                        var v20 int = (fi16 + (-(fi16)))
                        _ = v20
                    default:
                        if false {
                            _ = len(v1)
                            v3 += 192
                        }
                    }
                    _ = fi17
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS4{F0: (v3 + v3), F1: 2, F2: 212}
                    v13 = (v13 ^ 252)
                    _ = fi21
                    _ = v22
                }
                v3 = (^(v3))
            }
            if true {
                var v23 string = (((("hello" + (("hello" + "") + "foo bar")) + ("foo bar" + "a")) + "") + ("" + "a"))
                var v24 int64 = 1
                fmt.Println(v23, v13)
                _ = v23
                _ = v24
            }
            _ = fi16
        }
        _ = v15
    }
    v13 -= v13
    var v25 uint32 = 2147483647
    if (!false) {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        v25 = (^(v25))
    }
    if false {
        v2[32768] = "hello"
    }
    var v26 uint = 529
    v27 := v2[-58]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
