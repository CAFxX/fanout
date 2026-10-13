package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint16, p1 bool, p2 uint16) uint {
    p2 = p2
    var v6 uint32 = 2
    p2 = p2
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1(p0 uint32) int {
    v7 := GenS1{F0: 983, F1: 32768, F2: 2147483647}
    v8 := GenS1{F0: 0, F1: 32767, F2: -546}
    var v9 int16 = 935
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(v9, v9, v9)
            if ((18446744073709551615 == 2) && true) {
                var v12 int8 = -128
                var v13 string = ""
                switch (^(fi11)) {
                case 32767:
                    v12 -= v12
                }
                v14 := GenS1{F0: (^(fi10)), F1: fi11, F2: 774}
                fmt.Println(p0)
                _ = v12
                _ = v13
                _ = v14
            } else {
                var v15 int32 = 1
                v16 := []int{}
                _ = v15
                _ = v16
            }
            if (!(944 > -857)) {
                v17 := GenS1{F0: fi10, F1: fi11, F2: int32(v9)}
                _ = v17
            } else {
                switch fi10 {
                case 127:
                    var v18 uint32 = (p0 + 42)
                    _ = v18
                case 1:
                    p0 = 1
                case 32768:
                    var v19 uint8 = 255
                    _ = v19
                default:
                    v9 = 42
                }
                var v20 int = 4294967295
                v21 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615, F2: 0}
                _ = v20
                _ = v21
            }
            _ = fi11
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint8 = 255
            if true {
                fmt.Println(fi22, fi10)
                if true {
                    v24 := []int{32767, 549, (-(fi10))}
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    v25 := GenS2{F0: -935, F1: 0, F2: 65535}
                    if (-9223372036854775808 == -662) {
                        v23 -= v23
                        var v26 float64 = ((0.0 + (-2.25 - 1.5)) * -2.25)
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        _ = v26
                    } else {
                        _ = len(v24)
                    }
                    _ = v24
                    _ = v25
                } else {
                    v27 := []int{127}
                    _ = v27
                }
            } else {
                var v28 float32 = 0.0
                _ = v28
            }
            _ = fi22
            _ = v23
        }
        if (2 <= 1) {
            if true {
                fmt.Println(v9, fi10)
            } else {
                var v29 float64 = ((-2.25 * -2.25) - 0.0)
                p0 += (p0 & p0)
                _ = v29
            }
            fmt.Println(p0, v9, fi10)
            var v30 bool = (118 != -26)
            _ = v30
        }
        _ = fi10
    }
    v31 := GenS2{F0: -639, F1: 180, F2: 2}
    _ = p0
    _ = v31
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    v1 := []uint32{261}
    var v2 uint32 = 398
    var v3 int64 = -476
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    fmt.Println(v3)
    if (1 >= 712) {
        if true {
            var v32 int = 42
            if (!false) {
                v33 := v5["a"]
                if (true || v33) {
                    _ = len(v1)
                    var v34 float32 = -2.25
                    _ = v34
                }
                v2 = v2
                _ = v33
            }
            _ = v32
        }
        var v35 uint = 18446744073709551615
        var v36 uint8 = uint8(v35)
        _ = v35
        _ = v36
    }
    v1 = append(v1, v2)
    for fi37 := 0; fi37 < 12; fi37++ {
        fmt.Println(v2)
        v2 += v2
        v3 -= v3
        _ = fi37
    }
    v38 := GenS3{F0: 128, F1: v3}
    var v39 uint8 = 128
    var v40 bool = (-1 >= v3)
    switch 1 {
    case -1:
        var v41 uint32 = 65535
        _ = v41
    case -123:
        if v40 {
            if (false && v40) {
                v3 = 42
                for fi42 := 0; fi42 < 4; fi42++ {
                    var v43 int16 = 1
                    var v44 uint8 = 240
                    var v45 uint32 = v2
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                }
            } else {
                fmt.Println(v40)
                var v46 int8 = 42
                _ = v46
            }
        }
    }
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v47 int64 = -774
    var v48 int64 = v47
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v48
    _ = v5
}
