package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 string; F2 uint }
func genFn0(p0 uint8) int {
    p0 = p0
    var v5 int8 = 100
    _ = p0
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v6 string = ""
    if (!(255 < 4294967295)) {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint8 = 1
            if (!false) {
                var v9 float32 = 1.5
                var v10 uint64 = 42
                var v11 uint32 = 2147483647
                _ = v10
                _ = v11
                _ = v9
            } else {
                v6 = v6
                fmt.Println(v6, v6)
                if true {
                    v6 = v6
                }
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                delete(v1, (v6 + v6))
                v8 += v8
                _ = fi12
            }
            _ = fi7
            _ = v8
        }
        fmt.Println(v6)
        switch v6 {
        case "":
            switch 0 {
            case 32767:
                if false {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v6 = (("hello" + (v6 + (v6 + ("\n\t" + ((v6 + ((v6 + ((("" + "\n\t") + v6) + "foo bar")) + (v6 + v6))) + v6))))) + v6)
                        _ = fi13
                    }
                } else {
                    var v14 uint32 = 2
                    for fi15 := 0; fi15 < 12; fi15++ {
                        if ((42 < 65535) || true) {
                            for fi16 := 0; fi16 < 18; fi16++ {
                                fmt.Println(v14, fi16)
                                v6 = (v6 + (v6 + v6))
                                v17 := GenS1{F0: v14}
                                _ = fi16
                                _ = v17
                            }
                        }
                        var v18 uint = 205
                        _ = fi15
                        _ = v18
                    }
                    _ = v14
                }
            default:
                var v19 int = -982
                _ = v19
            }
        }
    } else {
        v20 := GenS2{F0: 42, F1: 1}
        switch v6 {
        case "\n\t":
            v6 = v6
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            v6 = ((v6 + v6) + v6)
            if true {
                if true {
                    if true {
                        switch v6 {
                        case "":
                            if false {
                                v6 = v6
                            }
                        }
                        v6 = ((v6 + v6) + v6)
                    }
                } else {
                    v6 = "foo bar"
                    var v22 int32 = 2147483647
                    _ = v22
                }
                var v23 float64 = 0.1
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    }
    var v24 bool = (0 > 18446744073709551615)
    switch genFn0(2) {
    case 32768:
        v25 := GenS3{F0: 42, F1: 556}
        _ = v25
    }
    v26 := []int{42, 0, 4294967295, -116}
    v6 = v6
    fmt.Println(v24)
    v27 := GenS4{F0: 42}
    _ = len(v26)
    if len(v26) > 0 {
        v26[0] = 814
    }
    fmt.Println(v24, v6)
    fmt.Println(v24, v24)
    v28 := GenS4{F0: -1}
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            switch 81 {
            case 32768:
                fmt.Println(v24, v6, v24)
            case 414:
                var v31 int64 = 9223372036854775807
                _ = v31
            default:
                var v32 float64 = 1.5
                _ = v32
            }
            v33 := GenS5{F0: 4294967295}
            var v34 int8 = 42
            _ = fi30
            _ = v33
            _ = v34
        }
        _ = fi29
    }
    v35 := GenS6{F0: (2147483647 == 2147483647)}
    var v36 bool = (1 == 688)
    for fi37 := 0; fi37 < 2; fi37++ {
        if (2147483648 == 620) {
            if len(v26) > 0 {
                v26[0] = (fi37 / ((fi37) | 1))
            }
            fmt.Println(v24, fi37, v6)
            var v38 uint8 = 42
            _ = v38
        } else {
            v6 = ((v6 + (v6 + v6)) + "\n\t")
            var v39 float64 = ((0.1 / 0.1) / float64((((0.0 * 1e10) * 1e10) + 1e10)))
            if ((1 > 34) || false) {
                v36 = (v36 && v36)
                fmt.Println(v39, v39)
                var v40 bool = true
                _ = v40
            } else {
                fmt.Println(v36)
                v36 = true
                v41 := GenS7{F0: v6, F1: ("" + v6), F2: 122}
                _ = v41
            }
            _ = v39
        }
        _ = fi37
    }
    var v42 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v6
}
