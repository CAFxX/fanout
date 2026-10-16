package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint16) uint64 {
    fmt.Println(p0, p0)
    fmt.Println(p0, p1)
    var v4 uint8 = 155
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint32, p1 uint) int64 {
    fmt.Println(p1, p1)
    var v5 uint64 = 711
    var v6 int = 9223372036854775807
    for fi7 := 0; fi7 < 5; fi7++ {
        switch ("a" + "\n\t") {
        case "\n\t":
            p1 += p1
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 string = ("foo bar" + ("hello" + "foo bar"))
            switch (((v9 + (("foo bar" + "hello") + v9)) + "") + (v9 + "hello")) {
            case "a":
                switch (v9 + "hello") {
                case "hello":
                    fmt.Println(v5, p0)
                default:
                    v5 = 42
                    if (fi7 < (v6 % ((0) | 1))) {
                        if false {
                            fmt.Println(p0, fi8, v9)
                            if false {
                                p1 -= (^(p1))
                                v10 := []int{(-(fi8)), (fi7 / ((fi8) | 1)), fi7}
                                v9 = (v9 + v9)
                                _ = v10
                            } else {
                                v5 += v5
                                v11 := []int{}
                                var v12 uint64 = v5
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                _ = v11
                                _ = v12
                            }
                            v13 := []int{(fi8 % (((fi8 / ((42) | 1))) | 1)), fi8}
                            _ = v13
                        }
                    } else {
                        v14 := []int{(^(v6)), (-(fi7)), 2147483647}
                        if false {
                            var v15 int8 = 42
                            v5 = v5
                            _ = v15
                        } else {
                            var v16 uint64 = (v5 << ((p1) & 63))
                            v17 := GenS1{F0: 0, F1: 1, F2: 42}
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                    }
                    v18 := []int{4294967295, fi8, v6, (v6 + (fi7 - (fi8 + (-(fi7)))))}
                    _ = v18
                }
            case "\n\t":
                if (0 >= 997) {
                    fmt.Println(p0, p0, fi7)
                    if true {
                        v19 := GenS1{F0: 255, F1: 1, F2: p1}
                        _ = v19
                    }
                }
            }
            v20 := GenS2{F0: true, F1: 128}
            _ = fi8
            _ = v20
            _ = v9
        }
        _ = fi7
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -918
}

func main() {
    var v1 uint64 = 0
    var v2 uint32 = uint32(v1)
    var v3 string = ""
    v2 = (v2 & (v2 ^ 42))
    v21 := []int{}
    var v22 int16 = 32767
    fmt.Println(v3, v3, v1)
    if false {
        var v23 int64 = 0
        v23 -= (v23 - v23)
        _ = v23
    }
    v2 = (v2 & v2)
    fmt.Println(v22, v2, v22)
    if len(v21) > 0 {
        v21[0] = 32768
    }
    fmt.Println(v2, v2, v2)
    v1 += v1
    fmt.Println(v3, v1, v3)
    for fi24 := 0; fi24 < 11; fi24++ {
        v1 = v1
        for fi25 := 0; fi25 < 2; fi25++ {
            switch (fi25 + (fi25 << ((65535) & 63))) {
            case -347:
                if len(v21) > 0 {
                    v21[0] = 32768
                }
            case 32767:
                var v26 float64 = (((1e10 / 3.14159) - float64(float32(-2.25))) - ((3.14159 + 1e10) + 0.1))
                _ = v26
            default:
                v21 = append(v21, 9223372036854775807)
            }
            var v27 int16 = (v22 + (v22 & v22))
            fmt.Println(v27)
            _ = fi25
            _ = v27
        }
        _ = fi24
    }
    v28 := GenS3{F0: 18446744073709551615}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v29 := GenS2{F0: true, F1: 70}
    switch -2147483648 {
    case 128:
        for fi30 := 0; fi30 < 18; fi30++ {
            v1 = (v1 ^ (v1 << ((uint(v1)) & 63)))
            _ = fi30
        }
    default:
        var v31 int16 = v22
        _ = v31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int8 = -15
        switch -60 {
        case -865:
            var v34 uint16 = 259
            _ = v34
        case 1:
            var v35 int8 = v33
            _ = v35
        case 128:
            v36 := GenS2{F0: false, F1: 128}
            _ = v36
        }
        v37 := GenS3{F0: v1}
        _ = fi32
        _ = v33
        _ = v37
    }
    v38 := GenS2{F0: true, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v38
}
