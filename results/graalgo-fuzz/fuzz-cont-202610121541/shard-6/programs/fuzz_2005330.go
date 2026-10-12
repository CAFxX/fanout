package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 bool }

func main() {
    var v1 uint8 = 1
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    var v5 float32 = (float32((3.14159 * 0.1)) + float32((-2.25 / 3.14159)))
    for fi6 := 0; fi6 < 6; fi6++ {
        delete(v3, 0)
        _ = fi6
    }
    delete(v3, 973)
    v5 += 0.1
    _ = len(v2)
    for fi7 := 0; fi7 < 19; fi7++ {
        if ("foo bar" < ("\n\t" + "foo bar")) {
            if (312 < 32767) {
                _ = len(v3)
                fmt.Println(v1)
            }
        }
        _ = fi7
    }
    fmt.Println(v4, v1)
    v1 = (v1 / (((v1 & (v1 / ((128) | 1)))) | 1))
    fmt.Println(v1, v4, v4)
    switch "\n\t" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    case "a":
        if (982 <= 669) {
            fmt.Println(v4, v4)
            var v8 int = 128
            _ = v8
        } else {
            v4 = (^(v4))
        }
    case "hello":
        var v9 int32 = int32(v1)
        _ = v9
    }
    var v10 uint32 = 2
    for fi11 := 0; fi11 < 19; fi11++ {
        if false {
            var v12 uint8 = v1
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS1{F0: 0, F1: 1, F2: (true && (!(v4 == int8(fi13))))}
                var v15 float32 = (v5 + -2.25)
                _ = fi13
                _ = v14
                _ = v15
            }
            v4 += v4
            _ = v12
        } else {
            v10 -= v10
            delete(v3, 4294967295)
            var v16 uint8 = (v1 + (v1 >> ((42) & 63)))
            _ = v16
        }
        if (1 <= v10) {
            if (true && (v5 > (v5 - (v5 - (v5 * float32(0.1)))))) {
                v17 := GenS2{F0: 25, F1: (^(v1))}
                v4 += (v4 % (((v4 + 1)) | 1))
                _ = v17
            } else {
                fmt.Println(fi11, fi11)
            }
            v18 := GenS1{F0: 967, F1: 0, F2: false}
            _ = v18
        } else {
            v2 = append(v2, ("hello" + ""))
        }
        _ = fi11
    }
    if (!true) {
        if false {
            if false {
                for fi19 := 0; fi19 < 5; fi19++ {
                    v10 += 43
                    fmt.Println(fi19, fi19, fi19)
                    _ = fi19
                }
                switch "foo bar" {
                case "":
                    v10 = v10
                case "foo bar":
                    if (v10 <= (^(v10))) {
                        switch "\n\t" {
                        case "a":
                            var v20 int8 = 100
                            _ = v20
                        case "":
                            for fi21 := 0; fi21 < 2; fi21++ {
                                var v22 int = fi21
                                _ = fi21
                                _ = v22
                            }
                        default:
                            v23 := GenS3{F0: "a", F1: false}
                            _ = v23
                        }
                    } else {
                        v10 = v10
                    }
                default:
                    v24 := GenS2{F0: 255, F1: 69}
                    _ = v24
                }
                _ = len(v2)
            }
            delete(v3, 32767)
        }
        var v25 int64 = 42
        _ = v25
    } else {
        var v26 uint = 18446744073709551615
        _ = v26
    }
    var v27 uint16 = 1
    v4 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
