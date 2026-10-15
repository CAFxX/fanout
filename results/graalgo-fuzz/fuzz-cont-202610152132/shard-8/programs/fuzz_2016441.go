package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 string }
func genFn0(p0 uint32, p1 uint) int32 {
    if false {
        v4 := []int{42, 128}
        var v5 int64 = -1
        switch 32767 {
        case 42:
            if true {
                fmt.Println(p0, v5, p0)
            } else {
                p0 = (p0 | p0)
                fmt.Println(p1, v5)
            }
        case -2147483648:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 32768:
            v5 = (v5 % ((0) | 1))
        }
        _ = v4
        _ = v5
    }
    switch "hello" {
    case "\n\t":
        switch "" {
        case "foo bar":
            p1 = p1
        case "\n\t":
            var v6 string = ("" + ((("" + ("a" + ("" + "\n\t"))) + "a") + "a"))
            _ = v6
        case "":
            switch -9223372036854775808 {
            case 127:
                p0 = 42
            case -1:
                p1 = p1
            case 565:
                p1 = (p1 % ((p1) | 1))
            }
        }
    }
    _ = p0
    _ = p1
    return 896
}

func main() {
    v1 := []string{"a", "hello"}
    var v2 float64 = 3.14159
    var v3 float64 = 1.5
    v3 = 0.1
    if true {
        if true {
            switch "\n\t" {
            case "foo bar":
                var v7 int32 = 0
                _ = v7
            }
            var v8 uint8 = 127
            _ = v8
        }
    } else {
        v9 := GenS1{F0: "foo bar", F1: -1, F2: "\n\t"}
        var v10 uint16 = 255
        _ = v10
        _ = v9
    }
    v2 = 3.14159
    for fi11 := 0; fi11 < 1; fi11++ {
        switch ("a" + (("a" + "hello") + "")) {
        case "a":
            if (!true) {
                switch "foo bar" {
                case "\n\t":
                    var v12 uint64 = 42
                    _ = v12
                case "":
                    v13 := GenS1{F0: "foo bar", F1: 94, F2: "\n\t"}
                    _ = v13
                case "hello":
                    switch 0 {
                    case -757:
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    case 7:
                        var v14 uint = 2
                        _ = v14
                    case 32768:
                        var v15 float64 = v2
                        _ = v15
                    default:
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                }
                var v16 int8 = 59
                var v17 int64 = 0
                _ = v16
                _ = v17
            }
        case "":
            for fi18 := 0; fi18 < 11; fi18++ {
                fmt.Println(v3)
                var v19 string = "\n\t"
                _ = fi18
                _ = v19
            }
        case "hello":
            v3 -= (v3 - v2)
        default:
            v3 += (v3 / ((v3 - 1.5) - v3))
        }
        v3 = v2
        _ = fi11
    }
    fmt.Println(v3, v2)
    switch "" {
    case "\n\t":
        v3 = v3
    }
    var v20 int8 = -1
    v2 -= 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
