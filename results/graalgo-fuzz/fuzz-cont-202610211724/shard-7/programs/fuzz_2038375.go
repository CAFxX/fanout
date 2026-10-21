package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint; F1 int; F2 int32 }
func genFn0(p0 int) uint16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 0, F1: ("a" + "a")}
        var v7 string = ("" + (("" + ("\n\t" + "hello")) + ("a" + "hello")))
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 514, F1: (p0 - p0), F2: -2147483648}
    var v9 int = p0
    v10 := []int{v9, v9}
    fmt.Println(v9, p0)
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 1
    v3 := make(map[string]int)
    var v4 string = ("a" + (("\n\t" + "\n\t") + "a"))
    var v11 uint32 = uint32(v2)
    v4 = "foo bar"
    v12 := v3[((v4 + v4) + "\n\t")]
    v13 := GenS2{F0: 42, F1: v12, F2: 1}
    v1 = 1e10
    for fi14 := 0; fi14 < 6; fi14++ {
        switch (fi14 ^ fi14) {
        case 32767:
            v11 = 2
        case 9223372036854775807:
            v3[v4] = v12
        case -964:
            if (false || (409 <= uint16(v2))) {
                switch fi14 {
                case -411:
                    v15 := GenS1{F0: 1, F1: "foo bar"}
                    _ = v15
                case -742:
                    v1 += v1
                case -797:
                    var v16 int32 = -607
                    _ = v16
                default:
                    var v17 int64 = -313
                    _ = v17
                }
                v12 = (v12 | (fi14 << ((18446744073709551615) & 63)))
            } else {
                switch v12 {
                case 4294967295:
                    switch ("\n\t" + "") {
                    case "foo bar":
                        if ((255 < 128) && false) {
                            switch (-(v12)) {
                            case 143:
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    if false {
                                        fmt.Println(v11, v1)
                                        if (!true) {
                                            v19 := GenS2{F0: 255, F1: fi14, F2: -1}
                                            v1 = v1
                                            _ = v19
                                        }
                                        var v20 float32 = -2.25
                                        v21 := GenS1{F0: 1000, F1: v4}
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = fi18
                                }
                            case -853:
                                var v22 uint16 = 32767
                                _ = v22
                            default:
                                delete(v3, v4)
                            }
                            fmt.Println(v4, v11)
                        } else {
                            var v23 int32 = -1
                            _ = v23
                        }
                    }
                default:
                    v24 := GenS1{F0: 1000, F1: (v4 + v4)}
                    _ = v24
                }
                v25 := GenS1{F0: 32767, F1: v4}
                _ = v25
            }
        }
        var v26 uint32 = v11
        v12 += v12
        _ = fi14
        _ = v26
    }
    fmt.Println(v12, v1, v4)
    var v27 string = v4
    for fi28 := 0; fi28 < 3; fi28++ {
        v29 := GenS1{F0: -32768, F1: "hello"}
        _ = fi28
        _ = v29
    }
    v2 -= (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
