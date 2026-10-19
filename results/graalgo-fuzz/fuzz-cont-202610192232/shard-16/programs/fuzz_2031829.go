package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 uint = 2
            _ = fi5
            _ = v6
        }
        var v7 float64 = (((1e10 + 3.14159) + 0.1) / 0.1)
        v8 := GenS1{F0: ("a" + ("\n\t" + ("\n\t" + ("a" + ("" + "foo bar")))))}
        _ = fi4
        _ = v7
        _ = v8
    }
    var v9 int = -2147483648
    var v10 int64 = -892
    _ = v10
    _ = v9
    return 255
}

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    var v3 uint8 = 127
    if (v1 || true) {
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v1)
            fmt.Println(fi11, v3)
            var v12 int64 = -9223372036854775808
            _ = fi11
            _ = v12
        }
        v13 := GenS2{F0: 0, F1: -1, F2: 0}
        v3 -= (v3 * 128)
        _ = v13
    } else {
        v3 -= v3
        v14 := []int{}
        _ = v14
    }
    if v1 {
        v15 := []int{}
        _ = v15
    }
    switch 0 {
    case 32767:
        v1 = v1
    default:
        v2[-827] = 1
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        switch ((("a" + (("hello" + "") + ("" + ("\n\t" + "\n\t")))) + "") + "hello") {
        case "a":
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(fi17, v1, v1)
                v3 = 132
                var v18 int32 = 0
                var v19 float64 = 3.14159
                var v20 uint = 65535
                v21 := v2[(fi16 & 768)]
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            }
        default:
            v22 := []int{399, -169, (^(fi16)), -1}
            _ = v22
        }
        _ = fi16
    }
    if (!v1) {
        var v23 int = 128
        v23 = (v23 << ((0) & 63))
        var v24 string = ("foo bar" + "foo bar")
        _ = v23
        _ = v24
    } else {
        fmt.Println(v1)
        v3 = 127
    }
    switch "foo bar" {
    case "\n\t":
        v3 = v3
    case "a":
        var v25 uint = 18446744073709551615
        _ = v25
    default:
        var v26 uint16 = 65535
        _ = v26
    }
    if v1 {
        v27 := GenS3{F0: -49, F1: false, F2: 127}
        v28 := GenS2{F0: -9223372036854775808, F1: -2147483648, F2: 959}
        v3 -= 128
        _ = v27
        _ = v28
    } else {
        var v29 int16 = 42
        v30 := []int{}
        var v31 uint8 = 2
        _ = v29
        _ = v30
        _ = v31
    }
    v3 = (^(v3))
    v1 = ((v1 && (!false)) && v1)
    for fi32 := 0; fi32 < 1; fi32++ {
        if true {
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 uint8 = v3
                v34 = 0
                var v35 uint16 = 390
                _ = fi33
                _ = v34
                _ = v35
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                fmt.Println(v1, v3)
                _ = fi36
            }
            var v37 int64 = 547
            _ = v37
        }
        _ = fi32
    }
    v1 = v1
    var v38 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
}
