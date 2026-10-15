package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int; F2 int16 }

func main() {
    v1 := []int{-516, -9223372036854775808}
    var v2 string = "\n\t"
    v3 := []uint32{658, 877, 2}
    var v4 int = 2147483647
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: "a", F1: 192}
    if (!(-1 >= -682)) {
        var v8 bool = true
        switch v2 {
        case "\n\t":
            v8 = v8
        case "a":
            v9 := v5[v2]
            _ = v9
        case "":
            _ = len(v6)
        default:
            fmt.Println(v4, v4)
        }
        _ = v8
    }
    if false {
        v10 := GenS1{F0: v2, F1: 2147483647}
        fmt.Println(v2)
        _ = v10
    } else {
        if (-2147483648 > 1) {
            v2 = v2
            if true {
                var v11 uint32 = 484
                switch "a" {
                case "foo bar":
                    var v12 bool = (v2 == v2)
                    _ = v12
                }
                _ = len(v5)
                _ = v11
            } else {
                var v13 uint16 = 200
                delete(v6, "hello")
                v4 += v4
                _ = v13
            }
            switch "hello" {
            case "\n\t":
                v14 := GenS2{F0: -9223372036854775808, F1: (v4 | (v4 ^ v4))}
                _ = v14
            case "hello":
                v15 := GenS3{F0: 1000, F1: 103}
                _ = v15
            case "foo bar":
                var v16 int32 = -2147483648
                _ = v16
            }
        } else {
            v2 = v2
            v4 -= v4
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            v2 = v2
            for fi18 := 0; fi18 < 1; fi18++ {
                if (true || (32767 != 693)) {
                    var v19 float64 = -2.25
                    v20 := v6["foo bar"]
                    var v21 int = 32767
                    _ = v19
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi18, fi18)
                _ = fi18
            }
            _ = fi17
        }
        delete(v6, (v2 + v2))
    }
    var v22 int16 = 1
    fmt.Println(v2, v2, v4)
    if (!true) {
        var v23 int16 = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 14; fi24++ {
            if len(v3) > 0 {
                v3[0] = 818
            }
            switch -288 {
            case 127:
                v22 += (v22 / ((int16(fi24)) | 1))
            case 9223372036854775807:
                if (true && (!(!false))) {
                    var v25 uint64 = 0
                    _ = v25
                } else {
                    v22 -= 0
                    fmt.Println(v2, fi24, v4)
                    delete(v5, (v2 + v2))
                }
            case 4294967295:
                for fi26 := 0; fi26 < 11; fi26++ {
                    v4 += (fi26 % (((fi26 ^ (fi26 * fi24))) | 1))
                    var v27 uint16 = 787
                    _ = len(v5)
                    _ = fi26
                    _ = v27
                }
            default:
                fmt.Println(fi24, v4)
            }
            _ = fi24
        }
    }
    v22 -= v22
    var v28 uint64 = 18446744073709551615
    fmt.Println(v22, v22, v28)
    switch (v2 + v2) {
    case "hello":
        var v29 float64 = ((0.0 + 0.1) + -2.25)
        _ = v29
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        v30 := GenS4{F0: -33}
        _ = v30
    }
    v4 -= (v4 & v4)
    if true {
        var v31 float64 = ((0.0 + -2.25) + 1e10)
        var v32 uint64 = v28
        for fi33 := 0; fi33 < 2; fi33++ {
            _ = len(v1)
            var v34 uint8 = 2
            v35 := GenS3{F0: 1, F1: v34}
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v31
        _ = v32
    } else {
        switch (v4 % ((v4) | 1)) {
        case 128:
            for fi36 := 0; fi36 < 19; fi36++ {
                if false {
                    v22 -= v22
                    v4 = v4
                    if (true && true) {
                        for fi37 := 0; fi37 < 9; fi37++ {
                            v28 = 935
                            v38 := GenS3{F0: (v22 + v22), F1: 71}
                            _ = fi37
                            _ = v38
                        }
                        v1 = append(v1, fi36)
                    }
                }
                _ = fi36
            }
        case -281:
            v2 = v2
        case -1:
            for fi39 := 0; fi39 < 8; fi39++ {
                v40 := GenS5{F0: ((0 < -40) && false), F1: fi39, F2: 878}
                var v41 uint8 = 1
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        v4 = 9223372036854775807
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
