package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 bool; F1 uint; F2 int }
func genFn0(p0 int8, p1 uint8) int16 {
    switch 0 {
    case 128:
        p1 = 0
    case 582:
        var v5 float64 = 0.0
        _ = v5
    }
    if true {
        v6 := GenS1{F0: (("" + "foo bar") + "\n\t")}
        _ = v6
    }
    v7 := GenS2{F0: 381, F1: 0}
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 string = "a"
    var v3 int64 = -1
    v4 := make(map[uint32]uint32)
    v8 := []int{1, 9223372036854775807, 383}
    v2 = "a"
    v8 = append(v8, 32767)
    if (!(128 != 237)) {
        var v9 uint8 = 128
        v3 = int64(v9)
        if (7 <= 501) {
            v2 = "\n\t"
        } else {
            var v10 int64 = -198
            for fi11 := 0; fi11 < 9; fi11++ {
                switch (v2 + "a") {
                case "\n\t":
                    v12 := GenS2{F0: 284, F1: -1}
                    _ = v12
                }
                _ = fi11
            }
            _ = v10
        }
        _ = v9
    } else {
        v2 = v2
        if ((2 < 218) || (1e10 >= 0.1)) {
            switch "\n\t" {
            case "\n\t":
                var v13 string = v2
                _ = v13
            case "a":
                v14 := v1[v2]
                _ = v14
            case "hello":
                for fi15 := 0; fi15 < 14; fi15++ {
                    v4[0] = 1
                    delete(v4, 1)
                    _ = fi15
                }
            }
            switch ("hello" + "") {
            case "":
                v4[357] = 922
            case "a":
                fmt.Println(v2, v2)
            default:
                if (v3 == v3) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    if false {
                        var v16 bool = (-2 != 42)
                        _ = v16
                    }
                    if true {
                        v17 := GenS3{F0: "\n\t", F1: (v3 / ((v3) | 1))}
                        v2 = (v2 + "\n\t")
                        v2 = ""
                        _ = v17
                    }
                }
            }
            v18 := GenS1{F0: v2}
            _ = v18
        } else {
            if true {
                _ = len(v8)
            }
        }
    }
    var v19 float64 = 1.5
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            if true {
                var v22 float64 = v19
                _ = v22
            } else {
                if (!((((((0.0 + 1e10) - 1.5) * 0.0) / 1.0) * (0.0 * 0.0)) >= 3.14159)) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    if true {
                        var v23 int = fi21
                        _ = v23
                    } else {
                        var v24 float32 = ((0.0 + 0.1) + 3.14159)
                        v3 = (-(v3))
                        v3 -= -431
                        _ = v24
                    }
                    var v25 int8 = -72
                    var v26 uint64 = 0
                    v27 := GenS2{F0: 130, F1: 9}
                    _ = v25
                    _ = v26
                    _ = v27
                }
                if (!false) {
                    v8 = append(v8, fi20)
                    v19 -= v19
                    var v28 uint16 = 1
                    _ = v28
                } else {
                    fmt.Println(v3, fi21)
                }
            }
            var v29 int64 = v3
            _ = fi21
            _ = v29
        }
        _ = fi20
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            if (!false) {
                _ = len(v1)
                delete(v1, "\n\t")
                switch fi30 {
                case -912:
                    v32 := GenS1{F0: ""}
                    _ = v32
                case 128:
                    var v33 int = (fi31 ^ -412)
                    _ = v33
                case -375:
                    var v34 int64 = -9223372036854775808
                    _ = v34
                }
                fmt.Println(fi30, v19, fi30)
                var v35 float32 = -2.25
                if (v2 == ("a" + (("" + "hello") + v2))) {
                    v35 = 3.14159
                    v36 := GenS4{F0: (18446744073709551615 >= 1), F1: 1, F2: (-(fi30))}
                    var v37 int8 = 127
                    _ = v36
                    _ = v37
                }
                _ = v35
            }
            _ = fi31
        }
        for fi38 := 0; fi38 < 4; fi38++ {
            v39 := GenS4{F0: false, F1: 4294967295, F2: fi30}
            var v40 int64 = v3
            var v41 uint8 = 127
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi30
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v42 := GenS4{F0: (65535 < 42), F1: 0, F2: 4294967295}
    var v43 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v8
}
