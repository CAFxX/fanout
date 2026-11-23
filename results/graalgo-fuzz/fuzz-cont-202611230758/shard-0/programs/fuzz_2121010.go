package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 uint32, p1 uint8, p2 int) int64 {
    if true {
        switch p2 {
        case 0:
            fmt.Println(p0, p2, p1)
        case 46:
            if (!true) {
                var v6 float32 = 1.5
                for fi7 := 0; fi7 < 13; fi7++ {
                    v8 := []int{(p2 | fi7), (p2 & (^(p2))), (fi7 - (p2 - 42)), (fi7 - int(p1))}
                    _ = fi7
                    _ = v8
                }
                _ = v6
            }
        case 2147483647:
            var v9 int32 = 1
            _ = v9
        default:
            for fi10 := 0; fi10 < 13; fi10++ {
                p0 = (^(p0))
                _ = fi10
            }
        }
        v11 := []int{128, p2}
        _ = v11
    } else {
        fmt.Println(p1)
        v12 := []int{(p2 >> ((1) & 63))}
        _ = v12
    }
    v13 := GenS1{F0: 255, F1: (!(p2 > 42)), F2: 42}
    v14 := GenS2{F0: -32768, F1: 1}
    var v15 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 1
}
func genFn1(p0 uint32, p1 int, p2 int16) int {
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int32 = -2147483648
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: 0}
    v19 := []int{p1}
    var v20 int32 = 0
    var v21 uint = 778
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    _ = v21
    return 723
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]int)
    var v3 uint8 = 127
    var v4 int64 = 0
    var v5 uint = 255
    v4 = 9223372036854775807
    _ = len(v2)
    for fi22 := 0; fi22 < 16; fi22++ {
        if (!(true && ((!(((4294967295 >= (v1 ^ 2)) || (uint16(v1) <= 351)) || true)) && false))) {
            if (genFn1(0, fi22, 42) != (fi22 % (((fi22 ^ 42)) | 1))) {
                v4 = v4
            } else {
                v23 := []int{genFn1((v1 * (v1 >> ((653) & 63))), fi22, -32768), 32768}
                _ = v23
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                v4 += v4
                _ = fi24
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                if true {
                    fmt.Println(v5)
                    var v26 float32 = 1e10
                    fmt.Println(fi22)
                    _ = v26
                }
                if (v5 == (v5 ^ v5)) {
                    v27 := []int{(fi25 & (fi25 * fi25)), (fi25 << (((v5 ^ 54)) & 63))}
                    _ = v27
                }
                _ = fi25
            }
        }
        var v28 uint32 = 0
        if false {
            switch fi22 {
            case 32767:
                v28 += 4294967295
            case 1:
                switch "\n\t" {
                case "foo bar":
                    var v29 uint64 = 4294967295
                    _ = v29
                default:
                    if (int16(v28) < 42) {
                        var v30 uint = (v5 ^ 42)
                        _ = v30
                    }
                }
            case 4294967295:
                fmt.Println(fi22, v5)
            default:
                fmt.Println(v5, v4, v3)
            }
            v31 := GenS4{F0: (128 == 136)}
            for fi32 := 0; fi32 < 20; fi32++ {
                fmt.Println(v1, v1, v4)
                v28 = (v1 ^ v1)
                fmt.Println(v28, v1, v1)
                _ = fi32
            }
            _ = v31
        }
        _ = fi22
        _ = v28
    }
    v3 -= v3
    v2[(^(v1))] = genFn1(2, 9223372036854775807, 1)
    var v33 int64 = (v4 + 9223372036854775807)
    var v34 bool = true
    var v35 int32 = 729
    if v34 {
        var v36 uint16 = 0
        fmt.Println(v4, v34, v4)
        _ = v36
    }
    if true {
        v37 := GenS4{F0: false}
        if (v34 && v34) {
            switch -683 {
            case 4294967295:
                for fi38 := 0; fi38 < 20; fi38++ {
                    var v39 uint64 = 18446744073709551615
                    v39 = (v39 + v39)
                    _ = fi38
                    _ = v39
                }
            }
            for fi40 := 0; fi40 < 4; fi40++ {
                _ = len(v2)
                _ = fi40
            }
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v1, v3)
            default:
                var v41 uint8 = 2
                _ = v41
            }
            v42 := []int{}
            _ = v42
        }
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 int32 = (v35 / (((^(v35))) | 1))
            _ = len(v2)
            _ = fi43
            _ = v44
        }
        _ = v37
    } else {
        if v34 {
            var v45 uint8 = v3
            var v46 int64 = 0
            fmt.Println(v4)
            _ = v45
            _ = v46
        }
    }
    if v34 {
        v47 := []int{genFn1(16, genFn1(v1, -83, 0), -572)}
        v48 := GenS5{F0: v5, F1: "hello"}
        v3 = (v3 + 128)
        _ = v47
        _ = v48
    }
    var v49 float32 = ((3.14159 + (3.14159 / 0.1)) - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v49
    _ = v5
}
