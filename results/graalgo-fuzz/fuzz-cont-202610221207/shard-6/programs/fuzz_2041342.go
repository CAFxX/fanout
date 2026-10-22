package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 317, F1: (^(fi7))}
        var v9 int32 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 2, F1: 1}
    _ = v10
    return 255
}
func genFn1() int64 {
    v11 := []int{-9223372036854775808, -2147483648, -2147483648}
    var v12 uint8 = 0
    var v13 string = ("hello" + "a")
    v13 = (v13 + (v13 + "hello"))
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 42
    v3 := []string{("\n\t" + "hello")}
    var v4 uint16 = 32767
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    v4 -= v4
    v14 := GenS1{F0: 849, F1: 2147483647}
    var v15 uint = 255
    var v16 int = 42
    v17 := v6[2147483648]
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(v17)
            switch "hello" {
            case "a":
                v19 := GenS1{F0: 4294967295, F1: 127}
                _ = v19
            case "":
                if true {
                    v4 += v4
                }
            case "hello":
                if (42 > -71) {
                    var v20 uint8 = 255
                    _ = v20
                }
            }
            _ = fi18
        }
        if true {
            if true {
                var v21 uint8 = 128
                var v22 int64 = (v2 + genFn1())
                v17 -= (v17 << ((v15) & 63))
                _ = v21
                _ = v22
            } else {
                _ = len(v6)
                switch "\n\t" {
                case "":
                    delete(v1, (v16 << ((v15) & 63)))
                default:
                    var v23 uint16 = 2
                    _ = v23
                }
            }
            v15 = v15
            var v24 uint32 = 65535
            _ = v24
        } else {
            var v25 int32 = -1
            var v26 uint8 = 42
            _ = v25
            _ = v26
        }
        switch "hello" {
        case "a":
            v4 -= v4
        case "hello":
            for fi27 := 0; fi27 < 8; fi27++ {
                if (-535 <= 42) {
                    v16 += (^(v17))
                } else {
                    v28 := GenS1{F0: 214, F1: 127}
                    _ = v28
                }
                var v29 float32 = 1e10
                _ = fi27
                _ = v29
            }
        case "foo bar":
            v2 += (v2 - (-(v2)))
        default:
            fmt.Println(v15, v2, v4)
        }
    } else {
        var v30 int16 = -32768
        if false {
            switch "" {
            case "\n\t":
                if (0.1 >= 1e10) {
                    v16 = (v17 / ((0) | 1))
                    if (1 <= 2147483647) {
                        v31 := GenS1{F0: 367, F1: (v17 ^ v17)}
                        _ = v31
                    }
                }
            case "a":
                v30 = (v30 * 32767)
            case "hello":
                _ = len(v3)
            }
            var v32 uint64 = 18446744073709551615
            switch (v17 << ((531) & 63)) {
            case 32768:
                v4 = (^(v4))
            default:
                v33 := GenS1{F0: uint64(v4), F1: v16}
                _ = v33
            }
            _ = v32
        }
        switch v16 {
        case -2147483648:
            var v34 int32 = 597
            _ = v34
        default:
            v3 = append(v3, "foo bar")
        }
        _ = v30
    }
    v35 := GenS2{F0: 1}
    v4 += (^(v4))
    for fi36 := 0; fi36 < 10; fi36++ {
        v37 := GenS3{F0: 0, F1: (-32768 <= 1000), F2: (^(v4))}
        _ = fi36
        _ = v37
    }
    if (!(100 <= 127)) {
        var v38 int32 = 0
        for fi39 := 0; fi39 < 8; fi39++ {
            v40 := v1[v16]
            v41 := GenS2{F0: 4294967295}
            _ = fi39
            _ = v40
            _ = v41
        }
        v42 := GenS3{F0: -83, F1: false, F2: v4}
        _ = v38
        _ = v42
    }
    v43 := v5["\n\t"]
    for fi44 := 0; fi44 < 9; fi44++ {
        v15 += uint(v43)
        var v45 uint32 = (v43 ^ v43)
        if (!true) {
            v46 := GenS2{F0: (v45 + 2147483648)}
            _ = v46
        }
        _ = fi44
        _ = v45
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        var v48 bool = ((1 != (^(v4))) || (127 < 0))
        v2 = (v2 / (((^(v2))) | 1))
        _ = fi47
        _ = v48
    }
    fmt.Println(v4)
    if true {
        v49 := GenS1{F0: 2, F1: v16}
        _ = v49
    } else {
        var v50 float64 = ((-2.25 - 0.1) / 1.5)
        for fi51 := 0; fi51 < 1; fi51++ {
            if (!(!(!(-2.25 != 3.14159)))) {
                v52 := GenS3{F0: 127, F1: (!((-(v2)) < (v2 >> ((v15) & 63)))), F2: (v4 | (^(v4)))}
                v53 := GenS3{F0: 1, F1: false, F2: v4}
                _ = v52
                _ = v53
            } else {
                var v54 int8 = 91
                for fi55 := 0; fi55 < 19; fi55++ {
                    fmt.Println(v43, fi55, v16)
                    _ = fi55
                }
                _ = v54
            }
            _ = fi51
        }
        _ = v50
    }
    fmt.Println(v17, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
