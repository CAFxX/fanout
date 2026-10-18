package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }
func genFn0(p0 int8) uint32 {
    var v4 uint16 = 32767
    if true {
        var v5 int = 847
        _ = v5
    } else {
        v6 := []int{-2147483648, 127}
        _ = v6
    }
    switch "a" {
    case "\n\t":
        v4 -= (v4 << ((530) & 63))
    }
    _ = p0
    _ = v4
    return 799
}
func genFn1() int {
    var v7 uint32 = 4294967295
    var v8 uint16 = 212
    var v9 string = ""
    for fi10 := 0; fi10 < 8; fi10++ {
        v8 += (^(v8))
        _ = fi10
    }
    _ = v7
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = 1000
    var v3 int64 = 42
    if (((((((("hello" + "") + ("" + ("hello" + "a"))) + "\n\t") + "") + "\n\t") + "foo bar") + "a") != (("foo bar" + "foo bar") + "hello")) {
        v11 := []int{-511, 42, 42}
        if false {
            var v12 int8 = 30
            _ = v12
        } else {
            v13 := GenS1{F0: true}
            _ = v13
        }
        _ = v11
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{genFn1(), fi14}
        var v16 int16 = (-(v2))
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint = 18446744073709551615
    for fi18 := 0; fi18 < 7; fi18++ {
        if false {
            var v19 int32 = 0
            if ((-(v19)) != (-(v19))) {
                v1 += v1
                v17 -= (v17 ^ 879)
                v2 -= v2
            }
            _ = v19
        } else {
            if ((!true) || (true && (!true))) {
                if false {
                    fmt.Println(v3, v3)
                    fmt.Println(fi18, v17)
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    v17 = 0
                    v21 := GenS2{F0: true}
                    v17 += 65535
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 int = fi18
                switch fi18 {
                case 9223372036854775807:
                    v23 := GenS1{F0: (true || (0 != 421))}
                    _ = v23
                default:
                    var v24 int = (^(v22))
                    _ = v24
                }
                _ = v22
            }
            fmt.Println(v2)
        }
        fmt.Println(v2, v3, v3)
        v25 := GenS3{F0: int16(v3), F1: v1, F2: -1}
        _ = fi18
        _ = v25
    }
    v26 := GenS2{F0: true}
    var v27 int64 = 0
    var v28 uint32 = 2147483648
    for fi29 := 0; fi29 < 10; fi29++ {
        if false {
            var v30 uint = v17
            _ = v30
        }
        fmt.Println(fi29, v27)
        var v31 float64 = 1.5
        _ = fi29
        _ = v31
    }
    switch genFn1() {
    case 42:
        if (!true) {
            v32 := []int{-841, 0}
            var v33 int8 = 127
            _ = v32
            _ = v33
        } else {
            var v34 uint16 = 37
            switch genFn1() {
            case -1:
                v34 = (v34 >> (((v17 >> ((v17) & 63))) & 63))
            }
            _ = v34
        }
    default:
        var v35 string = ""
        _ = v35
    }
    fmt.Println(v28, v17, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
