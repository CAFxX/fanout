package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{(fi7 ^ (fi7 & 694)), fi7, fi7}
        var v9 uint = 18446744073709551615
        fmt.Println(fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = (0.1 + (3.14159 - ((((3.14159 - (1.5 / 0.1)) + 3.14159) + 0.0) * (float64(3.14159) - (-2.25 - 0.1)))))
    _ = v10
    return 189
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (-1 < 49)
    var v3 uint16 = 42
    var v4 uint32 = 2147483648
    v5 := []uint32{4294967295, (v4 << ((18446744073709551615) & 63)), v4}
    var v6 uint32 = v4
    if ((0 <= 2) || ((^(v3)) != (v3 % (((v3 ^ 65535)) | 1)))) {
        fmt.Println(v3)
        var v11 float64 = (((0.0 / (1e10 * 3.14159)) / 3.14159) / 0.1)
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if (3.14159 <= (1.5 + 0.0)) {
            switch 9223372036854775807 {
            case 128:
                var v13 uint = 673
                _ = v13
            case 32768:
                v14 := GenS1{F0: 2}
                _ = v14
            default:
                v2 = v2
            }
        }
        if ((((297 == 923) && ((fi12 & (fi12 / (((fi12 % (((fi12 ^ 42)) | 1))) | 1))) == fi12)) || (!(488 != 4294967295))) && v2) {
            if (false || v2) {
                var v15 uint8 = 128
                v15 = 42
                var v16 int64 = 0
                _ = v15
                _ = v16
            }
        } else {
            v1["hello"] = true
            v3 -= 1
            v3 = (v3 >> ((0) & 63))
        }
        if len(v5) > 0 {
            v5[0] = v6
        }
        _ = fi12
    }
    switch -1 {
    case 32767:
        var v17 uint16 = v3
        _ = v17
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := GenS1{F0: 1}
                v21 := v1["a"]
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                switch 42 {
                case 42:
                    var v23 int32 = 1
                    if false {
                        var v24 uint64 = 917
                        _ = v24
                    }
                    _ = v23
                default:
                    var v25 int16 = 32767
                    _ = v25
                }
                var v26 uint64 = 18446744073709551615
                v27 := GenS1{F0: (v26 & (v26 % (((v26 + (v26 ^ 0))) | 1)))}
                var v28 uint16 = 0
                _ = fi22
                _ = v26
                _ = v27
                _ = v28
            }
            _ = fi18
        }
    }
    v29 := GenS1{F0: 1}
    v30 := GenS1{F0: 4294967295}
    fmt.Println(v6, v6)
    var v31 int = -445
    v2 = false
    v32 := GenS2{F0: "a"}
    v4 -= (^(v6))
    v3 += v3
    var v33 int = v31
    var v34 float32 = 1e10
    for fi35 := 0; fi35 < 17; fi35++ {
        if v2 {
            v36 := GenS2{F0: "\n\t"}
            delete(v1, (("\n\t" + (("hello" + (("\n\t" + "hello") + "\n\t")) + "hello")) + (("\n\t" + "\n\t") + "hello")))
            if false {
                v34 -= v34
                v37 := GenS2{F0: "a"}
                _ = v37
            }
            _ = v36
        } else {
            var v38 int64 = -9223372036854775808
            _ = v38
        }
        _ = fi35
    }
    var v39 uint32 = (v4 ^ v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
