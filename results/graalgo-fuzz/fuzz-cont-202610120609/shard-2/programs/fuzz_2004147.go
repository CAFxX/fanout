package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0() uint32 {
    switch "\n\t" {
    case "":
        v7 := GenS1{F0: -507, F1: (0.0 == 1.5)}
        _ = v7
    }
    var v8 string = ("hello" + ("\n\t" + (("\n\t" + "\n\t") + "")))
    v9 := GenS1{F0: -32768, F1: (1 <= 100)}
    var v10 uint16 = 65535
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []int{-2147483648, -241, 127}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 int16 = -1
    var v5 string = "hello"
    var v6 uint = 255
    if true {
        fmt.Println(v6, v6, v6)
    }
    v5 = v5
    var v11 string = (v5 + "")
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v4, v5)
        if ((1e10 * 3.14159) > 0.1) {
            var v13 string = v5
            v14 := GenS1{F0: int16(v6), F1: (2 <= 522)}
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    var v15 int16 = (-(v4))
    var v16 uint = v6
    if ((0 < -1) || true) {
        v17 := GenS1{F0: -587, F1: true}
        if true {
            switch "\n\t" {
            case "\n\t":
                v6 += (v6 % ((42) | 1))
            case "foo bar":
                v11 = ((v11 + v5) + v11)
            case "a":
                switch v5 {
                case "hello":
                    var v18 uint16 = 0
                    _ = v18
                default:
                    v19 := GenS1{F0: (v4 % ((v15) | 1)), F1: (42 != 1)}
                    _ = v19
                }
            }
        } else {
            v20 := GenS2{F0: 18446744073709551615}
            _ = v20
        }
        _ = v17
    }
    var v21 uint32 = 2147483647
    v22 := GenS3{F0: 127}
    var v23 uint = (v6 ^ (v6 ^ 0))
    var v24 uint = (v16 >> ((v6) & 63))
    v4 -= v15
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint64 = 42
        if false {
            if (1.5 >= (-2.25 * 3.14159)) {
                for fi27 := 0; fi27 < 10; fi27++ {
                    switch 4294967295 {
                    case 128:
                        var v28 uint = (v16 & v24)
                        _ = v28
                    case 4294967295:
                        fmt.Println(v23, fi27)
                    case -9223372036854775808:
                        _ = len(v2)
                        v6 -= v6
                    }
                    var v29 uint8 = 255
                    _ = fi27
                    _ = v29
                }
                for fi30 := 0; fi30 < 20; fi30++ {
                    fmt.Println(v23, fi25)
                    _ = fi30
                }
            } else {
                v15 = 0
                var v31 uint32 = (^(v21))
                if true {
                    v16 = (v16 >> ((v6) & 63))
                }
                _ = v31
            }
            var v32 float64 = 0.1
            v26 = v26
            _ = v32
        } else {
            v15 = v15
        }
        v21 += (v21 % (((v21 % ((v21) | 1))) | 1))
        _ = fi25
        _ = v26
    }
    v21 = uint32(v24)
    for fi33 := 0; fi33 < 7; fi33++ {
        if (!false) {
            if ((-1 <= (v15 | (v4 & (v4 ^ (v4 >> ((1) & 63)))))) || (true || (true || false))) {
                var v34 string = ((v11 + "foo bar") + "\n\t")
                v35 := GenS4{F0: (v6 * (v24 >> ((2) & 63))), F1: 255}
                _ = v34
                _ = v35
            }
            fmt.Println(v23, v5, v5)
            var v36 int16 = (v15 * (v15 * v15))
            _ = v36
        } else {
            var v37 float64 = float64(0.1)
            var v38 uint64 = 4294967295
            _ = v37
            _ = v38
        }
        _ = len(v3)
        _ = fi33
    }
    v16 += (^(v6))
    var v39 int64 = -9223372036854775808
    var v40 int8 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
