package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int16; F2 int64 }
func genFn0(p0 int16, p1 uint32) int32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 = p0
            var v9 bool = true
            _ = fi8
            _ = v9
        }
        p0 -= 42
        _ = fi7
    }
    v10 := []int{1}
    p1 += p1
    p1 += (p1 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = v10
    return -958
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 int = 127
    v4 := make(map[uint32]string)
    var v5 uint32 = 2
    v6 := []int{4294967295, (v3 % ((v3) | 1)), (v3 * -2147483648)}
    v3 = v3
    if (32767 > 1000) {
        var v11 int32 = 42
        _ = v11
    }
    if true {
        if len(v6) > 0 {
            v6[0] = v3
        }
        if false {
            if len(v6) > 0 {
                v6[0] = -2147483648
            }
            switch v3 {
            case -342:
                var v12 float32 = (0.1 * 1e10)
                _ = v12
            case 32767:
                if false {
                    v13 := GenS1{F0: 2, F1: "foo bar", F2: 455}
                    _ = v13
                } else {
                    var v14 uint = 493
                    v15 := GenS1{F0: 2, F1: (("\n\t" + ("\n\t" + "a")) + "foo bar"), F2: 2}
                    _ = v14
                    _ = v15
                }
            case 42:
                if false {
                    if (!(951 <= 470)) {
                        v5 = 2147483647
                    }
                }
            }
            fmt.Println(v3)
        }
    } else {
        v3 -= (v3 + v3)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(fi16, v3)
        for fi17 := 0; fi17 < 8; fi17++ {
            v5 = v5
            v18 := GenS2{F0: v5}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v3 -= v3
    if true {
        v19 := v2[(v3 % (((v3 << ((672) & 63))) | 1))]
        fmt.Println(v5, v5, v5)
        v20 := GenS2{F0: v5}
        _ = v19
        _ = v20
    }
    v5 = (v5 / ((v5) | 1))
    var v21 uint64 = 1
    fmt.Println(v5, v5)
    for fi22 := 0; fi22 < 16; fi22++ {
        v5 = 42
        switch v3 {
        case -2147483648:
            v23 := GenS1{F0: 42, F1: (((("a" + "\n\t") + ("" + ("" + "hello"))) + "a") + "a"), F2: (v21 << ((4294967295) & 63))}
            _ = v23
        case 1:
            var v24 uint = 0
            _ = v24
        case 9223372036854775807:
            var v25 int16 = 1000
            _ = v25
        }
        switch "" {
        case "foo bar":
            fmt.Println(v5, fi22, v3)
        default:
            var v26 int = (v3 | 9223372036854775807)
            _ = v26
        }
        _ = fi22
    }
    v27 := GenS1{F0: 286, F1: "\n\t", F2: 18446744073709551615}
    for fi28 := 0; fi28 < 3; fi28++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v5 = v5
        _ = fi28
    }
    v3 -= v3
    fmt.Println(v3, v5)
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            v31 := v2[1]
            if (true || false) {
                if (true || (int8(v3) < 0)) {
                    if (0 <= 2147483647) {
                        switch "a" {
                        case "\n\t":
                            _ = len(v6)
                        case "foo bar":
                            fmt.Println(v3, v31)
                        case "":
                            v31 = (v31 + v31)
                        default:
                            switch fi29 {
                            case -2147483648:
                                var v32 uint64 = v21
                                v33 := v1[(v5 / ((65535) | 1))]
                                var v34 int = (fi30 + (fi30 % ((v3) | 1)))
                                _ = v32
                                _ = v33
                                _ = v34
                            case 770:
                                _ = len(v6)
                            }
                            v5 = 2147483648
                        }
                        v5 = v5
                    }
                } else {
                    v21 = v21
                    v35 := GenS3{F0: -205, F1: 1, F2: 930}
                    _ = v35
                }
            }
            _ = fi30
            _ = v31
        }
        var v36 int32 = 2147483647
        v37 := GenS1{F0: 42, F1: "foo bar", F2: (v21 >> ((18446744073709551615) & 63))}
        _ = fi29
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
