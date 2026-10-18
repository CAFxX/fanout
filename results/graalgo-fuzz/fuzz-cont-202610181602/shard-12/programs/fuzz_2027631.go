package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 uint; F2 int8 }
type GenS7 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 uint, p1 int32) uint64 {
    var v4 float64 = (0.0 / 0.1)
    if ((42 < 18446744073709551615) && (((((0 >= p1) || false) || false) || ((!false) && true)) || (439 >= 127))) {
        for fi5 := 0; fi5 < 5; fi5++ {
            fmt.Println(p0, p0)
            v6 := []int{fi5, (fi5 * (fi5 % ((fi5) | 1)))}
            p0 += (^(p0))
            _ = fi5
            _ = v6
        }
        var v7 int64 = 9223372036854775807
        v7 = v7
        _ = v7
    }
    var v8 uint = 2
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 18446744073709551615
}
func genFn1() int32 {
    var v9 bool = false
    v10 := GenS1{F0: 2}
    var v11 int16 = 42
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]string)
    var v3 uint8 = 128
    v1 -= v1
    var v12 uint64 = 467
    var v13 int8 = -1
    v12 = genFn0(883, 1)
    v14 := GenS2{F0: 139}
    v15 := GenS2{F0: 65535}
    if true {
        if true {
            var v16 float32 = 3.14159
            var v17 uint32 = 42
            if true {
                v18 := GenS3{F0: 4294967295, F1: -32768, F2: genFn1()}
                v19 := GenS3{F0: 42, F1: 1, F2: -610}
                _ = v18
                _ = v19
            } else {
                if false {
                    fmt.Println(v3)
                    switch 504 {
                    case -951:
                        v17 += v17
                    case 1:
                        var v20 string = "a"
                        _ = v20
                    }
                }
                var v21 uint8 = (v3 ^ (v3 + 42))
                v2["\n\t"] = ("hello" + "a")
                _ = v21
            }
            _ = v16
            _ = v17
        }
    } else {
        v22 := []int{}
        if (42 == (v12 / ((genFn0(42, 603)) | 1))) {
            for fi23 := 0; fi23 < 17; fi23++ {
                v12 -= genFn0(42, -1)
                for fi24 := 0; fi24 < 16; fi24++ {
                    v25 := GenS4{F0: (fi24 * fi24)}
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 5; fi26++ {
                    v27 := v2[("a" + "\n\t")]
                    _ = fi26
                    _ = v27
                }
                _ = fi23
            }
        } else {
            fmt.Println(v1, v13, v12)
            v2["hello"] = ""
            v28 := GenS5{F0: 255}
            _ = v28
        }
        switch "foo bar" {
        case "a":
            fmt.Println(v12, v12)
        case "":
            v3 += (v3 % (((v3 & 25)) | 1))
        case "hello":
            if false {
                switch -5 {
                case -9223372036854775808:
                    v3 = 1
                }
            } else {
                if true {
                    v1 = (v1 & (v1 << ((18446744073709551615) & 63)))
                } else {
                    v29 := GenS6{F0: 42, F1: 847, F2: (v13 << ((224) & 63))}
                    _ = v29
                }
                v13 -= int8(v1)
            }
        default:
            var v30 float64 = 0.1
            _ = v30
        }
        _ = v22
    }
    if false {
        for fi31 := 0; fi31 < 15; fi31++ {
            for fi32 := 0; fi32 < 17; fi32++ {
                v12 -= v12
                if true {
                    v33 := GenS5{F0: v1}
                    v12 -= (v12 | 2)
                    _ = v33
                }
                v34 := GenS3{F0: 255, F1: int16(v12), F2: 594}
                _ = fi32
                _ = v34
            }
            var v35 int32 = -2147483648
            v36 := GenS7{F0: (^(v3)), F1: fi31, F2: (v12 * (v12 | genFn0(18446744073709551615, v35)))}
            _ = fi31
            _ = v35
            _ = v36
        }
        switch "" {
        case "":
            switch 42 {
            case 4294967295:
                for fi37 := 0; fi37 < 1; fi37++ {
                    v1 += 42
                    for fi38 := 0; fi38 < 6; fi38++ {
                        switch fi37 {
                        case 42:
                            v39 := []int{}
                            _ = v39
                        case 1:
                            fmt.Println(v1)
                        default:
                            v40 := GenS2{F0: 2}
                            _ = v40
                        }
                        var v41 int8 = (v13 ^ v13)
                        _ = fi38
                        _ = v41
                    }
                    _ = fi37
                }
            case 458:
                _ = len(v2)
            }
        }
        v42 := []int{906, 9223372036854775807, 9223372036854775807, -460}
        _ = v42
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        var v44 int32 = 42
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
