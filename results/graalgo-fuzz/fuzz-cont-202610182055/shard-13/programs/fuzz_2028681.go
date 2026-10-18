package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    v6 := []int{-285, int(p0), 647, 435}
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    _ = p0
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 bool = false
    var v4 int32 = 1
    var v5 float32 = 0.1
    var v7 uint16 = 288
    var v8 uint8 = 1
    v9 := v2["foo bar"]
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            if (1 != v4) {
                fmt.Println(fi11)
                v5 = (v5 + (v5 * 0.0))
            } else {
                var v12 uint = 2
                fmt.Println(v7, v5, fi11)
                var v13 float32 = v5
                _ = len(v2)
                switch "" {
                case "a":
                    var v14 uint = (v12 % ((v12) | 1))
                    switch "" {
                    case "a":
                        var v15 int64 = -851
                        _ = v15
                    case "\n\t":
                        switch fi11 {
                        case -1:
                            switch (fi11 & fi11) {
                            case 5:
                                var v16 int16 = -606
                                _ = v16
                            case -2147483648:
                                fmt.Println(v4)
                                v8 -= 127
                            default:
                                fmt.Println(v4, v14)
                            }
                        case 0:
                            v17 := GenS1{F0: "a", F1: (v7 | v7), F2: 2}
                            _ = v17
                        }
                    case "foo bar":
                        var v18 int8 = -31
                        _ = v18
                    }
                    _ = v14
                default:
                    var v19 int8 = -1
                    _ = v19
                }
                _ = v12
                _ = v13
            }
            switch "hello" {
            case "hello":
                var v20 float32 = v5
                _ = v20
            default:
                if ((1 < 0) && true) {
                    var v21 int8 = 1
                    var v22 uint16 = (v7 >> ((uint(v7)) & 63))
                    fmt.Println(v21, v3, v5)
                    _ = v21
                    _ = v22
                }
            }
            var v23 int8 = -126
            _ = fi11
            _ = v23
        }
        var v24 uint64 = 4294967295
        _ = fi10
        _ = v24
    }
    var v25 int = 32768
    v26 := GenS2{F0: 4294967295}
    var v27 uint8 = v8
    if v3 {
        v7 = v7
        v28 := v1[2147483647]
        _ = v28
    }
    if ((v8 < v27) || true) {
        fmt.Println(v9, v9)
        var v29 int32 = (^(v4))
        for fi30 := 0; fi30 < 8; fi30++ {
            v9 = false
            fmt.Println(v8)
            _ = fi30
        }
        _ = v29
    }
    var v31 bool = (v9 || v9)
    v25 = v25
    v32 := GenS3{F0: 2, F1: 0, F2: 2147483648}
    v8 += 148
    for fi33 := 0; fi33 < 20; fi33++ {
        switch -580 {
        case 2147483647:
            var v34 string = ""
            _ = v34
        case 32767:
            for fi35 := 0; fi35 < 4; fi35++ {
                v31 = true
                fmt.Println(v8, v5, fi35)
                var v36 uint8 = v27
                _ = fi35
                _ = v36
            }
        case 0:
            fmt.Println(v27, v3, v9)
        default:
            var v37 uint64 = 389
            _ = v37
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 string = "a"
            _ = fi38
            _ = v39
        }
        var v40 float64 = ((-2.25 / (1e10 + 3.14159)) / float64(v5))
        _ = fi33
        _ = v40
    }
    var v41 int = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
