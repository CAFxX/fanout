package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
func genFn0() int8 {
    var v7 string = ""
    if true {
        var v8 uint = 910
        var v9 float64 = 1e10
        v8 -= 87
        _ = v8
        _ = v9
    } else {
        if true {
            v10 := []int{128, 0, 42}
            _ = v10
        }
        var v11 int = 1
        var v12 int16 = 1000
        _ = v11
        _ = v12
    }
    v13 := []int{-182}
    var v14 string = v7
    _ = v13
    _ = v14
    _ = v7
    return 100
}

func main() {
    var v1 float32 = (3.14159 / 3.14159)
    v2 := []int{-2147483648}
    v3 := make(map[uint32]bool)
    v4 := []uint32{758, 928, 0}
    v5 := []int{128, -27, -767, 4294967295}
    var v6 int32 = 42
    fmt.Println(v1)
    var v15 float32 = 3.14159
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            v1 -= 0.0
            _ = fi17
        }
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint = 0
            v6 = v6
            if len(v2) > 0 {
                v2[0] = fi18
            }
            if (false || ((!(127 != 45)) || true)) {
                _ = len(v3)
            } else {
                v1 = 0.1
                _ = len(v3)
            }
            _ = fi18
            _ = v19
        }
        if true {
            v20 := GenS1{F0: 2, F1: 1}
            if ((-2.25 + float64((((1.5 + v1) + ((v1 + ((0.0 * ((0.1 + v15) * v1)) - float32(-2.25))) - -2.25)) * v1))) >= 1e10) {
                fmt.Println(v1)
                v6 = (v6 >> ((255) & 63))
                if false {
                    var v21 string = ""
                    _ = v21
                }
            } else {
                v22 := GenS1{F0: 32767, F1: 36}
                var v23 string = "hello"
                _ = v22
                _ = v23
            }
            if (!false) {
                for fi24 := 0; fi24 < 10; fi24++ {
                    fmt.Println(fi16)
                    _ = fi24
                }
                v25 := GenS1{F0: 513, F1: -608}
                fmt.Println(v15)
                _ = v25
            } else {
                delete(v3, 4294967295)
            }
            _ = v20
        }
        _ = fi16
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 int8 = 127
        var v28 uint64 = 4294967295
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        fmt.Println(fi29, v6, v1)
        var v30 int64 = 42
        _ = fi29
        _ = v30
    }
    switch "" {
    case "hello":
        switch ("hello" + "a") {
        case "\n\t":
            v5 = append(v5, 127)
        case "foo bar":
            v5 = append(v5, 682)
        case "":
            if false {
                switch (((((("\n\t" + "") + "a") + ("\n\t" + "a")) + ("\n\t" + "\n\t")) + "") + "a") {
                case "":
                    v15 -= (v1 / 1.0)
                default:
                    if len(v2) > 0 {
                        v2[0] = 128
                    }
                }
                for fi31 := 0; fi31 < 16; fi31++ {
                    if true {
                        var v32 string = "a"
                        var v33 bool = (-128 != -74)
                        v34 := GenS1{F0: 738, F1: 0}
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        v1 += float32(1e10)
                    }
                    var v35 int16 = -32768
                    _ = fi31
                    _ = v35
                }
                v36 := GenS2{F0: true, F1: 32767, F2: false}
                _ = v36
            }
        }
    default:
        fmt.Println(v6, v1)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        v38 := GenS2{F0: (131 <= 242), F1: 255, F2: (-273 == -9223372036854775808)}
        if (!false) {
            delete(v3, 794)
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
