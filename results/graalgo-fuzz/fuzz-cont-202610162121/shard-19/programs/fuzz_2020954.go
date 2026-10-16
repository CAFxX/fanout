package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int64 }
func genFn0(p0 string) int {
    var v5 uint = 0
    v6 := []int{}
    for fi7 := 0; fi7 < 13; fi7++ {
        if ((false && (414 == 4294967295)) && true) {
            v6 = append(v6, fi7)
            var v8 uint16 = 769
            _ = v8
        } else {
            var v9 uint16 = 255
            if true {
                var v10 int8 = 10
                fmt.Println(p0, fi7)
                v11 := GenS1{F0: 579}
                _ = v10
                _ = v11
            }
            var v12 string = p0
            _ = v12
            _ = v9
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v6 = append(v6, fi13)
        for fi14 := 0; fi14 < 1; fi14++ {
            if (0.1 >= ((-2.25 / 1e10) + (0.1 + 3.14159))) {
                var v15 uint8 = 42
                var v16 uint = (v5 << ((837) & 63))
                fmt.Println(fi14, fi13)
                _ = v15
                _ = v16
            }
            var v17 string = p0
            _ = len(v6)
            v6 = append(v6, fi14)
            _ = fi14
            _ = v17
        }
        if len(v6) > 0 {
            v6[0] = fi13
        }
        _ = fi13
    }
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}
func genFn1(p0 string, p1 int8) string {
    p0 = p0
    v18 := GenS2{F0: 215, F1: 543, F2: 255}
    _ = p0
    _ = p1
    _ = v18
    return p0
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = 1
    var v3 int32 = 1
    var v4 uint32 = 65535
    for fi19 := 0; fi19 < 2; fi19++ {
        v4 = v4
        v3 += v2
        var v20 uint16 = 546
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := []int{958, int(v3), genFn0("a"), (fi21 * -2147483648)}
        fmt.Println(v1, fi21, v1)
        switch "\n\t" {
        case "hello":
            v23 := GenS2{F0: 32767, F1: -387, F2: 235}
            _ = v23
        case "\n\t":
            var v24 uint16 = 255
            _ = v24
        case "foo bar":
            fmt.Println(fi21, v2)
        default:
            v22 = append(v22, (fi21 | (fi21 & genFn0("foo bar"))))
        }
        _ = fi21
        _ = v22
    }
    switch ("hello" + "hello") {
    case "hello":
        var v25 int16 = 0
        _ = v25
    case "":
        if (!(65535 <= 701)) {
            v26 := GenS2{F0: 65535, F1: 1000, F2: 255}
            _ = v26
        } else {
            var v27 int32 = (v3 + (v3 | (v3 * v2)))
            for fi28 := 0; fi28 < 17; fi28++ {
                switch "foo bar" {
                case "hello":
                    v4 = (v4 >> ((4294967295) & 63))
                }
                var v29 uint = 18446744073709551615
                _ = fi28
                _ = v29
            }
            switch "a" {
            case "hello":
                for fi30 := 0; fi30 < 2; fi30++ {
                    for fi31 := 0; fi31 < 4; fi31++ {
                        if ((v1 * v1) != (-2.25 + v1)) {
                            v32 := GenS1{F0: (^(v4))}
                            v4 = (v4 & 42)
                            v33 := []int{(fi31 ^ 0), (fi30 / (((fi31 - -2147483648)) | 1)), (-(fi31))}
                            _ = v32
                            _ = v33
                        }
                        _ = fi31
                    }
                    fmt.Println(v27, v27)
                    _ = fi30
                }
            case "a":
                switch "foo bar" {
                case "a":
                    v27 = v2
                case "\n\t":
                    v34 := GenS1{F0: v4}
                    _ = v34
                case "":
                    if false {
                        var v35 int = 1
                        v36 := GenS3{F0: 9223372036854775807, F1: (v4 / ((v4) | 1)), F2: 660}
                        _ = v35
                        _ = v36
                    } else {
                        var v37 uint8 = 255
                        for fi38 := 0; fi38 < 2; fi38++ {
                            if ((^(v37)) <= v37) {
                                fmt.Println(v37, fi38)
                                v39 := GenS1{F0: v4}
                                var v40 int64 = 9223372036854775807
                                _ = v39
                                _ = v40
                            }
                            for fi41 := 0; fi41 < 20; fi41++ {
                                v42 := []int{genFn0((("a" + "\n\t") + "a")), (fi38 + (fi38 / ((128) | 1))), (fi38 * (fi38 * fi41))}
                                v4 = (v4 % (((v4 >> ((65535) & 63))) | 1))
                                _ = fi41
                                _ = v42
                            }
                            _ = fi38
                        }
                        _ = v37
                    }
                default:
                    v1 = (float64((1.5 - 0.1)) + (0.0 * -2.25))
                }
            }
            _ = v27
        }
    case "foo bar":
        v4 += v4
    default:
        var v43 uint32 = 2147483647
        _ = v43
    }
    if true {
        for fi44 := 0; fi44 < 16; fi44++ {
            v45 := GenS4{F0: -20}
            var v46 bool = (!(float32(v1) != ((1e10 * (0.0 + 1.5)) / 0.1)))
            _ = fi44
            _ = v45
            _ = v46
        }
    } else {
        v47 := GenS5{F0: 9223372036854775807}
        _ = v47
    }
    v3 += (v2 << ((65535) & 63))
    v48 := []int{-1, 2147483647, 408, 32768}
    if false {
        if false {
            var v49 uint64 = 672
            _ = v49
        }
    }
    v50 := GenS6{F0: 233, F1: 2}
    if true {
        if ("\n\t" != "a") {
            for fi51 := 0; fi51 < 9; fi51++ {
                _ = len(v48)
                _ = fi51
            }
            v2 = 2147483647
        }
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v4, v2, v3)
        }
    }
    v3 = 42
    v3 += (^(v2))
    v52 := GenS7{F0: 997}
    v53 := GenS1{F0: (v4 % ((2) | 1))}
    v54 := GenS6{F0: 463, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v50
    _ = v52
    _ = v53
    _ = v54
}
