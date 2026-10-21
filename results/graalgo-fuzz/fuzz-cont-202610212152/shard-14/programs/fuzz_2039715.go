package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int64; F1 int64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    v3 := []uint32{42, 42, 949, 308}
    v1 += v1
    for fi4 := 0; fi4 < 6; fi4++ {
        v1 = 167
        switch (^(fi4)) {
        case 0:
            v1 += v1
        default:
            v1 = v1
        }
        _ = fi4
    }
    var v5 int16 = 0
    if (v1 <= v1) {
        for fi6 := 0; fi6 < 11; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v8 := GenS1{F0: 18446744073709551615, F1: 42}
                var v9 uint16 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
            fmt.Println(fi6, v5)
            _ = fi6
        }
        switch "hello" {
        case "\n\t":
            v5 -= v5
        case "a":
            var v10 float32 = 1.5
            _ = v10
        case "hello":
            var v11 int16 = v5
            _ = v11
        }
    } else {
        v1 -= 128
        v12 := GenS1{F0: 2, F1: -466}
        _ = v12
    }
    v1 -= (v1 | (v1 << ((255) & 63)))
    var v13 int64 = -9223372036854775808
    v14 := GenS1{F0: 4294967295, F1: -411}
    v13 += v13
    var v15 int16 = 32767
    v13 -= (^(v13))
    v16 := GenS2{F0: (v13 * 223), F1: 42}
    for fi17 := 0; fi17 < 9; fi17++ {
        switch "foo bar" {
        case "a":
            if true {
                if (129 < v1) {
                    fmt.Println(v13, fi17, v13)
                }
                _ = len(v2)
                var v18 int32 = 2147483647
                _ = v18
            } else {
                var v19 uint64 = 0
                _ = v19
            }
        }
        switch 4294967295 {
        case -601:
            v20 := GenS3{F0: 0, F1: 28}
            _ = v20
        case -345:
            v21 := v2[-833]
            _ = v21
        default:
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := GenS4{F0: "\n\t", F1: 32767, F2: 228}
                switch 9223372036854775807 {
                case -622:
                    v5 = (v5 % (((v5 ^ 0)) | 1))
                case 4294967295:
                    v5 -= v15
                }
                var v24 int = 518
                v15 += (v15 | (v5 - v15))
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        if (v1 < 255) {
            var v25 int = fi17
            v26 := GenS5{F0: (v13 * v13), F1: (v13 - 42)}
            _ = v25
            _ = v26
        } else {
            if true {
                for fi27 := 0; fi27 < 3; fi27++ {
                    fmt.Println(v5, fi17, fi17)
                    _ = fi27
                }
                fmt.Println(v1, v1)
            }
        }
        _ = fi17
    }
    switch "a" {
    case "hello":
        v15 += v5
    case "foo bar":
        v15 += v15
    case "\n\t":
        for fi28 := 0; fi28 < 7; fi28++ {
            switch "\n\t" {
            case "foo bar":
                v1 = 255
            }
            _ = fi28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
