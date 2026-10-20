package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0() bool {
    var v4 float64 = (-2.25 - (1e10 * 1e10))
    var v5 uint32 = 969
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 uint16) uint32 {
    p0 += 569
    v6 := GenS1{F0: 255, F1: "\n\t", F2: -822}
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 1
    var v3 int32 = 42
    v3 = (v3 << ((65535) & 63))
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float32 = 1e10
        _ = fi7
        _ = v8
    }
    fmt.Println(v3)
    v3 = v3
    var v9 int = v2
    var v10 float32 = 1e10
    if true {
        v10 = v10
        v3 = v3
        if false {
            v9 = -2147483648
            switch (^(v9)) {
            case 127:
                var v11 int32 = v3
                _ = v11
            case -1:
                v12 := []int{}
                _ = v12
            default:
                v13 := []int{(v2 - (v2 | v2)), v9, v9, 127}
                _ = v13
            }
        } else {
            switch (v9 << ((4294967295) & 63)) {
            case 32767:
                v14 := GenS1{F0: 65535, F1: "hello", F2: 687}
                _ = v14
            case -9223372036854775808:
                var v15 uint = 369
                _ = v15
            case -277:
                v3 -= v3
            default:
                var v16 float32 = 1e10
                _ = v16
            }
        }
    } else {
        v3 += v3
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        switch 32768 {
        case 42:
            v18 := []int{(v2 & (v2 * v2))}
            _ = v18
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := []int{v2, v2, 32768, v9}
            _ = fi19
            _ = v20
        }
        var v21 int16 = 32767
        _ = fi17
        _ = v21
    }
    var v22 uint = 65535
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 int8 = -1
        v2 += v2
        if (42 <= -1) {
            v24 = -128
        } else {
            for fi25 := 0; fi25 < 5; fi25++ {
                fmt.Println(v22)
                fmt.Println(v3, v3)
                switch "\n\t" {
                case "foo bar":
                    if ((("foo bar" + "\n\t") < "\n\t") && (-47 <= (v24 * (v24 >> ((v22) & 63))))) {
                        v22 -= v22
                        var v26 uint8 = 0
                        _ = v26
                    } else {
                        switch "" {
                        case "foo bar":
                            delete(v1, 2)
                        case "":
                            var v27 uint64 = 0
                            _ = v27
                        }
                        fmt.Println(fi23)
                        v2 = fi23
                    }
                case "a":
                    fmt.Println(fi23, v9, v10)
                }
                _ = fi25
            }
            switch v9 {
            case 934:
                v3 += v3
            }
        }
        _ = fi23
        _ = v24
    }
    if true {
        v22 = (v22 ^ 65535)
        v2 = v2
    }
    v9 -= int(v3)
    for fi28 := 0; fi28 < 5; fi28++ {
        switch "\n\t" {
        case "a":
            v3 += (v3 - (-(v3)))
        case "hello":
            delete(v1, 2147483648)
        case "foo bar":
            fmt.Println(v3, v10, v2)
        }
        switch "" {
        case "a":
            v10 -= (v10 + 0.0)
        case "foo bar":
            v29 := GenS1{F0: v22, F1: "foo bar", F2: (v9 >> (((v22 - 345)) & 63))}
            _ = v29
        case "":
            v22 = (v22 * (v22 / ((v22) | 1)))
        default:
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := GenS1{F0: 387, F1: "\n\t", F2: (fi30 ^ fi28)}
                v10 = v10
                if (1 != 42) {
                    v32 := GenS2{F0: 42, F1: 4294967295, F2: 2}
                    _ = v32
                } else {
                    fmt.Println(v10, v3, fi30)
                    if true {
                        v33 := GenS1{F0: 18446744073709551615, F1: ("foo bar" + "foo bar"), F2: -509}
                        v34 := GenS2{F0: 0, F1: 18446744073709551615, F2: 765}
                        _ = v33
                        _ = v34
                    }
                    delete(v1, 0)
                    var v35 uint16 = 255
                    _ = v35
                }
                _ = fi30
                _ = v31
            }
        }
        _ = fi28
    }
    var v36 int64 = -9223372036854775808
    var v37 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v9
}
