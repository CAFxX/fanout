package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int64; F2 int32 }
func genFn0(p0 uint, p1 int64) int16 {
    if (-1 != 344) {
        var v7 uint32 = 664
        var v8 int8 = -1
        v8 += (^(v8))
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if (42 > 634) {
            p1 = (p1 << (((p0 | (p0 + (p0 ^ p0)))) & 63))
        }
        _ = fi9
    }
    var v10 uint32 = 1
    _ = p0
    _ = p1
    _ = v10
    return 778
}

func main() {
    var v1 int8 = 100
    var v2 int16 = 619
    v3 := []uint32{65535, 4294967295}
    v4 := make(map[int]int)
    v5 := []uint32{2147483648, 200, 42}
    v6 := []string{"foo bar", ((("" + "\n\t") + ("\n\t" + "a")) + (("hello" + ("hello" + "a")) + "\n\t"))}
    v1 = v1
    switch "\n\t" {
    case "":
        v11 := GenS1{F0: 2147483647}
        _ = v11
    case "a":
        switch 664 {
        case -2147483648:
            if (197 == 55) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        v2 = (^(v2))
                        v14 := GenS1{F0: 42}
                        var v15 uint16 = 42
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v2 += v2
                    var v16 bool = true
                    _ = fi12
                    _ = v16
                }
                var v17 uint16 = 818
                _ = v17
            }
        }
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            v2 += genFn0(2, -322)
            _ = fi18
        }
    }
    if (!((1.5 + 1e10) != 0.0)) {
        for fi19 := 0; fi19 < 11; fi19++ {
            switch 88 {
            case 0:
                v1 = -65
            }
            _ = len(v4)
            _ = fi19
        }
        v20 := GenS1{F0: 0}
        _ = v20
    }
    _ = len(v4)
    v2 -= (^(v2))
    v21 := GenS1{F0: 2147483648}
    v2 = genFn0(65535, 9223372036854775807)
    if ((!true) && (!false)) {
        var v22 uint32 = 2147483648
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS2{F0: (-(fi24))}
                v1 = v1
                v2 += -351
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 uint16 = 32767
                v28 := GenS3{F0: fi26}
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 float64 = 1e10
            _ = fi23
            _ = v29
        }
        _ = v22
    } else {
        var v30 uint64 = 0
        _ = v30
    }
    v31 := GenS4{F0: (^(v1)), F1: 477, F2: -1}
    v32 := GenS2{F0: 42}
    var v33 uint64 = 975
    var v34 int16 = int16(v1)
    v1 = (^(v1))
    for fi35 := 0; fi35 < 10; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 uint = 255
            delete(v4, (fi36 & fi35))
            _ = fi36
            _ = v37
        }
        var v38 float64 = 3.14159
        for fi39 := 0; fi39 < 19; fi39++ {
            v40 := GenS1{F0: 4294967295}
            _ = fi39
            _ = v40
        }
        _ = fi35
        _ = v38
    }
    delete(v4, -32)
    if (0.0 >= 3.14159) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (v34 == int16(v33)) {
            v5 = append(v5, 873)
            var v41 uint = 18446744073709551615
            _ = v41
        }
    }
    v42 := GenS3{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
