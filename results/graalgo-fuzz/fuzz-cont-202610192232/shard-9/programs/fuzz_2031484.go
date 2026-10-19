package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 int16 }
func genFn0(p0 int16, p1 int32) uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{}
        v7 := GenS1{F0: -128}
        v8 := GenS1{F0: 42}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    p0 = 0
    var v9 int64 = -464
    _ = p0
    _ = p1
    _ = v9
    return 4294967295
}
func genFn1(p0 uint64, p1 int16, p2 int8) uint {
    p1 = 32767
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float32 = 3.14159
        var v12 int32 = 2147483647
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = 2
    v3 := make(map[string]string)
    var v4 int64 = 1
    v2 = v1
    if false {
        if true {
            fmt.Println(v2, v2)
            if false {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := GenS1{F0: -66}
                    var v15 int32 = -2147483648
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v2 = v2
                        _ = fi16
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v2 = (v1 / (((v2 % ((v2) | 1))) | 1))
            }
        } else {
            var v17 int16 = 0
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 uint64 = 42
                var v20 uint16 = v2
                var v21 float32 = (-2.25 * 1.5)
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v17
        }
        switch "\n\t" {
        case "\n\t":
            switch (("\n\t" + ("hello" + "\n\t")) + "") {
            case "a":
                var v22 uint = 18446744073709551615
                _ = v22
            case "hello":
                var v23 uint = genFn1(402, 32767, 42)
                _ = v23
            default:
                v4 -= (-(v4))
            }
        default:
            var v24 int = 249
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 13; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 int64 = (v4 << ((0) & 63))
                var v28 string = "hello"
                fmt.Println(v28, fi25)
                _ = fi26
                _ = v27
                _ = v28
            }
            v1 = v1
            _ = fi25
        }
        var v29 int64 = (^(v4))
        delete(v3, "a")
        _ = v29
    }
    var v30 uint8 = 247
    if (!(!(true && (2 != 42)))) {
        var v31 int64 = v4
        _ = v31
    }
    var v32 int16 = 1
    fmt.Println(v30)
    var v33 int = 0
    fmt.Println(v33, v4)
    v34 := []int{}
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 uint32 = 65535
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    var v38 uint16 = v1
    for fi39 := 0; fi39 < 9; fi39++ {
        switch ("hello" + ("foo bar" + (((("" + "foo bar") + "hello") + "\n\t") + "foo bar"))) {
        case "hello":
            var v40 uint = 0
            _ = v40
        default:
            if (false || ((127 < 126) && true)) {
                if ((-2147483648 < 277) || true) {
                    var v41 string = "hello"
                    for fi42 := 0; fi42 < 16; fi42++ {
                        var v43 int64 = -1
                        _ = fi42
                        _ = v43
                    }
                    _ = v41
                }
                v33 += fi39
            } else {
                fmt.Println(v38)
                switch fi39 {
                case -9223372036854775808:
                    v32 = v32
                case -594:
                    v34 = append(v34, (fi39 / ((v33) | 1)))
                case 9223372036854775807:
                    v32 = -1
                default:
                    if false {
                        var v44 int64 = (v4 - 0)
                        v32 -= v32
                        _ = v44
                    } else {
                        v45 := GenS2{F0: (v32 >> ((2) & 63)), F1: ("foo bar" + "\n\t"), F2: -305}
                        fmt.Println(v2, v4)
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v47 := GenS1{F0: int8(v30)}
                            _ = fi46
                            _ = v47
                        }
                        _ = v45
                    }
                }
                fmt.Println(v30, v2)
            }
        }
        _ = fi39
    }
    var v48 float32 = (3.14159 * ((1e10 - 0.0) + 0.0))
    v30 = v30
    v49 := v3["foo bar"]
    _ = len(v34)
    var v50 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v48
    _ = v49
    _ = v50
}
