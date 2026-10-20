package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 int16, p2 uint32) int32 {
    if (!(-2.25 != (-2.25 + 3.14159))) {
        v6 := []int{-361, -2147483648, 1}
        p2 += 1
        _ = v6
    } else {
        if true {
            v7 := []int{32768}
            if (-840 == 88) {
                _ = len(v7)
                var v8 float64 = 0.1
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := GenS1{F0: "\n\t", F1: (fi9 << ((65535) & 63))}
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
            _ = v7
        } else {
            var v11 int32 = -178
            p0 = (-(p0))
            _ = v11
        }
        p2 -= p2
        v12 := []int{}
        _ = v12
    }
    var v13 uint32 = p2
    v13 -= 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -2147483648
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]string)
    var v3 uint64 = 0
    var v4 int8 = 22
    var v5 uint = 147
    var v14 uint8 = 42
    v3 -= (^(v1))
    fmt.Println(v14, v3)
    v15 := GenS1{F0: "foo bar", F1: 1}
    v4 -= 33
    if ("\n\t" == ("a" + "a")) {
        v14 = v14
    }
    switch 42 {
    case 128:
        var v16 uint16 = 32767
        _ = v16
    default:
        var v17 uint64 = v3
        _ = v17
    }
    v1 += (v3 >> (((v5 & v5)) & 63))
    v18 := []int{}
    v19 := GenS1{F0: "", F1: 387}
    var v20 float64 = -2.25
    v2["foo bar"] = ("\n\t" + ("a" + "hello"))
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS1{F0: ("foo bar" + (("hello" + ("hello" + (("\n\t" + "foo bar") + ("\n\t" + ("hello" + "\n\t"))))) + "\n\t")), F1: (-(fi21))}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        if (true || (2 >= 0)) {
            if (true || (-1 != 2147483647)) {
                fmt.Println(v3)
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                var v24 uint8 = (v14 * (v14 ^ uint8(v5)))
                _ = v24
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 int64 = 1
                _ = fi25
                _ = v26
            }
            v20 -= v20
        } else {
            v1 += uint64(v5)
        }
        v4 = -1
        var v27 uint32 = 42
        _ = fi23
        _ = v27
    }
    var v28 int8 = 59
    switch 42 {
    case 42:
        if (v14 < v14) {
            _ = len(v18)
            v1 -= v1
            switch "\n\t" {
            case "":
                if len(v18) > 0 {
                    v18[0] = 32767
                }
            case "hello":
                var v29 int32 = -1
                _ = v29
            case "\n\t":
                v30 := GenS2{F0: 661}
                _ = v30
            default:
                for fi31 := 0; fi31 < 17; fi31++ {
                    fmt.Println(v1, v4)
                    if (uint32(v3) >= 2147483648) {
                        fmt.Println(fi31, v28, v14)
                    } else {
                        v32 := GenS2{F0: uint16(v14)}
                        if true {
                            var v33 int64 = 1
                            _ = v33
                        }
                        _ = v32
                    }
                    var v34 int64 = 752
                    _ = fi31
                    _ = v34
                }
            }
        }
    case 9223372036854775807:
        if false {
            var v35 int = 414
            _ = v35
        }
    }
    var v36 bool = true
    if v36 {
        for fi37 := 0; fi37 < 18; fi37++ {
            v36 = v36
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
