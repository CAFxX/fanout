package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32, p1 int64) int32 {
    p1 = p1
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 bool = (9223372036854775807 >= -1)
    var v2 int32 = -2147483648
    var v3 uint32 = 4294967295
    v3 += (v3 / ((2) | 1))
    for fi4 := 0; fi4 < 1; fi4++ {
        fmt.Println(v3, v1)
        v1 = true
        _ = fi4
    }
    v1 = ((-2147483648 < -895) || (-2.25 != float32((float64((1.5 + 1e10)) + 1.5))))
    var v5 int = 0
    fmt.Println(v5, v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v5 = 127
            switch (^(v5)) {
            case 32768:
                var v8 uint = 1
                _ = v8
            case -1:
                if false {
                    switch (fi6 ^ fi7) {
                    case 1:
                        var v9 uint = 0
                        _ = v9
                    }
                } else {
                    v10 := GenS1{F0: -836}
                    _ = v10
                }
            case -9223372036854775808:
                v11 := GenS2{F0: 42}
                _ = v11
            default:
                v3 -= (v3 ^ v3)
            }
            _ = fi7
        }
        _ = fi6
    }
    v12 := GenS1{F0: 32767}
    switch "foo bar" {
    case "":
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 float64 = float64(((float32(((0.0 / 3.14159) / (1.5 * -2.25))) - (float32(((1e10 * (float64(1e10) * 0.0)) + -2.25)) * 1e10)) - (1e10 - (-2.25 - float32(1e10)))))
            v1 = v1
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v15 := []int{4294967295, v5}
        _ = v15
    }
    switch v5 {
    case -2147483648:
        v16 := GenS1{F0: 18}
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            v3 = 2
            var v19 uint = 368
            var v20 float64 = 0.1
            switch "" {
            case "\n\t":
                var v21 string = ""
                _ = v21
            default:
                var v22 uint8 = 1
                _ = v22
            }
            v23 := GenS1{F0: 1}
            _ = fi18
            _ = v19
            _ = v20
            _ = v23
        }
        _ = fi17
    }
    v2 -= genFn0(0, -221)
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 uint64 = 1
        var v26 uint = 18446744073709551615
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            v1 = ((2 > 18446744073709551615) || false)
            if (v3 <= v3) {
                v29 := []int{v5}
                if len(v29) > 0 {
                    v29[0] = (fi27 >> ((0) & 63))
                }
                v30 := GenS3{F0: v5, F1: 255}
                v3 = 63
                v5 -= 32768
                _ = v29
                _ = v30
            } else {
                v5 = fi27
            }
            v31 := GenS4{F0: (-(v2))}
            _ = fi28
            _ = v31
        }
        var v32 int16 = 0
        var v33 int16 = v32
        _ = fi27
        _ = v32
        _ = v33
    }
    var v34 uint64 = 4294967295
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := GenS5{F0: 1, F1: int16(v3)}
        v37 := []int{fi35}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint16 = 32767
    if true {
        for fi39 := 0; fi39 < 12; fi39++ {
            switch fi39 {
            case -2147483648:
                if v1 {
                    var v40 uint8 = 33
                    v41 := []int{v5}
                    _ = v40
                    _ = v41
                } else {
                    if v1 {
                        var v42 int = -763
                        _ = v42
                    } else {
                        v43 := GenS1{F0: -164}
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 9; fi44++ {
                        v45 := []int{fi39}
                        v1 = v1
                        _ = fi44
                        _ = v45
                    }
                }
            case 42:
                v3 = v3
            case 32767:
                v5 = (v5 >> ((575) & 63))
            default:
                var v46 uint32 = v3
                _ = v46
            }
            fmt.Println(fi39, v2)
            _ = fi39
        }
    }
    v34 = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v5
}
