package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 uint32) int8 {
    var v5 int8 = 127
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            p0 = -2147483648
            _ = fi7
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{}
            fmt.Println(v5, fi6, v5)
            _ = fi8
            _ = v9
        }
        if ((("a" + "foo bar") > "\n\t") && true) {
            for fi10 := 0; fi10 < 16; fi10++ {
                if (!(false && (!false))) {
                    if false {
                        if (true && true) {
                            v11 := GenS1{F0: 0, F1: 754}
                            _ = v11
                        }
                        switch "\n\t" {
                        case "":
                            v12 := GenS2{F0: 0}
                            p0 = (p0 / ((42) | 1))
                            _ = v12
                        case "foo bar":
                            v5 = v5
                        default:
                            var v13 float32 = ((0.1 * 1.5) / float32((1e10 + -2.25)))
                            _ = v13
                        }
                    } else {
                        v14 := []int{4294967295, fi10, fi10}
                        v14 = append(v14, fi10)
                        if true {
                            var v15 int64 = -1
                            _ = v15
                        }
                        _ = v14
                    }
                }
                v5 = (v5 ^ -1)
                _ = fi10
            }
            var v16 uint16 = 734
            _ = v16
        }
        _ = fi6
    }
    var v17 int32 = p0
    p0 += 1
    _ = p0
    _ = p1
    _ = v17
    _ = v5
    return v5
}

func main() {
    var v1 uint = 2
    var v2 float32 = -2.25
    v3 := []int{1}
    v4 := make(map[int]int)
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "\n\t" {
        case "hello":
            _ = len(v3)
        case "a":
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS2{F0: 42}
                _ = fi19
                _ = v20
            }
        case "foo bar":
            v1 = 2
        default:
            switch (-(fi18)) {
            case 0:
                for fi21 := 0; fi21 < 19; fi21++ {
                    v1 = (v1 / (((v1 % (((^(v1))) | 1))) | 1))
                    _ = fi21
                }
            case 32768:
                var v22 uint16 = 979
                _ = v22
            }
        }
        var v23 bool = (false || true)
        _ = fi18
        _ = v23
    }
    fmt.Println(v2)
    var v24 int32 = 623
    for fi25 := 0; fi25 < 15; fi25++ {
        if (42 == 1) {
            fmt.Println(v2)
        }
        var v26 int = fi25
        for fi27 := 0; fi27 < 9; fi27++ {
            switch "" {
            case "foo bar":
                var v28 int64 = -9223372036854775808
                _ = v28
            default:
                _ = len(v3)
            }
            _ = len(v4)
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    if (true && (!false)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v29 := GenS3{F0: 32767, F1: 32768}
    v24 = v24
    fmt.Println(v2, v24, v1)
    if (-173 == 1000) {
        var v30 int8 = 1
        _ = v30
    } else {
        var v31 int16 = 32767
        _ = v31
    }
    if true {
        v1 -= (v1 | (v1 << ((v1) & 63)))
        v24 = -230
        var v32 string = ("a" + ("" + "foo bar"))
        _ = v32
    } else {
        v2 = ((float32(1.5) / ((1e10 / ((((v2 - (v2 + v2)) / v2) + (float32(float64(v2)) + 1.5)) + 0.1)) + -2.25)) / (v2 * v2))
        for fi33 := 0; fi33 < 14; fi33++ {
            fmt.Println(v1)
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
}
