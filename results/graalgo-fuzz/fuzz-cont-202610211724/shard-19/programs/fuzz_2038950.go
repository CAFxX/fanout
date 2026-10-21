package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    var v1 float32 = 1.5
    v2 := []int{1, 127, -1, -9223372036854775808}
    var v3 float32 = 1.5
    v3 += v3
    var v4 float32 = (((v1 / (0.0 * v1)) * v3) / v3)
    if (true && false) {
        fmt.Println(v3, v3, v1)
    } else {
        if true {
            fmt.Println(v4, v3)
            v2 = append(v2, 32767)
        }
        fmt.Println(v1, v1, v4)
    }
    v1 = v3
    if true {
        fmt.Println(v3)
    }
    v4 = 0.1
    fmt.Println(v1, v1)
    v1 -= (v4 - (v4 + 3.14159))
    if len(v2) > 0 {
        v2[0] = 32767
    }
    v5 := GenS1{F0: false, F1: -669, F2: 0}
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (-1 > 9223372036854775807) {
                var v8 float64 = 1.5
                _ = v8
            }
            v9 := GenS2{F0: 198}
            _ = fi7
            _ = v9
        }
        if (((fi6 & (-(fi6))) != 42) || (!((("hello" + "\n\t") > "") && (1e10 < (1.5 * -2.25))))) {
            fmt.Println(fi6, v4)
            switch ("" + "a") {
            case "a":
                v2 = append(v2, fi6)
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                switch "a" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = fi6
                    }
                case "hello":
                    v2 = append(v2, fi6)
                }
            }
        } else {
            var v10 string = "\n\t"
            v11 := GenS1{F0: (((v1 < v4) && true) || false), F1: -1, F2: 0}
            switch (v10 + (v10 + "\n\t")) {
            case "foo bar":
                switch (fi6 + fi6) {
                case 0:
                    var v12 bool = false
                    _ = v12
                case -47:
                    for fi13 := 0; fi13 < 7; fi13++ {
                        var v14 string = v10
                        _ = fi13
                        _ = v14
                    }
                }
            }
            _ = v10
            _ = v11
        }
        switch "a" {
        case "":
            v2 = append(v2, (fi6 >> ((42) & 63)))
        default:
            _ = len(v2)
        }
        _ = fi6
    }
    v1 -= -2.25
    var v15 int8 = -97
    fmt.Println(v4)
    if (255 > 2) {
        v16 := GenS3{F0: 2147483647, F1: (-1 <= 0)}
        if false {
            var v17 uint32 = 42
            fmt.Println(v15, v3)
            var v18 string = "a"
            _ = v17
            _ = v18
        } else {
            v15 = v15
            v15 -= v15
            var v19 string = "foo bar"
            _ = v19
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            v1 = 1e10
            fmt.Println(v3)
            var v21 int64 = 9223372036854775807
            _ = fi20
            _ = v21
        }
        _ = len(v2)
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            fmt.Println(v4)
        }
        fmt.Println(fi22, v1)
        var v23 int = (fi22 | (fi22 + 9223372036854775807))
        _ = fi22
        _ = v23
    }
    var v24 int16 = 0
    for fi25 := 0; fi25 < 12; fi25++ {
        if true {
            var v26 uint64 = 0
            _ = v26
        } else {
            v27 := GenS2{F0: 2}
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 uint = 2
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            if false {
                v31 := GenS2{F0: 18446744073709551615}
                var v32 int = fi25
                v33 := GenS3{F0: fi25, F1: (float64(v4) >= 1.5)}
                var v34 string = "a"
                if len(v2) > 0 {
                    v2[0] = fi30
                }
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            }
            fmt.Println(v3)
            var v35 bool = false
            _ = fi30
            _ = v35
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
