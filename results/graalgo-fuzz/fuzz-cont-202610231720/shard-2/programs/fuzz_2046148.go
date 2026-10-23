package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 int16 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }
func genFn0(p0 int16) string {
    v6 := GenS1{F0: 17}
    switch 0 {
    case 860:
        var v7 float32 = ((-2.25 + 3.14159) + (0.1 + 1e10))
        _ = v7
    case 42:
        fmt.Println(p0)
    default:
        switch "hello" {
        case "foo bar":
            fmt.Println(p0)
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return "a"
}

func main() {
    v1 := []uint32{0}
    v2 := []string{"hello"}
    v3 := make(map[int]int)
    var v4 int = -808
    var v5 uint = 0
    var v8 uint32 = 2147483647
    v9 := GenS2{F0: 1}
    if false {
        var v10 int64 = 9223372036854775807
        v2 = append(v2, ("a" + "hello"))
        _ = v10
    } else {
        switch ("foo bar" + ((((((("foo bar" + "") + "foo bar") + "foo bar") + (("foo bar" + ("\n\t" + "")) + "hello")) + (("" + ("foo bar" + "\n\t")) + ("" + ("foo bar" + ("" + ("\n\t" + "foo bar")))))) + ("a" + "a")) + (("\n\t" + (("hello" + "") + "a")) + "\n\t"))) {
        case "foo bar":
            var v11 float32 = 0.0
            _ = v11
        case "hello":
            v12 := GenS3{F0: -610, F1: 42, F2: 0}
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS3{F0: -1, F1: -2147483648, F2: -32768}
            v5 = (v5 - (^(v5)))
            _ = fi13
            _ = v14
        }
        fmt.Println(v4)
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v8, fi15)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "\n\t" {
        case "foo bar":
            v16 := GenS3{F0: -285, F1: 175, F2: 425}
            _ = v16
        case "":
            var v17 int64 = 979
            _ = v17
        }
        _ = fi15
    }
    switch ("a" + ("\n\t" + (("foo bar" + ("\n\t" + "\n\t")) + "foo bar"))) {
    case "\n\t":
        v18 := GenS4{F0: "foo bar", F1: 1000, F2: 42}
        _ = v18
    case "hello":
        v8 -= 0
    }
    _ = len(v2)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            switch "" {
            case "a":
                var v20 uint16 = 255
                _ = v20
            case "\n\t":
                if false {
                    if (v8 > v8) {
                        v3[(fi19 | 32767)] = (-(fi19))
                        var v21 uint8 = 20
                        _ = v21
                    } else {
                        v5 = (v5 << ((0) & 63))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v22 := GenS4{F0: ((((((((("" + "") + "foo bar") + ("foo bar" + "hello")) + "foo bar") + "a") + "a") + ("a" + "")) + (("a" + "foo bar") + "")) + ("a" + (("hello" + (("" + "a") + "foo bar")) + ("" + (("\n\t" + "") + ("\n\t" + "hello")))))), F1: -158, F2: 65535}
                        _ = v22
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi23 := 0; fi23 < 9; fi23++ {
                        v5 = (v5 >> (((v5 / (((^(v5))) | 1))) & 63))
                        if true {
                            v5 += (v5 - (^(v5)))
                            fmt.Println(fi23)
                            v24 := GenS1{F0: 100}
                            fmt.Println(v5, v8)
                            _ = v24
                        } else {
                            _ = len(v2)
                            delete(v3, (fi19 - -569))
                            fmt.Println(v4, fi19, v4)
                        }
                        _ = len(v2)
                        _ = fi23
                    }
                    fmt.Println(v4, v8)
                }
            }
            _ = fi19
        }
        delete(v3, (v4 * 930))
    }
    v8 -= 995
    v3[(v4 + v4)] = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
