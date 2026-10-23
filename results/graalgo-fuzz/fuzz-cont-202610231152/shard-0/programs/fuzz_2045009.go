package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32, p1 uint8) string {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 += (^(p0))
        _ = fi5
    }
    fmt.Println(p1, p0)
    v6 := []int{181, -165, 2147483647}
    if false {
        if ((!(p1 < p1)) || true) {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 float32 = 1.5
                _ = fi7
                _ = v8
            }
            var v9 float64 = 3.14159
            switch 128 {
            case 128:
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 bool = (int16(p1) < 1)
                    _ = fi10
                    _ = v11
                }
            case -2147483648:
                for fi12 := 0; fi12 < 20; fi12++ {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v9 += ((v9 / v9) * v9)
                        v14 := GenS1{F0: (false || (!false))}
                        var v15 float64 = (v9 * ((1e10 + v9) - 3.14159))
                        switch (-(fi12)) {
                        case -9223372036854775808:
                            switch (fi13 | fi12) {
                            case -1:
                                if (668 >= 255) {
                                    p1 = p1
                                    v16 := GenS2{F0: 2147483647, F1: -1, F2: -9223372036854775808}
                                    _ = v16
                                }
                            case 32768:
                                var v17 int8 = int8(p1)
                                _ = v17
                            default:
                                v9 -= (v9 / -2.25)
                            }
                        default:
                            switch fi13 {
                            case 32768:
                                var v18 float32 = float32(v15)
                                _ = v18
                            }
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = fi12
                }
            default:
                var v19 uint8 = (p1 - 127)
                _ = v19
            }
            _ = v9
        }
        if ((p0 % ((1) | 1)) > 159) {
            fmt.Println(p1, p0)
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 float64 = (float64(float32(float64(0.0))) / (0.1 * (1.5 / 1.0)))
                v6 = append(v6, fi20)
                _ = fi20
                _ = v21
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}
func genFn1(p0 uint8) string {
    if (!(42 == 18446744073709551615)) {
        p0 = 87
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint64 = 18446744073709551615
            _ = fi23
            _ = v24
        }
        if true {
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS2{F0: 42, F1: 1, F2: -1}
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                p0 = p0
                p0 = (p0 >> ((0) & 63))
                switch ("\n\t" + "foo bar") {
                case "a":
                    p0 = 42
                case "hello":
                    var v28 float64 = 3.14159
                    _ = v28
                case "\n\t":
                    p0 = 2
                default:
                    if (false || true) {
                        p0 -= 220
                        v29 := []int{42, (fi22 - fi27)}
                        _ = v29
                    } else {
                        p0 += 197
                        if ((-128 <= 0) && false) {
                            p0 -= p0
                        }
                    }
                }
                if true {
                    var v30 float64 = 0.0
                    _ = v30
                }
                _ = fi27
            }
        } else {
            v31 := []int{127, fi22, fi22}
            _ = v31
        }
        var v32 float32 = -2.25
        _ = fi22
        _ = v32
    }
    var v33 float64 = ((1e10 * (3.14159 / 3.14159)) - 0.1)
    _ = p0
    _ = v33
    return "hello"
}

func main() {
    var v1 int32 = 1
    var v2 float32 = 0.0
    v3 := make(map[uint32]string)
    var v4 uint32 = 43
    v34 := GenS3{F0: 42, F1: 1}
    var v35 string = ("\n\t" + "hello")
    if true {
        fmt.Println(v35, v35, v2)
        switch -627 {
        case 49:
            v1 += 42
        default:
            v35 = v35
        }
    } else {
        v1 = v1
        var v36 uint = 690
        _ = v36
    }
    var v37 float32 = v2
    var v38 int32 = (v1 * (v1 & v1))
    v38 += v38
    for fi39 := 0; fi39 < 14; fi39++ {
        switch fi39 {
        case 2147483647:
            if (18446744073709551615 > 2) {
                v2 += v2
                v40 := v3[(v4 * (v4 - v4))]
                _ = v40
            }
        case 127:
            v2 = v37
        case -1:
            var v41 int32 = v38
            _ = v41
        }
        v42 := GenS2{F0: v38, F1: (v1 + v38), F2: 1}
        _ = fi39
        _ = v42
    }
    v37 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
}
