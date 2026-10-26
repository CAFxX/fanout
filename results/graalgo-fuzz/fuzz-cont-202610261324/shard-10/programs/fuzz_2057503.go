package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int8 }
func genFn0() int {
    var v7 uint16 = 1
    v8 := GenS1{F0: 0, F1: 4294967295, F2: "hello"}
    if (("\n\t" + "foo bar") >= ("" + (("a" + "hello") + "foo bar"))) {
        if false {
            v7 += v7
            for fi9 := 0; fi9 < 5; fi9++ {
                if false {
                    for fi10 := 0; fi10 < 12; fi10++ {
                        fmt.Println(fi10)
                        if (false && true) {
                            fmt.Println(fi10, fi9)
                            v11 := []int{}
                            _ = v11
                        } else {
                            v7 = (^(v7))
                            switch (fi9 | 0) {
                            case 42:
                                v12 := GenS1{F0: 2147483647, F1: 1, F2: ""}
                                _ = v12
                            case 907:
                                var v13 int8 = -118
                                _ = v13
                            }
                            var v14 uint8 = 2
                            _ = v14
                        }
                        _ = fi10
                    }
                }
                if ((((float32((0.1 * 1.5)) - float32(3.14159)) * (0.1 + 0.0)) * 1e10) != 0.0) {
                    var v15 uint8 = 2
                    _ = v15
                }
                _ = fi9
            }
            fmt.Println(v7)
        }
    } else {
        var v16 int32 = -1
        v16 += v16
        _ = v16
    }
    var v17 int32 = -1
    _ = v17
    _ = v7
    _ = v8
    return -828
}

func main() {
    var v1 int32 = -2147483648
    v2 := []string{"foo bar", "a"}
    v3 := make(map[uint32]int)
    var v4 bool = (!(1e10 == (-2.25 / 1.5)))
    var v5 uint32 = 1
    var v6 float64 = 1.5
    fmt.Println(v5, v6)
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS1{F0: (v5 - (v5 ^ v5)), F1: 526, F2: "hello"}
            if (!v4) {
                if (v4 && false) {
                    v21 := GenS1{F0: v5, F1: 41, F2: "hello"}
                    fmt.Println(v4, fi18, v1)
                    _ = v21
                } else {
                    if (!false) {
                        var v22 uint = 897
                        v23 := GenS1{F0: 796, F1: 4294967295, F2: "\n\t"}
                        var v24 float32 = (1.5 + 3.14159)
                        delete(v3, (v5 | (v5 ^ (v5 + (v5 * 1)))))
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
                if (!(-128 > -128)) {
                    v25 := GenS1{F0: v5, F1: 4294967295, F2: "hello"}
                    v26 := GenS2{F0: 100}
                    _ = v25
                    _ = v26
                }
                _ = len(v3)
            }
            v1 -= -1
            _ = fi19
            _ = v20
        }
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint64 = 66
            _ = fi27
            _ = v28
        }
        v6 = (((float64((((1.5 * 1e10) * 0.0) - (((float32(0.0) / -2.25) / (float32((v6 - 1.5)) / 3.14159)) - 1e10))) * v6) - v6) - (v6 * (-2.25 * v6)))
        _ = fi18
    }
    var v29 uint8 = 42
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS2{F0: -63}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "foo bar" {
        case "\n\t":
            v29 = 42
        case "foo bar":
            var v32 float64 = v6
            _ = v32
        default:
            var v33 float64 = v6
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    v6 += 0.0
    v34 := GenS2{F0: 0}
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v3[(v5 ^ v5)] = (fi35 & (fi35 >> ((42) & 63)))
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
