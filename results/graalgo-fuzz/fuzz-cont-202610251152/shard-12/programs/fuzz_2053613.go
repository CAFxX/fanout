package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 bool }
func genFn0() int64 {
    switch ("hello" + "hello") {
    case "hello":
        if (2 <= 65535) {
            switch "hello" {
            case "":
                var v6 int8 = 1
                _ = v6
            case "\n\t":
                v7 := []int{-1}
                _ = v7
            }
            var v8 int = -2147483648
            v8 -= v8
            _ = v8
        }
    }
    switch -2147483648 {
    case 32768:
        var v9 float64 = -2.25
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint16 = 1
            _ = fi10
            _ = v11
        }
    default:
        var v12 uint = 540
        _ = v12
    }
    return -9223372036854775808
}
func genFn1(p0 int) uint {
    v13 := GenS1{F0: 216, F1: -434, F2: (p0 * (p0 + (p0 * 9223372036854775807)))}
    fmt.Println(p0)
    p0 = p0
    var v14 uint32 = 517
    v14 += v14
    _ = p0
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    var v1 int8 = 0
    v2 := []string{"hello", "a", ((("a" + "\n\t") + "\n\t") + "\n\t")}
    var v3 uint32 = 2
    var v4 float32 = 3.14159
    v5 := make(map[int]bool)
    if true {
        var v15 int = 829
        v1 = v1
        _ = v15
    } else {
        var v16 uint64 = 576
        v1 -= v1
        v17 := GenS1{F0: 549, F1: 9223372036854775807, F2: 9223372036854775807}
        _ = v16
        _ = v17
    }
    v3 = v3
    _ = len(v5)
    v18 := GenS1{F0: uint16(v3), F1: -9223372036854775808, F2: 42}
    delete(v5, 32768)
    v2 = append(v2, "")
    fmt.Println(v1, v4)
    v1 = (v1 + 1)
    if (0.1 < 1e10) {
        v19 := GenS2{F0: 1, F1: 42}
        var v20 uint64 = 781
        v4 -= 0.0
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 string = "\n\t"
        v23 := v5[(^(fi21))]
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 255, F1: 42}
    v25 := GenS1{F0: 498, F1: genFn0(), F2: 32767}
    fmt.Println(v4, v1, v1)
    if false {
        v5[32768] = ((!("a" >= "hello")) || true)
        for fi26 := 0; fi26 < 2; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 float64 = 1e10
                switch "foo bar" {
                case "hello":
                    switch ("a" + "foo bar") {
                    case "a":
                        fmt.Println(v28, v1, fi26)
                    }
                default:
                    if ((false && ((-(fi27)) > (fi27 ^ (fi27 % ((int(v1)) | 1))))) || ((915 != -514) && ((202 == 2) || true))) {
                        fmt.Println(v3)
                        var v29 string = "foo bar"
                        var v30 int64 = 9223372036854775807
                        fmt.Println(v1, v4, fi26)
                        v1 = (v1 * v1)
                        _ = v29
                        _ = v30
                    }
                }
                v28 -= (1.5 - 1e10)
                _ = fi27
                _ = v28
            }
            if (("a" + (("" + "foo bar") + "\n\t")) > "") {
                for fi31 := 0; fi31 < 5; fi31++ {
                    v4 += v4
                    _ = fi31
                }
                v32 := GenS3{F0: 42, F1: int64(v1), F2: false}
                var v33 int8 = v1
                _ = v32
                _ = v33
            } else {
                switch fi26 {
                case 32767:
                    v34 := GenS1{F0: 1, F1: 9223372036854775807, F2: 2147483647}
                    _ = v34
                default:
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
                var v35 string = ("foo bar" + (("hello" + "\n\t") + ""))
                _ = v35
            }
            _ = fi26
        }
    } else {
        var v36 uint = 18446744073709551615
        switch int(v36) {
        case 127:
            v37 := GenS1{F0: 65535, F1: genFn0(), F2: 682}
            _ = v37
        case 4294967295:
            switch ("hello" + ("\n\t" + "foo bar")) {
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                v2 = append(v2, ("hello" + "foo bar"))
            default:
                if len(v2) > 0 {
                    v2[0] = (("\n\t" + ("a" + (("" + "a") + "a"))) + ("\n\t" + ("foo bar" + ((("foo bar" + "foo bar") + "\n\t") + ("hello" + ("\n\t" + (("" + (("" + "hello") + "\n\t")) + "hello")))))))
                }
            }
        case 476:
            v38 := GenS1{F0: 255, F1: 42, F2: 127}
            _ = v38
        }
        _ = v36
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
