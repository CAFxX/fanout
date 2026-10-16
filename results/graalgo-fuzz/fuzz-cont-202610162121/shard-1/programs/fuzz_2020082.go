package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
func genFn0() int32 {
    switch ("\n\t" + "") {
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi7, fi7)
                    var v9 uint32 = 860
                    _ = v9
                case "":
                    var v10 string = ("foo bar" + "hello")
                    _ = v10
                case "foo bar":
                    var v11 uint8 = 0
                    _ = v11
                default:
                    fmt.Println(fi7, fi8)
                }
                _ = fi8
            }
            var v12 uint8 = 10
            if (-1 > -223) {
                v13 := GenS1{F0: "\n\t", F1: (^(fi7)), F2: 127}
                _ = v13
            } else {
                v14 := GenS2{F0: 627, F1: v12, F2: ("a" + "foo bar")}
                _ = v14
            }
            _ = fi7
            _ = v12
        }
    case "a":
        v15 := GenS2{F0: 909, F1: 127, F2: "hello"}
        _ = v15
    case "\n\t":
        var v16 bool = true
        _ = v16
    }
    var v17 int = -9223372036854775808
    _ = v17
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 128
    var v3 uint32 = 174
    v4 := make(map[string]int)
    v5 := []int{-9223372036854775808, -9223372036854775808}
    var v6 string = "a"
    var v18 int8 = 127
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS3{F0: v18, F1: (v2 >= 1)}
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v22 := v4[v6]
    var v23 uint16 = (v21 >> ((51) & 63))
    fmt.Println(v1)
    delete(v4, (v6 + "\n\t"))
    if ((v3 + (v3 - v3)) > (v3 >> ((2) & 63))) {
        v18 += (v18 + (v18 / ((v18) | 1)))
    } else {
        switch v6 {
        case "hello":
            var v24 uint32 = v3
            _ = v24
        case "\n\t":
            for fi25 := 0; fi25 < 19; fi25++ {
                switch ((v6 + v6) + v6) {
                case "a":
                    fmt.Println(fi25)
                default:
                    v22 = (v22 / (((-(v22))) | 1))
                }
                switch v6 {
                case "\n\t":
                    v26 := GenS4{F0: v2}
                    _ = v26
                default:
                    v23 += v23
                }
                var v27 uint64 = 1
                _ = fi25
                _ = v27
            }
        }
        v6 = ("a" + ("a" + ((v6 + (v6 + v6)) + "\n\t")))
        var v28 uint16 = v21
        _ = v28
    }
    var v29 uint64 = 14
    fmt.Println(v18)
    _ = len(v5)
    var v30 uint64 = (v29 - (v29 ^ 172))
    if true {
        fmt.Println(v1, v2)
        switch v22 {
        case 127:
            var v31 int8 = v18
            _ = v31
        case -156:
            v3 -= 590
        default:
            _ = len(v5)
        }
    } else {
        v29 = 1
    }
    if (float32(1e10) < 0.1) {
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 float32 = 1.5
            for fi34 := 0; fi34 < 18; fi34++ {
                if true {
                    fmt.Println(v2, fi32)
                    v30 = (v29 | v30)
                } else {
                    v33 = v33
                    v21 = v21
                    v35 := GenS5{F0: "foo bar", F1: (v2 * v2), F2: ((3.14159 > 1.5) || true)}
                    _ = v35
                }
                var v36 uint64 = v29
                _ = fi34
                _ = v36
            }
            _ = fi32
            _ = v33
        }
    } else {
        fmt.Println(v3, v21, v3)
        v21 -= v23
        if (0 > 999) {
            var v37 uint32 = (v3 + (v3 - (v3 - v3)))
            var v38 int16 = 32767
            _ = v37
            _ = v38
        } else {
            delete(v4, v6)
            var v39 int32 = (v1 / (((v1 + (v1 >> ((25) & 63)))) | 1))
            _ = v39
        }
    }
    var v40 uint64 = 837
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
