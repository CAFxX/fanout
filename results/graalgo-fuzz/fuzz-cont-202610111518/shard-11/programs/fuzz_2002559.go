package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0() int {
    if (32767 == 128) {
        for fi5 := 0; fi5 < 4; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := []int{fi5, fi6, fi6, (^(fi6))}
                v8 := GenS1{F0: 1, F1: false}
                var v9 int32 = -627
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            _ = fi5
        }
        v10 := []int{4294967295, 32768, 42, -355}
        v11 := GenS1{F0: 1, F1: true}
        _ = v10
        _ = v11
    } else {
        switch -605 {
        case 910:
            var v12 int32 = 1
            _ = v12
        case -1:
            v13 := GenS1{F0: 1000, F1: (true || true)}
            _ = v13
        case -868:
            if (0.0 > ((0.1 / 1e10) - float32((-2.25 - (float64(float32((0.1 / 1e10))) + (1.5 + -2.25)))))) {
                var v14 uint16 = 255
                _ = v14
            }
        }
    }
    var v15 bool = false
    for fi16 := 0; fi16 < 11; fi16++ {
        v15 = v15
        v17 := GenS1{F0: 1000, F1: false}
        v18 := GenS1{F0: -1, F1: false}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float32 = 1e10
    _ = v15
    _ = v19
    return 32767
}
func genFn1(p0 uint, p1 uint32) int8 {
    switch "" {
    case "a":
        var v20 bool = false
        _ = v20
    default:
        p0 += 255
    }
    switch 1 {
    case 128:
        if (-128 == 1) {
            if ((1 <= 1) || false) {
                p0 = (^(p0))
                v21 := []int{32768, 32768, -130}
                var v22 bool = true
                _ = v21
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 float64 = -2.25
                _ = fi23
                _ = v24
            }
            if true {
                var v25 uint = (^(p0))
                _ = v25
            } else {
                v26 := []int{genFn0()}
                _ = v26
            }
            fmt.Println(p1, p0)
        }
    case 0:
        p1 += 599
    default:
        v27 := GenS1{F0: 0, F1: false}
        _ = v27
    }
    p1 = (^(p1))
    var v28 string = ("" + "foo bar")
    _ = p0
    _ = p1
    _ = v28
    return 100
}

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]string)
    var v4 uint64 = (v2 << ((18446744073709551615) & 63))
    v29 := GenS1{F0: -32768, F1: false}
    for fi30 := 0; fi30 < 6; fi30++ {
        fmt.Println(v1)
        _ = fi30
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            v4 -= 292
            if (1e10 >= 3.14159) {
                switch "foo bar" {
                case "a":
                    var v33 uint64 = 18446744073709551615
                    _ = v33
                case "foo bar":
                    var v34 uint8 = 2
                    _ = v34
                case "":
                    var v35 int16 = 42
                    _ = v35
                }
            } else {
                v4 -= v4
            }
            var v36 float32 = (0.1 * float32(0.1))
            _ = fi32
            _ = v36
        }
        _ = fi31
    }
    v37 := GenS1{F0: 134, F1: true}
    var v38 uint64 = v4
    var v39 uint32 = 2147483648
    if true {
        var v40 uint32 = 588
        if false {
            for fi41 := 0; fi41 < 6; fi41++ {
                if ("foo bar" <= "\n\t") {
                    v42 := GenS1{F0: 25, F1: (1 < 4294967295)}
                    v43 := []int{fi41}
                    _ = v42
                    _ = v43
                } else {
                    if true {
                        if false {
                            var v44 int8 = genFn1(42, (^(v39)))
                            v38 += 5
                            _ = v44
                        } else {
                            switch (("hello" + ("foo bar" + ("" + "\n\t"))) + "\n\t") {
                            case "foo bar":
                                v45 := GenS1{F0: -32768, F1: ((128 == 127) || true)}
                                _ = v45
                            }
                        }
                    }
                }
                v38 = (v38 | 18446744073709551615)
                switch -1 {
                case 32767:
                    v46 := GenS1{F0: -32768, F1: (("foo bar" + "hello") != ("" + ("foo bar" + "a")))}
                    _ = v46
                case 0:
                    var v47 int32 = -1
                    _ = v47
                case 32768:
                    v48 := GenS1{F0: 228, F1: false}
                    _ = v48
                default:
                    v40 += 0
                }
                _ = fi41
            }
        } else {
            v38 = v2
        }
        var v49 string = "foo bar"
        _ = v40
        _ = v49
    }
    var v50 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v50
}
