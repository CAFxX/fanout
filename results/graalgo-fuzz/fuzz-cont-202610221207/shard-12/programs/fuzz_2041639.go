package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 bool; F1 bool }

func main() {
    v1 := []string{("foo bar" + "a"), "a"}
    var v2 int16 = 1000
    v3 := []int{547, 32768}
    v4 := GenS1{F0: (!(926 > 18446744073709551615)), F1: 771, F2: -9223372036854775808}
    v2 = v2
    switch 412 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 11; fi5++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                switch (fi5 * 261) {
                case 127:
                    v2 = v2
                }
            } else {
                v1 = append(v1, ("a" + "a"))
                var v6 uint8 = 42
                v7 := GenS2{F0: 2}
                _ = v6
                _ = v7
            }
            if (!false) {
                v2 -= (v2 % (((v2 - v2)) | 1))
                v8 := GenS1{F0: (fi5 != fi5), F1: 0, F2: int64(fi5)}
                _ = v8
            }
            _ = fi5
        }
    case 32767:
        v9 := GenS1{F0: ((-2.25 * 0.1) > 1e10), F1: -2147483648, F2: 1}
        _ = v9
    }
    var v10 float64 = 1e10
    _ = len(v3)
    var v11 int = -2147483648
    if (734 >= -1) {
        for fi12 := 0; fi12 < 6; fi12++ {
            v10 = (v10 + 3.14159)
            for fi13 := 0; fi13 < 12; fi13++ {
                if (107 < 0) {
                    if ((274 == 2) || (301 <= 165)) {
                        var v14 uint64 = 541
                        var v15 int32 = -586
                        _ = v14
                        _ = v15
                    } else {
                        fmt.Println(fi12)
                        v16 := GenS1{F0: true, F1: 74, F2: -483}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v16
                    }
                }
                v2 += v2
                switch fi13 {
                case 42:
                    var v17 int = v11
                    _ = v17
                case 2147483647:
                    v18 := GenS3{F0: 798, F1: 255}
                    _ = v18
                case 4294967295:
                    v2 = v2
                default:
                    if len(v3) > 0 {
                        v3[0] = 128
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
    } else {
        v19 := GenS4{F0: (2147483648 == 1), F1: (v10 <= 0.1)}
        _ = v19
    }
    var v20 int32 = 152
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
