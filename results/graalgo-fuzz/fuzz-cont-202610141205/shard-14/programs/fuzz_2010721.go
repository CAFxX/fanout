package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 uint8) int32 {
    p0 += p0
    p0 = p0
    switch -343 {
    case 0:
        var v7 uint = uint(p0)
        _ = v7
    case -2147483648:
        fmt.Println(p0)
    default:
        var v8 float64 = 1e10
        _ = v8
    }
    v9 := GenS1{F0: 617, F1: 65535, F2: "a"}
    _ = p0
    _ = v9
    return 2147483647
}
func genFn1() int64 {
    switch (("" + ("\n\t" + "")) + "a") {
    case "a":
        switch -931 {
        case 166:
            var v10 uint8 = 57
            _ = v10
        case 32768:
            var v11 int = 128
            _ = v11
        case 128:
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int32 = 2147483647
                for fi14 := 0; fi14 < 7; fi14++ {
                    switch "" {
                    case "\n\t":
                        if false {
                            v15 := GenS1{F0: 65535, F1: 0, F2: ""}
                            v16 := GenS2{F0: 117, F1: 42, F2: 4294967295}
                            _ = v15
                            _ = v16
                        }
                    case "":
                        v13 += 0
                    case "hello":
                        var v17 uint = 798
                        _ = v17
                    }
                    if false {
                        var v18 int64 = -1
                        _ = v18
                    }
                    _ = fi14
                }
                switch fi12 {
                case 32768:
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v13 = (v13 % (((v13 ^ (v13 | (v13 / (((v13 & -470)) | 1))))) | 1))
                        v20 := GenS3{F0: false}
                        var v21 uint = 2
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi12
                _ = v13
            }
        default:
            var v22 int8 = -113
            _ = v22
        }
    case "\n\t":
        var v23 int32 = genFn0(222)
        _ = v23
    }
    var v24 int8 = -1
    v24 = v24
    v24 += 15
    v24 = (v24 + 127)
    _ = v24
    return -1
}

func main() {
    var v1 uint8 = 73
    var v2 float64 = (0.1 + 1.5)
    v3 := make(map[string]int)
    var v4 string = ("\n\t" + "hello")
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS2{F0: (^(v1)), F1: 688, F2: 1}
        _ = fi25
        _ = v26
    }
    if true {
        var v27 int8 = -128
        var v28 int64 = genFn1()
        var v29 int16 = 1000
        _ = v27
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 20; fi30++ {
            fmt.Println(v4, v1, v1)
            _ = fi30
        }
    }
    v31 := GenS1{F0: 563, F1: 65535, F2: v4}
    v4 = v4
    fmt.Println(v1)
    v1 += 133
    switch 4294967295 {
    case 80:
        v32 := GenS1{F0: 255, F1: 634, F2: v4}
        _ = v32
    case -222:
        fmt.Println(v2, v4)
    case 1:
        v33 := GenS1{F0: 871, F1: 2147483647, F2: v4}
        _ = v33
    default:
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 float64 = v2
            v36 := []int{389, (fi34 | (^(fi34))), (fi34 & fi34)}
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    v37 := GenS3{F0: (false && false)}
    for fi38 := 0; fi38 < 11; fi38++ {
        v2 = v2
        _ = fi38
    }
    var v39 uint16 = 255
    var v40 uint16 = v39
    v39 -= v39
    var v41 uint = 698
    v1 = v1
    v42 := GenS4{F0: 42, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
