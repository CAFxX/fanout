package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 bool }
type GenS4 struct { F0 string; F1 uint16; F2 string }
type GenS5 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint = 681
    var v2 bool = false
    var v3 uint64 = 42
    var v4 uint16 = 2
    var v5 bool = false
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint = v1
        fmt.Println(v7, v2)
        _ = fi6
        _ = v7
    }
    if v2 {
        v8 := []int{-508}
        _ = v8
    } else {
        v9 := GenS1{F0: 1000, F1: 2147483647}
        var v10 int64 = 9223372036854775807
        v11 := []int{}
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(v4, v1)
    v12 := []int{56, 32768, 9223372036854775807}
    for fi13 := 0; fi13 < 11; fi13++ {
        v5 = v5
        _ = fi13
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        var v14 uint16 = v4
        _ = v14
    case "\n\t":
        v4 = v4
    case "a":
        fmt.Println(v4, v4)
    default:
        var v15 uint8 = 42
        _ = v15
    }
    var v16 uint8 = 15
    if (127 > -91) {
        switch "" {
        case "hello":
            v17 := GenS1{F0: 329, F1: 2147483648}
            _ = v17
        default:
            v18 := GenS1{F0: 2, F1: 2147483647}
            _ = v18
        }
    } else {
        v19 := GenS2{F0: v1, F1: uint32(v4)}
        var v20 uint64 = 18446744073709551615
        _ = v19
        _ = v20
    }
    if v5 {
        v21 := GenS3{F0: v4, F1: uint64(v16), F2: v5}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 12; fi22++ {
            v1 += (v1 & (v1 % ((uint(v16)) | 1)))
            if v5 {
                v23 := GenS3{F0: (v4 / ((v4) | 1)), F1: 42, F2: v5}
                _ = v23
            } else {
                if ((4294967295 >= (v3 | 42)) && (9223372036854775807 < -768)) {
                    var v24 int64 = -469
                    _ = v24
                } else {
                    v2 = (!v5)
                    v12 = append(v12, fi22)
                }
                fmt.Println(v5)
                v12 = append(v12, 128)
            }
            _ = fi22
        }
    }
    var v25 int64 = 1
    var v26 float64 = float64(0.0)
    fmt.Println(v3, v3, v26)
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            v4 = 0
            v2 = v2
            _ = fi28
        }
        var v29 int32 = 2147483647
        _ = fi27
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        switch fi30 {
        case -1:
            for fi31 := 0; fi31 < 7; fi31++ {
                if true {
                    switch "a" {
                    case "hello":
                        var v32 uint32 = 613
                        _ = v32
                    }
                    var v33 bool = (!v5)
                    v34 := GenS4{F0: "hello", F1: uint16(v3), F2: "a"}
                    switch fi30 {
                    case 42:
                        var v35 uint64 = 898
                        _ = v35
                    case 128:
                        _ = len(v12)
                    }
                    v25 = v25
                    var v36 string = "\n\t"
                    var v37 int16 = 294
                    _ = v33
                    _ = v34
                    _ = v36
                    _ = v37
                } else {
                    v26 = ((v26 * v26) + v26)
                    v26 = float64(1.5)
                    var v38 float64 = v26
                    _ = v38
                }
                var v39 uint32 = 2147483648
                _ = fi31
                _ = v39
            }
        }
        _ = fi30
    }
    switch ("a" + "foo bar") {
    case "":
        var v40 string = "hello"
        _ = v40
    case "hello":
        switch "a" {
        case "a":
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 string = "\n\t"
                v43 := GenS5{F0: 756, F1: 1}
                _ = fi41
                _ = v42
                _ = v43
            }
        case "\n\t":
            switch "\n\t" {
            case "a":
                for fi44 := 0; fi44 < 17; fi44++ {
                    var v45 uint64 = (v3 - v3)
                    v16 -= v16
                    _ = len(v12)
                    _ = fi44
                    _ = v45
                }
            case "\n\t":
                _ = len(v12)
            }
        case "":
            switch "foo bar" {
            case "\n\t":
                for fi46 := 0; fi46 < 16; fi46++ {
                    fmt.Println(v2)
                    _ = fi46
                }
            case "":
                fmt.Println(v4, v25, v5)
            case "a":
                var v47 string = ("a" + "")
                _ = v47
            default:
                if v5 {
                    var v48 float32 = -2.25
                    for fi49 := 0; fi49 < 20; fi49++ {
                        if (v2 && (343 >= 4294967295)) {
                            v26 += 0.1
                            v1 -= (v1 % ((v1) | 1))
                        }
                        _ = fi49
                    }
                    var v50 int8 = 42
                    _ = v48
                    _ = v50
                }
            }
        default:
            var v51 int64 = v25
            _ = v51
        }
    case "\n\t":
        v25 -= (v25 - (-(v25)))
    default:
        switch 701 {
        case 1:
            fmt.Println(v16, v5)
        case 4294967295:
            var v52 int = 890
            _ = v52
        }
    }
    var v53 float64 = v26
    if true {
        for fi54 := 0; fi54 < 10; fi54++ {
            var v55 int8 = 100
            _ = fi54
            _ = v55
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        if len(v12) > 0 {
            v12[0] = 9223372036854775807
        }
    }
    if len(v12) > 0 {
        v12[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
