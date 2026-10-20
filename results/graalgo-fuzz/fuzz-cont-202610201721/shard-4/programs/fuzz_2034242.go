package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int8, p1 int) uint {
    p0 = (p0 % (((p0 / (((p0 % ((1) | 1))) | 1))) | 1))
    p0 = (^(p0))
    var v5 uint8 = 128
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 int32 = -799
    var v2 string = "a"
    var v3 float64 = ((3.14159 * 1e10) / 3.14159)
    v4 := make(map[int]bool)
    fmt.Println(v2, v3, v3)
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            if (("foo bar" <= "hello") || true) {
                var v7 float32 = (0.0 + (1.5 / (1e10 / 1.5)))
                switch fi6 {
                case 157:
                    var v8 uint = 65535
                    _ = v8
                case -2147483648:
                    v9 := GenS1{F0: (855 <= 65535)}
                    _ = v9
                case 32768:
                    _ = len(v4)
                }
                _ = v7
            }
            fmt.Println(v3, v3, v2)
            fmt.Println(v2)
        } else {
            v10 := []int{(fi6 ^ fi6), -2147483648, (fi6 + fi6), fi6}
            _ = v10
        }
        fmt.Println(v2, v3, v3)
        _ = fi6
    }
    v1 += v1
    if false {
        v11 := GenS1{F0: true}
        var v12 int = -74
        v3 = 0.0
        _ = v11
        _ = v12
    } else {
        v13 := []int{}
        if false {
            switch v2 {
            case "foo bar":
                var v14 int8 = -29
                _ = v14
            case "hello":
                fmt.Println(v3, v1, v1)
            case "a":
                v2 = ((v2 + v2) + "a")
            default:
                var v15 int = -664
                _ = v15
            }
            switch 32768 {
            case 127:
                var v16 float64 = (v3 / v3)
                _ = v16
            case 0:
                var v17 int32 = -179
                _ = v17
            case 490:
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 int32 = 0
                    v2 = v2
                    fmt.Println(v1, fi18)
                    _ = fi18
                    _ = v19
                }
            default:
                switch -9223372036854775808 {
                case 1:
                    if false {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v4[fi20] = true
                            var v21 int64 = -1
                            fmt.Println(v21)
                            _ = fi20
                            _ = v21
                        }
                        var v22 uint8 = 42
                        _ = v22
                    } else {
                        v1 += (-(v1))
                        v2 = (v2 + v2)
                        if (-912 == (v1 + v1)) {
                            v23 := GenS1{F0: false}
                            v1 -= -2147483648
                            v24 := GenS2{F0: 2147483647}
                            _ = v23
                            _ = v24
                        }
                    }
                default:
                    v25 := GenS1{F0: false}
                    _ = v25
                }
            }
        } else {
            v26 := GenS1{F0: (!true)}
            _ = v26
        }
        _ = v13
    }
    v1 = 327
    v27 := GenS1{F0: (!true)}
    for fi28 := 0; fi28 < 17; fi28++ {
        switch v2 {
        case "foo bar":
            v29 := GenS3{F0: (v1 + v1)}
            _ = v29
        }
        _ = fi28
    }
    v30 := GenS4{F0: 4294967295}
    var v31 int8 = -45
    v2 = "foo bar"
    var v32 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
