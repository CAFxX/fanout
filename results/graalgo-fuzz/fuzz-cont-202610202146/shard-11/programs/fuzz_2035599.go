package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0(p0 uint32, p1 int32) uint16 {
    var v6 uint32 = p0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v6
    return 296
}
func genFn1(p0 int) uint {
    var v7 uint16 = 255
    var v8 int16 = -1
    fmt.Println(v8, p0, v7)
    v9 := []int{}
    var v10 int64 = 367
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 402
}

func main() {
    var v1 uint8 = 1
    var v2 bool = true
    v3 := make(map[string]uint32)
    var v4 uint64 = 0
    v5 := []int{32767, 821, 32768}
    v1 += (v1 >> ((42) & 63))
    var v11 int32 = 2147483647
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v12 := GenS1{F0: 91, F1: 42}
    var v13 uint16 = 2
    var v14 int64 = 0
    switch "\n\t" {
    case "":
        if v2 {
            for fi15 := 0; fi15 < 11; fi15++ {
                fmt.Println(v14, v1)
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := v3[("" + ("\n\t" + (("\n\t" + "\n\t") + "")))]
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            var v18 uint32 = 42
            _ = v18
        }
    case "hello":
        var v19 int = 424
        _ = v19
    default:
        v13 -= uint16(v1)
    }
    v1 += 1
    for fi20 := 0; fi20 < 13; fi20++ {
        if v2 {
            if v2 {
                var v21 string = ("" + "")
                _ = v21
            }
            v22 := GenS2{F0: v14, F1: (v1 + v1)}
            switch ("" + "a") {
            case "foo bar":
                var v23 uint64 = (v4 - 2)
                _ = v23
            default:
                if (!false) {
                    var v24 bool = v2
                    _ = v24
                } else {
                    switch "foo bar" {
                    case "":
                        var v25 int32 = v11
                        _ = v25
                    case "foo bar":
                        var v26 uint = 1
                        _ = v26
                    }
                    _ = len(v3)
                }
            }
            _ = v22
        } else {
            if true {
                switch "\n\t" {
                case "hello":
                    var v27 bool = v2
                    _ = v27
                }
            }
        }
        v2 = true
        switch "" {
        case "\n\t":
            switch (fi20 % (((-(fi20))) | 1)) {
            case 32768:
                _ = len(v5)
            case 805:
                var v28 bool = (((^(v4)) <= (v4 / ((v4) | 1))) && v2)
                _ = v28
            }
        case "hello":
            v3["hello"] = 1
        case "":
            v5 = append(v5, fi20)
        default:
            if ((v2 && v2) && ((v11 >> ((genFn1((fi20 - int(v4)))) & 63)) != (v11 >> ((742) & 63)))) {
                var v29 uint = 2
                _ = v29
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(v2, v4, v13)
                case "a":
                    if false {
                        var v30 float32 = (3.14159 * -2.25)
                        v30 = 1.5
                        _ = v30
                    }
                default:
                    var v31 bool = (v2 && v2)
                    _ = v31
                }
                v4 -= 681
            }
        }
        _ = fi20
    }
    if (uint16(v1) < v13) {
        v4 = v4
        var v32 string = "hello"
        var v33 int8 = 0
        _ = v32
        _ = v33
    } else {
        fmt.Println(v11, v14, v14)
        var v34 float32 = (1e10 * 0.1)
        var v35 int = -978
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: v14, F1: 42}
    v37 := GenS4{F0: 4294967295, F1: 978, F2: 32767}
    fmt.Println(v2, v4, v2)
    v38 := GenS5{F0: (^(v13)), F1: 1000, F2: -1}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
