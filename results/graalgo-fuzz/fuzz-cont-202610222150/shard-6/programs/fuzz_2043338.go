package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 int32 }
func genFn0() uint {
    var v6 float32 = 1e10
    switch 462 {
    case 2147483647:
        switch ("a" + ("" + "")) {
        case "":
            fmt.Println(v6)
        default:
            v6 = v6
        }
    case -440:
        v6 += v6
    case 1:
        var v7 uint = 65535
        _ = v7
    }
    switch ("" + "") {
    case "a":
        v8 := []int{}
        _ = v8
    case "foo bar":
        if true {
            v6 -= 3.14159
        } else {
            var v9 bool = true
            fmt.Println(v6)
            _ = v9
        }
    }
    if false {
        fmt.Println(v6)
        if false {
            for fi10 := 0; fi10 < 15; fi10++ {
                v6 = (float32((-2.25 / (float64(v6) / -2.25))) - float32((3.14159 + -2.25)))
                _ = fi10
            }
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int32 = -2147483648
                v13 := []int{}
                for fi14 := 0; fi14 < 2; fi14++ {
                    v6 -= v6
                    fmt.Println(fi11)
                    var v15 string = ""
                    _ = fi14
                    _ = v15
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v16 := GenS1{F0: 255, F1: 42, F2: (("hello" + "\n\t") + ("\n\t" + ((("" + "foo bar") + (("a" + "\n\t") + "foo bar")) + "foo bar")))}
        _ = v16
    } else {
        v6 += v6
        var v17 bool = (42 != 1)
        var v18 int = -9223372036854775808
        _ = v17
        _ = v18
    }
    _ = v6
    return 65535
}
func genFn1(p0 bool) uint16 {
    switch "foo bar" {
    case "hello":
        p0 = true
    case "a":
        var v19 int16 = -212
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        switch "hello" {
        case "hello":
            for fi21 := 0; fi21 < 8; fi21++ {
                fmt.Println(fi20, fi21)
                _ = fi21
            }
        case "a":
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 uint = genFn0()
                v23 -= v23
                _ = fi22
                _ = v23
            }
        }
        p0 = (569 == -1)
        if (p0 && (821 == -1)) {
            var v24 uint32 = 336
            fmt.Println(p0, v24)
            _ = v24
        } else {
            switch fi20 {
            case 490:
                v25 := GenS1{F0: 42, F1: -128, F2: "\n\t"}
                _ = v25
            default:
                fmt.Println(fi20, fi20)
            }
        }
        _ = fi20
    }
    fmt.Println(p0)
    var v26 int32 = 752
    _ = p0
    _ = v26
    return 0
}

func main() {
    var v1 int32 = -91
    v2 := make(map[string]uint32)
    var v3 uint32 = 2
    v4 := make(map[int]string)
    var v5 int8 = 9
    v27 := []int{-182, -2147483648, int(v5)}
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            fmt.Println(v3, fi29, v3)
            v3 = v3
            fmt.Println(fi28, fi28)
            var v30 float64 = -2.25
            _ = fi29
            _ = v30
        }
        _ = len(v27)
        _ = fi28
    }
    if (v3 != v3) {
        fmt.Println(v5, v3, v3)
        v5 += (v5 >> ((0) & 63))
        var v31 int64 = 175
        _ = v31
    }
    fmt.Println(v1, v5)
    var v32 uint32 = (^(v3))
    v2["foo bar"] = v32
    switch ("\n\t" + "") {
    case "\n\t":
        v32 = v3
    }
    if true {
        switch 1 {
        case -180:
            v33 := GenS1{F0: 0, F1: 27, F2: "foo bar"}
            _ = v33
        default:
            v32 -= 244
        }
        v32 = v3
        for fi34 := 0; fi34 < 20; fi34++ {
            if true {
                v3 += (v3 % (((v3 + v3)) | 1))
            }
            v35 := GenS2{F0: 0}
            var v36 uint64 = 18446744073709551615
            _ = fi34
            _ = v35
            _ = v36
        }
    } else {
        for fi37 := 0; fi37 < 20; fi37++ {
            v32 = v32
            _ = fi37
        }
    }
    var v38 int16 = -1
    fmt.Println(v5)
    switch 2147483647 {
    case 329:
        var v39 uint = 1
        _ = v39
    case -479:
        var v40 uint64 = 951
        _ = v40
    }
    _ = len(v4)
    var v41 uint16 = genFn1((v1 == (^(v1))))
    v38 -= v38
    v27 = append(v27, -597)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
