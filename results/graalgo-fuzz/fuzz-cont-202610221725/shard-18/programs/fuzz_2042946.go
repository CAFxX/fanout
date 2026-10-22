package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) int64 {
    v4 := GenS1{F0: -9223372036854775808, F1: 1, F2: (-396 < -9223372036854775808)}
    if (false && true) {
        v5 := []int{-652}
        _ = v5
    }
    if ((!true) && (128 < 778)) {
        switch 149 {
        case 286:
            var v6 uint64 = 88
            _ = v6
        case -976:
            if ((32767 <= 32768) || ((!false) && true)) {
                var v7 string = (("\n\t" + "a") + "foo bar")
                if false {
                    p0 = (p0 | 564)
                } else {
                    if true {
                        p0 = p0
                    } else {
                        fmt.Println(p2, v7, p2)
                        v8 := GenS2{F0: (p0 / (((^(p0))) | 1)), F1: uint8(p2)}
                        _ = v8
                    }
                }
                v9 := []int{32768, 120, -482, 9223372036854775807}
                _ = v7
                _ = v9
            } else {
                var v10 int32 = -774
                v11 := GenS3{F0: -1}
                fmt.Println(p2)
                _ = v10
                _ = v11
            }
        case -1:
            p1 -= 127
        }
        switch 0 {
        case 4294967295:
            v12 := GenS1{F0: 9223372036854775807, F1: 0, F2: true}
            _ = v12
        default:
            if (true || false) {
                switch 128 {
                case 4294967295:
                    switch -927 {
                    case -9223372036854775808:
                        v13 := []int{2147483647, int(p1), -9223372036854775808}
                        _ = v13
                    default:
                        v14 := []int{-906}
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint = 0
                    p2 = p2
                    _ = fi15
                    _ = v16
                }
                p0 = 42
            }
        }
    } else {
        if (-1 != 437) {
            v17 := GenS1{F0: 0, F1: 195, F2: false}
            for fi18 := 0; fi18 < 15; fi18++ {
                switch "hello" {
                case "":
                    v19 := []int{fi18, (fi18 & fi18)}
                    _ = v19
                case "a":
                    var v20 int64 = 317
                    _ = v20
                case "hello":
                    var v21 int32 = -1
                    _ = v21
                default:
                    v22 := []int{(-(fi18)), -701}
                    _ = v22
                }
                _ = fi18
            }
            switch (((("\n\t" + "") + "a") + "\n\t") + "foo bar") {
            case "a":
                var v23 float64 = (0.0 * (0.0 - 1e10))
                _ = v23
            case "hello":
                var v24 string = ((("" + ("foo bar" + "\n\t")) + "hello") + ("\n\t" + ""))
                _ = v24
            default:
                var v25 float32 = 1.5
                _ = v25
            }
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -9223372036854775808
}
func genFn1() uint8 {
    var v26 uint32 = 2147483648
    v27 := []int{}
    _ = len(v27)
    v26 = v26
    _ = v26
    _ = v27
    return 196
}

func main() {
    v1 := []int{739, 2147483647}
    var v2 uint64 = 42
    v3 := []string{"\n\t"}
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS4{F0: 56}
        if len(v1) > 0 {
            v1[0] = 0
        }
        _ = fi28
        _ = v29
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    if true {
        v30 := GenS1{F0: 26, F1: 1, F2: true}
        v31 := GenS5{F0: "a", F1: -421, F2: 255}
        _ = v30
        _ = v31
    }
    if true {
        switch "" {
        case "\n\t":
            var v32 string = ""
            _ = v32
        case "hello":
            switch "" {
            case "":
                var v33 int32 = 2147483647
                _ = v33
            case "foo bar":
                if (357 >= 0) {
                    fmt.Println(v2)
                    fmt.Println(v2)
                    var v34 uint8 = 255
                    _ = v34
                } else {
                    v35 := GenS5{F0: "a", F1: 69, F2: 596}
                    var v36 string = ("a" + "hello")
                    var v37 int32 = 442
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
        default:
            var v38 int = 668
            _ = v38
        }
        v2 += (v2 + v2)
        v2 += (^(v2))
    }
    var v39 uint = 225
    var v40 int16 = 0
    v40 += (v40 << ((2) & 63))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
