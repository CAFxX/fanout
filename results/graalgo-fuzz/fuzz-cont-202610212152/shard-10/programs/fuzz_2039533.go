package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int64 }
func genFn0() int32 {
    var v5 int64 = -1
    v6 := GenS1{F0: 0, F1: 42, F2: -1}
    fmt.Println(v5)
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = 42
        var v9 uint = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 bool) uint16 {
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 bool = (p0 && (42 >= 101))
        v12 := []int{(fi10 % ((fi10) | 1)), fi10}
        _ = fi10
        _ = v11
        _ = v12
    }
    if p0 {
        v13 := []int{4294967295, 127}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float32 = 0.0
            var v16 int = (fi14 | fi14)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        if true {
            if p0 {
                v17 := []int{}
                _ = v17
            }
        } else {
            v18 := GenS2{F0: 42, F1: 42, F2: -9223372036854775808}
            v19 := GenS2{F0: 709, F1: 0, F2: 766}
            _ = v18
            _ = v19
        }
    }
    var v20 int8 = 120
    _ = p0
    _ = v20
    return 304
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 float32 = 0.0
    v4 := []uint32{}
    if (!v2) {
        for fi21 := 0; fi21 < 6; fi21++ {
            if false {
                var v22 uint = 0
                v22 = 4294967295
                var v23 uint = 65535
                _ = v22
                _ = v23
            }
            _ = fi21
        }
        var v24 float64 = ((0.1 * (0.1 - 1.5)) / -2.25)
        _ = v24
    } else {
        fmt.Println(v3, v3)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v2, v3)
    v3 = v3
    if true {
        v3 += (((3.14159 / ((((v3 * (v3 + v3)) / v3) - (((1e10 + v3) + (v3 + v3)) - 0.1)) * ((v3 * 3.14159) * v3))) * v3) * float32((float64(v3) + (-2.25 / 0.1))))
    }
    v2 = true
    var v25 uint32 = 1
    _ = len(v4)
    for fi26 := 0; fi26 < 12; fi26++ {
        if true {
            fmt.Println(fi26, v25)
        } else {
            v2 = false
        }
        if v2 {
            if v2 {
                for fi27 := 0; fi27 < 4; fi27++ {
                    var v28 int32 = -1
                    var v29 uint64 = 2
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                switch "a" {
                case "":
                    v3 -= v3
                case "a":
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v1["\n\t"] = fi30
                        v31 := GenS2{F0: (v25 ^ (v25 >> ((18446744073709551615) & 63))), F1: 127, F2: 502}
                        _ = fi30
                        _ = v31
                    }
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    v2 = (-9223372036854775808 == -241)
                    var v33 float32 = float32((3.14159 / 3.14159))
                    _ = fi32
                    _ = v33
                }
            }
            v34 := GenS3{F0: 2}
            v25 = (^(v25))
            _ = v34
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 uint32 = 42
            _ = fi35
            _ = v36
        }
        _ = fi26
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 uint64 = 2
            if v2 {
                var v40 bool = (3.14159 != (((1e10 / 3.14159) * 1e10) * ((0.1 - 0.0) * 0.1)))
                v41 := GenS4{F0: (-2147483648 <= 42), F1: 9223372036854775807}
                v39 += 1
                _ = v40
                _ = v41
            }
            var v42 float64 = ((-2.25 / (1.5 * (0.1 / (0.1 - ((0.0 - (3.14159 / 1.5)) / ((-2.25 - 0.0) / (0.0 / ((-2.25 - 1.5) / float64(v3))))))))) / 1.5)
            _ = fi38
            _ = v39
            _ = v42
        }
        for fi43 := 0; fi43 < 17; fi43++ {
            fmt.Println(fi43, v3, v3)
            var v44 uint16 = 1
            _ = fi43
            _ = v44
        }
        for fi45 := 0; fi45 < 6; fi45++ {
            switch "a" {
            case "":
                v46 := v1[("a" + ("a" + ("a" + "")))]
                _ = v46
            default:
                if v2 {
                    delete(v1, "hello")
                    v1["a"] = 128
                    v47 := GenS1{F0: 2, F1: 128, F2: -587}
                    _ = v47
                }
            }
            _ = fi45
        }
        _ = fi37
    }
    switch "hello" {
    case "hello":
        v25 = (v25 | v25)
    case "":
        switch -1 {
        case 0:
            for fi48 := 0; fi48 < 4; fi48++ {
                v49 := GenS5{F0: 518, F1: 2147483647}
                _ = fi48
                _ = v49
            }
        case 812:
            for fi50 := 0; fi50 < 11; fi50++ {
                v51 := GenS6{F0: 0}
                var v52 int32 = -1
                _ = fi50
                _ = v51
                _ = v52
            }
        case 9223372036854775807:
            v53 := v1[((("hello" + ("hello" + "a")) + "hello") + "a")]
            _ = v53
        }
    }
    fmt.Println(v25)
    _ = len(v4)
    if v2 {
        if true {
            for fi54 := 0; fi54 < 19; fi54++ {
                if len(v4) > 0 {
                    v4[0] = (^(v25))
                }
                v2 = v2
                _ = fi54
            }
            v3 = (v3 - float32(0.0))
            switch 127 {
            case 4294967295:
                for fi55 := 0; fi55 < 1; fi55++ {
                    fmt.Println(v3, v2)
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                    v2 = (((v2 || true) || (!(genFn1((((-395 == 9223372036854775807) && true) || v2)) < 1))) || ((364 == 874) && ((fi55 << ((787) & 63)) == (fi55 ^ -220))))
                    _ = fi55
                }
            case 1:
                _ = len(v1)
            case 32768:
                var v56 uint32 = v25
                _ = v56
            default:
                var v57 uint16 = 32767
                _ = v57
            }
        }
    }
    v3 += v3
    v3 = (v3 + -2.25)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
