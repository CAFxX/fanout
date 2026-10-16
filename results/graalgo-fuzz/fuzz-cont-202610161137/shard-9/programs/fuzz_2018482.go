package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int8 {
    v4 := []int{-1}
    switch "a" {
    case "a":
        fmt.Println(p0)
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if true {
            v5 := GenS1{F0: 1, F1: -96}
            if len(v4) > 0 {
                v4[0] = 32768
            }
            p0 = (p0 >> ((1) & 63))
            _ = v5
        } else {
            var v6 float32 = (1.5 + 1e10)
            var v7 int = -2147483648
            var v8 float32 = v6
            _ = v6
            _ = v7
            _ = v8
        }
    }
    if (-867 >= -935) {
        for fi9 := 0; fi9 < 5; fi9++ {
            if (fi9 > fi9) {
                var v10 float32 = 3.14159
                fmt.Println(v10)
                _ = v10
            } else {
                var v11 float32 = (float32(float64(0.0)) * 0.1)
                if true {
                    p0 = 128
                } else {
                    var v12 int64 = -9223372036854775808
                    var v13 int = fi9
                    _ = v12
                    _ = v13
                }
                _ = v11
            }
            _ = fi9
        }
        fmt.Println(p0)
        _ = len(v4)
    } else {
        v14 := GenS1{F0: uint64(p0), F1: -119}
        _ = v14
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 bool = (2 < 32767)
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    var v15 uint32 = 807
    if v1 {
        if ((0.0 + float32((float64(1.5) * float64((1e10 * 3.14159))))) != (3.14159 / 3.14159)) {
            v16 := GenS1{F0: 2, F1: 28}
            _ = v16
        } else {
            if false {
                var v17 float32 = -2.25
                var v18 int = -1
                v19 := []int{(^(v18)), v18, (v18 / ((-810) | 1)), 9223372036854775807}
                _ = v17
                _ = v18
                _ = v19
            }
        }
    }
    var v20 uint32 = 120
    v21 := GenS1{F0: 18446744073709551615, F1: 1}
    v22 := GenS1{F0: 42, F1: genFn0(127)}
    v2 += v2
    var v23 float32 = 0.1
    v24 := GenS1{F0: uint64(v2), F1: -1}
    v15 += v15
    var v25 float64 = 0.0
    var v26 uint64 = 0
    var v27 int8 = genFn0(1)
    if (!(false && true)) {
        var v28 int64 = -1
        for fi29 := 0; fi29 < 1; fi29++ {
            fmt.Println(v1, v2, v26)
            switch "hello" {
            case "\n\t":
                switch (fi29 + 521) {
                case -9223372036854775808:
                    if ((!(uint16(v26) <= 0)) && v1) {
                        v2 = 254
                        for fi30 := 0; fi30 < 14; fi30++ {
                            var v31 uint = (v2 << ((v2) & 63))
                            _ = fi30
                            _ = v31
                        }
                        v32 := v3[("foo bar" + "\n\t")]
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 19; fi33++ {
                            v27 = 127
                            _ = fi33
                        }
                    }
                }
            case "a":
                var v34 int = fi29
                _ = v34
            }
            v35 := []int{}
            _ = fi29
            _ = v35
        }
        _ = v28
    }
    v23 = v23
    var v36 uint32 = 65535
    var v37 uint64 = v26
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := GenS2{F0: (v1 && v1)}
        _ = fi38
        _ = v39
    }
    delete(v3, "")
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
}
