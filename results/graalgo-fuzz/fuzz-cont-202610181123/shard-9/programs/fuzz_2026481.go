package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
func genFn0() int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, (fi7 | (fi7 / ((32767) | 1)))}
        switch "" {
        case "foo bar":
            switch (("a" + "\n\t") + ("\n\t" + "\n\t")) {
            case "":
                if false {
                    switch "hello" {
                    case "foo bar":
                        v9 := GenS1{F0: true}
                        _ = v9
                    case "\n\t":
                        v8 = append(v8, (^(fi7)))
                    }
                    switch fi7 {
                    case 128:
                        v10 := GenS1{F0: true}
                        _ = v10
                    case 604:
                        var v11 uint = 782
                        _ = v11
                    case -9223372036854775808:
                        for fi12 := 0; fi12 < 2; fi12++ {
                            if true {
                                var v13 uint32 = 2
                                switch fi12 {
                                case 1:
                                    var v14 float64 = -2.25
                                    _ = v14
                                }
                                _ = v13
                            } else {
                                var v15 uint = 456
                                _ = v15
                            }
                            _ = fi12
                        }
                    }
                    var v16 uint8 = 1
                    _ = v16
                }
            case "hello":
                if len(v8) > 0 {
                    v8[0] = -405
                }
            case "a":
                var v17 int32 = 0
                _ = v17
            default:
                var v18 uint16 = 1
                _ = v18
            }
        case "hello":
            fmt.Println(fi7)
        default:
            v19 := GenS2{F0: 703}
            _ = v19
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint64 = 560
            var v22 string = "\n\t"
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi7
        _ = v8
    }
    var v23 float64 = 1e10
    v24 := GenS2{F0: -903}
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(v23)
        v26 := []int{}
        v23 = v23
        _ = fi25
        _ = v26
    }
    _ = v23
    _ = v24
    return 42
}

func main() {
    var v1 uint64 = 497
    v2 := []string{(("a" + "a") + "a")}
    var v3 int8 = 100
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v6 string = "\n\t"
    switch v6 {
    case "a":
        v27 := GenS3{F0: 925, F1: (v3 ^ 100)}
        _ = v27
    case "hello":
        var v28 uint64 = 60
        _ = v28
    }
    switch (v6 + v6) {
    case "a":
        v3 += (v3 << ((18446744073709551615) & 63))
    default:
        var v29 float32 = float32(float64(0.0))
        _ = v29
    }
    var v30 float32 = 1.5
    var v31 uint32 = 2147483648
    v6 = v6
    if len(v2) > 0 {
        v2[0] = (v6 + ((v6 + v6) + (v6 + ((v6 + v6) + (v6 + v6)))))
    }
    if true {
        var v32 uint16 = 32767
        v31 += 64
        _ = v32
    }
    fmt.Println(v3, v6, v30)
    v31 -= 2147483647
    v30 = (((1e10 - v30) / -2.25) * 1e10)
    var v33 int = -380
    if true {
        fmt.Println(v1, v30, v6)
        var v34 uint32 = v31
        v1 = 2
        _ = v34
    }
    switch 751 {
    case 32768:
        fmt.Println(v31, v3)
    case 342:
        if (2 != (v1 + (v1 ^ v1))) {
            var v35 bool = (65535 <= 32767)
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = v35
        }
    case 459:
        v36 := GenS4{F0: 65535, F1: v1, F2: (v6 + "")}
        _ = v36
    default:
        fmt.Println(v3, v1)
    }
    var v37 uint16 = 42
    v5[128] = (0 == 42)
    switch v33 {
    case 748:
        if (!false) {
            v31 = v31
        }
    case -226:
        fmt.Println(v6, v1)
    default:
        if false {
            v33 = 32767
        }
    }
    var v38 int8 = (v3 % ((v3) | 1))
    v39 := GenS3{F0: (v31 | v31), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
