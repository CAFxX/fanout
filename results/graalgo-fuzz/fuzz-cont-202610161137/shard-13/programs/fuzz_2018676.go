package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint32 }
func genFn0(p0 string) string {
    if false {
        v7 := GenS1{F0: 42}
        var v8 int16 = 42
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -550
        v10 := []int{-9223372036854775808}
        _ = v10
        _ = v9
    }
    var v11 float64 = 0.1
    _ = p0
    _ = v11
    return "foo bar"
}

func main() {
    var v1 int32 = 0
    var v2 string = ""
    v3 := []string{v2, ((v2 + "a") + v2), v2, (v2 + v2)}
    var v4 uint = 577
    var v5 float32 = ((1e10 + ((0.0 * 0.0) + (-2.25 - float32(-2.25)))) - (1e10 - (3.14159 / 1.5)))
    v6 := []int{0, 127}
    var v12 uint = (v4 ^ v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v13 := GenS1{F0: 128}
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 9; fi15++ {
        if (1.5 != 3.14159) {
            v16 := GenS1{F0: 1}
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v17 int8 = 127
                _ = v17
            }
            _ = v16
        } else {
            v18 := GenS1{F0: 97}
            if (!true) {
                fmt.Println(v4, v2)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = v18
        }
        _ = fi15
    }
    v19 := GenS2{F0: 128, F1: (v1 >> ((0) & 63))}
    v2 = ("\n\t" + (v2 + (v2 + "\n\t")))
    if false {
        fmt.Println(v1)
        v20 := GenS3{F0: uint8(v12)}
        _ = len(v6)
        _ = v20
    } else {
        v1 += v1
    }
    v1 += 0
    if true {
        fmt.Println(v5)
    } else {
        if (false || true) {
            var v21 string = (v2 + v2)
            v14 = v14
            for fi22 := 0; fi22 < 17; fi22++ {
                if len(v6) > 0 {
                    v6[0] = (fi22 & 42)
                }
                _ = fi22
            }
            _ = v21
        } else {
            if false {
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 int16 = 494
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 string = v2
                        if true {
                            v3 = append(v3, v2)
                        } else {
                            if len(v3) > 0 {
                                v3[0] = v26
                            }
                            v27 := GenS4{F0: 0, F1: 770, F2: 2147483647}
                            _ = v27
                        }
                        if true {
                            fmt.Println(v14, v24)
                            if true {
                                _ = len(v6)
                                if (true || (250 < 0)) {
                                    var v28 uint16 = 664
                                    v29 := GenS3{F0: 127}
                                    _ = v28
                                    _ = v29
                                }
                                var v30 bool = true
                                v1 -= -1
                                _ = v30
                            }
                        }
                        if false {
                            if len(v3) > 0 {
                                v3[0] = ((v26 + v26) + "\n\t")
                            }
                            v3 = append(v3, v2)
                            var v31 int = (fi23 / (((fi25 ^ (fi25 / ((215) | 1)))) | 1))
                            _ = v31
                        } else {
                            fmt.Println(v24)
                            var v32 uint = 255
                            _ = v32
                        }
                        _ = fi25
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                }
            }
        }
        v2 = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
