package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string; F2 int32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint) uint64 {
    p0 += p0
    switch 32767 {
    case 0:
        fmt.Println(p0)
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            switch "" {
            case "hello":
                v8 := GenS1{F0: (fi7 >> (((^(p0))) & 63))}
                _ = v8
            }
            fmt.Println(fi7)
            _ = fi7
        }
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 int8 = 42
    var v2 int8 = v1
    var v3 int8 = v1
    var v4 float32 = float32(1e10)
    var v5 int8 = v2
    var v6 uint8 = 228
    fmt.Println(v2, v3, v1)
    v3 = (v3 & v5)
    var v9 uint16 = 0
    fmt.Println(v6, v5)
    v10 := []int{}
    v5 = v5
    var v11 uint32 = 550
    if len(v10) > 0 {
        v10[0] = 2147483647
    }
    if len(v10) > 0 {
        v10[0] = -2147483648
    }
    v12 := GenS1{F0: 4294967295}
    for fi13 := 0; fi13 < 13; fi13++ {
        _ = len(v10)
        v14 := GenS1{F0: int(v6)}
        switch "\n\t" {
        case "":
            var v15 uint32 = v11
            _ = v15
        case "\n\t":
            _ = len(v10)
        case "foo bar":
            var v16 bool = ((^(v11)) == v11)
            _ = v16
        default:
            v17 := GenS1{F0: (fi13 - (-(fi13)))}
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v11)
    var v18 uint16 = v9
    _ = len(v10)
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            switch "\n\t" {
            case "foo bar":
                if false {
                    var v21 uint64 = 919
                    switch 127 {
                    case -2147483648:
                        v22 := GenS1{F0: (fi20 * (fi20 << ((uint(v11)) & 63)))}
                        _ = v22
                    }
                    _ = v21
                } else {
                    if (true && false) {
                        var v23 int64 = -1
                        v5 -= (v5 & -128)
                        if true {
                            var v24 int32 = int32(v2)
                            if (((42 < 713) || true) || (true || false)) {
                                v4 = v4
                                if (false && false) {
                                    v9 -= (v18 / ((v18) | 1))
                                }
                            } else {
                                var v25 int16 = 1000
                                _ = v25
                            }
                            _ = v24
                        } else {
                            var v26 int16 = -32768
                            v27 := GenS2{F0: (fi20 / (((-(fi19))) | 1)), F1: ("a" + "\n\t"), F2: 80}
                            _ = v26
                            _ = v27
                        }
                        _ = v23
                    }
                }
            case "":
                if len(v10) > 0 {
                    v10[0] = fi19
                }
            }
            _ = len(v10)
            _ = fi20
        }
        _ = fi19
    }
    v18 += (v9 * (v9 - (^(v9))))
    var v28 string = ("foo bar" + "\n\t")
    if (0 != 0) {
        v29 := GenS1{F0: 4294967295}
        var v30 float64 = 3.14159
        if false {
            v4 -= (3.14159 - (v4 + -2.25))
            _ = len(v10)
            var v31 int32 = 78
            _ = v31
        }
        _ = v29
        _ = v30
    } else {
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS3{F0: v28, F1: 781}
            var v34 uint32 = v11
            var v35 string = v28
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
