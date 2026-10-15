package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[string]string)
    var v3 bool = (!(!(float32(1.5) < v1)))
    var v4 int16 = 32767
    v5 := make(map[int]string)
    v1 = float32(1.5)
    v1 += ((-2.25 * (v1 * v1)) * (v1 + (3.14159 + 0.0)))
    v6 := GenS1{F0: -242, F1: -2147483648}
    v7 := GenS1{F0: 42, F1: 493}
    var v8 int8 = 125
    v3 = (v3 && (v3 && (512 != 42)))
    v9 := []int{127}
    var v10 uint8 = 255
    v11 := GenS1{F0: -1, F1: -1}
    var v12 bool = true
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v4, v10)
            if (!(uint16(v10) <= 0)) {
                fmt.Println(fi14, v3)
                v15 := GenS2{F0: 240, F1: 4294967295}
                if v3 {
                    delete(v5, (fi14 ^ fi13))
                    v12 = v3
                    if false {
                        var v16 int = (fi13 << ((887) & 63))
                        var v17 int16 = v4
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            v1 = ((((v1 - 0.1) + 0.1) - v1) + v1)
                        case "hello":
                            delete(v5, (fi13 * (^(fi14))))
                        case "foo bar":
                            v9 = append(v9, (-(fi14)))
                        }
                        _ = v16
                        _ = v17
                    } else {
                        if true {
                            fmt.Println(fi14, v1, v8)
                            if (0 != 42) {
                                v18 := GenS3{F0: (v4 | (v4 * (v4 | (^(v4))))), F1: "hello"}
                                v8 += (v8 | (v8 << ((843) & 63)))
                                delete(v5, int(v8))
                                _ = v18
                            }
                            switch fi14 {
                            case 242:
                                var v19 int16 = int16(v10)
                                _ = v19
                            case 4294967295:
                                v10 = v10
                            default:
                                var v20 uint16 = 42
                                _ = v20
                            }
                        }
                        var v21 int16 = int16(v10)
                        _ = v21
                    }
                } else {
                    var v22 bool = v12
                    _ = v22
                }
                _ = v15
            } else {
                fmt.Println(fi14, v1, fi14)
                v23 := GenS4{F0: -2147483648}
                _ = v23
            }
            _ = fi14
        }
        var v24 string = "hello"
        if (false && false) {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = ((("a" + "hello") + v24) + (((v24 + "foo bar") + ("" + v24)) + v24))
                var v27 uint64 = 23
                v28 := GenS2{F0: (v10 % (((v10 | v10)) | 1)), F1: (v27 - uint64(v10))}
                fmt.Println(v3, v27)
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                _ = len(v9)
                _ = len(v5)
                v30 := GenS5{F0: v8, F1: 78}
                _ = fi29
                _ = v30
            }
        }
        _ = fi13
        _ = v24
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v8, v10)
    case "hello":
        var v31 float64 = 3.14159
        _ = v31
    case "a":
        var v32 int = 1
        _ = v32
    }
    var v33 int = int(v10)
    var v34 int16 = (v4 ^ -154)
    v4 += (v4 << ((0) & 63))
    var v35 uint16 = 32767
    switch "hello" {
    case "foo bar":
        v35 = v35
    case "a":
        v8 += v8
    case "\n\t":
        fmt.Println(v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
