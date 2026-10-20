package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint32) uint32 {
    if ((^(p0)) <= p0) {
        fmt.Println(p0)
        if (-9223372036854775808 < 42) {
            for fi7 := 0; fi7 < 12; fi7++ {
                switch "a" {
                case "a":
                    fmt.Println(p0, p0)
                case "hello":
                    v8 := GenS1{F0: 127, F1: ("foo bar" + "hello"), F2: 221}
                    _ = v8
                default:
                    var v9 uint32 = 1
                    _ = v9
                }
                _ = fi7
            }
            fmt.Println(p0)
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                p0 += p0
                var v11 uint8 = 128
                v12 := []int{(fi10 << ((817) & 63)), 4294967295, fi10}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    } else {
        if (255 < 182) {
            v13 := GenS1{F0: 255, F1: "hello", F2: 1}
            fmt.Println(p0)
            _ = v13
        }
        switch -2147483648 {
        case 687:
            v14 := []int{-708}
            _ = v14
        default:
            for fi15 := 0; fi15 < 2; fi15++ {
                switch -9223372036854775808 {
                case -846:
                    fmt.Println(p0, fi15)
                default:
                    var v16 int16 = -1
                    _ = v16
                }
                _ = fi15
            }
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int64 = -1
        _ = fi17
        _ = v18
    }
    p0 += (^(p0))
    p0 = (p0 % ((p0) | 1))
    if (("\n\t" + "foo bar") < ("" + "\n\t")) {
        if (0 > 4294967295) {
            v19 := GenS1{F0: 128, F1: ("a" + ("foo bar" + ("a" + "a"))), F2: 0}
            _ = v19
        }
    }
    _ = p0
    return 65535
}

func main() {
    var v1 int = -321
    var v2 int16 = 410
    var v3 uint64 = 0
    var v4 uint16 = 255
    var v5 int8 = 118
    v6 := make(map[uint32]string)
    v20 := GenS1{F0: uint8(v3), F1: "foo bar", F2: -653}
    fmt.Println(v5, v3, v3)
    v1 -= v1
    v2 -= (v2 << ((175) & 63))
    v21 := GenS2{F0: 47, F1: 0, F2: (v2 * v2)}
    fmt.Println(v1, v1)
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 bool = false
        var v24 int16 = (^(v2))
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: 5, F1: 128, F2: 885}
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := v6[920]
        var v28 int32 = -191
        v4 += (v4 | (v4 << ((4294967295) & 63)))
        _ = fi26
        _ = v27
        _ = v28
    }
    if true {
        fmt.Println(v5, v2, v4)
        if (42 <= 0) {
            var v29 float64 = 0.0
            _ = v29
        }
    } else {
        if (!(127 != 128)) {
            switch v1 {
            case -1:
                _ = len(v6)
            case 697:
                if true {
                    if true {
                        v30 := GenS2{F0: 91, F1: 255, F2: int16(v4)}
                        _ = v30
                    } else {
                        v5 -= v5
                    }
                } else {
                    fmt.Println(v5)
                    fmt.Println(v1, v2, v3)
                }
            case 9223372036854775807:
                v31 := v6[104]
                _ = v31
            }
            var v32 int8 = 127
            v33 := GenS3{F0: -1}
            _ = v32
            _ = v33
        } else {
            v5 = (v5 | v5)
            if false {
                v4 -= 1
                v34 := GenS3{F0: 653}
                v35 := GenS3{F0: 9223372036854775807}
                _ = v34
                _ = v35
            }
        }
        if (-2.25 != float64(float32(1.5))) {
            v36 := []int{(v1 / ((v1) | 1)), v1, v1, v1}
            _ = v36
        } else {
            v37 := []int{v1, v1, v1, (v1 * 127)}
            v1 = v1
            _ = v37
        }
        v38 := GenS3{F0: -9223372036854775808}
        _ = v38
    }
    v39 := v6[2147483647]
    var v40 int = v1
    switch "hello" {
    case "\n\t":
        var v41 int8 = (-(v5))
        _ = v41
    default:
        v42 := GenS2{F0: uint8(v3), F1: 128, F2: v2}
        _ = v42
    }
    var v43 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
