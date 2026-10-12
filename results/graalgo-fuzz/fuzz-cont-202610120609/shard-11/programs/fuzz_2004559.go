package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64) int16 {
    var v4 int32 = -257
    switch 127 {
    case 814:
        switch "" {
        case "\n\t":
            var v5 float32 = (0.1 + -2.25)
            _ = v5
        case "foo bar":
            var v6 int8 = -61
            _ = v6
        default:
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int32 = v4
                v8 -= (v8 | v4)
                var v9 float64 = (1.5 + 0.0)
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    case 127:
        var v10 uint32 = 42
        _ = v10
    default:
        v11 := GenS1{F0: v4}
        _ = v11
    }
    v12 := GenS1{F0: v4}
    v13 := GenS2{F0: 4294967295}
    if true {
        p0 -= (p0 >> ((18446744073709551615) & 63))
    } else {
        v4 = -2147483648
        var v14 bool = true
        if false {
            v14 = false
            v15 := []int{}
            if len(v15) > 0 {
                v15[0] = -226
            }
            _ = v15
        } else {
            if v14 {
                for fi16 := 0; fi16 < 8; fi16++ {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v14 = v14
                        v18 := []int{2147483647, (^(fi17)), (fi17 - fi16)}
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                var v19 int64 = (-(p0))
                var v20 uint8 = 0
                _ = v19
                _ = v20
            } else {
                v21 := []int{}
                v21 = append(v21, 127)
                _ = v21
            }
            p0 = (-(p0))
        }
        _ = v14
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v4
    return 1
}

func main() {
    var v1 uint32 = 435
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v1 = (^(v1))
    var v22 uint16 = 0
    var v23 int16 = 1000
    v22 = v22
    fmt.Println(v1)
    if false {
        v24 := []int{}
        _ = v24
    } else {
        switch ("\n\t" + ("foo bar" + "foo bar")) {
        case "":
            switch ("" + "foo bar") {
            case "hello":
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 int64 = -1
                    v26 -= v26
                    _ = fi25
                    _ = v26
                }
            }
        default:
            v22 += 2
        }
        switch 0 {
        case 32767:
            v23 -= 1
        default:
            var v27 uint16 = (v22 / ((v22) | 1))
            _ = v27
        }
        var v28 uint8 = 42
        _ = v28
    }
    if ((3.14159 - ((1.5 - (1e10 * 1.5)) + 0.0)) != 3.14159) {
        v1 = v1
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS1{F0: 1}
        v31 := GenS1{F0: -2147483648}
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := GenS1{F0: -645}
    v1 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
}
