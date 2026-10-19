package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int, p1 int16, p2 uint16) int16 {
    if false {
        p2 = 471
    }
    v5 := []int{p0}
    for fi6 := 0; fi6 < 18; fi6++ {
        v5 = append(v5, (-(p0)))
        fmt.Println(fi6)
        v7 := GenS1{F0: p2}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1() uint32 {
    if (440 > 682) {
        var v8 int = 128
        v8 += v8
        _ = v8
    } else {
        v9 := []int{42, 2147483647, 844}
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: 2}
        if true {
            switch "" {
            case "":
                var v12 bool = true
                _ = v12
            case "hello":
                var v13 float64 = (0.0 * 1.5)
                _ = v13
            default:
                var v14 float32 = (float32(1.5) / 3.14159)
                _ = v14
            }
        }
        _ = fi10
        _ = v11
    }
    return 2
}

func main() {
    var v1 int32 = 1
    var v2 int32 = -2147483648
    v3 := []uint32{2147483647}
    v4 := []string{"a"}
    switch 2147483647 {
    case 32768:
        var v15 int16 = 1000
        _ = v15
    case 4294967295:
        if true {
            var v16 bool = (true || true)
            for fi17 := 0; fi17 < 15; fi17++ {
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                _ = fi17
            }
            _ = v16
        } else {
            var v18 bool = false
            v2 -= (v1 & (v2 & (v1 * 1)))
            _ = v18
        }
    case 32767:
        var v19 string = "hello"
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch ((("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + ("foo bar" + "\n\t")) + "a") {
            case "a":
                _ = len(v4)
            case "foo bar":
                if true {
                    if (0.0 < ((0.0 * -2.25) - 0.1)) {
                        var v22 string = "\n\t"
                        _ = v22
                    } else {
                        var v23 bool = ((!(!false)) && (false && ((0 != genFn0((fi20 / ((fi20) | 1)), 42, 2)) && (630 < 2))))
                        v2 = v1
                        v1 += (v1 | v2)
                        _ = v23
                    }
                    var v24 string = "a"
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v24
                }
            case "\n\t":
                var v25 int32 = v2
                _ = v25
            }
            v2 -= (v2 / (((^(v1))) | 1))
            _ = fi21
        }
        var v26 bool = true
        _ = len(v4)
        _ = fi20
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 int8 = -1
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 uint16 = 369
        var v31 int32 = 0
        v32 := GenS2{F0: 1, F1: (v30 % (((v30 - v30)) | 1))}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println(v2, v2)
    v2 = (v1 | v2)
    if (42 > 32767) {
        v33 := GenS1{F0: 65535}
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(v2)
        _ = fi34
    }
    v1 += v1
    var v35 string = "\n\t"
    if false {
        switch 1 {
        case -72:
            switch 32768 {
            case 26:
                fmt.Println(v35, v1, v35)
            case -297:
                v36 := GenS3{F0: 222, F1: 42}
                _ = v36
            }
        }
        fmt.Println(v1)
    } else {
        switch 9223372036854775807 {
        case -930:
            fmt.Println(v35, v2, v1)
        default:
            for fi37 := 0; fi37 < 18; fi37++ {
                v35 = "foo bar"
                v38 := GenS4{F0: 283, F1: 42}
                v1 = v1
                _ = fi37
                _ = v38
            }
        }
    }
    var v39 uint8 = 182
    fmt.Println(v1, v39, v2)
    for fi40 := 0; fi40 < 17; fi40++ {
        fmt.Println(v35)
        v41 := GenS3{F0: 0, F1: (v39 + v39)}
        fmt.Println(fi40, v39)
        _ = fi40
        _ = v41
    }
    var v42 string = v35
    var v43 uint32 = 200
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v44 uint = 4294967295
        _ = v44
    } else {
        v1 += v2
        v42 = ("a" + "a")
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v42
    _ = v43
}
