package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int64 }
func genFn0(p0 uint32, p1 uint64) int32 {
    fmt.Println(p1)
    v7 := GenS1{F0: p1, F1: 0}
    var v8 uint8 = 0
    if (176 < uint8(p1)) {
        var v9 int8 = -110
        var v10 string = "foo bar"
        v11 := GenS1{F0: 0, F1: v8}
        _ = v10
        _ = v11
        _ = v9
    } else {
        switch -1 {
        case 42:
            v12 := GenS1{F0: 0, F1: v8}
            _ = v12
        case 2147483647:
            if true {
                if (!(false || (!true))) {
                    if (false && true) {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            p0 = (p0 | 2147483648)
                            var v14 int64 = -1
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        var v15 int = 9223372036854775807
                        switch (v15 >> ((4294967295) & 63)) {
                        case -961:
                            v16 := []int{(^(v15))}
                            _ = v16
                        }
                        _ = v15
                    }
                    for fi17 := 0; fi17 < 20; fi17++ {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            var v19 uint8 = v8
                            v19 += 20
                            var v20 uint32 = 536
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                    var v21 uint32 = 4294967295
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        p0 = 909
                        var v23 int64 = 826
                        _ = fi22
                        _ = v23
                    }
                    p0 -= 224
                }
            }
        }
        v24 := []int{156, 128, 0, 32767}
        _ = v24
    }
    if ((539 < 1) || false) {
        v25 := GenS2{F0: 10, F1: 42}
        v26 := GenS3{F0: (v8 * uint8(p0)), F1: v8, F2: 1}
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 string = ("\n\t" + ((("a" + ("" + "\n\t")) + "hello") + "hello"))
    var v4 uint16 = 118
    v5 := make(map[string]string)
    var v6 uint64 = (v2 * 2)
    fmt.Println(v4)
    var v27 uint64 = (v2 - v6)
    if (!(!false)) {
        var v28 uint64 = (v6 | 42)
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := []int{(fi29 + fi29), fi29}
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        v6 = 4294967295
        v3 = v3
    }
    fmt.Println(v4, v2)
    if false {
        v2 += (v6 >> ((uint(v4)) & 63))
    }
    v6 -= 2
    for fi31 := 0; fi31 < 18; fi31++ {
        fmt.Println(v2)
        switch v3 {
        case "a":
            v32 := GenS1{F0: (v27 << ((229) & 63)), F1: 65}
            _ = v32
        case "\n\t":
            fmt.Println(v27, v4)
        default:
            v3 = v3
        }
        v4 -= (v4 + (v4 << ((567) & 63)))
        _ = fi31
    }
    if true {
        var v33 uint = 525
        v34 := GenS1{F0: 834, F1: 0}
        if (-2.25 < (1.5 - (0.1 - 3.14159))) {
            v27 -= (v27 ^ v27)
        }
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
