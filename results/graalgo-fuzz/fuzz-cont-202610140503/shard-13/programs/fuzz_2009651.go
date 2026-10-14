package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 int8) int64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{-9223372036854775808, fi6}
        p1 -= (^(p1))
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int8) bool {
    p0 -= (p0 % ((p0) | 1))
    if (-2147483648 < -1) {
        if true {
            if false {
                v8 := GenS1{F0: (p0 + p0), F1: 9223372036854775807}
                p0 = p0
                _ = v8
            } else {
                var v9 int8 = p0
                _ = v9
            }
        }
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint = 102
        for fi12 := 0; fi12 < 19; fi12++ {
            if ((p0 >> (((v11 / ((v11) | 1))) & 63)) >= (p0 << ((4294967295) & 63))) {
                if ((fi10 | fi10) < (fi12 / ((-167) | 1))) {
                    if true {
                        v11 = v11
                        p0 += (p0 / ((-128) | 1))
                        fmt.Println(fi10, fi12, fi12)
                        fmt.Println(p0, v11, p0)
                    } else {
                        var v13 int = (fi12 >> (((v11 % ((1) | 1))) & 63))
                        _ = v13
                    }
                    fmt.Println(fi10)
                    fmt.Println(p0)
                } else {
                    v14 := GenS1{F0: p0, F1: -35}
                    _ = v14
                }
            }
            if true {
                v11 += v11
                var v15 uint32 = 2147483647
                _ = v15
            } else {
                if ((378 == 1) && true) {
                    v16 := []int{42}
                    _ = len(v16)
                    _ = v16
                }
                var v17 bool = ((v11 > (v11 << ((276) & 63))) && ((-1 > 849) && true))
                _ = v17
            }
            if (((32 <= 1) || ((v11 >> (((v11 >> ((212) & 63))) & 63)) >= (v11 - 807))) || (1 > 2)) {
                var v18 int = -402
                var v19 uint = 2
                _ = v18
                _ = v19
            }
            _ = fi12
        }
        v20 := GenS2{F0: p0}
        _ = fi10
        _ = v11
        _ = v20
    }
    v21 := []int{}
    var v22 int8 = p0
    _ = p0
    _ = v21
    _ = v22
    return (false || (127 <= 997))
}

func main() {
    var v1 uint64 = 42
    var v2 bool = (true || true)
    var v3 int8 = 59
    var v4 uint64 = v1
    var v5 uint32 = 65535
    v2 = true
    var v23 int64 = -396
    v24 := []int{775}
    v4 = (v4 & (v4 << ((0) & 63)))
    var v25 float64 = 1.5
    var v26 uint16 = 0
    var v27 uint16 = v26
    v28 := GenS1{F0: (v3 << ((1) & 63)), F1: v23}
    v24 = append(v24, 596)
    if v2 {
        switch 128 {
        case 4294967295:
            v1 = (^(v4))
        case 42:
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 uint32 = v5
                _ = fi29
                _ = v30
            }
        case 9223372036854775807:
            v31 := GenS2{F0: 42}
            _ = v31
        default:
            if false {
                for fi32 := 0; fi32 < 19; fi32++ {
                    v3 -= -1
                    v27 = v27
                    _ = fi32
                }
                v33 := GenS1{F0: v3, F1: v23}
                for fi34 := 0; fi34 < 18; fi34++ {
                    v4 = (v1 - v4)
                    for fi35 := 0; fi35 < 20; fi35++ {
                        var v36 uint32 = v5
                        fmt.Println(v5, v5)
                        if false {
                            v24 = append(v24, (-(fi34)))
                            if ((v26 < v26) && (!(v2 || v2))) {
                                v36 = (v5 >> ((65535) & 63))
                                v4 = 42
                                v4 += (v1 & 664)
                            } else {
                                v3 = 82
                                v5 += (v36 ^ (v36 % (((^(v5))) | 1)))
                                var v37 int32 = -1
                                _ = v37
                            }
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                        }
                        _ = fi35
                        _ = v36
                    }
                    _ = fi34
                }
                _ = v33
            }
        }
    }
    if v2 {
        v2 = true
        var v38 int64 = v23
        _ = v38
    } else {
        v2 = true
        v1 += (^(v1))
        switch "" {
        case "foo bar":
            for fi39 := 0; fi39 < 20; fi39++ {
                fmt.Println(v3, v3, v25)
                v23 = 42
                v24 = append(v24, fi39)
                _ = fi39
            }
        }
    }
    switch 0 {
    case 619:
        var v40 uint16 = (v26 ^ 0)
        _ = v40
    case 574:
        if v2 {
            v2 = (!(1e10 != v25))
            switch -2147483648 {
            case 997:
                _ = len(v24)
            case -1:
                var v41 int64 = (^(v23))
                _ = v41
            }
            v2 = true
        }
    default:
        v42 := GenS3{F0: (1e10 == (0.0 * (1.5 * ((float32(v25) - -2.25) * 3.14159))))}
        _ = v42
    }
    var v43 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
