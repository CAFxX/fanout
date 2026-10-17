package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 string }
func genFn0(p0 string, p1 uint8, p2 int8) int8 {
    p0 = p0
    v4 := GenS1{F0: ("foo bar" + "foo bar"), F1: p2, F2: ""}
    switch "hello" {
    case "a":
        p2 -= (p2 - p2)
    case "":
        var v5 int32 = -2147483648
        _ = v5
    case "\n\t":
        if (1.5 == (3.14159 - ((0.1 - -2.25) + 0.0))) {
            v6 := []int{1, 128, -746, 32768}
            _ = v6
        }
    }
    p1 = p1
    var v7 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return p2
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    fmt.Println(v3)
    if (!(!((1 == 2) && false))) {
        v3 = (v3 + v3)
        _ = len(v2)
        v8 := GenS1{F0: "foo bar", F1: 100, F2: v3}
        _ = v8
    } else {
        var v9 uint64 = 4294967295
        v10 := GenS1{F0: ("foo bar" + v3), F1: genFn0((v3 + "hello"), 255, 42), F2: "foo bar"}
        _ = v10
        _ = v9
    }
    if false {
        v3 = ""
        v11 := GenS1{F0: v3, F1: -78, F2: (((v3 + v3) + "hello") + "")}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            v3 = v3
            var v13 uint64 = 339
            _ = fi12
            _ = v13
        }
    }
    var v14 int32 = 0
    delete(v2, 812)
    v14 -= (v14 & 524)
    for fi15 := 0; fi15 < 17; fi15++ {
        v3 = (v3 + "a")
        _ = fi15
    }
    v3 = "hello"
    if false {
        fmt.Println(v3)
        if (1 > 255) {
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS1{F0: v3, F1: -1, F2: v3}
                switch "hello" {
                case "\n\t":
                    var v18 uint8 = 228
                    _ = v18
                case "hello":
                    v3 = v3
                case "foo bar":
                    fmt.Println(v3)
                default:
                    if len(v1) > 0 {
                        v1[0] = (fi16 & fi16)
                    }
                }
                v3 = v3
                _ = fi16
                _ = v17
            }
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            v3 = "\n\t"
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi19
        }
    }
    fmt.Println(v14)
    if true {
        for fi20 := 0; fi20 < 14; fi20++ {
            if true {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v14 = (v14 / (((v14 % (((v14 % ((v14) | 1))) | 1))) | 1))
                    fmt.Println(fi20, v14, v3)
                    var v22 int8 = 1
                    _ = fi21
                    _ = v22
                }
                if false {
                    var v23 uint = 386
                    var v24 int = (fi20 - 9223372036854775807)
                    _ = v23
                    _ = v24
                }
            }
            switch fi20 {
            case 128:
                var v25 uint16 = 96
                _ = v25
            case 0:
                if false {
                    fmt.Println(fi20, v14, fi20)
                    v14 = -2147483648
                    v3 = ""
                }
            }
            fmt.Println(v3, v3, v14)
            _ = fi20
        }
        v3 = v3
        if true {
            for fi26 := 0; fi26 < 11; fi26++ {
                _ = len(v1)
                v2[742] = ("foo bar" + v3)
                v3 = (v3 + v3)
                _ = fi26
            }
            v27 := GenS1{F0: ((v3 + v3) + v3), F1: 0, F2: ""}
            _ = v27
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = ""
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
