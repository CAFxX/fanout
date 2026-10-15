package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint8 = 121
    var v2 float64 = 0.0
    var v3 uint16 = 255
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint64 = 2
        fmt.Println(v5)
        _ = fi4
        _ = v5
    }
    if false {
        if (-1 > 42) {
            v2 -= 3.14159
            var v6 bool = true
            _ = v6
        }
        v7 := []int{0, 32768, -246}
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                if (224 < 42) {
                    v10 := GenS1{F0: 18446744073709551615, F1: 42}
                    _ = v10
                } else {
                    v11 := GenS2{F0: -88, F1: 178}
                    var v12 int8 = 127
                    if (((true || false) || ("foo bar" <= "")) || (!(v3 > (v3 * 65535)))) {
                        var v13 uint = 18446744073709551615
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v7 = append(v7, fi8)
                        _ = v13
                    }
                    _ = v11
                    _ = v12
                }
                if false {
                    v3 -= v3
                }
                _ = fi9
            }
            v14 := GenS1{F0: 65535, F1: 42}
            switch "\n\t" {
            case "a":
                v15 := GenS1{F0: 1, F1: -12}
                _ = v15
            case "\n\t":
                _ = len(v7)
            }
            _ = fi8
            _ = v14
        }
        _ = v7
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v3 += (^(v3))
            if (!((!true) && (18446744073709551615 > 2))) {
                var v18 bool = (((18446744073709551615 > 530) && true) || false)
                _ = v18
            } else {
                v3 = v3
                if false {
                    var v19 int64 = -9223372036854775808
                    _ = v19
                } else {
                    var v20 bool = (18446744073709551615 == 2)
                    _ = v20
                }
            }
            var v21 int64 = 42
            _ = fi17
            _ = v21
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            switch (fi22 / (((^(fi22))) | 1)) {
            case 408:
                v2 = (((-2.25 / 1.5) / v2) - v2)
            case 0:
                v23 := GenS3{F0: fi16, F1: 2}
                _ = v23
            case 4294967295:
                v2 -= 0.0
            default:
                v2 += v2
            }
            _ = fi22
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 uint32 = 0
            if ((!true) && (519 >= 0)) {
                v26 := []int{(fi24 / (((fi24 << ((uint(v3)) & 63))) | 1)), fi16}
                fmt.Println(v25)
                v3 -= (^(v3))
                _ = v26
            }
            var v27 int32 = 572
            _ = fi24
            _ = v25
            _ = v27
        }
        _ = fi16
    }
    v2 = (v2 + v2)
    if true {
        var v28 uint64 = 1
        if false {
            switch 32767 {
            case -208:
                v3 += 69
            case 9223372036854775807:
                v29 := []int{-872, 976, 2147483647}
                _ = v29
            case 2147483647:
                v30 := GenS3{F0: -722, F1: 97}
                _ = v30
            default:
                switch 42 {
                case 128:
                    fmt.Println(v1, v28, v28)
                }
            }
            fmt.Println(v3, v28, v28)
        }
        _ = v28
    } else {
        v31 := GenS1{F0: 4294967295, F1: 32767}
        var v32 uint16 = v3
        for fi33 := 0; fi33 < 20; fi33++ {
            if true {
                var v34 uint16 = v32
                _ = v34
            } else {
                v35 := GenS3{F0: fi33, F1: 856}
                if (!true) {
                    var v36 float32 = ((-2.25 - 0.1) - float32(float64(1e10)))
                    for fi37 := 0; fi37 < 2; fi37++ {
                        switch ("hello" + "\n\t") {
                        case "":
                            if (!((v36 - v36) < 0.1)) {
                                v38 := []int{}
                                var v39 uint64 = 962
                                _ = v38
                                _ = v39
                            } else {
                                fmt.Println(v36, fi37)
                            }
                        case "\n\t":
                            fmt.Println(v3, fi33)
                        case "hello":
                            v36 += ((float32(v2) + 0.0) * ((3.14159 + 0.0) * (float32(v2) - v36)))
                        }
                        v1 = 64
                        var v40 int64 = 42
                        _ = fi37
                        _ = v40
                    }
                    _ = v36
                }
                _ = v35
            }
            _ = fi33
        }
        _ = v31
        _ = v32
    }
    v41 := []int{-200, 933, 1, 42}
    var v42 bool = (0 != -439)
    v43 := GenS1{F0: 924, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
