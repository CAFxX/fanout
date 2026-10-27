package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    p0 = (p0 & (p0 - p0))
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int32 = 161
    v4 := make(map[int]bool)
    var v5 int = -2147483648
    if false {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint32 = 844
            v5 = 128
            _ = fi6
            _ = v7
        }
    } else {
        v8 := GenS1{F0: 0, F1: 119}
        _ = v8
    }
    fmt.Println(v5, v3)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v5, v3)
    case "hello":
        for fi9 := 0; fi9 < 8; fi9++ {
            v3 = v3
            v2[4294967295] = (!(false && false))
            if (42 != 557) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 int64 = 940
                    v3 = (v3 - (v3 * 677))
                    if (1 != 443) {
                        var v12 int16 = 1000
                        _ = v12
                    } else {
                        fmt.Println(v11, v11, v11)
                    }
                    _ = fi10
                    _ = v11
                }
            } else {
                var v13 int16 = -1
                var v14 bool = (false || false)
                _ = v13
                _ = v14
            }
            _ = fi9
        }
    default:
        if (2147483647 == 612) {
            fmt.Println(v3, v5)
            var v15 uint32 = 2
            switch v5 {
            case -828:
                var v16 int32 = 2147483647
                _ = v16
            case 128:
                var v17 bool = false
                _ = v17
            case 32767:
                v5 -= 32768
            }
            _ = v15
        } else {
            var v18 int8 = 43
            var v19 uint64 = 18446744073709551615
            _ = v18
            _ = v19
        }
    }
    var v20 uint32 = 4294967295
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS1{F0: 0, F1: (^(v20))}
        _ = fi21
        _ = v22
    }
    var v23 bool = false
    var v24 int8 = -114
    if true {
        v25 := GenS1{F0: -280, F1: 42}
        _ = v25
    } else {
        v26 := GenS2{F0: 65535}
        _ = v26
    }
    var v27 int8 = 20
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 uint32 = (v20 & v20)
        for fi30 := 0; fi30 < 1; fi30++ {
            v1[v5] = v29
            v3 -= (v3 - (v3 ^ v3))
            v27 = (v27 << ((42) & 63))
            _ = fi30
        }
        v5 = -198
        _ = fi28
        _ = v29
    }
    var v31 int16 = 351
    if (true && (!(!true))) {
        v24 = (v24 ^ 42)
        v3 += v3
    } else {
        v32 := GenS3{F0: "\n\t", F1: -1, F2: 0}
        delete(v4, (v5 >> ((0) & 63)))
        _ = v32
    }
    v23 = v23
    if (v23 || v23) {
        v27 = v24
        if v23 {
            v33 := GenS1{F0: -1, F1: (v20 >> ((1) & 63))}
            v4[958] = false
            var v34 bool = v23
            _ = v33
            _ = v34
        } else {
            v5 = 0
            v24 = (^(v27))
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 float64 = 1e10
                v37 := []int{128, 128, fi35}
                if (!v23) {
                    v24 = -128
                } else {
                    v37 = append(v37, (v5 >> ((0) & 63)))
                    if v23 {
                        for fi38 := 0; fi38 < 12; fi38++ {
                            v39 := GenS2{F0: 2}
                            switch (^(fi35)) {
                            case -9223372036854775808:
                                v27 = -87
                            }
                            delete(v4, (fi35 - fi38))
                            v36 += v36
                            _ = fi38
                            _ = v39
                        }
                        switch "" {
                        case "":
                            var v40 int = v5
                            _ = v40
                        case "\n\t":
                            v41 := GenS2{F0: 65535}
                            _ = v41
                        }
                    } else {
                        v23 = v23
                        var v42 int = (fi35 ^ 2147483647)
                        v31 += v31
                        _ = v42
                    }
                }
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    } else {
        v43 := GenS4{F0: 958}
        var v44 uint32 = v20
        _ = len(v2)
        _ = v43
        _ = v44
    }
    if v23 {
        v45 := GenS5{F0: (v3 % (((^(v3))) | 1))}
        _ = v45
    } else {
        v5 += v5
        v27 += v24
        switch v5 {
        case 361:
            v46 := GenS1{F0: -32768, F1: v20}
            _ = v46
        case -719:
            fmt.Println(v3, v24)
        case -9223372036854775808:
            v47 := []int{v5}
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        v27 = 127
        v24 = v27
        v27 = v27
        _ = fi48
    }
    var v49 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v49
    _ = v5
}
