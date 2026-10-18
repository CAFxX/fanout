package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0() int8 {
    var v6 uint32 = 2147483647
    v7 := []int{2147483647, -1, 42, 42}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 100
}
func genFn1() uint {
    switch "foo bar" {
    case "\n\t":
        var v8 uint16 = 168
        _ = v8
    case "":
        var v9 uint8 = 137
        _ = v9
    default:
        v10 := []int{140, 858}
        _ = v10
    }
    var v11 int64 = -9223372036854775808
    _ = v11
    return 65535
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    var v5 int = 1
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int32 = 42
        switch (fi12 << ((255) & 63)) {
        case -2147483648:
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 int8 = 0
                v16 := GenS1{F0: 916}
                var v17 float32 = ((float32(1.5) + float32((1e10 - 3.14159))) - -2.25)
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        default:
            var v18 float32 = 1.5
            _ = v18
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v1)
    var v19 bool = true
    if v19 {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 float64 = (0.0 * 0.1)
            var v22 uint32 = 2147483647
            if v19 {
                var v23 float32 = (1.5 / (0.1 * (1.5 - 0.0)))
                _ = v23
            } else {
                v22 += (v22 - 2)
                for fi24 := 0; fi24 < 2; fi24++ {
                    fmt.Println(v5, v5)
                    var v25 int64 = -1
                    v25 -= (v25 & v25)
                    _ = fi24
                    _ = v25
                }
                if (!v19) {
                    v1 -= v1
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 bool = (!true)
                        _ = fi26
                        _ = v27
                    }
                    v28 := v2["a"]
                    _ = v28
                } else {
                    if v19 {
                        v29 := []int{fi20, 118, (v5 + (v5 << ((0) & 63)))}
                        v1 = 1
                        _ = len(v2)
                        _ = v29
                    } else {
                        v22 -= (v22 / (((v22 - v22)) | 1))
                    }
                    v1 = (^(v1))
                }
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 8; fi31++ {
                v19 = (2 != 65535)
                fmt.Println(v1)
                v32 := GenS2{F0: 42, F1: 433, F2: 1}
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
    }
    var v33 int32 = 42
    v34 := []int{498, 4294967295}
    for fi35 := 0; fi35 < 18; fi35++ {
        if v19 {
            var v36 float64 = 0.0
            if false {
                var v37 float32 = (0.1 - 1.5)
                v5 = v5
                _ = v37
            }
            _ = v36
        }
        v33 = (v33 + -1)
        _ = fi35
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        switch v5 {
        case 2147483647:
            var v39 int64 = 762
            _ = v39
        case 1:
            if v19 {
                var v40 int32 = -2147483648
                for fi41 := 0; fi41 < 5; fi41++ {
                    if v19 {
                        var v42 int = v5
                        _ = v42
                    } else {
                        if v19 {
                            v5 = fi38
                            var v43 int = (fi41 * (fi41 / ((fi41) | 1)))
                            var v44 float64 = 0.0
                            var v45 int32 = 1
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = fi41
                }
                _ = v40
            }
        default:
            switch "\n\t" {
            case "a":
                var v46 float64 = ((1e10 + 1.5) + -2.25)
                _ = v46
            case "":
                switch int(v1) {
                case 127:
                    if (3.14159 < 0.1) {
                        v1 -= 1
                        for fi47 := 0; fi47 < 18; fi47++ {
                            v33 -= -1
                            _ = fi47
                        }
                    }
                case -602:
                    v33 = int32(v1)
                }
            default:
                v34 = append(v34, (-(fi38)))
            }
        }
        v33 += -397
        _ = fi38
    }
    _ = len(v3)
    if true {
        if len(v34) > 0 {
            v34[0] = (v5 % ((32767) | 1))
        }
        v48 := GenS1{F0: 621}
        _ = v48
    }
    switch v5 {
    case 499:
        for fi49 := 0; fi49 < 8; fi49++ {
            var v50 float64 = (0.1 + ((float64((-2.25 / (1e10 - 0.0))) / -2.25) / 0.1))
            v51 := GenS3{F0: (v33 * v33), F1: (^(v33))}
            _ = fi49
            _ = v50
            _ = v51
        }
    }
    fmt.Println(v1)
    var v52 bool = v19
    v1 = v1
    switch ((("hello" + (("a" + "\n\t") + "a")) + (("\n\t" + "a") + "\n\t")) + "a") {
    case "\n\t":
        v33 += (v33 >> ((551) & 63))
    case "a":
        var v53 int64 = -1
        _ = v53
    case "hello":
        if false {
            v52 = true
            v54 := GenS1{F0: 301}
            var v55 int8 = 127
            _ = v54
            _ = v55
        } else {
            var v56 int8 = -128
            v57 := GenS3{F0: v33, F1: (v33 - 159)}
            _ = v56
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v52
}
