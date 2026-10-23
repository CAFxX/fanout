package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0() int8 {
    var v7 float64 = -2.25
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int = (fi8 - fi8)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v7 += (1e10 - v7)
        if false {
            v11 := []int{(fi10 - (^(fi10)))}
            var v12 string = ("\n\t" + "hello")
            _ = v11
            _ = v12
        } else {
            v7 = v7
        }
        _ = fi10
    }
    var v13 uint16 = 366
    _ = v13
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{420, 1, 128, 1}
    var v3 uint16 = 986
    var v4 int = 0
    var v5 int8 = -21
    var v6 int = (v4 << ((634) & 63))
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 bool = false
        v6 -= fi14
        _ = fi14
        _ = v15
    }
    v6 += v6
    v5 += (v5 + v5)
    v5 -= 127
    v5 -= v5
    var v16 uint8 = 51
    var v17 uint64 = 2
    switch (v6 - (v4 | v6)) {
    case -565:
        var v18 int32 = -2147483648
        _ = v18
    case 0:
        switch "foo bar" {
        case "":
            v19 := GenS1{F0: -1, F1: (v5 / (((v5 << ((124) & 63))) | 1)), F2: 600}
            _ = v19
        default:
            var v20 string = ""
            _ = v20
        }
    default:
        if ((0 > 705) || true) {
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 int16 = 32767
                v16 = 127
                _ = fi21
                _ = v22
            }
            v4 = v6
            if (!true) {
                for fi23 := 0; fi23 < 18; fi23++ {
                    fmt.Println(v5, v3, v17)
                    _ = fi23
                }
                if (true || (2147483647 == 1)) {
                    var v24 int32 = -724
                    _ = v24
                } else {
                    var v25 float64 = float64((1.5 - 1e10))
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 string = "hello"
                        for fi28 := 0; fi28 < 2; fi28++ {
                            if true {
                                if false {
                                    v27 = (v27 + v27)
                                    fmt.Println(v6, v5, v27)
                                } else {
                                    var v29 bool = false
                                    if (!v29) {
                                        v30 := v1[(fi26 - fi28)]
                                        var v31 int32 = 1
                                        _ = v30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                            } else {
                                v25 += 1e10
                                v32 := GenS2{F0: 2, F1: (v3 / ((v3) | 1)), F2: 2}
                                v2 = append(v2, int(v16))
                                _ = v32
                            }
                            _ = fi28
                        }
                        fmt.Println(v27, v4, v5)
                        _ = fi26
                        _ = v27
                    }
                    v25 += 0.0
                    _ = v25
                }
            } else {
                var v33 uint8 = (v16 >> ((255) & 63))
                var v34 float64 = (float64(0.1) * 1e10)
                _ = v33
                _ = v34
            }
        } else {
            if false {
                delete(v1, v6)
                v2 = append(v2, 42)
            } else {
                var v35 int64 = -9223372036854775808
                _ = v35
            }
            var v36 float64 = (0.0 * 0.0)
            v6 -= v6
            _ = v36
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v37 float32 = 0.0
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
