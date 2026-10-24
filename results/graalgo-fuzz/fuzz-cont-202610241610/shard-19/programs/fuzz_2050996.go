package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64 }

func main() {
    var v1 float64 = -2.25
    var v2 bool = (1 != 1000)
    var v3 string = "a"
    var v4 float64 = (1e10 * v1)
    v5 := []int{-952, 739, -1, 1}
    switch -2147483648 {
    case -83:
        var v6 int64 = 42
        _ = v6
    case 42:
        _ = len(v5)
    }
    v7 := GenS1{F0: 948, F1: 255, F2: 18}
    var v8 int16 = 0
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = 0
    }
    switch 32767 {
    case 540:
        var v9 int32 = -711
        _ = v9
    case 128:
        v10 := GenS1{F0: 128, F1: 846, F2: 255}
        _ = v10
    case -2147483648:
        v4 += v1
    default:
        v1 += ((((float64(1e10) * v4) + v1) / v1) - 3.14159)
    }
    if (((1 >= 32767) || (!v2)) || false) {
        v4 += float64(-2.25)
        for fi11 := 0; fi11 < 10; fi11++ {
            if false {
                var v12 uint = 655
                switch 9223372036854775807 {
                case 4294967295:
                    if (!true) {
                        var v13 uint = v12
                        if (-1 <= 1) {
                            v14 := GenS1{F0: -1, F1: 2, F2: 1}
                            v2 = ((!v2) || (!v2))
                            _ = len(v5)
                            _ = v14
                        } else {
                            var v15 int16 = (v8 & 549)
                            _ = v15
                        }
                        var v16 int = (-(fi11))
                        _ = v13
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v18 := GenS1{F0: (-(fi11)), F1: 42, F2: 182}
                            _ = fi17
                            _ = v18
                        }
                    }
                case -1:
                    v3 = (v3 + (v3 + "foo bar"))
                case -212:
                    v19 := GenS1{F0: fi11, F1: 255, F2: 109}
                    _ = v19
                }
                v12 -= 494
                _ = v12
            }
            if false {
                var v20 bool = (202 != 42)
                var v21 uint64 = 191
                _ = v20
                _ = v21
            }
            _ = fi11
        }
    } else {
        v8 = -32768
    }
    var v22 int64 = 9223372036854775807
    v8 = (v8 << ((424) & 63))
    _ = len(v5)
    var v23 int16 = 32767
    if v2 {
        if true {
            var v24 uint64 = 347
            fmt.Println(v8, v23, v2)
            _ = v24
        } else {
            _ = len(v5)
        }
        switch v3 {
        case "a":
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 int32 = 215
                v5 = append(v5, fi25)
                _ = fi25
                _ = v26
            }
        case "":
            var v27 uint32 = 2
            _ = v27
        case "\n\t":
            v28 := GenS1{F0: 42, F1: 1, F2: 42}
            _ = v28
        }
    }
    if v2 {
        for fi29 := 0; fi29 < 1; fi29++ {
            if (!false) {
                switch fi29 {
                case -9223372036854775808:
                    if len(v5) > 0 {
                        v5[0] = (fi29 | fi29)
                    }
                }
                if (!((128 < 4294967295) || v2)) {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        if ((-(v23)) <= v8) {
                            var v31 int16 = v8
                            v2 = (1.5 == -2.25)
                            _ = v31
                        }
                        var v32 uint16 = 0
                        v3 = v3
                        _ = fi30
                        _ = v32
                    }
                    switch v3 {
                    case "foo bar":
                        v23 += v23
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v2 = v2
                    case "foo bar":
                        var v33 string = (v3 + (v3 + ""))
                        _ = v33
                    }
                }
                v34 := GenS2{F0: 2147483647, F1: v22}
                _ = v34
            } else {
                fmt.Println(v1)
                _ = len(v5)
            }
            _ = fi29
        }
        fmt.Println(v3)
        v35 := GenS2{F0: 457, F1: (v22 ^ v22)}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
