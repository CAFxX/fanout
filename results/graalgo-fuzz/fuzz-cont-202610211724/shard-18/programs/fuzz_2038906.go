package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint) int32 {
    p0 += p0
    var v6 int8 = 0
    v7 := GenS1{F0: 196, F1: 116}
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 int16) int8 {
    v8 := GenS1{F0: 0, F1: 2}
    var v9 bool = ("foo bar" > "\n\t")
    v10 := GenS1{F0: 255, F1: 42}
    v9 = ((-2.25 >= 1e10) && (0 <= 492))
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int16 = 42
    v2 := []int{}
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 uint64 = 2
    v5 -= (v5 + v5)
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            delete(v3, "\n\t")
            var v12 int = fi11
            fmt.Println(fi11, fi11)
            _ = v12
        } else {
            for fi13 := 0; fi13 < 3; fi13++ {
                v2 = append(v2, fi13)
                v1 = 0
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi13
            }
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 uint8 = 127
        fmt.Println(v15)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 32767}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (uint8(v5) > 0) {
        v1 -= v1
        var v17 float32 = 1e10
        _ = len(v2)
        _ = v17
    }
    var v18 int64 = -9223372036854775808
    if (float64((-2.25 / (((0.1 * (0.0 + 0.1)) + 1.5) + (0.0 + -2.25)))) == (0.1 - 3.14159)) {
        switch -2147483648 {
        case 2147483647:
            v19 := GenS1{F0: 0, F1: 32767}
            _ = v19
        }
        v18 += 1
    } else {
        var v20 uint = 255
        v2 = append(v2, -1)
        _ = v20
    }
    v5 += v5
    v4[32767] = -277
    var v21 int8 = 100
    if (2147483648 < 65535) {
        if (-640 <= -9223372036854775808) {
            v22 := GenS1{F0: 255, F1: 65535}
            _ = v22
        } else {
            v1 -= 0
        }
    }
    if ((!false) || true) {
        v21 += genFn1(int16(v21))
        if false {
            v23 := GenS1{F0: 128, F1: 605}
            fmt.Println(v1, v18)
            fmt.Println(v21, v18, v21)
            _ = v23
        }
        if true {
            for fi24 := 0; fi24 < 8; fi24++ {
                for fi25 := 0; fi25 < 7; fi25++ {
                    switch (("\n\t" + "\n\t") + "a") {
                    case "\n\t":
                        if (!(!false)) {
                            v26 := GenS1{F0: 87, F1: 0}
                            if true {
                                v5 = v5
                                v5 -= (v5 | v5)
                                v21 = v21
                                var v27 uint64 = (v5 & (v5 >> ((42) & 63)))
                                _ = v27
                            }
                            _ = v26
                        } else {
                            if false {
                                var v28 uint8 = 2
                                v29 := GenS1{F0: (v28 ^ v28), F1: 42}
                                var v30 bool = ((v28 + v28) >= 255)
                                v18 += v18
                                var v31 uint8 = v28
                                v30 = v30
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v18, v21)
                            }
                            if false {
                                if (!true) {
                                    v5 = v5
                                    if false {
                                        fmt.Println(fi24)
                                        v1 += (v1 + v1)
                                    }
                                    v4[128] = fi24
                                } else {
                                    v1 += v1
                                    v21 -= 2
                                    v32 := GenS1{F0: 35, F1: 0}
                                    v18 -= (v18 / ((v18) | 1))
                                    _ = v32
                                }
                            }
                            var v33 int64 = (v18 ^ 1)
                            _ = v33
                        }
                        switch (fi25 + (fi25 % ((1) | 1))) {
                        case 2147483647:
                            var v34 float32 = 0.1
                            _ = v34
                        }
                        v21 += (v21 | v21)
                        var v35 int32 = -2147483648
                        _ = v35
                    }
                    v2 = append(v2, (^(fi24)))
                    _ = fi25
                }
                var v36 bool = false
                fmt.Println(v5, v21)
                _ = fi24
                _ = v36
            }
        }
    }
    v37 := GenS1{F0: 128, F1: 32767}
    v38 := GenS1{F0: 177, F1: 334}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
