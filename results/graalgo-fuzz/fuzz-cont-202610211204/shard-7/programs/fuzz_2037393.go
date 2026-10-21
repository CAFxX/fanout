package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int16 = 42
    var v4 uint64 = 4294967295
    var v5 int16 = v3
    _ = len(v1)
    var v6 string = "hello"
    var v7 uint32 = 65535
    if true {
        v7 = (v7 << ((4294967295) & 63))
        v8 := GenS1{F0: v6, F1: v6, F2: 4294967295}
        _ = v8
    } else {
        v7 = 0
        if true {
            var v9 uint16 = 2
            v4 = v4
            var v10 string = (v6 + "a")
            _ = v10
            _ = v9
        }
        if ((849 >= v4) || true) {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int = fi11
                fmt.Println(v6)
                _ = fi11
                _ = v12
            }
        }
    }
    switch ("" + "foo bar") {
    case "\n\t":
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                v5 -= -32768
                fmt.Println(v3, v6, v7)
                fmt.Println(v4, v3, v4)
            } else {
                v14 := []int{(-(fi13)), (fi13 + fi13), (-(fi13))}
                fmt.Println(v4, v3, v7)
                v7 -= (v7 + 935)
                _ = v14
            }
            var v15 uint64 = 279
            _ = fi13
            _ = v15
        }
    case "foo bar":
        v3 = v5
    default:
        v16 := []int{128, -1, -457, 32768}
        _ = v16
    }
    if true {
        v17 := []int{395}
        _ = v17
    } else {
        if true {
            switch 128 {
            case -399:
                v18 := GenS2{F0: -728, F1: 32767}
                _ = v18
            default:
                if false {
                    fmt.Println(v6, v6, v6)
                    switch -1 {
                    case 127:
                        fmt.Println(v7, v4, v3)
                    default:
                        v6 = v6
                    }
                } else {
                    var v19 uint32 = (v7 | v7)
                    v20 := v1[(v19 / ((v19) | 1))]
                    _ = v19
                    _ = v20
                }
            }
        }
        var v21 int32 = 23
        _ = v21
    }
    var v22 int16 = 1
    v3 += v3
    if ((!false) || true) {
        for fi23 := 0; fi23 < 9; fi23++ {
            if false {
                v24 := []int{-751, (fi23 ^ (fi23 | 4294967295)), (fi23 % ((707) | 1))}
                _ = v24
            } else {
                var v25 bool = (v7 == 70)
                _ = v25
            }
            fmt.Println(v5, v7)
            _ = fi23
        }
        switch 32768 {
        case 32768:
            var v26 int8 = -128
            _ = v26
        case -721:
            v27 := GenS1{F0: "\n\t", F1: ((v6 + ((v6 + (v6 + "a")) + "\n\t")) + (v6 + v6)), F2: uint32(v4)}
            _ = v27
        default:
            var v28 uint32 = v7
            _ = v28
        }
    }
    _ = len(v1)
    delete(v2, v7)
    for fi29 := 0; fi29 < 15; fi29++ {
        if (!false) {
            for fi30 := 0; fi30 < 20; fi30++ {
                var v31 bool = (42 != -706)
                delete(v1, (^(v7)))
                _ = fi30
                _ = v31
            }
        }
        var v32 int64 = int64(v7)
        if (!(v6 != ((v6 + "a") + "hello"))) {
            v33 := GenS2{F0: v32, F1: (-(fi29))}
            var v34 uint8 = 10
            for fi35 := 0; fi35 < 6; fi35++ {
                if true {
                    v36 := GenS2{F0: (v32 - v32), F1: fi35}
                    var v37 float64 = 0.1
                    fmt.Println(fi35, fi29)
                    v37 = 3.14159
                    switch (fi35 >> ((919) & 63)) {
                    case -152:
                        var v38 uint32 = v7
                        _ = v38
                    case -708:
                        fmt.Println(fi35, v34, v37)
                    case 9223372036854775807:
                        if (-9223372036854775808 > (-(v32))) {
                            v39 := GenS2{F0: 9223372036854775807, F1: (fi29 >> ((139) & 63))}
                            if true {
                                v37 = v37
                                var v40 uint8 = 0
                                if (false || false) {
                                    var v41 int32 = 42
                                    _ = v41
                                }
                                _ = v40
                            } else {
                                _ = len(v2)
                            }
                            v42 := []int{}
                            _ = v39
                            _ = v42
                        }
                    }
                    _ = v36
                    _ = v37
                } else {
                    v43 := GenS3{F0: fi29}
                    v32 = (v32 & (^(v32)))
                    var v44 float32 = 0.0
                    _ = v43
                    _ = v44
                }
                var v45 int32 = -2147483648
                delete(v1, v7)
                _ = fi35
                _ = v45
            }
            _ = v33
            _ = v34
        }
        _ = fi29
        _ = v32
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        for fi47 := 0; fi47 < 2; fi47++ {
            var v48 uint8 = 130
            if (!(9223372036854775807 < 9223372036854775807)) {
                v6 = ("hello" + v6)
                v22 += -1
                var v49 int8 = -75
                _ = v49
            } else {
                v4 = (v4 ^ (v4 ^ v4))
            }
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
