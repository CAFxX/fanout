package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 string, p1 uint32, p2 int32) uint8 {
    var v5 string = (p0 + "a")
    v6 := []int{-175, -281, -393}
    p2 = 2147483647
    switch p0 {
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "foo bar":
        var v7 bool = false
        _ = v7
    case "\n\t":
        switch 42 {
        case 369:
            if false {
                p1 = (p1 - (p1 * 1))
                p1 += p1
            } else {
                p1 += (p1 / ((p1) | 1))
            }
        }
    }
    p0 = v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int8, p1 int8, p2 int16) int16 {
    v8 := GenS1{F0: 537}
    if true {
        var v9 int8 = p0
        _ = v9
    }
    var v10 uint64 = 18446744073709551615
    p1 += (p0 + -1)
    v11 := []int{892, 298}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    return 32767
}

func main() {
    var v1 int16 = -1000
    var v2 float32 = 3.14159
    var v3 int32 = 808
    v4 := make(map[string]string)
    v3 = v3
    v12 := GenS1{F0: 2}
    if true {
        fmt.Println(v3, v1)
        fmt.Println(v1)
        fmt.Println(v2, v2, v3)
    }
    if (42 == 127) {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v3, v1, v3)
        default:
            v3 += (v3 >> ((42) & 63))
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            v2 = v2
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS1{F0: 587}
                v3 = v3
                v2 += v2
                _ = fi14
                _ = v15
            }
            v2 = 1e10
            _ = fi13
        }
    } else {
        v16 := GenS1{F0: 895}
        _ = v16
    }
    var v17 int64 = 931
    v18 := []int{}
    v3 -= v3
    if false {
        if false {
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(v17, v17)
                _ = len(v4)
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = fi19
            }
            switch "hello" {
            case "foo bar":
                v20 := GenS1{F0: 557}
                _ = v20
            case "\n\t":
                if true {
                    v18 = append(v18, 4294967295)
                }
            }
            var v21 int = -9223372036854775808
            _ = v21
        }
    } else {
        v17 -= v17
    }
    var v22 int64 = v17
    var v23 int16 = genFn1(-128, -128, v1)
    delete(v4, "")
    fmt.Println(v23, v17)
    for fi24 := 0; fi24 < 6; fi24++ {
        if true {
            v25 := GenS2{F0: 127, F1: 524, F2: (65535 < 0)}
            switch "foo bar" {
            case "a":
                v26 := GenS3{F0: genFn0("\n\t", 0, (v3 >> ((0) & 63))), F1: 0}
                _ = v26
            case "\n\t":
                for fi27 := 0; fi27 < 7; fi27++ {
                    v4["hello"] = "foo bar"
                    v1 += -783
                    _ = len(v4)
                    _ = fi27
                }
            case "":
                var v28 string = ""
                _ = v28
            default:
                if false {
                    var v29 uint = 776
                    _ = v29
                } else {
                    v22 = v17
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v31 := GenS2{F0: -128, F1: (-(fi30)), F2: ((v17 / ((9223372036854775807) | 1)) <= -9223372036854775808)}
                        _ = fi30
                        _ = v31
                    }
                }
            }
            v3 -= (^(v3))
            _ = v25
        }
        switch (fi24 ^ -582) {
        case 127:
            v32 := v4[("" + "\n\t")]
            _ = v32
        }
        v17 = int64(v3)
        _ = fi24
    }
    v17 -= (-(v22))
    v18 = append(v18, 4294967295)
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS2{F0: -1, F1: (^(fi33)), F2: true}
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS4{F0: v22, F1: -98}
            v37 := v4["a"]
            v3 = v3
            _ = fi35
            _ = v36
            _ = v37
        }
        v17 += v17
        _ = fi33
        _ = v34
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := v4["\n\t"]
        v23 = 72
        v1 -= 1000
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
