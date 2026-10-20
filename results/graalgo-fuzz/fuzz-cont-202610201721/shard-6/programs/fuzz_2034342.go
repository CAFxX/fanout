package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 uint64 }
type GenS4 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    var v3 int = -243
    var v4 string = ("" + "\n\t")
    v3 -= 42
    var v5 bool = (((-1 < -2147483648) || true) && (1.5 != 1e10))
    v6 := GenS1{F0: (v3 + 588), F1: -32768}
    v7 := GenS1{F0: 32767, F1: 0}
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint16 = 65535
        var v10 uint16 = v9
        v11 := v1[fi8]
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v4 = ""
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v2, fi12, v4)
        _ = fi12
    }
    var v13 float64 = 1.5
    var v14 int16 = -32768
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            if true {
                if ((!(fi16 >= 457)) && (!v5)) {
                    v13 -= v13
                    v17 := []int{0}
                    _ = v17
                } else {
                    var v18 bool = true
                    fmt.Println(v14)
                    _ = v18
                }
                v5 = (68 > 128)
                var v19 int = fi15
                _ = v19
            } else {
                if (!((156 == 0) && false)) {
                    v20 := GenS2{F0: 0, F1: (v14 >> ((18446744073709551615) & 63)), F2: 127}
                    _ = v20
                } else {
                    var v21 uint64 = 2
                    _ = v21
                }
                var v22 int64 = -1
                var v23 int16 = -32768
                v24 := []int{v3, 726, (v3 | (fi16 >> ((1) & 63))), (fi15 - 128)}
                v25 := GenS3{F0: 611, F1: 276, F2: 251}
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            }
            if (42 <= 1) {
                switch v4 {
                case "":
                    v1[559] = v5
                default:
                    fmt.Println(v14, v5)
                }
                fmt.Println(v3, v3, fi16)
                v2 = (v2 - v2)
                fmt.Println(v5, fi16)
            }
            _ = fi16
        }
        v1[(v3 % ((fi15) | 1))] = (42 == 18446744073709551615)
        _ = fi15
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if v5 {
            v14 -= v14
            var v27 int32 = 2147483647
            v28 := GenS3{F0: 255, F1: fi26, F2: 42}
            _ = v27
            _ = v28
        } else {
            fmt.Println(v4, v3)
            v29 := v1[(fi26 + -9223372036854775808)]
            _ = v29
        }
        _ = fi26
    }
    var v30 uint64 = 488
    v2 += v2
    switch -2147483648 {
    case -723:
        switch int(v30) {
        case 914:
            v31 := []int{}
            _ = v31
        case 42:
            var v32 uint64 = (^(v30))
            _ = v32
        case 2147483647:
            delete(v1, 9223372036854775807)
        default:
            if v5 {
                if v5 {
                    v4 = "\n\t"
                }
                v33 := GenS2{F0: 224, F1: 283, F2: 44}
                var v34 int32 = int32(v3)
                _ = v33
                _ = v34
            }
        }
    case -391:
        if true {
            var v35 int8 = -111
            _ = v35
        } else {
            var v36 uint8 = 127
            switch ((v4 + (v4 + v4)) + ("a" + (v4 + ("a" + (v4 + v4))))) {
            case "\n\t":
                v37 := GenS4{F0: v4, F1: 127}
                _ = v37
            case "a":
                v38 := GenS3{F0: 65535, F1: (v3 % ((int(v30)) | 1)), F2: (v30 << ((uint(v36)) & 63))}
                _ = v38
            case "hello":
                var v39 float64 = v13
                _ = v39
            }
            var v40 uint8 = v36
            _ = v36
            _ = v40
        }
    }
    var v41 bool = v5
    fmt.Println(v13, v2)
    v1[(v3 | 1)] = v41
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
