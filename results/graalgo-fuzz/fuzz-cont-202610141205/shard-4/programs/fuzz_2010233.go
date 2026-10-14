package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = ("\n\t" + ("a" + ("" + "foo bar")))
    var v3 uint64 = 891
    v4 := make(map[string]bool)
    var v5 uint = 1
    v1 -= v1
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch fi7 {
            case 1:
                var v8 bool = (float64((1.5 - 0.1)) == ((3.14159 * 0.1) - 1e10))
                _ = v8
            case 9223372036854775807:
                v9 := []int{}
                _ = v9
            default:
                v10 := []int{(-(fi6))}
                _ = v10
            }
            if ((v2 + v2) != "") {
                if (!true) {
                    v2 = v2
                    fmt.Println(fi6)
                }
                var v11 uint32 = 637
                _ = v11
            }
            _ = fi7
        }
        fmt.Println(v1, v2, v5)
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 = (v1 * v1)
            fmt.Println(fi6, v5, fi12)
            var v13 int8 = -27
            _ = fi12
            _ = v13
        }
        _ = fi6
    }
    fmt.Println(v3, v3, v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        switch ("foo bar" + v2) {
        case "":
            if true {
                fmt.Println(v2, v1)
                v15 := GenS1{F0: (fi14 * 9223372036854775807), F1: 944, F2: 2147483647}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := GenS2{F0: (128 != uint8(v5)), F1: -128, F2: (v1 << (((^(v5))) & 63))}
                    _ = fi16
                    _ = v17
                }
                v5 -= 18446744073709551615
            }
        }
        v18 := v4[""]
        for fi19 := 0; fi19 < 13; fi19++ {
            v4[(v2 + v2)] = (v18 || v18)
            var v20 float64 = ((((3.14159 * ((1.5 * 3.14159) / 1.0)) - (1e10 + (float64(float32(float64(1.5))) + (1e10 * 3.14159)))) + float64(1e10)) / 1.0)
            var v21 uint8 = uint8(v3)
            var v22 bool = (v18 && (-682 > 0))
            v5 = v5
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi14
        _ = v18
    }
    if false {
        switch 703 {
        case 87:
            if false {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    if ((992 != v5) && true) {
                        switch 298 {
                        case 2147483647:
                            v23 := GenS1{F0: 4294967295, F1: 42, F2: 1}
                            _ = v23
                        default:
                            v24 := []int{-494}
                            _ = v24
                        }
                    }
                case "":
                    var v25 float64 = 0.1
                    _ = v25
                case "hello":
                    var v26 int64 = -9223372036854775808
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 8; fi27++ {
                        fmt.Println(v5, v1, v3)
                        if (true || true) {
                            v1 -= (^(v1))
                        }
                        _ = fi27
                    }
                }
                v5 -= 42
            } else {
                v2 = v2
                v28 := GenS2{F0: (-1 <= 127), F1: 0, F2: -1}
                var v29 float64 = float64(0.0)
                _ = v28
                _ = v29
            }
        case 127:
            var v30 int = 32767
            _ = v30
        }
    }
    fmt.Println(v3)
    v31 := GenS1{F0: 371, F1: -1, F2: -569}
    for fi32 := 0; fi32 < 20; fi32++ {
        v33 := GenS2{F0: ((42 != 255) || (868 != int16(v5))), F1: -105, F2: v1}
        _ = fi32
        _ = v33
    }
    fmt.Println(v3)
    if false {
        fmt.Println(v1, v1)
        fmt.Println(v5)
        var v34 float32 = -2.25
        _ = v34
    } else {
        v2 = v2
    }
    v5 = v5
    fmt.Println(v3, v3, v1)
    fmt.Println(v2, v3)
    var v35 uint8 = 42
    var v36 int64 = v1
    var v37 float64 = 3.14159
    var v38 uint = (v5 % ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
