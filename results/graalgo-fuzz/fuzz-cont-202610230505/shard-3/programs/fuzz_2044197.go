package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 bool, p1 string, p2 bool) uint32 {
    p2 = (61 >= 0)
    switch p1 {
    case "foo bar":
        var v7 int16 = 1000
        _ = v7
    }
    if (0 != 128) {
        var v8 int64 = -438
        if p2 {
            var v9 uint64 = 814
            fmt.Println(p1, p1)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 float64 = ((((1e10 * (1.5 + 0.1)) * -2.25) - 0.1) - 1e10)
                fmt.Println(p2)
                v12 := GenS1{F0: "\n\t", F1: 146, F2: -1}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint = 893
            fmt.Println(p0, p2, v8)
            _ = v13
        }
        var v14 int8 = -1
        _ = v14
        _ = v8
    } else {
        var v15 uint8 = 1
        v16 := GenS2{F0: 0, F1: 9223372036854775807}
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint16) uint {
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 bool = true
        var v19 int = (fi17 % ((fi17) | 1))
        _ = fi17
        _ = v18
        _ = v19
    }
    switch 42 {
    case 127:
        if true {
            p0 = (p0 & p0)
            v20 := []int{}
            fmt.Println(p0)
            _ = v20
        } else {
            fmt.Println(p0)
            if (!false) {
                var v21 int16 = -1
                _ = v21
            }
            v22 := []int{-369, int(p0), 224}
            _ = v22
        }
    case 9223372036854775807:
        var v23 float64 = 1e10
        _ = v23
    default:
        fmt.Println(p0)
    }
    _ = p0
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{("" + (("" + "") + "foo bar")), ("hello" + ((("foo bar" + ("a" + (((("" + ("a" + ("" + ("\n\t" + ("a" + "a"))))) + "a") + "foo bar") + ""))) + (("a" + "\n\t") + ("a" + "a"))) + "foo bar")), ("hello" + "a")}
    v3 := make(map[uint32]int)
    var v4 uint8 = 127
    v5 := make(map[string]string)
    var v6 bool = (!false)
    fmt.Println(v4)
    switch "hello" {
    case "\n\t":
        switch 0 {
        case 1:
            var v24 uint32 = 2147483647
            _ = v24
        case 9223372036854775807:
            switch "\n\t" {
            case "a":
                fmt.Println(v6, v4)
            default:
                var v25 int8 = -79
                _ = v25
            }
        }
    case "":
        v4 += 255
    default:
        v26 := GenS2{F0: -9223372036854775808, F1: 0}
        _ = v26
    }
    _ = len(v5)
    var v27 float32 = 1.5
    v27 -= 1e10
    var v28 float32 = (v27 - v27)
    var v29 int32 = 1
    var v30 int32 = 1
    fmt.Println(v27, v29, v28)
    v31 := GenS1{F0: "foo bar", F1: 2, F2: 0}
    v30 = 2147483647
    switch "foo bar" {
    case "":
        if (-9223372036854775808 != 9223372036854775807) {
            if (-128 > 100) {
                switch "hello" {
                case "hello":
                    if v6 {
                        var v32 float32 = -2.25
                        if v6 {
                            var v33 bool = true
                            _ = v33
                        }
                        var v34 int = 891
                        _ = v32
                        _ = v34
                    }
                }
                v6 = false
                var v35 uint = 660
                _ = v35
            }
            v27 = v28
            var v36 uint32 = 65535
            _ = v36
        } else {
            var v37 float64 = 0.1
            _ = v37
        }
    case "foo bar":
        var v38 int16 = -32768
        _ = v38
    case "a":
        switch "foo bar" {
        case "":
            v3[genFn0(v6, "foo bar", (42 > 32767))] = 435
        case "\n\t":
            v39 := GenS3{F0: 42, F1: 32767}
            _ = v39
        case "foo bar":
            if v6 {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = len(v2)
            }
        default:
            v27 -= float32(1.5)
        }
    default:
        delete(v3, 1)
    }
    if (false || true) {
        delete(v3, 2147483648)
        v4 = 2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
