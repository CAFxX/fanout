package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint; F1 uint; F2 uint }

func main() {
    v1 := []uint32{42, 0, 99}
    var v2 bool = true
    var v3 uint = 0
    v4 := []uint32{235, 2147483647, 2147483647}
    var v5 int64 = 1
    var v6 uint64 = 18446744073709551615
    switch 9223372036854775807 {
    case 520:
        v5 = (^(v5))
    case 127:
        var v7 string = "foo bar"
        _ = v7
    }
    if v2 {
        v4 = append(v4, 655)
        switch ("" + ("hello" + "hello")) {
        case "hello":
            var v8 float64 = ((0.1 + 1e10) * -2.25)
            _ = v8
        case "":
            if (-2.25 <= (3.14159 + (float64(3.14159) * 1.5))) {
                var v9 float64 = -2.25
                var v10 int32 = 0
                _ = v10
                _ = v9
            }
        case "\n\t":
            fmt.Println(v2, v6)
        default:
            v11 := GenS1{F0: "hello"}
            _ = v11
        }
        var v12 int32 = -718
        _ = v12
    } else {
        if v2 {
            fmt.Println(v3)
        } else {
            var v13 float32 = 0.1
            for fi14 := 0; fi14 < 7; fi14++ {
                _ = len(v1)
                _ = fi14
            }
            _ = v13
        }
        v3 = 255
        v15 := GenS1{F0: ""}
        _ = v15
    }
    if false {
        switch ((((("a" + ("\n\t" + "a")) + "foo bar") + ("hello" + "")) + "foo bar") + "") {
        case "":
            switch 32768 {
            case 127:
                var v16 uint = v3
                _ = v16
            case -9223372036854775808:
                fmt.Println(v2)
            }
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 int8 = 0
                _ = fi18
                _ = v19
            }
            fmt.Println(v6)
            _ = fi17
        }
        var v20 bool = v2
        _ = v20
    } else {
        var v21 int32 = 2147483647
        _ = len(v1)
        for fi22 := 0; fi22 < 14; fi22++ {
            fmt.Println(fi22)
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int32 = int32(v5)
                _ = fi23
                _ = v24
            }
            v25 := GenS2{F0: uint32(v3)}
            _ = fi22
            _ = v25
        }
        _ = v21
    }
    var v26 int64 = 789
    if v2 {
        v27 := GenS3{F0: "\n\t", F1: -128, F2: ""}
        v26 -= v26
        if v2 {
            var v28 int = 2147483647
            var v29 int = v28
            _ = v28
            _ = v29
        } else {
            fmt.Println(v2, v2, v5)
            var v30 int = -9223372036854775808
            v31 := GenS4{F0: 465, F1: v26}
            _ = v30
            _ = v31
        }
        _ = v27
    }
    v1 = append(v1, 1)
    v5 -= (v26 % ((v26) | 1))
    _ = len(v4)
    if v2 {
        if (127 > 2147483647) {
            v32 := GenS2{F0: 0}
            v33 := GenS5{F0: (v3 - v3), F1: 18446744073709551615, F2: (v3 >> ((v3) & 63))}
            v2 = (false || (-158 >= 873))
            _ = v32
            _ = v33
        } else {
            var v34 uint32 = 973
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
