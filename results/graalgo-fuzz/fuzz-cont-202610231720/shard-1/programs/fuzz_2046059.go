package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int64 }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483647
    v3 := []int{127, 127, 42}
    var v4 string = ((("hello" + ("\n\t" + "a")) + "foo bar") + ("hello" + "hello"))
    if true {
        fmt.Println(v2)
    }
    v5 := GenS1{F0: 1, F1: -1, F2: 0}
    if (42 <= 72) {
        _ = len(v1)
        v6 := GenS2{F0: -987, F1: 65535, F2: true}
        _ = v6
    }
    v7 := GenS3{F0: 1}
    switch 42 {
    case -2147483648:
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                _ = len(v1)
                if true {
                    if true {
                        _ = len(v3)
                        v9 := GenS4{F0: 40}
                        _ = v9
                    } else {
                        var v10 int64 = 0
                        for fi11 := 0; fi11 < 20; fi11++ {
                            var v12 uint64 = 255
                            v13 := GenS5{F0: v2}
                            var v14 uint16 = 1
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = v10
                    }
                    v2 -= (v2 << ((20) & 63))
                    v3 = append(v3, fi8)
                }
                var v15 uint64 = 810
                _ = fi8
                _ = v15
            }
            _ = len(v3)
        } else {
            switch -859 {
            case -738:
                switch ("" + "a") {
                case "\n\t":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        var v17 float32 = float32((0.1 + 0.0))
                        _ = fi16
                        _ = v17
                    }
                default:
                    var v18 int32 = 438
                    _ = v18
                }
            }
            if true {
                if true {
                    v4 = v4
                    var v19 int64 = 465
                    _ = v19
                }
            } else {
                var v20 uint32 = v2
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := GenS1{F0: -381, F1: -96, F2: 9223372036854775807}
                    v23 := GenS6{F0: 4294967295, F1: -32768, F2: (822 <= 32767)}
                    v4 = "a"
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v2 += (v20 & (^(v20)))
                _ = v20
            }
            v4 = ("hello" + (v4 + v4))
        }
    case -320:
        if ((589 >= 342) || true) {
            v2 += (v2 | (v2 ^ 2))
            if true {
                if (1 == 737) {
                    v2 -= (v2 & (v2 | (^(v2))))
                    var v24 float64 = -2.25
                    var v25 uint64 = 2
                    _ = v24
                    _ = v25
                }
                if true {
                    var v26 bool = true
                    var v27 int8 = -128
                    v26 = v26
                    _ = v26
                    _ = v27
                }
                if len(v3) > 0 {
                    v3[0] = 9223372036854775807
                }
            } else {
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 int8 = 127
                    _ = fi28
                    _ = v29
                }
                switch (((("\n\t" + ("a" + ((("\n\t" + v4) + "foo bar") + v4))) + "hello") + "hello") + "a") {
                case "\n\t":
                    v2 = (^(v2))
                case "foo bar":
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v4 = "\n\t"
                        var v31 int = int(v2)
                        _ = fi30
                        _ = v31
                    }
                }
            }
        } else {
            var v32 int32 = -2147483648
            v2 += (^(v2))
            _ = v32
        }
    }
    v2 = v2
    v1[-766] = v2
    var v33 int32 = -426
    v2 -= v2
    fmt.Println(v2, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v7
}
