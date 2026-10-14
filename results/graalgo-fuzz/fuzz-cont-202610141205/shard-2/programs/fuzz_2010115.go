package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0() int8 {
    if (float32((float64(1.5) - 0.0)) != 3.14159) {
        var v6 int = 1
        var v7 int16 = -202
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                var v9 float64 = 1.5
                v10 := GenS1{F0: 31}
                fmt.Println(fi8, v6)
                _ = v10
                _ = v9
            }
            _ = fi8
        }
        _ = v6
        _ = v7
    }
    v11 := []int{232, 32768}
    v12 := GenS1{F0: 151}
    if len(v11) > 0 {
        v11[0] = 1
    }
    _ = v11
    _ = v12
    return -29
}
func genFn1(p0 uint8, p1 int, p2 int16) uint8 {
    v13 := []int{303, (p1 * (-(p1)))}
    switch (("\n\t" + "foo bar") + ("foo bar" + "")) {
    case "\n\t":
        if true {
            p1 = p1
        } else {
            v14 := GenS1{F0: 32767}
            _ = v14
        }
    case "":
        switch "hello" {
        case "hello":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case "foo bar":
            p0 -= p0
        case "":
            v15 := GenS2{F0: -32768}
            _ = v15
        default:
            var v16 uint64 = 995
            _ = v16
        }
    case "foo bar":
        v17 := GenS1{F0: 2}
        _ = v17
    }
    switch "a" {
    case "a":
        var v18 int32 = 2147483647
        _ = v18
    case "":
        switch ("a" + ("\n\t" + "")) {
        case "hello":
            for fi19 := 0; fi19 < 12; fi19++ {
                switch ((("\n\t" + "") + "foo bar") + "") {
                case "\n\t":
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        var v21 uint = 42
                        var v22 uint64 = 4294967295
                        v21 = v21
                        if (-2147483648 < 792) {
                            switch -2147483648 {
                            case -1:
                                var v23 int = -43
                                _ = v23
                            case 2147483647:
                                var v24 string = ("" + ("a" + "a"))
                                _ = v24
                            }
                            p0 += (^(p0))
                        }
                        var v25 uint32 = 48
                        var v26 int32 = -698
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v25
                        _ = v26
                    }
                case "":
                    v27 := GenS3{F0: (false || true), F1: p0}
                    _ = v27
                }
                switch "\n\t" {
                case "hello":
                    var v28 float64 = (((float64(1e10) / 1.0) - 0.1) * 3.14159)
                    _ = v28
                }
                _ = fi19
            }
        case "a":
            fmt.Println(p0, p1, p2)
        case "":
            if true {
                v13 = append(v13, 350)
            } else {
                v29 := GenS1{F0: 255}
                fmt.Println(p2, p1)
                fmt.Println(p0, p1, p2)
                _ = v29
            }
        }
    case "foo bar":
        var v30 int8 = 100
        _ = v30
    default:
        fmt.Println(p0)
    }
    var v31 uint32 = 2
    var v32 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v31
    _ = v32
    return p0
}

func main() {
    var v1 int8 = 100
    var v2 uint64 = 4294967295
    v3 := []int{128, 2147483647}
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 uint64 = 18446744073709551615
        delete(v5, 0)
        _ = fi33
        _ = v34
    }
    var v35 string = (("a" + "") + "\n\t")
    var v36 bool = true
    v1 = genFn0()
    fmt.Println(v1, v2, v1)
    var v37 uint16 = 71
    var v38 int = 1
    v39 := v4["\n\t"]
    _ = len(v3)
    var v40 int = v38
    v39 = (((v35 + (v35 + "hello")) + (v35 + (v35 + ("" + ((v35 + v35) + (("" + v35) + (v35 + v35))))))) + ("a" + v39))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
