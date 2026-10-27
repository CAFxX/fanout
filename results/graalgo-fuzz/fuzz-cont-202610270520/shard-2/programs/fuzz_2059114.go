package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint16 {
    v5 := []int{-2147483648, 4294967295, 0, 42}
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 180, F1: 1, F2: 529}
        var v8 int16 = 346
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = p0
    if false {
        switch ("\n\t" + "hello") {
        case "foo bar":
            var v9 string = "a"
            _ = v9
        default:
            v5 = append(v5, -9223372036854775808)
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            if false {
                var v11 int32 = 483
                v12 := GenS1{F0: 87, F1: 255, F2: (-(v11))}
                v13 := GenS1{F0: 93, F1: 255, F2: (v11 / (((^(v11))) | 1))}
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    } else {
        p0 = ("" <= ("\n\t" + "a"))
        fmt.Println(p0)
        if (p0 || ((p0 || true) || (!p0))) {
            p0 = (!p0)
            for fi14 := 0; fi14 < 18; fi14++ {
                for fi15 := 0; fi15 < 11; fi15++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v16 uint64 = 42
                    _ = fi15
                    _ = v16
                }
                var v17 float64 = 1e10
                _ = fi14
                _ = v17
            }
            switch "hello" {
            case "a":
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
            default:
                if (1 > -1) {
                    if len(v5) > 0 {
                        v5[0] = 32768
                    }
                    if p0 {
                        v5 = append(v5, 4294967295)
                    }
                    p0 = p0
                } else {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v20 := GenS1{F0: 2, F1: 127, F2: 42}
                            var v21 int8 = -128
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        if (!(!(-128 == 1))) {
                            p0 = true
                            var v22 float32 = (((((-2.25 - 1e10) - 1.5) - -2.25) + 0.0) + 0.1)
                            _ = v22
                        }
                        _ = fi18
                    }
                    _ = len(v5)
                }
            }
        }
    }
    if (p0 || p0) {
        var v23 int16 = 32767
        fmt.Println(p0)
        _ = v23
    } else {
        v24 := GenS1{F0: 1, F1: 42, F2: -2147483648}
        for fi25 := 0; fi25 < 1; fi25++ {
            p0 = (1 >= -119)
            p0 = (p0 && p0)
            var v26 float32 = 1.5
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    _ = p0
    _ = v5
    return 255
}
func genFn1() uint {
    var v27 uint64 = 18446744073709551615
    if (((2 > 42) || (!false)) && false) {
        var v28 uint8 = 42
        _ = v28
    } else {
        v29 := []int{-82, 32768, -2147483648, -1}
        v27 = 18446744073709551615
        _ = v29
    }
    v30 := []int{32768}
    _ = v27
    _ = v30
    return 2
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    v4 := []string{(("\n\t" + "hello") + ("a" + "\n\t")), "foo bar", "\n\t"}
    var v31 string = "a"
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 uint64 = 0
        v34 := GenS1{F0: 255, F1: 1, F2: 130}
        if true {
            fmt.Println(fi32, v3)
            v35 := GenS1{F0: 127, F1: 1, F2: 162}
            _ = v35
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    if false {
        v36 := GenS1{F0: 0, F1: genFn0((-810 == 0)), F2: 42}
        _ = v36
    } else {
        v37 := GenS2{F0: 65535}
        var v38 int16 = 32767
        _ = v37
        _ = v38
    }
    _ = len(v4)
    var v39 uint64 = 4294967295
    var v40 uint8 = 42
    v39 = v39
    var v41 int16 = -1
    fmt.Println(v31, v31, v3)
    for fi42 := 0; fi42 < 10; fi42++ {
        v40 = 255
        var v43 uint64 = 18446744073709551615
        _ = len(v4)
        _ = fi42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
