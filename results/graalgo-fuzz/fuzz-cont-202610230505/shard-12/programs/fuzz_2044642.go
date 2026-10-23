package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 string) int8 {
    v7 := GenS1{F0: 1}
    v8 := []int{412, -698, 769}
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 1
    v3 := []uint32{76, 2147483647, 0}
    v4 := make(map[string]int)
    var v5 int8 = -23
    v6 := []int{}
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            if (true || true) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint16 = 823
                    if (0.0 == -2.25) {
                        if false {
                            if ((532 != (^(v11))) && false) {
                                v12 := GenS1{F0: v2}
                                _ = v12
                            } else {
                                var v13 int16 = -1
                                delete(v1, -491)
                                var v14 int64 = 935
                                _ = v13
                                _ = v14
                            }
                            if ((^(v11)) > (v11 << ((184) & 63))) {
                                v11 = v11
                                v15 := GenS1{F0: 42}
                                _ = v15
                            }
                        } else {
                            var v16 int8 = genFn0("foo bar")
                            _ = v16
                        }
                    } else {
                        var v17 uint32 = 704
                        var v18 int16 = 1000
                        _ = v17
                        _ = v18
                    }
                    _ = fi10
                    _ = v11
                }
                v19 := GenS2{F0: 0, F1: 32767, F2: "foo bar"}
                _ = v19
            }
            _ = fi9
        }
    } else {
        if true {
            var v20 int = 2147483647
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v20
        } else {
            var v21 int16 = 489
            v21 = int16(v2)
            _ = v21
        }
    }
    var v22 uint16 = 255
    fmt.Println(v22, v5, v5)
    v23 := GenS3{F0: -1, F1: 65535}
    fmt.Println(v22, v5)
    var v24 int16 = 42
    v25 := GenS2{F0: 1, F1: v22, F2: ""}
    v26 := GenS3{F0: (v24 | (v24 >> ((0) & 63))), F1: v22}
    v22 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
