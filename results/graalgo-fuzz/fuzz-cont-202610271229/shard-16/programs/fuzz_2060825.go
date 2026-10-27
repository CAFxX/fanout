package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 int32) string {
    if ((!(!true)) && (2 == 4294967295)) {
        v7 := []int{917, 32767, -2147483648, 1}
        _ = v7
    } else {
        v8 := []int{-1, 32767, 127, 663}
        if (2147483648 != 556) {
            _ = len(v8)
        } else {
            v8 = append(v8, -1)
            switch "" {
            case "":
                p0 = p0
            default:
                if (!(32767 == 970)) {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 int32 = (p0 << ((435) & 63))
                        p0 += v10
                        v11 := GenS1{F0: 0, F1: 119}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                } else {
                    var v12 int16 = 1
                    for fi13 := 0; fi13 < 20; fi13++ {
                        switch fi13 {
                        case 127:
                            v14 := GenS1{F0: 562, F1: 128}
                            _ = v14
                        case -766:
                            for fi15 := 0; fi15 < 12; fi15++ {
                                if true {
                                    var v16 uint16 = 42
                                    v17 := GenS2{F0: (!((false && (!true)) && false)), F1: 42}
                                    var v18 float32 = 1.5
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 uint8 = 128
                                    var v20 int64 = 0
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint16 = 32767
                                fmt.Println(v21, fi13)
                                _ = fi15
                                _ = v21
                            }
                        case -2147483648:
                            var v22 uint32 = 424
                            _ = v22
                        default:
                            if len(v8) > 0 {
                                v8[0] = (fi13 & fi13)
                            }
                        }
                        v23 := GenS2{F0: (-66 <= 42), F1: 1}
                        _ = fi13
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v12 = v12
                        _ = fi24
                    }
                    _ = v12
                }
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(fi25)
                v26 := GenS1{F0: 65535, F1: 128}
                if ("" > "") {
                    if len(v8) > 0 {
                        v8[0] = (fi25 / ((42) | 1))
                    }
                }
                _ = fi25
                _ = v26
            }
        }
        _ = v8
    }
    var v27 uint8 = 1
    var v28 int = -2147483648
    var v29 int32 = p0
    v27 = v27
    _ = p0
    _ = v27
    _ = v28
    _ = v29
    return (("a" + "\n\t") + "\n\t")
}
func genFn1(p0 int16, p1 int64) uint8 {
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS2{F0: true, F1: p1}
        v32 := GenS2{F0: (-128 < 100), F1: (p1 * (p1 / ((p1) | 1)))}
        _ = fi30
        _ = v31
        _ = v32
    }
    if (!true) {
        v33 := GenS2{F0: ((^(p0)) > 0), F1: (^(p1))}
        v34 := GenS1{F0: 4294967295, F1: 127}
        _ = v33
        _ = v34
    }
    var v35 float64 = -2.25
    var v36 bool = ((0 == 420) || true)
    _ = p0
    _ = p1
    _ = v35
    _ = v36
    return 255
}

func main() {
    var v1 bool = (!((("foo bar" + "\n\t") <= "\n\t") && (!(4294967295 > 0))))
    v2 := []string{"foo bar", ("hello" + "hello"), "foo bar"}
    v3 := make(map[uint32]string)
    v4 := []string{"a"}
    var v5 uint = 686
    var v6 float64 = float64((1.5 / (1e10 - (1.5 / 1.5))))
    for fi37 := 0; fi37 < 16; fi37++ {
        v38 := v3[65535]
        var v39 int64 = 0
        v38 = v38
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 string = "\n\t"
    v1 = false
    fmt.Println(v1)
    v5 = v5
    switch 4294967295 {
    case 4294967295:
        switch ((v40 + v40) + "hello") {
        case "foo bar":
            fmt.Println(v40, v6, v6)
        case "":
            v1 = ((v6 * ((0.1 * v6) - 1e10)) <= (((v6 * 0.1) + ((v6 * (v6 * ((((float64(0.1) - v6) * v6) * v6) * v6))) - -2.25)) * v6))
        case "hello":
            v1 = (!v1)
        default:
            v41 := v3[4294967295]
            _ = v41
        }
    }
    var v42 float32 = (3.14159 * 0.0)
    v1 = (!(v1 && true))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
