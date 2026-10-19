package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int64; F2 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    v3 := []int{}
    var v4 uint16 = 32767
    var v5 uint = 431
    if (true || false) {
        if (!false) {
            v4 = v4
            if ((((-2.25 * float64((0.0 - 1.5))) * -2.25) < 1e10) && true) {
                var v6 int16 = 906
                if true {
                    v4 = (v4 - (v4 << ((18446744073709551615) & 63)))
                    v6 -= 32767
                    v6 += v6
                }
                _ = v6
            }
            var v7 float32 = 0.0
            _ = v7
        }
        v8 := GenS1{F0: "a"}
        _ = v8
    } else {
        fmt.Println(v4, v4)
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v4)
        v5 = v5
        _ = fi9
    }
    if true {
        v10 := GenS1{F0: "a"}
        _ = v10
    } else {
        v4 = (v4 + 738)
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        switch fi11 {
        case 42:
            var v12 uint32 = 2
            _ = v12
        }
        v4 = v4
        if false {
            switch 924 {
            case 26:
                v4 += v4
            case 32768:
                var v13 int8 = 0
                _ = v13
            case 0:
                v14 := GenS2{F0: 123, F1: -200, F2: 825}
                _ = v14
            }
        }
        _ = fi11
    }
    switch "a" {
    case "\n\t":
        var v15 uint64 = 190
        _ = v15
    default:
        var v16 float64 = 0.0
        _ = v16
    }
    v4 += (v4 & (v4 & uint16(v5)))
    v4 -= 42
    switch 342 {
    case 1:
        fmt.Println(v4, v4)
    case 42:
        switch 32767 {
        case 0:
            var v17 bool = false
            _ = v17
        case 2147483647:
            var v18 int = 0
            _ = v18
        default:
            v19 := GenS2{F0: 255, F1: -9223372036854775808, F2: v5}
            _ = v19
        }
    }
    switch 42 {
    case 9223372036854775807:
        v5 = (v5 * v5)
    case 127:
        v4 -= 238
    case -9223372036854775808:
        fmt.Println(v4)
    default:
        v20 := GenS1{F0: ""}
        _ = v20
    }
    v21 := GenS3{F0: 826, F1: 0, F2: 1}
    switch 9223372036854775807 {
    case -173:
        if false {
            v22 := GenS3{F0: 957, F1: 42, F2: 42}
            _ = v22
        } else {
            fmt.Println(v4)
        }
    case -667:
        fmt.Println(v4, v4)
    case -564:
        v23 := GenS4{F0: 32767}
        _ = v23
    }
    v24 := GenS4{F0: -2147483648}
    _ = len(v1)
    fmt.Println(v5)
    switch -1 {
    case 32768:
        v5 += v5
    default:
        for fi25 := 0; fi25 < 18; fi25++ {
            if false {
                var v26 int8 = 127
                switch ("hello" + "\n\t") {
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = int(v4)
                    }
                case "":
                    fmt.Println(v5, fi25, v5)
                default:
                    if (true && (true && (!(42 == 466)))) {
                        switch fi25 {
                        case 32768:
                            v27 := GenS1{F0: ""}
                            _ = v27
                        }
                        switch "a" {
                        case "hello":
                            for fi28 := 0; fi28 < 11; fi28++ {
                                delete(v1, -960)
                                var v29 int16 = 887
                                _ = fi28
                                _ = v29
                            }
                        case "":
                            v30 := GenS3{F0: 65535, F1: -2147483648, F2: 1}
                            _ = v30
                        case "\n\t":
                            for fi31 := 0; fi31 < 6; fi31++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi31
                            }
                        }
                    } else {
                        var v32 int32 = -838
                        _ = v32
                    }
                }
                v33 := GenS1{F0: ""}
                _ = v26
                _ = v33
            }
            v34 := GenS5{F0: 92, F1: 2147483647}
            _ = fi25
            _ = v34
        }
    }
    var v35 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
