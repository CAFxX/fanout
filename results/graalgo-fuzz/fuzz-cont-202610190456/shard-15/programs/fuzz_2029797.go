package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
func genFn0(p0 int64) string {
    v7 := GenS1{F0: true, F1: 1000}
    if true {
        var v8 float64 = 1e10
        p0 = p0
        _ = v8
    } else {
        p0 -= p0
    }
    if true {
        var v9 bool = false
        _ = v9
    }
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int = -9223372036854775808
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v7
    return ((((("foo bar" + "\n\t") + "hello") + "\n\t") + (("\n\t" + ("a" + ("foo bar" + "hello"))) + ("hello" + ("" + ("hello" + "hello"))))) + ("a" + "a"))
}
func genFn1(p0 uint16, p1 uint8, p2 uint) uint32 {
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 uint64 = 2
                _ = fi14
                _ = v15
            }
            fmt.Println(p0, p2)
            for fi16 := 0; fi16 < 11; fi16++ {
                fmt.Println(p2, p0)
                _ = fi16
            }
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 int16 = 1
                if false {
                    if true {
                        v19 := GenS1{F0: false, F1: (v18 & v18)}
                        p2 -= p2
                        if (1 != 0) {
                            v20 := GenS1{F0: (!((float64(-2.25) != 1e10) || true)), F1: -1}
                            var v21 int64 = 0
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS1{F0: ((!false) || (!true)), F1: v18}
                        _ = v19
                        _ = v22
                    } else {
                        var v23 float64 = 0.1
                        var v24 int8 = 0
                        _ = v23
                        _ = v24
                    }
                    if true {
                        var v25 int8 = 127
                        var v26 uint8 = (p1 - 128)
                        var v27 string = (("hello" + "a") + "")
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 string = "foo bar"
                    _ = v28
                }
                _ = fi17
                _ = v18
            }
        }
        _ = fi13
    }
    p2 += 2
    var v29 uint16 = uint16(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return 764
}

func main() {
    var v1 float64 = 1.5
    var v2 string = "a"
    var v3 int64 = 968
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]string)
    v30 := GenS1{F0: false, F1: 1000}
    fmt.Println(v1)
    v3 -= v3
    var v31 uint16 = 805
    for fi32 := 0; fi32 < 17; fi32++ {
        _ = len(v5)
        var v33 float32 = float32(v1)
        if (42 != 1000) {
            v34 := []int{}
            v35 := GenS2{F0: 0, F1: 770, F2: -1}
            _ = v34
            _ = v35
        } else {
            v2 = "foo bar"
        }
        _ = fi32
        _ = v33
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 int8 = -64
        _ = fi36
        _ = v37
    }
    var v38 float64 = v1
    v3 = (v3 / ((v3) | 1))
    var v39 string = "foo bar"
    for fi40 := 0; fi40 < 16; fi40++ {
        v1 += ((((v1 - (float64(((0.1 / (0.1 * 3.14159)) + 3.14159)) + (3.14159 * v1))) / v38) / float64(-2.25)) + v38)
        _ = fi40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v3 = 42
        fmt.Println(fi41, v31)
        switch "a" {
        case "a":
            var v42 string = (v39 + v39)
            _ = v42
        case "":
            var v43 bool = ((v3 | 42) != v3)
            _ = v43
        default:
            var v44 uint8 = 0
            _ = v44
        }
        _ = fi41
    }
    delete(v6, ((((v2 + "a") + v2) + v39) + v39))
    var v45 int8 = 42
    for fi46 := 0; fi46 < 6; fi46++ {
        v45 = -1
        fmt.Println(v45)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
