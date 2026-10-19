package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
func genFn0(p0 uint64, p1 bool) int8 {
    if (42 <= 42) {
        var v7 int = 9223372036854775807
        switch (v7 - (v7 + v7)) {
        case -2147483648:
            v8 := []int{4294967295, v7, (v7 / (((v7 << ((1) & 63))) | 1))}
            _ = v8
        default:
            var v9 int32 = 816
            _ = v9
        }
        _ = v7
    } else {
        switch -814 {
        case 334:
            if false {
                v10 := []int{-9223372036854775808, 525}
                switch -1 {
                case 32768:
                    p1 = (42 != 42)
                default:
                    for fi11 := 0; fi11 < 3; fi11++ {
                        p1 = p1
                        var v12 uint64 = 852
                        if (p1 || false) {
                            fmt.Println(fi11)
                            switch "hello" {
                            case "":
                                var v13 float32 = 1e10
                                _ = v13
                            case "foo bar":
                                if false {
                                    v12 += v12
                                    p1 = (1 == 1)
                                    if (70 == 0) {
                                        var v14 uint32 = 65535
                                        var v15 string = "\n\t"
                                        for fi16 := 0; fi16 < 4; fi16++ {
                                            var v17 float32 = (0.0 * 0.1)
                                            v17 += 3.14159
                                            _ = fi16
                                            _ = v17
                                        }
                                        _ = v14
                                        _ = v15
                                    } else {
                                        v18 := GenS1{F0: 1}
                                        _ = v18
                                    }
                                } else {
                                    var v19 string = "\n\t"
                                    _ = v19
                                }
                            case "hello":
                                var v20 uint = 18446744073709551615
                                _ = v20
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                }
                _ = v10
            }
        case 128:
            v21 := []int{4294967295, 9223372036854775807, -373}
            _ = v21
        }
        var v22 uint8 = 2
        _ = v22
    }
    var v23 int64 = -207
    p1 = ((p1 || true) && (-297 < 1000))
    var v24 int = 2147483647
    var v25 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return 1
}

func main() {
    v1 := []string{"foo bar", "a"}
    var v2 int16 = -624
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 65535
    var v5 int8 = -9
    var v6 bool = ((-(v2)) != v2)
    var v26 uint16 = 0
    v27 := GenS1{F0: 42}
    var v28 uint32 = (v4 - (v4 / ((1) | 1)))
    delete(v3, (^(v4)))
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if v6 {
        v29 := GenS1{F0: 978}
        if false {
            var v30 int8 = (v5 / ((-1) | 1))
            v6 = ((!(!((!(!(((1.5 - -2.25) >= ((((0.1 + float32(3.14159)) - -2.25) - (0.0 - (0.0 / 1.5))) / (1.5 + 0.1))) && (v6 && (255 <= 0))))) && (1 <= -2147483648)))) || ((v2 >> ((0) & 63)) < (v2 + (v2 & v2))))
            _ = v30
        } else {
            v31 := GenS1{F0: 1}
            if (851 <= 473) {
                fmt.Println(v6, v26, v4)
                delete(v3, v28)
            }
            _ = len(v1)
            _ = v31
        }
        _ = v29
    }
    fmt.Println(v26, v26, v4)
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 uint16 = v26
        v34 := v3[(v4 ^ uint32(v26))]
        var v35 uint = 42
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    if v6 {
        var v36 float64 = -2.25
        v6 = v6
        v37 := GenS1{F0: -9223372036854775808}
        _ = v36
        _ = v37
    } else {
        v26 = v26
        v4 = uint32(v26)
        var v38 bool = (!v6)
        _ = v38
    }
    v39 := GenS2{F0: (v5 - genFn0(42, v6)), F1: "", F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
