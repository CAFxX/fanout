package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 int16, p1 int16, p2 uint32) int32 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{(fi7 ^ (fi7 * fi7)), (-(fi7)), fi7, (fi7 << ((18446744073709551615) & 63))}
        if (1 < p1) {
            var v9 int = (^(fi7))
            fmt.Println(fi7)
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 bool = (-13 < 42)
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 string = "\n\t"
            _ = fi12
            _ = v13
        }
        switch fi11 {
        case 127:
            fmt.Println(v10, p0)
        case 288:
            p0 += (p0 & (p0 | (-(p0))))
        case 32767:
            var v14 float64 = ((0.0 * 0.0) + (float64(0.1) - ((1.5 / float64(float32(float64(1.5)))) / 0.1)))
            _ = v14
        default:
            switch "foo bar" {
            case "":
                v15 := []int{(fi11 / ((fi11) | 1)), 9223372036854775807, (fi11 | fi11), (fi11 / ((fi11) | 1))}
                _ = v15
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 int16) int {
    fmt.Println(p0)
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 uint16 = 65535
            if (!(255 < (v17 >> ((493) & 63)))) {
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := []int{}
                    _ = fi18
                    _ = v19
                }
                var v20 float32 = ((0.1 - 3.14159) * 0.1)
                var v21 int8 = 0
                _ = v20
                _ = v21
            } else {
                if false {
                    p0 = -665
                    v22 := GenS1{F0: p0, F1: false, F2: false}
                    v23 := GenS2{F0: 384, F1: ((!false) || true)}
                    _ = v22
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        p0 = p0
                        if false {
                            p0 = (-(p0))
                        } else {
                            fmt.Println(p0, fi24)
                        }
                        v25 := GenS3{F0: -9223372036854775808, F1: 4294967295, F2: 2}
                        _ = fi24
                        _ = v25
                    }
                    v26 := []int{fi16, fi16, fi16, (fi16 + (fi16 << ((255) & 63)))}
                    v27 := GenS1{F0: p0, F1: false, F2: (!(32767 <= p0))}
                    _ = v26
                    _ = v27
                }
            }
            _ = fi16
            _ = v17
        }
    } else {
        var v28 bool = false
        _ = v28
    }
    _ = p0
    return 127
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 uint = 18446744073709551615
    var v6 uint = 65535
    if (false || true) {
        for fi29 := 0; fi29 < 15; fi29++ {
            for fi30 := 0; fi30 < 11; fi30++ {
                fmt.Println(v6)
                v5 += 18446744073709551615
                var v31 float64 = ((1.5 * 0.0) - float64((0.0 / 3.14159)))
                var v32 uint64 = 42
                var v33 int64 = -9223372036854775808
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            if (fi29 != genFn1(-32768)) {
                v5 = v5
            } else {
                var v34 bool = false
                var v35 uint32 = uint32(v6)
                v35 += (v35 | (v35 >> (((v5 % (((^(v6))) | 1))) & 63)))
                _ = v34
                _ = v35
            }
            _ = fi29
        }
        if false {
            var v36 int64 = 0
            _ = v36
        } else {
            if true {
                var v37 uint8 = 24
                var v38 uint = v6
                _ = v37
                _ = v38
            }
            fmt.Println(v5, v1)
            if true {
                if (!(819 < 722)) {
                    v5 += v6
                } else {
                    v6 = 180
                }
            } else {
                v39 := GenS4{F0: 42, F1: 0}
                _ = v39
            }
        }
    } else {
        v4[genFn1(42)] = 0
        v40 := []int{291, -434, -470}
        _ = v40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        fmt.Println(fi41, v6, v6)
        var v42 uint64 = 4294967295
        _ = fi41
        _ = v42
    }
    if false {
        for fi43 := 0; fi43 < 18; fi43++ {
            if (754 > 4294967295) {
                var v44 string = "\n\t"
                _ = v44
            }
            if false {
                var v45 uint = v6
                var v46 int16 = 42
                fmt.Println(v1, v46, fi43)
                _ = v45
                _ = v46
            }
            _ = fi43
        }
        fmt.Println(v5, v5)
        fmt.Println(v6, v5, v6)
    }
    v6 = v6
    v47 := GenS5{F0: 0, F1: -32768}
    var v48 int8 = 100
    v1 = (v1 - v1)
    v48 = 1
    var v49 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
