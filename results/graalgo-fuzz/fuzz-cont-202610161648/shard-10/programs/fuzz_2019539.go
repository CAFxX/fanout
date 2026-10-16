package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
func genFn0(p0 uint) uint {
    p0 += 922
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(p0, p0)
        p0 = (^(p0))
        _ = fi6
    }
    if (32767 == 32767) {
        v7 := GenS1{F0: ""}
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{fi8}
            var v10 float64 = (0.1 - ((1.5 * 3.14159) * float64(((-2.25 * float32(float64(((1e10 / 1e10) + -2.25)))) * (((1.5 + -2.25) - (1e10 + (((float32(1e10) + 1.5) * float32(1.5)) * 0.1))) + 3.14159)))))
            if false {
                if true {
                    v11 := GenS1{F0: "a"}
                    _ = v11
                }
                switch "foo bar" {
                case "":
                    var v12 uint8 = 255
                    _ = v12
                case "hello":
                    v9 = append(v9, fi8)
                case "foo bar":
                    v13 := GenS2{F0: 0}
                    _ = v13
                }
                var v14 uint32 = 2147483648
                _ = v14
            } else {
                switch (fi8 + (fi8 | (fi8 * int(p0)))) {
                case 127:
                    v15 := GenS1{F0: "hello"}
                    _ = v15
                default:
                    var v16 int8 = 127
                    _ = v16
                }
                var v17 uint32 = 183
                v17 = (v17 * 2147483647)
                _ = v17
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v18 := []int{}
        _ = v18
        _ = v7
    }
    _ = p0
    return (p0 - 255)
}
func genFn1(p0 uint32, p1 string) bool {
    fmt.Println(p0, p1)
    if false {
        switch p1 {
        case "foo bar":
            fmt.Println(p1)
        case "":
            v19 := []int{42, 99, -2147483648, 32768}
            _ = v19
        default:
            p1 = (p1 + "")
        }
        switch "\n\t" {
        case "":
            v20 := []int{}
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 7; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int = 4294967295
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        default:
            if false {
                switch (("\n\t" + ("" + "a")) + ("foo bar" + ("hello" + p1))) {
                case "foo bar":
                    fmt.Println(p0, p1)
                default:
                    p0 -= p0
                }
            }
        }
    }
    p0 -= (p0 * (p0 | p0))
    _ = p0
    _ = p1
    return false
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 42
    v3 := make(map[string]int)
    v4 := []int{}
    var v5 bool = false
    var v24 uint64 = 4294967295
    if (v5 && v5) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch 4294967295 {
    case 1:
        fmt.Println(v5, v24)
    }
    var v25 int16 = 32767
    v24 = (v24 | v24)
    switch "" {
    case "":
        var v26 int64 = int64(v24)
        _ = v26
    case "\n\t":
        v25 -= v25
    case "foo bar":
        fmt.Println(v2)
    }
    fmt.Println(v5, v24)
    fmt.Println(v24, v2)
    if v5 {
        _ = len(v1)
        for fi27 := 0; fi27 < 5; fi27++ {
            fmt.Println(v24, v2)
            var v28 uint64 = (v24 - v24)
            var v29 int = fi27
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v25 = v25
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := v3["hello"]
        var v32 uint16 = 798
        v2 -= 130
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 uint16 = 2
    v24 = 0
    fmt.Println(v33, v5, v24)
    var v34 uint64 = (^(v24))
    v33 -= 710
    switch ("hello" + "foo bar") {
    case "":
        v34 = (v24 - 0)
    case "hello":
        for fi35 := 0; fi35 < 2; fi35++ {
            v36 := v1[(fi35 / ((fi35) | 1))]
            _ = fi35
            _ = v36
        }
    default:
        v37 := v1[2147483647]
        _ = v37
    }
    switch 32767 {
    case -9223372036854775808:
        if false {
            v1[42] = 707
            fmt.Println(v5, v24, v5)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case -140:
        var v38 uint32 = 642
        _ = v38
    case 1:
        fmt.Println(v2, v2)
    default:
        if (v5 || (2 > genFn0(genFn0(65535)))) {
            var v39 int32 = int32(v2)
            _ = v39
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
