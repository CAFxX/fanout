package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 377
    v3 := []uint32{}
    var v4 string = ""
    v5 := make(map[int]int)
    var v6 float64 = 1.5
    v7 := v1[-387]
    if (false && false) {
        if false {
            fmt.Println(v2, v6, v7)
            var v8 float32 = -2.25
            _ = v8
        } else {
            var v9 int32 = 2147483647
            if false {
                var v10 uint16 = 539
                v4 = "\n\t"
                var v11 int16 = 1000
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 20; fi12++ {
                    if (-898 >= v9) {
                        fmt.Println(v7, v4)
                        if (9223372036854775807 != -1) {
                            v4 = v4
                            for fi13 := 0; fi13 < 13; fi13++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (!true) {
                                    v14 := GenS1{F0: (100 < 127), F1: (^(v7)), F2: (((("foo bar" + v4) + v4) + "\n\t") + ((v4 + v4) + "hello"))}
                                    v15 := GenS1{F0: true, F1: 0, F2: ((("\n\t" + ((v4 + ("hello" + ("foo bar" + "\n\t"))) + v4)) + v4) + "a")}
                                    v16 := GenS1{F0: true, F1: v7, F2: ((v4 + "a") + v4)}
                                    fmt.Println(v6, v7)
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    v17 := GenS1{F0: (1e10 != -2.25), F1: v7, F2: v4}
                                    var v18 uint64 = 2
                                    _ = v17
                                    _ = v18
                                }
                                var v19 float32 = -2.25
                                _ = fi13
                                _ = v19
                            }
                            v7 += v7
                        } else {
                            v6 -= 0.1
                            v7 = v7
                            var v20 int64 = 0
                            _ = v20
                        }
                    }
                    var v21 int8 = 100
                    var v22 float32 = float32(((float64(3.14159) + 0.0) + 1e10))
                    _ = fi12
                    _ = v21
                    _ = v22
                }
            }
            _ = v9
        }
    } else {
        v6 = (((-2.25 / float64(((1e10 - -2.25) - (0.1 + 1e10)))) / -2.25) + (v6 + (0.1 * v6)))
        for fi23 := 0; fi23 < 6; fi23++ {
            delete(v1, fi23)
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 uint16 = uint16(v7)
                v26 := GenS2{F0: v4, F1: false}
                var v27 float64 = v6
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi28
            }
            _ = fi23
        }
    }
    var v29 int32 = 2147483647
    var v30 uint32 = 2147483647
    v31 := GenS2{F0: "a", F1: (true && false)}
    for fi32 := 0; fi32 < 1; fi32++ {
        v6 += float64((0.1 + 0.1))
        var v33 int = 797
        _ = fi32
        _ = v33
    }
    v5[4294967295] = -733
    v4 = (v4 + ((v4 + ("\n\t" + v4)) + (v4 + v4)))
    delete(v1, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
