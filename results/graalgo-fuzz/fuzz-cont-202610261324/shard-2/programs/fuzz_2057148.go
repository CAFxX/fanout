package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 uint8) uint16 {
    var v5 float32 = 0.0
    var v6 int32 = -2147483648
    fmt.Println(v5, p0)
    v6 += v6
    _ = p0
    _ = v5
    _ = v6
    return 32767
}
func genFn1() uint64 {
    var v7 int8 = 42
    v7 += 100
    for fi8 := 0; fi8 < 19; fi8++ {
        switch "foo bar" {
        case "\n\t":
            v9 := GenS1{F0: 736, F1: 151, F2: 42}
            _ = v9
        case "hello":
            v10 := GenS1{F0: 942, F1: 0, F2: 0}
            _ = v10
        }
        _ = fi8
    }
    var v11 uint8 = 128
    _ = v11
    _ = v7
    return 0
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 int32 = -2147483648
    v4 -= -507
    if ((true || true) || (42 >= int64(v1))) {
        if (genFn1() != 4294967295) {
            v2[719] = v1
        } else {
            switch (v1 << ((18446744073709551615) & 63)) {
            case 32768:
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS1{F0: 42, F1: 105, F2: -9223372036854775808}
                    fmt.Println(fi12, v1, v4)
                    _ = fi12
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 3; fi14++ {
                    if true {
                        v15 := GenS1{F0: 849, F1: 1, F2: 42}
                        _ = v15
                    } else {
                        switch (fi14 / ((v1) | 1)) {
                        case 1:
                            switch 9223372036854775807 {
                            case 4294967295:
                                var v16 int8 = -22
                                _ = v16
                            case 143:
                                v1 -= (^(v1))
                            }
                        case 600:
                            v4 = (v4 << ((0) & 63))
                        case -9223372036854775808:
                            var v17 int32 = (-(v4))
                            _ = v17
                        }
                        var v18 float32 = 3.14159
                        fmt.Println(v1)
                        _ = v18
                    }
                    _ = fi14
                }
            }
            fmt.Println(v1, v4)
        }
        delete(v2, 65535)
        v19 := GenS1{F0: 65535, F1: 2, F2: 42}
        _ = v19
    }
    var v20 uint16 = 255
    v1 += (v1 + v1)
    var v21 int8 = 1
    if false {
        v22 := []int{}
        v21 = 100
        _ = v22
    } else {
        var v23 int16 = 1000
        v4 -= v4
        _ = v23
    }
    v4 += (-(v4))
    fmt.Println(v20, v1, v20)
    if false {
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(v21, v1)
            _ = fi24
        }
        v1 = (v1 + (v1 % (((v1 ^ (v1 / ((v1) | 1)))) | 1)))
    }
    var v25 bool = true
    var v26 string = "hello"
    switch (v1 % ((-707) | 1)) {
    case -1:
        var v27 string = v26
        _ = v27
    case -403:
        v20 = (v20 - v20)
    default:
        var v28 float32 = 0.1
        _ = v28
    }
    v29 := GenS1{F0: 65535, F1: 201, F2: -1}
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
