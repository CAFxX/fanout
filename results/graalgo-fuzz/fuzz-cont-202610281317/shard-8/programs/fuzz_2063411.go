package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int8, p2 int64) int16 {
    switch p0 {
    case 32767:
        if (p1 < int8(p2)) {
            fmt.Println(p1, p2, p2)
            v6 := []int{(p0 / ((42) | 1)), (^(p0))}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 8; fi7++ {
                switch 206 {
                case -2147483648:
                    if (!false) {
                        p0 += 0
                        if (!(154 != -326)) {
                            var v8 float64 = 0.0
                            for fi9 := 0; fi9 < 11; fi9++ {
                                fmt.Println(fi9)
                                if (0 >= 2147483648) {
                                    var v10 int16 = 1000
                                    _ = v10
                                }
                                var v11 uint8 = 128
                                _ = fi9
                                _ = v11
                            }
                            v12 := GenS1{F0: p2}
                            _ = v12
                            _ = v8
                        } else {
                            p2 -= p2
                        }
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v14 := GenS1{F0: -1}
                            v15 := GenS2{F0: 2}
                            if false {
                                if ((1.5 * 0.1) < -2.25) {
                                    p0 = fi7
                                }
                            } else {
                                p2 += p2
                                p2 = p2
                            }
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    } else {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            var v17 int16 = 0
                            p0 -= 127
                            _ = fi16
                            _ = v17
                        }
                        p1 = (^(p1))
                    }
                case 42:
                    p2 += (p2 + (p2 / ((p2) | 1)))
                case 880:
                    v18 := GenS1{F0: (^(p2))}
                    _ = v18
                default:
                    var v19 float64 = (-2.25 - -2.25)
                    _ = v19
                }
                v20 := GenS1{F0: (p2 / ((p2) | 1))}
                for fi21 := 0; fi21 < 17; fi21++ {
                    v22 := []int{32768}
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    v23 := GenS2{F0: 128}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi7
                _ = v20
            }
        }
    }
    var v24 uint8 = 255
    var v25 uint16 = 255
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 92
}

func main() {
    var v1 string = "a"
    var v2 string = v1
    v3 := make(map[string]uint32)
    var v4 float64 = 1.5
    var v5 bool = true
    var v26 int32 = -576
    fmt.Println(v5, v2)
    if v5 {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int = fi27
            for fi29 := 0; fi29 < 13; fi29++ {
                v30 := GenS2{F0: 79}
                var v31 int64 = 507
                v2 = "foo bar"
                v5 = v5
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi27
            _ = v28
        }
    }
    if false {
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 13; fi33++ {
                v34 := []int{(fi33 - fi33), 843}
                var v35 uint = 1
                var v36 float64 = (v4 + (v4 * v4))
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi32
        }
        fmt.Println(v4)
        v4 += (v4 * float64(3.14159))
    }
    v5 = true
    fmt.Println(v4, v4, v26)
    _ = len(v3)
    var v37 int32 = (v26 % (((v26 / (((v26 >> ((42) & 63))) | 1))) | 1))
    v4 = 0.0
    var v38 uint16 = 0
    var v39 int32 = 544
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
