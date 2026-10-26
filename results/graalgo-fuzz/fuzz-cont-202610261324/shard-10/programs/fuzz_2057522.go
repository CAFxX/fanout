package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint8 }
func genFn0() bool {
    v5 := []int{206, 970, 1, -471}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint32 = 65535
        if false {
            for fi8 := 0; fi8 < 6; fi8++ {
                _ = len(v5)
                var v9 int = 42
                v10 := GenS1{F0: 149, F1: 2}
                var v11 int32 = -861
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            if false {
                var v12 int32 = 0
                v12 = (v12 << ((42) & 63))
                var v13 float32 = 0.0
                _ = v12
                _ = v13
            }
        }
        switch -1 {
        case 2147483647:
            var v14 uint16 = 2
            _ = v14
        case -1:
            v15 := GenS2{F0: (v7 | 65535)}
            _ = v15
        case -2147483648:
            if true {
                for fi16 := 0; fi16 < 7; fi16++ {
                    if false {
                        if false {
                            v5 = append(v5, fi6)
                        }
                        v7 += (v7 & 65535)
                    } else {
                        v17 := GenS1{F0: 255, F1: (v7 + 2147483648)}
                        _ = v17
                    }
                    _ = fi16
                }
                var v18 bool = false
                fmt.Println(v18, v18)
                _ = v18
            } else {
                _ = len(v5)
                v7 = v7
            }
        default:
            v19 := GenS2{F0: (v7 >> ((797) & 63))}
            _ = v19
        }
        _ = fi6
        _ = v7
    }
    var v20 bool = (42 != 127)
    v21 := GenS2{F0: 2}
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 uint32 = 255
        fmt.Println(v23, v20)
        _ = fi22
        _ = v23
    }
    _ = v20
    _ = v21
    _ = v5
    return false
}
func genFn1(p0 uint, p1 bool) int8 {
    v24 := GenS1{F0: p0, F1: 2147483648}
    p0 = 2
    v25 := []int{-1, -439, 42, 915}
    if (!false) {
        var v26 float32 = -2.25
        var v27 uint = 2
        _ = v26
        _ = v27
    }
    p1 = (!p1)
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    return -128
}

func main() {
    var v1 string = ""
    var v2 float32 = -2.25
    var v3 uint8 = 2
    var v4 uint64 = 18446744073709551615
    var v28 int = 745
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(fi29, v28, v4)
        fmt.Println(v4)
        _ = fi29
    }
    fmt.Println(v2, v3)
    v30 := GenS1{F0: 0, F1: 1}
    var v31 uint32 = 2147483647
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 int32 = 0
        _ = fi32
        _ = v33
    }
    v3 = (v3 / ((uint8(v4)) | 1))
    switch v28 {
    case -927:
        for fi34 := 0; fi34 < 1; fi34++ {
            fmt.Println(v31, v31, v31)
            if (42 == uint(v3)) {
                var v35 int32 = -513
                _ = v35
            } else {
                if ((!(!false)) && (v3 <= (v3 >> ((995) & 63)))) {
                    v36 := GenS3{F0: (v31 + v31), F1: (v4 - 2), F2: (v3 + 32)}
                    if true {
                        switch (v28 >> ((0) & 63)) {
                        case 562:
                            fmt.Println(fi34, v28, v3)
                        default:
                            var v37 bool = false
                            _ = v37
                        }
                        var v38 uint16 = 815
                        v39 := []int{}
                        _ = v38
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 18; fi40++ {
                            switch v28 {
                            case -514:
                                v31 -= v31
                            default:
                                fmt.Println(v3)
                            }
                            fmt.Println(v28)
                            _ = fi40
                        }
                        if false {
                            var v41 int16 = 920
                            v42 := []int{}
                            _ = v41
                            _ = v42
                        } else {
                            var v43 int16 = 1000
                            var v44 uint32 = uint32(v4)
                            _ = v43
                            _ = v44
                        }
                    }
                    _ = v36
                }
                v45 := []int{fi34, (fi34 ^ fi34)}
                _ = v45
            }
            var v46 int32 = 0
            _ = fi34
            _ = v46
        }
    }
    var v47 int64 = 809
    switch v1 {
    case "\n\t":
        v48 := []int{}
        _ = v48
    default:
        var v49 float32 = 1e10
        _ = v49
    }
    v4 = 747
    if ((0 < 1000) || true) {
        var v50 uint = 596
        fmt.Println(v2)
        _ = v50
    }
    switch (v28 / ((v28) | 1)) {
    case 528:
        fmt.Println(v2, v31)
    case 127:
        v4 -= v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v47
}
