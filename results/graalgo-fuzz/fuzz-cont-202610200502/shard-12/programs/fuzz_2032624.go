package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := []uint32{2}
    var v4 float64 = ((1.5 * (1.5 + 3.14159)) - 3.14159)
    v5 := []int{128, 9223372036854775807, 127, 0}
    v6 := []string{"hello", "hello", "foo bar", "foo bar"}
    v6 = append(v6, "foo bar")
    switch -2147483648 {
    case -1:
        v7 := GenS1{F0: 2147483648, F1: "\n\t"}
        _ = v7
    }
    v4 -= (float64(0.1) / v4)
    _ = len(v3)
    var v8 string = "foo bar"
    v9 := GenS1{F0: 330, F1: (v8 + v8)}
    v10 := GenS1{F0: 4294967295, F1: v8}
    v4 = (v4 - v4)
    if true {
        var v11 int32 = -2147483648
        for fi12 := 0; fi12 < 20; fi12++ {
            if (-2.25 <= 3.14159) {
                var v13 int8 = -128
                _ = v13
            } else {
                _ = len(v6)
            }
            v5 = append(v5, (fi12 * (fi12 % ((fi12) | 1))))
            _ = fi12
        }
        _ = v11
    } else {
        switch 9223372036854775807 {
        case -9223372036854775808:
            v8 = ("" + ("hello" + "a"))
        default:
            v4 = 1e10
        }
        v4 = 0.1
    }
    v14 := GenS1{F0: 654, F1: (v8 + v8)}
    var v15 int64 = -1
    v16 := GenS2{F0: v8}
    var v17 uint32 = 2
    switch -2147483648 {
    case -97:
        if (v8 > "\n\t") {
            if (-2147483648 != 42) {
                fmt.Println(v15, v15)
            } else {
                var v18 int16 = 1
                switch 32768 {
                case 674:
                    v19 := v2["\n\t"]
                    _ = v19
                case 32768:
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v15 = 9223372036854775807
                        var v21 uint16 = 32767
                        if false {
                            v15 = v15
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                v3 = append(v3, 4294967295)
                _ = v18
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                if len(v3) > 0 {
                    v3[0] = 42
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v23 := GenS1{F0: (^(v17)), F1: (v8 + "a")}
                _ = fi22
                _ = v23
            }
            switch (v8 + "a") {
            case "a":
                v15 += (v15 >> ((4294967295) & 63))
            case "foo bar":
                var v24 uint64 = 1
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 uint = 110
                _ = fi25
                _ = v26
            }
            if false {
                v2[(v8 + v8)] = false
                var v27 string = v8
                _ = v27
            } else {
                v28 := GenS2{F0: v8}
                _ = v28
            }
            v29 := GenS3{F0: 0}
            _ = v29
        }
    default:
        if false {
            if (-229 > 0) {
                v17 = v17
            }
        } else {
            var v30 int64 = 105
            _ = v30
        }
    }
    fmt.Println(v4, v8)
    if true {
        v31 := v1[v17]
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 7; fi33++ {
                fmt.Println(fi32)
                var v34 int32 = 0
                if (v31 && false) {
                    var v35 float32 = float32(v4)
                    _ = v35
                } else {
                    switch 32768 {
                    case 128:
                        if false {
                            var v36 int8 = 42
                            _ = v36
                        } else {
                            switch -915 {
                            case -9223372036854775808:
                                if (255 < 255) {
                                    var v37 uint64 = 42
                                    v38 := GenS3{F0: 0}
                                    var v39 uint = 4294967295
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    v31 = false
                                    v40 := GenS4{F0: 42}
                                    delete(v1, (^(v17)))
                                    _ = v40
                                }
                            case 42:
                                v41 := GenS1{F0: 4294967295, F1: ("" + v8)}
                                _ = v41
                            default:
                                v8 = "a"
                            }
                        }
                    case -236:
                        var v42 uint = 0
                        _ = v42
                    case 1:
                        v43 := GenS1{F0: (v17 | (v17 | v17)), F1: v8}
                        _ = v43
                    default:
                        if (v31 && (v17 < (v17 | v17))) {
                            v44 := v2[(((("\n\t" + v8) + v8) + (("foo bar" + "foo bar") + "\n\t")) + "hello")]
                            if true {
                                var v45 float64 = float64((float32(float64(1.5)) / float32(1e10)))
                                v46 := GenS4{F0: v15}
                                v1[(v17 & (v17 % ((v17) | 1)))] = false
                                _ = v45
                                _ = v46
                            }
                            v47 := v1[(v17 ^ 42)]
                            _ = v44
                            _ = v47
                        }
                        v34 -= v34
                    }
                    v17 += 0
                }
                v17 -= 507
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
