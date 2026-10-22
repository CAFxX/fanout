package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
func genFn0() string {
    if (!(true || false)) {
        var v7 int = -729
        _ = v7
    } else {
        if false {
            var v8 uint = 279
            _ = v8
        } else {
            var v9 int32 = -121
            _ = v9
        }
    }
    var v10 uint = 992
    _ = v10
    return "a"
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v6 := make(map[uint32]int)
    var v11 uint64 = 4294967295
    if (false || true) {
        for fi12 := 0; fi12 < 13; fi12++ {
            _ = len(v2)
            _ = fi12
        }
        switch "hello" {
        case "\n\t":
            v11 += (v11 * 18446744073709551615)
        default:
            switch 2147483647 {
            case 42:
                var v13 uint32 = 65535
                _ = v13
            case 215:
                var v14 string = ((("" + "\n\t") + (("" + "") + "\n\t")) + ("a" + ("\n\t" + ("foo bar" + (("" + ("hello" + ("a" + "foo bar"))) + "")))))
                _ = v14
            default:
                switch 4294967295 {
                case -708:
                    fmt.Println(v1, v11)
                }
            }
        }
        if true {
            v1 = (v1 % ((18446744073709551615) | 1))
        }
    }
    v11 -= 42
    var v15 float64 = 0.0
    var v16 uint64 = 2
    v11 -= 48
    for fi17 := 0; fi17 < 14; fi17++ {
        switch "hello" {
        case "foo bar":
            switch "a" {
            case "\n\t":
                fmt.Println(v1, v11)
            case "hello":
                var v18 float64 = float64((3.14159 + 0.1))
                _ = v18
            }
        case "a":
            var v19 uint64 = (v16 - (^(v11)))
            _ = v19
        default:
            if (65535 != 200) {
                if false {
                    var v20 uint8 = 2
                    _ = v20
                }
                switch fi17 {
                case -773:
                    v21 := GenS1{F0: 42, F1: 100, F2: 18446744073709551615}
                    _ = v21
                case 46:
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 float64 = v15
                        if true {
                            v16 -= v16
                        }
                        var v24 int8 = 1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                var v25 bool = (0 <= 9223372036854775807)
                _ = v25
            }
        }
        _ = fi17
    }
    v16 -= (v16 >> ((4294967295) & 63))
    switch 2147483647 {
    case 522:
        var v26 uint64 = (v1 << ((4294967295) & 63))
        _ = v26
    case 0:
        switch 42 {
        case 9223372036854775807:
            fmt.Println(v15, v15)
        default:
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := []int{(fi27 * fi27), -2147483648, fi27}
                fmt.Println(fi27)
                var v29 uint64 = v1
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    case 32767:
        for fi30 := 0; fi30 < 20; fi30++ {
            v11 = v1
            var v31 uint64 = 4294967295
            _ = fi30
            _ = v31
        }
    default:
        v32 := []int{32767, -1, 128}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
