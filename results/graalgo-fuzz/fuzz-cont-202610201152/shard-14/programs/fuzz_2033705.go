package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint32; F2 string }
type GenS6 struct { F0 int32; F1 uint64 }
func genFn0(p0 int64, p1 uint16) uint {
    if (0 < 784) {
        var v4 int8 = 42
        _ = v4
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        if ((!(920 > 255)) && (((41 <= 0) && (129 != 979)) || true)) {
            var v6 int = fi5
            _ = v6
        } else {
            v7 := []int{}
            v8 := GenS1{F0: 42}
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int, p1 uint8, p2 int8) uint32 {
    v9 := GenS2{F0: 2, F1: (p0 | (p0 | p0)), F2: 255}
    if false {
        if (65535 >= 255) {
            p1 = p1
            v10 := GenS3{F0: ("hello" + ("hello" + "a"))}
            _ = v10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint64 = 328
            if false {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS4{F0: 783}
                    v15 := []int{(p0 / ((403) | 1)), (fi11 - p0), p0}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS5{F0: p1, F1: 2147483647, F2: "foo bar"}
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        var v17 uint = 4294967295
        _ = v17
    }
    var v18 uint = 65535
    if (int8(v18) >= p2) {
        v19 := GenS6{F0: -527, F1: 489}
        _ = v19
    } else {
        v20 := GenS5{F0: p1, F1: 2147483648, F2: "foo bar"}
        for fi21 := 0; fi21 < 14; fi21++ {
            p2 = p2
            _ = fi21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 14; fi23++ {
                if (2 == 0) {
                    if false {
                        var v24 uint32 = 0
                        if (!(!(int64(p2) > -183))) {
                            var v25 bool = false
                            var v26 int16 = -32768
                            v26 = -546
                            _ = v25
                            _ = v26
                        }
                        if (!(!true)) {
                            fmt.Println(p0)
                            v27 := []int{}
                            _ = len(v27)
                            _ = v27
                        }
                        _ = v24
                    }
                } else {
                    if ((!(9223372036854775807 == 852)) || (p1 < p1)) {
                        v18 = genFn0(-9223372036854775808, 32767)
                    } else {
                        switch (p0 - (^(fi22))) {
                        case -910:
                            p1 -= (^(p1))
                        case 453:
                            p2 += -1
                        case 2147483647:
                            p1 = (p1 / ((p1) | 1))
                        }
                    }
                    var v28 float64 = float64((0.0 / 3.14159))
                    _ = v28
                }
                switch fi23 {
                case -347:
                    if (p2 >= (p2 / ((1) | 1))) {
                        p1 -= 128
                        if (true && true) {
                            if false {
                                v29 := []int{p0, 32767}
                                p1 = p1
                                _ = v29
                            } else {
                                v30 := GenS3{F0: "foo bar"}
                                v31 := []int{}
                                _ = v30
                                _ = v31
                            }
                            p1 += (p1 + 0)
                            v32 := []int{(fi22 | (fi22 + (fi23 << ((v18) & 63))))}
                            _ = v32
                        }
                        var v33 uint = v18
                        v34 := []int{p0}
                        v18 += v33
                        _ = v33
                        _ = v34
                    }
                case -9223372036854775808:
                    if ((0 < p1) || true) {
                        v35 := GenS6{F0: -2147483648, F1: 214}
                        _ = v35
                    } else {
                        v36 := GenS1{F0: 1000}
                        _ = v36
                    }
                case 111:
                    var v37 float32 = 1e10
                    p0 = (p0 + fi23)
                    v37 = ((((((-2.25 - v37) / ((1.5 - float32(3.14159)) * 0.1)) + ((1.5 * 0.0) + v37)) * (v37 / -2.25)) / v37) / float32(3.14159))
                    _ = v37
                }
                _ = fi23
            }
            v18 = v18
            v18 += 42
            _ = fi22
        }
        _ = v20
    }
    var v38 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v38
    _ = v9
    return 1
}

func main() {
    var v1 bool = true
    var v2 bool = ("foo bar" < ("\n\t" + "foo bar"))
    var v3 float32 = 1.5
    v39 := GenS2{F0: 65535, F1: 32768, F2: 255}
    var v40 uint = 65535
    if false {
        fmt.Println(v1, v3, v1)
    } else {
        for fi41 := 0; fi41 < 18; fi41++ {
            v1 = false
            for fi42 := 0; fi42 < 3; fi42++ {
                fmt.Println(v3, v3)
                fmt.Println(v1)
                var v43 int16 = -449
                _ = fi42
                _ = v43
            }
            v40 += (v40 | (v40 * genFn0(1, 0)))
            _ = fi41
        }
        v44 := []int{32768, -197, 9223372036854775807}
        _ = v44
    }
    switch 2147483647 {
    case 589:
        if false {
            var v45 float64 = (0.1 + (((1.5 - float64(v3)) / 1.0) * float64(v3)))
            var v46 uint8 = 255
            var v47 bool = (0.0 == v3)
            _ = v45
            _ = v46
            _ = v47
        }
    }
    v2 = (!(!v2))
    v3 += float32(0.1)
    v3 = -2.25
    v40 = v40
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
