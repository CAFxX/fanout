package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    var v3 uint = 306
    var v4 uint16 = 1
    var v5 uint8 = 255
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 string = "a"
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: v3, F1: 42}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 bool = (false && ((v1 | (v1 | (v1 << (((v3 ^ (v3 % (((v3 % (((v3 >> ((1) & 63))) | 1))) | 1)))) & 63)))) != v1))
        v12 := v2[0]
        if (0 != 1) {
            fmt.Println(v3, v4, v3)
        } else {
            if v11 {
                v5 = (^(v5))
            }
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "hello"
    if false {
        if false {
            var v14 int32 = -603
            fmt.Println(v5)
            _ = v14
        } else {
            v15 := []int{42}
            if (float64(-2.25) < (0.1 + 0.1)) {
                fmt.Println(v13, v1, v13)
                v16 := GenS2{F0: 42}
                v15 = append(v15, 445)
                _ = v16
            }
            _ = v15
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v1 -= v1
        _ = fi17
    }
    v3 = v3
    if false {
        v1 = (^(v1))
    } else {
        switch "\n\t" {
        case "a":
            delete(v2, 1)
        default:
            var v18 float32 = 1e10
            _ = v18
        }
        if true {
            v6[9223372036854775807] = 42
        } else {
            v13 = (((v13 + (v13 + ((v13 + "hello") + ""))) + v13) + (("foo bar" + "hello") + v13))
            switch 858 {
            case -2147483648:
                v19 := GenS1{F0: (v3 ^ (v3 - v3)), F1: -857}
                _ = v19
            case 1:
                var v20 int64 = 1
                _ = v20
            }
            if (true && true) {
                switch (v13 + "hello") {
                case "hello":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if (((!true) && false) && (((float32(0.1) * 0.1) + 0.0) == (1.5 * -2.25))) {
                            v5 += 127
                        }
                        if ((!false) && false) {
                            var v22 int = 0
                            var v23 uint = 65
                            _ = v22
                            _ = v23
                        } else {
                            v5 += v5
                        }
                        var v24 bool = (!false)
                        _ = fi21
                        _ = v24
                    }
                }
                fmt.Println(v1, v5, v5)
                v25 := GenS2{F0: 42}
                _ = v25
            }
        }
    }
    var v26 uint8 = (v5 ^ (v5 - (v5 % ((v5) | 1))))
    for fi27 := 0; fi27 < 18; fi27++ {
        switch 2147483647 {
        case -1:
            var v28 int = (fi27 ^ (^(fi27)))
            _ = v28
        case -784:
            fmt.Println(v5, v3)
        case 2147483647:
            v5 = v26
        default:
            var v29 string = "\n\t"
            _ = v29
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v26, v4, v13)
        _ = fi30
    }
    var v31 float32 = float32(1e10)
    v5 = (^(v5))
    var v32 bool = false
    if ((v32 && v32) || true) {
        v13 = v13
    } else {
        v33 := []int{4294967295, int(v26), 87, 32768}
        v13 = ("foo bar" + (v13 + "foo bar"))
        _ = v33
    }
    switch 127 {
    case 82:
        v34 := []int{-2147483648, 1, -536}
        _ = v34
    case 42:
        v35 := []int{2147483647, 127, 9223372036854775807, 127}
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        switch 32768 {
        case 46:
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 uint8 = uint8(v1)
                _ = fi37
                _ = v38
            }
        case 128:
            fmt.Println(v3, fi36, v4)
        }
        if true {
            for fi39 := 0; fi39 < 1; fi39++ {
                v3 = uint(v4)
                var v40 int8 = 127
                _ = fi39
                _ = v40
            }
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
