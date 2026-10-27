package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 3.14159
    var v3 string = ""
    v4 := []string{("foo bar" + v3), v3, ("\n\t" + "foo bar"), ("a" + (v3 + v3))}
    v5 := make(map[int]uint32)
    var v6 uint16 = 133
    v6 = v6
    _ = len(v4)
    var v7 int32 = 42
    v8 := GenS1{F0: 255}
    fmt.Println(v6, v6)
    if false {
        v2 += v2
        v7 = (v7 / ((v7) | 1))
    } else {
        v9 := GenS2{F0: false, F1: 4294967295}
        _ = v9
    }
    fmt.Println(v6, v6)
    if ((true || true) && false) {
        v10 := GenS1{F0: 154}
        switch 128 {
        case -190:
            for fi11 := 0; fi11 < 14; fi11++ {
                switch (fi11 % (((-(fi11))) | 1)) {
                case 307:
                    v3 = v3
                case 127:
                    var v12 uint16 = 0
                    _ = v12
                }
                _ = fi11
            }
        case 2147483647:
            v13 := GenS3{F0: v6}
            _ = v13
        case 949:
            switch "foo bar" {
            case "":
                var v14 float32 = 0.0
                _ = v14
            case "hello":
                var v15 string = v3
                _ = v15
            case "foo bar":
                for fi16 := 0; fi16 < 12; fi16++ {
                    v2 = 0.1
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v18 := GenS4{F0: -584, F1: 760}
                        delete(v5, (fi17 / ((4294967295) | 1)))
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
            default:
                var v19 string = (v3 + ("hello" + (v3 + v3)))
                _ = v19
            }
        default:
            var v20 bool = true
            _ = v20
        }
        delete(v5, 32768)
        _ = v10
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3 = ""
    }
    var v21 bool = true
    v22 := GenS1{F0: 23}
    v23 := GenS3{F0: (v6 / ((v6) | 1))}
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 uint8 = 202
            v4 = append(v4, (v3 + "hello"))
            fmt.Println(fi25, v3)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    var v27 string = "\n\t"
    v28 := v1[9223372036854775807]
    v29 := GenS3{F0: v6}
    switch 766 {
    case 657:
        switch -782 {
        case 454:
            v1[32768] = true
        case 9223372036854775807:
            var v30 float32 = ((1.5 - (0.1 / (0.1 - (1e10 / 0.1)))) + (1e10 - 3.14159))
            _ = v30
        default:
            var v31 uint16 = (v6 / (((v6 + (^(v6)))) | 1))
            _ = v31
        }
    case 754:
        for fi32 := 0; fi32 < 13; fi32++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi32
        }
    case 32767:
        v33 := GenS4{F0: -883, F1: -9223372036854775808}
        _ = v33
    }
    var v34 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
