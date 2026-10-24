package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 17; fi6++ {
        if false {
            var v7 bool = true
            var v8 float64 = 0.1
            for fi9 := 0; fi9 < 13; fi9++ {
                if v7 {
                    if ((18446744073709551615 == 0) || (!(65535 < 658))) {
                        var v10 uint32 = 88
                        _ = v10
                    }
                    var v11 int16 = 0
                    v12 := GenS1{F0: 128, F1: 127}
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(v7, fi9)
                    var v13 int32 = -2147483648
                    _ = v13
                }
                _ = fi9
            }
            _ = v7
            _ = v8
        }
        var v14 bool = ("a" < "foo bar")
        v15 := []int{(^(fi6)), (fi6 / ((fi6) | 1)), fi6}
        _ = fi6
        _ = v14
        _ = v15
    }
    var v16 uint32 = 1
    v17 := GenS1{F0: 127, F1: 255}
    _ = v16
    _ = v17
    return 42
}
func genFn1(p0 int16, p1 int8) uint32 {
    var v18 int16 = p0
    p1 -= (^(p1))
    fmt.Println(v18)
    _ = p0
    _ = p1
    _ = v18
    return 576
}

func main() {
    var v1 float32 = ((0.1 + 0.0) * float32(0.0))
    var v2 uint64 = 42
    v3 := []uint32{144}
    var v4 float32 = (float32(3.14159) + ((v1 / (v1 + v1)) - v1))
    v5 := make(map[string]uint32)
    var v19 float64 = 1e10
    fmt.Println(v4, v19, v2)
    if (true && true) {
        v20 := GenS2{F0: 18446744073709551615, F1: ("hello" + ("hello" + "foo bar")), F2: 231}
        for fi21 := 0; fi21 < 19; fi21++ {
            v19 = 0.0
            v1 -= v4
            v4 = ((3.14159 / v4) * v1)
            _ = fi21
        }
        _ = v20
    } else {
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS2{F0: 217, F1: "\n\t", F2: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v24 uint16 = 32767
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 int64 = -1
        for fi26 := 0; fi26 < 14; fi26++ {
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 bool = false
                _ = fi27
                _ = v28
            }
            var v29 int32 = 685
            _ = fi26
            _ = v29
        }
        _ = v25
    }
    if (662 <= 153) {
        fmt.Println(v2, v2)
        var v30 uint8 = 111
        switch -762 {
        case 649:
            v31 := GenS2{F0: 18446744073709551615, F1: (((((("" + "foo bar") + "hello") + ("foo bar" + "foo bar")) + "\n\t") + "") + "hello"), F2: -295}
            _ = v31
        default:
            if true {
                var v32 string = (("" + ("foo bar" + "foo bar")) + ("hello" + ("hello" + "a")))
                var v33 uint32 = 2147483647
                _ = v32
                _ = v33
            }
        }
        _ = v30
    } else {
        v2 -= (v2 | (v2 * 42))
    }
    v19 += (v19 * 1e10)
    v1 = (v1 * v4)
    if (2 >= 1) {
        fmt.Println(v19, v2, v2)
        v34 := GenS2{F0: 65535, F1: "\n\t", F2: 0}
        if (650 != 0) {
            if (0 != 0) {
                var v35 uint16 = 2
                v36 := GenS2{F0: 42, F1: "a", F2: -1}
                _ = v35
                _ = v36
            } else {
                v37 := v5[""]
                fmt.Println(v4, v2, v37)
                _ = v37
            }
            delete(v5, "\n\t")
        }
        _ = v34
    }
    v19 -= (3.14159 - 1e10)
    v38 := v5["hello"]
    var v39 uint64 = v2
    if false {
        v40 := GenS3{F0: -589}
        var v41 uint = 722
        _ = v40
        _ = v41
    }
    var v42 int = 264
    v4 += float32(((1e10 * -2.25) / -2.25))
    for fi43 := 0; fi43 < 6; fi43++ {
        v44 := GenS3{F0: -9223372036854775808}
        for fi45 := 0; fi45 < 14; fi45++ {
            var v46 int32 = -1
            v42 -= v42
            _ = fi45
            _ = v46
        }
        for fi47 := 0; fi47 < 18; fi47++ {
            var v48 float64 = float64(0.1)
            v49 := GenS4{F0: 531, F1: 791}
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
