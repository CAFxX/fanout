package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int64 }
func genFn0() int64 {
    switch -908 {
    case 2147483647:
        switch -9223372036854775808 {
        case 4294967295:
            if (false && false) {
                v6 := GenS1{F0: 998}
                v7 := GenS1{F0: 42}
                switch 9223372036854775807 {
                case 0:
                    var v8 uint8 = 255
                    _ = v8
                default:
                    var v9 float64 = ((0.0 / -2.25) - 0.1)
                    _ = v9
                }
                _ = v6
                _ = v7
            }
        case -233:
            v10 := GenS2{F0: "foo bar", F1: -221}
            _ = v10
        }
    }
    var v11 int32 = 195
    v12 := GenS1{F0: 65535}
    _ = v11
    _ = v12
    return -542
}
func genFn1() int8 {
    var v13 float64 = 1e10
    if true {
        v13 = v13
    } else {
        v13 += ((v13 - v13) / (((v13 * v13) / (v13 - (((v13 - v13) / -2.25) / 1.5))) / (0.0 - (((v13 - float64((3.14159 - ((-2.25 * 3.14159) - -2.25)))) + 1.5) - ((1.5 - (v13 - -2.25)) + v13)))))
    }
    if true {
        var v14 uint16 = 42
        if true {
            v14 += v14
        }
        var v15 int64 = 0
        _ = v14
        _ = v15
    } else {
        if true {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint16 = 1
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v13, fi16)
                    _ = fi18
                }
                v19 := []int{(^(fi16)), 127, fi16, (fi16 >> ((835) & 63))}
                _ = fi16
                _ = v17
                _ = v19
            }
        } else {
            switch 128 {
            case -9223372036854775808:
                v13 = (3.14159 - float64((((-2.25 * float32((v13 / 1.0))) + 1e10) * 0.0)))
            default:
                var v20 uint8 = 0
                _ = v20
            }
        }
    }
    fmt.Println(v13)
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(fi21, v13)
        var v22 float32 = 1e10
        _ = fi21
        _ = v22
    }
    _ = v13
    return -87
}

func main() {
    var v1 float64 = (float64(0.0) - (-2.25 - float64((1e10 - (0.1 * (float32((1.5 * 1.5)) * (0.0 + -2.25)))))))
    var v2 int = 0
    v3 := []uint32{0, 903}
    var v4 int = (-(v2))
    var v5 int32 = -2147483648
    switch ("" + ("" + "")) {
    case "\n\t":
        v23 := GenS2{F0: "foo bar", F1: v5}
        _ = v23
    case "":
        _ = len(v3)
    default:
        switch v2 {
        case -2147483648:
            v4 -= (v2 >> ((42) & 63))
        default:
            v24 := GenS3{F0: "foo bar", F1: 840}
            _ = v24
        }
    }
    fmt.Println(v2, v5, v1)
    if len(v3) > 0 {
        v3[0] = 674
    }
    v25 := GenS3{F0: "hello", F1: 261}
    if (255 == 42) {
        if true {
            v4 = (v4 << ((0) & 63))
            var v26 uint = 438
            switch (v2 ^ 759) {
            case 62:
                if false {
                    fmt.Println(v26, v26)
                }
            case 32767:
                v5 = 2147483647
            default:
                v27 := GenS3{F0: ("foo bar" + "\n\t"), F1: v26}
                _ = v27
            }
            _ = v26
        }
        var v28 string = ((("" + "hello") + ("a" + "hello")) + (("foo bar" + "\n\t") + "hello"))
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint16 = 530
        switch v2 {
        case 4294967295:
            v31 := GenS3{F0: "a", F1: 128}
            _ = v31
        case -2147483648:
            for fi32 := 0; fi32 < 6; fi32++ {
                if false {
                    if true {
                        v3 = append(v3, 2147483648)
                    }
                    v33 := GenS1{F0: 65535}
                    var v34 float32 = 1e10
                    _ = v33
                    _ = v34
                }
                v3 = append(v3, 4294967295)
                v35 := GenS4{F0: -709}
                _ = fi32
                _ = v35
            }
        }
        _ = fi29
        _ = v30
    }
    v4 = (v4 >> ((168) & 63))
    var v36 int8 = -67
    if true {
        for fi37 := 0; fi37 < 15; fi37++ {
            fmt.Println(v2, v4, v5)
            v1 += ((v1 + v1) + (1e10 / 3.14159))
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = fi37
        }
    }
    v2 += (v4 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
