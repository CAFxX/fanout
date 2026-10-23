package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 int8) uint16 {
    v6 := GenS1{F0: (^(p1))}
    v7 := GenS1{F0: (p1 ^ (p1 % (((p1 * p1)) | 1)))}
    var v8 int = int(p0)
    if false {
        if true {
            if (!true) {
                p0 += p0
            } else {
                p0 -= 42
                if false {
                    if (false && true) {
                        for fi9 := 0; fi9 < 4; fi9++ {
                            if true {
                                v10 := GenS1{F0: p1}
                                p1 += p1
                                _ = v10
                            }
                            _ = fi9
                        }
                        v11 := []int{1, 1, (-(v8)), (v8 << ((213) & 63))}
                        v12 := GenS1{F0: (p1 - (p1 / ((int8(v8)) | 1)))}
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS1{F0: (p1 / ((-1) | 1))}
                    v14 := GenS1{F0: (p1 + (p1 & 100))}
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: p1}
                    _ = v15
                }
            }
            if false {
                v16 := []int{1}
                if (false && false) {
                    if true {
                        var v17 int16 = 42
                        _ = v17
                    }
                    if len(v16) > 0 {
                        v16[0] = v8
                    }
                }
                _ = v16
            }
        } else {
            p0 = p0
        }
    }
    v18 := []int{(v8 << ((65535) & 63))}
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1() int32 {
    var v19 int32 = -870
    v20 := []int{}
    if true {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    v21 := GenS2{F0: 42, F1: 0, F2: 4294967295}
    _ = v19
    _ = v20
    _ = v21
    return (v19 % ((1) | 1))
}

func main() {
    var v1 bool = (!true)
    v2 := []string{"hello"}
    v3 := []string{"a", "foo bar", "\n\t"}
    var v4 int32 = 42
    var v5 float32 = 3.14159
    v22 := GenS3{F0: "hello"}
    v5 -= 3.14159
    var v23 int32 = v4
    fmt.Println(v4, v1, v5)
    if true {
        v5 = (1.5 - (v5 * v5))
    }
    v4 = v23
    v5 = (v5 * v5)
    fmt.Println(v1)
    v4 = v23
    fmt.Println(v4)
    var v24 bool = ((130 > 0) || v1)
    fmt.Println(v1, v23, v5)
    v23 = (-(v23))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
