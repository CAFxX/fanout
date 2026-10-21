package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 uint32, p1 uint32, p2 uint64) uint {
    v7 := []int{9223372036854775807, -1, -9223372036854775808, 128}
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            var v9 int32 = 2147483647
            p0 += p1
            if (int8(v9) >= 0) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    if (1 != 255) {
                        var v11 float32 = 0.1
                        _ = v11
                    }
                    _ = fi10
                }
                for fi12 := 0; fi12 < 20; fi12++ {
                    if len(v7) > 0 {
                        v7[0] = (^(fi8))
                    }
                    v13 := GenS1{F0: 1, F1: (p2 ^ (p2 + p2)), F2: (p2 * p2)}
                    _ = fi12
                    _ = v13
                }
            }
            _ = v9
        } else {
            if false {
                if true {
                    v7 = append(v7, 32767)
                }
            }
            v14 := GenS1{F0: 255, F1: 2, F2: p2}
            var v15 int16 = -32768
            _ = v14
            _ = v15
        }
        if false {
            if (true || false) {
                var v16 bool = (-657 < 1)
                if v16 {
                    v16 = ((p2 * 387) != 2)
                    var v17 uint32 = p1
                    v17 += (v17 + (p0 >> ((1) & 63)))
                    _ = v17
                }
                if v16 {
                    switch "\n\t" {
                    case "foo bar":
                        var v18 int64 = 9223372036854775807
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v20 := GenS1{F0: 0, F1: 4294967295, F2: p2}
                        if ((fi8 << ((4294967295) & 63)) > fi8) {
                            var v21 int64 = 113
                            _ = v21
                        }
                        _ = fi19
                        _ = v20
                    }
                } else {
                    p0 += 0
                }
                _ = v16
            }
            p2 = (^(p2))
        } else {
            if true {
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 uint64 = 969
                    _ = fi22
                    _ = v23
                }
            }
            fmt.Println(p0)
        }
        _ = fi8
    }
    if false {
        switch "" {
        case "":
            var v24 int = 9223372036854775807
            _ = v24
        case "foo bar":
            _ = len(v7)
        case "\n\t":
            p2 = p2
        }
        var v25 string = ("a" + "hello")
        _ = v25
    }
    var v26 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v7
    return 255
}

func main() {
    var v1 int = -113
    v2 := make(map[int]uint32)
    var v3 string = ("\n\t" + "")
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    var v6 int64 = 1
    v1 = 2147483647
    if (65535 != 42) {
        v27 := []int{int(v6)}
        var v28 int16 = 0
        _ = v27
        _ = v28
    } else {
        v1 = v1
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v2[fi29] = 4294967295
        _ = fi29
    }
    if (255 <= 1) {
        v30 := GenS1{F0: 0, F1: 134, F2: 18446744073709551615}
        _ = v30
    }
    switch (v3 + v3) {
    case "hello":
        if (int16(v1) >= 1000) {
            var v31 int64 = v6
            var v32 float64 = -2.25
            _ = v31
            _ = v32
        } else {
            fmt.Println(v6, v1)
            delete(v2, 0)
        }
    case "\n\t":
        v3 = "hello"
    case "":
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := GenS2{F0: 42, F1: 0}
            _ = fi33
            _ = v34
        }
    default:
        var v35 string = (v3 + v3)
        _ = v35
    }
    v1 = 9223372036854775807
    if (true || (v3 > (("foo bar" + v3) + v3))) {
        switch (v1 << ((18446744073709551615) & 63)) {
        case 0:
            v36 := []int{9223372036854775807, (v1 ^ v1), (v1 | v1), v1}
            _ = v36
        case 47:
            v37 := GenS2{F0: 884, F1: 41}
            _ = v37
        case 32768:
            fmt.Println(v1, v6, v3)
        default:
            v38 := GenS3{F0: v6, F1: 127, F2: -33}
            _ = v38
        }
        var v39 float64 = 1e10
        var v40 int16 = 465
        _ = v39
        _ = v40
    }
    delete(v4, v1)
    var v41 uint32 = 42
    fmt.Println(v41, v3)
    var v42 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
