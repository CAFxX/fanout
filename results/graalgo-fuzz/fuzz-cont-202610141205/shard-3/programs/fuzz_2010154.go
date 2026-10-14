package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 string; F1 int32; F2 string }
func genFn0() uint32 {
    v5 := []int{128}
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 = append(v5, (fi6 | 32767))
        var v7 uint32 = 2147483648
        fmt.Println(fi6, v7)
        _ = fi6
        _ = v7
    }
    var v8 int = -82
    var v9 uint8 = 22
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(fi10, v8)
        var v11 uint8 = v9
        _ = fi10
        _ = v11
    }
    _ = v5
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int64 = 204
    var v2 int8 = 42
    v3 := make(map[uint32]string)
    var v4 int64 = (v1 % (((v1 ^ v1)) | 1))
    fmt.Println(v4)
    v4 = 757
    v2 = v2
    v12 := []int{127, 2147483647, 1, 2147483647}
    var v13 float32 = (0.1 / (1.5 / 1e10))
    var v14 int = 0
    if ((546 > 0) && (65535 == 801)) {
        v15 := GenS1{F0: (18446744073709551615 != 0)}
        _ = v15
    }
    if false {
        fmt.Println(v2, v2, v1)
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint16 = 1
            v1 = 1
            _ = fi16
            _ = v17
        }
        fmt.Println(v4, v1)
    } else {
        v2 += 100
        v2 += 103
        v1 = v4
    }
    var v18 int = (v14 >> ((255) & 63))
    var v19 bool = true
    switch "hello" {
    case "":
        var v20 float32 = (v13 - v13)
        _ = v20
    }
    switch "" {
    case "\n\t":
        if v19 {
            for fi21 := 0; fi21 < 5; fi21++ {
                if ((!v19) && v19) {
                    v22 := GenS1{F0: true}
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 string = ""
                        switch v24 {
                        case "":
                            v13 += v13
                        case "foo bar":
                            fmt.Println(v24)
                        case "a":
                            v13 = v13
                        }
                        _ = fi23
                        _ = v24
                    }
                    var v25 int = -55
                    _ = v22
                    _ = v25
                }
                _ = len(v12)
                _ = fi21
            }
            var v26 int32 = 292
            _ = v26
        } else {
            fmt.Println(v14)
            if v19 {
                var v27 string = "\n\t"
                _ = v27
            }
        }
    case "hello":
        var v28 float32 = -2.25
        _ = v28
    }
    v29 := GenS2{F0: 42, F1: 32767}
    v14 = (v14 + v14)
    for fi30 := 0; fi30 < 5; fi30++ {
        v4 += (v1 + v1)
        for fi31 := 0; fi31 < 5; fi31++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            v32 := GenS3{F0: "hello", F1: 617, F2: ("\n\t" + "")}
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
