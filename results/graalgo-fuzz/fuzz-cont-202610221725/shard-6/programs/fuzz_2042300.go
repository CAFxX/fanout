package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint32; F1 int32; F2 int8 }
type GenS7 struct { F0 string; F1 uint16 }
func genFn0(p0 int16) uint64 {
    p0 += -625
    fmt.Println(p0)
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{-1, -1, fi5, 9223372036854775807}
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi7)
            v6 = append(v6, fi7)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    var v8 uint8 = 128
    _ = p0
    _ = v8
    return 4294967295
}
func genFn1(p0 uint64, p1 uint8) uint16 {
    if false {
        fmt.Println(p0, p0)
        var v9 uint16 = uint16(p0)
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int64 = 42
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 float64 = (float64(1e10) / float64(-2.25))
                if true {
                    v14 := GenS1{F0: 1}
                    var v15 float32 = (1e10 / float32(v13))
                    _ = v14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    p0 += (p0 % ((p0) | 1))
    p0 = genFn0(42)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = (false && (false && false))
    v3 := make(map[string]int)
    var v4 uint32 = 2
    v2 = true
    v16 := GenS2{F0: 127}
    var v17 int64 = 1
    v4 = (^(v4))
    var v18 string = "\n\t"
    fmt.Println(v18, v4, v2)
    switch 674 {
    case 127:
        var v19 uint64 = 18446744073709551615
        _ = v19
    case 1:
        fmt.Println(v1)
    case 0:
        delete(v3, (v18 + v18))
    }
    var v20 uint16 = 255
    switch (v18 + v18) {
    case "\n\t":
        v4 = (^(v4))
    case "a":
        var v21 int8 = 42
        _ = v21
    case "":
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                switch ("" + v18) {
                case "hello":
                    var v24 uint = 1
                    _ = v24
                case "\n\t":
                    switch v18 {
                    case "hello":
                        fmt.Println(v2, v20)
                    case "a":
                        switch v18 {
                        case "a":
                            delete(v3, v18)
                        case "hello":
                            v25 := v3[(((v18 + "hello") + (v18 + v18)) + "foo bar")]
                            _ = v25
                        case "\n\t":
                            v26 := GenS2{F0: 100}
                            _ = v26
                        default:
                            v2 = (307 < v17)
                        }
                    case "foo bar":
                        fmt.Println(v1, v1, fi23)
                    }
                }
                _ = fi23
            }
            _ = fi22
        }
    default:
        v27 := v3[v18]
        _ = v27
    }
    if true {
        switch v18 {
        case "":
            if (v4 == (v4 % ((673) | 1))) {
                var v28 uint16 = v1
                v18 = (("\n\t" + (v18 + v18)) + "hello")
                var v29 int32 = -2147483648
                _ = v28
                _ = v29
            } else {
                var v30 uint64 = 872
                if true {
                    v3[("foo bar" + v18)] = int(v4)
                    v31 := GenS3{F0: uint16(v30)}
                    _ = v31
                } else {
                    if v2 {
                        v32 := GenS4{F0: v4}
                        if v2 {
                            for fi33 := 0; fi33 < 12; fi33++ {
                                v34 := GenS5{F0: 42, F1: 128, F2: 48}
                                v1 = v20
                                v35 := []int{(fi33 >> ((uint(v4)) & 63))}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v18, v18, v30)
                            v18 = (v18 + v18)
                        }
                        _ = v32
                    } else {
                        v2 = true
                    }
                    for fi36 := 0; fi36 < 12; fi36++ {
                        for fi37 := 0; fi37 < 18; fi37++ {
                            var v38 uint64 = (v30 ^ (v30 << ((65535) & 63)))
                            var v39 uint64 = (^(v30))
                            var v40 bool = ((^(v39)) != v30)
                            _ = fi37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        v41 := []int{fi36, fi36, (fi36 / ((fi36) | 1)), -9223372036854775808}
                        v1 += v20
                        _ = fi36
                        _ = v41
                    }
                }
                switch 923 {
                case 32768:
                    var v42 float64 = ((-2.25 + float64(1e10)) + ((((1e10 / 3.14159) / float64(((0.1 - 3.14159) / 1.5))) + (0.0 + 0.1)) / 1.5))
                    _ = v42
                default:
                    switch ("a" + (("" + v18) + ((v18 + "a") + v18))) {
                    case "foo bar":
                        fmt.Println(v2)
                    case "hello":
                        if false {
                            for fi43 := 0; fi43 < 20; fi43++ {
                                for fi44 := 0; fi44 < 2; fi44++ {
                                    fmt.Println(v20, v30, v2)
                                    v45 := GenS5{F0: 3, F1: 2, F2: 2}
                                    _ = fi44
                                    _ = v45
                                }
                                var v46 uint64 = v30
                                var v47 int32 = int32(v30)
                                _ = fi43
                                _ = v46
                                _ = v47
                            }
                            v2 = (211 != 42)
                            switch 42 {
                            case -1:
                                fmt.Println(v1, v30)
                            }
                        } else {
                            switch ("" + v18) {
                            case "hello":
                                var v48 float32 = 0.0
                                _ = v48
                            case "":
                                v49 := GenS1{F0: v17}
                                _ = v49
                            }
                        }
                    }
                }
                _ = v30
            }
        case "foo bar":
            var v50 uint8 = uint8(v1)
            _ = v50
        case "\n\t":
            v51 := []int{-821, 151, -584}
            _ = v51
        default:
            fmt.Println(v2, v1)
        }
        var v52 bool = v2
        _ = v52
    } else {
        if (118 <= 1) {
            v53 := GenS5{F0: 1, F1: 1, F2: 2}
            _ = v53
        }
        v54 := GenS6{F0: 2147483648, F1: 2147483647, F2: 0}
        for fi55 := 0; fi55 < 10; fi55++ {
            v18 = (v18 + ("" + v18))
            _ = fi55
        }
        _ = v54
    }
    v56 := GenS7{F0: (("hello" + v18) + "\n\t"), F1: (^(v20))}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v56
}
