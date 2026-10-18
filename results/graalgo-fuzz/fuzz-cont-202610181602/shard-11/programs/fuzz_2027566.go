package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 int8, p1 int32, p2 int64) uint64 {
    var v5 int = 32768
    p1 -= -755
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 642
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 0
    var v3 float32 = -2.25
    var v4 int = 0
    switch 96 {
    case 32768:
        var v6 uint64 = 568
        _ = v6
    case -889:
        var v7 uint64 = 876
        _ = v7
    default:
        var v8 uint32 = (v2 % (((v2 >> ((4294967295) & 63))) | 1))
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v3 = v3
        switch "" {
        case "\n\t":
            var v10 int32 = -2147483648
            _ = v10
        case "hello":
            fmt.Println(v3, v3)
        case "a":
            var v11 uint8 = 127
            _ = v11
        default:
            var v12 int32 = -2147483648
            _ = v12
        }
        v13 := GenS1{F0: genFn0(-128, 42, 42), F1: 15}
        _ = fi9
        _ = v13
    }
    v14 := []int{v4}
    var v15 int64 = 813
    switch (v4 % ((1) | 1)) {
    case 977:
        if false {
            var v16 uint16 = 65535
            var v17 uint32 = 4294967295
            v15 = (^(v15))
            _ = v16
            _ = v17
        } else {
            v3 -= -2.25
            fmt.Println(v3)
            v15 += (v15 << ((839) & 63))
        }
    case -2147483648:
        v3 += v3
    default:
        v3 += v3
    }
    switch (("foo bar" + ((("\n\t" + "") + "hello") + ("a" + ""))) + ("" + (("\n\t" + ("" + (("hello" + "a") + ("hello" + "foo bar")))) + ("foo bar" + "a")))) {
    case "hello":
        switch v4 {
        case -2147483648:
            var v18 uint32 = v2
            _ = v18
        case 29:
            switch "a" {
            case "foo bar":
                fmt.Println(v2, v4)
            default:
                _ = len(v14)
            }
        }
    case "":
        var v19 bool = ((^(v2)) <= 42)
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 1; fi20++ {
            _ = len(v14)
            var v21 uint64 = 1
            _ = fi20
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 3; fi22++ {
            v14 = append(v14, (fi22 * (^(v4))))
            _ = fi22
        }
    }
    v1 = (v1 | v1)
    var v23 int = (v4 << ((2) & 63))
    if (!false) {
        if len(v14) > 0 {
            v14[0] = v23
        }
    } else {
        if (0 > 128) {
            v23 = v4
            if false {
                var v24 bool = (!true)
                _ = v24
            } else {
                var v25 uint32 = (v2 * v2)
                if true {
                    fmt.Println(v25)
                    v4 += (^(v4))
                }
                _ = v25
            }
        } else {
            if (-9223372036854775808 < v15) {
                switch ("hello" + "foo bar") {
                case "a":
                    if true {
                        if false {
                            var v26 string = ("\n\t" + "a")
                            for fi27 := 0; fi27 < 17; fi27++ {
                                fmt.Println(v2, fi27, fi27)
                                _ = fi27
                            }
                            if len(v14) > 0 {
                                v14[0] = (^(v23))
                            }
                            _ = v26
                        } else {
                            fmt.Println(v2, v3, v1)
                        }
                        for fi28 := 0; fi28 < 5; fi28++ {
                            fmt.Println(fi28, v2)
                            var v29 int32 = -665
                            switch v4 {
                            case 42:
                                var v30 string = (("foo bar" + "foo bar") + ("a" + "hello"))
                                _ = v30
                            case 32767:
                                v31 := GenS2{F0: false, F1: ("hello" + "foo bar"), F2: 565}
                                _ = v31
                            default:
                                fmt.Println(v4, v1, v4)
                            }
                            _ = fi28
                            _ = v29
                        }
                        v14 = append(v14, v23)
                    }
                case "foo bar":
                    var v32 float32 = (v3 * -2.25)
                    _ = v32
                case "hello":
                    var v33 float64 = (1e10 * (((1e10 + ((0.1 + (1.5 + ((1.5 + 1e10) + 0.1))) + 0.1)) * 0.1) + -2.25))
                    _ = v33
                default:
                    var v34 int = 32767
                    _ = v34
                }
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 bool = (true || ("foo bar" <= "foo bar"))
                    fmt.Println(v2, fi35)
                    _ = fi35
                    _ = v36
                }
                _ = len(v14)
            }
        }
    }
    v1 -= v1
    v37 := GenS3{F0: 65535, F1: (v4 & v4), F2: -1}
    v15 += 9223372036854775807
    if len(v14) > 0 {
        v14[0] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v4
}
