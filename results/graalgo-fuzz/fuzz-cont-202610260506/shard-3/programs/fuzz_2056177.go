package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
func genFn0(p0 int8) int32 {
    switch 1 {
    case 968:
        p0 -= -105
    case -510:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    return 2147483647
}
func genFn1() uint8 {
    var v6 bool = true
    v7 := GenS1{F0: 65535}
    fmt.Println(v6)
    v6 = v6
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int64 = -9223372036854775808
        var v10 string = ("" + "a")
        switch ("foo bar" + v10) {
        case "\n\t":
            var v11 uint = 18446744073709551615
            _ = v11
        default:
            if (!v6) {
                if v6 {
                    var v12 float64 = 0.0
                    _ = v12
                } else {
                    fmt.Println(fi8, v9)
                    switch ("foo bar" + ((v10 + "a") + ("" + v10))) {
                    case "":
                        var v13 uint64 = 4294967295
                        _ = v13
                    case "a":
                        for fi14 := 0; fi14 < 15; fi14++ {
                            v6 = v6
                            fmt.Println(v10, fi14, v10)
                            _ = fi14
                        }
                    }
                    v9 -= 42
                }
                v10 = v10
                fmt.Println(fi8)
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := []string{"", "hello"}
    var v5 bool = false
    switch 2147483647 {
    case 4294967295:
        delete(v3, 128)
    case -874:
        v5 = (-2.25 >= -2.25)
    case 1:
        var v15 int32 = 0
        _ = v15
    }
    fmt.Println(v5)
    v5 = ((-1 > 2147483647) && false)
    var v16 int32 = 0
    var v17 float64 = 0.1
    var v18 int32 = 0
    var v19 bool = (!(-9223372036854775808 < -58))
    var v20 bool = v5
    var v21 int8 = 70
    v20 = v20
    if (-1 > 2147483647) {
        if false {
            v5 = true
            if v5 {
                v16 -= genFn0(-3)
                fmt.Println(v19)
                var v22 int = 0
                _ = v22
            }
            fmt.Println(v20, v19)
        } else {
            v23 := GenS2{F0: 9223372036854775807, F1: 182, F2: (^(v21))}
            if true {
                for fi24 := 0; fi24 < 20; fi24++ {
                    v16 = (v16 - (^(v18)))
                    v25 := GenS3{F0: -202, F1: 42}
                    switch (fi24 | (fi24 - (fi24 + fi24))) {
                    case 127:
                        if ((v19 || v5) && v5) {
                            var v26 float32 = 0.0
                            v27 := v1[fi24]
                            _ = v26
                            _ = v27
                        }
                    default:
                        var v28 uint8 = 2
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                v4 = append(v4, "hello")
            }
            _ = v23
        }
    }
    v19 = true
    for fi29 := 0; fi29 < 16; fi29++ {
        switch (fi29 >> ((1) & 63)) {
        case -2147483648:
            for fi30 := 0; fi30 < 10; fi30++ {
                v31 := GenS4{F0: 942}
                _ = fi30
                _ = v31
            }
        case 32767:
            if (v16 == -2147483648) {
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 int8 = v21
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v16, v17, v20)
                v34 := GenS5{F0: 934}
                _ = v34
            } else {
                v21 = v21
                fmt.Println(v20, v21)
            }
        default:
            v35 := GenS6{F0: false}
            _ = v35
        }
        _ = fi29
    }
    var v36 uint = 686
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
