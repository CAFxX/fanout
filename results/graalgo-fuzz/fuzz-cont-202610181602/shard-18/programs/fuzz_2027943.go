package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
func genFn0(p0 bool, p1 int) int8 {
    v7 := GenS1{F0: 32767}
    if (p0 && p0) {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint16 = 0
            _ = fi8
            _ = v9
        }
    } else {
        p0 = false
        if p0 {
            switch 867 {
            case -1:
                for fi10 := 0; fi10 < 11; fi10++ {
                    if p0 {
                        var v11 uint32 = 1
                        var v12 string = "foo bar"
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
            case -836:
                p1 += 4294967295
            }
        } else {
            v13 := GenS1{F0: -795}
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            switch 42 {
            case -2147483648:
                var v16 int64 = 1
                _ = v16
            case -374:
                var v17 uint8 = 127
                _ = v17
            case 4294967295:
                if p0 {
                    if (((1.5 - 1.5) == -2.25) && true) {
                        if (true && false) {
                            v18 := []int{}
                            if p0 {
                                p1 -= -349
                                _ = len(v18)
                            }
                            _ = v18
                        } else {
                            if p0 {
                                var v19 int16 = 1
                                fmt.Println(fi15, v19, v19)
                                _ = v19
                            }
                            var v20 int64 = -1
                            p1 -= (p1 << ((4294967295) & 63))
                            _ = v20
                        }
                        p0 = p0
                        if p0 {
                            p1 -= (fi15 << ((353) & 63))
                            if (fi15 <= 648) {
                                fmt.Println(fi14, fi15)
                                p0 = (!(1 < -2147483648))
                                v21 := []int{p1}
                                _ = v21
                            }
                        } else {
                            var v22 bool = (!((p0 || (255 > 998)) && (-128 > 42)))
                            _ = v22
                        }
                    }
                    if (p0 || ((!(1 >= 220)) && p0)) {
                        v23 := []int{fi14, fi15, (p1 % (((fi14 / ((98) | 1))) | 1))}
                        fmt.Println(p0, p0)
                        _ = v23
                    }
                    var v24 uint64 = 18446744073709551615
                    _ = v24
                }
            }
            var v25 string = ("a" + "hello")
            _ = fi15
            _ = v25
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v7
    return -128
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{42, 2147483647, 794, 2}
    var v3 int8 = 1
    var v4 uint32 = 2
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    v4 = 492
    for fi26 := 0; fi26 < 11; fi26++ {
        v4 = v4
        v27 := GenS2{F0: 255, F1: 2, F2: 1}
        _ = fi26
        _ = v27
    }
    var v28 uint = 344
    var v29 int32 = 241
    if (1 <= 233) {
        fmt.Println(v29)
        v30 := GenS1{F0: -32768}
        switch 1 {
        case 128:
            _ = len(v6)
        case 32767:
            v4 = v4
        default:
            fmt.Println(v1, v3, v3)
        }
        _ = v30
    }
    var v31 uint = (^(v28))
    v29 = v29
    v3 -= (^(v3))
    for fi32 := 0; fi32 < 7; fi32++ {
        delete(v6, (fi32 / ((9223372036854775807) | 1)))
        for fi33 := 0; fi33 < 2; fi33++ {
            if (false || true) {
                v1 = ((("foo bar" + ((v1 + (v1 + ("foo bar" + v1))) + ((v1 + "foo bar") + "\n\t"))) + v1) + v1)
                fmt.Println(v31)
            } else {
                if (2 == 18446744073709551615) {
                    var v34 int32 = (v29 * v29)
                    v35 := GenS1{F0: 196}
                    fmt.Println(v1)
                    _ = v34
                    _ = v35
                } else {
                    v36 := GenS2{F0: 65535, F1: 2, F2: 361}
                    var v37 string = v1
                    delete(v5, v4)
                    _ = v36
                    _ = v37
                }
            }
            v28 = (v31 >> ((v28) & 63))
            _ = fi33
        }
        _ = fi32
    }
    fmt.Println(v4)
    v29 = 988
    v31 = 46
    var v38 float64 = (3.14159 + (1e10 / 1.0))
    var v39 int64 = -452
    for fi40 := 0; fi40 < 8; fi40++ {
        v29 = (v29 * (v29 / ((v29) | 1)))
        _ = fi40
    }
    var v41 float64 = (v38 / (((v38 / 0.1) * (v38 * float64(1e10))) - (v38 / v38)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
