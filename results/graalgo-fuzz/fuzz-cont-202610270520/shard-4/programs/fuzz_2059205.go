package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 uint32 }
func genFn0() uint16 {
    var v5 uint = 817
    switch "\n\t" {
    case "hello":
        v6 := []int{}
        _ = v6
    default:
        var v7 int32 = 1
        _ = v7
    }
    var v8 uint8 = 2
    for fi9 := 0; fi9 < 9; fi9++ {
        v8 = v8
        v10 := GenS1{F0: 255, F1: ""}
        fmt.Println(v8, v8, v8)
        _ = fi9
        _ = v10
    }
    _ = v5
    _ = v8
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 65535
    v3 := []int{733, 2147483647, -9223372036854775808}
    var v4 float32 = 0.1
    v11 := GenS1{F0: genFn0(), F1: "a"}
    var v12 bool = true
    if false {
        var v13 uint8 = 128
        _ = v13
    }
    v2 = 323
    var v14 float64 = -2.25
    delete(v1, 11)
    var v15 uint64 = 600
    fmt.Println(v12, v14, v4)
    v15 = 42
    var v16 int16 = 1
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint = 18446744073709551615
            if v12 {
                fmt.Println(fi17, v14)
                v20 := GenS2{F0: 2, F1: (^(fi18)), F2: "a"}
                v21 := GenS3{F0: 681, F1: 1}
                _ = v20
                _ = v21
            } else {
                var v22 int = (fi18 % ((fi17) | 1))
                v2 -= v2
                fmt.Println(v16, v12, fi18)
                switch v22 {
                case -9223372036854775808:
                    var v23 uint64 = v15
                    _ = v23
                case -672:
                    if v12 {
                        v19 = v19
                        if (!(v2 != v2)) {
                            var v24 float32 = (v4 * (((v4 / 0.1) + 1.5) - v4))
                            var v25 int64 = 257
                            _ = v24
                            _ = v25
                        } else {
                            var v26 float32 = (v4 * ((v4 + float32(((v14 * float64(((3.14159 - (-2.25 / (float32((((v14 - float64(float32(((0.0 + v14) - v14)))) / float64(v4)) * 0.1)) / (v4 / v4)))) / float32(v14)))) + v14))) * v4))
                            _ = v26
                        }
                        var v27 uint64 = v15
                        _ = v27
                    } else {
                        v22 -= (v22 % (((fi17 ^ fi17)) | 1))
                    }
                case -2147483648:
                    switch (^(fi18)) {
                    case -60:
                        v19 = 1
                    case 979:
                        var v28 string = "a"
                        _ = v28
                    }
                }
                v29 := GenS4{F0: 151, F1: 194}
                _ = v22
                _ = v29
            }
            _ = fi18
            _ = v19
        }
        v3 = append(v3, fi17)
        _ = fi17
    }
    if (!false) {
        v16 = v16
    } else {
        switch 2147483647 {
        case 1:
            var v30 int64 = 939
            _ = v30
        case 128:
            v31 := GenS5{F0: genFn0()}
            _ = v31
        default:
            v2 = (v2 >> ((42) & 63))
        }
    }
    fmt.Println(v4)
    v14 = 0.1
    var v32 int32 = 0
    var v33 int32 = (v32 * (v32 ^ v32))
    if v12 {
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 int = (fi34 ^ (fi34 & -2147483648))
                v37 := GenS6{F0: ("a" + "a")}
                if (v12 && true) {
                    var v38 uint64 = v15
                    switch ("a" + "a") {
                    case "foo bar":
                        var v39 uint8 = 2
                        _ = v39
                    }
                    v40 := GenS7{F0: (^(v15)), F1: 737}
                    _ = v38
                    _ = v40
                } else {
                    fmt.Println(v12)
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            var v41 string = "a"
            _ = fi34
            _ = v41
        }
        v32 -= v32
    } else {
        v2 = (v2 - 42)
        if v12 {
            var v42 bool = false
            _ = v42
        }
        fmt.Println(v32, v33, v33)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
