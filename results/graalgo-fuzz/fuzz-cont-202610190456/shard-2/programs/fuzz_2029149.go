package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint64) uint64 {
    var v6 int = -618
    v7 := []int{v6, v6}
    if false {
        var v8 int8 = 24
        _ = v8
    } else {
        if len(v7) > 0 {
            v7[0] = v6
        }
    }
    v9 := GenS1{F0: 42, F1: 823, F2: 300}
    v10 := GenS2{F0: 0, F1: 244, F2: 969}
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    _ = v9
    return 42
}

func main() {
    var v1 string = ("foo bar" + ("\n\t" + ""))
    v2 := []int{-9223372036854775808, -555, 9223372036854775807}
    var v3 int8 = 42
    v4 := []uint32{104, 65535, 56}
    v5 := []uint32{208}
    var v11 int64 = 282
    if (32 >= -32768) {
        v12 := GenS3{F0: ((v1 < (v1 + ("\n\t" + v1))) || (3.14159 == (-2.25 - 1.5))), F1: 2}
        var v13 int32 = -756
        if true {
            switch (v1 + v1) {
            case "foo bar":
                if true {
                    switch 42 {
                    case 32767:
                        var v14 uint8 = 0
                        _ = v14
                    case 2147483647:
                        var v15 int16 = 0
                        _ = v15
                    default:
                        var v16 uint = 18446744073709551615
                        _ = v16
                    }
                } else {
                    var v17 uint32 = 1
                    _ = v17
                }
            default:
                v18 := GenS2{F0: 436, F1: 0, F2: 344}
                _ = v18
            }
            v1 = v1
            var v19 uint32 = 2
            _ = v19
        } else {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int16 = 1000
                v22 := GenS2{F0: (v21 % ((v21) | 1)), F1: genFn0(18446744073709551615), F2: 2}
                if (1 == 1) {
                    v1 = v1
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = v12
        _ = v13
    }
    var v23 uint32 = 247
    var v24 int32 = int32(v11)
    v25 := GenS4{F0: v1}
    if false {
        switch (v1 + (((v1 + "hello") + (v1 + (v1 + v1))) + (v1 + "\n\t"))) {
        case "a":
            for fi26 := 0; fi26 < 12; fi26++ {
                if true {
                    v5 = append(v5, v23)
                    var v27 int = fi26
                    if len(v2) > 0 {
                        v2[0] = 712
                    }
                    _ = v27
                } else {
                    var v28 string = v1
                    if true {
                        v29 := GenS2{F0: 1, F1: 1, F2: 1}
                        v24 = -2147483648
                        fmt.Println(v28, v11, v1)
                        _ = v29
                    } else {
                        if false {
                            v28 = ""
                        } else {
                            var v30 int16 = 1
                            _ = v30
                        }
                        var v31 int8 = (-(v3))
                        _ = v31
                    }
                    _ = v28
                }
                var v32 float64 = -2.25
                _ = fi26
                _ = v32
            }
        case "foo bar":
            var v33 uint16 = 65535
            _ = v33
        }
        if len(v4) > 0 {
            v4[0] = v23
        }
    }
    v3 += (v3 >> ((0) & 63))
    var v34 float32 = 1.5
    v1 = v1
    var v35 int = -452
    var v36 uint16 = 2
    for fi37 := 0; fi37 < 14; fi37++ {
        fmt.Println(v1)
        for fi38 := 0; fi38 < 19; fi38++ {
            if false {
                fmt.Println(fi38, v1)
            }
            _ = len(v5)
            _ = fi38
        }
        switch v1 {
        case "\n\t":
            var v39 int32 = v24
            _ = v39
        default:
            fmt.Println(v11)
        }
        _ = fi37
    }
    var v40 bool = true
    _ = len(v4)
    for fi41 := 0; fi41 < 17; fi41++ {
        v42 := GenS5{F0: -32768}
        if ((1 < 18446744073709551615) || v40) {
            v40 = (("" == v1) && (!true))
            fmt.Println(fi41, v34, v3)
            for fi43 := 0; fi43 < 9; fi43++ {
                v11 -= (v11 << ((18446744073709551615) & 63))
                _ = fi43
            }
        } else {
            v35 -= fi41
        }
        v44 := GenS5{F0: 40}
        _ = fi41
        _ = v42
        _ = v44
    }
    switch v1 {
    case "hello":
        v34 -= 0.1
    default:
        fmt.Println(v36)
    }
    var v45 uint64 = 70
    v46 := GenS2{F0: 32767, F1: 0, F2: (v36 - v36)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v5
}
