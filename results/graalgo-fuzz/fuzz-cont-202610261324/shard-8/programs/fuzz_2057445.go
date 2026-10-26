package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0(p0 int, p1 uint) int64 {
    v5 := GenS1{F0: p1, F1: (p1 - (p1 - p1)), F2: 42}
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(fi6, fi6, fi7)
            var v8 uint = 993
            _ = fi7
            _ = v8
        }
        var v9 float64 = float64((3.14159 - 1.5))
        _ = fi6
        _ = v9
    }
    var v10 float64 = (((0.1 - 3.14159) - (0.1 / 0.1)) / -2.25)
    switch p0 {
    case 543:
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := []int{-441, p0, (p0 - (^(p0))), (p0 & fi11)}
            _ = fi11
            _ = v12
        }
    case -353:
        v13 := []int{(p0 % (((p0 + (p0 / (((p0 | p0)) | 1)))) | 1)), 4294967295, p0}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 2
    v2 := []string{}
    var v3 int32 = 224
    var v4 uint16 = (v1 / ((0) | 1))
    v14 := GenS1{F0: 597, F1: 234, F2: 0}
    v3 = -2147483648
    var v15 int32 = 256
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 = (-(v3))
        var v17 float32 = (0.1 * 3.14159)
        _ = fi16
        _ = v17
    }
    v15 = v15
    if (float32(1.5) > 0.0) {
        for fi18 := 0; fi18 < 19; fi18++ {
            if (!false) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi19 := 0; fi19 < 10; fi19++ {
                    v2 = append(v2, "")
                    v20 := GenS1{F0: 404, F1: 0, F2: 1}
                    v4 = 32767
                    _ = fi19
                    _ = v20
                }
            } else {
                fmt.Println(v4, v15)
                var v21 float64 = 0.1
                _ = v21
            }
            _ = fi18
        }
        v15 += v3
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 uint32 = 65535
            v23 -= (v23 + (v23 >> ((1) & 63)))
            _ = fi22
            _ = v23
        }
    } else {
        var v24 uint = 255
        v3 -= (v3 / (((^(v3))) | 1))
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if false {
            var v26 string = ("" + "")
            if false {
                v27 := GenS1{F0: 65535, F1: 18446744073709551615, F2: 0}
                _ = v27
            } else {
                if (42 > 18446744073709551615) {
                    var v28 int = (fi25 / (((^(fi25))) | 1))
                    v29 := GenS2{F0: "\n\t", F1: -128}
                    _ = v28
                    _ = v29
                } else {
                    var v30 uint32 = 1
                    var v31 uint32 = 0
                    var v32 string = "a"
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
            _ = v26
        } else {
            v33 := GenS3{F0: 263, F1: 127}
            v34 := GenS4{F0: (^(v4)), F1: 2147483648}
            _ = v33
            _ = v34
        }
        v4 -= (^(v4))
        var v35 int64 = -9223372036854775808
        _ = fi25
        _ = v35
    }
    var v36 float64 = 0.1
    fmt.Println(v4, v4, v15)
    if (129 > 1) {
        v3 = (v15 + (v3 >> ((4294967295) & 63)))
    }
    if (false && (2147483648 == 65535)) {
        fmt.Println(v15, v3, v36)
    } else {
        if (((!(!true)) && false) && false) {
            v36 = v36
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 string = ("a" + "")
                var v40 int8 = 36
                if true {
                    var v41 int32 = (v3 << ((189) & 63))
                    v42 := GenS5{F0: 65535, F1: v4, F2: 1}
                    _ = v41
                    _ = v42
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
