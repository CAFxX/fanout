package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16; F1 string }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7)
        v8 := GenS1{F0: ("\n\t" + "foo bar")}
        v9 := GenS2{F0: 65535, F1: 2, F2: 65535}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-9223372036854775808, 128, 9223372036854775807, -9223372036854775808}
    _ = v10
    return 0
}
func genFn1(p0 int16, p1 int64) int8 {
    for fi11 := 0; fi11 < 19; fi11++ {
        p0 += p0
        v12 := []int{fi11, 2147483647}
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: genFn0(), F1: 128, F2: 0}
    v14 := []int{-321, 9223372036854775807, 2147483647, 4294967295}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 string = ("a" + "a")
    v3 := make(map[int]uint32)
    var v4 uint8 = 1
    var v5 string = (v2 + ("hello" + v2))
    var v6 int16 = -651
    fmt.Println(v2, v6, v6)
    v15 := []int{-786}
    if true {
        fmt.Println(v5)
        if true {
            switch (v5 + v5) {
            case "":
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 uint32 = 2
                    if (2 == 42) {
                        v2 = v2
                    } else {
                        v4 -= (^(v4))
                        if true {
                            if false {
                                if false {
                                    var v18 int = -1
                                    _ = v18
                                } else {
                                    for fi19 := 0; fi19 < 1; fi19++ {
                                        var v20 int64 = 9223372036854775807
                                        var v21 uint32 = (v17 % (((v17 >> ((42) & 63))) | 1))
                                        _ = fi19
                                        _ = v20
                                        _ = v21
                                    }
                                }
                            }
                            v22 := GenS2{F0: (v17 % ((4294967295) | 1)), F1: v1, F2: 32767}
                            switch v2 {
                            case "foo bar":
                                v23 := GenS2{F0: 2147483648, F1: 233, F2: 65535}
                                _ = v23
                            default:
                                v24 := GenS1{F0: "foo bar"}
                                _ = v24
                            }
                            _ = v22
                        } else {
                            var v25 float32 = 3.14159
                            _ = v25
                        }
                    }
                    _ = fi16
                    _ = v17
                }
            case "a":
                v1 = (v4 & (v4 << ((1) & 63)))
            default:
                var v26 float32 = -2.25
                _ = v26
            }
            v15 = append(v15, 0)
        }
    }
    v5 = (("a" + (v5 + ("a" + (v2 + v5)))) + ("hello" + "\n\t"))
    switch 32767 {
    case 9223372036854775807:
        v6 = 1000
    default:
        v27 := GenS3{F0: v6, F1: ("hello" + v5)}
        _ = v27
    }
    var v28 uint64 = 0
    v2 = v2
    v4 = v1
    if len(v15) > 0 {
        v15[0] = 32767
    }
    v29 := GenS1{F0: ((v2 + v2) + v2)}
    fmt.Println(v5, v6, v5)
    v28 += v28
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
