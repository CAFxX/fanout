package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 uint8) int64 {
    var v5 int8 = 0
    v6 := GenS1{F0: p0}
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16) uint8 {
    var v7 int = 805
    v8 := GenS2{F0: (v7 << ((18446744073709551615) & 63))}
    switch int(p0) {
    case 0:
        if false {
            var v9 float32 = ((0.0 + 1e10) + 3.14159)
            var v10 int64 = 289
            _ = v10
            _ = v9
        } else {
            v7 += (v7 * (v7 * 127))
            fmt.Println(v7)
        }
    case -779:
        v11 := []int{(v7 - v7), 882}
        _ = v11
    default:
        fmt.Println(p0, p0)
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        p0 += (p0 & (p0 - p0))
        v7 -= v7
        _ = fi12
    }
    p0 = (p0 & 155)
    _ = p0
    _ = v7
    _ = v8
    return 111
}

func main() {
    var v1 int64 = -796
    v2 := []string{("" + ("" + ("hello" + ("\n\t" + ("\n\t" + "\n\t"))))), "foo bar", "foo bar", "a"}
    var v3 uint32 = 358
    var v4 float64 = (3.14159 - ((1e10 + (0.1 - (-2.25 * (((1e10 / 1.5) - ((((0.0 - 0.0) / 1.0) / 1e10) + (3.14159 * 0.0))) + (0.0 - 3.14159))))) + ((((-2.25 + float64(0.1)) + 3.14159) - float64(0.0)) + 1e10)))
    if false {
        v3 += 2147483648
    } else {
        if true {
            switch 1 {
            case 127:
                fmt.Println(v4, v1, v4)
            case 32767:
                v13 := GenS2{F0: -2147483648}
                _ = v13
            default:
                if (((0.1 * (float64(0.1) / v4)) / 1.0) >= (v4 - (float64(1.5) - (1e10 / v4)))) {
                    v1 += (v1 << ((991) & 63))
                    for fi14 := 0; fi14 < 2; fi14++ {
                        if (float64(float32(v4)) >= v4) {
                            var v15 int32 = -383
                            _ = v15
                        }
                        var v16 bool = (!(255 >= 896))
                        var v17 float32 = 0.0
                        _ = fi14
                        _ = v16
                        _ = v17
                    }
                }
            }
            v18 := GenS1{F0: int32(v3)}
            var v19 uint64 = 18446744073709551615
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 1; fi20++ {
                v3 -= 2147483647
                _ = fi20
            }
            switch 0 {
            case 2147483647:
                var v21 float32 = 3.14159
                _ = v21
            case 4294967295:
                _ = len(v2)
            case -828:
                v1 -= v1
            default:
                if false {
                    v22 := GenS2{F0: 127}
                    _ = v22
                }
            }
        }
        _ = len(v2)
    }
    v2 = append(v2, ("" + ("\n\t" + "foo bar")))
    if true {
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 bool = true
            _ = fi23
            _ = v24
        }
        var v25 uint = 896
        _ = v25
    } else {
        v3 -= (v3 >> ((65535) & 63))
        if (false && false) {
            if ((!(!((1 >= 996) || false))) || (804 < 173)) {
                v1 = int64(v3)
                var v26 uint8 = 1
                _ = v26
            }
            var v27 int64 = (v1 >> ((229) & 63))
            _ = v27
        } else {
            var v28 int32 = 1
            _ = v28
        }
        switch 1 {
        case -9223372036854775808:
            v29 := GenS2{F0: -2147483648}
            _ = v29
        case 916:
            if (100 == -1) {
                var v30 float32 = (-2.25 * 1e10)
                v30 = 0.0
                _ = v30
            } else {
                var v31 uint16 = 2
                _ = v31
            }
        case -408:
            var v32 float64 = float64(((((0.1 + (3.14159 + 0.0)) + 1e10) - (-2.25 - 3.14159)) + 0.1))
            _ = v32
        default:
            v3 += 281
        }
    }
    var v33 float64 = (float64(3.14159) - -2.25)
    var v34 float64 = v4
    v3 = v3
    if true {
        if ((!(3.14159 <= float64(-2.25))) && (("hello" > "a") && (65535 >= 1))) {
            for fi35 := 0; fi35 < 17; fi35++ {
                for fi36 := 0; fi36 < 17; fi36++ {
                    var v37 int = 1
                    var v38 uint64 = 1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = fi35
            }
            switch -982 {
            case 32767:
                _ = len(v2)
            case 2147483647:
                if false {
                    var v39 uint32 = 0
                    _ = v39
                }
            default:
                v40 := GenS3{F0: "hello", F1: "foo bar"}
                _ = v40
            }
        } else {
            v41 := GenS4{F0: -355, F1: -42, F2: v3}
            _ = v41
        }
    } else {
        v1 = v1
        var v42 uint32 = v3
        _ = v42
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        switch (("foo bar" + ("foo bar" + "foo bar")) + "a") {
        case "\n\t":
            if false {
                var v44 float32 = (0.0 * ((((float32(0.1) * ((0.1 * (0.0 + 0.0)) / (float32(v33) / 1.5))) + (0.1 + 0.1)) - (3.14159 + 0.1)) / float32((v33 * v4))))
                _ = v44
            } else {
                v45 := GenS1{F0: int32(v3)}
                if true {
                    _ = len(v2)
                    var v46 uint = 18446744073709551615
                    _ = v46
                }
                _ = v45
            }
        case "hello":
            switch ("" + "foo bar") {
            case "\n\t":
                v3 = (v3 << ((875) & 63))
            default:
                var v47 uint64 = 0
                _ = v47
            }
        default:
            v48 := GenS1{F0: -725}
            _ = v48
        }
        _ = fi43
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (("foo bar" + "foo bar") + "hello") {
    case "hello":
        _ = len(v2)
    }
    v49 := GenS1{F0: 0}
    var v50 uint16 = 84
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v50
}
