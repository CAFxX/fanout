package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 bool; F2 string }
func genFn0() uint32 {
    if (!false) {
        v7 := GenS1{F0: 1, F1: "\n\t"}
        _ = v7
    }
    var v8 string = "hello"
    v8 = "\n\t"
    _ = v8
    return 42
}
func genFn1(p0 uint8) int64 {
    v9 := []int{127}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int8 = 100
        p0 = (p0 >> ((0) & 63))
        _ = fi10
        _ = v11
    }
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := []int{415, 4294967295, -782, 9223372036854775807}
    var v2 int64 = -1
    var v3 int64 = (v2 / ((v2) | 1))
    var v4 bool = (float32(3.14159) < -2.25)
    var v5 uint16 = 32767
    var v6 uint32 = 0
    var v12 bool = ((3.14159 + 3.14159) != 0.1)
    v12 = v4
    fmt.Println(v12, v5, v3)
    v4 = false
    v13 := GenS1{F0: 1, F1: "hello"}
    for fi14 := 0; fi14 < 3; fi14++ {
        v4 = (((1000 <= 1000) && (4294967295 != 42)) && v12)
        _ = fi14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            if (!v12) {
                fmt.Println(v3, v2, fi15)
            }
            var v17 uint16 = (v5 << ((42) & 63))
            v18 := GenS1{F0: 116, F1: "\n\t"}
            var v19 uint = 956
            v20 := GenS1{F0: -182, F1: ("foo bar" + "")}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
        }
        v2 += (v3 / ((v2) | 1))
        _ = fi15
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 string = "foo bar"
        for fi23 := 0; fi23 < 14; fi23++ {
            v2 = (^(v3))
            var v24 uint64 = 42
            _ = fi23
            _ = v24
        }
        v3 += (^(v3))
        _ = fi21
        _ = v22
    }
    var v25 bool = false
    v26 := GenS2{F0: (!(!v12)), F1: (true || v4), F2: "hello"}
    var v27 float64 = 0.1
    switch -466 {
    case 32768:
        fmt.Println(v12, v5)
    case 927:
        fmt.Println(v27, v6, v27)
    case -704:
        if (!v4) {
            v28 := GenS2{F0: v4, F1: v25, F2: (("hello" + "foo bar") + "")}
            _ = v28
        } else {
            if v4 {
                for fi29 := 0; fi29 < 13; fi29++ {
                    v4 = v25
                    v25 = v12
                    _ = fi29
                }
                var v30 int = -9223372036854775808
                _ = v30
            } else {
                if v12 {
                    var v31 uint32 = v6
                    _ = v31
                }
                v4 = true
                if (0 <= 32767) {
                    var v32 uint16 = (v5 + v5)
                    if true {
                        v1 = append(v1, -2147483648)
                    }
                    _ = v32
                } else {
                    var v33 uint16 = (v5 >> ((394) & 63))
                    _ = v33
                }
            }
            var v34 int8 = 100
            _ = v34
        }
    }
    v27 -= (v27 * v27)
    var v35 string = ""
    var v36 int = -353
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
