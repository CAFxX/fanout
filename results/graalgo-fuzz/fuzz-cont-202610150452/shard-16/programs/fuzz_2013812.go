package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 int16, p1 uint32) uint16 {
    if (float64(float32((float64(3.14159) - (1e10 / 1.0)))) != float64(3.14159)) {
        v5 := []int{0, 1, -535}
        v6 := GenS1{F0: (^(p0)), F1: uint64(p1)}
        _ = v5
        _ = v6
    }
    if true {
        var v7 uint64 = 18446744073709551615
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(v7)
            _ = fi8
        }
        _ = v7
    } else {
        p1 = 208
        var v9 bool = true
        v10 := GenS1{F0: p0, F1: 18446744073709551615}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return 198
}
func genFn1(p0 int64, p1 int8) int32 {
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            var v12 float32 = 3.14159
            if true {
                v13 := []int{991, fi11}
                var v14 uint64 = 0
                _ = v13
                _ = v14
            }
            _ = v12
        }
        if true {
            p1 = int8(p0)
        } else {
            var v15 uint8 = 21
            _ = v15
        }
        _ = fi11
    }
    if false {
        var v16 uint32 = 371
        var v17 bool = true
        _ = v16
        _ = v17
    } else {
        fmt.Println(p1, p0)
        p1 -= (p1 * (p1 % ((p1) | 1)))
        if ((-854 > 2147483647) && (!true)) {
            var v18 uint = 42
            switch ((("hello" + (("hello" + "foo bar") + "")) + ("\n\t" + "foo bar")) + "hello") {
            case "hello":
                v19 := []int{32767}
                _ = v19
            case "\n\t":
                p1 += p1
            default:
                var v20 float64 = -2.25
                _ = v20
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS2{F0: true, F1: false, F2: "a"}
                p0 -= p0
                _ = fi21
                _ = v22
            }
            _ = v18
        }
    }
    v23 := []int{32767, -196, 2147483647, int(p0)}
    switch 0 {
    case -2147483648:
        for fi24 := 0; fi24 < 6; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                v23 = append(v23, (fi25 - -953))
                switch (^(fi24)) {
                case -352:
                    switch "\n\t" {
                    case "hello":
                        p1 = (p1 >> ((360) & 63))
                    case "":
                        p0 += p0
                    case "foo bar":
                        fmt.Println(fi25, p1)
                    }
                default:
                    v26 := GenS3{F0: 0, F1: 128, F2: 4294967295}
                    _ = v26
                }
                _ = fi25
            }
            var v27 int16 = -32768
            var v28 uint16 = genFn0(v27, 2147483648)
            _ = fi24
            _ = v27
            _ = v28
        }
    case 0:
        fmt.Println(p1, p1)
    case 2147483647:
        v29 := GenS4{F0: "a", F1: 65535, F2: 18446744073709551615}
        _ = v29
    default:
        var v30 uint64 = 42
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v23
    return -2147483648
}

func main() {
    var v1 uint64 = 2
    var v2 uint32 = 0
    v3 := make(map[int]uint32)
    var v4 uint32 = v2
    var v31 int16 = 0
    var v32 int8 = 42
    v2 -= 2147483647
    for fi33 := 0; fi33 < 10; fi33++ {
        v2 = v2
        v34 := []int{(^(fi33))}
        _ = fi33
        _ = v34
    }
    v2 += (v2 + (^(v2)))
    if (0 != 1) {
        v1 = (^(v1))
    }
    fmt.Println(v32)
    var v35 uint64 = v1
    if true {
        if (true && true) {
            switch 127 {
            case 32767:
                fmt.Println(v2)
            case -9223372036854775808:
                v36 := GenS4{F0: "\n\t", F1: 204, F2: 4294967295}
                _ = v36
            }
            v37 := []int{}
            for fi38 := 0; fi38 < 2; fi38++ {
                v32 -= (v32 - v32)
                switch (fi38 >> ((65535) & 63)) {
                case 0:
                    var v39 uint16 = 853
                    _ = v39
                case 9223372036854775807:
                    v3[(fi38 ^ (fi38 % (((-(fi38))) | 1)))] = 42
                case 128:
                    var v40 int8 = v32
                    _ = v40
                default:
                    fmt.Println(v4, v2)
                }
                _ = fi38
            }
            _ = v37
        }
    } else {
        if false {
            v35 = v1
            var v41 float32 = 1.5
            _ = v41
        }
        fmt.Println(v32, v31, v1)
    }
    var v42 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v42
}
