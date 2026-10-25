package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int8) uint64 {
    p0 = (p0 % ((-128) | 1))
    fmt.Println(p0)
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            p0 = (p0 | (p0 << ((18446744073709551615) & 63)))
            v7 := GenS1{F0: false, F1: "\n\t"}
            fmt.Println(p0, fi6)
            p0 = (^(p0))
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 bool = (!((3.14159 >= (float32(0.1) + 0.1)) || true))
    v9 := []int{2147483647, 1}
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 42
    var v3 int32 = 956
    var v4 string = "a"
    if false {
        var v10 int32 = (v3 + v3)
        fmt.Println(v2, v4)
        for fi11 := 0; fi11 < 1; fi11++ {
            switch v4 {
            case "a":
                var v12 float32 = float32(-2.25)
                _ = v12
            case "":
                fmt.Println(v3)
            case "\n\t":
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint8 = 195
                    if false {
                        fmt.Println(v4)
                        if true {
                            v15 := GenS2{F0: 340, F1: 255}
                            var v16 bool = (!(true || true))
                            if true {
                                v14 -= (v14 | (^(v14)))
                                var v17 float32 = (1.5 * 3.14159)
                                var v18 int64 = 1
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint = 18446744073709551615
                                var v20 float32 = -2.25
                                if (float64(v20) != 0.1) {
                                    var v21 uint32 = 4294967295
                                    v22 := GenS2{F0: 374, F1: 0}
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                                _ = v20
                            }
                            _ = v15
                            _ = v16
                        } else {
                            if (!true) {
                                var v23 int64 = 9223372036854775807
                                _ = v23
                            } else {
                                v24 := GenS1{F0: ((v2 / ((v2) | 1)) >= (v2 / ((v2) | 1))), F1: (v4 + "\n\t")}
                                var v25 int16 = -931
                                v26 := GenS2{F0: 257, F1: 255}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        }
                        var v27 uint16 = 32767
                        _ = v27
                    }
                    fmt.Println(fi11, v2, v4)
                    var v28 bool = true
                    _ = fi13
                    _ = v14
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := []int{fi11, 1, fi11}
                v10 = v3
                _ = fi29
                _ = v30
            }
            _ = fi11
        }
        _ = v10
    } else {
        var v31 float32 = ((float32(0.1) * float32(3.14159)) - float32(-2.25))
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS1{F0: (int32(v2) < (v3 >> ((1) & 63))), F1: ((v4 + "hello") + "\n\t")}
            if false {
                for fi34 := 0; fi34 < 5; fi34++ {
                    v2 += 2147483647
                    _ = fi34
                }
            } else {
                v4 = (v4 + v4)
                v1[v4] = (!false)
            }
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    v3 += -1
    var v35 string = ((v4 + (v4 + v4)) + v4)
    var v36 uint = 4294967295
    v4 = (("foo bar" + (v4 + "a")) + ((v4 + (v35 + (v4 + v4))) + "a"))
    v37 := v1["\n\t"]
    v38 := []int{42, 32767, -796, 128}
    v39 := v1[v35]
    var v40 float64 = ((0.0 / (3.14159 / 1.5)) - 0.1)
    var v41 uint16 = 898
    _ = len(v38)
    v40 = 1e10
    fmt.Println(v37)
    for fi42 := 0; fi42 < 12; fi42++ {
        if v39 {
            for fi43 := 0; fi43 < 3; fi43++ {
                v44 := GenS1{F0: (!(42 != -1)), F1: (("a" + (((v35 + "hello") + "foo bar") + v35)) + v35)}
                _ = fi43
                _ = v44
            }
            fmt.Println(v35, v41, fi42)
            var v45 uint8 = 2
            _ = v45
        }
        v36 -= v36
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
