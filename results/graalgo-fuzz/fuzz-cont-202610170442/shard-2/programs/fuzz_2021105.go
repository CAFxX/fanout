package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
func genFn0(p0 uint) uint32 {
    fmt.Println(p0)
    p0 = p0
    p0 += 255
    var v4 uint8 = 255
    _ = p0
    _ = v4
    return 326
}

func main() {
    v1 := []int{}
    var v2 int16 = 1000
    v3 := make(map[uint32]bool)
    v5 := GenS1{F0: (!false), F1: 1, F2: 715}
    for fi6 := 0; fi6 < 18; fi6++ {
        delete(v3, 42)
        _ = fi6
    }
    if false {
        var v7 bool = (!(!(!(!false))))
        v2 = v2
        v2 -= v2
        _ = v7
    }
    v2 += -714
    v8 := GenS2{F0: 1, F1: genFn0(365), F2: 769}
    switch "hello" {
    case "":
        fmt.Println(v2)
    }
    fmt.Println(v2)
    var v9 uint64 = 623
    for fi10 := 0; fi10 < 3; fi10++ {
        switch fi10 {
        case 1:
            v11 := GenS2{F0: -127, F1: uint32(v9), F2: 1}
            _ = v11
        case 969:
            switch ("a" + "hello") {
            case "foo bar":
                var v12 uint32 = genFn0(255)
                _ = v12
            case "\n\t":
                if false {
                    v2 += v2
                } else {
                    v13 := GenS2{F0: 577, F1: 2147483648, F2: 116}
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v15 := GenS1{F0: ((-2.25 - 0.0) > (0.1 + (1e10 - 1.5))), F1: 434, F2: (v9 << ((65535) & 63))}
                        v2 += (v2 + 32767)
                        var v16 float32 = 0.0
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                }
            }
        }
        v2 = v2
        v17 := GenS1{F0: true, F1: 65535, F2: 42}
        _ = fi10
        _ = v17
    }
    v9 -= (v9 & v9)
    v9 += v9
    v18 := GenS1{F0: ((v9 ^ (v9 * 881)) != 42), F1: 0, F2: 4294967295}
    delete(v3, 2147483648)
    var v19 bool = false
    for fi20 := 0; fi20 < 7; fi20++ {
        switch (("a" + ("a" + "foo bar")) + (("\n\t" + "\n\t") + ("foo bar" + "\n\t"))) {
        case "hello":
            var v21 int8 = -128
            _ = v21
        case "foo bar":
            if v19 {
                v22 := GenS2{F0: int64(v9), F1: 4294967295, F2: 2}
                _ = v22
            }
        case "":
            v23 := GenS2{F0: -9223372036854775808, F1: 42, F2: 612}
            _ = v23
        }
        _ = fi20
    }
    v9 += (^(v9))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
