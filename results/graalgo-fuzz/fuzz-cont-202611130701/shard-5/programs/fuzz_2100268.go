package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0() uint {
    if (!true) {
        if (!(false || (65535 <= 1))) {
            var v7 int16 = 0
            switch -9223372036854775808 {
            case 573:
                if (!false) {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        var v9 uint32 = 4294967295
                        v9 -= 11
                        v10 := []int{fi8, fi8, -282}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    fmt.Println(v7)
                }
            default:
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 float32 = -2.25
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 float64 = float64(3.14159)
                        var v15 int8 = 42
                        var v16 uint = 501
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = fi11
                    _ = v12
                }
            }
            _ = v7
        } else {
            v17 := []int{506, 417, -76}
            _ = v17
        }
        switch ("foo bar" + ((("" + "\n\t") + ("\n\t" + "foo bar")) + "\n\t")) {
        case "\n\t":
            var v18 float64 = 0.1
            _ = v18
        }
        switch "" {
        case "a":
            v19 := GenS1{F0: 14, F1: -292, F2: "foo bar"}
            _ = v19
        }
    }
    v20 := []int{9223372036854775807, 42}
    var v21 int32 = -844
    if false {
        var v22 int = 491
        v21 -= 2147483647
        _ = v22
    }
    v20 = append(v20, 4294967295)
    _ = v20
    _ = v21
    return 2
}

func main() {
    v1 := []string{}
    var v2 int16 = 0
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    v5 := []string{"", (("" + "") + ("\n\t" + ("a" + ""))), ""}
    var v6 uint64 = 590
    v2 -= v2
    var v23 uint32 = 178
    switch 1 {
    case -647:
        fmt.Println(v6)
    case -94:
        _ = len(v5)
    default:
        v4[(v23 - (v23 | 0))] = v23
    }
    var v24 int8 = 42
    for fi25 := 0; fi25 < 10; fi25++ {
        v24 = (v24 - 0)
        var v26 int8 = v24
        if false {
            v24 = -1
            if (792 == genFn0()) {
                var v27 int16 = 321
                var v28 uint32 = 0
                v2 = v2
                _ = v27
                _ = v28
            } else {
                v6 -= (^(v6))
                v29 := GenS1{F0: (v23 | (v23 << ((42) & 63))), F1: -711, F2: "foo bar"}
                _ = v29
            }
            var v30 int32 = -838
            _ = v30
        }
        _ = fi25
        _ = v26
    }
    v23 = (v23 + v23)
    v31 := GenS2{F0: 0, F1: (v24 + v24)}
    var v32 int = 32767
    fmt.Println(v32, v32)
    var v33 int16 = (v2 | (v2 & v2))
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint64 = 0
            v1 = append(v1, "hello")
            var v37 int32 = -1
            _ = fi35
            _ = v36
            _ = v37
        }
        v33 = v33
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi34
    }
    v38 := GenS1{F0: (v23 >> ((0) & 63)), F1: 1, F2: "hello"}
    var v39 uint8 = 128
    var v40 string = ("" + ("" + ""))
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
