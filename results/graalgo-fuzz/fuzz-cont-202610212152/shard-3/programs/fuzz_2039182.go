package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32) bool {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float64 = -2.25
            p0 -= (p0 - p0)
            _ = fi7
            _ = v8
        }
        p0 -= (^(p0))
        _ = fi6
    }
    v9 := GenS1{F0: (p0 / ((p0) | 1))}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int = (fi10 + fi10)
        p0 = (p0 & p0)
        _ = fi10
        _ = v11
    }
    p0 += 0
    fmt.Println(p0)
    _ = p0
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -9223372036854775808
    var v3 float64 = 3.14159
    v4 := make(map[uint32]bool)
    var v5 int16 = 855
    v12 := GenS2{F0: 2, F1: 42}
    var v13 string = (("" + "hello") + ("" + "foo bar"))
    v14 := []int{-648}
    if true {
        fmt.Println(v13, v5, v3)
        if (32767 > (-(v5))) {
            for fi15 := 0; fi15 < 6; fi15++ {
                v5 += (v5 ^ (v5 + 1))
                v5 = (v5 % ((32767) | 1))
                _ = fi15
            }
            if (!false) {
                v2 -= v2
                v5 += v5
                v16 := GenS3{F0: 4294967295}
                _ = v16
            }
            var v17 int64 = (v2 / ((v2) | 1))
            _ = v17
        } else {
            v13 = v13
        }
        v18 := GenS4{F0: 220}
        _ = v18
    }
    switch -906 {
    case -842:
        var v19 int32 = -1
        _ = v19
    case -506:
        for fi20 := 0; fi20 < 5; fi20++ {
            if (false && true) {
                v5 += v5
                v2 += v2
            } else {
                switch fi20 {
                case -2147483648:
                    var v21 int8 = -1
                    _ = v21
                }
                var v22 float32 = -2.25
                _ = v22
            }
            if (int32(v2) != 1) {
                if (false && true) {
                    if (4294967295 == 42) {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        fmt.Println(v3, v13)
                        v14 = append(v14, fi20)
                    }
                    if ((453 != 2147483647) || false) {
                        _ = len(v14)
                        var v23 float64 = v3
                        _ = v23
                    }
                } else {
                    if true {
                        if len(v14) > 0 {
                            v14[0] = -2147483648
                        }
                        if false {
                            v3 = ((v3 / ((1.5 - (v3 * v3)) / v3)) + 0.0)
                            v13 = ("\n\t" + (v13 + v13))
                            fmt.Println(v5, v3)
                        } else {
                            var v24 float64 = v3
                            _ = v24
                        }
                    }
                }
            }
            _ = fi20
        }
    case -348:
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 string = (("a" + "\n\t") + (v13 + (v13 + v13)))
            v2 = (v2 - (v2 | (^(v2))))
            _ = fi25
            _ = v26
        }
    default:
        var v27 int8 = 100
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v1[fi28] = v13
        var v29 float64 = -2.25
        _ = fi28
        _ = v29
    }
    if (1e10 == -2.25) {
        switch (v13 + ((v13 + v13) + ("hello" + (v13 + v13)))) {
        case "":
            if len(v14) > 0 {
                v14[0] = -86
            }
        case "hello":
            delete(v1, 0)
        default:
            var v30 uint16 = 65535
            _ = v30
        }
        if (42 != 127) {
            if (255 == 255) {
                var v31 int8 = 0
                _ = v31
            } else {
                var v32 float64 = float64(-2.25)
                v2 = v2
                _ = v32
            }
        } else {
            var v33 string = v13
            _ = v33
        }
        var v34 int16 = v5
        _ = v34
    }
    v3 = v3
    if true {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
