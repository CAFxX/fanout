package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 uint32) uint16 {
    if ((825 < 813) && false) {
        p1 += 2
        v6 := GenS1{F0: 152, F1: (p0 + p0)}
        v7 := []int{128}
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 127, F1: p0}
        switch "foo bar" {
        case "\n\t":
            p1 += 65535
        case "foo bar":
            var v9 int16 = -32768
            _ = v9
        }
        _ = v8
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int8 = (p0 << ((973) & 63))
        var v12 int = 948
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    var v4 bool = true
    var v5 int64 = 42
    for fi13 := 0; fi13 < 4; fi13++ {
        if ((v3 << ((255) & 63)) <= (v3 * (v3 | (v3 + v3)))) {
            var v14 string = (("\n\t" + "foo bar") + "")
            v14 = ("" + (v14 + (("\n\t" + (((v14 + (v14 + v14)) + ("foo bar" + (v14 + v14))) + "\n\t")) + (v14 + ((v14 + v14) + v14)))))
            _ = v14
        } else {
            v4 = v4
            var v15 uint8 = v1
            var v16 int = fi13
            _ = v15
            _ = v16
        }
        if v4 {
            v5 += -9223372036854775808
        } else {
            switch fi13 {
            case 128:
                v17 := v2[0]
                _ = v17
            case -9223372036854775808:
                var v18 int32 = -1
                _ = v18
            case 419:
                switch "foo bar" {
                case "\n\t":
                    v19 := []int{fi13}
                    _ = v19
                case "foo bar":
                    var v20 int64 = 1
                    _ = v20
                case "hello":
                    v3 += v3
                default:
                    if (false || true) {
                        var v21 float64 = 3.14159
                        _ = v21
                    }
                }
            }
        }
        if v4 {
            switch fi13 {
            case 0:
                if v4 {
                    v5 += v5
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 int8 = 100
                        _ = fi22
                        _ = v23
                    }
                    if (524 == 77) {
                        var v24 uint32 = 816
                        var v25 int16 = (-(v3))
                        v26 := GenS2{F0: 127}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                } else {
                    v27 := GenS3{F0: 2}
                    _ = v27
                }
            case 128:
                v5 -= v5
            }
            var v28 uint16 = 65535
            if false {
                fmt.Println(v5, v28, v4)
            }
            _ = v28
        } else {
            for fi29 := 0; fi29 < 14; fi29++ {
                v5 = -937
                fmt.Println(fi13, v1)
                _ = fi29
            }
        }
        _ = fi13
    }
    fmt.Println(v1, v3, v5)
    switch 0 {
    case 128:
        v3 += (v3 * 374)
    case 4294967295:
        if ((!v4) && (!((v5 / (((v5 | (v5 & v5))) | 1)) >= v5))) {
            fmt.Println(v3)
            v30 := []int{}
            _ = v30
        } else {
            fmt.Println(v1)
        }
    case -896:
        v31 := []int{761, 0, 9223372036854775807, 4294967295}
        _ = v31
    }
    v32 := v2[0]
    fmt.Println(v4, v5)
    v33 := GenS4{F0: (("\n\t" + v32) + ((v32 + v32) + "\n\t"))}
    var v34 int = 306
    var v35 int8 = 42
    v2[2147483647] = v32
    switch v32 {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        v1 += (v1 >> ((65535) & 63))
    case "\n\t":
        v36 := GenS2{F0: 75}
        _ = v36
    default:
        if false {
            if ((!(1 >= 9223372036854775807)) || v4) {
                v37 := []int{v34}
                if v4 {
                    var v38 uint = 51
                    _ = v38
                } else {
                    v34 = (v34 / (((v34 | (v34 * int(v3)))) | 1))
                }
                _ = v37
            }
        }
    }
    v1 += 0
    if ((v34 ^ (-(v34))) == (v34 >> ((4294967295) & 63))) {
        v39 := []int{v34, (^(v34)), int(v1), (v34 - 182)}
        _ = v39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        if (757 > 18446744073709551615) {
            switch v32 {
            case "foo bar":
                v41 := []int{fi40, (^(v34))}
                _ = v41
            default:
                v1 += v1
            }
            v42 := GenS1{F0: (v1 ^ v1), F1: (v35 * (v35 + v35))}
            var v43 uint64 = 2
            _ = v42
            _ = v43
        } else {
            v1 = v1
        }
        _ = fi40
    }
    v34 = (v34 << ((635) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
