package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int8) int8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := []int{(fi4 >> ((2) & 63)), 9223372036854775807, (fi4 ^ 1), fi4}
        _ = fi4
        _ = v5
    }
    var v6 int32 = -551
    _ = p0
    _ = p1
    _ = v6
    return (p1 + (^(p1)))
}

func main() {
    var v1 int32 = 111
    v2 := make(map[uint32]uint32)
    v3 := []int{-9223372036854775808, 9223372036854775807, 0}
    v1 -= v1
    fmt.Println(v1)
    v1 += (-(v1))
    if false {
        if (true && (false && true)) {
            for fi7 := 0; fi7 < 14; fi7++ {
                v1 = v1
                var v8 float32 = float32(1.5)
                _ = fi7
                _ = v8
            }
            if (!(((18446744073709551615 == 0) && true) && true)) {
                v9 := v2[2147483647]
                _ = v9
            } else {
                for fi10 := 0; fi10 < 19; fi10++ {
                    var v11 int32 = (v1 ^ v1)
                    _ = fi10
                    _ = v11
                }
            }
        } else {
            switch 295 {
            case 128:
                if ((1.5 * -2.25) < 1e10) {
                    v1 = v1
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 int32 = (v1 * (-(v1)))
                        _ = fi12
                        _ = v13
                    }
                }
            }
            _ = len(v3)
            if true {
                v14 := v2[2147483647]
                fmt.Println(v1, v14)
                if true {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        fmt.Println(v1)
                        if (0 == fi15) {
                            v1 += (v1 + 339)
                            v16 := v2[(v14 & (v14 | v14))]
                            _ = v16
                        }
                        if (!false) {
                            var v17 uint32 = (v14 - v14)
                            var v18 uint = 18446744073709551615
                            v18 += 42
                            _ = v17
                            _ = v18
                        } else {
                            v1 = (v1 % ((-1) | 1))
                        }
                        _ = fi15
                    }
                    fmt.Println(v1, v14)
                }
                _ = v14
            }
        }
        v1 = -1
        v19 := GenS1{F0: 127, F1: true, F2: 42}
        _ = v19
    } else {
        if (false && (0 != -32768)) {
            v1 += v1
            if (42 > 0) {
                if false {
                    var v20 float32 = 1.5
                    _ = v20
                }
                v3 = append(v3, 4294967295)
                fmt.Println(v1)
            } else {
                var v21 uint64 = 293
                v1 += (v1 / ((v1) | 1))
                _ = v21
            }
        }
        delete(v2, 65535)
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 string = "foo bar"
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 bool = true
        var v26 int8 = 28
        _ = fi24
        _ = v25
        _ = v26
    }
    switch "hello" {
    case "hello":
        v1 = v1
    case "":
        v1 += (^(v1))
    case "\n\t":
        fmt.Println(v1)
    }
    var v27 int = 404
    v28 := GenS2{F0: 17, F1: -332}
    if (100 != 0) {
        var v29 bool = false
        v30 := GenS1{F0: v27, F1: (!v29), F2: 2}
        var v31 uint16 = 0
        _ = v29
        _ = v30
        _ = v31
    }
    v1 += (^(v1))
    for fi32 := 0; fi32 < 7; fi32++ {
        if (-9223372036854775808 == 9223372036854775807) {
            v33 := GenS3{F0: 156}
            _ = v33
        }
        _ = fi32
    }
    if (4294967295 <= 65535) {
        var v34 uint8 = 139
        _ = v34
    } else {
        switch -775 {
        case 754:
            var v35 uint16 = 255
            _ = v35
        case -2147483648:
            v36 := GenS1{F0: (v27 + -566), F1: true, F2: 128}
            _ = v36
        }
    }
    fmt.Println(v27, v27)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
