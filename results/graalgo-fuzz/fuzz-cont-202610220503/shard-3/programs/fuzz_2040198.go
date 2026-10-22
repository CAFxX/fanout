package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 int8, p1 bool, p2 uint32) uint32 {
    var v6 uint = 835
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 int16, p1 int16) uint64 {
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = p1
        for fi8 := 0; fi8 < 10; fi8++ {
            p1 = p0
            var v9 uint16 = 255
            p1 += -671
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint32 = 4294967295
    v11 := GenS1{F0: v10, F1: ("\n\t" + "foo bar"), F2: (v10 >> ((2) & 63))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := []uint32{}
    var v2 string = "foo bar"
    var v3 int64 = 1
    var v4 int32 = 1
    var v5 int = 0
    if (!false) {
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 string = ("foo bar" + v2)
            _ = fi12
            _ = v13
        }
        var v14 int64 = v3
        _ = v14
    } else {
        var v15 int32 = (v4 ^ 0)
        switch v5 {
        case 617:
            v15 = v4
        case 548:
            if (2 > 18446744073709551615) {
                v5 -= (-(v5))
                var v16 int8 = -123
                _ = v16
            }
        case -472:
            var v17 uint8 = 42
            _ = v17
        default:
            v3 = 797
        }
        var v18 uint32 = genFn0(-93, (v5 != v5), 791)
        _ = v15
        _ = v18
    }
    v19 := GenS1{F0: 2, F1: v2, F2: 991}
    var v20 bool = (0 >= 65535)
    switch (v2 + (("\n\t" + (("\n\t" + (v2 + v2)) + v2)) + "hello")) {
    case "hello":
        var v21 uint32 = 2147483648
        _ = v21
    case "\n\t":
        v22 := GenS1{F0: 4294967295, F1: v2, F2: 4294967295}
        _ = v22
    default:
        v23 := GenS1{F0: 2147483648, F1: "\n\t", F2: 827}
        _ = v23
    }
    if v20 {
        _ = len(v1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v5 += v5
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v3, v5, v2)
            _ = fi24
        }
        var v25 int64 = v3
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 uint16 = 65535
            if ((0.1 < (3.14159 * 0.0)) || v20) {
                if (7 < 18446744073709551615) {
                    var v28 int16 = -32768
                    fmt.Println(v3, fi26, v5)
                    _ = v28
                } else {
                    switch (fi26 ^ (fi26 >> ((1) & 63))) {
                    case -9223372036854775808:
                        var v29 int16 = 1
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 14; fi30++ {
                        var v31 int16 = 1
                        var v32 float32 = (0.1 + 3.14159)
                        v33 := GenS1{F0: 2147483648, F1: v2, F2: 42}
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v27 = (v27 | (v27 >> ((2) & 63)))
                }
            }
            var v34 string = (v2 + "hello")
            _ = fi26
            _ = v27
            _ = v34
        }
        _ = v25
    }
    var v35 int8 = 127
    v36 := GenS2{F0: 4294967295, F1: v4}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v37 int = (v5 ^ v5)
    for fi38 := 0; fi38 < 7; fi38++ {
        if (0 <= int16(v35)) {
            switch 32767 {
            case 460:
                fmt.Println(v37, v35, v5)
            case 127:
                var v39 uint8 = 0
                _ = v39
            case 4294967295:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                switch v5 {
                case 4294967295:
                    var v40 int16 = 1000
                    _ = v40
                }
            }
            v4 = (v4 ^ (v4 ^ 2147483647))
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
        var v41 bool = false
        _ = fi38
        _ = v41
    }
    v2 = (v2 + (v2 + v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
