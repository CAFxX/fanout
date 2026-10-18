package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 int32 = -47
    v5 := []int{}
    var v6 int16 = -32768
    var v7 int = 0
    if false {
        v8 := GenS1{F0: 1}
        fmt.Println(v6)
        _ = v8
    } else {
        v2[69] = -395
        if (!false) {
            fmt.Println(v7)
            var v9 int32 = v4
            _ = v9
        } else {
            var v10 string = (((("foo bar" + "foo bar") + "") + "\n\t") + "hello")
            switch (v7 | (-(v7))) {
            case 128:
                var v11 string = (v10 + v10)
                _ = v11
            case -2147483648:
                for fi12 := 0; fi12 < 10; fi12++ {
                    v10 = v10
                    _ = fi12
                }
            }
            if true {
                v1 -= v1
                var v13 uint64 = 94
                _ = v13
            }
            _ = v10
        }
        v14 := GenS2{F0: v7, F1: uint(v1), F2: 4294967295}
        _ = v14
    }
    v6 += -224
    v7 -= (^(v7))
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            _ = len(v5)
            _ = fi15
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int32 = v4
                if (!true) {
                    switch fi17 {
                    case 347:
                        var v19 float32 = 1e10
                        _ = v19
                    case 128:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -1:
                        if (42 != 42) {
                            v1 -= v1
                            v2[4294967295] = (v7 + (fi17 * (-(fi16))))
                            if true {
                                var v20 uint32 = 0
                                v21 := GenS1{F0: 2}
                                _ = v20
                                _ = v21
                            } else {
                                v6 += (^(v6))
                                delete(v2, 65535)
                                v6 = (v6 ^ 42)
                            }
                            v18 = (v18 - 1)
                            v22 := GenS2{F0: fi16, F1: 0, F2: 42}
                            _ = v22
                        }
                    }
                    v23 := GenS2{F0: v7, F1: 2, F2: 100}
                    var v24 string = "a"
                    _ = v23
                    _ = v24
                }
                var v25 int64 = 42
                _ = fi17
                _ = v18
                _ = v25
            }
            _ = fi16
        }
    } else {
        var v26 uint8 = 42
        v1 = (v1 % ((v1) | 1))
        var v27 uint = 2
        _ = v26
        _ = v27
    }
    var v28 bool = true
    v3[""] = 65535
    v28 = v28
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := GenS3{F0: (v7 / ((fi29) | 1))}
        v31 := GenS1{F0: 68}
        var v32 int32 = 42
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
