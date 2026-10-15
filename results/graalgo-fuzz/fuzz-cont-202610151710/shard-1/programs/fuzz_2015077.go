package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint32; F1 uint16 }
func genFn0() uint {
    v6 := []int{}
    v7 := GenS1{F0: ((65535 != 65535) && (0.0 != 3.14159))}
    _ = v6
    _ = v7
    return 836
}
func genFn1() uint8 {
    var v8 float64 = 3.14159
    v9 := []int{}
    if (609 <= 127) {
        switch -1 {
        case 197:
            v10 := GenS1{F0: true}
            _ = v10
        }
    }
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 42
    var v3 bool = true
    v4 := []uint32{2147483648}
    var v5 float64 = -2.25
    if (v3 && ((v1 + "\n\t") > v1)) {
        var v11 float64 = float64((0.1 + (1e10 + ((3.14159 + 1e10) / 1e10))))
        if len(v4) > 0 {
            v4[0] = 0
        }
        if v3 {
            v12 := GenS1{F0: (!(!((^(v2)) != (v2 >> ((18446744073709551615) & 63)))))}
            var v13 uint32 = 0
            if v3 {
                switch 1 {
                case 0:
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 uint = 773
                        fmt.Println(v11, v13, v3)
                        _ = fi14
                        _ = v15
                    }
                case 128:
                    var v16 int8 = -1
                    _ = v16
                default:
                    switch "\n\t" {
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                    default:
                        var v17 bool = (741 != -1)
                        _ = v17
                    }
                }
            } else {
                var v18 float64 = 0.1
                var v19 int32 = 0
                v13 = 324
                _ = v18
                _ = v19
            }
            _ = v12
            _ = v13
        } else {
            v20 := GenS1{F0: v3}
            v3 = v3
            v1 = ""
            _ = v20
        }
        _ = v11
    }
    v2 = v2
    v1 = v1
    v21 := GenS2{F0: 1}
    v4 = append(v4, 4294967295)
    var v22 bool = v3
    v23 := GenS2{F0: 1}
    fmt.Println(v1)
    for fi24 := 0; fi24 < 13; fi24++ {
        v2 = genFn1()
        v5 = v5
        _ = fi24
    }
    v5 += (float64(1e10) / v5)
    v1 = v1
    if v22 {
        v4 = append(v4, 4294967295)
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 uint = 941
            v27 := GenS1{F0: (float64(0.1) != (float64(3.14159) + v5))}
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        var v28 float32 = 3.14159
        switch v1 {
        case "":
            if v3 {
                fmt.Println(v1, v28)
                fmt.Println(v1, v2, v22)
            } else {
                var v29 int64 = 0
                fmt.Println(v28)
                _ = v29
            }
        case "foo bar":
            for fi30 := 0; fi30 < 6; fi30++ {
                v22 = true
                var v31 uint32 = 2147483648
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 uint16 = 0
                    v2 -= v2
                    v3 = (v2 < (v2 >> ((1) & 63)))
                    _ = fi32
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = 0
            }
        default:
            v4 = append(v4, 1)
        }
        v34 := GenS3{F0: 0, F1: v1}
        _ = v28
        _ = v34
    }
    v35 := GenS4{F0: 0, F1: 255}
    var v36 float32 = 0.0
    v1 = ""
    var v37 int32 = -1
    for fi38 := 0; fi38 < 18; fi38++ {
        if (42 >= 2) {
            _ = len(v4)
            var v39 int64 = 9223372036854775807
            _ = v39
        }
        fmt.Println(v22, fi38)
        var v40 int32 = v37
        _ = fi38
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
