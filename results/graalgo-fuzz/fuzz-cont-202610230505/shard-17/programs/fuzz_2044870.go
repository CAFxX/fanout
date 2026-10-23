package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint32) int64 {
    if false {
        if false {
            for fi5 := 0; fi5 < 9; fi5++ {
                var v6 float64 = (-2.25 / (1.5 / ((((3.14159 / ((float64(-2.25) / 1.5) - float64(1.5))) - 3.14159) * 0.0) + (1.5 + (-2.25 * (0.0 + (((0.1 + (0.1 / 0.1)) + float64((((0.0 - 1e10) / 0.1) / 0.1))) + -2.25)))))))
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 uint = 18446744073709551615
                    _ = fi7
                    _ = v8
                }
                _ = fi5
                _ = v6
            }
        }
    }
    fmt.Println(p0)
    var v9 uint64 = 824
    fmt.Println(v9, p0)
    if true {
        v10 := GenS1{F0: 1}
        _ = v10
    } else {
        v9 -= (v9 ^ 18446744073709551615)
        switch 42 {
        case 4294967295:
            fmt.Println(p0, v9)
        }
    }
    _ = p0
    _ = v9
    return 0
}
func genFn1(p0 uint32, p1 int32, p2 string) int32 {
    p1 -= (p1 % (((p1 - (p1 - p1))) | 1))
    var v11 uint16 = 255
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int64 = 0
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p1
}

func main() {
    var v1 uint32 = 0
    v2 := []string{("a" + ("" + "\n\t")), "hello", (("a" + "hello") + ("\n\t" + "foo bar")), "\n\t"}
    v3 := make(map[uint32]int)
    var v4 int64 = 1
    if ((true && false) || (((0 != 0) && false) || false)) {
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 uint8 = 2
                if true {
                    v18 := GenS1{F0: v4}
                    _ = v18
                } else {
                    var v19 int8 = -105
                    _ = v19
                }
                fmt.Println(v1)
                _ = fi16
                _ = v17
            }
            var v20 uint16 = 207
            _ = fi15
            _ = v20
        }
        if (2 == 2) {
            v21 := GenS1{F0: (v4 % (((^(v4))) | 1))}
            for fi22 := 0; fi22 < 15; fi22++ {
                for fi23 := 0; fi23 < 3; fi23++ {
                    v1 += 1
                    v1 -= 0
                    v1 += 1
                    _ = fi23
                }
                var v24 string = "hello"
                _ = fi22
                _ = v24
            }
            v1 -= v1
            _ = v21
        }
        var v25 int = 640
        _ = v25
    } else {
        var v26 uint = 255
        _ = v26
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v27 bool = (!true)
    switch ("\n\t" + "foo bar") {
    case "a":
        if (42 <= 2) {
            var v28 int8 = 0
            fmt.Println(v4, v28)
            _ = v28
        }
    default:
        switch -9223372036854775808 {
        case 521:
            for fi29 := 0; fi29 < 8; fi29++ {
                v4 = v4
                _ = fi29
            }
        case 2147483647:
            v30 := GenS2{F0: 252, F1: v1, F2: 0}
            _ = v30
        default:
            if v27 {
                var v31 uint = 631
                _ = v31
            }
        }
    }
    v32 := GenS3{F0: 18446744073709551615}
    _ = len(v3)
    fmt.Println(v27)
    for fi33 := 0; fi33 < 14; fi33++ {
        switch fi33 {
        case 42:
            var v34 int = fi33
            _ = v34
        default:
            _ = len(v3)
        }
        fmt.Println(v4)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
