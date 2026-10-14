package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int16; F2 uint32 }

func main() {
    var v1 float32 = float32(0.1)
    var v2 int64 = 0
    v3 := make(map[string]int)
    var v4 float32 = v1
    for fi5 := 0; fi5 < 1; fi5++ {
        if (!(127 <= 1)) {
            var v6 uint32 = 83
            v7 := GenS1{F0: 238, F1: v6}
            _ = v6
            _ = v7
        }
        if false {
            fmt.Println(v2, v4)
            v2 += (v2 | v2)
            v8 := GenS1{F0: 1, F1: 841}
            _ = v8
        } else {
            switch fi5 {
            case 706:
                var v9 string = "hello"
                _ = v9
            }
            if (32767 == -32768) {
                var v10 uint = 42
                _ = v10
            } else {
                v11 := []int{fi5, (fi5 / ((0) | 1))}
                for fi12 := 0; fi12 < 8; fi12++ {
                    _ = len(v3)
                    var v13 int16 = -1
                    _ = fi12
                    _ = v13
                }
                v1 = v4
                _ = v11
            }
        }
        v14 := GenS2{F0: 65535, F1: -128}
        _ = fi5
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            if (!(false && (0 == 42))) {
                if ((-32768 != 42) || (100 <= 127)) {
                    var v17 uint8 = 127
                    if false {
                        v1 -= (v1 + v1)
                        var v18 uint8 = 213
                        v19 := GenS3{F0: false}
                        _ = v18
                        _ = v19
                    }
                    var v20 string = "\n\t"
                    _ = v17
                    _ = v20
                }
                if true {
                    v21 := GenS4{F0: "foo bar"}
                    var v22 int = (fi15 % ((fi15) | 1))
                    if ((!(false && false)) && true) {
                        switch (fi15 & 4294967295) {
                        case 42:
                            var v23 uint16 = 2
                            _ = v23
                        case -201:
                            v4 -= 0.1
                        default:
                            v4 = v4
                        }
                        var v24 bool = (fi16 <= (fi15 / ((127) | 1)))
                        _ = v24
                    } else {
                        v25 := GenS2{F0: 1, F1: 100}
                        v26 := []int{v22}
                        if true {
                            _ = len(v3)
                        }
                        _ = v25
                        _ = v26
                    }
                    _ = v21
                    _ = v22
                }
                v3["\n\t"] = fi15
            }
            _ = fi16
        }
        v27 := GenS5{F0: 264, F1: 710, F2: 998}
        var v28 uint32 = 334
        _ = fi15
        _ = v27
        _ = v28
    }
    v1 = 0.1
    var v29 uint8 = 42
    var v30 uint64 = 42
    v31 := []int{825, int(v29)}
    fmt.Println(v29, v29)
    v4 -= (v1 / v1)
    if ((-467 != -201) && (!((!(!(!(-2147483648 == -2147483648)))) && true))) {
        var v32 uint32 = 760
        _ = v32
    } else {
        fmt.Println(v2, v30, v1)
        v4 = 1.5
    }
    var v33 int32 = 95
    v30 += v30
    v1 -= v4
    var v34 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
}
