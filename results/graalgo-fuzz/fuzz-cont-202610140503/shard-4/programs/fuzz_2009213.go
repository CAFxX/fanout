package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0() int64 {
    var v7 uint32 = 691
    switch -1 {
    case 127:
        switch 127 {
        case 2147483647:
            v8 := []int{}
            _ = v8
        case 127:
            v7 = 42
        case -2147483648:
            v9 := GenS1{F0: true, F1: "hello"}
            _ = v9
        default:
            switch "a" {
            case "foo bar":
                var v10 int64 = 9223372036854775807
                _ = v10
            }
        }
    case 9223372036854775807:
        var v11 bool = false
        _ = v11
    case -2147483648:
        if (354 <= 255) {
            v12 := GenS1{F0: false, F1: "\n\t"}
            if (0.0 > ((3.14159 * 1.5) + (0.0 * (0.0 * 0.1)))) {
                v7 = 747
            }
            v7 = (v7 % (((v7 & v7)) | 1))
            _ = v12
        }
    default:
        v7 += 2147483648
    }
    _ = v7
    return -1
}

func main() {
    v1 := []string{"a"}
    var v2 uint16 = 42
    var v3 int = -2147483648
    var v4 float64 = 0.1
    var v5 int32 = 1
    var v6 uint8 = 0
    fmt.Println(v2, v3)
    var v13 float32 = ((0.0 + 3.14159) + (-2.25 * -2.25))
    v2 += (v2 | v2)
    for fi14 := 0; fi14 < 6; fi14++ {
        switch fi14 {
        case 869:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 0:
            if (-273 >= -736) {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            } else {
                for fi15 := 0; fi15 < 8; fi15++ {
                    fmt.Println(v4, fi15, v5)
                    var v16 int = v3
                    _ = fi15
                    _ = v16
                }
                var v17 uint = 808
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS2{F0: 1}
                    var v20 float64 = v4
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            }
        case 42:
            var v21 uint64 = 0
            _ = v21
        default:
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 bool = false
                if (473 != 2147483648) {
                    if v23 {
                        v5 = v5
                        var v24 int16 = 1
                        v2 -= v2
                        _ = v24
                    }
                    if v23 {
                        if (v23 && (!(!((v5 % ((42) | 1)) > v5)))) {
                            switch (fi14 % (((fi22 >> ((255) & 63))) | 1)) {
                            case 42:
                                switch 4294967295 {
                                case 127:
                                    fmt.Println(v4, v5, fi14)
                                }
                            case 504:
                                var v25 float64 = (v4 * 1.5)
                                _ = v25
                            case -9223372036854775808:
                                var v26 bool = (v23 || v23)
                                v27 := GenS1{F0: false, F1: ("\n\t" + ("foo bar" + ("hello" + "foo bar")))}
                                v28 := GenS2{F0: 4294967295}
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            var v29 int64 = genFn0()
                            _ = v29
                        }
                    }
                } else {
                    v13 -= float32(((v4 / v4) + v4))
                    v30 := GenS2{F0: 2}
                    if v23 {
                        var v31 float64 = 1.5
                        if true {
                            if (v3 >= (v3 << ((189) & 63))) {
                                var v32 int64 = -1
                                _ = v32
                            }
                        }
                        var v33 string = (((("foo bar" + ("foo bar" + "foo bar")) + "") + ("foo bar" + "foo bar")) + "a")
                        _ = v31
                        _ = v33
                    }
                    _ = v30
                }
                var v34 float32 = v13
                v35 := GenS1{F0: v23, F1: "hello"}
                _ = len(v1)
                var v36 int8 = 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi22
                _ = v23
                _ = v34
                _ = v35
                _ = v36
            }
        }
        if true {
            v37 := GenS1{F0: true, F1: ("" + "foo bar")}
            _ = v37
        } else {
            v6 = v6
        }
        _ = fi14
    }
    var v38 uint32 = 4294967295
    fmt.Println(v4)
    var v39 int = v3
    fmt.Println(v2, v38, v4)
    var v40 int32 = 0
    _ = len(v1)
    v4 -= (1e10 / float64(v13))
    v41 := GenS3{F0: -2147483648, F1: (v38 % ((v38) | 1))}
    var v42 uint8 = 203
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
