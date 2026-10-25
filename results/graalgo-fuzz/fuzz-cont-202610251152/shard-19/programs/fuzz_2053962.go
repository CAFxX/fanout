package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    var v7 uint8 = 1
    fmt.Println(p0, p0)
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "hello":
            var v8 bool = false
            _ = v8
        case "foo bar":
            var v9 int = 42
            _ = v9
        }
    case "hello":
        var v10 int = -1
        _ = v10
    case "":
        v7 = 0
    default:
        v11 := []int{95}
        _ = v11
    }
    _ = p0
    _ = v7
    return -787
}
func genFn1(p0 int8) uint16 {
    p0 -= (p0 + 0)
    for fi12 := 0; fi12 < 6; fi12++ {
        if (9223372036854775807 >= -9223372036854775808) {
            for fi13 := 0; fi13 < 19; fi13++ {
                if (!((((1e10 / (0.1 / (-2.25 + ((3.14159 * 0.0) + (-2.25 - -2.25))))) * -2.25) * (float64(1.5) / 1.5)) != ((1e10 / 3.14159) + 0.0))) {
                    if ((!true) && (65535 != 2)) {
                        var v14 bool = true
                        _ = v14
                    }
                    p0 += int8(fi13)
                    p0 -= (p0 | -1)
                    p0 = (p0 | -1)
                }
                if true {
                    v15 := []int{843, fi13, 32767, fi12}
                    v16 := GenS1{F0: 42, F1: 42, F2: 1}
                    if (fi13 > fi13) {
                        switch fi12 {
                        case -2147483648:
                            var v17 int64 = 352
                            _ = v17
                        case 128:
                            v18 := GenS1{F0: 2, F1: 127, F2: 4294967295}
                            _ = v18
                        }
                    } else {
                        if len(v15) > 0 {
                            v15[0] = (fi12 << ((42) & 63))
                        }
                        switch "" {
                        case "\n\t":
                            fmt.Println(p0)
                        default:
                            v19 := GenS1{F0: 65535, F1: 2, F2: 1}
                            _ = v19
                        }
                    }
                    _ = v15
                    _ = v16
                }
                p0 = (p0 / ((p0) | 1))
                _ = fi13
            }
        } else {
            if (0.0 >= 3.14159) {
                if (65535 == 537) {
                    var v20 uint64 = 0
                    _ = v20
                } else {
                    v21 := []int{}
                    _ = v21
                }
                fmt.Println(fi12, p0)
                p0 = p0
            } else {
                switch (fi12 ^ (fi12 >> ((2) & 63))) {
                case 459:
                    fmt.Println(p0, p0)
                }
                var v22 float64 = 1e10
                _ = v22
            }
        }
        _ = fi12
    }
    var v23 int32 = -408
    var v24 uint32 = 0
    for fi25 := 0; fi25 < 2; fi25++ {
        if false {
            v24 -= v24
        }
        v26 := []int{(fi25 >> ((65535) & 63)), fi25}
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = v23
    _ = v24
    return 32767
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 312
    var v3 int16 = 32767
    var v4 int64 = 0
    v5 := make(map[uint32]string)
    v6 := []uint32{v2}
    var v27 string = ("hello" + "\n\t")
    v28 := v5[1]
    v29 := GenS2{F0: v2}
    v30 := GenS1{F0: 18446744073709551615, F1: 42, F2: 18446744073709551615}
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(fi31)
        var v32 bool = (false || true)
        var v33 int8 = -1
        _ = fi31
        _ = v32
        _ = v33
    }
    v1 += -2147483648
    v2 = (v2 % ((v2) | 1))
    var v34 uint = 18446744073709551615
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 string = v27
        _ = fi35
        _ = v36
    }
    var v37 string = ""
    var v38 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
