package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32, p1 bool, p2 uint) bool {
    fmt.Println(p1)
    v7 := []int{32768, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 int = 1
    v2 := []int{}
    var v3 uint64 = 42
    v4 := []uint32{762, 675}
    var v5 bool = true
    var v6 uint16 = 705
    fmt.Println(v6, v3)
    var v8 int32 = 2147483647
    v5 = v5
    v6 = (^(v6))
    fmt.Println(v5, v5)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if v5 {
        switch ("a" + "") {
        case "a":
            if (!false) {
                fmt.Println(v6, v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if v5 {
                    v1 = 4294967295
                    for fi9 := 0; fi9 < 2; fi9++ {
                        if (!v5) {
                            v1 = (fi9 ^ fi9)
                        } else {
                            var v10 int16 = 1000
                            var v11 uint8 = 2
                            _ = v10
                            _ = v11
                        }
                        v12 := GenS1{F0: 42, F1: v6}
                        if v5 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v3 -= v3
                        } else {
                            var v13 uint64 = 1
                            var v14 int32 = v8
                            v15 := GenS1{F0: 4294967295, F1: 2}
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        _ = fi9
                        _ = v12
                    }
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 uint8 = 255
                        v8 += v8
                        _ = fi16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 int32 = 461
                        v20 := GenS1{F0: 65535, F1: v6}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 bool = v5
                    v5 = (v21 || (32767 >= 1000))
                    _ = v21
                }
            } else {
                v6 -= (^(v6))
                v4 = append(v4, 868)
                if (-361 < (v1 & v1)) {
                    v5 = (v3 != 42)
                    if (18446744073709551615 <= v3) {
                        var v22 int16 = 1
                        _ = len(v2)
                        _ = v22
                    } else {
                        _ = len(v4)
                        var v23 uint64 = v3
                        _ = v23
                    }
                }
            }
        }
        for fi24 := 0; fi24 < 8; fi24++ {
            if false {
                if ((-(fi24)) <= fi24) {
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if (false || v5) {
                            fmt.Println(fi25, v3, v1)
                            v26 := GenS1{F0: 390, F1: v6}
                            fmt.Println(v5, v3, v1)
                            v8 += (v8 * v8)
                            _ = v26
                        }
                        v3 -= (v3 + v3)
                        if len(v4) > 0 {
                            v4[0] = 54
                        }
                        _ = fi25
                    }
                    var v27 uint = 4294967295
                    v6 = 255
                    _ = v27
                }
            } else {
                v6 = v6
            }
            v2 = append(v2, fi24)
            for fi28 := 0; fi28 < 13; fi28++ {
                var v29 uint64 = (v3 / (((v3 - v3)) | 1))
                var v30 string = "a"
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi24
        }
    } else {
        switch v1 {
        case -2147483648:
            for fi31 := 0; fi31 < 3; fi31++ {
                v32 := GenS1{F0: 93, F1: 32767}
                _ = fi31
                _ = v32
            }
        case 128:
            _ = len(v2)
        default:
            var v33 int8 = int8(v3)
            _ = v33
        }
        v2 = append(v2, (v1 * -9223372036854775808))
        v2 = append(v2, v1)
    }
    v5 = v5
    if ((-(v1)) <= 974) {
        var v34 float64 = 1e10
        v8 += v8
        _ = v34
    }
    v35 := GenS2{F0: v1}
    var v36 int32 = (v8 & (v8 & (-(v8))))
    switch -929 {
    case 4294967295:
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 uint8 = 0
            for fi39 := 0; fi39 < 20; fi39++ {
                fmt.Println(v6)
                _ = fi39
            }
            _ = fi37
            _ = v38
        }
    }
    var v40 int = v1
    v1 += v40
    var v41 uint32 = 42
    v42 := GenS1{F0: 215, F1: 1}
    v40 += (v1 / ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v8
}
