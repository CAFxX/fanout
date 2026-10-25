package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 int }
func genFn0() uint16 {
    switch ("hello" + "foo bar") {
    case "":
        var v4 bool = true
        _ = v4
    default:
        var v5 string = "hello"
        _ = v5
    }
    var v6 float64 = 1e10
    v6 -= v6
    switch 0 {
    case 0:
        fmt.Println(v6)
    case 42:
        if true {
            var v7 int8 = 115
            _ = v7
        }
    case 943:
        fmt.Println(v6)
    default:
        fmt.Println(v6)
    }
    _ = v6
    return 1
}
func genFn1(p0 int) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    var v8 int64 = 512
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = 1
            fmt.Println(fi10, v8)
            _ = fi10
            _ = v11
        }
        var v12 uint16 = 396
        _ = fi9
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch fi13 {
        case -686:
            fmt.Println(p0)
        case -1:
            var v14 float64 = (-2.25 + 3.14159)
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v8 = v8
            switch p0 {
            case -2147483648:
                v8 -= 9223372036854775807
            case 1:
                var v16 int16 = -32768
                _ = v16
            case -1:
                var v17 bool = false
                _ = v17
            default:
                fmt.Println(fi13)
            }
            if (!false) {
                fmt.Println(fi13)
                fmt.Println(fi13, fi13, fi13)
                p0 -= int(v8)
            }
            _ = fi15
        }
        if (42 > 42) {
            if false {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v8 = v8
                    var v19 uint64 = 689
                    var v20 uint8 = 141
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            } else {
                var v21 float64 = float64(-2.25)
                var v22 int32 = 42
                var v23 uint8 = 1
                _ = v21
                _ = v22
                _ = v23
            }
            v8 += v8
        } else {
            p0 = p0
            if true {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS1{F0: 65535, F1: 0}
                    v26 := GenS1{F0: 65535, F1: 4294967295}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 int64 = (v8 | (v8 * v8))
                _ = v27
            }
        }
        _ = fi13
    }
    _ = p0
    _ = v8
    return false
}

func main() {
    var v1 int64 = -1
    var v2 bool = false
    var v3 bool = v2
    v2 = v2
    switch 127 {
    case 4294967295:
        var v28 uint64 = 462
        _ = v28
    case 42:
        if (0 != -246) {
            if (0 != 422) {
                for fi29 := 0; fi29 < 11; fi29++ {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        fmt.Println(v2)
                        v31 := GenS1{F0: 276, F1: 18446744073709551615}
                        v1 += v1
                        _ = fi30
                        _ = v31
                    }
                    v32 := []int{(fi29 % ((fi29) | 1)), (^(fi29))}
                    _ = fi29
                    _ = v32
                }
            }
        }
    default:
        v1 = v1
    }
    switch 616 {
    case 127:
        var v33 uint16 = genFn0()
        _ = v33
    default:
        v34 := GenS1{F0: 1, F1: 42}
        _ = v34
    }
    v3 = (1 < -128)
    fmt.Println(v3)
    var v35 float32 = 0.0
    for fi36 := 0; fi36 < 20; fi36++ {
        v2 = (("\n\t" + ("foo bar" + "\n\t")) <= ("hello" + "a"))
        for fi37 := 0; fi37 < 5; fi37++ {
            switch (fi37 ^ fi36) {
            case -998:
                v3 = ("a" == ("\n\t" + ""))
            default:
                v35 = v35
            }
            if v2 {
                var v38 uint32 = 499
                v39 := []int{(fi37 + fi37)}
                _ = v38
                _ = v39
            }
            v1 += 1
            _ = fi37
        }
        switch (fi36 + fi36) {
        case 1:
            var v40 float64 = 0.0
            _ = v40
        case 32768:
            v35 -= 1e10
        case -2147483648:
            fmt.Println(v1, v35)
        }
        _ = fi36
    }
    if v3 {
        if v2 {
            v41 := []int{847, 32767}
            var v42 int8 = -100
            _ = v41
            _ = v42
        } else {
            v43 := GenS2{F0: 18446744073709551615, F1: 127, F2: -1}
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
