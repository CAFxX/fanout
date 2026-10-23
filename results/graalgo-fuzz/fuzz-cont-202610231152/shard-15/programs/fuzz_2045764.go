package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS6 struct { F0 string; F1 int; F2 int64 }
type GenS7 struct { F0 int32; F1 uint32 }
func genFn0(p0 string, p1 int16) int8 {
    var v4 uint16 = 244
    v4 = 65535
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 uint, p1 uint32) uint32 {
    v5 := GenS1{F0: -32768, F1: 2}
    v6 := []int{2147483647}
    p0 += (^(p0))
    p0 = 237
    var v7 int8 = -1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := []int{-110}
    var v2 int16 = -1
    v3 := []string{"hello"}
    var v8 bool = (100 <= 127)
    if len(v1) > 0 {
        v1[0] = 223
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if v8 {
            var v10 int8 = 78
            var v11 uint16 = 255
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: 42, F1: 203}
                if true {
                    v14 := GenS1{F0: 1, F1: uint8(v11)}
                    var v15 float64 = float64(-2.25)
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: 32767, F1: 195}
                    if (!false) {
                        v17 := GenS1{F0: v2, F1: 224}
                        v18 := GenS1{F0: (v2 % ((v2) | 1)), F1: 68}
                        _ = v17
                        _ = v18
                    } else {
                        switch fi12 {
                        case 0:
                            v19 := GenS2{F0: "\n\t"}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v20 int16 = (v2 & (v2 | 32767))
                            _ = v19
                            _ = v20
                        case 82:
                            _ = len(v1)
                        }
                    }
                    _ = v16
                }
                if v8 {
                    var v21 uint64 = 2
                    var v22 uint64 = v21
                    _ = v21
                    _ = v22
                }
                _ = fi12
                _ = v13
            }
            _ = v10
            _ = v11
        } else {
            v2 += (v2 % (((v2 - v2)) | 1))
        }
        var v23 int64 = -9223372036854775808
        if ((^(v2)) > 0) {
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := GenS3{F0: 65535, F1: 2, F2: 1}
                if false {
                    if ((v23 * 9223372036854775807) >= v23) {
                        v3 = append(v3, (("" + "") + (("a" + "a") + "a")))
                        v2 += v2
                        fmt.Println(fi24, fi24)
                    } else {
                        var v26 string = "a"
                        var v27 int8 = -101
                        _ = v26
                        _ = v27
                    }
                } else {
                    v1 = append(v1, fi24)
                    switch "" {
                    case "\n\t":
                        var v28 int32 = 0
                        _ = v28
                    default:
                        var v29 int8 = 100
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        switch "\n\t" {
                        case "hello":
                            v30 := GenS4{F0: 32767, F1: 201}
                            _ = v30
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v29
                    }
                }
                fmt.Println(v8)
                _ = fi24
                _ = v25
            }
        }
        _ = fi9
        _ = v23
    }
    v31 := GenS5{F0: 1, F1: 2147483647, F2: 298}
    if v8 {
        var v32 float64 = float64(1e10)
        var v33 int64 = 67
        var v34 int64 = (v33 + v33)
        _ = v32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: -243, F1: 255}
    var v36 uint = 1
    if (9223372036854775807 > 831) {
        for fi37 := 0; fi37 < 1; fi37++ {
            v8 = true
            _ = fi37
        }
        fmt.Println(v2)
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        if v8 {
            if len(v1) > 0 {
                v1[0] = (-(fi38))
            }
            v39 := GenS6{F0: "foo bar", F1: (^(fi38)), F2: 42}
            var v40 float64 = 0.1
            _ = v39
            _ = v40
        }
        fmt.Println(v8, fi38, v8)
        _ = fi38
    }
    var v41 float32 = (0.0 - 1.5)
    var v42 string = "a"
    var v43 uint = (^(v36))
    for fi44 := 0; fi44 < 20; fi44++ {
        if (-966 != 1) {
            var v45 int64 = -9223372036854775808
            v46 := GenS7{F0: 1, F1: 258}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v45
            _ = v46
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch (fi44 << ((475) & 63)) {
        case 4294967295:
            var v47 string = (v42 + (((v42 + "") + v42) + ""))
            _ = v47
        case -1:
            v48 := GenS2{F0: v42}
            _ = v48
        }
        _ = fi44
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        switch fi49 {
        case 695:
            switch (fi49 - int(v2)) {
            case 523:
                fmt.Println(fi49, v36)
            case 849:
                for fi50 := 0; fi50 < 17; fi50++ {
                    var v51 int = (fi49 | (fi49 + fi50))
                    v41 += 1e10
                    v2 = 1000
                    fmt.Println(v43)
                    _ = fi50
                    _ = v51
                }
            case 42:
                var v52 string = "hello"
                _ = v52
            }
        }
        var v53 string = v42
        _ = fi49
        _ = v53
    }
    switch (v42 + (v42 + v42)) {
    case "":
        v41 += float32((0.1 - 1e10))
    case "\n\t":
        if (942 > 784) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        fmt.Println(v8, v42)
    }
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v41
    _ = v42
    _ = v43
    _ = v8
}
