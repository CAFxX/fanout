package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int, p1 uint) int32 {
    p0 = int(p1)
    for fi5 := 0; fi5 < 20; fi5++ {
        if false {
            p0 = fi5
            p1 += (p1 + 65535)
            var v6 float32 = -2.25
            _ = v6
        } else {
            fmt.Println(p0, p1)
        }
        var v7 int16 = -264
        _ = fi5
        _ = v7
    }
    _ = p0
    _ = p1
    return 792
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 208
    var v3 bool = (2147483647 == -2147483648)
    v4 := make(map[uint32]string)
    if (!v3) {
        var v8 bool = (32767 == 793)
        v9 := []int{1, 32768, -17}
        v10 := GenS1{F0: 255, F1: (!(!(v8 || true)))}
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := []int{32767, -586, 128, 617}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int64 = 0
            _ = fi13
            _ = v14
        }
        fmt.Println(v3, v2)
        v1[-1] = fi12
        _ = fi12
    }
    v3 = (!false)
    v15 := GenS2{F0: 2, F1: "\n\t", F2: (v2 >> ((42) & 63))}
    if (genFn0(1, 1) != 42) {
        var v16 uint16 = 0
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 string = "foo bar"
            v19 := GenS3{F0: 0, F1: -504}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        v20 := []int{128, 128}
        if v3 {
            _ = len(v4)
            for fi21 := 0; fi21 < 16; fi21++ {
                v3 = (-1 >= -2147483648)
                if (42 == 162) {
                    if len(v20) > 0 {
                        v20[0] = fi21
                    }
                } else {
                    v2 += v2
                    v22 := GenS3{F0: 42, F1: -9223372036854775808}
                    switch "" {
                    case "":
                        switch 203 {
                        case 698:
                            var v23 float64 = 1e10
                            _ = v23
                        case -527:
                            var v24 int32 = -2147483648
                            _ = v24
                        }
                    case "foo bar":
                        if (!(v3 || (-2147483648 == int32(fi21)))) {
                            var v25 uint8 = 128
                            switch (fi21 << ((290) & 63)) {
                            case 42:
                                for fi26 := 0; fi26 < 6; fi26++ {
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    _ = fi26
                                }
                            case 763:
                                if (false || false) {
                                    v27 := GenS2{F0: 65535, F1: "hello", F2: 42}
                                    if len(v20) > 0 {
                                        v20[0] = (fi21 ^ 128)
                                    }
                                    _ = v27
                                } else {
                                    v3 = false
                                    fmt.Println(v2, v25)
                                }
                            default:
                                v3 = true
                            }
                            _ = v25
                        }
                    case "a":
                        var v28 uint64 = 21
                        _ = v28
                    default:
                        var v29 uint16 = 0
                        _ = v29
                    }
                    _ = v22
                }
                _ = fi21
            }
        }
        if (!v3) {
            v2 += v2
            fmt.Println(v3, v2)
            if (!(42 != 0)) {
                v30 := GenS1{F0: 1, F1: false}
                var v31 uint8 = v2
                v32 := GenS2{F0: 42, F1: "\n\t", F2: v31}
                _ = v30
                _ = v31
                _ = v32
            }
        } else {
            var v33 uint8 = v2
            v34 := GenS2{F0: 587, F1: "foo bar", F2: v33}
            _ = v33
            _ = v34
        }
        _ = v20
    }
    v35 := []int{9223372036854775807, -988, 128, -9223372036854775808}
    var v36 uint = 42
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 uint8 = (^(v2))
        var v39 uint64 = 0
        _ = fi37
        _ = v38
        _ = v39
    }
    v3 = true
    var v40 uint32 = 42
    v41 := GenS4{F0: -483}
    if len(v35) > 0 {
        v35[0] = 127
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 int64 = 165
        _ = len(v35)
        _ = fi42
        _ = v43
    }
    v44 := GenS3{F0: -9223372036854775808, F1: 802}
    var v45 float64 = 1.5
    v46 := GenS4{F0: -32768}
    var v47 int16 = 619
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
