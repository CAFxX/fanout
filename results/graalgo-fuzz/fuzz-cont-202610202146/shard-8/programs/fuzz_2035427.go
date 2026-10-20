package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint = 255
    var v2 int = -236
    var v3 uint8 = 127
    v4 := make(map[uint32]string)
    var v5 int16 = 32767
    v6 := GenS1{F0: 42, F1: v3}
    switch (v2 >> ((v1) & 63)) {
    case 9223372036854775807:
        if false {
            v7 := GenS2{F0: (v2 + (v2 % ((96) | 1))), F1: v5, F2: -1}
            if false {
                var v8 float32 = 0.1
                _ = v8
            } else {
                if (475 < 42) {
                    v9 := GenS3{F0: (^(v3)), F1: -9223372036854775808, F2: -1}
                    var v10 uint16 = 0
                    if false {
                        var v11 uint = (v1 * v1)
                        var v12 float32 = -2.25
                        _ = v11
                        _ = v12
                    } else {
                        var v13 uint = (v1 + (v1 | v1))
                        _ = v13
                    }
                    _ = v10
                    _ = v9
                } else {
                    v14 := v4[911]
                    _ = v14
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    _ = len(v4)
                    v16 := []int{}
                    _ = fi15
                    _ = v16
                }
                fmt.Println(v3, v3, v5)
            }
            v2 += 1
            _ = v7
        }
    case 974:
        if false {
            var v17 float32 = 0.0
            switch ("\n\t" + "") {
            case "foo bar":
                var v18 int64 = -346
                _ = v18
            case "\n\t":
                for fi19 := 0; fi19 < 6; fi19++ {
                    v1 += (v1 | v1)
                    for fi20 := 0; fi20 < 14; fi20++ {
                        fmt.Println(v2, v17, v5)
                        v17 -= 1.5
                        fmt.Println(v2, v17)
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v22 := GenS1{F0: -1, F1: (v3 << ((65535) & 63))}
                        switch ("hello" + "\n\t") {
                        case "a":
                            if (false || ((v5 << (((v1 + (v1 << (((v1 / ((65535) | 1))) & 63)))) & 63)) <= v5)) {
                                _ = len(v4)
                                var v23 string = (("a" + "hello") + (((("hello" + "\n\t") + "a") + "a") + "a"))
                                _ = v23
                            }
                        default:
                            var v24 float32 = v17
                            _ = v24
                        }
                        _ = fi21
                        _ = v22
                    }
                    _ = fi19
                }
            case "":
                var v25 int32 = 42
                _ = v25
            }
            if false {
                v26 := GenS2{F0: v2, F1: v5, F2: -1}
                v27 := []int{128, (-(v2)), v2}
                v27 = append(v27, 128)
                _ = v26
                _ = v27
            } else {
                fmt.Println(v1)
            }
            _ = v17
        }
    case 4294967295:
        v1 -= v1
    }
    var v28 float32 = 0.0
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := GenS3{F0: v3, F1: 9223372036854775807, F2: 22}
        _ = fi29
        _ = v30
    }
    if true {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 int8 = 14
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 uint64 = 18446744073709551615
                v35 := v4[42]
                v34 = (v34 | v34)
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi31
            _ = v32
        }
        if true {
            v3 += (v3 >> (((v1 & 502)) & 63))
            v36 := GenS4{F0: (v5 ^ 1)}
            _ = v36
        }
    }
    switch "\n\t" {
    case "":
        v37 := []int{(^(v2)), 17, -1, (v2 * v2)}
        _ = v37
    case "\n\t":
        v38 := []int{v2, v2}
        _ = v38
    case "hello":
        v39 := GenS1{F0: 925, F1: 164}
        _ = v39
    }
    v40 := GenS1{F0: -9223372036854775808, F1: 239}
    for fi41 := 0; fi41 < 20; fi41++ {
        v1 = (v1 + (v1 << (((v1 & 1)) & 63)))
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS5{F0: 65535}
            _ = fi42
            _ = v43
        }
        if ((((v28 * 0.0) / v28) < v28) || true) {
            switch (fi41 << ((65535) & 63)) {
            case 772:
                var v44 int32 = 2147483647
                _ = v44
            case 128:
                v45 := []int{v2, fi41}
                _ = v45
            case 32767:
                var v46 uint32 = 8
                _ = v46
            default:
                fmt.Println(v28, v5, fi41)
            }
        }
        _ = fi41
    }
    delete(v4, 2147483648)
    for fi47 := 0; fi47 < 14; fi47++ {
        if true {
            v48 := GenS1{F0: -1, F1: (v3 * v3)}
            _ = v48
        }
        _ = fi47
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        if true {
            v50 := GenS4{F0: 122}
            v51 := GenS5{F0: 32767}
            if ((((-2.25 / 0.1) / (-2.25 * 3.14159)) - (3.14159 / 1.5)) >= 0.1) {
                v52 := GenS4{F0: -637}
                _ = v52
            }
            _ = v50
            _ = v51
        } else {
            fmt.Println(v28, v5)
        }
        var v53 int32 = 1
        for fi54 := 0; fi54 < 16; fi54++ {
            v55 := v4[65535]
            var v56 uint32 = 1
            fmt.Println(v28, v53)
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi49
        _ = v53
    }
    var v57 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v57
    _ = v6
}
