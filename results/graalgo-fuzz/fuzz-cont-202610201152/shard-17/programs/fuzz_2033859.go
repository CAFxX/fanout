package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 int) int16 {
    var v6 uint8 = 127
    if (true || true) {
        v7 := []int{p0}
        _ = len(v7)
        fmt.Println(v6, p0)
        _ = v7
    }
    _ = p0
    _ = v6
    return -32768
}

func main() {
    var v1 uint64 = 703
    var v2 float64 = (-2.25 + ((0.1 / (1.5 / (0.1 * 3.14159))) / (((-2.25 * (3.14159 - 3.14159)) + (-2.25 + 0.0)) / 1e10)))
    var v3 bool = false
    var v4 int8 = 51
    var v5 uint16 = 255
    fmt.Println(v4, v5, v2)
    var v8 float32 = (-2.25 / (0.1 / -2.25))
    var v9 int = -1
    var v10 uint16 = 65535
    v11 := []int{(v9 >> ((0) & 63)), v9}
    fmt.Println(v8)
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: 651, F1: 2, F2: 0}
        if (((v8 / 1.0) + 1.5) <= v8) {
            fmt.Println(v5)
        } else {
            v2 = ((v2 + v2) * (v2 * (v2 + 1e10)))
            var v14 int32 = -145
            _ = v14
        }
        switch "\n\t" {
        case "a":
            v11 = append(v11, v9)
        case "foo bar":
            v10 = 1
        default:
            var v15 string = ("hello" + "foo bar")
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    var v16 uint16 = (^(v5))
    v5 = v10
    v17 := GenS2{F0: (-(v4)), F1: 1}
    v10 += 65535
    fmt.Println(v4, v9)
    if false {
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint16 = v5
                if true {
                    v20 += v10
                    if (1 >= 1) {
                        if (2 < 65535) {
                            fmt.Println(v3)
                        }
                        var v21 int64 = 1
                        var v22 float32 = v8
                        var v23 float32 = ((float32(3.14159) + float32(v2)) / v8)
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v1 -= v1
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if ((621 == 130) && (!true)) {
                        if true {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            v16 = v10
                        }
                    }
                    v10 = 1
                }
                _ = fi19
                _ = v20
            }
            switch (("" + ("" + "\n\t")) + ("" + "\n\t")) {
            case "foo bar":
                if v3 {
                    switch "" {
                    case "\n\t":
                        v1 += (v1 | 18446744073709551615)
                    case "":
                        var v24 int = -439
                        _ = v24
                    case "foo bar":
                        var v25 uint16 = v5
                        _ = v25
                    }
                    fmt.Println(v1, fi18, v5)
                    var v26 int16 = 42
                    _ = v26
                }
            case "a":
                v27 := GenS1{F0: 1, F1: 127, F2: 1}
                _ = v27
            }
            _ = fi18
        }
    }
    fmt.Println(v16, v9)
    if (!(-2.25 >= v2)) {
        var v28 int8 = int8(v1)
        _ = v28
    } else {
        v9 += 9223372036854775807
        for fi29 := 0; fi29 < 5; fi29++ {
            if (v3 || (v3 && (v2 == v2))) {
                v5 -= (v10 * 255)
            }
            _ = fi29
        }
        fmt.Println(v1, v5)
    }
    v30 := GenS1{F0: 892, F1: 127, F2: 0}
    v10 += (v10 + v5)
    var v31 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
