package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 uint32 = 2
    var v5 bool = (v4 >= v4)
    delete(v2, v4)
    fmt.Println(v4)
    switch "hello" {
    case "hello":
        switch "hello" {
        case "\n\t":
            v6 := GenS1{F0: 65535}
            _ = v6
        case "a":
            var v7 uint64 = 18446744073709551615
            _ = v7
        case "foo bar":
            v8 := []int{}
            _ = v8
        }
    case "\n\t":
        v4 = (v4 | (v4 ^ v4))
    case "":
        v9 := []int{4294967295}
        _ = v9
    }
    var v10 uint64 = 353
    v1 = 0
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint64 = 772
        v13 := GenS2{F0: 1}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 164
    v4 += v4
    v1 += v1
    if true {
        if (-32768 > 1) {
            for fi15 := 0; fi15 < 13; fi15++ {
                v14 = (v14 ^ (v14 % ((v14) | 1)))
                _ = fi15
            }
            fmt.Println(v14, v5, v10)
        }
        var v16 float32 = -2.25
        _ = v16
    } else {
        if v5 {
            var v17 uint = 599
            _ = v17
        } else {
            var v18 float32 = 0.1
            var v19 string = ("hello" + "hello")
            v1 -= v1
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int16 = 42
            v14 -= (v14 << ((216) & 63))
            v22 := []int{(fi20 - fi20), (^(fi20))}
            _ = fi20
            _ = v21
            _ = v22
        }
        if (!v5) {
            var v23 bool = (("foo bar" + "\n\t") > "")
            _ = v23
        } else {
            if true {
                if v5 {
                    v24 := []int{32768, 980}
                    _ = v24
                } else {
                    _ = len(v3)
                }
            }
            v10 -= (^(v10))
            if (65535 > 822) {
                v25 := v2[v4]
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 int = (fi26 & fi26)
                    _ = fi26
                    _ = v27
                }
                var v28 int64 = v1
                _ = v25
                _ = v28
            } else {
                v5 = (!(!v5))
                v29 := GenS1{F0: 42}
                for fi30 := 0; fi30 < 20; fi30++ {
                    if v5 {
                        v14 += (v14 + v14)
                        fmt.Println(v14)
                        var v31 int8 = 100
                        _ = v31
                    } else {
                        v1 = v1
                    }
                    if v5 {
                        for fi32 := 0; fi32 < 20; fi32++ {
                            v4 = v4
                            _ = fi32
                        }
                        for fi33 := 0; fi33 < 15; fi33++ {
                            if v5 {
                                if (v4 != v4) {
                                    if (3.14159 < 1e10) {
                                        v14 -= v14
                                    }
                                } else {
                                    v14 = v14
                                }
                                fmt.Println(v1, fi33)
                                if (!(18446744073709551615 <= 42)) {
                                    v1 = (^(v1))
                                    v4 -= (v4 >> ((18446744073709551615) & 63))
                                }
                                var v34 uint32 = v4
                                _ = v34
                            }
                            switch ("" + ("\n\t" + "hello")) {
                            case "foo bar":
                                switch 4294967295 {
                                case 128:
                                    if (42 >= (v4 | 4294967295)) {
                                        v5 = (v1 <= 360)
                                        v35 := GenS2{F0: -32768}
                                        v4 -= (^(v4))
                                        v10 = v10
                                        _ = v35
                                    } else {
                                        v36 := []int{(fi33 ^ fi33), (fi30 + (fi33 * fi30)), (fi30 | (fi30 ^ (fi30 ^ 641)))}
                                        _ = v36
                                    }
                                default:
                                    v5 = v5
                                }
                            case "a":
                                v5 = v5
                            case "":
                                v37 := []int{(fi30 & 2147483647)}
                                _ = v37
                            }
                            _ = fi33
                        }
                    } else {
                        if (!v5) {
                            v38 := GenS2{F0: 42}
                            _ = v38
                        } else {
                            switch fi30 {
                            case 198:
                                v39 := v2[42]
                                _ = v39
                            }
                            if true {
                                switch -2147483648 {
                                case -9223372036854775808:
                                    var v40 float32 = 1.5
                                    _ = v40
                                default:
                                    v41 := GenS3{F0: 2}
                                    _ = v41
                                }
                                switch "" {
                                case "hello":
                                    var v42 float64 = 1e10
                                    _ = v42
                                case "\n\t":
                                    fmt.Println(v1, v5, v5)
                                case "a":
                                    var v43 int16 = -32768
                                    _ = v43
                                default:
                                    v1 -= int64(v4)
                                }
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    v14 = (v14 + v14)
                                    _ = fi44
                                }
                            }
                            _ = len(v3)
                        }
                    }
                    v45 := GenS3{F0: uint16(v14)}
                    _ = fi30
                    _ = v45
                }
                _ = v29
            }
        }
    }
    if (!v5) {
        v5 = (-545 > 0)
    } else {
        switch "foo bar" {
        case "":
            if v5 {
                var v46 int32 = 1
                _ = v46
            }
        case "\n\t":
            v14 = (v14 % (((v14 >> ((2) & 63))) | 1))
        }
        v14 -= (v14 / (((v14 % (((^(v14))) | 1))) | 1))
    }
    var v47 bool = (true || (0 <= 2147483647))
    var v48 uint64 = v10
    if v47 {
        for fi49 := 0; fi49 < 12; fi49++ {
            if (0.1 == (((float64(((((0.0 / (0.1 / -2.25)) + 3.14159) + 0.1) + float32(1.5))) - 3.14159) / (3.14159 - -2.25)) / 1.5)) {
                if (v4 != v4) {
                    var v50 uint8 = (v14 & (v14 & 128))
                    v50 = v14
                    _ = v50
                }
                var v51 uint = 1
                _ = v51
            }
            v52 := GenS2{F0: -32768}
            _ = fi49
            _ = v52
        }
    }
    var v53 bool = ((-32768 >= -32768) && (v5 || v47))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v53
}
