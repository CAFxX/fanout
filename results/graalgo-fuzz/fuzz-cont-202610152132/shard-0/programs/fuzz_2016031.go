package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool; F1 int16; F2 uint }

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]uint32)
    v3 := []uint32{0, 65535, 0, 42}
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    v1 = (v1 << ((370) & 63))
    if true {
        v6 := v4[-2147483648]
        v7 := GenS1{F0: -1, F1: 65535}
        _ = v6
        _ = v7
    } else {
        var v8 int16 = 32767
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v9 uint8 = (^(v1))
    switch "hello" {
    case "\n\t":
        var v10 uint8 = v9
        _ = v10
    case "hello":
        if true {
            if false {
                switch ("\n\t" + "a") {
                case "hello":
                    v9 = 128
                case "a":
                    if true {
                        v9 -= v9
                    }
                }
                if (true || true) {
                    var v11 uint8 = v9
                    v12 := GenS2{F0: ((!(!(100 == 42))) || ((("" + "foo bar") == (("a" + "foo bar") + "hello")) || true)), F1: 1000, F2: 1}
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS2{F0: (1.5 < 1e10), F1: 0, F2: 1}
                    fmt.Println(v1, v9)
                    _ = v13
                }
                for fi14 := 0; fi14 < 1; fi14++ {
                    v9 -= 127
                    _ = fi14
                }
            } else {
                v1 = (^(v9))
                if (127 == 971) {
                    fmt.Println(v9)
                    fmt.Println(v1, v9)
                    var v15 float64 = (1.5 - 0.1)
                    _ = v15
                }
                fmt.Println(v1, v1)
            }
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                v1 = (v9 - 127)
                _ = fi16
            }
            _ = len(v3)
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 bool = (!(0 == 990))
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: true, F1: -368, F2: 24}
    v1 = (^(v1))
    v1 = 76
    if true {
        v20 := GenS2{F0: true, F1: -1, F2: 1}
        _ = v20
    }
    if (true && (1 <= 0)) {
        var v21 int64 = 1
        v22 := v4[-2147483648]
        var v23 int8 = 127
        _ = v21
        _ = v22
        _ = v23
    }
    if true {
        switch 1 {
        case 32767:
            fmt.Println(v9)
        case 1:
            var v24 int16 = 1000
            _ = v24
        }
        v9 += v9
    }
    v1 = (v1 << ((470) & 63))
    v1 -= 255
    v9 -= (^(v9))
    if (v1 <= (v9 * (v1 >> ((871) & 63)))) {
        var v25 uint32 = 0
        v3 = append(v3, v25)
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 float32 = 1.5
            var v28 int16 = 0
            if len(v3) > 0 {
                v3[0] = (v25 - 2)
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    v29 := GenS2{F0: true, F1: 0, F2: 42}
    var v30 string = (("hello" + "") + ("" + "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
