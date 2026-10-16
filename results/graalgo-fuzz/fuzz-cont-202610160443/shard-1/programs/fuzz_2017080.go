package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint64 }

func main() {
    v1 := []string{"hello", ("" + "a"), (("" + "") + "\n\t")}
    var v2 int32 = -928
    var v3 string = "a"
    v4 := make(map[string]string)
    var v5 uint64 = 349
    if true {
        fmt.Println(v5)
        v5 = 347
        var v6 string = v3
        _ = v6
    }
    _ = len(v1)
    var v7 int = 430
    v8 := GenS1{F0: uint(v5)}
    if ((!(!true)) && false) {
        var v9 uint = 441
        var v10 int8 = -128
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 = (v3 + v3)
        v2 += (v2 * v2)
        if (1e10 != (-2.25 / ((0.1 + (float32(((0.1 / 1.0) - ((0.1 + ((float64(float32(0.1)) - 3.14159) - 1e10)) / 0.1))) / 1.5)) - 3.14159))) {
            v2 -= (v2 << ((65535) & 63))
            var v12 uint32 = 308
            v7 += fi11
            _ = v12
        } else {
            v13 := GenS2{F0: 109}
            _ = v13
        }
        _ = fi11
    }
    var v14 bool = false
    var v15 string = "a"
    v5 = 2
    var v16 int16 = 42
    if ((98 <= 32767) || (!(2147483647 > v7))) {
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                if (0 >= -1) {
                    if v14 {
                        v5 = v5
                        v19 := GenS3{F0: v7}
                        var v20 uint = 2
                        _ = v19
                        _ = v20
                    }
                    if v14 {
                        if v14 {
                            v21 := GenS4{F0: 962}
                            var v22 float32 = 1.5
                            v2 = v2
                            _ = v21
                            _ = v22
                        }
                        if (!false) {
                            v14 = v14
                        } else {
                            var v23 bool = v14
                            v24 := GenS1{F0: 988}
                            _ = v23
                            _ = v24
                        }
                        if true {
                            fmt.Println(v14, v14)
                            v16 += v16
                            switch (v7 - fi17) {
                            case 745:
                                v7 = v7
                            case 40:
                                if (v14 || v14) {
                                    switch (fi17 & (v7 % ((458) | 1))) {
                                    case -888:
                                        v25 := v4[""]
                                        _ = v25
                                    case 0:
                                        switch v15 {
                                        case "hello":
                                            v14 = true
                                        }
                                        var v26 bool = false
                                        var v27 uint32 = 2147483648
                                        _ = v26
                                        _ = v27
                                    case 128:
                                        _ = len(v1)
                                        var v28 string = "hello"
                                        _ = v28
                                    }
                                } else {
                                    v29 := GenS5{F0: (^(v7)), F1: (^(v5))}
                                    _ = v29
                                }
                            case 1:
                                v30 := GenS4{F0: 2147483648}
                                _ = v30
                            default:
                                var v31 bool = v14
                                _ = v31
                            }
                        }
                        var v32 int16 = v16
                        _ = v32
                    } else {
                        var v33 uint8 = 255
                        switch (fi18 ^ fi18) {
                        case 416:
                            fmt.Println(v7)
                        case 333:
                            _ = len(v1)
                        case 2147483647:
                            var v34 bool = ((v5 % ((v5) | 1)) < 2)
                            _ = v34
                        default:
                            var v35 uint16 = uint16(v5)
                            _ = v35
                        }
                        _ = v33
                    }
                }
                var v36 float64 = float64(((0.0 + ((float32(0.1) / 3.14159) * -2.25)) * 1e10))
                fmt.Println(v3)
                _ = fi18
                _ = v36
            }
            _ = fi17
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            if false {
                if ((32767 != 739) && v14) {
                    var v38 uint32 = 525
                    fmt.Println(v38)
                    _ = v38
                } else {
                    v1 = append(v1, v15)
                    v39 := GenS4{F0: 4294967295}
                    _ = v39
                }
                var v40 uint8 = 42
                var v41 float64 = ((float64(((float32(3.14159) + 0.0) * (0.0 / 1e10))) - 0.0) - 0.0)
                _ = v40
                _ = v41
            } else {
                v14 = (0 != 985)
                v14 = ((v16 * int16(v2)) >= 1)
                for fi42 := 0; fi42 < 7; fi42++ {
                    if (fi37 <= 128) {
                        switch (fi37 | fi37) {
                        case -1:
                            v7 += (^(fi37))
                        }
                        v43 := GenS2{F0: 18}
                        _ = v43
                    }
                    var v44 int64 = -364
                    _ = fi42
                    _ = v44
                }
            }
            for fi45 := 0; fi45 < 11; fi45++ {
                v46 := v4[v3]
                _ = fi45
                _ = v46
            }
            var v47 int64 = int64(v5)
            _ = fi37
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        var v49 int8 = -72
        _ = fi48
        _ = v49
    }
    fmt.Println(v15)
    for fi50 := 0; fi50 < 5; fi50++ {
        v51 := GenS1{F0: 42}
        v15 = "a"
        if v14 {
            v52 := GenS2{F0: 5}
            fmt.Println(v5)
            for fi53 := 0; fi53 < 9; fi53++ {
                v54 := GenS3{F0: fi53}
                v4[v3] = (v3 + "\n\t")
                _ = len(v1)
                _ = fi53
                _ = v54
            }
            _ = v52
        } else {
            v55 := v4["hello"]
            fmt.Println(v55)
            _ = v55
        }
        _ = fi50
        _ = v51
    }
    v56 := v4[v15]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v7
    _ = v8
}
