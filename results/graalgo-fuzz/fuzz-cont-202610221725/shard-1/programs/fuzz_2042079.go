package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int) int8 {
    if false {
        p0 = (p0 | p0)
    } else {
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 uint64 = 2
            switch "hello" {
            case "\n\t":
                v6 := GenS1{F0: p0, F1: 100}
                _ = v6
            }
            var v7 int32 = 143
            _ = fi4
            _ = v5
            _ = v7
        }
        v8 := []int{(^(p0)), 127, p0}
        _ = v8
    }
    v9 := []int{(p0 << ((1) & 63)), (p0 << ((140) & 63)), 2147483647}
    var v10 int32 = 42
    p0 -= 42
    var v11 string = "hello"
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 1
    var v3 int8 = -1
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: fi12, F1: -97}
        _ = fi12
        _ = v13
    }
    v3 = (v3 & 102)
    v1 -= 1
    var v14 uint32 = 1
    if false {
        var v15 int16 = 42
        if false {
            var v16 bool = (((^(v15)) >= v15) && (!false))
            v17 := GenS2{F0: v3, F1: v2}
            _ = v16
            _ = v17
        } else {
            v18 := []int{9223372036854775807, 341, -715}
            for fi19 := 0; fi19 < 10; fi19++ {
                fmt.Println(v15, v1)
                _ = fi19
            }
            _ = v18
        }
        _ = v15
    }
    if (!(!(true || (0.0 >= 0.0)))) {
        var v20 int32 = -2147483648
        _ = v20
    }
    if ((((v3 | (-(v3))) >= v3) || true) || false) {
        for fi21 := 0; fi21 < 15; fi21++ {
            v3 += -115
            v22 := GenS3{F0: 1, F1: 72}
            _ = fi21
            _ = v22
        }
        v1 -= v1
        if (genFn0(-792) != (v3 << ((v2) & 63))) {
            if (!true) {
                switch (("\n\t" + "hello") + "a") {
                case "\n\t":
                    if false {
                        var v23 uint = 65535
                        if ((v14 / ((uint32(v2)) | 1)) == 2147483647) {
                            v1 += (v1 & v1)
                            var v24 uint = (v23 * v23)
                            v25 := GenS2{F0: 100, F1: 184}
                            _ = v24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        fmt.Println(v14, v14, v3)
                    }
                case "a":
                    var v26 int = 1
                    _ = v26
                default:
                    var v27 int32 = 42
                    _ = v27
                }
                v28 := GenS4{F0: uint8(v2), F1: 39}
                if (!(("\n\t" + "a") < ("" + ("\n\t" + "hello")))) {
                    var v29 int8 = (v3 / ((v3) | 1))
                    var v30 int8 = (v29 - v3)
                    _ = v29
                    _ = v30
                } else {
                    v14 = v14
                }
                _ = v28
            } else {
                switch (("" + "hello") + "\n\t") {
                case "hello":
                    fmt.Println(v14, v2, v14)
                default:
                    fmt.Println(v1, v2)
                }
                fmt.Println(v1, v1, v1)
                if (v3 > v3) {
                    var v31 bool = (0.1 < (0.0 * ((0.0 * -2.25) + 3.14159)))
                    _ = v31
                }
            }
        } else {
            if false {
                fmt.Println(v2)
                for fi32 := 0; fi32 < 10; fi32++ {
                    var v33 int8 = 0
                    v34 := []int{fi32, (fi32 ^ fi32), (fi32 % ((607) | 1)), fi32}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v1 = (-(v1))
            } else {
                v35 := []int{1, -1}
                var v36 bool = false
                _ = v35
                _ = v36
            }
            if false {
                if false {
                    v3 = (v3 >> ((v2) & 63))
                } else {
                    var v37 int64 = (v1 >> ((4294967295) & 63))
                    _ = v37
                }
                v38 := []int{-378}
                _ = v38
            } else {
                v39 := []int{42}
                _ = v39
            }
        }
    } else {
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := GenS3{F0: 4294967295, F1: uint8(v2)}
            v42 := GenS4{F0: 2, F1: (v14 >> (((^(v2))) & 63))}
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    v43 := GenS3{F0: 18446744073709551615, F1: 107}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v43
}
