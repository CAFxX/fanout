package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() int8 {
    var v6 int16 = 832
    var v7 uint32 = 65535
    switch "hello" {
    case "hello":
        var v8 int = 127
        _ = v8
    case "a":
        var v9 uint32 = v7
        _ = v9
    default:
        v10 := GenS1{F0: 507, F1: 33}
        _ = v10
    }
    var v11 int32 = 42
    _ = v11
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 bool, p1 int, p2 int64) string {
    v12 := []int{}
    var v13 int = p1
    var v14 bool = (p0 || p0)
    var v15 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return ("\n\t" + ("" + "foo bar"))
}

func main() {
    var v1 uint64 = 1
    var v2 uint = 4294967295
    var v3 uint8 = 0
    v4 := make(map[int]string)
    var v5 int32 = 485
    var v16 float32 = float32(-2.25)
    if ((true && false) && false) {
        v17 := v4[-575]
        fmt.Println(v3, v17)
        for fi18 := 0; fi18 < 1; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                v1 = v1
                var v20 float32 = 0.1
                v21 := []int{}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    }
    switch -961 {
    case 127:
        v22 := GenS1{F0: 2, F1: v3}
        _ = v22
    }
    var v23 int16 = 1000
    v24 := GenS2{F0: 0}
    if (!(!(true || (v2 >= (^(v2)))))) {
        var v25 int64 = -9223372036854775808
        _ = v25
    } else {
        v4[127] = "a"
    }
    var v26 int64 = 803
    if (v5 <= (v5 * (v5 | (v5 / ((-1) | 1))))) {
        v23 = (^(v23))
        if false {
            v27 := []int{2147483647, -449}
            v28 := v4[-9223372036854775808]
            if (v28 > "a") {
                v1 -= 18446744073709551615
                v23 = v23
                v16 = v16
            }
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 13; fi29++ {
                fmt.Println(v16, v23, v5)
                for fi30 := 0; fi30 < 13; fi30++ {
                    v2 = 1
                    v31 := GenS1{F0: (v2 << ((1) & 63)), F1: 127}
                    _ = fi30
                    _ = v31
                }
                v32 := GenS1{F0: v2, F1: v3}
                _ = fi29
                _ = v32
            }
        }
        switch -67 {
        case -382:
            if true {
                var v33 float64 = float64(((1e10 + v16) - (0.1 + (0.1 * (-2.25 / (v16 - (float32(0.0) - v16)))))))
                _ = v33
            }
        case 32768:
            var v34 uint8 = (^(v3))
            _ = v34
        }
    } else {
        fmt.Println(v16, v1)
    }
    v23 += (v23 & v23)
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := v4[fi35]
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := v4[fi37]
            if (!(3.14159 <= (-2.25 / 1.0))) {
                if (1 >= 427) {
                    var v39 uint32 = 927
                    fmt.Println(v36)
                    var v40 uint8 = (v3 + 1)
                    _ = v39
                    _ = v40
                } else {
                    fmt.Println(v23)
                    v23 -= v23
                }
                if (true && false) {
                    v41 := GenS1{F0: 381, F1: v3}
                    _ = v41
                } else {
                    v3 -= (^(v3))
                }
                v42 := v4[(fi35 | (fi35 - int(v1)))]
                _ = v42
            }
            v43 := GenS2{F0: 23}
            var v44 uint = uint(v1)
            fmt.Println(v26)
            _ = fi37
            _ = v38
            _ = v43
            _ = v44
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v16, v3)
    for fi45 := 0; fi45 < 11; fi45++ {
        var v46 string = ("hello" + "foo bar")
        var v47 int16 = -580
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
