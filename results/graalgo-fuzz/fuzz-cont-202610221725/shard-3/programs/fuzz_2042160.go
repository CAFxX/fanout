package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool }
func genFn0(p0 string, p1 int16, p2 bool) int8 {
    switch -992 {
    case 4294967295:
        p1 += p1
    default:
        if p2 {
            var v6 bool = (753 != -1)
            v7 := GenS1{F0: -2147483648, F1: 4294967295, F2: 2147483647}
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                for fi9 := 0; fi9 < 7; fi9++ {
                    p0 = ("\n\t" + p0)
                    if (p1 != (p1 % (((p1 ^ p1)) | 1))) {
                        v10 := GenS1{F0: 2147483647, F1: 2147483647, F2: 2147483647}
                        if true {
                            if p2 {
                                fmt.Println(p2)
                            } else {
                                var v11 uint16 = 1
                                v11 = (v11 / (((v11 - v11)) | 1))
                                v12 := []int{fi9, fi9}
                                _ = v11
                                _ = v12
                            }
                            v13 := []int{}
                            _ = v13
                        } else {
                            v14 := []int{(^(fi8)), -1}
                            _ = v14
                        }
                        if (p1 != p1) {
                            var v15 int16 = 524
                            _ = v15
                        } else {
                            if (("hello" + "a") < (p0 + p0)) {
                                p1 -= (p1 ^ (p1 * (^(p1))))
                            }
                            if (1 <= 715) {
                                p2 = true
                            } else {
                                p0 = "hello"
                            }
                            var v16 bool = p2
                            p2 = v16
                            _ = v16
                        }
                        _ = v10
                    } else {
                        var v17 int32 = 816
                        var v18 float64 = 1.5
                        var v19 int32 = 2147483647
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi9
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 float64 = (1e10 * float64(1.5))
                    if (true || true) {
                        var v22 int = -2147483648
                        _ = v22
                    }
                    switch (p0 + (p0 + (p0 + ((p0 + p0) + "\n\t")))) {
                    case "\n\t":
                        var v23 int16 = 32767
                        _ = v23
                    default:
                        p1 += p1
                    }
                    v24 := GenS1{F0: 42, F1: 2147483647, F2: 166}
                    var v25 int16 = (p1 << ((681) & 63))
                    p1 += (p1 / (((^(v25))) | 1))
                    _ = fi20
                    _ = v21
                    _ = v24
                    _ = v25
                }
                _ = fi8
            }
            var v26 string = p0
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 uint = 1
                if true {
                    v29 := GenS1{F0: 2147483647, F1: 2147483647, F2: 391}
                    _ = v29
                }
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    }
    v30 := []int{}
    v31 := GenS2{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v31
    return 0
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 uint64 = 623
    v3 := []int{}
    v4 := make(map[uint32]string)
    var v5 uint = 4294967295
    var v32 int16 = 1
    v33 := GenS1{F0: 0, F1: 4294967295, F2: 42}
    fmt.Println(v2, v32, v32)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (1 > 0) {
        var v34 uint64 = 2
        _ = v34
    }
    if len(v3) > 0 {
        v3[0] = 32768
    }
    v35 := GenS2{F0: false}
    switch v1 {
    case "hello":
        var v36 uint8 = 0
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
