package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 bool }
func genFn0() string {
    var v7 uint8 = 228
    var v8 int64 = -736
    var v9 int64 = (-(v8))
    _ = v7
    _ = v8
    _ = v9
    return ("a" + (("foo bar" + ("foo bar" + "hello")) + ""))
}

func main() {
    var v1 float64 = 1.5
    var v2 int = 4294967295
    var v3 int32 = 601
    v4 := make(map[int]bool)
    var v5 string = ("foo bar" + ("a" + "hello"))
    var v6 bool = ((4294967295 > 4294967295) || true)
    fmt.Println(v5)
    if ((2 >= 845) || true) {
        v10 := GenS1{F0: 44}
        _ = v10
    } else {
        v11 := []int{(v2 + v2), 0, v2, v2}
        switch (v2 / (((v2 * 1)) | 1)) {
        case 32768:
            v3 -= (v3 & (v3 - -2147483648))
        case 9223372036854775807:
            fmt.Println(v2, v6)
        }
        _ = v11
    }
    var v12 uint16 = 32767
    var v13 float64 = 1e10
    var v14 int32 = -2147483648
    v15 := GenS1{F0: 0}
    if (!((3.14159 + 1.5) > -2.25)) {
        if v6 {
            fmt.Println(v3, v12)
        } else {
            var v16 float32 = float32(3.14159)
            var v17 bool = v6
            _ = v16
            _ = v17
        }
        v2 = (v2 | v2)
    } else {
        switch v5 {
        case "":
            v5 = v5
        case "hello":
            if true {
                fmt.Println(v1, v6, v3)
                v18 := v4[(v2 / (((v2 >> ((0) & 63))) | 1))]
                _ = v18
            }
        }
        var v19 float32 = 3.14159
        _ = v19
    }
    if v6 {
        var v20 int = v2
        _ = len(v4)
        var v21 int64 = 1
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 float64 = float64(1e10)
            fmt.Println(v2)
            v6 = v6
            _ = fi22
            _ = v23
        }
        if (uint32(v12) >= 1) {
            switch (v5 + v5) {
            case "hello":
                v13 = ((v13 - 3.14159) - v13)
            case "":
                if true {
                    v2 += (^(v2))
                    var v24 int8 = 0
                    _ = v24
                } else {
                    if (0 < 60) {
                        var v25 int16 = 1000
                        if true {
                            fmt.Println(v13, v5, v12)
                            for fi26 := 0; fi26 < 5; fi26++ {
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    v28 := GenS2{F0: 18446744073709551615, F1: fi26, F2: true}
                                    var v29 int16 = (v25 | v25)
                                    if v6 {
                                        delete(v4, v2)
                                    } else {
                                        v30 := GenS1{F0: 217}
                                        var v31 int64 = 9223372036854775807
                                        v32 := GenS2{F0: 65535, F1: 9223372036854775807, F2: v6}
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                v1 += -2.25
                                v33 := v4[v2]
                                _ = fi26
                                _ = v33
                            }
                        }
                        _ = v25
                    }
                }
            default:
                var v34 float32 = float32(v13)
                _ = v34
            }
        }
    }
    v35 := []int{0, v2}
    fmt.Println(v14, v1)
    delete(v4, (v2 / ((v2) | 1)))
    v5 = (v5 + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
