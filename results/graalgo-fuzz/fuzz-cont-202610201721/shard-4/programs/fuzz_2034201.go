package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint8) int32 {
    var v6 int64 = 42
    var v7 uint8 = (p0 >> ((65535) & 63))
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int = 493
    v2 := make(map[uint32]uint32)
    var v3 bool = ((255 == 80) || (42 >= 127))
    var v4 int32 = -772
    v5 := make(map[uint32]bool)
    v8 := v5[2147483648]
    if ((0.0 * -2.25) >= 0.0) {
        var v9 uint8 = 127
        v4 += -536
        switch ("a" + "\n\t") {
        case "a":
            fmt.Println(v4)
        }
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        switch "foo bar" {
        case "a":
            v10 := []int{v1}
            _ = v10
        case "hello":
            _ = len(v5)
        }
    case "\n\t":
        v3 = v3
    case "a":
        v1 += (v1 / ((v1) | 1))
    default:
        v11 := []int{4294967295, v1}
        _ = v11
    }
    v12 := GenS1{F0: -558, F1: 42}
    switch ("" + ("" + "\n\t")) {
    case "a":
        if v3 {
            var v13 uint = 797
            v4 = 0
            _ = v13
        }
    }
    v14 := []int{}
    var v15 uint8 = 57
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint8 = v15
        for fi18 := 0; fi18 < 4; fi18++ {
            _ = len(v14)
            _ = fi18
        }
        v15 = (v15 ^ v17)
        _ = fi16
        _ = v17
    }
    switch v1 {
    case -372:
        v19 := GenS2{F0: v15}
        _ = v19
    case -1:
        v1 = v1
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        v3 = ((!true) || v3)
        var v21 uint = 2
        switch fi20 {
        case 0:
            for fi22 := 0; fi22 < 1; fi22++ {
                fmt.Println(v3, v3)
                v15 -= (v15 | v15)
                if v3 {
                    _ = len(v14)
                    if ("foo bar" >= "") {
                        v23 := GenS1{F0: v4, F1: 4294967295}
                        switch fi22 {
                        case 128:
                            switch fi22 {
                            case 32768:
                                var v24 int32 = (v4 * 343)
                                var v25 uint32 = 4294967295
                                _ = v24
                                _ = v25
                            case -2147483648:
                                fmt.Println(fi22)
                                v15 += v15
                                v26 := GenS2{F0: 0}
                                _ = v26
                            default:
                                var v27 int8 = -104
                                _ = v27
                            }
                        case 610:
                            v28 := GenS3{F0: (v4 ^ genFn0(v15)), F1: 500, F2: 42}
                            _ = v28
                        }
                        _ = v23
                    } else {
                        var v29 int64 = 9223372036854775807
                        _ = v29
                    }
                } else {
                    if (2 == 2) {
                        v30 := GenS4{F0: 2, F1: 2}
                        _ = v30
                    }
                    v5[42] = ((v15 / (((v15 + v15)) | 1)) == 1)
                    v1 = (v1 % ((v1) | 1))
                }
                _ = fi22
            }
        case 32768:
            v3 = true
        case -966:
            var v31 uint32 = 2
            _ = v31
        }
        _ = fi20
        _ = v21
    }
    switch (("\n\t" + "hello") + "\n\t") {
    case "\n\t":
        if false {
            fmt.Println(v3, v8, v4)
            switch 42 {
            case -884:
                var v32 int8 = 0
                _ = v32
            case -2147483648:
                for fi33 := 0; fi33 < 9; fi33++ {
                    var v34 bool = false
                    v14 = append(v14, (fi33 & fi33))
                    _ = fi33
                    _ = v34
                }
            case 0:
                fmt.Println(v15, v3, v15)
            }
        } else {
            _ = len(v5)
        }
    case "a":
        fmt.Println(v1, v4)
    case "":
        var v35 int16 = 1
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
