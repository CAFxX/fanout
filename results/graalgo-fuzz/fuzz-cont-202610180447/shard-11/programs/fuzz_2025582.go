package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 uint8; F1 uint32; F2 string }
type GenS4 struct { F0 int }

func main() {
    var v1 int32 = 42
    var v2 int8 = 76
    var v3 float64 = 3.14159
    v4 := GenS1{F0: 432, F1: v2}
    fmt.Println(v1)
    switch 32768 {
    case 9223372036854775807:
        var v5 int = -9223372036854775808
        _ = v5
    case 28:
        switch "" {
        case "hello":
            v6 := []int{2147483647}
            _ = v6
        }
    default:
        if true {
            switch 32767 {
            case 128:
                if (255 > 4294967295) {
                    var v7 int32 = v1
                    _ = v7
                } else {
                    v8 := []int{32767, -953, -2147483648, -2147483648}
                    for fi9 := 0; fi9 < 12; fi9++ {
                        for fi10 := 0; fi10 < 15; fi10++ {
                            v3 = v3
                            v11 := GenS1{F0: 32767, F1: (v2 + 51)}
                            v8 = append(v8, fi9)
                            _ = fi10
                            _ = v11
                        }
                        switch ("foo bar" + ("a" + ((("hello" + "hello") + "\n\t") + "\n\t"))) {
                        case "a":
                            switch ("a" + ("hello" + (("foo bar" + ("" + "\n\t")) + "a"))) {
                            case "a":
                                v12 := GenS1{F0: 32767, F1: (v2 * (v2 % ((v2) | 1)))}
                                _ = v12
                            case "hello":
                                var v13 uint8 = 128
                                _ = v13
                            default:
                                v1 = (v1 * v1)
                            }
                        case "foo bar":
                            v14 := GenS2{F0: 2, F1: (v1 << ((42) & 63)), F2: 32767}
                            _ = v14
                        default:
                            v15 := GenS2{F0: 18446744073709551615, F1: -939, F2: (fi9 / ((32768) | 1))}
                            _ = v15
                        }
                        _ = fi9
                    }
                    if (65535 > 799) {
                        fmt.Println(v1, v1)
                        var v16 uint64 = 419
                        _ = v16
                    } else {
                        if true {
                            if false {
                                if len(v8) > 0 {
                                    v8[0] = -9223372036854775808
                                }
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                            var v17 float32 = ((-2.25 + 0.1) * (((1.5 + 1.5) / 1.0) + 3.14159))
                            _ = v17
                        }
                        if true {
                            v1 = 0
                            if false {
                                var v18 int32 = v1
                                _ = v18
                            }
                            var v19 uint16 = 2
                            _ = v19
                        }
                    }
                    _ = v8
                }
            case 2147483647:
                v20 := GenS1{F0: 32767, F1: v2}
                _ = v20
            default:
                var v21 bool = false
                _ = v21
            }
            v3 += (1.5 / v3)
        }
    }
    fmt.Println(v2, v1)
    v1 = -558
    if false {
        if (1 <= 42) {
            for fi22 := 0; fi22 < 9; fi22++ {
                fmt.Println(fi22)
                v23 := []int{(fi22 / ((708) | 1))}
                _ = fi22
                _ = v23
            }
        }
        v24 := []int{42, -493, 127}
        _ = v24
    } else {
        v3 -= v3
        switch ((("a" + "\n\t") + ("" + "a")) + "a") {
        case "hello":
            v1 -= v1
        default:
            switch 9223372036854775807 {
            case 128:
                v3 = 3.14159
            }
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 bool = (!(true || (!false)))
            switch (fi25 % ((fi25) | 1)) {
            case 9223372036854775807:
                v26 = false
            default:
                for fi27 := 0; fi27 < 16; fi27++ {
                    switch (fi27 / ((fi25) | 1)) {
                    case -928:
                        v28 := []int{-42}
                        _ = v28
                    case 4294967295:
                        v2 = v2
                    case -9223372036854775808:
                        var v29 int = 2147483647
                        _ = v29
                    default:
                        if (2 >= 607) {
                            v1 += -350
                            var v30 float32 = (-2.25 + (float32(((3.14159 - ((v3 * float64(0.1)) * 1.5)) + 0.1)) + float32((v3 - (v3 + 0.1)))))
                            var v31 int32 = v1
                            v32 := []int{(fi27 | -595), 32767, (fi25 + 4294967295)}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                    v1 += (v1 | (v1 & 1))
                    switch fi25 {
                    case 32767:
                        v33 := []int{(fi27 + 127), fi27, (fi25 / (((fi25 & fi27)) | 1))}
                        _ = v33
                    case 179:
                        v34 := []int{fi25}
                        _ = v34
                    case 585:
                        v35 := GenS3{F0: 45, F1: 42, F2: "a"}
                        v2 -= (v2 % ((v2) | 1))
                        _ = v35
                    default:
                        switch fi27 {
                        case 40:
                            v36 := GenS2{F0: 2, F1: (v1 ^ (v1 % (((v1 ^ int32(v2))) | 1))), F2: 9223372036854775807}
                            _ = v36
                        }
                    }
                    _ = fi27
                }
            }
            v37 := []int{}
            _ = fi25
            _ = v26
            _ = v37
        }
    }
    v1 = v1
    var v38 int32 = -827
    for fi39 := 0; fi39 < 20; fi39++ {
        fmt.Println(fi39, v1)
        _ = fi39
    }
    v2 = 1
    v1 -= v1
    v40 := GenS4{F0: 128}
    var v41 float32 = (float32(v3) * ((-2.25 + 1e10) / ((0.1 / (-2.25 + (0.1 + (3.14159 - 0.0)))) - 0.0)))
    var v42 uint64 = 970
    var v43 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
