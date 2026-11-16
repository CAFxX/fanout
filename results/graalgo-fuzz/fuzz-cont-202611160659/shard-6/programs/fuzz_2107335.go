package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 int16 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0(p0 string, p1 int16) int8 {
    if (-9223372036854775808 >= 9223372036854775807) {
        var v7 float32 = 1.5
        _ = v7
    } else {
        p1 = p1
    }
    if false {
        v8 := []int{-868, 2147483647, 127}
        p0 = p0
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v8
    } else {
        switch -757 {
        case -1:
            var v9 int = -2147483648
            _ = v9
        }
    }
    v10 := []int{615, -32, 1, 2147483647}
    for fi11 := 0; fi11 < 2; fi11++ {
        p1 += (p1 % (((p1 / ((p1) | 1))) | 1))
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: (p1 % ((-1) | 1))}
            if (42 != 947) {
                fmt.Println(p1, p0)
                v14 := GenS1{F0: 1000}
                if false {
                    v15 := GenS1{F0: (p1 * (p1 % ((-293) | 1)))}
                    var v16 uint8 = 1
                    _ = v15
                    _ = v16
                } else {
                    if len(v10) > 0 {
                        v10[0] = -780
                    }
                    p1 += p1
                }
                _ = v14
            } else {
                v17 := GenS1{F0: (^(p1))}
                if true {
                    if false {
                        var v18 int16 = (^(p1))
                        var v19 bool = false
                        var v20 uint32 = 340
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    if true {
                        var v21 int32 = 2147483647
                        _ = v21
                    }
                }
                _ = v17
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v10 = append(v10, -937)
    _ = p0
    _ = p1
    _ = v10
    return 6
}
func genFn1() int32 {
    v22 := GenS1{F0: 0}
    var v23 int32 = 0
    _ = v22
    _ = v23
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 uint = 1
    v4 := []uint32{2147483647, 465}
    v5 := make(map[uint32]string)
    v6 := []uint32{620, 2147483648, 2147483648}
    v3 = 255
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 float32 = 0.1
        var v26 int8 = -76
        var v27 bool = false
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: -185}
    if (!(255 == 0)) {
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS1{F0: 91}
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v3)
    for fi31 := 0; fi31 < 12; fi31++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi32 := 0; fi32 < 6; fi32++ {
            _ = len(v1)
            if (fi31 > -2147483648) {
                var v33 uint = 4294967295
                switch 2147483647 {
                case 0:
                    v34 := GenS2{F0: -2147483648, F1: 2, F2: 0}
                    _ = v34
                case -9223372036854775808:
                    var v35 float32 = 0.0
                    _ = v35
                case -2147483648:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v33
            }
            v36 := GenS3{F0: "a", F1: 2}
            _ = fi32
            _ = v36
        }
        _ = fi31
    }
    var v37 uint = 255
    _ = len(v5)
    var v38 int = 433
    if true {
        var v39 int16 = -1
        _ = v39
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    if false {
        v5[65535] = "foo bar"
    } else {
        switch "hello" {
        case "":
            if false {
                var v40 int32 = 0
                var v41 uint8 = 128
                _ = v40
                _ = v41
            }
        case "hello":
            _ = len(v2)
        case "\n\t":
            for fi42 := 0; fi42 < 17; fi42++ {
                v3 = (^(v3))
                _ = fi42
            }
        }
        switch (v38 - -674) {
        case 736:
            v38 -= v38
        case 127:
            fmt.Println(v37, v37)
        default:
            v43 := GenS4{F0: 0, F1: false, F2: 103}
            _ = v43
        }
    }
    v44 := GenS1{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
