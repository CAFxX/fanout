package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0(p0 int16) bool {
    if (65535 > 0) {
        fmt.Println(p0)
    } else {
        switch "hello" {
        case "a":
            for fi4 := 0; fi4 < 13; fi4++ {
                var v5 int16 = p0
                p0 = (v5 & 42)
                fmt.Println(fi4)
                _ = fi4
                _ = v5
            }
        default:
            p0 = (^(p0))
        }
    }
    if true {
        fmt.Println(p0)
        p0 = 887
    }
    switch "\n\t" {
    case "\n\t":
        p0 -= -1
    }
    v6 := []int{0}
    _ = p0
    _ = v6
    return true
}
func genFn1(p0 int32, p1 int16, p2 uint16) int {
    v7 := GenS1{F0: "a", F1: 127}
    v8 := []int{128, -736}
    if false {
        var v9 bool = true
        switch "\n\t" {
        case "\n\t":
            p2 = 671
        case "hello":
            var v10 int8 = -95
            _ = v10
        }
        if len(v8) > 0 {
            v8[0] = 415
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]string)
    var v3 uint64 = 18446744073709551615
    var v11 int64 = (v1 - v1)
    if false {
        v12 := []int{genFn1(42, 0, 721), 9223372036854775807, 4294967295, 748}
        _ = v12
    }
    var v13 int32 = 0
    v13 = (v13 & (v13 % (((v13 ^ (v13 << ((420) & 63)))) | 1)))
    var v14 int32 = 2147483647
    for fi15 := 0; fi15 < 1; fi15++ {
        switch "" {
        case "\n\t":
            v16 := GenS1{F0: ("" + "\n\t"), F1: 0}
            _ = v16
        case "a":
            fmt.Println(v13, v14, v13)
        case "":
            v11 = v11
        default:
            var v17 uint = 226
            _ = v17
        }
        if true {
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 uint = 65535
                var v20 uint = v19
                v21 := GenS2{F0: -927, F1: v11}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        } else {
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint = 1
                v13 -= v13
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                fmt.Println(fi15, v11)
                fmt.Println(v3, v3)
                _ = fi24
            }
            var v25 uint = 18446744073709551615
            _ = v25
        }
        fmt.Println(v13, fi15)
        _ = fi15
    }
    v26 := GenS2{F0: v1, F1: v1}
    var v27 bool = true
    for fi28 := 0; fi28 < 11; fi28++ {
        if v27 {
            v1 += (v11 << ((18446744073709551615) & 63))
        } else {
            if ((v27 || true) && v27) {
                var v29 uint = 2
                var v30 int = fi28
                _ = v29
                _ = v30
            } else {
                var v31 float32 = -2.25
                v32 := GenS1{F0: "", F1: 255}
                v1 = v1
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                switch "hello" {
                case "foo bar":
                    if v27 {
                        v27 = (2147483648 >= 2147483648)
                        v3 += (v3 >> ((2) & 63))
                    } else {
                        v34 := GenS1{F0: "\n\t", F1: 1}
                        _ = v34
                    }
                case "\n\t":
                    v27 = false
                case "":
                    v35 := []int{fi33, (fi33 | fi33), genFn1((v14 * (v14 * (v14 * (v13 * (v13 - (^(v14))))))), 1000, 255), (fi28 << ((4294967295) & 63))}
                    _ = v35
                }
                v14 = v13
                v36 := []int{fi28, fi28, 4294967295, fi28}
                fmt.Println(fi28, v27, v1)
                _ = fi33
                _ = v36
            }
        }
        _ = fi28
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v27 = (((-849 <= (v14 / ((-756) | 1))) || (0 > 128)) && v27)
        _ = fi37
    }
    v38 := GenS3{F0: 32767}
    switch 554 {
    case 127:
        v11 -= (v1 / (((v11 - v11)) | 1))
    default:
        if ((v14 * (v13 >> ((789) & 63))) < v13) {
            switch ("" + ("foo bar" + ("foo bar" + "\n\t"))) {
            case "\n\t":
                delete(v2, "hello")
            case "hello":
                fmt.Println(v27, v27, v3)
            case "foo bar":
                v39 := GenS4{F0: v27, F1: 195, F2: v1}
                _ = v39
            }
        }
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v38
}
