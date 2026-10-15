package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8) uint32 {
    var v7 string = "foo bar"
    v8 := GenS1{F0: 127, F1: 100, F2: (("foo bar" + ((v7 + "") + ("\n\t" + (v7 + "a")))) + (v7 + v7))}
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    v4 := []int{1, 1, 32767}
    v5 := make(map[string]string)
    var v6 uint32 = 4294967295
    v9 := GenS2{F0: 1}
    fmt.Println(v2)
    if (0.1 != 1.5) {
        if ((4294967295 < 126) || false) {
            v2 = v2
        }
        v2 = (v2 - v2)
    }
    var v10 int64 = -1
    v6 -= v6
    for fi11 := 0; fi11 < 13; fi11++ {
        v10 = v10
        _ = fi11
    }
    v12 := GenS2{F0: 0}
    v13 := GenS1{F0: 4294967295, F1: v2, F2: "\n\t"}
    for fi14 := 0; fi14 < 2; fi14++ {
        if true {
            var v15 int32 = 2147483647
            _ = v15
        } else {
            v6 = v6
            v16 := GenS2{F0: 0}
            v17 := v5[""]
            _ = v16
            _ = v17
        }
        if ((v2 << ((4294967295) & 63)) < 1) {
            v4 = append(v4, 32767)
            v10 = v10
        } else {
            var v18 float64 = (0.0 / 1.5)
            switch -2147483648 {
            case 9223372036854775807:
                if (114 <= 1) {
                    v6 -= (v6 * 48)
                    if (!(!((v6 | (v6 | (v6 << ((255) & 63)))) < 714))) {
                        var v19 uint64 = 430
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v20 := GenS2{F0: int32(v2)}
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                switch fi21 {
                                case -9223372036854775808:
                                    if (!("a" > "hello")) {
                                        var v22 uint16 = 453
                                        _ = v22
                                    }
                                }
                                v23 := GenS2{F0: int32(v6)}
                                if len(v4) > 0 {
                                    v4[0] = (^(fi21))
                                }
                                _ = fi21
                                _ = v23
                            }
                        } else {
                            var v24 uint32 = genFn0((v2 | (v2 / (((v2 / ((-21) | 1))) | 1))))
                            _ = v24
                        }
                    }
                }
            default:
                for fi25 := 0; fi25 < 5; fi25++ {
                    v2 = v2
                    if (false && false) {
                        v26 := GenS1{F0: fi14, F1: (v2 % ((-128) | 1)), F2: ""}
                        if len(v4) > 0 {
                            v4[0] = fi14
                        }
                        _ = v26
                    }
                    switch 32767 {
                    case 32768:
                        v27 := GenS1{F0: (fi25 | fi14), F1: v2, F2: (("foo bar" + "hello") + "\n\t")}
                        _ = v27
                    }
                    _ = fi25
                }
            }
            v10 = (v10 >> ((872) & 63))
            _ = v18
        }
        v28 := GenS1{F0: fi14, F1: (v2 + (v2 / (((v2 % ((-94) | 1))) | 1))), F2: "a"}
        _ = fi14
        _ = v28
    }
    switch 42 {
    case 2147483647:
        v10 += v10
    case 200:
        if (true && (true && (0.0 >= 1.5))) {
            var v29 uint = 4294967295
            _ = v29
        }
    case 697:
        v10 = v10
    }
    fmt.Println(v10, v2)
    switch ((("hello" + "\n\t") + "hello") + ("" + "foo bar")) {
    case "a":
        v2 = 42
    case "":
        if false {
            v2 = 42
            v2 -= v2
        }
    default:
        if false {
            v6 -= v6
            v2 -= (v2 / ((0) | 1))
        }
    }
    v30 := v1[369]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
