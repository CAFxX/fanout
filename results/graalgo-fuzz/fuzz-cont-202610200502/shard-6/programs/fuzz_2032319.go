package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int32; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 1
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 int64 = v2
    v6 := []string{"foo bar"}
    v7 := GenS1{F0: 100}
    if true {
        v5 += v2
        v8 := GenS1{F0: 127}
        var v9 uint16 = 65535
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        if false {
            v11 := GenS2{F0: 21}
            switch 4294967295 {
            case 42:
                v2 += v5
            }
            _ = v11
        }
        delete(v3, fi10)
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint16 = 42
            var v14 int64 = (^(v5))
            v5 = (v14 * (^(v14)))
            v13 += v13
            var v15 float32 = 3.14159
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    v5 -= (v2 % (((v5 - v5)) | 1))
    var v16 uint16 = 2
    for fi17 := 0; fi17 < 7; fi17++ {
        _ = len(v6)
        _ = fi17
    }
    v18 := GenS2{F0: 4294967295}
    var v19 int16 = -187
    var v20 uint64 = 18446744073709551615
    var v21 bool = false
    var v22 uint16 = (v16 ^ 32767)
    for fi23 := 0; fi23 < 2; fi23++ {
        v16 += v16
        var v24 uint = 18446744073709551615
        var v25 uint8 = 1
        _ = fi23
        _ = v24
        _ = v25
    }
    if (v2 >= v2) {
        if ((128 == -109) || v21) {
            var v26 int16 = (v19 + v19)
            if v21 {
                var v27 uint16 = (v22 << ((255) & 63))
                if v21 {
                    if v21 {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v29 := GenS3{F0: "a", F1: -890}
                            for fi30 := 0; fi30 < 8; fi30++ {
                                v31 := GenS4{F0: 1, F1: "a", F2: (^(v20))}
                                if ((v16 | (v16 * (^(v27)))) == (v22 ^ v22)) {
                                    fmt.Println(v27, v21, v16)
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                _ = fi30
                                _ = v31
                            }
                            v26 = -1
                            _ = fi28
                            _ = v29
                        }
                    }
                } else {
                    _ = len(v1)
                    for fi32 := 0; fi32 < 7; fi32++ {
                        var v33 int32 = 2147483647
                        _ = fi32
                        _ = v33
                    }
                    if (v5 == v5) {
                        fmt.Println(v5, v16)
                    }
                }
                var v34 uint = 65535
                _ = v27
                _ = v34
            } else {
                v6 = append(v6, "foo bar")
                var v35 int32 = 0
                _ = v35
            }
            if len(v6) > 0 {
                v6[0] = "hello"
            }
            _ = v26
        }
        fmt.Println(v22, v20, v5)
    }
    fmt.Println(v20, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
