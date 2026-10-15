package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 int8, p2 string) bool {
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5, p0, fi5)
        if ((32767 > 144) && (!true)) {
            switch fi5 {
            case -850:
                for fi6 := 0; fi6 < 2; fi6++ {
                    v7 := GenS1{F0: -1}
                    _ = fi6
                    _ = v7
                }
            }
        } else {
            var v8 float32 = (1.5 + 3.14159)
            var v9 int32 = 42
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: -78}
        _ = fi5
        _ = v10
    }
    v11 := GenS1{F0: 1}
    var v12 uint64 = 104
    switch 854 {
    case 32768:
        var v13 uint64 = (^(v12))
        _ = v13
    case 127:
        var v14 uint = 255
        _ = v14
    default:
        var v15 int8 = 104
        _ = v15
    }
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (32767 <= -32768)
}
func genFn1(p0 int8, p1 uint8, p2 uint8) int32 {
    p2 = 255
    p2 -= (p1 | (p2 / (((p2 & p1)) | 1)))
    var v16 int16 = 0
    v17 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -165
}

func main() {
    var v1 uint = 0
    v2 := make(map[int]bool)
    var v3 int64 = 0
    var v4 float64 = 3.14159
    v18 := GenS2{F0: 4294967295, F1: 115}
    v19 := GenS3{F0: 0, F1: 255}
    v1 = v1
    switch 409 {
    case 32768:
        v20 := GenS3{F0: 65535, F1: 255}
        _ = v20
    case -601:
        var v21 float32 = ((3.14159 * 0.1) + 1e10)
        _ = v21
    default:
        if true {
            var v22 int32 = -1
            var v23 string = "a"
            if true {
                switch ((v23 + v23) + "foo bar") {
                case "hello":
                    if false {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            for fi25 := 0; fi25 < 18; fi25++ {
                                var v26 int = (^(fi24))
                                v26 = 27
                                _ = len(v2)
                                _ = fi25
                                _ = v26
                            }
                            if true {
                                var v27 uint8 = 128
                                v28 := []int{}
                                _ = v27
                                _ = v28
                            }
                            fmt.Println(v23, v23, v22)
                            _ = fi24
                        }
                    } else {
                        fmt.Println(v3)
                        var v29 int64 = -9223372036854775808
                        switch v23 {
                        case "foo bar":
                            switch "\n\t" {
                            case "foo bar":
                                var v30 uint = (v1 / ((v1) | 1))
                                _ = v30
                            }
                        case "a":
                            v31 := v2[640]
                            _ = v31
                        }
                        _ = v29
                    }
                }
            }
            _ = v22
            _ = v23
        } else {
            fmt.Println(v4, v3, v4)
        }
    }
    if (!false) {
        v32 := []int{int(v1), 4294967295, 32768, 32767}
        var v33 uint8 = 255
        _ = v32
        _ = v33
    }
    fmt.Println(v4, v1)
    var v34 string = "hello"
    v34 = (v34 + "foo bar")
    var v35 string = v34
    v36 := GenS4{F0: 127, F1: 128, F2: 286}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
