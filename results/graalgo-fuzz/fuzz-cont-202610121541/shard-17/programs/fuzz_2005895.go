package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 uint, p1 uint8) uint32 {
    p0 -= 411
    switch 32767 {
    case -1:
        for fi4 := 0; fi4 < 1; fi4++ {
            var v5 uint16 = 550
            _ = fi4
            _ = v5
        }
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 686
    var v3 int32 = -2147483648
    for fi6 := 0; fi6 < 17; fi6++ {
        if (33 == 100) {
            switch fi6 {
            case 1:
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 int64 = 429
                    fmt.Println(v8)
                    _ = fi7
                    _ = v8
                }
            case -151:
                v2 = 42
            }
            var v9 uint8 = 42
            _ = v9
        } else {
            if true {
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch "hello" {
                    case "a":
                        v2 = (v2 << ((255) & 63))
                    }
                    if (1 >= 0) {
                        fmt.Println(v2)
                        if (!true) {
                            var v11 float32 = (((1e10 / ((0.0 + ((float32(3.14159) + 1.5) + 0.0)) - -2.25)) / float32(((3.14159 / 3.14159) - (0.0 + 1.5)))) - (0.0 + 0.1))
                            var v12 int64 = -9223372036854775808
                            if ("a" > ("" + "")) {
                                _ = len(v1)
                            }
                            v13 := []int{}
                            v11 += v11
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: 2, F1: 28}
                        _ = v14
                    } else {
                        if ((v3 & 635) < (v3 * (-(v3)))) {
                            v2 -= (v2 + 255)
                            v3 += (v3 << ((255) & 63))
                        }
                        v2 += 77
                        if ((true && true) && true) {
                            fmt.Println(fi10)
                            if (true || false) {
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v2 += (^(v2))
                                    default:
                                        v2 = 487
                                        var v15 bool = (true && true)
                                        v15 = (71 >= 111)
                                        _ = v15
                                    }
                                }
                                fmt.Println(v2, fi10)
                            }
                            v16 := []int{fi10, fi10, (fi10 * (fi10 * fi10))}
                            _ = v16
                        } else {
                            delete(v1, "\n\t")
                            if false {
                                if (-1 < 42) {
                                    var v17 int = fi10
                                    fmt.Println(fi6, fi6, v17)
                                    switch (fi10 - -629) {
                                    case 127:
                                        var v18 bool = (1 >= -128)
                                        _ = v18
                                    default:
                                        v3 += (v3 ^ (v3 ^ v3))
                                    }
                                    _ = v17
                                }
                                if (890 < 0) {
                                    var v19 float64 = ((1.5 / 1e10) * 0.1)
                                    if (!true) {
                                        v20 := GenS1{F0: 0, F1: (v3 ^ (^(v3)))}
                                        var v21 uint64 = 0
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = v19
                                } else {
                                    var v22 int32 = -214
                                    var v23 float64 = 1.5
                                    _ = v22
                                    _ = v23
                                }
                            }
                            if (!true) {
                                var v24 int = (fi10 * fi6)
                                v25 := GenS1{F0: 0, F1: v3}
                                _ = v24
                                _ = v25
                            }
                            var v26 uint32 = 2
                            _ = v26
                        }
                    }
                    _ = fi10
                }
                v2 -= 1
            }
        }
        _ = fi6
    }
    v3 = -267
    v3 -= (v3 & v3)
    for fi27 := 0; fi27 < 14; fi27++ {
        if true {
            for fi28 := 0; fi28 < 5; fi28++ {
                v29 := GenS2{F0: -1, F1: 1}
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 int8 = 42
                var v32 uint64 = 4294967295
                var v33 float32 = float32(float64(3.14159))
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        _ = fi27
    }
    var v34 int8 = 42
    switch ("hello" + (("foo bar" + "") + ("" + "hello"))) {
    case "":
        for fi35 := 0; fi35 < 9; fi35++ {
            v36 := GenS2{F0: 471, F1: -963}
            _ = fi35
            _ = v36
        }
    default:
        var v37 uint16 = (^(v2))
        _ = v37
    }
    fmt.Println(v2, v3)
    v38 := GenS1{F0: 0, F1: v3}
    v3 = -2147483648
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 int32 = v3
        v40 += v40
        _ = fi39
        _ = v40
    }
    v41 := GenS2{F0: v3, F1: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v41
}
