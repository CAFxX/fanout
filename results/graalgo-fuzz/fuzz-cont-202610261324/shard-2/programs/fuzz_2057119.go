package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int64; F2 int }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int32; F1 int; F2 int8 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 uint8; F1 int; F2 uint8 }
type GenS9 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -1
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    if false {
        var v7 float32 = 0.1
        v7 = v7
        switch 42 {
        case -852:
            var v8 uint64 = 0
            _ = v8
        case -9223372036854775808:
            switch 42 {
            case 9223372036854775807:
                if (0 <= 0) {
                    var v9 uint64 = 29
                    switch ("a" + "\n\t") {
                    case "a":
                        v10 := GenS1{F0: 16, F1: "a", F2: -946}
                        _ = v10
                    default:
                        v9 = v9
                    }
                    if true {
                        _ = len(v5)
                    }
                    _ = v9
                }
            case 4294967295:
                switch 1 {
                case 32767:
                    v11 := GenS1{F0: 231, F1: "", F2: 0}
                    _ = v11
                case 4294967295:
                    switch 2147483647 {
                    case 32767:
                        v12 := GenS2{F0: 846, F1: 32768, F2: 524}
                        _ = v12
                    case -837:
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v14 := GenS2{F0: 18446744073709551615, F1: fi13, F2: 2147483647}
                            fmt.Println(fi13, v7, v2)
                            _ = fi13
                            _ = v14
                        }
                    default:
                        v15 := []int{9223372036854775807}
                        _ = v15
                    }
                }
            }
        default:
            if true {
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS1{F0: 65535, F1: "a", F2: 2147483647}
                    _ = fi16
                    _ = v17
                }
                _ = len(v6)
                v18 := GenS1{F0: 65535, F1: "", F2: 640}
                _ = v18
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS3{F0: 749, F1: v2, F2: 127}
                    if (100 <= 1) {
                        v21 := GenS4{F0: 112, F1: -659, F2: (fi19 * (-(fi19)))}
                        _ = v21
                    } else {
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    }
                    switch (("hello" + (("hello" + "") + "")) + ("foo bar" + ("" + (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "\n\t")))) {
                    case "hello":
                        v23 := GenS5{F0: 32767, F1: 605}
                        _ = v23
                    }
                    _ = fi19
                    _ = v20
                }
            }
        }
        _ = v7
    }
    v24 := GenS6{F0: 852, F1: -378, F2: 100}
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 float64 = (0.0 * float64(-2.25))
        _ = fi25
        _ = v26
    }
    v27 := GenS4{F0: 4294967295, F1: -517, F2: 9223372036854775807}
    if false {
        var v28 float32 = 1e10
        v2 += (-(v2))
        _ = v28
    }
    switch -689 {
    case -2147483648:
        _ = len(v3)
    case 32768:
        v2 -= v2
    default:
        if false {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 uint16 = 2
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(v2, fi31, fi29)
                    _ = fi31
                }
                _ = fi29
                _ = v30
            }
            switch -1 {
            case 127:
                fmt.Println(v2)
            case 128:
                var v32 bool = true
                _ = v32
            case -728:
                for fi33 := 0; fi33 < 10; fi33++ {
                    if true {
                        if ((v2 >> ((255) & 63)) != v2) {
                            v34 := GenS7{F0: -2147483648, F1: 1}
                            _ = v34
                        }
                    }
                    _ = len(v4)
                    _ = fi33
                }
            }
        } else {
            fmt.Println(v2)
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        v2 = v2
        if false {
            switch (fi35 / (((fi35 | fi35)) | 1)) {
            case 2147483647:
                fmt.Println(v2, fi35)
            case -492:
                if ("foo bar" >= "") {
                    for fi36 := 0; fi36 < 4; fi36++ {
                        switch "a" {
                        case "hello":
                            delete(v3, "a")
                        case "":
                            if false {
                                switch -9223372036854775808 {
                                case 2147483647:
                                    if (42 <= 174) {
                                        switch 4294967295 {
                                        case 127:
                                            var v37 uint32 = 65535
                                            _ = v37
                                        default:
                                            v38 := []int{fi36, (fi35 % (((fi35 * fi36)) | 1))}
                                            _ = v38
                                        }
                                    } else {
                                        v39 := GenS8{F0: 42, F1: fi36, F2: 0}
                                        _ = v39
                                    }
                                }
                                v40 := []int{(fi35 ^ (fi36 >> ((901) & 63))), (fi35 + (^(fi36))), fi36, fi36}
                                _ = v40
                            }
                        case "a":
                            _ = len(v6)
                        }
                        v41 := []int{(fi36 / ((4294967295) | 1)), (^(fi35))}
                        _ = fi36
                        _ = v41
                    }
                    switch (fi35 >> ((708) & 63)) {
                    case 128:
                        v2 = v2
                    }
                    var v42 uint8 = 216
                    _ = v42
                }
            case 127:
                var v43 uint = 125
                _ = v43
            }
        }
        v2 += v2
        _ = fi35
    }
    v2 -= (v2 * (v2 * v2))
    v44 := GenS9{F0: 9223372036854775807}
    v45 := []int{127, -2147483648, -9223372036854775808}
    fmt.Println(v2)
    for fi46 := 0; fi46 < 5; fi46++ {
        v2 = (v2 * v2)
        var v47 int32 = 2147483647
        var v48 int64 = 9223372036854775807
        _ = fi46
        _ = v47
        _ = v48
    }
    v2 = (-(v2))
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
