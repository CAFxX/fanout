package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string }
func genFn0(p0 string) uint32 {
    var v4 int32 = 2147483647
    if (!(!(true && (1 > 100)))) {
        p0 = (p0 + ("foo bar" + p0))
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint16 = 2
            _ = fi5
            _ = v6
        }
    } else {
        if (-1 >= 802) {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 uint = 1
                _ = fi7
                _ = v8
            }
        }
        v4 += (v4 << ((320) & 63))
        var v9 bool = (-1 <= -752)
        _ = v9
    }
    if false {
        if (!(!true)) {
            var v10 int8 = -128
            _ = v10
        } else {
            var v11 int16 = 32767
            p0 = ""
            fmt.Println(v4, v4, v11)
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        if (-32768 <= 1) {
            v13 := GenS1{F0: 1, F1: 65535, F2: 4294967295}
            p0 = p0
            fmt.Println(v4, fi12, v4)
            _ = v13
        }
        if (!true) {
            for fi14 := 0; fi14 < 3; fi14++ {
                if (!false) {
                    if true {
                        v15 := []int{fi14, fi14, fi14, fi12}
                        fmt.Println(fi14, p0, p0)
                        _ = v15
                    } else {
                        v16 := GenS1{F0: -1, F1: 390, F2: 4294967295}
                        if false {
                            fmt.Println(fi12, p0, fi12)
                            v17 := GenS1{F0: 0, F1: 776, F2: 2147483648}
                            _ = v17
                        }
                        v4 += -314
                        _ = v16
                    }
                }
                v18 := []int{32768, fi12}
                v4 += v4
                _ = fi14
                _ = v18
            }
            v4 -= v4
        }
        _ = fi12
    }
    var v19 bool = (!(false || (float32(0.1) <= 0.1)))
    _ = p0
    _ = v19
    _ = v4
    return 169
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = -2.25
    v3 := make(map[uint32]bool)
    v1 = (v1 - v1)
    var v20 uint = 4294967295
    if true {
        for fi21 := 0; fi21 < 19; fi21++ {
            if false {
                v20 += v20
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 string = (("foo bar" + "\n\t") + "")
                    var v24 bool = true
                    var v25 uint16 = 149
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            v26 := GenS1{F0: -1, F1: 435, F2: 746}
            _ = fi21
            _ = v26
        }
        var v27 int32 = -1
        fmt.Println(v27, v1, v1)
        _ = v27
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS1{F0: 242, F1: 0, F2: genFn0(("" + "\n\t"))}
            v1 = (v1 << ((v20) & 63))
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS1{F0: 9223372036854775807, F1: 65535, F2: 146}
    v2 = v2
    var v31 int64 = 208
    v2 = ((1.5 / ((0.0 + (v2 * v2)) / 1e10)) * v2)
    var v32 int = 9223372036854775807
    if false {
        for fi33 := 0; fi33 < 14; fi33++ {
            for fi34 := 0; fi34 < 7; fi34++ {
                switch -19 {
                case 32768:
                    v35 := GenS1{F0: 9223372036854775807, F1: uint32(v20), F2: 186}
                    _ = v35
                case -1:
                    v36 := GenS1{F0: (^(v31)), F1: genFn0("foo bar"), F2: 2147483647}
                    _ = v36
                case 1:
                    if false {
                        var v37 string = "hello"
                        fmt.Println(fi34)
                        v38 := []int{fi33, (v32 << (((v20 & 255)) & 63)), fi34, 628}
                        _ = v37
                        _ = v38
                    } else {
                        v1 -= v1
                        _ = len(v3)
                        v20 = v20
                        var v39 uint16 = 0
                        _ = v39
                    }
                default:
                    v32 -= (v32 ^ 0)
                }
                v40 := GenS2{F0: genFn0((("" + "") + "\n\t")), F1: (^(v20))}
                _ = fi34
                _ = v40
            }
            v41 := GenS1{F0: 0, F1: uint32(v20), F2: genFn0("\n\t")}
            v20 = 1
            _ = fi33
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            if (true || false) {
                v2 = float64(3.14159)
            }
            v31 = v31
            for fi43 := 0; fi43 < 18; fi43++ {
                v20 += v20
                _ = fi43
            }
            _ = fi42
        }
        v44 := GenS3{F0: "hello"}
        _ = v44
    } else {
        fmt.Println(v32, v31)
    }
    var v45 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v45
}
