package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 string }

func main() {
    var v1 int16 = 635
    v2 := make(map[string]uint32)
    var v3 bool = true
    var v4 float32 = 1e10
    switch -2147483648 {
    case -1:
        v4 += v4
    case 0:
        v3 = (v3 && v3)
    }
    var v5 int = -312
    v6 := v2["hello"]
    if true {
        delete(v2, "\n\t")
        fmt.Println(v4)
    }
    switch "a" {
    case "hello":
        v7 := GenS1{F0: 1}
        _ = v7
    default:
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int = v5
            if false {
                var v10 uint32 = v6
                v6 = (v10 - 277)
                v2[(("a" + "") + (("" + (("foo bar" + "foo bar") + "foo bar")) + ""))] = v6
                _ = v10
            } else {
                v5 = (v5 + -286)
            }
            _ = fi8
            _ = v9
        }
    }
    if ("" != "a") {
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := []int{}
            var v13 float64 = ((-2.25 * 3.14159) * 0.0)
            v14 := GenS1{F0: 4294967295}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    }
    if true {
        var v15 string = "foo bar"
        _ = v15
    } else {
        fmt.Println(v6, v6, v1)
    }
    _ = len(v2)
    for fi16 := 0; fi16 < 19; fi16++ {
        if (v3 || ((1 < 42) && (2 <= 255))) {
            v17 := GenS1{F0: 1}
            _ = v17
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v4 = (float32(1e10) / v4)
            v4 -= (((v4 / ((v4 - float32((1.5 + 1e10))) * (0.0 / v4))) + ((v4 * ((v4 / v4) - v4)) * 0.0)) * v4)
            switch "" {
            case "\n\t":
                delete(v2, "")
            case "hello":
                v4 = (0.1 + v4)
            case "foo bar":
                v4 = 1.5
            default:
                if ((v3 || ((fi16 / ((fi18) | 1)) < (v5 ^ (fi16 % ((921) | 1))))) || (190 != 1)) {
                    if (-56 != 1) {
                        if false {
                            v4 = 3.14159
                        } else {
                            v6 = 655
                            var v19 int64 = -1
                            _ = v19
                        }
                    }
                    v1 = 0
                } else {
                    v20 := []int{fi18, -360, 128, v5}
                    _ = v20
                }
            }
            _ = fi18
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(fi21)
            if (919 < 222) {
                switch 1 {
                case -2147483648:
                    fmt.Println(v1, fi16, v3)
                case 2147483647:
                    v1 -= (v1 * v1)
                }
            }
            var v22 string = ("foo bar" + "a")
            _ = fi21
            _ = v22
        }
        _ = fi16
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 int16 = v1
            var v26 string = "a"
            v27 := GenS1{F0: 894}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v28 int32 = 257
            _ = v28
        }
        _ = fi23
    }
    v29 := GenS2{F0: 1, F1: v3, F2: "\n\t"}
    for fi30 := 0; fi30 < 17; fi30++ {
        if (v3 || (241 == 1)) {
            fmt.Println(v6, v5)
            if (!((fi30 ^ fi30) != v5)) {
                fmt.Println(v4, fi30, v1)
            } else {
                var v31 uint8 = 178
                var v32 uint32 = v6
                _ = v31
                _ = v32
            }
        } else {
            var v33 int16 = (v1 / (((v1 * v1)) | 1))
            _ = v33
        }
        _ = fi30
    }
    fmt.Println(v1, v1)
    for fi34 := 0; fi34 < 16; fi34++ {
        switch ("hello" + "hello") {
        case "a":
            var v35 uint = 0
            _ = v35
        default:
            v36 := GenS1{F0: 1}
            _ = v36
        }
        _ = fi34
    }
    v37 := GenS2{F0: 64, F1: (v3 || ((0 < -622) && v3)), F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
