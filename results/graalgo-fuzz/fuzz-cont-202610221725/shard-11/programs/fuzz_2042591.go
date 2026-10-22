package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 uint64 }
func genFn0(p0 int) int16 {
    v5 := []int{(p0 & (p0 & p0))}
    p0 -= p0
    fmt.Println(p0)
    v6 := GenS1{F0: 51, F1: "hello"}
    var v7 int8 = 127
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 int32, p1 bool) uint64 {
    p1 = (80 >= 42)
    if p1 {
        var v8 int = 4294967295
        _ = v8
    } else {
        p1 = (!p1)
        p1 = (931 == 42)
        switch 516 {
        case 894:
            if (genFn0(-26) != genFn0(42)) {
                p0 = p0
                v9 := []int{298, 32767}
                for fi10 := 0; fi10 < 19; fi10++ {
                    p1 = false
                    _ = len(v9)
                    for fi11 := 0; fi11 < 7; fi11++ {
                        fmt.Println(p1, p1, p1)
                        switch ("" + (("\n\t" + "hello") + ("\n\t" + ("foo bar" + ("\n\t" + ("\n\t" + "")))))) {
                        case "\n\t":
                            v9 = append(v9, fi11)
                        default:
                            var v12 int16 = 32767
                            _ = v12
                        }
                        _ = fi11
                    }
                    _ = fi10
                }
                _ = v9
            }
        case 32768:
            if true {
                for fi13 := 0; fi13 < 10; fi13++ {
                    fmt.Println(p0, fi13)
                    _ = fi13
                }
                p1 = (!(!p1))
            } else {
                v14 := []int{42}
                p1 = (255 == 1)
                if len(v14) > 0 {
                    v14[0] = -275
                }
                _ = v14
            }
        default:
            if p1 {
                for fi15 := 0; fi15 < 16; fi15++ {
                    p1 = (p1 || p1)
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v17 := []int{fi15, -9223372036854775808, (fi15 ^ 0), -561}
                        fmt.Println(fi16, fi15)
                        fmt.Println(p1)
                        _ = fi16
                        _ = v17
                    }
                    if (!(!(-2.25 == (-2.25 - -2.25)))) {
                        var v18 int8 = -1
                        var v19 float64 = (0.0 + (((-2.25 * (0.0 / 1.5)) + 1e10) - (((0.0 - 0.1) / 0.1) - (0.0 * 0.0))))
                        _ = v18
                        _ = v19
                    } else {
                        p1 = p1
                        if p1 {
                            for fi20 := 0; fi20 < 20; fi20++ {
                                if p1 {
                                    if (42 < 2) {
                                        switch (^(fi15)) {
                                        case 128:
                                            p0 = (p0 * p0)
                                        case 741:
                                            if p1 {
                                                v21 := []int{63, (fi15 ^ fi15), fi15}
                                                v22 := GenS2{F0: -9223372036854775808, F1: 656}
                                                _ = v21
                                                _ = v22
                                            }
                                        case 42:
                                            var v23 bool = (792 == -906)
                                            _ = v23
                                        default:
                                            fmt.Println(fi15)
                                        }
                                    }
                                }
                                var v24 uint = 0
                                _ = fi20
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 17; fi25++ {
                                if p1 {
                                    if ((p1 && false) && (-2147483648 < -782)) {
                                        v26 := GenS1{F0: 0, F1: "hello"}
                                        _ = v26
                                    }
                                    var v27 bool = ((-2.25 - 1.5) == (((float32(-2.25) / (((1.5 - -2.25) + float32((-2.25 - 0.0))) * float32(0.1))) * -2.25) / (0.0 + -2.25)))
                                    if p1 {
                                        v28 := GenS2{F0: 0, F1: 689}
                                        _ = v28
                                    }
                                    v27 = true
                                    _ = v27
                                }
                                v29 := GenS1{F0: 42, F1: "\n\t"}
                                v30 := []int{(fi15 | fi25), (fi25 & (^(fi25)))}
                                _ = fi25
                                _ = v29
                                _ = v30
                            }
                        }
                    }
                    _ = fi15
                }
                fmt.Println(p1)
            } else {
                if true {
                    v31 := []int{}
                    _ = v31
                }
                fmt.Println(p1)
            }
        }
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    var v3 uint32 = 772
    var v4 int8 = 127
    v1 = 0
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 int = (fi32 + fi32)
        var v34 string = ("\n\t" + (("hello" + "") + (("" + "") + "foo bar")))
        fmt.Println(v1, v1, v3)
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 uint64 = 625
    _ = len(v2)
    v1 += v1
    fmt.Println(v1, v3, v4)
    var v36 bool = (-2.25 != -2.25)
    v3 = v3
    for fi37 := 0; fi37 < 18; fi37++ {
        v35 = (v35 | v35)
        fmt.Println(v35, v4, fi37)
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
