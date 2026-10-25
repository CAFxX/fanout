package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 uint16, p1 uint8, p2 int32) uint64 {
    v6 := GenS1{F0: 100}
    var v7 int32 = (^(p2))
    p2 += (v7 * v7)
    var v8 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 614
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    v4 := []uint32{2147483647, 233}
    v5 := make(map[string]string)
    var v9 float32 = 0.1
    switch (((("" + "\n\t") + ("foo bar" + "\n\t")) + "") + "foo bar") {
    case "foo bar":
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS1{F0: -1}
                if true {
                    if (227 <= 131) {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v1, v9, v1)
                        case "foo bar":
                            if (1 < 42) {
                                fmt.Println(v9, fi10, fi11)
                                v1 = (v1 / (((v1 * (v1 | v1))) | 1))
                                v1 += 779
                            }
                        case "\n\t":
                            fmt.Println(v1, fi10)
                            v13 := GenS2{F0: 18446744073709551615, F1: -2147483648}
                            v2 = (v2 % ((genFn0(42, 189, -530)) | 1))
                            _ = v13
                        }
                        v9 += v9
                    }
                    var v14 float32 = 0.1
                    if (255 < uint(v2)) {
                        var v15 int16 = (v1 % ((-675) | 1))
                        _ = len(v5)
                        _ = v15
                    }
                    _ = v14
                } else {
                    v4 = append(v4, 2147483648)
                    var v16 uint64 = (v2 << ((261) & 63))
                    _ = v16
                }
                if false {
                    var v17 int = (fi11 >> ((982) & 63))
                    _ = v17
                } else {
                    var v18 uint64 = v2
                    _ = v18
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    v5["hello"] = ("\n\t" + ("\n\t" + "\n\t"))
    for fi19 := 0; fi19 < 17; fi19++ {
        v9 = (v9 * v9)
        for fi20 := 0; fi20 < 5; fi20++ {
            v9 -= (float32(0.0) * (v9 / v9))
            _ = fi20
        }
        _ = fi19
    }
    switch 0 {
    case 9223372036854775807:
        v1 = v1
    default:
        if (v9 == 1e10) {
            if (100 <= 21) {
                switch 2147483647 {
                case 4294967295:
                    v3[1] = ((("\n\t" + "a") + "a") + "foo bar")
                }
            }
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 bool = ("hello" >= ("\n\t" + "hello"))
                v23 := GenS1{F0: -106}
                for fi24 := 0; fi24 < 15; fi24++ {
                    v2 = v2
                    _ = fi24
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            v1 -= 436
        }
    }
    var v25 int = -196
    var v26 uint16 = 65535
    v27 := GenS3{F0: (v25 | (-(v25))), F1: ((-2.25 + 1e10) == -2.25)}
    var v28 uint16 = (v26 | (^(v26)))
    if false {
        v2 = 18446744073709551615
        var v29 uint64 = v2
        _ = v29
    } else {
        for fi30 := 0; fi30 < 11; fi30++ {
            if true {
                v31 := GenS3{F0: fi30, F1: (!true)}
                v28 = v28
                _ = v31
            } else {
                _ = len(v3)
            }
            v32 := GenS3{F0: 951, F1: false}
            switch "foo bar" {
            case "foo bar":
                for fi33 := 0; fi33 < 7; fi33++ {
                    v9 = 1.5
                    v26 += v28
                    _ = fi33
                }
            default:
                v4 = append(v4, 2)
            }
            _ = fi30
            _ = v32
        }
        v28 = (^(v26))
    }
    var v34 int16 = v1
    fmt.Println(v9, v34)
    delete(v3, 2)
    var v35 int64 = 9223372036854775807
    v36 := v5[("a" + "\n\t")]
    v37 := GenS1{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
