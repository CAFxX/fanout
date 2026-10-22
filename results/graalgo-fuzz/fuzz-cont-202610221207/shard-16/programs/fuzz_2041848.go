package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint32 = 0
    var v2 uint8 = 34
    v3 := make(map[uint32]int)
    var v4 uint16 = 483
    var v5 float32 = ((1.5 + 1e10) * (-2.25 * 3.14159))
    v6 := []string{("hello" + "hello"), (((("\n\t" + "a") + "hello") + ("\n\t" + "")) + "hello"), "", "hello"}
    if (127 == 1) {
        v5 = ((-2.25 + (v5 + v5)) - float32(1.5))
    } else {
        if true {
            switch "\n\t" {
            case "a":
                var v7 uint8 = (^(v2))
                _ = v7
            case "\n\t":
                v8 := GenS1{F0: 9223372036854775807}
                _ = v8
            }
        } else {
            v9 := GenS1{F0: 1}
            _ = v9
        }
        var v10 int8 = 42
        for fi11 := 0; fi11 < 1; fi11++ {
            if (!(!((v10 / ((v10) | 1)) < (v10 << ((2) & 63))))) {
                fmt.Println(v2, v4)
                v10 -= (v10 % ((v10) | 1))
            }
            for fi12 := 0; fi12 < 10; fi12++ {
                switch ("a" + ((("\n\t" + (("\n\t" + "foo bar") + "hello")) + "") + (((("a" + "hello") + ("hello" + "")) + "foo bar") + "\n\t"))) {
                case "":
                    v6 = append(v6, "hello")
                case "\n\t":
                    v10 += 42
                default:
                    var v13 int16 = 1000
                    _ = v13
                }
                _ = fi12
            }
            _ = fi11
        }
        _ = v10
    }
    v4 += (^(v4))
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int8 = 114
        v3[4294967295] = (fi14 | fi14)
        _ = fi14
        _ = v15
    }
    v1 -= v1
    if true {
        if (!true) {
            fmt.Println(v2, v2, v2)
        } else {
            v16 := v3[v1]
            _ = v16
        }
    } else {
        if (float32(1.5) > 3.14159) {
            for fi17 := 0; fi17 < 16; fi17++ {
                v4 = 32767
                var v18 string = (("\n\t" + "a") + "")
                _ = fi17
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        delete(v3, (v1 / (((^(v1))) | 1)))
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int64 = 747
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v1 += v1
    if true {
        _ = len(v3)
        var v22 uint16 = (v4 >> ((18446744073709551615) & 63))
        v23 := GenS2{F0: 42}
        _ = v22
        _ = v23
    }
    if (v5 == v5) {
        fmt.Println(v4, v2)
        var v24 float32 = ((v5 * v5) * v5)
        v4 = (v4 + (v4 & v4))
        _ = v24
    } else {
        var v25 float32 = v5
        v2 -= (^(v2))
        v26 := v3[v1]
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS1{F0: -9223372036854775808}
        fmt.Println(fi27, fi27, v2)
        _ = fi27
        _ = v28
    }
    v4 += (v4 >> ((1) & 63))
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 string = "a"
        if ((1 >= (v2 | (v2 + uint8(v1)))) || false) {
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(v1, v30, fi29)
                var v32 string = ((v30 + "foo bar") + (v30 + "\n\t"))
                v30 = v32
                _ = fi31
                _ = v32
            }
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v5, v1)
    v6 = append(v6, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
