package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool) int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 = (2 >= 18446744073709551615)
        _ = fi7
    }
    fmt.Println(p0)
    _ = p0
    return -1
}
func genFn1(p0 uint, p1 int64) int32 {
    v8 := GenS1{F0: 64, F1: 79}
    v9 := []int{0, 0, 4294967295, 211}
    var v10 uint32 = 1
    var v11 int8 = -41
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -216
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{127, 0}
    v3 := make(map[uint32]string)
    v4 := []int{32767, 201, 32768}
    var v5 int8 = 1
    v6 := make(map[int]string)
    var v12 int64 = -892
    v13 := GenS2{F0: 161, F1: 861}
    v12 = 1
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v5)
        v14 += 1e10
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(fi16)
            if (!(42 == 0)) {
                switch "\n\t" {
                case "a":
                    switch "a" {
                    case "foo bar":
                        if (-2147483648 > 616) {
                            var v17 uint = 344
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(fi15, fi15, v17)
                            _ = v17
                        }
                    default:
                        delete(v3, 1)
                    }
                }
                var v18 int8 = -128
                _ = v18
            }
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println(v14)
    switch 127 {
    case 127:
        switch ("hello" + "\n\t") {
        case "a":
            v19 := GenS2{F0: 2147483647, F1: 1}
            _ = v19
        case "":
            _ = len(v2)
        case "foo bar":
            var v20 float64 = (1.5 * 1e10)
            _ = v20
        }
    case -1:
        var v21 int32 = -1
        _ = v21
    case 1:
        if true {
            if (false && false) {
                v6[511] = ""
                v14 += (v14 - 0.1)
            } else {
                var v22 bool = (v5 <= v5)
                if true {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 int64 = v12
                        _ = len(v1)
                        switch "foo bar" {
                        case "foo bar":
                            switch fi23 {
                            case 128:
                                for fi25 := 0; fi25 < 20; fi25++ {
                                    switch (fi25 / (((-(fi23))) | 1)) {
                                    case 989:
                                        v26 := GenS2{F0: 1, F1: 42}
                                        _ = v26
                                    case 0:
                                        v27 := GenS3{F0: int64(fi25), F1: 823}
                                        _ = v27
                                    case -395:
                                        var v28 int64 = (-(v24))
                                        _ = v28
                                    }
                                    v3[1] = "\n\t"
                                    v29 := GenS1{F0: 65535, F1: (v5 << ((1) & 63))}
                                    fmt.Println(v14, v5, v14)
                                    v30 := GenS1{F0: 2, F1: (v5 + (v5 * (v5 / (((v5 & (v5 << ((1) & 63)))) | 1))))}
                                    _ = fi25
                                    _ = v29
                                    _ = v30
                                }
                            case 9223372036854775807:
                                v31 := GenS2{F0: 111, F1: 42}
                                _ = v31
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case "\n\t":
                            var v32 bool = (true && v22)
                            _ = v32
                        }
                        _ = fi23
                        _ = v24
                    }
                    v6[32767] = "\n\t"
                }
                _ = v22
            }
            var v33 uint64 = 0
            _ = v33
        } else {
            v34 := GenS2{F0: 213, F1: 1}
            v14 -= float32(3.14159)
            _ = v34
        }
    }
    v35 := GenS4{F0: 0}
    var v36 int = -856
    var v37 int16 = 32767
    v36 -= 2147483647
    for fi38 := 0; fi38 < 17; fi38++ {
        v14 = v14
        _ = fi38
    }
    v39 := GenS3{F0: (^(v12)), F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
