package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64) int {
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int8 = 101
        switch (^(fi7)) {
        case 32768:
            if (127 < v6) {
                if true {
                    v8 += 100
                    fmt.Println(v8, fi7)
                    fmt.Println(p0, p0, fi7)
                }
            }
        default:
            v9 := GenS1{F0: 124}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 32768
}
func genFn1(p0 uint32, p1 int32) bool {
    var v10 bool = true
    p1 = (p1 | (p1 / (((p1 * (^(p1)))) | 1)))
    switch -1 {
    case 942:
        v10 = (9223372036854775807 >= 566)
    case 372:
        fmt.Println(v10)
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v10
    return (v10 || true)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 1000
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 string = (("\n\t" + "hello") + "foo bar")
    if (true && (!false)) {
        fmt.Println(v2, v2)
    } else {
        var v11 float64 = (((0.1 - (float64(1.5) / float64(float32(-2.25)))) + 1.5) / (1e10 * 3.14159))
        _ = v11
    }
    v12 := v4[393]
    v2 = (v2 & v2)
    if true {
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                v14 := GenS1{F0: 127}
                for fi15 := 0; fi15 < 5; fi15++ {
                    v16 := []int{(-(fi13)), genFn0(2), fi13}
                    var v17 int8 = 0
                    v5 = v5
                    if (191 >= 0) {
                        v4[860] = ("\n\t" + "a")
                        switch "foo bar" {
                        case "a":
                            var v18 uint8 = 2
                            _ = v18
                        case "":
                            v19 := GenS1{F0: v17}
                            _ = v19
                        case "\n\t":
                            switch fi13 {
                            case -274:
                                v5 = v12
                            case 32767:
                                var v20 int64 = -431
                                _ = v20
                            case 252:
                                v21 := v3[32767]
                                _ = v21
                            }
                        }
                        v22 := GenS2{F0: (true && ((1 < 707) || (!(true || ((-(fi13)) == fi15))))), F1: (v2 % ((32767) | 1))}
                        _ = v22
                    } else {
                        if false {
                            _ = len(v16)
                        }
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v23 := []int{-634, fi13, 142}
                _ = v14
                _ = v23
            }
            v24 := []int{}
            _ = fi13
            _ = v24
        }
    } else {
        if (2 >= 714) {
            if (287 == -2147483648) {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 bool = (!true)
                    v2 += (-(v2))
                    v12 = (("hello" + v12) + v5)
                    _ = fi25
                    _ = v26
                }
                v5 = v5
                v27 := GenS1{F0: -128}
                _ = v27
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                if false {
                    v1[fi28] = true
                }
                fmt.Println(v2, v12)
                _ = fi28
            }
        } else {
            for fi29 := 0; fi29 < 1; fi29++ {
                v30 := []int{(fi29 * (fi29 / ((467) | 1))), (fi29 % ((32767) | 1))}
                v2 -= v2
                _ = fi29
                _ = v30
            }
            v12 = v12
            if (!((3.14159 > float64(1.5)) || true)) {
                for fi31 := 0; fi31 < 17; fi31++ {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v2 += (v2 % (((v2 % ((-32768) | 1))) | 1))
                        v5 = v5
                        var v33 uint = 42
                        if (127 < 128) {
                            v33 -= 65535
                            var v34 string = v5
                            _ = v34
                        } else {
                            v1[fi32] = true
                            v12 = ""
                        }
                        var v35 int = genFn0(1)
                        var v36 uint64 = 1
                        _ = fi32
                        _ = v33
                        _ = v35
                        _ = v36
                    }
                    v2 = v2
                    _ = fi31
                }
            } else {
                if (398 != 42) {
                    if (((1e10 + ((3.14159 - 1.5) / 1.0)) <= 0.1) && (1 <= -1)) {
                        var v37 uint16 = 1
                        v38 := v3[genFn0(1)]
                        _ = v37
                        _ = v38
                    } else {
                        var v39 uint8 = 127
                        v40 := []int{9223372036854775807, 32767, -9223372036854775808}
                        _ = v39
                        _ = v40
                    }
                    if false {
                        v2 -= (-(v2))
                    } else {
                        v41 := []int{42}
                        _ = v41
                    }
                }
            }
        }
        v42 := []int{-2147483648, -149, 9223372036854775807}
        for fi43 := 0; fi43 < 11; fi43++ {
            fmt.Println(fi43, fi43)
            _ = fi43
        }
        _ = v42
    }
    fmt.Println(v5)
    v44 := GenS2{F0: false, F1: v2}
    fmt.Println(v2, v5)
    if true {
        fmt.Println(v2, v2)
        v2 = 1000
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
