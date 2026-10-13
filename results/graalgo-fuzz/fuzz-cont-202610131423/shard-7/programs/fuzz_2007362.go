package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
func genFn0(p0 uint8) string {
    var v4 int64 = 1000
    fmt.Println(v4, p0)
    _ = p0
    _ = v4
    return "\n\t"
}
func genFn1(p0 bool, p1 uint32) int32 {
    if p0 {
        switch -1 {
        case 551:
            v5 := []int{2147483647, 4294967295, 9223372036854775807, -686}
            _ = v5
        }
        fmt.Println(p1, p1)
    } else {
        v6 := []int{-910}
        var v7 uint16 = 715
        var v8 int64 = 1
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 32767
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 0
    var v3 uint64 = v2
    switch 32767 {
    case -9223372036854775808:
        v10 := []int{}
        _ = v10
    case -2147483648:
        var v11 bool = true
        _ = v11
    }
    var v12 float64 = 3.14159
    if false {
        if false {
            fmt.Println(v12, v1)
        }
        var v13 uint16 = 65535
        v14 := GenS1{F0: 32767, F1: 127, F2: "foo bar"}
        _ = v13
        _ = v14
    }
    v15 := []int{348, 2147483647, -2147483648, -2147483648}
    if false {
        if ((v3 % ((v2) | 1)) == (v3 + (v3 - 4294967295))) {
            v15 = append(v15, 209)
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS2{F0: 258, F1: 2, F2: v2}
                _ = fi16
                _ = v17
            }
        } else {
            fmt.Println(v12)
            fmt.Println(v12, v1)
        }
    } else {
        if false {
            var v18 float64 = float64(3.14159)
            for fi19 := 0; fi19 < 3; fi19++ {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 float64 = v18
                    if (true && (-351 >= 1)) {
                        v21 -= (v21 + v18)
                        switch "hello" {
                        case "foo bar":
                            _ = len(v15)
                        case "a":
                            var v22 string = ""
                            _ = v22
                        default:
                            var v23 int64 = 42
                            _ = v23
                        }
                    } else {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        var v24 int16 = -32768
                        _ = v24
                    }
                    var v25 bool = true
                    _ = fi20
                    _ = v21
                    _ = v25
                }
                _ = fi19
            }
            _ = v18
        }
    }
    if (!(2147483647 >= 585)) {
        fmt.Println(v1, v2)
        v2 = 4294967295
    } else {
        switch 0 {
        case 182:
            var v26 int8 = 100
            _ = v26
        case 1:
            fmt.Println(v1, v1)
        default:
            v27 := GenS3{F0: true, F1: 213, F2: "hello"}
            _ = v27
        }
        var v28 int = -2147483648
        var v29 string = "\n\t"
        _ = v28
        _ = v29
    }
    if (false && (100 < 1)) {
        fmt.Println(v2)
        for fi30 := 0; fi30 < 5; fi30++ {
            v3 += (v2 / ((2) | 1))
            _ = fi30
        }
        v2 = 0
    } else {
        v2 = 42
        if false {
            var v31 int = 626
            var v32 uint32 = 4294967295
            for fi33 := 0; fi33 < 17; fi33++ {
                var v34 int16 = 497
                _ = fi33
                _ = v34
            }
            _ = v31
            _ = v32
        } else {
            fmt.Println(v2, v1)
            fmt.Println(v12, v2)
            if true {
                if (42 >= -1) {
                    _ = len(v15)
                    var v35 uint = 255
                    _ = v35
                }
                var v36 int8 = 0
                v3 = 4294967295
                _ = v36
            }
        }
    }
    var v37 int = -594
    fmt.Println(v37, v3, v3)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v38 float64 = (1.5 + float64((0.1 - 3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
}
