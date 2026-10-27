package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    var v1 string = "\n\t"
    var v2 int32 = 804
    v3 := []uint32{}
    fmt.Println(v1, v2)
    var v4 string = v1
    if (0 < 1000) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v5 := GenS1{F0: 9223372036854775807, F1: (^(v2)), F2: 128}
        fmt.Println(v1, v1, v2)
        _ = v5
    }
    if false {
        v4 = "hello"
    } else {
        switch 2147483647 {
        case -274:
            var v6 int8 = 1
            _ = v6
        case 2147483647:
            v4 = v1
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v2 += v2
    }
    v1 = (v4 + "\n\t")
    var v7 int64 = 0
    v8 := GenS2{F0: -128}
    if (((!(((677 != 255) || false) || false)) || (42 < 1)) && true) {
        v9 := GenS3{F0: -1, F1: v2, F2: 42}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v10 float64 = ((0.1 * 3.14159) / (1e10 + (-2.25 + 3.14159)))
        _ = v10
        _ = v9
    } else {
        v1 = ((v4 + v1) + (v4 + v1))
        var v11 uint = 65535
        _ = v11
    }
    if (686 <= 1) {
        for fi12 := 0; fi12 < 20; fi12++ {
            if (325 != 2) {
                _ = len(v3)
                v2 = (v2 / ((-448) | 1))
            }
            _ = fi12
        }
        var v13 uint8 = 255
        _ = v13
    } else {
        var v14 uint32 = 4294967295
        fmt.Println(v4)
        _ = v14
    }
    switch 127 {
    case 32768:
        _ = len(v3)
    }
    var v15 int64 = 0
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            v4 = ("hello" + "")
            var v17 uint64 = 4294967295
            v3 = append(v3, 364)
            _ = v17
        } else {
            v4 = ((v1 + v4) + ("\n\t" + v1))
            var v18 uint = 42
            switch (v4 + (("foo bar" + "\n\t") + v4)) {
            case "":
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(v1, fi19, v4)
                    fmt.Println(v2, v7, v2)
                    var v20 int64 = int64(fi19)
                    _ = fi19
                    _ = v20
                }
            default:
                var v21 bool = false
                _ = v21
            }
            _ = v18
        }
        v2 += 42
        var v22 bool = true
        _ = fi16
        _ = v22
    }
    v23 := GenS4{F0: 42, F1: 136, F2: 42}
    v24 := GenS5{F0: (204 < 0), F1: 2}
    v25 := GenS2{F0: 42}
    for fi26 := 0; fi26 < 18; fi26++ {
        fmt.Println(v1)
        _ = fi26
    }
    var v27 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
