package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint; F2 int16 }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 uint32, p1 uint16, p2 int8) uint64 {
    fmt.Println(p2)
    switch 32768 {
    case 2147483647:
        p2 = (p2 ^ p2)
    case -323:
        v4 := []int{32767}
        _ = v4
    case -1:
        v5 := []int{-1, -1}
        _ = v5
    default:
        fmt.Println(p2, p0)
    }
    var v6 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 uint64) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        if true {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: -141, F1: (fi7 | fi7)}
                switch -9223372036854775808 {
                case -2147483648:
                    var v10 uint8 = 0
                    _ = v10
                case 9223372036854775807:
                    var v11 uint16 = 1
                    _ = v11
                }
                v12 := GenS2{F0: 848}
                v13 := GenS3{F0: 42, F1: -400, F2: -9223372036854775808}
                if (0 != 1) {
                    v14 := []int{}
                    _ = v14
                }
                v15 := GenS4{F0: ((float32(float64(3.14159)) == -2.25) || ((714 != 65535) || ("foo bar" >= "")))}
                _ = fi8
                _ = v12
                _ = v13
                _ = v15
                _ = v9
            }
            switch (-(fi7)) {
            case 145:
                p0 = (^(p0))
            case 32767:
                if (true || false) {
                    p0 = p0
                    for fi16 := 0; fi16 < 12; fi16++ {
                        p0 += 662
                        if (true || true) {
                            fmt.Println(fi16, p0, p0)
                            v17 := GenS5{F0: -90}
                            v18 := []int{fi16, fi7, fi7, fi16}
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS5{F0: 0}
                            fmt.Println(p0)
                            var v20 int32 = 2147483647
                            _ = v19
                            _ = v20
                        }
                        switch "foo bar" {
                        case "":
                            var v21 uint16 = 42
                            _ = v21
                        case "hello":
                            switch (fi7 >> ((2) & 63)) {
                            case -668:
                                var v22 uint64 = (p0 ^ (p0 << ((42) & 63)))
                                _ = v22
                            case 32768:
                                fmt.Println(fi7)
                                p0 = p0
                            }
                        }
                        _ = fi16
                    }
                }
            default:
                var v23 int64 = -652
                _ = v23
            }
            var v24 int = fi7
            _ = v24
        }
        _ = fi7
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := GenS1{F0: 1, F1: fi25}
        _ = fi25
        _ = v26
    }
    switch 128 {
    case 0:
        if (2 <= 102) {
            p0 += p0
            var v27 string = "\n\t"
            _ = v27
        }
    default:
        p0 += p0
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 54
    v3 := make(map[int]bool)
    v2 = (-(v2))
    if (float32(((float64(float32(3.14159)) - 0.1) * float64(float32(-2.25)))) < (float32(float64(1e10)) / 3.14159)) {
        v28 := GenS6{F0: "foo bar", F1: 255, F2: -1}
        _ = v28
    }
    v2 = (-(v2))
    v2 = (v2 >> ((65535) & 63))
    v3[448] = (false || false)
    if (1 <= -2147483648) {
        if true {
            v2 -= v2
        } else {
            v1 += 32767
        }
        if true {
            v1 = (v1 / ((v1) | 1))
            v1 += (v1 & 1)
        } else {
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 int32 = 0
                var v31 int8 = 2
                var v32 string = "hello"
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            fmt.Println(v2)
            v1 -= v1
        }
        fmt.Println(v2, v1)
    } else {
        v33 := []int{42, 274, -9223372036854775808, 9223372036854775807}
        v34 := GenS3{F0: 441, F1: -9223372036854775808, F2: -1}
        var v35 uint = 1
        _ = v33
        _ = v34
        _ = v35
    }
    if false {
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 uint = uint(v1)
            v38 := []int{(fi36 - (fi36 + fi36)), 32767, fi36}
            v2 -= 38
            _ = fi36
            _ = v37
            _ = v38
        }
    } else {
        v39 := GenS6{F0: ("hello" + "foo bar"), F1: 65535, F2: 1000}
        _ = v39
    }
    v40 := GenS7{F0: -1, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v40
}
