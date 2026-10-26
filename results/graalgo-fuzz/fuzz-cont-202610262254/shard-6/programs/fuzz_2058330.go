package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 int64, p2 int8) uint8 {
    p1 -= (p1 % ((1) | 1))
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]int)
    var v3 int32 = 1
    var v4 string = "foo bar"
    v5 := []uint32{}
    var v6 uint8 = 73
    v4 = "hello"
    switch 32768 {
    case 32767:
        v7 := GenS1{F0: -128, F1: (^(v1)), F2: 0}
        _ = v7
    }
    fmt.Println(v1, v6)
    if (42 == 2) {
        switch (v4 + "hello") {
        case "":
            v4 = "hello"
        case "foo bar":
            var v8 int32 = 548
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS2{F0: -128, F1: -1}
                v2[(^(fi10))] = (fi10 - (fi10 / (((fi9 % ((fi10) | 1))) | 1)))
                if (!(42 < 529)) {
                    v4 = "\n\t"
                    var v12 int16 = 1000
                    if (false && (0 != 4294967295)) {
                        var v13 int32 = (^(v3))
                        var v14 int64 = 779
                        var v15 uint32 = 65535
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int8 = 100
                        fmt.Println(fi10)
                        v1 = v1
                        _ = v16
                    }
                    if true {
                        fmt.Println(v6, fi10)
                        var v17 float32 = 3.14159
                        _ = v17
                    } else {
                        if (!false) {
                            switch v4 {
                            case "a":
                                var v18 uint32 = 2
                                _ = v18
                            }
                            v6 = (v6 & 0)
                        } else {
                            v1 = uint64(v6)
                        }
                        v19 := v2[fi9]
                        _ = v19
                    }
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            var v20 int8 = 100
            _ = fi9
            _ = v20
        }
        if true {
            switch v4 {
            case "\n\t":
                fmt.Println(v1)
            case "hello":
                var v21 uint64 = v1
                _ = v21
            case "":
                if true {
                    var v22 int16 = 32767
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 uint64 = v1
                        var v25 bool = ((v24 | v1) > 4294967295)
                        var v26 uint8 = v6
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    if true {
                        _ = len(v2)
                    } else {
                        var v27 int16 = v22
                        _ = v27
                    }
                    _ = v22
                }
            }
            v4 = (v4 + v4)
        }
    }
    var v28 uint = 42
    for fi29 := 0; fi29 < 15; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            v1 -= v1
            if true {
                v31 := GenS3{F0: 1, F1: 288}
                _ = len(v2)
                _ = v31
            } else {
                v3 = v3
            }
            _ = fi30
        }
        delete(v2, (-(fi29)))
        _ = fi29
    }
    if false {
        v4 = ("\n\t" + ((v4 + v4) + v4))
        var v32 string = v4
        _ = v32
    }
    var v33 int8 = 0
    v1 -= (v1 & 4294967295)
    v6 += v6
    var v34 string = "a"
    v3 = v3
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 float32 = ((0.1 - ((0.1 + 3.14159) / (3.14159 * -2.25))) * (3.14159 + 3.14159))
        var v37 int = fi35
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
