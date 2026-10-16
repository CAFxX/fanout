package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 uint32 }
func genFn0() uint32 {
    v7 := []int{928, -585, 0}
    var v8 uint32 = 824
    fmt.Println(v8)
    _ = v7
    _ = v8
    return v8
}
func genFn1() uint16 {
    var v9 string = "foo bar"
    v10 := []int{42}
    var v11 float64 = float64(-2.25)
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    v1 := []int{127, 32768}
    v2 := make(map[string]int)
    v3 := []uint32{0, 563, 2147483648}
    v4 := make(map[string]string)
    var v5 uint = 149
    var v6 uint32 = 1
    v5 += (^(v5))
    if false {
        var v12 int64 = int64(v6)
        var v13 uint8 = 55
        _ = v12
        _ = v13
    }
    var v14 int32 = -304
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(v14, v14, v6)
        _ = fi15
    }
    v16 := GenS1{F0: 18446744073709551615}
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 uint32 = (v6 << ((v5) & 63))
        v6 += v6
        v19 := GenS2{F0: (^(v5)), F1: "a"}
        _ = fi17
        _ = v18
        _ = v19
    }
    v5 = (^(v5))
    v6 -= (v6 >> ((18446744073709551615) & 63))
    v6 = v6
    if (255 != 0) {
        if (!true) {
            v20 := GenS2{F0: 529, F1: "hello"}
            v21 := v4["\n\t"]
            v22 := GenS3{F0: 65535, F1: 42, F2: (v5 % ((v5) | 1))}
            _ = v20
            _ = v21
            _ = v22
        } else {
            if false {
                v14 -= v14
            } else {
                var v23 int16 = -766
                _ = len(v1)
                switch ("\n\t" + ("" + (("" + ("foo bar" + "\n\t")) + "\n\t"))) {
                case "":
                    var v24 bool = (0 > 127)
                    _ = v24
                case "\n\t":
                    for fi25 := 0; fi25 < 11; fi25++ {
                        fmt.Println(fi25)
                        _ = fi25
                    }
                default:
                    if (!false) {
                        v6 = v6
                        v14 = v14
                        for fi26 := 0; fi26 < 16; fi26++ {
                            var v27 uint = 65535
                            switch "" {
                            case "\n\t":
                                switch (fi26 & -363) {
                                case 32767:
                                    v23 += v23
                                }
                            case "a":
                                v6 += (v6 & genFn0())
                            default:
                                var v28 uint64 = 2
                                _ = v28
                            }
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        v29 := GenS4{F0: v6}
                        _ = v29
                    }
                }
                _ = v23
            }
            if ("a" >= "foo bar") {
                var v30 float32 = (1.5 * (((float32(0.1) * float32((1e10 / (1e10 / 1.5)))) / float32(((0.0 + 3.14159) - (1.5 / 1.0)))) / (3.14159 + (0.0 / 3.14159))))
                _ = v30
            }
        }
        delete(v2, ("hello" + ""))
    } else {
        fmt.Println(v6, v6)
    }
    v3 = append(v3, 42)
    for fi31 := 0; fi31 < 10; fi31++ {
        v14 = (^(v14))
        v5 -= v5
        var v32 uint32 = 42
        _ = fi31
        _ = v32
    }
    var v33 float64 = 1e10
    fmt.Println(v33)
    if true {
        switch (("\n\t" + ("" + ("" + "hello"))) + "hello") {
        case "a":
            switch "foo bar" {
            case "a":
                v1 = append(v1, -2147483648)
            case "\n\t":
                var v34 float32 = float32(0.1)
                _ = v34
            default:
                if len(v1) > 0 {
                    v1[0] = -728
                }
            }
        default:
            delete(v2, "hello")
        }
        v14 = 0
    }
    if ((!true) || false) {
        v35 := GenS2{F0: v5, F1: "\n\t"}
        v36 := GenS2{F0: v5, F1: "hello"}
        _ = v35
        _ = v36
    } else {
        fmt.Println(v6, v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
