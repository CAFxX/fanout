package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int16 }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    v7 := []int{32767, 137, 2147483647}
    var v8 float64 = ((3.14159 / -2.25) / 0.1)
    if len(v7) > 0 {
        v7[0] = 127
    }
    if (2 <= 32767) {
        var v9 int16 = 594
        if len(v7) > 0 {
            v7[0] = 127
        }
        var v10 float32 = float32(v8)
        _ = v10
        _ = v9
    } else {
        fmt.Println(v8)
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 128
    var v3 float32 = (1e10 + float32(0.1))
    var v4 uint8 = (v2 | v2)
    var v5 int16 = 1
    var v6 uint32 = 4294967295
    var v11 uint16 = 0
    var v12 bool = (v4 >= v2)
    if v12 {
        var v13 int16 = v5
        _ = v13
    }
    var v14 float32 = (float32(1.5) - ((v3 / 1.0) + v3))
    v11 = (v11 + (v11 ^ v11))
    var v15 bool = (!true)
    v11 = (^(v11))
    if false {
        v16 := []int{}
        _ = v16
    }
    if (9223372036854775807 == -9223372036854775808) {
        fmt.Println(v6, v12, v12)
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint8 = 42
            _ = fi17
            _ = v18
        }
        if (0.0 > ((1.5 + (float64(v14) * (0.0 / 3.14159))) + 1.5)) {
            delete(v1, "")
            fmt.Println(v3)
        } else {
            delete(v1, "foo bar")
            v12 = (!v12)
        }
    } else {
        fmt.Println(v12, v5, v2)
        v19 := GenS1{F0: v11, F1: "a", F2: v5}
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 float64 = 1.5
        fmt.Println(fi20, v21, v12)
        _ = fi20
        _ = v21
    }
    if (float64(float32((float64((v14 - (1.5 + 0.0))) / ((1.5 - 1.5) + 1.5)))) != -2.25) {
        if (!v15) {
            v3 = v3
            v22 := []int{128, -471, 42, 127}
            for fi23 := 0; fi23 < 6; fi23++ {
                fmt.Println(v4, v3, v14)
                v3 -= v3
                var v24 float64 = (((3.14159 + ((3.14159 + 1.5) * 1e10)) - 1e10) / (1.5 / 1.0))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        if false {
            v1[(("" + "a") + "a")] = (v6 / (((v6 - 4294967295)) | 1))
        }
    }
    switch "foo bar" {
    case "foo bar":
        v25 := []int{-763, 2147483647, int(v6), 357}
        _ = v25
    case "":
        var v26 int32 = -686
        _ = v26
    case "\n\t":
        v12 = false
    }
    if v15 {
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := []int{}
            fmt.Println(fi27, v4)
            var v29 uint8 = v4
            _ = fi27
            _ = v28
            _ = v29
        }
        if true {
            for fi30 := 0; fi30 < 9; fi30++ {
                v6 = v6
                var v31 uint16 = (v11 % ((v11) | 1))
                var v32 bool = true
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 string = "a"
            var v34 string = v33
            _ = v33
            _ = v34
        } else {
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 int16 = (v5 * v5)
                v11 = (v11 & 65535)
                for fi37 := 0; fi37 < 13; fi37++ {
                    switch -312 {
                    case -976:
                        if true {
                            var v38 int32 = 2147483647
                            var v39 int = fi35
                            _ = v38
                            _ = v39
                        }
                    case 63:
                        v12 = true
                    case -226:
                        v40 := []int{(fi35 & (fi35 + (fi35 | fi37)))}
                        _ = v40
                    }
                    var v41 uint32 = 183
                    _ = fi37
                    _ = v41
                }
                _ = fi35
                _ = v36
            }
        }
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        v4 += uint8(v6)
        _ = fi42
    }
    if (!(2 == 0)) {
        if v15 {
            for fi43 := 0; fi43 < 10; fi43++ {
                switch (fi43 << ((65535) & 63)) {
                case 0:
                    v15 = false
                case 32768:
                    v12 = (0 > 42)
                case 9223372036854775807:
                    var v44 float32 = v14
                    _ = v44
                }
                fmt.Println(v12, v11)
                _ = fi43
            }
            v45 := GenS1{F0: (^(v11)), F1: "foo bar", F2: -1}
            _ = v45
        }
        v46 := GenS1{F0: (v11 / ((462) | 1)), F1: "", F2: v5}
        fmt.Println(v14, v5, v12)
        _ = v46
    }
    if true {
        switch 176 {
        case -9223372036854775808:
            v47 := GenS2{F0: 1}
            _ = v47
        case -1:
            for fi48 := 0; fi48 < 16; fi48++ {
                v49 := GenS1{F0: (v11 >> ((529) & 63)), F1: ("" + "a"), F2: v5}
                _ = fi48
                _ = v49
            }
        }
        var v50 bool = (268 == v11)
        _ = v50
    } else {
        fmt.Println(v12, v6, v5)
        v51 := GenS2{F0: 4294967295}
        _ = v51
    }
    v11 = (v11 | v11)
    fmt.Println(v12, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
