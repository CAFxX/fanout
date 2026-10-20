package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
func genFn0(p0 string) uint64 {
    if true {
        p0 = ((p0 + (p0 + p0)) + "\n\t")
    } else {
        v5 := GenS1{F0: -1, F1: p0, F2: 0}
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        if false {
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := []int{}
                var v9 int32 = 42
                var v10 int16 = 0
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            fmt.Println(p0, p0)
        }
        _ = fi6
    }
    _ = p0
    return 0
}
func genFn1() uint8 {
    v11 := []int{-2147483648}
    if true {
        var v12 int32 = 42
        _ = v12
    } else {
        var v13 uint64 = 2
        _ = v13
    }
    var v14 float64 = (-2.25 / 1.0)
    var v15 float32 = 1.5
    _ = v11
    _ = v14
    _ = v15
    return 128
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 599
    var v3 bool = (0 >= 56)
    var v4 float32 = 1.5
    v4 = 0.0
    fmt.Println(v1)
    v16 := GenS1{F0: -9223372036854775808, F1: "hello", F2: 42}
    switch "" {
    case "a":
        v3 = v3
    case "hello":
        v17 := GenS2{F0: 45, F1: (v2 << ((18446744073709551615) & 63)), F2: -128}
        _ = v17
    default:
        fmt.Println(v2, v1)
    }
    v18 := []int{127, 4294967295, 42, 0}
    v3 = true
    switch ("\n\t" + (("a" + "hello") + "foo bar")) {
    case "\n\t":
        if false {
            fmt.Println(v2, v4, v3)
            var v19 float32 = ((float32((0.1 * (-2.25 - (float64(v4) - 3.14159)))) * -2.25) - v4)
            if true {
                switch -585 {
                case -2147483648:
                    if (((((^(v2)) < 0) && (-9223372036854775808 > -1)) && (-1 != 9223372036854775807)) || ((255 > 95) || false)) {
                        var v20 int8 = 127
                        v3 = true
                        _ = len(v18)
                        _ = v20
                    }
                default:
                    v4 = ((float32(0.0) / float32(0.1)) + v19)
                }
            }
            _ = v19
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                for fi22 := 0; fi22 < 2; fi22++ {
                    if (!(-107 > 1)) {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    }
                    _ = fi22
                }
                var v23 int32 = -740
                if v3 {
                    v24 := GenS1{F0: -1, F1: ("\n\t" + "\n\t"), F2: 1}
                    _ = v24
                } else {
                    v25 := GenS3{F0: (^(v2)), F1: 127, F2: 261}
                    var v26 int8 = -128
                    v18 = append(v18, fi21)
                    _ = v25
                    _ = v26
                }
                _ = fi21
                _ = v23
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 uint8 = 42
                _ = fi27
                _ = v28
            }
            fmt.Println(v2, v2, v3)
        }
    default:
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    var v29 float32 = float32(3.14159)
    var v30 float32 = -2.25
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS3{F0: genFn0(("" + ((("\n\t" + "hello") + "a") + ("foo bar" + "a")))), F1: 0, F2: 123}
        v33 := GenS3{F0: genFn0("\n\t"), F1: 127, F2: 0}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v4, v29, v29)
    if true {
        if len(v18) > 0 {
            v18[0] = 32768
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
