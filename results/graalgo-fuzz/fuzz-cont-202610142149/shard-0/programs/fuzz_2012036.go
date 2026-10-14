package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint }

func main() {
    var v1 uint64 = 598
    var v2 string = "hello"
    var v3 uint32 = 2
    v4 := make(map[uint32]int)
    var v5 uint8 = 0
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 32767}
        if (42 == 0) {
            for fi9 := 0; fi9 < 20; fi9++ {
                switch (fi7 - fi7) {
                case 4294967295:
                    _ = len(v4)
                case 836:
                    switch 42 {
                    case -2147483648:
                        if ((-2.25 + 0.1) < ((((float64(0.0) - float64((float32(3.14159) - 0.1))) + 1e10) * (0.1 * 0.0)) * 1e10)) {
                            if true {
                                delete(v4, 642)
                                v10 := []int{fi9}
                                var v11 float32 = ((0.0 + 1.5) + (0.0 - (-2.25 + float32((-2.25 * 1.5)))))
                                _ = v10
                                _ = v11
                            } else {
                                var v12 uint8 = (^(v5))
                                switch v2 {
                                case "a":
                                    if true {
                                        switch ((v2 + "hello") + (v2 + "\n\t")) {
                                        case "\n\t":
                                            switch fi9 {
                                            case -342:
                                                if false {
                                                    var v13 uint8 = 255
                                                    switch -2147483648 {
                                                    case 63:
                                                        v3 = uint32(v1)
                                                    case -9223372036854775808:
                                                        v14 := GenS1{F0: -486}
                                                        _ = v14
                                                    case 742:
                                                        fmt.Println(v3, v12)
                                                    default:
                                                        var v15 uint = 0
                                                        _ = v15
                                                    }
                                                    _ = v13
                                                }
                                            case 892:
                                                v16 := GenS2{F0: -2147483648, F1: (v5 & uint8(v3))}
                                                _ = v16
                                            }
                                        }
                                    }
                                case "":
                                    v17 := []int{}
                                    _ = v17
                                case "foo bar":
                                    var v18 uint32 = (^(v3))
                                    _ = v18
                                default:
                                    v19 := []int{(fi7 % ((fi9) | 1)), -1, (fi7 / (((fi9 - int(v1))) | 1))}
                                    _ = v19
                                }
                                _ = v12
                            }
                            v20 := []int{fi9, fi9, fi7, (fi7 * (-(fi9)))}
                            var v21 bool = ("hello" > v2)
                            _ = v20
                            _ = v21
                        }
                    case 825:
                        v2 = v2
                    default:
                        v2 = v2
                    }
                case -2147483648:
                    v22 := GenS3{F0: 0, F1: (v1 ^ v1)}
                    _ = v22
                default:
                    v1 = (v1 + (^(v1)))
                }
                v2 = v2
                _ = fi9
            }
            v2 = ("a" + "a")
        }
        _ = fi7
        _ = v8
    }
    _ = len(v4)
    fmt.Println(v2, v2, v2)
    v3 += (v3 ^ (^(v3)))
    var v23 int16 = 0
    switch v2 {
    case "foo bar":
        var v24 uint32 = v3
        _ = v24
    case "hello":
        if (false && ((!(!true)) && false)) {
            v25 := GenS1{F0: v23}
            for fi26 := 0; fi26 < 2; fi26++ {
                if (!true) {
                    if (false || (false && true)) {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            var v28 uint32 = v3
                            var v29 int = 2147483647
                            v5 = (v5 * (v5 / (((v5 & (v5 | v5))) | 1)))
                            if (42 != 114) {
                                var v30 float64 = (1e10 + (3.14159 + 1.5))
                                _ = v30
                            } else {
                                v23 += v23
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        var v31 int8 = -1
                        _ = v31
                    }
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 float32 = ((3.14159 * 0.1) + float32(((-2.25 * (0.0 * float64(1e10))) + 1.5)))
                    switch v2 {
                    case "\n\t":
                        var v34 uint32 = 2147483648
                        _ = v34
                    case "a":
                        v35 := []int{(fi32 << ((uint(v1)) & 63)), (fi32 / ((fi26) | 1)), int(v5)}
                        _ = v35
                    default:
                        var v36 int32 = 2147483647
                        _ = v36
                    }
                    _ = fi32
                    _ = v33
                }
                for fi37 := 0; fi37 < 8; fi37++ {
                    v38 := []int{2147483647, int(v23)}
                    var v39 uint64 = 2
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = fi26
            }
            for fi40 := 0; fi40 < 6; fi40++ {
                v41 := GenS4{F0: 100, F1: 328}
                _ = fi40
                _ = v41
            }
            _ = v25
        }
    case "a":
        switch "foo bar" {
        case "\n\t":
            v23 += (v23 + 42)
        case "hello":
            v1 = (v1 - (v1 - v1))
        case "a":
            fmt.Println(v1)
        }
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        v43 := GenS1{F0: 0}
        _ = fi42
        _ = v43
    }
    fmt.Println(v1, v2, v2)
    var v44 float64 = (1.5 + (0.1 * 0.1))
    var v45 uint = uint(v5)
    v5 = (^(v5))
    v44 = v44
    var v46 int16 = v23
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
