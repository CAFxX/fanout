package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 bool, p1 int8, p2 string) uint {
    v7 := []int{int(p1)}
    if p0 {
        v8 := GenS1{F0: 2, F1: 42}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v8
    } else {
        var v9 int64 = 9223372036854775807
        fmt.Println(p0, p2, p1)
        var v10 int32 = int32(v9)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(p2, p2, p2)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 918
}

func main() {
    var v1 int16 = 1000
    v2 := []int{1, 32767}
    var v3 int16 = (v1 % ((v1) | 1))
    v4 := make(map[string]string)
    var v5 uint16 = 42
    v6 := []string{"\n\t", ("" + ("hello" + "\n\t")), "foo bar", "a"}
    v5 -= v5
    if (!false) {
        var v12 int = 586
        for fi13 := 0; fi13 < 19; fi13++ {
            v5 += v5
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int8 = -1
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        v3 += v3
        _ = v12
    }
    var v16 bool = false
    v17 := v4["hello"]
    var v18 int8 = int8(v1)
    switch (((v17 + v17) + v17) + v17) {
    case "":
        if (0 > -270) {
            for fi19 := 0; fi19 < 8; fi19++ {
                if v16 {
                    switch "a" {
                    case "":
                        var v20 uint16 = 65535
                        _ = v20
                    case "a":
                        for fi21 := 0; fi21 < 16; fi21++ {
                            v22 := GenS2{F0: true, F1: v17, F2: 713}
                            _ = fi21
                            _ = v22
                        }
                    default:
                        var v23 int64 = 42
                        _ = v23
                    }
                    if ((!v16) || (!(!v16))) {
                        v17 = v17
                        var v24 uint = 42
                        v18 += v18
                        _ = v24
                    } else {
                        fmt.Println(v17, v16, v1)
                        if (v16 && (!v16)) {
                            _ = len(v4)
                            var v25 float64 = float64(1.5)
                            v16 = v16
                            _ = v25
                        } else {
                            var v26 float32 = 1e10
                            if ((v1 << ((255) & 63)) <= v3) {
                                v26 -= (v26 * ((((v26 * 3.14159) * (3.14159 / (v26 / (v26 * v26)))) * ((v26 - (1.5 - 0.1)) / ((0.0 - float32((0.0 * -2.25))) - -2.25))) / 1.0))
                                var v27 int64 = -294
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    var v29 int16 = v1
                                    var v30 int8 = -1
                                    v31 := GenS2{F0: false, F1: v17, F2: 744}
                                    var v32 float64 = 1.5
                                    v6 = append(v6, v17)
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                _ = v27
                            } else {
                                fmt.Println(v3)
                            }
                            _ = v26
                        }
                        v4[v17] = "\n\t"
                    }
                } else {
                    v33 := GenS3{F0: int(v5), F1: -1}
                    _ = v33
                }
                var v34 int64 = 9223372036854775807
                var v35 uint16 = 32767
                _ = fi19
                _ = v34
                _ = v35
            }
            if (4294967295 >= 9223372036854775807) {
                v1 = (v3 / (((-(v1))) | 1))
                var v36 uint = 1
                if false {
                    if false {
                        v6 = append(v6, v17)
                    } else {
                        v5 = v5
                        _ = len(v4)
                    }
                    fmt.Println(v18, v1)
                }
                _ = v36
            }
            if v16 {
                var v37 uint = 749
                var v38 uint64 = 4294967295
                _ = v37
                _ = v38
            }
        }
    default:
        var v39 uint32 = 2147483647
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        if ((v18 >> ((1) & 63)) <= -128) {
            v41 := GenS3{F0: (^(fi40)), F1: 1}
            v5 += (v5 & 1)
            _ = v41
        }
        delete(v4, v17)
        v42 := v4[v17]
        _ = fi40
        _ = v42
    }
    v3 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
