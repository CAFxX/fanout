package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint8 }
func genFn0(p0 bool, p1 int64) uint16 {
    v6 := []int{127}
    var v7 int = 941
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 670
}
func genFn1() uint8 {
    var v8 uint64 = 711
    var v9 uint32 = 4294967295
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: (v8 % ((v8) | 1)), F1: (v8 + (v8 / ((v8) | 1))), F2: (v9 ^ 456)}
            if false {
                v13 := []int{(fi11 << ((1) & 63)), (fi11 & (fi11 + fi10))}
                v8 = v8
                var v14 int = (fi11 >> ((971) & 63))
                _ = v13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        var v15 uint32 = 2147483647
        _ = fi10
        _ = v15
    }
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    v3 := make(map[uint32]int)
    v4 := []uint32{0, 159, 190, 0}
    v5 := make(map[int]bool)
    fmt.Println(v2)
    v2 = true
    fmt.Println(v2)
    v2 = (!true)
    if (2 != 4294967295) {
        var v16 string = "\n\t"
        v4 = append(v4, 817)
        _ = v16
    }
    var v17 bool = true
    switch "hello" {
    case "":
        var v18 uint16 = 786
        _ = v18
    case "foo bar":
        if ((!(-161 <= 852)) || true) {
            var v19 int32 = 0
            for fi20 := 0; fi20 < 3; fi20++ {
                v17 = v17
                _ = fi20
            }
            _ = v19
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                for fi22 := 0; fi22 < 12; fi22++ {
                    if v2 {
                        var v23 uint16 = 65535
                        v24 := GenS1{F0: 0, F1: 611, F2: 458}
                        v23 -= (v23 * 255)
                        if (((("hello" + "a") + "") <= "hello") && v17) {
                            fmt.Println(v2)
                            v25 := GenS2{F0: fi22, F1: 1}
                            if true {
                                fmt.Println(fi21, v17)
                            } else {
                                var v26 uint64 = 303
                                if v17 {
                                    v26 = (v26 / (((v26 & 0)) | 1))
                                } else {
                                    v27 := GenS1{F0: v26, F1: 1, F2: 968}
                                    _ = v27
                                }
                                var v28 int64 = -1
                                _ = v26
                                _ = v28
                            }
                            _ = v25
                        } else {
                            var v29 bool = true
                            _ = v29
                        }
                        var v30 uint = 255
                        _ = v23
                        _ = v24
                        _ = v30
                    }
                    _ = fi22
                }
                _ = fi21
            }
        }
    }
    var v31 int64 = 42
    v2 = v17
    v31 = 1
    v32 := GenS3{F0: 127, F1: 2, F2: genFn1()}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
