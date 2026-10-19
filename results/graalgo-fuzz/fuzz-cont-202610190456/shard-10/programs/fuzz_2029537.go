package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    var v3 uint64 = 895
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v6 := v5["a"]
    v6 -= (^(v6))
    if (1 == 65535) {
        if true {
            fmt.Println(v3)
            v7 := []int{v6, (-(v6)), -9223372036854775808}
            var v8 float64 = 3.14159
            _ = v7
            _ = v8
        }
    }
    switch "foo bar" {
    case "foo bar":
        switch "a" {
        case "hello":
            if true {
                v9 := GenS1{F0: 32767, F1: -95, F2: -82}
                _ = v9
            }
        case "\n\t":
            v2 -= (v2 << ((4294967295) & 63))
        default:
            fmt.Println(v3)
        }
    case "":
        var v10 uint64 = (v3 >> ((0) & 63))
        _ = v10
    default:
        var v11 string = "foo bar"
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v6 = fi13
            if true {
                v14 := []int{(fi13 / ((v6) | 1))}
                _ = v14
            } else {
                v15 := v1["foo bar"]
                _ = v15
            }
            v16 := GenS1{F0: 1, F1: -1, F2: -1}
            _ = fi13
            _ = v16
        }
        if false {
            var v17 string = "hello"
            _ = v17
        } else {
            if true {
                v2 = v2
                v3 = 520
                v18 := GenS1{F0: 450, F1: -83, F2: 2147483647}
                _ = v18
            }
            var v19 uint16 = 2
            _ = v19
        }
        v20 := GenS1{F0: 656, F1: -32, F2: -2147483648}
        _ = fi12
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            switch "\n\t" {
            case "":
                if true {
                    v6 -= fi21
                    v2 -= (v2 & -1)
                    var v23 uint32 = 42
                    var v24 uint16 = 32767
                    var v25 float64 = 3.14159
                    v26 := GenS2{F0: 0}
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            case "\n\t":
                delete(v5, "\n\t")
            default:
                switch fi21 {
                case 32768:
                    var v27 int32 = 2147483647
                    _ = v27
                case -858:
                    if ((65535 > 255) || true) {
                        v28 := GenS1{F0: 0, F1: 1, F2: 0}
                        if (fi22 > fi21) {
                            v29 := GenS2{F0: 9223372036854775807}
                            var v30 float64 = (0.1 * -2.25)
                            v31 := GenS3{F0: 4294967295}
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v28
                    }
                }
            }
            var v32 int8 = -1
            v3 = (^(v3))
            var v33 uint32 = 244
            _ = fi22
            _ = v32
            _ = v33
        }
        _ = fi21
    }
    var v34 bool = false
    if (((v2 >> ((497) & 63)) > (^(v2))) || (!((v3 <= 2) || (1 < 42)))) {
        delete(v4, 65535)
    } else {
        var v35 float64 = (((-2.25 + ((1e10 * 0.1) - 1.5)) - 3.14159) * (3.14159 + 1e10))
        v4[2147483647] = v34
        var v36 int32 = 333
        _ = v35
        _ = v36
    }
    fmt.Println(v2, v34)
    var v37 string = "a"
    for fi38 := 0; fi38 < 17; fi38++ {
        v3 = v3
        _ = fi38
    }
    var v39 int32 = 42
    if false {
        var v40 uint8 = 0
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 int8 = 42
            v43 := GenS4{F0: v39, F1: (v39 << ((255) & 63))}
            var v44 string = v37
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        v45 := v5["\n\t"]
        _ = v40
        _ = v45
    }
    fmt.Println(v37, v6)
    if v34 {
        fmt.Println(v3)
        v46 := GenS1{F0: 917, F1: -128, F2: v39}
        _ = v46
    }
    var v47 int32 = -2147483648
    var v48 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
