package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 56
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int32 = -2147483648
        var v7 uint32 = 2147483648
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 += (^(v2))
        switch (fi8 >> ((v2) & 63)) {
        case -683:
            switch (fi8 * fi8) {
            case 4294967295:
                v9 := GenS1{F0: 42, F1: 97, F2: 9223372036854775807}
                _ = v9
            case -1:
                v2 += v2
            default:
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(v2, fi8, v2)
                    v11 := []int{fi10, fi8, (fi8 << ((726) & 63)), 2147483647}
                    _ = fi10
                    _ = v11
                }
            }
        case 32768:
            if (!(1 >= 2147483647)) {
                v12 := []int{32768}
                _ = v12
            } else {
                var v13 int64 = 9223372036854775807
                fmt.Println(v2, v2, v13)
                _ = v13
            }
        }
        _ = fi8
    }
    v2 = (v2 >> ((v2) & 63))
    fmt.Println(v2)
    switch 2147483647 {
    case 32767:
        var v14 uint64 = 797
        _ = v14
    case 42:
        var v15 int32 = 907
        _ = v15
    case -2147483648:
        v16 := GenS1{F0: 18446744073709551615, F1: -76, F2: -1}
        _ = v16
    default:
        v17 := GenS1{F0: 306, F1: 1, F2: -9223372036854775808}
        _ = v17
    }
    var v18 int32 = 1
    v2 -= (v2 << ((2) & 63))
    var v19 int16 = 1
    if (-128 >= 42) {
        if ((!(3.14159 == 0.1)) || (-2.25 == float64(1e10))) {
            v19 = v19
            var v20 uint = 210
            v18 -= (v18 * (v18 - -2147483648))
            _ = v20
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int32 = v18
                if false {
                    var v23 bool = (0 <= 127)
                    var v24 uint64 = 23
                    v25 := v3[1]
                    _ = v23
                    _ = v24
                    _ = v25
                }
                if true {
                    v26 := []int{42, int(v22)}
                    if (!false) {
                        v27 := GenS1{F0: 2, F1: v19, F2: 1}
                        v22 = v18
                        _ = v27
                    }
                    switch ("a" + "") {
                    case "foo bar":
                        switch "foo bar" {
                        case "a":
                            v19 -= v19
                        case "hello":
                            v28 := v4[9223372036854775807]
                            _ = v28
                        default:
                            v2 -= v2
                        }
                    case "a":
                        if (653 <= 65535) {
                            var v29 string = "hello"
                            _ = v29
                        }
                    case "":
                        v19 = (v19 / ((1) | 1))
                    default:
                        v22 += v18
                    }
                    _ = v26
                } else {
                    var v30 float32 = 1e10
                    var v31 float32 = v30
                    switch fi21 {
                    case 42:
                        v22 = (-(v18))
                    default:
                        v32 := GenS1{F0: 18446744073709551615, F1: 42, F2: 1}
                        _ = v32
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi21
                _ = v22
            }
        }
        if true {
            v33 := []int{765, -290, 2147483647}
            _ = v33
        } else {
            v34 := []int{}
            _ = v34
        }
        v19 += (v19 / (((v19 & (v19 * v19))) | 1))
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v2 += 42
        _ = fi35
    }
    v2 = 65535
    if (1 >= 1) {
        for fi36 := 0; fi36 < 2; fi36++ {
            v4[fi36] = 904
            v37 := GenS1{F0: 18446744073709551615, F1: 1, F2: -1}
            _ = fi36
            _ = v37
        }
        fmt.Println(v2, v19)
    } else {
        _ = len(v4)
        v38 := GenS2{F0: (("hello" + "\n\t") + "\n\t")}
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
