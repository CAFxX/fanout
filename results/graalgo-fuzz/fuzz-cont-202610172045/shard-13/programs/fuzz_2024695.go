package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
type GenS4 struct { F0 int; F1 int16 }

func main() {
    var v1 uint32 = 1
    var v2 uint32 = v1
    var v3 bool = (false || ((!false) || false))
    var v4 int = 32768
    var v5 float32 = (3.14159 * (0.1 - 1e10))
    fmt.Println(v5, v3)
    switch (v4 >> ((18446744073709551615) & 63)) {
    case 1:
        if (!v3) {
            var v6 uint32 = v2
            v3 = v3
            switch "a" {
            case "\n\t":
                v7 := GenS1{F0: 4294967295, F1: v2}
                _ = v7
            case "hello":
                v8 := []int{v4, (v4 * v4)}
                _ = v8
            default:
                switch ("a" + (("\n\t" + (("\n\t" + "hello") + ((("foo bar" + ("hello" + "hello")) + "\n\t") + ("a" + "foo bar")))) + ("a" + "hello"))) {
                case "foo bar":
                    fmt.Println(v3, v6, v6)
                default:
                    var v9 float64 = 0.1
                    _ = v9
                }
            }
            _ = v6
        } else {
            fmt.Println(v1, v2, v2)
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 212, F1: (v2 ^ v1)}
                if (879 > 32767) {
                    v12 := []int{fi10, v4, -206, (fi10 | -9223372036854775808)}
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
        }
    case -1:
        var v13 float32 = (((float32(float64(v5)) - v5) - v5) - (float32(1e10) / (0.0 / v5)))
        _ = v13
    }
    if (true || false) {
        v14 := GenS1{F0: 1, F1: (v1 / ((v2) | 1))}
        _ = v14
    }
    var v15 string = ("foo bar" + "\n\t")
    var v16 string = ((v15 + ("a" + "")) + "foo bar")
    switch v16 {
    case "":
        v3 = (4294967295 > 169)
    case "hello":
        switch 9223372036854775807 {
        case 32768:
            var v17 int = -1
            _ = v17
        case -868:
            v4 = v4
        case -9223372036854775808:
            switch ((v16 + v15) + v16) {
            case "foo bar":
                v3 = (v2 > (v2 << ((795) & 63)))
            case "a":
                if v3 {
                    v1 -= (v2 % (((^(v1))) | 1))
                    if true {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            switch v15 {
                            case "foo bar":
                                var v19 uint8 = 0
                                _ = v19
                            default:
                                var v20 uint = 0
                                _ = v20
                            }
                            v2 += 1
                            v5 = v5
                            _ = fi18
                        }
                        if (211 != 312) {
                            v1 = (v2 * (v2 | v2))
                            v3 = v3
                            var v21 float64 = 0.0
                            _ = v21
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        if false {
                            switch (fi22 % (((v4 / ((v4) | 1))) | 1)) {
                            case 32767:
                                v23 := []int{fi22, (^(v4)), (fi22 | (v4 >> ((18446744073709551615) & 63)))}
                                _ = v23
                            }
                            fmt.Println(v15)
                        } else {
                            v24 := GenS2{F0: 32767}
                            var v25 string = v16
                            v26 := []int{fi22, v4, -649, 287}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        _ = fi22
                    }
                    var v27 uint8 = uint8(v2)
                    _ = v27
                }
            }
        }
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS1{F0: 2, F1: (v1 ^ 4294967295)}
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS3{F0: 375, F1: -2147483648, F2: -1}
    v31 := []int{v4, -400}
    fmt.Println(v5, v1)
    v2 -= (^(v1))
    var v32 uint8 = 128
    v33 := GenS4{F0: -967, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
