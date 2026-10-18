package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8) int {
    if (!true) {
        var v6 uint64 = 191
        fmt.Println(v6)
        if (1 >= 557) {
            var v7 int = 2147483647
            var v8 uint8 = 129
            p0 = (p0 | p0)
            _ = v7
            _ = v8
        }
        _ = v6
    }
    if false {
        var v9 uint64 = 2
        _ = v9
    }
    var v10 uint8 = 171
    v10 += v10
    _ = p0
    _ = v10
    return 9223372036854775807
}
func genFn1(p0 int16, p1 string) int {
    switch 127 {
    case 4294967295:
        p1 = (p1 + ((("foo bar" + "a") + "foo bar") + "foo bar"))
    default:
        p0 = p0
    }
    v11 := GenS1{F0: 2, F1: 42, F2: -128}
    if false {
        if true {
            var v12 int = -947
            switch (v12 ^ -475) {
            case 128:
                v13 := GenS2{F0: 2, F1: 407, F2: 0}
                _ = v13
            default:
                var v14 float64 = (1.5 / (3.14159 + 0.0))
                _ = v14
            }
            var v15 int32 = 0
            _ = v12
            _ = v15
        }
    } else {
        fmt.Println(p1, p1)
        p1 = (p1 + p1)
        if ((128 <= 56) || false) {
            if false {
                v16 := []int{}
                var v17 uint16 = 255
                _ = v16
                _ = v17
            }
            v18 := []int{}
            switch -2147483648 {
            case 127:
                for fi19 := 0; fi19 < 15; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 float32 = ((3.14159 - 0.0) + 0.0)
                        _ = fi20
                        _ = v21
                    }
                    p1 = ((p1 + p1) + "")
                    var v22 uint8 = 53
                    _ = fi19
                    _ = v22
                }
            }
            _ = v18
        } else {
            p0 = p0
            if (!false) {
                v23 := []int{128, 9223372036854775807, 967, genFn0(-89)}
                p0 = (p0 * 42)
                if (-1 <= 0) {
                    p0 = (p0 % (((p0 ^ p0)) | 1))
                    switch p1 {
                    case "hello":
                        if (true && (!false)) {
                            var v24 int = 403
                            var v25 int = (-(v24))
                            v24 += (v24 + (-(v24)))
                            _ = v24
                            _ = v25
                        }
                    case "\n\t":
                        var v26 int64 = -588
                        _ = v26
                    default:
                        v27 := GenS2{F0: 18446744073709551615, F1: 128, F2: 0}
                        _ = v27
                    }
                    fmt.Println(p0, p0)
                } else {
                    var v28 string = "hello"
                    var v29 bool = (65535 <= 710)
                    _ = v28
                    _ = v29
                }
                _ = v23
            }
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return int(p0)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 0
    var v3 uint32 = 0
    var v4 uint64 = 413
    var v5 int = v2
    _ = len(v1)
    v30 := GenS2{F0: (^(v4)), F1: 991, F2: 1}
    v3 += 2
    for fi31 := 0; fi31 < 5; fi31++ {
        v3 = uint32(v4)
        v32 := GenS1{F0: 83, F1: 832, F2: int8(fi31)}
        var v33 int64 = 0
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := GenS3{F0: 1}
    v3 = v3
    v35 := GenS1{F0: 2, F1: int64(v4), F2: -128}
    v36 := []int{(v5 & (v5 + (v2 & (v2 + (-(v2))))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
