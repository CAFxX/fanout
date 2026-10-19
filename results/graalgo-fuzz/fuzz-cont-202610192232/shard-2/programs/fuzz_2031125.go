package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32; F2 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint, p2 bool) int32 {
    v7 := GenS1{F0: -1}
    var v8 int8 = 100
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{(fi10 >> (((p1 % ((p1) | 1))) & 63))}
            var v12 int32 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int = (fi9 & (fi9 + (fi9 / ((0) | 1))))
        _ = fi9
        _ = v13
    }
    v14 := GenS1{F0: -875}
    if (p2 && (!(0 != 0))) {
        for fi15 := 0; fi15 < 4; fi15++ {
            switch ("hello" + "") {
            case "\n\t":
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := []int{}
                    var v18 uint8 = 236
                    var v19 int16 = int16(fi15)
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            default:
                v20 := GenS1{F0: 42}
                _ = v20
            }
            var v21 uint64 = 4294967295
            for fi22 := 0; fi22 < 9; fi22++ {
                v21 += (v21 + (^(v21)))
                v21 = (^(v21))
                p0 = 567
                _ = fi22
            }
            _ = fi15
            _ = v21
        }
    } else {
        p2 = (true || p2)
        if ((p0 & (p0 + p0)) <= (p0 + (p0 / ((p0) | 1)))) {
            var v23 float64 = 0.0
            v24 := []int{128, 128}
            _ = v23
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    _ = v8
    return 0
}
func genFn1() int32 {
    var v25 int16 = 923
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := []int{fi26}
        _ = fi26
        _ = v27
    }
    if true {
        for fi28 := 0; fi28 < 4; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                if (!((float32(1.5) / 0.1) <= float32(float64(-2.25)))) {
                    v30 := GenS1{F0: -405}
                    var v31 int32 = 42
                    var v32 int32 = v31
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    var v33 int8 = -128
                    _ = v33
                }
                if (!false) {
                    v25 += 0
                } else {
                    if (42 < 255) {
                        switch (fi28 >> ((0) & 63)) {
                        case 32767:
                            fmt.Println(v25)
                        case 128:
                            v25 = 658
                        default:
                            v34 := GenS2{F0: ("\n\t" + "\n\t"), F1: -583, F2: -1}
                            _ = v34
                        }
                        fmt.Println(fi28, fi28)
                        var v35 float64 = (3.14159 / -2.25)
                        _ = v35
                    }
                    v25 = v25
                    v36 := GenS3{F0: -454, F1: 65535, F2: (v25 << ((945) & 63))}
                    _ = v36
                }
                _ = fi29
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := []int{fi37, fi37, 1, fi28}
                if len(v38) > 0 {
                    v38[0] = (^(fi28))
                }
                v38 = append(v38, (fi37 - fi28))
                _ = fi37
                _ = v38
            }
            if (!false) {
                var v39 int32 = 679
                _ = v39
            } else {
                var v40 float32 = (3.14159 + (0.0 + ((3.14159 + 0.0) / (1.5 - 1e10))))
                _ = v40
            }
            _ = fi28
        }
        fmt.Println(v25)
    } else {
        if (0.1 < 0.0) {
            var v41 uint8 = 1
            fmt.Println(v25, v41)
            v25 = v25
            _ = v41
        }
    }
    _ = v25
    return 1
}

func main() {
    var v1 int8 = 0
    var v2 float64 = 1e10
    var v3 int64 = int64(v1)
    var v4 int32 = -2147483648
    v5 := []string{"a", "foo bar", "", "a"}
    v6 := make(map[uint32]uint32)
    v1 = v1
    fmt.Println(v3, v4, v4)
    v2 += 1e10
    v1 += v1
    v2 = (((0.0 / (v2 * (v2 * ((v2 / 3.14159) * v2)))) * v2) / 1e10)
    for fi42 := 0; fi42 < 6; fi42++ {
        v43 := GenS3{F0: (fi42 * 2147483647), F1: 42, F2: 32767}
        v2 = -2.25
        _ = fi42
        _ = v43
    }
    v3 += 549
    switch 2147483647 {
    case -2147483648:
        v1 += 42
    case -973:
        v3 -= 9223372036854775807
    }
    fmt.Println(v2)
    v44 := GenS4{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
