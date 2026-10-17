package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int32 }
func genFn0(p0 bool, p1 uint8, p2 string) int8 {
    if p0 {
        fmt.Println(p0, p2, p1)
        v5 := []int{235, -961, 32767, 127}
        for fi6 := 0; fi6 < 4; fi6++ {
            v5 = append(v5, (fi6 % (((fi6 + -369)) | 1)))
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: p2, F1: 0, F2: (^(p1))}
                v9 := GenS2{F0: 1, F1: 188, F2: 826}
                v10 := GenS3{F0: 42, F1: 259, F2: p1}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        _ = v5
    } else {
        v11 := GenS4{F0: -532}
        switch p2 {
        case "hello":
            if (true || true) {
                v12 := GenS5{F0: -1}
                _ = v12
            }
        case "foo bar":
            v13 := []int{805}
            _ = v13
        }
        _ = v11
    }
    v14 := []int{1}
    if true {
        p2 = "a"
    } else {
        if p0 {
            switch 9223372036854775807 {
            case -2147483648:
                switch p2 {
                case "a":
                    p1 = 24
                case "":
                    p1 -= (p1 % ((2) | 1))
                }
            case 913:
                switch p2 {
                case "\n\t":
                    v15 := GenS3{F0: 669, F1: 255, F2: (p1 | p1)}
                    _ = v15
                case "hello":
                    v16 := GenS6{F0: 0, F1: uint(p1), F2: 177}
                    _ = v16
                }
            case 1:
                var v17 float32 = 1.5
                _ = v17
            default:
                p1 -= (p1 ^ 128)
            }
        }
        var v18 bool = (p2 != p2)
        var v19 int32 = 2147483647
        _ = v18
        _ = v19
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}
func genFn1() bool {
    var v20 uint64 = 257
    v20 -= (v20 * 151)
    v21 := GenS3{F0: 631, F1: 255, F2: 42}
    var v22 uint64 = 25
    var v23 string = ("hello" + "\n\t")
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    return (357 >= 128)
}

func main() {
    var v1 int32 = 0
    v2 := []int{989, 32768}
    var v3 uint8 = 0
    v4 := make(map[string]int)
    switch ("" + "\n\t") {
    case "a":
        fmt.Println(v3, v1)
    default:
        if false {
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 uint8 = v3
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 float64 = float64(3.14159)
                    fmt.Println(v3, v3, v25)
                    if ((^(v1)) >= 42) {
                        fmt.Println(fi24, v1, v3)
                        switch fi26 {
                        case -808:
                            switch "\n\t" {
                            case "foo bar":
                                var v28 string = "\n\t"
                                _ = v28
                            case "hello":
                                switch (fi24 ^ 32768) {
                                case -2147483648:
                                    v27 = v27
                                case 128:
                                    v27 -= v27
                                }
                            }
                            var v29 int = (fi24 % ((-9223372036854775808) | 1))
                            v30 := v4["a"]
                            v31 := GenS7{F0: 9223372036854775807, F1: genFn0((42 <= v25), (^(v3)), ("a" + ("hello" + "foo bar")))}
                            _ = v29
                            _ = v30
                            _ = v31
                        case 128:
                            v3 -= v25
                        }
                        var v32 uint8 = (v3 - v3)
                        _ = v32
                    }
                    v27 += (1e10 / (3.14159 * v27))
                    var v33 string = ("hello" + "\n\t")
                    _ = fi26
                    _ = v27
                    _ = v33
                }
                _ = fi24
                _ = v25
            }
            fmt.Println(v1, v1)
            v34 := GenS2{F0: -659, F1: 2147483648, F2: 4294967295}
            _ = v34
        } else {
            switch -9223372036854775808 {
            case 1:
                v35 := GenS5{F0: 42}
                _ = v35
            case -9223372036854775808:
                var v36 int32 = 42
                _ = v36
            case 9223372036854775807:
                var v37 int16 = -439
                _ = v37
            default:
                v38 := GenS2{F0: 42, F1: 2, F2: 65535}
                _ = v38
            }
            var v39 uint = 806
            _ = v39
        }
    }
    var v40 bool = true
    v41 := GenS2{F0: 1, F1: 792, F2: 114}
    var v42 bool = v40
    var v43 int32 = (v1 * (v1 & 535))
    if (!(v40 || false)) {
        for fi44 := 0; fi44 < 10; fi44++ {
            v42 = v40
            _ = fi44
        }
    } else {
        fmt.Println(v42, v43, v42)
    }
    if v42 {
        for fi45 := 0; fi45 < 17; fi45++ {
            v40 = (0 > 0)
            _ = fi45
        }
        if (!v42) {
            v42 = (!(false || v40))
        }
    } else {
        switch 94 {
        case 4294967295:
            fmt.Println(v42)
        }
    }
    v46 := GenS8{F0: -2147483648}
    for fi47 := 0; fi47 < 17; fi47++ {
        var v48 string = "\n\t"
        fmt.Println(v3, fi47)
        _ = fi47
        _ = v48
    }
    v3 = 255
    var v49 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v46
    _ = v49
}
