package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 1e10
    var v3 int64 = 42
    if (498 <= 2147483647) {
        for fi4 := 0; fi4 < 4; fi4++ {
            for fi5 := 0; fi5 < 13; fi5++ {
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi5)
                }
                _ = fi5
            }
            for fi6 := 0; fi6 < 9; fi6++ {
                switch (^(fi6)) {
                case 1:
                    if (fi6 < fi4) {
                        v7 := GenS1{F0: -8}
                        _ = v7
                    } else {
                        if ((!((!((fi6 >> ((18446744073709551615) & 63)) == (fi4 / (((^(fi4))) | 1)))) || (!true))) || false) {
                            fmt.Println(v2)
                            fmt.Println(fi6, v2)
                        } else {
                            v8 := v1[0]
                            var v9 uint = 65535
                            v10 := GenS2{F0: 2}
                            _ = v10
                            _ = v8
                            _ = v9
                        }
                        v11 := GenS2{F0: 42}
                        switch fi6 {
                        case 32768:
                            fmt.Println(fi4, v3, fi6)
                        case 9223372036854775807:
                            if false {
                                v12 := []int{fi4, (fi6 >> ((65535) & 63)), 42, 4294967295}
                                v3 += 1
                                _ = v12
                            } else {
                                var v13 uint16 = 255
                                _ = v13
                            }
                        case 32767:
                            fmt.Println(v3, fi4, v3)
                        }
                        _ = v11
                    }
                }
                _ = fi6
            }
            _ = fi4
        }
    } else {
        v3 = 297
    }
    var v14 float32 = 0.1
    v15 := GenS1{F0: -1}
    var v16 int = -1
    if false {
        var v17 int = (v16 | v16)
        for fi18 := 0; fi18 < 8; fi18++ {
            _ = len(v1)
            var v19 uint16 = 818
            var v20 uint64 = 887
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        fmt.Println(v3)
        v3 = v3
        _ = fi21
    }
    v2 += ((float64((1e10 - 1.5)) + 0.0) * 0.0)
    var v22 bool = ((v3 == (v3 ^ v3)) && false)
    var v23 int16 = 32767
    _ = len(v1)
    switch v16 {
    case 0:
        var v24 int = (v16 / ((int(v23)) | 1))
        _ = v24
    case 32767:
        v25 := GenS3{F0: 255}
        _ = v25
    case -556:
        v26 := GenS4{F0: v23, F1: 4294967295, F2: ("foo bar" + "hello")}
        _ = v26
    default:
        v23 -= v23
    }
    fmt.Println(v22, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
