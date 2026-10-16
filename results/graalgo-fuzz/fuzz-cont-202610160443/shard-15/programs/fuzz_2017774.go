package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 int, p1 string, p2 uint) uint32 {
    p1 = "hello"
    p1 = "a"
    var v7 uint = p2
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(p0, p1, p2)
        _ = fi8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1() uint64 {
    var v9 int32 = -65
    v10 := []int{9223372036854775807, -929, -607, 2147483647}
    _ = v10
    _ = v9
    return 730
}

func main() {
    v1 := []string{("\n\t" + ""), "foo bar", "foo bar", "foo bar"}
    var v2 int8 = -1
    var v3 int8 = v2
    var v4 int64 = 1
    v5 := make(map[uint32]bool)
    var v6 float32 = (-2.25 + -2.25)
    v11 := GenS1{F0: genFn0(-419, ("a" + "foo bar"), 132), F1: -572}
    var v12 int16 = -1
    v3 = v3
    if true {
        v12 += (v12 + -1)
        v13 := GenS1{F0: 2, F1: 42}
        v2 -= 37
        _ = v13
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: 2, F1: 1}
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v12 -= v12
        var v17 int64 = (v4 / ((v4) | 1))
        for fi18 := 0; fi18 < 11; fi18++ {
            delete(v5, 666)
            var v19 string = "foo bar"
            var v20 float32 = (0.0 - (3.14159 / float32(-2.25)))
            var v21 uint8 = 165
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    var v22 int = 4294967295
    if false {
        _ = len(v5)
        fmt.Println(v12, v4)
        if ((v22 + 32767) != 128) {
            switch -482 {
            case 994:
                _ = len(v1)
            default:
                var v23 float32 = float32(0.0)
                _ = v23
            }
            var v24 int32 = 2147483647
            if true {
                var v25 float64 = -2.25
                var v26 uint64 = 18446744073709551615
                _ = v25
                _ = v26
            } else {
                var v27 uint64 = 170
                var v28 float32 = ((v6 + float32(0.1)) - (v6 - v6))
                _ = v27
                _ = v28
            }
            _ = v24
        } else {
            var v29 uint8 = 155
            for fi30 := 0; fi30 < 3; fi30++ {
                v12 -= 1
                var v31 uint = 65535
                _ = fi30
                _ = v31
            }
            _ = v29
        }
    }
    var v32 uint64 = 4294967295
    var v33 int32 = 1
    var v34 bool = (!false)
    v35 := GenS1{F0: genFn0((v22 * (v22 + (v22 << ((1) & 63)))), "\n\t", 65535), F1: 894}
    if len(v1) > 0 {
        v1[0] = ((("a" + "a") + "a") + "a")
    }
    delete(v5, 1)
    var v36 float32 = (v6 * 0.0)
    var v37 float64 = 0.1
    switch (("\n\t" + ("\n\t" + ("hello" + "\n\t"))) + ("foo bar" + "\n\t")) {
    case "a":
        v2 += 127
    case "foo bar":
        switch ("\n\t" + "a") {
        case "\n\t":
            switch int(v2) {
            case 2147483647:
                if ((1 == 127) && true) {
                    var v38 int64 = (v4 % (((-(v4))) | 1))
                    _ = v38
                }
            case -1:
                for fi39 := 0; fi39 < 6; fi39++ {
                    v34 = false
                    for fi40 := 0; fi40 < 9; fi40++ {
                        v12 = -1
                        v4 = v4
                        v3 = (v2 ^ (v3 % (((v2 - (v2 >> ((832) & 63)))) | 1)))
                        _ = fi40
                    }
                    _ = fi39
                }
            default:
                if (!((("\n\t" + "hello") + "foo bar") <= "\n\t")) {
                    var v41 float64 = 0.1
                    v33 = (v33 / ((1) | 1))
                    _ = v41
                }
            }
        case "foo bar":
            v32 -= (v32 / (((v32 >> ((4294967295) & 63))) | 1))
        case "":
            fmt.Println(v37, v33)
        default:
            fmt.Println(v34, v3)
        }
    }
    v36 -= (v6 * 1e10)
    if v34 {
        delete(v5, 65535)
        var v42 int16 = (v12 / ((v12) | 1))
        for fi43 := 0; fi43 < 15; fi43++ {
            var v44 uint64 = (v32 * v32)
            for fi45 := 0; fi45 < 6; fi45++ {
                v34 = false
                v42 -= v42
                fmt.Println(v34)
                _ = fi45
            }
            var v46 float32 = (((float32(v37) - ((((v6 + v6) / v6) * 1.5) - (0.1 - (v36 - 0.0)))) / v6) + v6)
            _ = fi43
            _ = v44
            _ = v46
        }
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
