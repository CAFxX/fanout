package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 int16 }
func genFn0(p0 uint16) string {
    for fi5 := 0; fi5 < 9; fi5++ {
        switch 825 {
        case 4294967295:
            v6 := []int{}
            _ = v6
        case -9223372036854775808:
            p0 += 255
        }
        fmt.Println(p0, fi5)
        _ = fi5
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return "\n\t"
}
func genFn1(p0 int16, p1 uint32, p2 uint16) uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = (^(p0))
        _ = fi7
    }
    switch -479 {
    case -709:
        p2 -= 0
    case -613:
        v8 := GenS1{F0: (p1 >> ((1) & 63)), F1: 152}
        _ = v8
    case 0:
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 string = ""
            fmt.Println(p0, v10, fi9)
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int = (fi11 - fi11)
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 42
    for fi14 := 0; fi14 < 3; fi14++ {
        switch -366 {
        case 309:
            fmt.Println(p0)
        case 9223372036854775807:
            v15 := []int{(fi14 - (fi14 << ((1) & 63))), (fi14 ^ 98)}
            _ = v15
        }
        var v16 int8 = -1
        v16 = (v16 ^ (v16 << ((4294967295) & 63)))
        _ = fi14
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := []uint32{1, 2, 310}
    v2 := make(map[uint32]string)
    var v3 float32 = ((3.14159 * 3.14159) - 3.14159)
    var v4 uint64 = 966
    var v17 float32 = (float32(-2.25) / 1.5)
    fmt.Println(v3, v4, v17)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint = 65535
        _ = fi18
        _ = v19
    }
    v2[1000] = ""
    fmt.Println(v17, v3, v17)
    var v20 uint8 = 127
    for fi21 := 0; fi21 < 16; fi21++ {
        v17 -= v17
        _ = fi21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 int32 = 1
            var v25 uint = 1
            _ = len(v1)
            var v26 int32 = (v24 | 2147483647)
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        switch "hello" {
        case "\n\t":
            if ((v20 ^ v20) == v20) {
                v17 += -2.25
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v27 := v2[2147483648]
                _ = v27
            }
        case "":
            v17 += 1.5
        case "foo bar":
            v28 := GenS1{F0: 644, F1: v20}
            _ = v28
        }
        _ = fi22
    }
    v17 -= v17
    if len(v1) > 0 {
        v1[0] = uint32(v4)
    }
    v29 := GenS1{F0: 0, F1: (v20 >> ((272) & 63))}
    v4 = 511
    if len(v1) > 0 {
        v1[0] = 2
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS1{F0: 2147483648, F1: (^(v20))}
        if (30 > -1) {
            delete(v2, 2147483648)
            var v32 float32 = float32(3.14159)
            v17 = ((3.14159 * 3.14159) * float32(float64((-2.25 * ((0.1 + -2.25) / v17)))))
            _ = v32
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                if (v3 != v3) {
                    v20 = (v20 % (((v20 | v20)) | 1))
                    var v34 float64 = -2.25
                    _ = v34
                } else {
                    v20 -= (v20 | (v20 & (v20 | 84)))
                }
                var v35 int = (fi30 / (((fi33 | int(v20))) | 1))
                _ = fi33
                _ = v35
            }
            v36 := GenS2{F0: 42, F1: ("" + "a"), F2: -372}
            for fi37 := 0; fi37 < 16; fi37++ {
                _ = len(v2)
                var v38 int16 = -861
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        _ = len(v2)
        _ = fi30
        _ = v31
    }
    v39 := GenS2{F0: -712, F1: "a", F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v39
    _ = v4
}
