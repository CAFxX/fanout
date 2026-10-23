package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int8 }
type GenS5 struct { F0 string; F1 string; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int8; F2 int }
func genFn0(p0 uint8) uint64 {
    v5 := GenS1{F0: 4294967295, F1: 0, F2: 2}
    var v6 uint32 = 790
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 611
    var v3 float32 = ((0.1 + -2.25) - 1.5)
    var v4 uint32 = 926
    switch 0 {
    case 2147483647:
        fmt.Println(v2, v3)
    default:
        _ = len(v1)
    }
    switch ("" + ("foo bar" + "")) {
    case "\n\t":
        for fi7 := 0; fi7 < 6; fi7++ {
            if (true && false) {
                v1[2147483647] = (v4 - v4)
                var v8 uint = 18446744073709551615
                v9 := GenS1{F0: 65535, F1: 124, F2: 970}
                _ = v8
                _ = v9
            }
            if true {
                fmt.Println(fi7, fi7, v4)
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := v1[fi10]
                _ = fi10
                _ = v11
            }
            _ = fi7
        }
    default:
        v2 = v2
    }
    var v12 int64 = 1
    var v13 float32 = float32(-2.25)
    if true {
        var v14 bool = (true || true)
        v15 := GenS1{F0: v4, F1: v2, F2: 255}
        _ = v14
        _ = v15
    } else {
        switch 128 {
        case 0:
            v4 += 2147483648
        case 9223372036854775807:
            fmt.Println(v4, v13, v2)
        case -1:
            var v16 string = ("hello" + ("foo bar" + "foo bar"))
            _ = v16
        }
        if false {
            if (3.14159 > v13) {
                v17 := GenS1{F0: (v4 << ((431) & 63)), F1: (v2 | (-(v2))), F2: 1}
                _ = v17
            } else {
                switch ("" + "hello") {
                case "foo bar":
                    _ = len(v1)
                case "\n\t":
                    var v18 int16 = v2
                    _ = v18
                default:
                    v2 += (v2 % ((v2) | 1))
                }
                if false {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 int = (^(fi20))
                            fmt.Println(v13)
                            v22 := v1[127]
                            v23 := v1[fi19]
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v3 = 3.14159
                        fmt.Println(v4, fi19)
                        _ = fi19
                    }
                    v24 := GenS1{F0: 2147483647, F1: (v2 * 1000), F2: 496}
                    _ = v24
                }
                var v25 uint32 = (v4 & (v4 | (v4 + (^(v4)))))
                _ = v25
            }
            var v26 uint = 255
            _ = v26
        }
    }
    if true {
        if true {
            v12 = (v12 * (^(v12)))
        } else {
            var v27 string = "foo bar"
            v28 := GenS2{F0: v27}
            switch -9223372036854775808 {
            case 127:
                if (v2 <= v2) {
                    v29 := v1[0]
                    delete(v1, 4294967295)
                    var v30 int16 = (v2 | v2)
                    _ = v29
                    _ = v30
                }
            case 9223372036854775807:
                for fi31 := 0; fi31 < 11; fi31++ {
                    if true {
                        for fi32 := 0; fi32 < 4; fi32++ {
                            fmt.Println(v13, v27, v3)
                            fmt.Println(v27, v12)
                            _ = fi32
                        }
                    }
                    _ = fi31
                }
            case -870:
                v13 = ((float32(3.14159) + (v3 - float32(float64(v3)))) / (v13 - (((float32(1.5) + ((3.14159 / v3) - (v3 + v3))) - v3) * 0.1)))
            }
            _ = v27
            _ = v28
        }
        fmt.Println(v2, v3)
        var v33 int = -561
        _ = v33
    }
    var v34 bool = (131 < 128)
    for fi35 := 0; fi35 < 17; fi35++ {
        v1[-9223372036854775808] = 42
        v12 += (v12 | 1)
        _ = fi35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := GenS2{F0: (("foo bar" + "\n\t") + ("hello" + "foo bar"))}
        if (!((1 >= 32) || (!v34))) {
            if true {
                v38 := GenS2{F0: "\n\t"}
                _ = v38
            }
            v39 := GenS2{F0: "hello"}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 16; fi40++ {
                v34 = (v4 >= v4)
                v41 := GenS3{F0: 42}
                v12 = v12
                v2 += v2
                _ = fi40
                _ = v41
            }
            fmt.Println(fi36, v4, v13)
            switch "" {
            case "hello":
                var v42 uint32 = (^(v4))
                _ = v42
            }
        }
        if v34 {
            v43 := GenS4{F0: v12, F1: 4294967295, F2: 127}
            _ = v43
        }
        _ = fi36
        _ = v37
    }
    fmt.Println(v13)
    fmt.Println(v12, v4)
    var v44 float64 = (1.5 / 1e10)
    v44 = v44
    var v45 int64 = (-(v12))
    if (true || (!v34)) {
        delete(v1, 127)
        switch -2147483648 {
        case 1:
            v2 = v2
        case 32767:
            v46 := GenS5{F0: "hello", F1: (("" + "hello") + "a"), F2: 42}
            _ = v46
        case 4294967295:
            v47 := GenS1{F0: v4, F1: v2, F2: 1}
            _ = v47
        }
        v34 = (!(65535 > 0))
    } else {
        var v48 int16 = 643
        _ = v48
    }
    fmt.Println(v12, v45, v12)
    v49 := []int{-472, 127}
    v50 := GenS6{F0: 255, F1: 127, F2: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v49
    _ = v50
}
