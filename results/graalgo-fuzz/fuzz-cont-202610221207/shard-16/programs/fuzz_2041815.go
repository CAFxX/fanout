package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0(p0 int8, p1 int16) int {
    v5 := []int{32767, -2147483648}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(p1, p1)
        _ = fi6
    }
    p1 = p1
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    _ = p0
    _ = p1
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 int8 = -1
    var v4 uint = 2
    var v7 int32 = 1
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float64 = ((1.5 / -2.25) + 3.14159)
        var v11 float64 = ((v10 + v10) - v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = v8
    if (42 <= v7) {
        for fi13 := 0; fi13 < 8; fi13++ {
            switch "foo bar" {
            case "\n\t":
                if false {
                    var v14 uint32 = 2147483648
                    v2[("a" + ("hello" + (("hello" + "\n\t") + "foo bar")))] = genFn0(v3, -32768)
                    if false {
                        v15 := []int{127}
                        _ = v15
                    } else {
                        v4 -= (v4 >> (((v4 + 42)) & 63))
                        fmt.Println(v7, v12)
                        if false {
                            if false {
                                fmt.Println(v14, v8, v8)
                            } else {
                                v16 := GenS1{F0: true, F1: (1 > 4294967295)}
                                for fi17 := 0; fi17 < 20; fi17++ {
                                    var v18 float64 = 1e10
                                    _ = fi17
                                    _ = v18
                                }
                                var v19 uint8 = 0
                                _ = v16
                                _ = v19
                            }
                            v3 = v3
                        }
                    }
                    _ = v14
                }
            }
            var v20 int64 = 1
            if true {
                var v21 float32 = ((1.5 + 3.14159) * -2.25)
                for fi22 := 0; fi22 < 12; fi22++ {
                    if (((1.5 / 3.14159) - v21) < v21) {
                        v21 = (1.5 / (1.5 / -2.25))
                    } else {
                        _ = len(v2)
                    }
                    _ = fi22
                }
                _ = v21
            } else {
                if false {
                    fmt.Println(v7)
                    delete(v2, "a")
                }
                fmt.Println(v7, v20, v4)
                v12 -= v8
            }
            _ = fi13
            _ = v20
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint64 = 1
            v4 = 2
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            v8 = 32767
            v26 := GenS2{F0: 493}
            _ = fi25
            _ = v26
        }
    } else {
        var v27 string = "a"
        var v28 int64 = 0
        var v29 uint = v4
        _ = v27
        _ = v28
        _ = v29
    }
    if true {
        var v30 int64 = 252
        _ = v30
    } else {
        var v31 bool = (!false)
        v7 = (v7 / ((v7) | 1))
        var v32 uint = (^(v4))
        _ = v31
        _ = v32
    }
    var v33 uint16 = v8
    var v34 uint = (v4 / ((v4) | 1))
    var v35 int32 = v7
    var v36 uint = (^(v34))
    var v37 float32 = 0.1
    fmt.Println(v7, v35, v35)
    fmt.Println(v36)
    v38 := GenS3{F0: 0, F1: true, F2: 18446744073709551615}
    switch 4294967295 {
    case -925:
        var v39 float64 = (float64(-2.25) - -2.25)
        _ = v39
    case 32768:
        if ((v3 + int8(v4)) <= v3) {
            for fi40 := 0; fi40 < 9; fi40++ {
                if false {
                    var v41 int32 = (-(v7))
                    fmt.Println(v8, v3)
                    _ = v41
                }
                if (!(v3 <= v3)) {
                    if (false && false) {
                        v42 := []int{genFn0((v3 - v3), 0)}
                        _ = v42
                    } else {
                        var v43 int16 = 557
                        var v44 uint16 = (v12 << ((18446744073709551615) & 63))
                        _ = v43
                        _ = v44
                    }
                    fmt.Println(v3, v8, v4)
                }
                _ = fi40
            }
            v45 := []int{-128, 1, 127}
            _ = v45
        } else {
            if true {
                for fi46 := 0; fi46 < 18; fi46++ {
                    var v47 uint16 = v8
                    _ = fi46
                    _ = v47
                }
            }
            var v48 string = "\n\t"
            var v49 int64 = 0
            _ = v48
            _ = v49
        }
    case -948:
        v7 -= (v7 * (v35 & v7))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v7
    _ = v8
}
