package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint8; F2 int32 }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
func genFn0() uint64 {
    var v7 bool = true
    v8 := []int{}
    fmt.Println(v7)
    var v9 int64 = -9223372036854775808
    v9 += (v9 % ((v9) | 1))
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 int32, p1 uint8, p2 uint16) uint32 {
    v10 := GenS1{F0: 18446744073709551615}
    v11 := GenS1{F0: 255}
    p0 = p0
    v12 := []int{}
    if ("hello" < ("" + (("a" + ("" + "hello")) + "hello"))) {
        fmt.Println(p2, p2)
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                if len(v12) > 0 {
                    v12[0] = fi13
                }
                _ = fi14
            }
            fmt.Println(p2, p0)
            _ = fi13
        }
    } else {
        var v15 int32 = p0
        if (4294967295 == 352) {
            var v16 int8 = 105
            for fi17 := 0; fi17 < 8; fi17++ {
                if ("\n\t" <= ("hello" + "")) {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if true {
                            var v19 int16 = 97
                            var v20 uint32 = 0
                            _ = v19
                            _ = v20
                        } else {
                            if ("foo bar" != "hello") {
                                _ = len(v12)
                                var v21 uint32 = 2
                                _ = v21
                            }
                            switch fi18 {
                            case -1:
                                p0 = (-(p0))
                            }
                        }
                        _ = fi18
                    }
                }
                _ = fi17
            }
            var v22 int32 = 1
            _ = v16
            _ = v22
        }
        v15 = -72
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    var v1 uint8 = 128
    var v2 int32 = -2147483648
    var v3 uint32 = 2147483648
    var v4 int = -9223372036854775808
    v5 := make(map[string]bool)
    var v6 int = (v4 << ((2) & 63))
    v23 := GenS2{F0: 9223372036854775807}
    v3 += uint32(v1)
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := []int{1, fi24, v6, v6}
            _ = len(v5)
            v1 += v1
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            fmt.Println(fi24)
            _ = fi27
        }
        var v28 uint8 = v1
        _ = fi24
        _ = v28
    }
    var v29 uint = 4294967295
    v30 := []int{128}
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 bool = true
        v29 = (^(v29))
        v33 := v5["foo bar"]
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int = (v4 + v6)
    if (1e10 < 1e10) {
        fmt.Println(v4, v1)
        _ = len(v30)
    } else {
        v5[""] = false
        v35 := GenS3{F0: v2, F1: 2, F2: (v2 & 42)}
        var v36 float32 = (((3.14159 + ((0.0 * (((0.1 - -2.25) + 0.0) + ((1.5 + ((-2.25 * 0.0) - 1.5)) + 1.5))) + 0.1)) / float32(-2.25)) / float32((1.5 + ((1e10 + -2.25) * (((((float64((3.14159 / (((0.1 / 0.1) * 3.14159) / 3.14159))) * (0.1 - ((1e10 - 3.14159) * (((3.14159 - 1.5) + (1.5 / float64(0.1))) / -2.25)))) * 0.1) * 1e10) / 1.5) + 0.0)))))
        _ = v35
        _ = v36
    }
    v6 = 128
    v2 += v2
    fmt.Println(v2)
    var v37 int16 = -1
    v1 -= v1
    switch "" {
    case "a":
        _ = len(v5)
    case "foo bar":
        var v38 int64 = -1
        _ = v38
    case "":
        _ = len(v30)
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 int64 = 969
        var v41 string = ("a" + "")
        v42 := GenS4{F0: v2, F1: 4294967295, F2: (^(v6))}
        _ = fi39
        _ = v40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
