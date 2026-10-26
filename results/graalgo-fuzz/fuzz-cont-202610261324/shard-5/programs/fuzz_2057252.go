package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0() string {
    var v5 uint8 = 1
    v5 = v5
    v6 := []int{128, -9223372036854775808, 37}
    if len(v6) > 0 {
        v6[0] = -863
    }
    _ = v5
    _ = v6
    return ("foo bar" + "")
}

func main() {
    var v1 int64 = 0
    v2 := []int{}
    v3 := []uint32{2147483648, 42}
    v4 := make(map[uint32]string)
    v1 -= v1
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7)
        var v8 int32 = -2147483648
        var v9 uint8 = 2
        _ = fi7
        _ = v8
        _ = v9
    }
    if (56 == -128) {
        if true {
            for fi10 := 0; fi10 < 1; fi10++ {
                if (false || true) {
                    var v11 uint64 = 2
                    if true {
                        v11 = (v11 >> ((0) & 63))
                        v11 += 2
                    } else {
                        switch (("\n\t" + ("" + (("" + (("hello" + (("\n\t" + "\n\t") + "foo bar")) + ("" + "foo bar"))) + ""))) + ("a" + "")) {
                        case "hello":
                            v12 := GenS1{F0: 0, F1: 52, F2: 1}
                            _ = v12
                        case "":
                            _ = len(v3)
                        case "\n\t":
                            var v13 uint16 = 878
                            _ = v13
                        }
                    }
                    v14 := GenS2{F0: 2147483647, F1: -32768}
                    _ = v11
                    _ = v14
                } else {
                    v1 = v1
                    var v15 string = "a"
                    v16 := GenS1{F0: 18446744073709551615, F1: 92, F2: 78}
                    _ = v15
                    _ = v16
                }
                v1 -= -1
                _ = fi10
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 int64 = (v1 << ((0) & 63))
                _ = len(v3)
                _ = fi17
                _ = v18
            }
        }
        var v19 float32 = (float32(0.1) - 0.1)
        _ = v19
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = 0
    }
    if false {
        v1 = (v1 ^ 268)
        var v20 bool = (1 == 0)
        v1 += (v1 - (v1 - v1))
        _ = v20
    } else {
        switch ("\n\t" + "") {
        case "foo bar":
            switch "a" {
            case "":
                fmt.Println(v1)
            case "hello":
                v1 = v1
            case "a":
                fmt.Println(v1)
            default:
                v2 = append(v2, -9223372036854775808)
            }
        case "":
            for fi21 := 0; fi21 < 9; fi21++ {
                v1 -= 1
                switch fi21 {
                case -635:
                    v1 = (-(v1))
                case 0:
                    if (false && (255 >= 1)) {
                        fmt.Println(v1)
                    }
                }
                _ = fi21
            }
        }
        v1 = -1
    }
    if (!true) {
        v1 += -9223372036854775808
        delete(v4, 280)
        v1 -= (^(v1))
    }
    if (((1.5 - 0.0) * -2.25) < 3.14159) {
        v1 += v1
    } else {
        if true {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 float32 = float32((1e10 + -2.25))
                _ = fi22
                _ = v23
            }
            fmt.Println(v1)
        }
        v1 = (v1 * (v1 - v1))
    }
    var v24 float32 = (1e10 - -2.25)
    for fi25 := 0; fi25 < 18; fi25++ {
        v1 = (v1 | 349)
        _ = fi25
    }
    var v26 string = "a"
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 9; fi28++ {
            fmt.Println(fi28, fi27)
            var v29 string = (v26 + v26)
            _ = fi28
            _ = v29
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v30 := GenS2{F0: 1, F1: -795}
        _ = fi27
        _ = v30
    }
    v4[42] = "foo bar"
    var v31 int = 52
    var v32 int8 = 100
    var v33 uint16 = 877
    v33 -= v33
    var v34 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
