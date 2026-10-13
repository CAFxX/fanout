package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]int)
    var v3 float64 = 3.14159
    var v4 uint8 = 31
    v5 := make(map[string]bool)
    var v6 uint8 = (v4 << ((18446744073709551615) & 63))
    if (v3 >= v3) {
        var v7 float64 = v3
        var v8 int32 = -1
        _ = v7
        _ = v8
    }
    v9 := []int{4294967295, 1, 9223372036854775807}
    if false {
        var v10 int8 = 127
        _ = v10
    } else {
        if false {
            var v11 int64 = -417
            var v12 uint16 = uint16(v4)
            _ = v11
            _ = v12
        }
        if false {
            v9 = append(v9, 780)
            v9 = append(v9, 42)
        } else {
            v4 -= (v4 << ((412) & 63))
        }
        var v13 bool = false
        _ = v13
    }
    var v14 int = 32767
    for fi15 := 0; fi15 < 11; fi15++ {
        if false {
            v9 = append(v9, fi15)
            if false {
                v9 = append(v9, (fi15 * 42))
                v1 = float32(v3)
                v9 = append(v9, 935)
            } else {
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint16 = 1
                    _ = fi16
                    _ = v17
                }
            }
            var v18 float64 = (v3 / v3)
            _ = v18
        }
        if false {
            if (true && true) {
                var v19 int64 = -9223372036854775808
                fmt.Println(v4, v19, v3)
                v6 = v4
                _ = v19
            } else {
                var v20 uint16 = 65535
                v21 := GenS1{F0: 2, F1: (v6 ^ v6)}
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 int8 = 127
                    fmt.Println(fi15)
                    _ = fi22
                    _ = v23
                }
                _ = v20
                _ = v21
            }
        } else {
            var v24 float32 = (v1 * v1)
            _ = v24
        }
        _ = fi15
    }
    v25 := GenS2{F0: v4}
    var v26 int32 = -2147483648
    v27 := GenS1{F0: 876, F1: v4}
    var v28 float64 = (v3 - (v3 * (-2.25 + (v3 - float64(1.5)))))
    switch (-(v14)) {
    case 127:
        var v29 string = ("hello" + "\n\t")
        _ = v29
    case 490:
        v30 := GenS3{F0: -1}
        _ = v30
    case 198:
        for fi31 := 0; fi31 < 10; fi31++ {
            fmt.Println(v3)
            v14 = v14
            _ = fi31
        }
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v6 -= 128
    default:
        if true {
            for fi32 := 0; fi32 < 19; fi32++ {
                v4 += (v6 ^ 1)
                _ = fi32
            }
            v14 -= v14
        } else {
            for fi33 := 0; fi33 < 20; fi33++ {
                _ = len(v2)
                _ = fi33
            }
            if (!false) {
                var v34 uint32 = 65535
                _ = v34
            } else {
                if (!(!true)) {
                    switch v14 {
                    case 0:
                        v14 = (v14 << ((255) & 63))
                    case -744:
                        fmt.Println(v4, v6)
                    case 2147483647:
                        for fi35 := 0; fi35 < 20; fi35++ {
                            v26 += (v26 / (((v26 - (v26 * (v26 << ((456) & 63))))) | 1))
                            v14 += (v14 / ((fi35) | 1))
                            v26 += (v26 & (^(v26)))
                            _ = fi35
                        }
                    default:
                        v9 = append(v9, (-(v14)))
                    }
                }
                var v36 int32 = (-(v26))
                v4 = v6
                _ = v36
            }
            v3 = -2.25
        }
    }
    v37 := GenS1{F0: 42, F1: 238}
    for fi38 := 0; fi38 < 9; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            v3 -= 0.1
            var v40 uint16 = 32767
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            fmt.Println(v4, v6, v14)
            _ = fi42
        }
        _ = fi41
    }
    v43 := GenS4{F0: -67, F1: 482}
    delete(v2, 2147483648)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v9
}
