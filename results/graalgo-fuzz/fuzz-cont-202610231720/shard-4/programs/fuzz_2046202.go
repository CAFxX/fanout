package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint }
func genFn0() int32 {
    v7 := GenS1{F0: 0}
    if (1 < 65535) {
        var v8 int = 4294967295
        if ((0.0 / 1.0) > 1.5) {
            fmt.Println(v8)
        }
        _ = v8
    }
    _ = v7
    return 42
}
func genFn1() uint32 {
    var v9 uint = 4294967295
    for fi10 := 0; fi10 < 14; fi10++ {
        v9 += (^(v9))
        v11 := GenS1{F0: 849}
        _ = fi10
        _ = v11
    }
    if false {
        fmt.Println(v9)
        v12 := []int{127, -846, 921}
        var v13 string = "foo bar"
        _ = v12
        _ = v13
    }
    var v14 uint16 = 0
    v9 += (v9 ^ v9)
    _ = v14
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 2
    var v2 int8 = 54
    v3 := make(map[string]uint32)
    var v4 uint64 = 910
    v5 := make(map[string]int)
    v6 := make(map[string]string)
    v1 = v1
    if (v2 <= v2) {
        v1 -= (v1 % ((v1) | 1))
    } else {
        switch ("" + (("a" + ((("" + "\n\t") + ("a" + "\n\t")) + (("foo bar" + "foo bar") + ""))) + "hello")) {
        case "\n\t":
            v15 := []int{4294967295, 773, 268, -1}
            _ = v15
        case "a":
            v1 = (^(v1))
        case "hello":
            v16 := GenS1{F0: -1}
            _ = v16
        default:
            v1 -= genFn1()
        }
        switch "\n\t" {
        case "foo bar":
            var v17 int = -323
            _ = v17
        default:
            fmt.Println(v2, v4)
        }
        if (((431 > 222) || (!(42 <= -2147483648))) || true) {
            if ((v2 >> ((4294967295) & 63)) >= v2) {
                if (255 != 0) {
                    var v18 int32 = 0
                    v4 = v4
                    _ = v18
                }
                if false {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v4 += (v4 & v4)
                        _ = fi19
                    }
                    var v20 int16 = 1000
                    _ = v20
                } else {
                    v21 := GenS2{F0: 255, F1: (v4 << ((18446744073709551615) & 63))}
                    for fi22 := 0; fi22 < 17; fi22++ {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 uint8 = 243
                            fmt.Println(v2)
                            v25 := GenS3{F0: 285, F1: -32768, F2: "foo bar"}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = fi22
                    }
                    _ = v21
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint = 65535
                    v28 := GenS2{F0: v27, F1: uint64(v27)}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            var v29 string = (("a" + ((("a" + (("\n\t" + "foo bar") + "foo bar")) + ("foo bar" + ("a" + "\n\t"))) + ("a" + ""))) + ((("foo bar" + (("" + ("" + "")) + "a")) + "foo bar") + "\n\t"))
            _ = v29
        } else {
            v30 := GenS4{F0: -9223372036854775808, F1: genFn0()}
            fmt.Println(v4)
            v4 += v4
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := []int{(fi31 ^ fi31), (fi31 | fi31)}
        v2 = v2
        v1 -= (v1 >> ((42) & 63))
        _ = fi31
        _ = v32
    }
    v4 += 2
    v4 -= 18446744073709551615
    v2 -= (v2 >> ((65535) & 63))
    if (!(2 >= 4294967295)) {
        v33 := GenS5{F0: (v4 / ((v4) | 1)), F1: 557, F2: 2}
        var v34 int16 = 32767
        _ = v33
        _ = v34
    }
    v35 := []int{}
    v2 = (v2 & (^(v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
