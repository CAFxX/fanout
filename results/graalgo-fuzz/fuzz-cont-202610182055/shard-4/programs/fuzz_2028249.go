package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = -2.25
    v2 := []string{}
    var v3 bool = true
    var v4 int16 = 1
    var v5 int8 = 126
    v6 := make(map[uint32]uint32)
    v4 = (-(v4))
    switch ("" + "foo bar") {
    case "":
        if (((float64(0.0) + 0.0) < v1) && v3) {
            if (65535 < 255) {
                var v7 uint64 = 2
                _ = v7
            }
        } else {
            v5 -= (-(v5))
            fmt.Println(v1, v5)
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: 42, F1: ("" + "hello")}
                _ = fi8
                _ = v9
            }
        }
    default:
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                v4 = (v4 * v4)
                _ = fi11
            }
            _ = fi10
        }
    }
    switch ("\n\t" + ("\n\t" + "a")) {
    case "\n\t":
        fmt.Println(v1)
    }
    var v12 int32 = -548
    v3 = (!(!((true && ("" == ("hello" + "hello"))) || true)))
    switch 393 {
    case 42:
        fmt.Println(v3, v12)
    case -1:
        var v13 uint8 = 172
        _ = v13
    case 1:
        var v14 float64 = v1
        _ = v14
    }
    var v15 int = 592
    fmt.Println(v3, v1)
    for fi16 := 0; fi16 < 15; fi16++ {
        v3 = ("" < ("a" + ((("\n\t" + ("foo bar" + "foo bar")) + "hello") + ("foo bar" + (("a" + "foo bar") + "\n\t")))))
        _ = fi16
    }
    fmt.Println(v4, v1, v1)
    var v17 float32 = 1.5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (v3 && (!v3)) {
        v18 := GenS1{F0: 0, F1: ("\n\t" + "hello")}
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v12, v17, fi19)
            if true {
                if false {
                    fmt.Println(v3, v17, v12)
                } else {
                    v12 += 42
                    v17 -= float32((((0.0 - v1) / v1) + float64(0.1)))
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 uint32 = 498
                        v22 := GenS1{F0: 65535, F1: "hello"}
                        if v3 {
                            v23 := GenS1{F0: 4294967295, F1: "\n\t"}
                            _ = v23
                        }
                        var v24 uint32 = 2
                        var v25 uint = 42
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v24
                        _ = v25
                    }
                }
                var v26 float64 = ((v1 * float64(-2.25)) - 0.0)
                var v27 int8 = 19
                _ = v26
                _ = v27
            } else {
                v4 -= (v4 ^ 0)
            }
            fmt.Println(v1, v5)
            _ = fi19
        }
        v28 := GenS2{F0: ""}
        _ = v18
        _ = v28
    }
    if v3 {
        var v29 uint32 = 2147483648
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
