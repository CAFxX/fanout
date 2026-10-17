package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0() bool {
    v7 := GenS1{F0: 1, F1: 53}
    if (1000 > -32768) {
        var v8 int32 = 907
        v8 = (v8 % ((v8) | 1))
        _ = v8
    } else {
        v9 := GenS1{F0: 1, F1: 0}
        var v10 int32 = -158
        _ = v10
        _ = v9
    }
    _ = v7
    return false
}
func genFn1(p0 uint16, p1 int8) int {
    if false {
        switch "a" {
        case "":
            fmt.Println(p1, p1)
        case "a":
            var v11 float32 = 1e10
            _ = v11
        case "foo bar":
            p0 -= p0
        }
        p1 -= p1
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: 0, F1: -128}
            var v15 int8 = (p1 | (p1 + p1))
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(p1)
            var v17 uint32 = 1
            fmt.Println(fi12, v17)
            _ = fi16
            _ = v17
        }
        _ = fi12
    }
    if (100 < (p1 * (p1 - p1))) {
        v18 := []int{32767, -860, -481, 127}
        var v19 bool = true
        if true {
            var v20 uint32 = 2147483648
            v21 := GenS1{F0: 1, F1: p1}
            _ = v20
            _ = v21
        }
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    return 276
}

func main() {
    v1 := []uint32{1}
    v2 := []string{"hello"}
    v3 := []string{}
    var v4 int = 39
    v5 := make(map[int]uint32)
    var v6 bool = true
    var v22 float32 = 3.14159
    switch (v4 | -109) {
    case 42:
        var v23 int64 = 0
        _ = v23
    case -2147483648:
        if (!true) {
            var v24 float32 = ((-2.25 * v22) * (v22 - ((v22 + v22) + -2.25)))
            _ = v24
        }
    }
    if (42 < 12) {
        v25 := v5[v4]
        _ = v25
    }
    var v26 uint32 = 580
    var v27 int16 = 67
    v4 = 127
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS1{F0: uint(v26), F1: 42}
        v6 = false
        _ = fi28
        _ = v29
    }
    if ("foo bar" <= "\n\t") {
        switch ((("" + "foo bar") + (("foo bar" + (("foo bar" + "\n\t") + "a")) + ("\n\t" + "foo bar"))) + ("hello" + "hello")) {
        case "":
            _ = len(v1)
        case "foo bar":
            for fi30 := 0; fi30 < 13; fi30++ {
                switch fi30 {
                case -112:
                    if (v6 || ("hello" > "")) {
                        v22 += float32(0.1)
                    } else {
                        var v31 int8 = int8(v26)
                        _ = v31
                    }
                case 299:
                    v22 -= v22
                case 32768:
                    v27 = 0
                default:
                    for fi32 := 0; fi32 < 2; fi32++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if v6 {
                            v26 = v26
                            v6 = ("hello" != "foo bar")
                            if (v6 && v6) {
                                var v33 uint8 = 127
                                fmt.Println(v27)
                                var v34 uint16 = 229
                                _ = v33
                                _ = v34
                            }
                        } else {
                            v35 := v5[v4]
                            var v36 string = ("foo bar" + "hello")
                            v6 = (1 > 128)
                            _ = v35
                            _ = v36
                        }
                        if v6 {
                            fmt.Println(fi32)
                        } else {
                            v37 := GenS1{F0: 2, F1: -60}
                            if (-92 == -1) {
                                var v38 float32 = (v22 - (v22 - float32(-2.25)))
                                _ = v38
                            } else {
                                _ = len(v1)
                            }
                            _ = v37
                        }
                        _ = fi32
                    }
                }
                var v39 int32 = -1
                v2 = append(v2, (("a" + "foo bar") + "a"))
                _ = fi30
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 18; fi40++ {
                for fi41 := 0; fi41 < 3; fi41++ {
                    v27 += (v27 ^ (v27 & v27))
                    var v42 int8 = int8(v26)
                    v26 -= 1
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        }
        for fi43 := 0; fi43 < 5; fi43++ {
            v44 := v5[fi43]
            var v45 int16 = v27
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        v27 = 0
        var v47 int64 = 42
        v47 = 9223372036854775807
        _ = fi46
        _ = v47
    }
    delete(v5, v4)
    v6 = (v26 == v26)
    v4 = (v4 & -876)
    if v6 {
        fmt.Println(v6, v26)
        delete(v5, (-(v4)))
    } else {
        v6 = v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
