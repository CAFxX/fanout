package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 int32, p1 int, p2 string) int8 {
    v6 := GenS1{F0: 790}
    v7 := GenS1{F0: 255}
    v8 := []int{p1, p1, p1}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 int16 = 0
    v2 := []int{730, -9223372036854775808}
    var v3 uint16 = 42
    var v4 uint8 = 0
    v5 := make(map[int]uint32)
    var v9 uint8 = v4
    v5[9223372036854775807] = 633
    switch int(v1) {
    case 32768:
        switch -807 {
        case 2147483647:
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 bool = ((v9 >= v9) || false)
                v1 = v1
                v12 := GenS1{F0: v3}
                _ = fi10
                _ = v11
                _ = v12
            }
        case -82:
            if true {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v2 = append(v2, fi13)
                    if false {
                        var v14 bool = true
                        _ = v14
                    }
                    _ = fi13
                }
                fmt.Println(v3, v1)
            } else {
                switch ("\n\t" + ("foo bar" + "hello")) {
                case "\n\t":
                    v4 += 42
                case "a":
                    var v15 uint8 = v4
                    _ = v15
                }
                for fi16 := 0; fi16 < 9; fi16++ {
                    fmt.Println(v9, v4, v1)
                    _ = fi16
                }
            }
        case 128:
            fmt.Println(v1)
        }
    default:
        if true {
            if len(v2) > 0 {
                v2[0] = 998
            }
            var v17 float32 = float32(1e10)
            if false {
                v18 := v5[128]
                delete(v5, 127)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    fmt.Println(v17, v9)
                    _ = fi19
                }
                if (((^(v1)) >= 32767) || true) {
                    fmt.Println(v9)
                    delete(v5, 618)
                } else {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 int32 = 429
                        for fi22 := 0; fi22 < 15; fi22++ {
                            if true {
                                v1 -= -268
                            } else {
                                _ = len(v5)
                                switch (fi20 & fi22) {
                                case -415:
                                    v4 = (v4 + (v4 & 144))
                                case -1:
                                    v3 = 638
                                    v21 -= v21
                                    var v23 bool = false
                                    _ = v23
                                default:
                                    fmt.Println(v4, v4, fi20)
                                }
                            }
                            if ((^(v3)) <= v3) {
                                v4 = v9
                                if (4294967295 == 109) {
                                    var v24 int16 = 1000
                                    fmt.Println(v21)
                                    _ = v24
                                }
                            }
                            _ = len(v2)
                            _ = fi22
                        }
                        v21 += v21
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v4, v17, v3)
                }
                var v25 bool = ((!false) && false)
                _ = v25
            }
            _ = v17
        }
    }
    v26 := GenS2{F0: v1, F1: v4, F2: -32}
    v1 = (v1 / ((0) | 1))
    _ = len(v2)
    var v27 uint32 = 2147483647
    if len(v2) > 0 {
        v2[0] = 32768
    }
    var v28 uint16 = 2
    var v29 uint32 = (v27 - uint32(v4))
    if (1.5 == -2.25) {
        if true {
            var v30 uint32 = v27
            switch "foo bar" {
            case "":
                v31 := GenS3{F0: 32767, F1: 42}
                _ = v31
            case "a":
                var v32 float32 = (0.1 * 3.14159)
                _ = v32
            case "hello":
                var v33 float64 = float64(0.0)
                _ = v33
            }
            v28 -= (^(v28))
            _ = v30
        }
        var v34 float32 = 1.5
        var v35 uint32 = v27
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
