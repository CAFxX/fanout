package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 219
    v3 := make(map[string]int)
    var v4 string = ("\n\t" + ("hello" + "a"))
    if false {
        if true {
            v5 := GenS1{F0: v2}
            v4 = (v4 + v4)
            _ = v5
        } else {
            fmt.Println(v2, v4)
        }
        var v6 float64 = 1e10
        _ = v6
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            v4 = v4
            _ = fi7
        }
        switch 0 {
        case 763:
            v8 := GenS2{F0: 2}
            _ = v8
        }
    }
    if (!false) {
        fmt.Println(v2)
        for fi9 := 0; fi9 < 7; fi9++ {
            v4 = ("hello" + "a")
            _ = fi9
        }
        var v10 string = v4
        _ = v10
    }
    v11 := []int{42, 4294967295, -2147483648, -63}
    var v12 int8 = 0
    var v13 float64 = 0.0
    var v14 int8 = v12
    if true {
        if true {
            for fi15 := 0; fi15 < 5; fi15++ {
                v4 = ("" + "a")
                var v16 float32 = float32(1.5)
                _ = fi15
                _ = v16
            }
        }
        var v17 float64 = 0.1
        _ = v17
    } else {
        if false {
            var v18 int8 = 0
            if (128 == 9223372036854775807) {
                v13 += (v13 * v13)
                if (true || true) {
                    v4 = "foo bar"
                    if false {
                        v12 -= (v12 - (v14 - v12))
                        fmt.Println(v13, v2)
                        v19 := GenS3{F0: 675}
                        _ = v19
                    } else {
                        var v20 float64 = v13
                        switch v4 {
                        case "hello":
                            _ = len(v11)
                        default:
                            for fi21 := 0; fi21 < 10; fi21++ {
                                v2 += (v2 + (v2 * (^(v2))))
                                var v22 int64 = v2
                                delete(v3, ("" + v4))
                                _ = fi21
                                _ = v22
                            }
                        }
                        _ = v20
                    }
                } else {
                    var v23 float64 = float64((1e10 + 1e10))
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 int64 = v2
                    var v26 uint64 = 42
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v12 -= v14
                if len(v11) > 0 {
                    v11[0] = 9223372036854775807
                }
            }
            _ = v18
        } else {
            var v27 float64 = v13
            var v28 uint32 = 2147483648
            _ = v27
            _ = v28
        }
        fmt.Println(v2, v14)
        v29 := GenS4{F0: 68}
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS5{F0: 65535, F1: 255, F2: -384}
        _ = fi30
        _ = v31
    }
    var v32 string = v4
    if (!(729 != 255)) {
        fmt.Println(v2, v14)
        var v33 uint64 = 18446744073709551615
        if false {
            for fi34 := 0; fi34 < 18; fi34++ {
                if (v12 != (v14 + 42)) {
                    _ = len(v3)
                }
                var v35 uint32 = 0
                _ = fi34
                _ = v35
            }
        }
        _ = v33
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS5{F0: 18446744073709551615, F1: 349, F2: 1}
        var v38 uint8 = 78
        if true {
            fmt.Println(v13, v38, fi36)
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
