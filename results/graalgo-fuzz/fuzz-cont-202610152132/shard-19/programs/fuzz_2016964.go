package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    var v3 uint32 = 4294967295
    var v4 float32 = float32((1e10 * (float64((1e10 - ((((float32((float64(0.1) + (((1e10 + (0.0 * (float64(0.0) + 0.1))) + 1.5) * float64((0.1 - 0.0))))) + 0.0) * (float32(((1.5 + 0.1) * 0.0)) / 1.5)) * 0.0) / 1e10))) * ((0.1 / 1e10) + 0.1))))
    var v5 uint64 = 694
    v6 := make(map[string]int)
    var v7 uint64 = (^(v5))
    switch ("\n\t" + ("a" + "\n\t")) {
    case "hello":
        fmt.Println(v4, v4, v5)
    default:
        v8 := GenS1{F0: uint16(v7), F1: 1, F2: (1 >= 255)}
        _ = v8
    }
    var v9 uint32 = v3
    var v10 float32 = (1.5 / v4)
    fmt.Println(v4, v3, v4)
    var v11 float32 = ((v10 / ((v4 / v10) - (v10 - v10))) * v4)
    v12 := GenS1{F0: 1, F1: 42, F2: ((0 != -39) || true)}
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v7, fi13)
        switch ("\n\t" + "\n\t") {
        case "hello":
            v14 := GenS1{F0: 2, F1: 1, F2: (v7 >= v7)}
            _ = v14
        case "foo bar":
            v5 += (v5 ^ v5)
        default:
            var v15 int64 = 706
            _ = v15
        }
        _ = fi13
    }
    var v16 int32 = v1
    v17 := GenS1{F0: 895, F1: 996, F2: false}
    var v18 float32 = v11
    for fi19 := 0; fi19 < 2; fi19++ {
        if false {
            v10 -= 1e10
        }
        _ = fi19
    }
    v20 := GenS2{F0: 42, F1: (v5 | 2), F2: 0}
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            if (true || false) {
                delete(v6, (("\n\t" + "a") + (("" + "") + ("foo bar" + "foo bar"))))
                var v23 bool = (("hello" + ("a" + "foo bar")) >= "\n\t")
                var v24 int = fi22
                _ = v23
                _ = v24
            } else {
                if ((^(v9)) != 2) {
                    var v25 bool = false
                    v26 := GenS3{F0: 42}
                    switch "foo bar" {
                    case "hello":
                        switch (fi22 | 160) {
                        case 4294967295:
                            if (v25 && (v25 && true)) {
                                v27 := GenS1{F0: 255, F1: 42, F2: v25}
                                fmt.Println(v18, v9)
                                _ = v27
                            } else {
                                v18 = float32(3.14159)
                            }
                        default:
                            if (100 > -40) {
                                v25 = false
                            } else {
                                v10 -= v11
                            }
                        }
                    case "a":
                        v3 = v9
                    }
                    _ = v25
                    _ = v26
                }
                var v28 bool = ((!false) && (128 >= fi22))
                _ = len(v6)
                _ = v28
            }
            _ = fi22
        }
        var v29 uint32 = (v9 / (((v3 - v9)) | 1))
        v1 = 528
        _ = fi21
        _ = v29
    }
    v30 := GenS3{F0: 1}
    v31 := GenS4{F0: "a", F1: "", F2: 1}
    var v32 int = -391
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
