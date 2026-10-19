package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0() int16 {
    v4 := []int{}
    v5 := GenS1{F0: 603, F1: 100, F2: 105}
    var v6 string = "hello"
    _ = v4
    _ = v5
    _ = v6
    return -416
}
func genFn1() uint64 {
    var v7 string = "hello"
    v8 := GenS1{F0: 42, F1: -128, F2: -128}
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    var v3 uint64 = 4294967295
    v3 = (^(v3))
    v9 := []int{-2147483648, 42, 62, -302}
    switch 806 {
    case 906:
        v3 = 42
    case 856:
        var v10 uint64 = 146
        _ = v10
    case -2147483648:
        for fi11 := 0; fi11 < 7; fi11++ {
            if ((fi11 / (((fi11 & fi11)) | 1)) >= (fi11 >> ((42) & 63))) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v3 = (v3 & (v3 >> ((4294967295) & 63)))
                    _ = fi12
                }
                if ((1.5 > -2.25) && ((!true) && true)) {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        switch (fi11 - fi13) {
                        case -507:
                            v14 := GenS1{F0: 42, F1: 42, F2: v1}
                            _ = v14
                        case -284:
                            var v15 uint16 = 65535
                            _ = v15
                        }
                        var v16 int16 = 1
                        _ = fi13
                        _ = v16
                    }
                    var v17 uint16 = 284
                    _ = v17
                } else {
                    switch (fi11 - fi11) {
                    case 2147483647:
                        fmt.Println(v3)
                    default:
                        var v18 uint16 = 32767
                        _ = v18
                    }
                }
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 string = "\n\t"
                v21 := GenS1{F0: int32(v1), F1: (v1 ^ 29), F2: (v1 ^ 33)}
                fmt.Println(v3)
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 42, F1: int8(v3), F2: 88}
            _ = fi11
            _ = v22
        }
    }
    v1 += v1
    fmt.Println(v3)
    v1 += v1
    v1 += v1
    switch "a" {
    case "a":
        v1 = v1
    case "\n\t":
        switch 42 {
        case -1:
            for fi23 := 0; fi23 < 3; fi23++ {
                v1 = -1
                fmt.Println(v1, fi23)
                _ = fi23
            }
        case 9223372036854775807:
            v24 := v2[("\n\t" + "hello")]
            _ = v24
        default:
            v3 -= (v3 >> ((4294967295) & 63))
        }
    case "foo bar":
        v3 -= genFn1()
    }
    if false {
        v25 := GenS1{F0: 2147483647, F1: (-(v1)), F2: (v1 - 127)}
        _ = v25
    } else {
        var v26 float64 = 1e10
        var v27 int32 = -2147483648
        _ = v26
        _ = v27
    }
    var v28 bool = (249 >= 892)
    delete(v2, "\n\t")
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int8 = (v1 & (v1 / (((v1 << ((654) & 63))) | 1)))
        switch fi29 {
        case -1:
            var v31 float32 = float32(-2.25)
            _ = v31
        case 127:
            if false {
                if true {
                    var v32 int16 = -768
                    _ = v32
                } else {
                    var v33 int64 = 42
                    _ = v33
                }
                v1 = v1
            }
        case -362:
            v1 -= v30
        default:
            var v34 uint16 = 412
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    v1 -= (v1 + (v1 / ((-54) | 1)))
    v1 = v1
    v35 := GenS2{F0: -32768, F1: (v1 ^ (v1 % ((v1) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v9
}
