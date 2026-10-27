package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0() int {
    var v6 float64 = 1.5
    v7 := []int{}
    v8 := GenS1{F0: 42, F1: 128}
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint32 {
    if false {
        if true {
            switch "foo bar" {
            case "foo bar":
                var v9 string = "foo bar"
                _ = v9
            default:
                if true {
                    for fi10 := 0; fi10 < 19; fi10++ {
                        if (!(0 < -1)) {
                            v11 := []int{-2147483648, 9223372036854775807, genFn0()}
                            var v12 int32 = -2147483648
                            _ = v11
                            _ = v12
                        }
                        var v13 float32 = 3.14159
                        _ = fi10
                        _ = v13
                    }
                } else {
                    v14 := GenS1{F0: 32767, F1: 586}
                    if (9223372036854775807 >= 0) {
                        var v15 uint64 = 42
                        v16 := []int{127, genFn0()}
                        var v17 string = "a"
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                }
            }
            var v18 uint8 = 128
            _ = v18
        } else {
            switch "foo bar" {
            case "hello":
                v19 := []int{}
                _ = v19
            default:
                var v20 float32 = float32(float64(1.5))
                _ = v20
            }
        }
    } else {
        var v21 int8 = 100
        _ = v21
    }
    v22 := []int{42, 239, -2147483648, -9223372036854775808}
    _ = v22
    return 144
}

func main() {
    var v1 int64 = 1
    var v2 string = "a"
    var v3 int64 = 42
    var v4 float64 = 0.0
    v5 := make(map[uint32]bool)
    var v23 float64 = v4
    v2 = v2
    v24 := GenS1{F0: -1, F1: -639}
    var v25 int16 = 513
    switch v2 {
    case "foo bar":
        v25 -= v25
    case "\n\t":
        v26 := []int{383, 136, -9223372036854775808, 9223372036854775807}
        _ = v26
    }
    var v27 uint32 = 65535
    for fi28 := 0; fi28 < 19; fi28++ {
        v2 = "foo bar"
        _ = fi28
    }
    v1 += (v3 / (((v1 >> ((42) & 63))) | 1))
    var v29 bool = (false && false)
    if v29 {
        v2 = ("\n\t" + v2)
        var v30 int64 = v3
        _ = v30
    }
    v31 := GenS2{F0: v1, F1: v27}
    var v32 uint64 = 0
    switch 9223372036854775807 {
    case -662:
        v27 = v27
    case -534:
        for fi33 := 0; fi33 < 11; fi33++ {
            fmt.Println(v3)
            switch fi33 {
            case -601:
                var v34 uint8 = uint8(v32)
                _ = v34
            default:
                fmt.Println(v3, v4)
            }
            _ = fi33
        }
    case 127:
        v25 = (v25 & int16(v3))
    }
    if (1 <= 42) {
        switch v2 {
        case "\n\t":
            var v35 int = -559
            _ = v35
        case "a":
            switch -330 {
            case -629:
                v32 = 2
            case 127:
                delete(v5, genFn1())
            case 128:
                v25 = (^(v25))
            }
        default:
            var v36 int16 = (v25 * 42)
            _ = v36
        }
        if v29 {
            if v29 {
                v37 := GenS2{F0: v3, F1: (v27 % ((v27) | 1))}
                switch ("" + (((v2 + v2) + (v2 + ("hello" + "hello"))) + (v2 + ""))) {
                case "hello":
                    if (3.14159 < 1e10) {
                        fmt.Println(v27, v2, v29)
                    } else {
                        var v38 uint = 440
                        for fi39 := 0; fi39 < 11; fi39++ {
                            _ = len(v5)
                            _ = fi39
                        }
                        _ = v38
                    }
                case "":
                    delete(v5, v27)
                }
                v5[v27] = false
                _ = v37
            } else {
                v40 := GenS2{F0: (v1 << ((550) & 63)), F1: genFn1()}
                _ = v40
            }
            for fi41 := 0; fi41 < 15; fi41++ {
                fmt.Println(v23, v25)
                var v42 uint = 65535
                _ = fi41
                _ = v42
            }
            fmt.Println(v2, v2, v3)
        }
    } else {
        var v43 float64 = v23
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
