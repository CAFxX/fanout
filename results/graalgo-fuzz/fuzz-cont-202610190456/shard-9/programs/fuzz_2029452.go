package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]uint32)
    var v3 uint32 = 2
    var v4 float64 = ((0.0 / 0.1) * -2.25)
    delete(v2, 2147483647)
    var v5 int64 = 0
    v6 := GenS1{F0: (("hello" + ("hello" + (("hello" + "\n\t") + (("foo bar" + ("a" + "foo bar")) + ("foo bar" + "\n\t"))))) + "\n\t"), F1: 42}
    if ((float32((v4 * v4)) >= (float32((float64((0.0 - 3.14159)) - (v4 / (-2.25 * v4)))) * 3.14159)) && (((v5 >= 0) && false) && false)) {
        var v7 float64 = v4
        if (-1 > 283) {
            var v8 uint8 = 2
            v9 := GenS2{F0: 42}
            _ = v8
            _ = v9
        }
        v10 := GenS3{F0: 42}
        _ = v10
        _ = v7
    } else {
        var v11 string = "hello"
        fmt.Println(v1)
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        _ = len(v2)
        _ = fi12
    }
    var v13 string = "hello"
    switch "foo bar" {
    case "a":
        var v14 bool = (v5 <= v5)
        _ = v14
    }
    var v15 bool = (94 < 1)
    switch (-(v1)) {
    case 891:
        v15 = v15
    case -710:
        switch 1 {
        case -215:
            var v16 float32 = 3.14159
            _ = v16
        case 692:
            switch v13 {
            case "foo bar":
                v5 = (v5 * v5)
            case "":
                v17 := GenS4{F0: 65535, F1: -1}
                _ = v17
            default:
                if true {
                    switch (v1 << ((255) & 63)) {
                    case -9223372036854775808:
                        fmt.Println(v1, v5)
                    case 127:
                        v18 := []int{}
                        _ = v18
                    default:
                        v15 = v15
                    }
                    for fi19 := 0; fi19 < 8; fi19++ {
                        switch 4294967295 {
                        case 9223372036854775807:
                            if (-659 != 0) {
                                var v20 uint8 = 127
                                if v15 {
                                    v13 = v13
                                    if ((!(v13 <= "foo bar")) && (128 > v20)) {
                                        if v15 {
                                            switch "\n\t" {
                                            case "a":
                                                var v21 int32 = -20
                                                _ = v21
                                            default:
                                                v22 := []int{(-(fi19)), (v1 - fi19), int(v3), (v1 | (^(v1)))}
                                                _ = v22
                                            }
                                            for fi23 := 0; fi23 < 20; fi23++ {
                                                v4 = v4
                                                v5 -= v5
                                                _ = fi23
                                            }
                                        } else {
                                            v24 := GenS1{F0: "a", F1: -656}
                                            var v25 bool = (false && v15)
                                            var v26 uint8 = (v20 / (((v20 - 182)) | 1))
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                    fmt.Println(v20, v5, v4)
                                } else {
                                    var v27 uint16 = 478
                                    fmt.Println(v20, v20, fi19)
                                    _ = v27
                                }
                                var v28 uint8 = v20
                                _ = v20
                                _ = v28
                            }
                        case 978:
                            for fi29 := 0; fi29 < 15; fi29++ {
                                var v30 uint64 = 1
                                fmt.Println(v4)
                                _ = fi29
                                _ = v30
                            }
                        case 206:
                            fmt.Println(v15, v5)
                        }
                        fmt.Println(v15, v15)
                        _ = fi19
                    }
                } else {
                    v3 -= (v3 * v3)
                    v2[2147483647] = v3
                }
            }
        case 42:
            v31 := []int{v1, (-(v1)), (^(v1)), 28}
            _ = v31
        default:
            v4 -= v4
        }
    case 709:
        v32 := GenS5{F0: 609}
        _ = v32
    }
    var v33 int8 = -40
    var v34 uint8 = 2
    var v35 bool = v15
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
