package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 bool) uint8 {
    v4 := []int{76, 291, 42}
    var v5 uint64 = 4294967295
    if (((1e10 * float32(((1.5 * 1e10) * ((0.0 / 1.0) - (((-2.25 / 1.5) / float64(3.14159)) * 1e10))))) == float32(0.1)) || (!p0)) {
        var v6 uint32 = 65535
        p0 = p0
        _ = v6
    } else {
        switch 128 {
        case 42:
            for fi7 := 0; fi7 < 10; fi7++ {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                var v8 int16 = -32768
                _ = fi7
                _ = v8
            }
        }
    }
    v9 := GenS1{F0: 42, F1: -113}
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS1{F0: 12, F1: 0}
        var v12 float64 = 0.0
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 bool = false
            _ = fi13
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v9
    return 128
}
func genFn1(p0 uint32, p1 uint64) string {
    var v15 bool = (42 <= int64(p0))
    v16 := []int{1, 0}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return (("foo bar" + ("a" + "foo bar")) + "")
}

func main() {
    var v1 bool = (!(-1 == -2147483648))
    var v2 uint = 2
    var v3 uint32 = 4294967295
    v1 = v1
    v1 = v1
    fmt.Println(v1)
    v17 := []int{-426, -168, -9223372036854775808, 32767}
    v1 = true
    v3 = v3
    if (v1 && ((1.5 != (float32((0.1 - 0.1)) - -2.25)) || (-758 <= 1))) {
        fmt.Println(v1, v2)
        for fi18 := 0; fi18 < 2; fi18++ {
            if v1 {
                _ = len(v17)
            }
            v19 := GenS1{F0: -128, F1: -128}
            var v20 int8 = 127
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v21 int16 = 0
    v2 = 660
    v3 = v3
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v17) > 0 {
            v17[0] = fi22
        }
        _ = fi22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 string = "\n\t"
            var v26 uint32 = (v3 >> (((v2 << ((v2) & 63))) & 63))
            v26 -= (v3 % (((^(v3))) | 1))
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS1{F0: -58, F1: -128}
            var v29 int = 32768
            if true {
                var v30 int64 = 9223372036854775807
                var v31 bool = false
                var v32 uint32 = 0
                _ = v30
                _ = v31
                _ = v32
            } else {
                v3 = (^(v3))
                if v1 {
                    var v33 int32 = -739
                    v34 := GenS1{F0: -1, F1: 0}
                    switch "foo bar" {
                    case "hello":
                        if ((true && v1) && false) {
                            v35 := GenS1{F0: 42, F1: -1}
                            _ = v35
                        }
                    case "":
                        var v36 uint16 = 42
                        _ = v36
                    case "\n\t":
                        switch fi23 {
                        case 796:
                            var v37 int32 = (v33 % ((-953) | 1))
                            _ = v37
                        case -2147483648:
                            var v38 int64 = 42
                            _ = v38
                        case 4294967295:
                            switch (v29 - 4294967295) {
                            case -9223372036854775808:
                                v3 = 190
                            default:
                                v17 = append(v17, int(v3))
                            }
                        }
                    default:
                        v1 = v1
                    }
                    v39 := GenS1{F0: -1, F1: -1}
                    _ = v33
                    _ = v34
                    _ = v39
                }
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
