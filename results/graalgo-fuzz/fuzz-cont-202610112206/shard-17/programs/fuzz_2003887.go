package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint8 }
func genFn0() int32 {
    v5 := GenS1{F0: 120, F1: 42}
    if false {
        v6 := GenS2{F0: 4294967295, F1: 2, F2: 42}
        var v7 float32 = 1e10
        _ = v6
        _ = v7
    }
    _ = v5
    return 1
}
func genFn1() string {
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            if (!false) {
                for fi9 := 0; fi9 < 1; fi9++ {
                    if true {
                        var v10 uint64 = 18446744073709551615
                        v10 = (^(v10))
                        v11 := GenS3{F0: v10, F1: 65535, F2: (v10 - v10)}
                        _ = v10
                        _ = v11
                    }
                    var v12 uint16 = 255
                    var v13 int8 = 0
                    _ = fi9
                    _ = v12
                    _ = v13
                }
            }
            _ = fi8
        }
    }
    var v14 int64 = -876
    var v15 int = 1
    _ = v14
    _ = v15
    return "foo bar"
}

func main() {
    v1 := []uint32{65535, 2147483648, 42}
    var v2 uint32 = 753
    var v3 uint16 = 255
    v4 := []int{18, -1}
    v2 = v2
    if (v2 != v2) {
        for fi16 := 0; fi16 < 16; fi16++ {
            switch ("a" + ("hello" + "\n\t")) {
            case "":
                if (!false) {
                    v17 := GenS4{F0: 4294967295, F1: v3, F2: 42}
                    _ = v17
                }
            case "a":
                v4 = append(v4, 558)
            case "\n\t":
                for fi18 := 0; fi18 < 19; fi18++ {
                    v2 -= v2
                    _ = fi18
                }
            }
            _ = fi16
        }
    } else {
        var v19 uint64 = 0
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int8 = 1
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: 65535, F1: v19}
        _ = v19
        _ = v22
    }
    if true {
        fmt.Println(v3, v2)
        v2 = (v2 - (v2 + v2))
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int64 = -559
            _ = fi23
            _ = v24
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        if (128 <= 0) {
            _ = len(v1)
            v2 = (v2 % ((2147483647) | 1))
            v25 := GenS3{F0: 2, F1: v3, F2: 917}
            _ = v25
        }
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    if (42 != -1) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch 483 {
        case 128:
            v3 -= 1
        case 127:
            v2 = v2
        case -2147483648:
            for fi26 := 0; fi26 < 3; fi26++ {
                if false {
                    if true {
                        var v27 uint = 42
                        _ = v27
                    } else {
                        fmt.Println(v2, v2)
                        var v28 int16 = -1
                        var v29 uint16 = v3
                        _ = v28
                        _ = v29
                    }
                    v3 = (v3 ^ v3)
                    v3 -= v3
                }
                switch (fi26 * fi26) {
                case 2147483647:
                    for fi30 := 0; fi30 < 18; fi30++ {
                        switch fi30 {
                        case 3:
                            v2 += (^(v2))
                        case 692:
                            fmt.Println(v2, fi26)
                        case 2147483647:
                            if false {
                                fmt.Println(fi30, fi26)
                            }
                        default:
                            var v31 uint32 = 2147483648
                            _ = v31
                        }
                        _ = fi30
                    }
                case 32767:
                    fmt.Println(fi26, v3)
                default:
                    switch (fi26 << ((26) & 63)) {
                    case -1:
                        var v32 int32 = 1000
                        _ = v32
                    case -392:
                        v3 += (v3 & 65535)
                    case 42:
                        for fi33 := 0; fi33 < 4; fi33++ {
                            v3 -= 32767
                            var v34 bool = ((true || (-184 > -1)) || true)
                            v1 = append(v1, (^(v2)))
                            v3 = v3
                            _ = fi33
                            _ = v34
                        }
                    }
                }
                _ = fi26
            }
        }
    }
    v2 = (v2 & v2)
    var v35 uint32 = (v2 & 2147483648)
    switch 42 {
    case 490:
        switch (("a" + ((("a" + "a") + "hello") + "a")) + (("\n\t" + ("hello" + "a")) + "hello")) {
        case "a":
            var v36 int16 = -206
            _ = v36
        default:
            _ = len(v4)
        }
    }
    var v37 float64 = ((-2.25 - 0.0) * (-2.25 / float64(1e10)))
    v35 -= 776
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v37 = (v37 + float64(float32(v37)))
    v37 = v37
    switch (("" + ("a" + (("" + "") + ""))) + "") {
    case "foo bar":
        fmt.Println(v35, v37, v3)
    case "hello":
        var v38 float32 = -2.25
        _ = v38
    case "\n\t":
        var v39 int8 = -1
        _ = v39
    default:
        v35 -= (^(v2))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
}
