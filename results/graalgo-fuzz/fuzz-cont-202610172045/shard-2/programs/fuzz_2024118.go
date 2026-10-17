package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint, p1 bool) int16 {
    var v6 uint8 = 1
    v7 := GenS1{F0: -94}
    var v8 bool = (p0 < 0)
    v8 = (true || (2 >= 18446744073709551615))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 908
}
func genFn1(p0 uint8, p1 uint32, p2 uint16) int64 {
    var v9 uint8 = p0
    switch 430 {
    case 4294967295:
        v10 := []int{4294967295, 0, 127}
        _ = v10
    case 694:
        v9 -= 127
    default:
        switch "a" {
        case "":
            v11 := GenS2{F0: -1, F1: int(p1), F2: (p2 / ((p2) | 1))}
            _ = v11
        default:
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 float64 = (0.0 * -2.25)
                var v14 uint16 = p2
                var v15 bool = true
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        }
    }
    var v16 float64 = 0.0
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v9
    return 933
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (!(0 == 42))
    v3 := []uint32{2147483647, 664, 4294967295, 26}
    var v4 uint8 = 67
    var v5 int64 = -705
    v17 := GenS1{F0: 9223372036854775807}
    if (v2 || (18446744073709551615 != 18446744073709551615)) {
        if (2 < 298) {
            v3 = append(v3, 101)
            var v18 uint64 = 42
            _ = v18
        } else {
            v4 -= 43
            v5 = -1
        }
        var v19 float32 = 1e10
        v3 = append(v3, 1)
        _ = v19
    } else {
        if ((v4 << ((36) & 63)) <= v4) {
            v20 := v1[(("" + "\n\t") + ((("\n\t" + "a") + "a") + ("a" + ("foo bar" + ("" + "")))))]
            _ = v20
        }
    }
    switch "a" {
    case "foo bar":
        switch int(v4) {
        case -9223372036854775808:
            if (1e10 != float32(0.1)) {
                v1[""] = 896
                if (59 != 1) {
                    _ = len(v3)
                    var v21 int32 = 91
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v1)
                    case "a":
                        for fi22 := 0; fi22 < 6; fi22++ {
                            var v23 uint16 = 32767
                            var v24 uint32 = 65535
                            var v25 float32 = -2.25
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    case "foo bar":
                        if false {
                            v26 := GenS3{F0: 123, F1: int64(v21), F2: 65535}
                            fmt.Println(v2, v2, v21)
                            fmt.Println(v4, v21)
                            _ = v26
                        } else {
                            v4 = v4
                        }
                    default:
                        v2 = v2
                    }
                    _ = v21
                } else {
                    var v27 string = "hello"
                    _ = v27
                }
                v5 = (-(v5))
            }
        case 32767:
            switch 127 {
            case 32767:
                v2 = v2
            default:
                v28 := GenS1{F0: 127}
                _ = v28
            }
        case 0:
            fmt.Println(v2, v5)
        default:
            fmt.Println(v4, v5)
        }
    case "\n\t":
        v29 := GenS4{F0: 981, F1: -1}
        _ = v29
    case "hello":
        v1[("" + "")] = 2147483648
    default:
        switch ("" + "foo bar") {
        case "\n\t":
            for fi30 := 0; fi30 < 1; fi30++ {
                for fi31 := 0; fi31 < 16; fi31++ {
                    var v32 uint = 255
                    _ = fi31
                    _ = v32
                }
                if v2 {
                    var v33 uint = 972
                    var v34 bool = true
                    _ = v33
                    _ = v34
                }
                _ = fi30
            }
        case "a":
            switch 128 {
            case -372:
                if (("" + "hello") <= (((((("hello" + "hello") + "foo bar") + ("hello" + "a")) + "\n\t") + "hello") + ("foo bar" + ("a" + ("hello" + "foo bar"))))) {
                    v35 := GenS1{F0: 4294967295}
                    v2 = (3.14159 != (3.14159 - (0.0 - 3.14159)))
                    _ = v35
                }
            default:
                switch ("a" + "\n\t") {
                case "\n\t":
                    var v36 int16 = -712
                    _ = v36
                case "":
                    v37 := GenS3{F0: 18446744073709551615, F1: genFn1((v4 ^ v4), 42, 1), F2: 32767}
                    _ = v37
                }
            }
        case "hello":
            for fi38 := 0; fi38 < 13; fi38++ {
                v4 += (v4 & (v4 & (v4 % ((127) | 1))))
                v5 += (v5 * (v5 / (((v5 | (v5 | -958))) | 1)))
                v2 = (!false)
                _ = fi38
            }
        default:
            if true {
                for fi39 := 0; fi39 < 17; fi39++ {
                    v40 := GenS3{F0: 0, F1: v5, F2: 2}
                    v3 = append(v3, 2)
                    _ = fi39
                    _ = v40
                }
                v2 = (731 != -1)
            }
        }
    }
    if (2147483648 == 255) {
        v3 = append(v3, 2147483647)
        v2 = ((v5 ^ (v5 / (((v5 & v5)) | 1))) < v5)
        var v41 bool = false
        _ = v41
    }
    v4 += 128
    for fi42 := 0; fi42 < 1; fi42++ {
        v4 = v4
        v4 += 101
        if true {
            var v43 uint = 862
            v4 -= (^(v4))
            _ = v43
        } else {
            var v44 uint16 = 32767
            _ = v44
        }
        _ = fi42
    }
    var v45 uint64 = 42
    v5 = (-(v5))
    fmt.Println(v4, v5)
    delete(v1, "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
