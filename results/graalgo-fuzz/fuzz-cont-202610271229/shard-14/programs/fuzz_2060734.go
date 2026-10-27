package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 bool; F1 uint16; F2 int8 }
type GenS4 struct { F0 int; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint64; F1 bool; F2 int }
func genFn0(p0 int16, p1 uint8, p2 uint64) uint16 {
    var v5 uint8 = (p1 ^ 2)
    v6 := GenS1{F0: -128, F1: ("hello" + "hello"), F2: 46}
    p0 += 531
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 114
}
func genFn1(p0 uint16, p1 int) uint64 {
    v7 := []int{p1, p1, (p1 - 1), (p1 % (((p1 - (p1 << ((972) & 63)))) | 1))}
    p0 -= (p0 & p0)
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2147483647
    var v3 bool = (2147483647 <= -1)
    v4 := make(map[int]string)
    v3 = v3
    v2 = v2
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 = (v2 >> ((688) & 63))
        v1[221] = "a"
        _ = fi8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := v4[-2147483648]
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(fi9, v3)
            _ = fi11
        }
        switch v10 {
        case "":
            var v12 uint32 = (v2 + 65535)
            _ = v12
        case "foo bar":
            v10 = v10
        }
        _ = fi9
        _ = v10
    }
    if v3 {
        v13 := GenS2{F0: 2, F1: 0, F2: "a"}
        _ = v13
    } else {
        switch ("foo bar" + "a") {
        case "hello":
            for fi14 := 0; fi14 < 12; fi14++ {
                for fi15 := 0; fi15 < 17; fi15++ {
                    if (18446744073709551615 != uint64(v2)) {
                        fmt.Println(v2, v2, v3)
                        if (42 > 3) {
                            var v16 string = ("hello" + "")
                            _ = v16
                        }
                    }
                    v17 := GenS3{F0: true, F1: 0, F2: -115}
                    v18 := GenS1{F0: -75, F1: "a", F2: 0}
                    _ = fi15
                    _ = v17
                    _ = v18
                }
                v19 := GenS3{F0: (v3 && (float64(-2.25) >= float64(((1e10 + (((0.1 + 3.14159) * (0.1 * 1e10)) / 1e10)) / -2.25)))), F1: 32767, F2: -128}
                _ = fi14
                _ = v19
            }
        case "a":
            v3 = (2147483647 != -2147483648)
        }
    }
    var v20 uint8 = 119
    if v3 {
        for fi21 := 0; fi21 < 8; fi21++ {
            if v3 {
                switch 2147483647 {
                case 1:
                    switch (fi21 % ((fi21) | 1)) {
                    case -731:
                        v22 := v4[-9223372036854775808]
                        _ = v22
                    case 128:
                        v23 := GenS4{F0: (fi21 % (((fi21 / (((fi21 & fi21)) | 1))) | 1)), F1: 801, F2: 255}
                        _ = v23
                    default:
                        _ = len(v4)
                    }
                case 32768:
                    v3 = (-857 < 0)
                default:
                    v20 -= (v20 * (^(v20)))
                }
                v24 := []int{fi21, fi21, fi21}
                v2 += 1
                _ = v24
            }
            _ = fi21
        }
        v2 = 0
        v25 := GenS5{F0: 42, F1: v20}
        _ = v25
    }
    if (1e10 >= float64(0.1)) {
        if (0.1 == float64(1e10)) {
            v26 := GenS6{F0: 119, F1: (genFn1(32767, 128) >= 807), F2: -139}
            _ = v26
        }
    }
    switch "hello" {
    case "foo bar":
        v27 := []int{-1, 32767}
        _ = v27
    case "hello":
        v2 -= (^(v2))
    case "a":
        v2 = v2
    }
    v20 += 132
    v3 = v3
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(fi28, fi28, v2)
        var v29 float64 = -2.25
        _ = fi28
        _ = v29
    }
    fmt.Println(v20, v3, v20)
    if v3 {
        v3 = v3
        v30 := []int{2147483647}
        var v31 uint16 = 261
        _ = v30
        _ = v31
    } else {
        var v32 int = 32768
        switch (v32 >> ((42) & 63)) {
        case 9223372036854775807:
            var v33 uint = 0
            _ = v33
        default:
            var v34 uint8 = (v20 << ((4294967295) & 63))
            _ = v34
        }
        var v35 bool = v3
        _ = v32
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
