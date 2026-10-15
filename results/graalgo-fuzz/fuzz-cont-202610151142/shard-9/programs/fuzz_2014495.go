package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 int; F1 int32; F2 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 bool) int {
    var v5 int = 32768
    if true {
        switch v5 {
        case -585:
            if (true || (!(((v5 - (v5 >> ((0) & 63))) >= (v5 / ((v5) | 1))) && (!(2 == 42))))) {
                var v6 float64 = ((((0.1 + -2.25) * 0.1) - 0.1) - 3.14159)
                v7 := GenS1{F0: 735}
                _ = v6
                _ = v7
            }
        }
    } else {
        v8 := GenS2{F0: 14, F1: p0, F2: 127}
        v5 -= (v5 % (((v5 ^ (v5 | v5))) | 1))
        _ = v8
    }
    var v9 int8 = -128
    v9 = -71
    _ = p0
    _ = v5
    _ = v9
    return 15
}
func genFn1(p0 string, p1 int8, p2 uint8) int64 {
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(fi10, fi11)
            _ = fi11
        }
        p0 = "foo bar"
        if true {
            v12 := []int{genFn0(false)}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            p1 = (p1 * (p1 * p1))
            _ = v12
        } else {
            for fi13 := 0; fi13 < 11; fi13++ {
                p0 = p0
                if false {
                    var v14 uint = 1
                    p2 = p2
                    switch (^(fi13)) {
                    case -762:
                        switch p0 {
                        case "hello":
                            var v15 uint8 = (p2 % (((p2 - p2)) | 1))
                            _ = v15
                        }
                    }
                    _ = v14
                }
                var v16 bool = ((-2.25 / 1e10) == (1.5 * 0.1))
                _ = fi13
                _ = v16
            }
            p1 += 106
            for fi17 := 0; fi17 < 16; fi17++ {
                p1 = p1
                fmt.Println(p0, p0, p1)
                _ = fi17
            }
        }
        _ = fi10
    }
    switch -1 {
    case 4294967295:
        var v18 uint8 = (p2 / ((1) | 1))
        _ = v18
    default:
        v19 := []int{genFn0((!(!(true || false))))}
        _ = v19
    }
    p1 = p1
    v20 := GenS1{F0: 437}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return -365
}

func main() {
    v1 := []int{128}
    var v2 string = "hello"
    v3 := make(map[string]bool)
    v4 := []string{(("\n\t" + ("hello" + (v2 + v2))) + (v2 + ("" + "hello"))), "hello", v2}
    v2 = v2
    for fi21 := 0; fi21 < 10; fi21++ {
        v2 = ((("\n\t" + "") + v2) + "\n\t")
        fmt.Println(fi21, v2)
        fmt.Println(fi21, v2)
        _ = fi21
    }
    fmt.Println(v2)
    v2 = v2
    if len(v4) > 0 {
        v4[0] = (v2 + (((("hello" + v2) + ((v2 + ("a" + v2)) + v2)) + "foo bar") + (v2 + ("foo bar" + "\n\t"))))
    }
    fmt.Println(v2)
    fmt.Println(v2)
    if (-96 >= 128) {
        var v22 string = ""
        switch "" {
        case "a":
            v23 := GenS3{F0: -9223372036854775808, F1: 176, F2: 2}
            _ = v23
        default:
            var v24 uint64 = 924
            _ = v24
        }
        _ = v22
    } else {
        if false {
            switch v2 {
            case "a":
                v2 = (("a" + (v2 + v2)) + (v2 + (v2 + (v2 + (v2 + (v2 + v2))))))
            }
        } else {
            var v25 uint8 = 42
            _ = v25
        }
    }
    fmt.Println(v2)
    fmt.Println(v2)
    for fi26 := 0; fi26 < 17; fi26++ {
        if len(v1) > 0 {
            v1[0] = (-(fi26))
        }
        _ = fi26
    }
    fmt.Println(v2)
    v27 := GenS2{F0: -105, F1: (0 != 516), F2: 71}
    switch "" {
    case "a":
        var v28 uint8 = 255
        _ = v28
    case "":
        switch ((v2 + ((v2 + v2) + "\n\t")) + (((v2 + v2) + (v2 + v2)) + v2)) {
        case "foo bar":
            switch 127 {
            case 9223372036854775807:
                delete(v3, v2)
            case -9223372036854775808:
                fmt.Println(v2)
            default:
                v29 := GenS4{F0: -951}
                _ = v29
            }
        default:
            v2 = v2
        }
    case "foo bar":
        switch v2 {
        case "\n\t":
            v2 = (v2 + "\n\t")
        case "":
            v3[v2] = false
        case "a":
            v30 := GenS4{F0: 502}
            _ = v30
        default:
            v2 = v2
        }
    }
    v31 := GenS3{F0: genFn0((42 == 1)), F1: 2147483647, F2: 746}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
