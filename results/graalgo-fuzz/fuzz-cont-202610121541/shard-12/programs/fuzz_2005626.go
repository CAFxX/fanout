package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
func genFn0(p0 int16, p1 int16, p2 int16) int64 {
    p2 -= (p1 * p0)
    if false {
        p2 = (p0 - p2)
        for fi7 := 0; fi7 < 10; fi7++ {
            if true {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 float32 = (3.14159 * 1e10)
                    _ = fi8
                    _ = v9
                }
                p2 -= (p0 + p0)
                v10 := []int{128}
                _ = v10
            }
            var v11 int32 = 0
            for fi12 := 0; fi12 < 8; fi12++ {
                switch "hello" {
                case "\n\t":
                    v13 := GenS1{F0: 1, F1: (v11 >> ((1) & 63)), F2: 2147483648}
                    _ = v13
                }
                if (!false) {
                    if true {
                        switch "\n\t" {
                        case "a":
                            p2 = p2
                        case "\n\t":
                            v11 = (v11 * 0)
                        default:
                            if false {
                                var v14 int64 = 42
                                _ = v14
                            }
                        }
                    } else {
                        if (((!true) || (4294967295 >= 18446744073709551615)) || true) {
                            v15 := GenS1{F0: (p1 & p1), F1: 2147483647, F2: 0}
                            var v16 int16 = p0
                            v16 = (p2 * p2)
                            _ = v15
                            _ = v16
                        } else {
                            var v17 uint16 = 42
                            var v18 bool = ((!(2 != 1)) && (-1 == 393))
                            v17 = v17
                            v19 := []int{(fi7 % (((^(fi7))) | 1)), fi7, (fi7 % (((fi7 - 2147483647)) | 1))}
                            v19 = append(v19, -875)
                            p2 = (p0 & (p0 + -624))
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        p2 += p2
                        v20 := GenS2{F0: 0, F1: 1, F2: 4294967295}
                        _ = v20
                    }
                    switch fi7 {
                    case 127:
                        v21 := []int{(fi12 + -868), (fi12 - -2147483648), fi12}
                        _ = v21
                    case -696:
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            p1 += p2
                            p1 += (^(p1))
                            p1 -= (p0 / ((p1) | 1))
                        }
                    }
                    var v22 uint16 = 65535
                    _ = v22
                }
                if false {
                    v23 := GenS1{F0: (p1 + int16(v11)), F1: (v11 % ((v11) | 1)), F2: 1}
                    p1 += p1
                    if true {
                        if false {
                            var v24 float32 = 0.1
                            fmt.Println(v24)
                            _ = v24
                        } else {
                            fmt.Println(fi12, p1, p0)
                            fmt.Println(p0, fi7, fi7)
                        }
                    } else {
                        p1 = (^(p2))
                        var v25 uint16 = 65535
                        _ = v25
                    }
                    _ = v23
                }
                v26 := []int{fi7}
                var v27 uint8 = 111
                _ = fi12
                _ = v26
                _ = v27
            }
            _ = fi7
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}
func genFn1() uint64 {
    if true {
        var v28 int = 1
        _ = v28
    } else {
        if (!(!false)) {
            var v29 int64 = 9223372036854775807
            v30 := GenS1{F0: 0, F1: 0, F2: 42}
            _ = v29
            _ = v30
        }
    }
    if true {
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint32 = 2
            if true {
                var v33 int64 = -9223372036854775808
                _ = v33
            }
            v32 -= (v32 ^ 2147483648)
            _ = fi31
            _ = v32
        }
    }
    return 904
}

func main() {
    var v1 float32 = (0.0 * 0.1)
    var v2 float32 = v1
    v3 := make(map[int]bool)
    var v4 int8 = -123
    v5 := make(map[uint32]bool)
    var v6 uint16 = 32767
    v1 += float32(0.0)
    var v34 uint32 = 211
    if false {
        fmt.Println(v2, v34, v6)
        v6 = 1
        var v35 uint8 = 128
        _ = v35
    }
    v1 += (0.0 + 1.5)
    for fi36 := 0; fi36 < 5; fi36++ {
        v3[2147483647] = false
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS1{F0: -568, F1: 0, F2: (^(v34))}
            v2 -= (v2 + v2)
            if false {
                v2 -= v1
            } else {
                fmt.Println(fi36, fi36)
                v1 = 1e10
                v6 -= v6
            }
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    var v39 int32 = -2147483648
    v1 = 1.5
    fmt.Println(v6, v6)
    v40 := GenS3{F0: "hello", F1: 2, F2: -32768}
    var v41 uint16 = v6
    fmt.Println(v34)
    v42 := GenS2{F0: (v4 | v4), F1: genFn1(), F2: v34}
    v3[128] = (v4 != 100)
    v43 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
