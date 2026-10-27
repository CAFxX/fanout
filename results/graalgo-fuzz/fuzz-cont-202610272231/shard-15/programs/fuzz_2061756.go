package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint32 }
func genFn0(p0 int32) int32 {
    var v6 uint32 = 350
    v7 := GenS1{F0: 65535, F1: 128}
    _ = p0
    _ = v6
    _ = v7
    return (p0 + (p0 % ((p0) | 1)))
}
func genFn1(p0 string, p1 uint16, p2 int8) int {
    for fi8 := 0; fi8 < 1; fi8++ {
        p2 += p2
        v9 := []int{(^(fi8)), (fi8 << ((757) & 63)), (fi8 + (fi8 / (((fi8 % ((fi8) | 1))) | 1))), fi8}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if (true && true) {
                var v12 float64 = 0.1
                if ((((((3.14159 + v12) + v12) - 3.14159) / v12) + v12) >= float64(3.14159)) {
                    var v13 uint32 = 4294967295
                    p1 = 2
                    v14 := GenS1{F0: p1, F1: (fi10 + (fi10 % ((fi11) | 1)))}
                    _ = v13
                    _ = v14
                }
                _ = v12
            }
            _ = fi11
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int16 = 42
        var v17 int8 = p2
        v16 += v16
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 int64 = 0
    v2 := []int{32767, 128}
    var v3 bool = (false && false)
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    var v18 int8 = 42
    v19 := GenS1{F0: 999, F1: -277}
    v18 += (v18 & (v18 / ((-128) | 1)))
    if (v18 != v18) {
        if (478 == 4294967295) {
            var v20 bool = true
            v5[-271] = (0 <= 0)
            _ = len(v4)
            _ = v20
        }
    }
    v21 := GenS1{F0: 13, F1: 9223372036854775807}
    if (v3 && (65535 >= 18446744073709551615)) {
        v3 = v3
    } else {
        v3 = false
        switch -650 {
        case -1:
            v22 := GenS1{F0: 2, F1: 32767}
            _ = v22
        }
        if v3 {
            switch "foo bar" {
            case "hello":
                v23 := v5[2147483647]
                _ = v23
            case "a":
                if (-1 < 32767) {
                    v18 = (v18 - (^(v18)))
                    if false {
                        if v3 {
                            for fi24 := 0; fi24 < 17; fi24++ {
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    if v3 {
                                        v2 = append(v2, (fi24 % (((-(fi25))) | 1)))
                                    }
                                    fmt.Println(fi25, fi24, v18)
                                    v26 := GenS1{F0: 329, F1: 128}
                                    _ = fi25
                                    _ = v26
                                }
                                v18 = 19
                                var v27 uint = 255
                                _ = fi24
                                _ = v27
                            }
                            switch 32768 {
                            case 1:
                                v18 += v18
                            }
                        }
                        var v28 int8 = (v18 * v18)
                        _ = v28
                    }
                } else {
                    v18 = int8(v1)
                    var v29 int = 461
                    v1 -= v1
                    _ = v29
                }
            case "":
                v3 = ((true || v3) && (128 <= 0))
            }
        }
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        if ((v18 % (((v18 & v18)) | 1)) < v18) {
            if (480 > -219) {
                v31 := GenS2{F0: 1, F1: 0}
                _ = v31
            } else {
                if false {
                    var v32 int64 = (v1 >> ((667) & 63))
                    var v33 float32 = 3.14159
                    _ = v32
                    _ = v33
                }
            }
            var v34 string = "\n\t"
            _ = v34
        } else {
            v35 := GenS3{F0: 826, F1: 1, F2: 4294967295}
            _ = v35
        }
        var v36 bool = (255 < 0)
        _ = fi30
        _ = v36
    }
    _ = len(v2)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
