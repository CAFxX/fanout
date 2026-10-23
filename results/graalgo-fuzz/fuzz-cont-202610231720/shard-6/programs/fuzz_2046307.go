package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{4294967295, 588, 590, 2147483648}
    v3 := []int{9223372036854775807, 1, -9223372036854775808}
    var v4 int = 127
    var v5 int = 380
    var v6 float64 = 0.1
    delete(v1, (v4 | (v5 >> ((847) & 63))))
    v7 := GenS1{F0: (-1 < 42), F1: (v5 < (v5 << ((364) & 63))), F2: 18446744073709551615}
    v6 -= v6
    var v8 int64 = 1
    if (float32((1e10 / v6)) <= 1.5) {
        _ = len(v3)
        if (v8 >= (v8 & v8)) {
            v9 := GenS1{F0: true, F1: ("a" != ""), F2: 42}
            if false {
                switch (^(v5)) {
                case 2147483647:
                    for fi10 := 0; fi10 < 7; fi10++ {
                        for fi11 := 0; fi11 < 15; fi11++ {
                            v8 = -9223372036854775808
                            v12 := GenS2{F0: 42}
                            v4 += -2147483648
                            _ = fi11
                            _ = v12
                        }
                        _ = fi10
                    }
                default:
                    for fi13 := 0; fi13 < 12; fi13++ {
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v8 -= (v8 | -865)
                            _ = fi14
                        }
                        for fi15 := 0; fi15 < 20; fi15++ {
                            if true {
                                delete(v1, fi15)
                            }
                            _ = fi15
                        }
                        _ = fi13
                    }
                }
                fmt.Println(v8, v8)
            } else {
                switch "a" {
                case "hello":
                    v6 = v6
                default:
                    if (1 <= 42) {
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v17 := v1[838]
                            _ = fi16
                            _ = v17
                        }
                        if (true || false) {
                            v6 = ((v6 / v6) * 1e10)
                        }
                    } else {
                        v8 = -9223372036854775808
                        if false {
                            v6 -= ((v6 - v6) / v6)
                            var v18 uint8 = 1
                            v18 = (^(v18))
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 1; fi19++ {
                            fmt.Println(v8, v5)
                            v4 += (v5 | (^(v5)))
                            _ = fi19
                        }
                    }
                }
                var v20 int = 127
                if true {
                    delete(v1, (v4 - (v4 - -843)))
                    v5 += v4
                }
                _ = v20
            }
            _ = v9
        } else {
            if (!(!true)) {
                v21 := GenS1{F0: true, F1: true, F2: 42}
                for fi22 := 0; fi22 < 11; fi22++ {
                    v23 := GenS2{F0: 1}
                    _ = fi22
                    _ = v23
                }
                if (255 < 225) {
                    if (1 >= -42) {
                        var v24 float64 = float64(1e10)
                        for fi25 := 0; fi25 < 15; fi25++ {
                            var v26 uint64 = 938
                            v6 -= 0.0
                            v27 := GenS1{F0: false, F1: true, F2: 0}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                    }
                }
                _ = v21
            }
        }
    }
    fmt.Println(v5)
    var v28 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
