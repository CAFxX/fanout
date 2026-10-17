package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint16, p1 string, p2 int) uint16 {
    if ((255 < 250) && (false || true)) {
        if (((float32(3.14159) - -2.25) * 0.1) == 1.5) {
            if true {
                p1 = p1
                var v6 float64 = 3.14159
                _ = v6
            }
            v7 := GenS1{F0: 128, F1: p2}
            _ = v7
        } else {
            if ((42 == 1) || true) {
                v8 := []int{(-(p2)), (p2 / ((32767) | 1)), p2}
                if len(v8) > 0 {
                    v8[0] = p2
                }
                _ = v8
            }
            fmt.Println(p1, p0, p1)
        }
        v9 := []int{(-(p2)), 0}
        _ = v9
    } else {
        var v10 float32 = (1e10 + 3.14159)
        var v11 uint32 = 0
        p1 = p1
        _ = v10
        _ = v11
    }
    if (true || false) {
        v12 := []int{(p2 & 32768), -122, (p2 << ((0) & 63))}
        _ = v12
    } else {
        p0 = p0
        switch p1 {
        case "":
            var v13 bool = true
            _ = v13
        case "foo bar":
            var v14 uint16 = (p0 >> ((18446744073709551615) & 63))
            _ = v14
        case "hello":
            p2 = (p2 * p2)
        }
    }
    v15 := []int{p2, (p2 % (((-(p2))) | 1))}
    for fi16 := 0; fi16 < 9; fi16++ {
        if false {
            for fi17 := 0; fi17 < 3; fi17++ {
                p0 += (^(p0))
                p1 = (("" + p1) + "foo bar")
                v18 := GenS2{F0: 255, F1: 4294967295, F2: -9223372036854775808}
                v19 := GenS1{F0: 255, F1: fi16}
                var v20 int32 = 0
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            if (4294967295 > 269) {
                if false {
                    _ = len(v15)
                }
                var v21 string = (p1 + ((p1 + "foo bar") + p1))
                p0 = 0
                _ = v21
            } else {
                var v22 uint32 = 844
                _ = v22
            }
        }
        v23 := GenS1{F0: 103, F1: (fi16 % (((fi16 << ((2) & 63))) | 1))}
        _ = fi16
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p0 | 255)
}

func main() {
    v1 := []int{128}
    v2 := make(map[uint32]string)
    var v3 uint16 = 32767
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    switch -1 {
    case 42:
        v3 = genFn0(v3, "\n\t", 2147483647)
    case 465:
        v24 := v2[2]
        _ = v24
    case 702:
        var v25 int = 127
        _ = v25
    default:
        if false {
            switch 4294967295 {
            case 42:
                var v26 float32 = (-2.25 * 3.14159)
                _ = v26
            }
        }
    }
    if true {
        switch 0 {
        case -147:
            switch ("hello" + ("hello" + ("" + (((("a" + "a") + ("\n\t" + "a")) + "") + ("a" + "hello"))))) {
            case "foo bar":
                switch "hello" {
                case "hello":
                    var v27 uint = 65535
                    _ = v27
                case "a":
                    delete(v2, 2147483647)
                case "foo bar":
                    for fi28 := 0; fi28 < 14; fi28++ {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            delete(v2, 0)
                            var v30 float64 = 3.14159
                            if len(v1) > 0 {
                                v1[0] = (fi29 + 1)
                            }
                            _ = fi29
                            _ = v30
                        }
                        _ = fi28
                    }
                default:
                    fmt.Println(v3)
                }
            case "\n\t":
                v5 = v5
            }
        case 128:
            var v31 uint32 = 4294967295
            _ = v31
        case 0:
            var v32 int16 = 162
            _ = v32
        }
        var v33 uint64 = (^(v5))
        if true {
            var v34 float64 = (1e10 * -2.25)
            var v35 uint = 467
            _ = v34
            _ = v35
        } else {
            v36 := GenS3{F0: 110, F1: 2147483647}
            v3 = (v3 % ((genFn0(genFn0((v3 / ((2) | 1)), "foo bar", 4294967295), "", -35)) | 1))
            v3 -= genFn0(genFn0(v3, ("" + ("hello" + ("hello" + "foo bar"))), 9223372036854775807), (("hello" + "hello") + "a"), -104)
            _ = v36
        }
        _ = v33
    }
    if false {
        v37 := GenS1{F0: 207, F1: 2147483647}
        for fi38 := 0; fi38 < 16; fi38++ {
            switch (fi38 / (((fi38 / ((617) | 1))) | 1)) {
            case 4294967295:
                var v39 uint16 = genFn0(v3, "hello", fi38)
                _ = v39
            default:
                fmt.Println(fi38)
            }
            var v40 int32 = 2147483647
            _ = fi38
            _ = v40
        }
        var v41 int32 = -2147483648
        _ = v37
        _ = v41
    }
    fmt.Println(v3)
    v42 := GenS3{F0: 42, F1: 65535}
    v2[820] = "a"
    v5 += v5
    fmt.Println(v5, v3)
    for fi43 := 0; fi43 < 15; fi43++ {
        fmt.Println(fi43, v5, fi43)
        v5 = v5
        if (!(!false)) {
            for fi44 := 0; fi44 < 10; fi44++ {
                v45 := GenS4{F0: 1}
                v3 += (v3 + (v3 * v3))
                v5 += (^(v5))
                _ = fi44
                _ = v45
            }
        }
        _ = fi43
    }
    var v46 uint64 = (v5 % ((v5) | 1))
    v5 += (v5 % ((v5) | 1))
    var v47 uint32 = 1
    v47 = v47
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v5
}
