package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }
type GenS4 struct { F0 int64; F1 int; F2 int64 }
func genFn0(p0 int, p1 int32) int8 {
    v4 := []int{}
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 bool = (false && true)
        fmt.Println(v6, fi5, p0)
        _ = fi5
        _ = v6
    }
    fmt.Println(p1, p0)
    p1 += 2147483647
    if false {
        var v7 int8 = -55
        if false {
            p0 = p0
        }
        fmt.Println(p0, v7)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            p1 -= p1
            v9 := GenS1{F0: fi8, F1: 42, F2: "a"}
            if len(v4) > 0 {
                v4[0] = fi8
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(p0, p0)
        v4 = append(v4, (p0 & 128))
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}
func genFn1() uint16 {
    var v10 uint64 = 2
    v11 := GenS1{F0: 1, F1: 2, F2: "a"}
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 214
    var v3 uint64 = 18446744073709551615
    var v12 uint = 718
    if false {
        var v13 bool = (65535 < 2)
        _ = v13
    }
    if true {
        var v14 uint8 = 182
        for fi15 := 0; fi15 < 13; fi15++ {
            if (-119 == 2147483647) {
                if false {
                    switch fi15 {
                    case 220:
                        var v16 uint64 = v3
                        _ = v16
                    case -152:
                        fmt.Println(v1, v1, v2)
                    case -1:
                        v3 = (v3 ^ (v3 << (((^(v12))) & 63)))
                    }
                    var v17 int64 = -1
                    _ = v17
                } else {
                    var v18 bool = (-53 == -842)
                    _ = v18
                }
            } else {
                v3 = v3
                v19 := []int{}
                var v20 float32 = 1e10
                _ = v19
                _ = v20
            }
            v21 := GenS1{F0: -2147483648, F1: v2, F2: "\n\t"}
            if (v2 > (v2 ^ v2)) {
                if (2147483648 <= v2) {
                    v22 := GenS1{F0: (fi15 - (fi15 % (((-(fi15))) | 1))), F1: v2, F2: "a"}
                    _ = v22
                } else {
                    fmt.Println(fi15, v3)
                }
            }
            _ = fi15
            _ = v21
        }
        var v23 string = ("foo bar" + "\n\t")
        _ = v14
        _ = v23
    }
    v24 := []int{-895, 127, 952, -1}
    v25 := GenS2{F0: (v2 ^ 180)}
    fmt.Println(v1)
    var v26 uint64 = 659
    if true {
        if (true || false) {
            switch -400 {
            case 9223372036854775807:
                if (32767 != uint16(v26)) {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v2 += v2
                        v28 := GenS3{F0: "a", F1: 111, F2: -9223372036854775808}
                        _ = fi27
                        _ = v28
                    }
                } else {
                    var v29 uint8 = 254
                    v1 -= 0.1
                    var v30 bool = true
                    _ = v29
                    _ = v30
                }
            case 1:
                var v31 string = ((("foo bar" + "foo bar") + ("hello" + "\n\t")) + "\n\t")
                _ = v31
            case 239:
                for fi32 := 0; fi32 < 5; fi32++ {
                    fmt.Println(v3, v2, v1)
                    _ = fi32
                }
            default:
                switch 4294967295 {
                case -1:
                    if true {
                        var v33 uint16 = genFn1()
                        _ = v33
                    }
                case 0:
                    v2 = v2
                default:
                    if true {
                        v1 = (float64(-2.25) + 0.0)
                        var v34 uint = (v12 % ((uint(v3)) | 1))
                        _ = v34
                    } else {
                        var v35 uint = (v12 << ((2) & 63))
                        switch 848 {
                        case 128:
                            v35 = 623
                        case 1:
                            var v36 int8 = 100
                            _ = v36
                        }
                        _ = v35
                    }
                }
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                _ = len(v24)
                _ = fi37
            }
            _ = len(v24)
        } else {
            v38 := GenS3{F0: "a", F1: 127, F2: -9223372036854775808}
            _ = v38
        }
    } else {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    }
    v39 := GenS4{F0: -1, F1: -2147483648, F2: 1}
    if false {
        var v40 int16 = 702
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v39
}
