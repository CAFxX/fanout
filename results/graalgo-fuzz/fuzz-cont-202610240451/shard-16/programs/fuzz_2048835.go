package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 bool; F1 int; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 32767
    v3 := []int{870, 2147483647}
    var v4 int32 = 42
    var v5 int64 = -526
    var v6 int8 = -1
    if false {
        fmt.Println(v4, v4)
        v2 = v2
    } else {
        v7 := GenS1{F0: 626, F1: -1, F2: (v2 << ((911) & 63))}
        _ = v7
    }
    switch (("hello" + "hello") + "") {
    case "a":
        var v8 uint = 4294967295
        _ = v8
    case "\n\t":
        if true {
            if false {
                if true {
                    v5 = (v5 * (v5 ^ v5))
                    v2 = (^(v2))
                    var v9 int32 = (v4 + v4)
                    _ = v9
                }
                v5 = v5
            }
        } else {
            v6 = v6
            if true {
                if len(v3) > 0 {
                    v3[0] = 32768
                }
            } else {
                v10 := GenS2{F0: v4, F1: 42, F2: "hello"}
                if false {
                    v11 := GenS3{F0: 177, F1: (!false)}
                    var v12 bool = (v2 >= v2)
                    _ = v11
                    _ = v12
                }
                _ = v10
            }
            if false {
                switch 14 {
                case 42:
                    v13 := GenS3{F0: (v4 & 2147483647), F1: (true || (!(0.0 >= 1.5)))}
                    _ = v13
                case 1:
                    v2 = 42
                }
                v5 -= -892
                if true {
                    _ = len(v3)
                }
            }
        }
    case "foo bar":
        if (!((0 == 1) && (0.0 <= 3.14159))) {
            for fi14 := 0; fi14 < 14; fi14++ {
                v5 += 9223372036854775807
                fmt.Println(v4, v6, fi14)
                v3 = append(v3, int(v4))
                _ = fi14
            }
        }
    default:
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint16 = 754
            fmt.Println(v5, v4)
            v16 = v16
            _ = fi15
            _ = v16
        }
    }
    var v17 string = ((("\n\t" + "foo bar") + (("hello" + "foo bar") + (("" + "") + (((("" + (("\n\t" + ("\n\t" + ((((("\n\t" + ("\n\t" + "\n\t")) + "hello") + "\n\t") + ("" + "a")) + ((("\n\t" + "a") + "foo bar") + "a")))) + "")) + "\n\t") + ("\n\t" + (("a" + "a") + "foo bar"))) + "\n\t")))) + "foo bar")
    var v18 float64 = -2.25
    v2 = 32767
    fmt.Println(v6, v2)
    v19 := GenS4{F0: (v2 >> ((18446744073709551615) & 63)), F1: 42}
    var v20 bool = true
    if v20 {
        v1[42] = v17
        switch -9223372036854775808 {
        case -31:
            v6 -= (^(v6))
        case -2147483648:
            var v21 int32 = int32(v5)
            _ = v21
        }
    }
    v18 = (v18 * v18)
    if v20 {
        if true {
            v22 := GenS5{F0: 745, F1: true}
            switch "hello" {
            case "":
                _ = len(v1)
            case "\n\t":
                v5 = -506
            }
            _ = v22
        }
        if v20 {
            v23 := GenS4{F0: v2, F1: -2147483648}
            var v24 float64 = ((1.5 * v18) + float64(0.0))
            _ = v23
            _ = v24
        } else {
            v25 := GenS6{F0: 65535, F1: 0}
            var v26 float32 = -2.25
            v27 := GenS7{F0: ("hello" > "\n\t"), F1: -547, F2: v2}
            _ = v25
            _ = v26
            _ = v27
        }
    }
    fmt.Println(v17, v5)
    v5 += (v5 % ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
