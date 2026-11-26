package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0(p0 uint16, p1 uint16) int64 {
    p1 = p0
    v4 := []int{9223372036854775807, 1, 163}
    fmt.Println(p1)
    for fi5 := 0; fi5 < 17; fi5++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(fi5, p1)
        var v6 uint16 = 1
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 uint8, p2 uint64) string {
    v7 := []int{-840}
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "hello"
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 42
    var v3 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: fi8, F1: (v1 * v1)}
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    v2 += v2
    v1 = int16(v2)
    if true {
        var v11 uint64 = 184
        if ((float32(0.0) + 0.1) == 0.0) {
            v1 = (v1 | (v1 * v1))
            var v12 uint32 = 65535
            _ = v12
        }
        _ = v11
    } else {
        switch "hello" {
        case "a":
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 float64 = 0.1
                _ = fi13
                _ = v14
            }
        case "hello":
            var v15 int32 = 520
            _ = v15
        default:
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int64 = v3
                _ = fi16
                _ = v17
            }
        }
        var v18 int8 = -60
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v2 = (v2 + (v2 % (((v2 ^ v2)) | 1)))
        v3 = (v3 / (((v3 - v3)) | 1))
        _ = fi19
    }
    v20 := GenS1{F0: 4294967295, F1: (v1 - (^(v1)))}
    for fi21 := 0; fi21 < 10; fi21++ {
        switch "" {
        case "a":
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v2)
                var v23 float32 = 0.1
                _ = fi22
                _ = v23
            }
        case "foo bar":
            v1 -= (^(v1))
        case "hello":
            v3 = 50
        }
        _ = fi21
    }
    if (0 >= 2147483648) {
        for fi24 := 0; fi24 < 12; fi24++ {
            if true {
                var v25 bool = true
                switch 0 {
                case 0:
                    v2 -= (v2 >> ((42) & 63))
                default:
                    _ = len(v10)
                }
                _ = v25
            } else {
                for fi26 := 0; fi26 < 17; fi26++ {
                    if (v1 == (v1 >> ((255) & 63))) {
                        v2 -= 255
                        var v27 float32 = 1.5
                        _ = v27
                    }
                    var v28 bool = true
                    var v29 float64 = 0.1
                    _ = fi26
                    _ = v28
                    _ = v29
                }
            }
            _ = fi24
        }
        switch (("hello" + "foo bar") + ("a" + (("a" + ("" + "a")) + ("foo bar" + ("hello" + "hello"))))) {
        case "a":
            switch "\n\t" {
            case "foo bar":
                if false {
                    _ = len(v10)
                }
            }
        case "foo bar":
            for fi30 := 0; fi30 < 19; fi30++ {
                if false {
                    v2 += v2
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := GenS1{F0: 127, F1: 42}
                        v33 := GenS1{F0: (fi31 - (fi31 / (((-(fi31))) | 1))), F1: v1}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                }
                if ((^(v2)) != (v2 << ((167) & 63))) {
                    v2 = v2
                    v34 := GenS2{F0: 127, F1: (v2 | 65535), F2: 0}
                    _ = v34
                }
                _ = fi30
            }
        case "":
            v1 = (v1 * (v1 * (v1 | (v1 / (((v1 / (((v1 + v1)) | 1))) | 1)))))
        }
        var v35 uint16 = v2
        _ = v35
    }
    v10 = append(v10, 4294967295)
    v3 = (^(v3))
    var v36 int32 = 42
    switch 32768 {
    case 32767:
        fmt.Println(v1, v2)
    case -971:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v36
}
