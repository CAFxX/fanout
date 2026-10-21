package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
func genFn0(p0 int64, p1 int) uint8 {
    fmt.Println(p1)
    v4 := []int{}
    switch p1 {
    case 4294967295:
        if ((-(p0)) == (p0 % (((p0 + p0)) | 1))) {
            var v5 float64 = 1e10
            _ = v5
        }
    default:
        switch "foo bar" {
        case "a":
            var v6 uint8 = 0
            _ = v6
        case "hello":
            var v7 float32 = -2.25
            _ = v7
        default:
            switch (("\n\t" + "") + "\n\t") {
            case "\n\t":
                var v8 int = -537
                _ = v8
            case "":
                fmt.Println(p1)
            case "foo bar":
                for fi9 := 0; fi9 < 6; fi9++ {
                    p1 -= p1
                    v10 := GenS1{F0: false, F1: 255, F2: "hello"}
                    _ = fi9
                    _ = v10
                }
            }
        }
    }
    fmt.Println(p0, p1)
    if (42 <= 2) {
        p1 += p1
    } else {
        if false {
            p1 -= p1
            switch "" {
            case "":
                if true {
                    if true {
                        for fi11 := 0; fi11 < 3; fi11++ {
                            p1 = fi11
                            var v12 uint64 = 4294967295
                            _ = fi11
                            _ = v12
                        }
                        p1 += (^(p1))
                    } else {
                        fmt.Println(p0, p0)
                        switch p1 {
                        case 32767:
                            p1 = p1
                        case 238:
                            var v13 bool = true
                            _ = v13
                        }
                    }
                } else {
                    v14 := GenS1{F0: (689 != 18446744073709551615), F1: 255, F2: "hello"}
                    if len(v4) > 0 {
                        v4[0] = p1
                    }
                    v15 := GenS1{F0: (!true), F1: 156, F2: ("" + "hello")}
                    _ = v14
                    _ = v15
                }
            default:
                v16 := GenS1{F0: (!(("" + "foo bar") <= ("hello" + ((("\n\t" + ("\n\t" + "a")) + "") + "\n\t")))), F1: 42, F2: ""}
                _ = v16
            }
            v17 := GenS2{F0: 100, F1: false, F2: 0}
            _ = v17
        } else {
            switch ("" + ("" + "foo bar")) {
            case "a":
                if false {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        p1 = (fi18 + p1)
                        _ = fi18
                    }
                }
            }
            if false {
                var v19 float64 = (((((3.14159 + 1.5) * 1e10) + (float64(float32(((1e10 / 0.1) - float64(1e10)))) + 1.5)) / float64(-2.25)) + -2.25)
                _ = v19
            }
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 128
}

func main() {
    var v1 float64 = 3.14159
    var v2 bool = (1 > 127)
    v3 := []uint32{}
    var v20 uint32 = 2147483647
    v2 = (!v2)
    v1 += 0.0
    for fi21 := 0; fi21 < 12; fi21++ {
        if (3.14159 > v1) {
            var v22 uint16 = 2
            _ = v22
        }
        v2 = v2
        _ = fi21
    }
    switch "a" {
    case "hello":
        v20 = v20
    default:
        var v23 bool = ((v20 >= (v20 | v20)) && (814 <= 42))
        _ = v23
    }
    var v24 uint32 = (v20 - 152)
    v25 := GenS3{F0: v2, F1: 125, F2: -1}
    fmt.Println(v1, v2, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
}
