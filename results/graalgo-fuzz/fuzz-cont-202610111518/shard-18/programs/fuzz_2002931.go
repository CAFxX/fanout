package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0() uint8 {
    v5 := []int{-808}
    v6 := GenS1{F0: -855}
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := []string{("" + "")}
    v2 := make(map[string]int)
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    var v7 bool = false
    fmt.Println(v7, v7)
    v8 := v4[4294967295]
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(fi10, v9, fi10)
        fmt.Println(v9)
        if (!v7) {
            if (true || v7) {
                if (-128 >= 13) {
                    var v11 int16 = v3
                    if (!v7) {
                        fmt.Println(v9)
                        for fi12 := 0; fi12 < 12; fi12++ {
                            if ((1.5 - 1e10) > 1e10) {
                                if (v3 <= (v11 % ((v11) | 1))) {
                                    v7 = true
                                    fmt.Println(v8, v8, fi10)
                                } else {
                                    v13 := GenS2{F0: -1, F1: 753}
                                    var v14 bool = v7
                                    v9 = "hello"
                                    _ = v13
                                    _ = v14
                                }
                                var v15 int = fi12
                                _ = v15
                            } else {
                                var v16 uint64 = 847
                                _ = v16
                            }
                            v17 := GenS2{F0: 1, F1: v8}
                            _ = fi12
                            _ = v17
                        }
                    } else {
                        v2[v9] = (fi10 % ((-84) | 1))
                        var v18 int32 = -611
                        fmt.Println(v11, v8)
                        _ = v18
                    }
                    _ = v11
                } else {
                    v19 := GenS1{F0: 0}
                    _ = v19
                }
                fmt.Println(fi10)
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := v4[(v8 >> ((2) & 63))]
                    _ = fi20
                    _ = v21
                }
            }
        }
        _ = fi10
    }
    v22 := GenS1{F0: -158}
    v23 := GenS2{F0: 1, F1: v8}
    fmt.Println(v8)
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS3{F0: (fi25 & fi25), F1: 42}
            v7 = v7
            v27 := GenS3{F0: 32767, F1: int8(v3)}
            v1 = append(v1, v9)
            var v28 uint64 = 0
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        fmt.Println(v9, v8, v3)
        _ = fi24
    }
    v29 := GenS1{F0: 653}
    var v30 uint64 = 289
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
