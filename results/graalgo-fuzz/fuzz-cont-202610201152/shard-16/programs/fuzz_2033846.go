package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0(p0 int8, p1 string, p2 uint16) int32 {
    p2 += 2
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -799
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 128
    var v3 int = -479
    var v4 int16 = -1
    var v5 uint8 = (v2 * 128)
    var v6 float32 = (-2.25 + -2.25)
    v4 = (v4 ^ (v4 - 32767))
    switch -1 {
    case 128:
        v8 := v1[127]
        _ = v8
    case 2147483647:
        v5 -= 42
    case -78:
        v9 := []int{42, (v3 ^ v3), (v3 << ((1) & 63))}
        _ = v9
    default:
        var v10 int64 = 382
        _ = v10
    }
    fmt.Println(v3, v6)
    fmt.Println(v4)
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint64 = 150
        for fi13 := 0; fi13 < 8; fi13++ {
            v5 -= (v5 - (v5 + v2))
            _ = fi13
        }
        v14 := []int{v3}
        _ = fi11
        _ = v12
        _ = v14
    }
    var v15 uint32 = 2
    switch ((("a" + "") + ("foo bar" + "foo bar")) + ("" + "foo bar")) {
    case "foo bar":
        for fi16 := 0; fi16 < 19; fi16++ {
            if (false && (-804 < -2147483648)) {
                var v17 int = fi16
                for fi18 := 0; fi18 < 16; fi18++ {
                    v3 += v3
                    _ = fi18
                }
                _ = v17
            } else {
                switch (("foo bar" + "") + "") {
                case "a":
                    if true {
                        v5 += v5
                        var v19 int16 = v4
                        if (-128 == -46) {
                            v5 += 0
                            v19 = v4
                            if (!false) {
                                v19 = v4
                                var v20 uint32 = 940
                                if true {
                                    v21 := GenS1{F0: (4294967295 == 910), F1: 1}
                                    var v22 float64 = 1.5
                                    if (255 < 42) {
                                        v23 := GenS1{F0: true, F1: 2}
                                        v19 -= v4
                                        v24 := []int{(fi16 * v3)}
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v6 -= (v6 / v6)
                                        var v25 bool = ((!true) || true)
                                        v26 := GenS1{F0: (v22 != v22), F1: 1}
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v21
                                    _ = v22
                                }
                                _ = v20
                            } else {
                                v19 += v4
                                v6 = (float32((float64((v6 + v6)) - 3.14159)) + v6)
                                fmt.Println(v6, v19, v5)
                            }
                        }
                        _ = v19
                    }
                case "\n\t":
                    fmt.Println(v6, v15)
                case "foo bar":
                    v5 = (v5 | uint8(v15))
                }
                v5 += v5
                if false {
                    var v27 string = ("" + "foo bar")
                    v28 := GenS1{F0: true, F1: 120}
                    if false {
                        v29 := GenS1{F0: false, F1: 255}
                        _ = v29
                    } else {
                        v30 := []int{(fi16 ^ (v3 / (((v3 / ((v3) | 1))) | 1))), v3, 840, (v3 ^ (fi16 ^ (fi16 ^ fi16)))}
                        if (fi16 != (v3 & (fi16 / ((v3) | 1)))) {
                            v3 = fi16
                        } else {
                            v31 := GenS1{F0: true, F1: 167}
                            fmt.Println(fi16)
                            v5 = v5
                            _ = v31
                        }
                        _ = v30
                    }
                    _ = v27
                    _ = v28
                }
            }
            var v32 float64 = -2.25
            for fi33 := 0; fi33 < 8; fi33++ {
                v34 := GenS1{F0: (v3 >= (fi16 ^ v3)), F1: 971}
                _ = fi33
                _ = v34
            }
            _ = fi16
            _ = v32
        }
    }
    v6 = float32(1.5)
    v3 += 9223372036854775807
    v2 = v2
    v35 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
