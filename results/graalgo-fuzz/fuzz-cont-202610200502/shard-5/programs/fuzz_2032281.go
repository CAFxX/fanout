package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{((("\n\t" + (("\n\t" + ("hello" + "hello")) + ("a" + ((("" + "hello") + "hello") + "\n\t")))) + ("" + "\n\t")) + ("\n\t" + "")), "foo bar", ("\n\t" + ""), "foo bar"}
    var v3 float64 = -2.25
    var v4 uint32 = 613
    var v5 float64 = (v3 - v3)
    v4 = (v4 >> ((2) & 63))
    var v6 uint32 = 2
    var v7 int16 = 1
    switch "a" {
    case "":
        var v8 float32 = 3.14159
        _ = v8
    case "hello":
        _ = len(v1)
    case "a":
        fmt.Println(v6, v4, v4)
    }
    v6 = 2147483648
    if (2 > 42) {
        v9 := GenS1{F0: 1, F1: 934, F2: 0}
        if true {
            if true {
                v10 := GenS2{F0: -1, F1: 4294967295}
                _ = v10
            }
            var v11 uint = 42
            delete(v1, (v6 % ((v4) | 1)))
            _ = v11
        } else {
            switch ("" + "\n\t") {
            case "a":
                var v12 int16 = 32767
                _ = v12
            case "":
                fmt.Println(v6, v6, v4)
            default:
                var v13 uint32 = 2147483648
                _ = v13
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(v4, fi14)
                var v15 uint = uint(v6)
                _ = fi14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            fmt.Println(v3, v6, v6)
            v17 := GenS3{F0: 255, F1: 432}
            fmt.Println(fi16, v6, v7)
            _ = fi16
            _ = v17
        }
        _ = v9
    } else {
        delete(v1, 2)
    }
    v1[v4] = (!((false && true) || (false || ("a" != "hello"))))
    var v18 int = 143
    for fi19 := 0; fi19 < 7; fi19++ {
        v7 = v7
        _ = fi19
    }
    fmt.Println(v6)
    if (true || false) {
        var v20 int = v18
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int8 = int8(v6)
            v7 -= (v7 - (^(v7)))
            _ = fi21
            _ = v22
        }
        var v23 bool = true
        _ = v20
        _ = v23
    } else {
        if len(v2) > 0 {
            v2[0] = ""
        }
        if (("foo bar" + "a") <= ("a" + ("foo bar" + "foo bar"))) {
            for fi24 := 0; fi24 < 1; fi24++ {
                v5 += 1.5
                _ = fi24
            }
            v1[(v6 << ((0) & 63))] = true
            var v25 uint32 = v4
            _ = v25
        } else {
            fmt.Println(v6, v7, v6)
        }
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v7 = (^(v7))
        var v27 float32 = 0.0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
