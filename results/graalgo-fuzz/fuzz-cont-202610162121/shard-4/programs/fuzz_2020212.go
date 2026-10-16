package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
func genFn0(p0 int8) uint32 {
    if ((!false) && true) {
        if (false && ((2147483647 != 141) && true)) {
            if true {
                fmt.Println(p0)
                for fi5 := 0; fi5 < 11; fi5++ {
                    v6 := []int{fi5}
                    p0 -= (p0 & p0)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi5
                    _ = v6
                }
                if (!((3.14159 - (3.14159 - (0.1 / 1.0))) >= (3.14159 - (1e10 / 1.5)))) {
                    fmt.Println(p0)
                    for fi7 := 0; fi7 < 15; fi7++ {
                        v8 := GenS1{F0: 80}
                        _ = fi7
                        _ = v8
                    }
                }
            }
            p0 -= p0
            fmt.Println(p0)
        }
        if true {
            switch ("" + "hello") {
            case "foo bar":
                switch "\n\t" {
                case "":
                    fmt.Println(p0)
                default:
                    var v9 float32 = (0.1 - 1.5)
                    _ = v9
                }
            default:
                var v10 int64 = -9223372036854775808
                _ = v10
            }
            if (32768 >= 32768) {
                fmt.Println(p0)
            } else {
                if (!(304 >= 263)) {
                    var v11 uint64 = 2
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 uint8 = 42
                        fmt.Println(fi12, v13)
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS2{F0: 128}
                    _ = v11
                    _ = v14
                } else {
                    var v15 int64 = -802
                    v15 = (-(v15))
                    _ = v15
                }
                v16 := GenS2{F0: 24}
                _ = v16
            }
        }
        if false {
            p0 = 100
            v17 := []int{128, -84, 32768, 303}
            _ = v17
        } else {
            v18 := []int{9223372036854775807}
            v19 := GenS2{F0: 128}
            var v20 uint8 = 182
            _ = v18
            _ = v19
            _ = v20
        }
    }
    fmt.Println(p0)
    v21 := GenS1{F0: 2}
    _ = p0
    _ = v21
    return 641
}
func genFn1() int16 {
    var v22 float64 = 3.14159
    v22 += float64(((-2.25 - 0.1) / 1.5))
    _ = v22
    return 32767
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 986
    var v3 bool = (-9223372036854775808 == -9223372036854775808)
    var v4 int8 = -128
    if ((!false) && v3) {
        var v23 bool = (815 == 931)
        _ = v23
    } else {
        v24 := v1[72]
        v25 := GenS1{F0: 2}
        v26 := []int{510}
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println(v4, v3, v3)
    v3 = v3
    var v27 int64 = 1
    for fi28 := 0; fi28 < 3; fi28++ {
        if true {
            delete(v1, fi28)
        }
        _ = fi28
    }
    v2 += (v2 + v2)
    v29 := GenS1{F0: 1}
    switch ("" + ("" + "\n\t")) {
    case "foo bar":
        v30 := GenS2{F0: 42}
        _ = v30
    case "\n\t":
        v31 := GenS2{F0: 226}
        _ = v31
    case "":
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 uint32 = 4294967295
            var v34 int32 = 0
            fmt.Println(v3, v2, v33)
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    fmt.Println(v27, v3, v27)
    v35 := []int{307, 742}
    for fi36 := 0; fi36 < 7; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            v3 = (float64(((float32((1.5 + 1.5)) + (1e10 * 0.0)) / (0.1 / float32((3.14159 + 1.5))))) < 1.5)
            fmt.Println(fi37, fi37, v3)
            v2 -= (v2 - (^(v2)))
            if (((v4 * (v4 | v4)) >= v4) && v3) {
                var v38 float32 = (-2.25 - 0.0)
                v39 := GenS1{F0: 42}
                switch (fi36 | -2147483648) {
                case 1:
                    var v40 uint8 = 255
                    _ = v40
                case 637:
                    if (879 < 32767) {
                        v3 = (v3 && v3)
                        fmt.Println(v38, v4)
                    } else {
                        v2 -= (v2 | 629)
                    }
                default:
                    v41 := GenS1{F0: 1}
                    _ = v41
                }
                _ = v38
                _ = v39
            }
            v42 := GenS3{F0: false, F1: "foo bar", F2: -128}
            fmt.Println(v3, v4, v2)
            _ = fi37
            _ = v42
        }
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
