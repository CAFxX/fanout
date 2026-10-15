package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 int64, p1 int64, p2 int64) uint32 {
    p2 -= (p2 >> ((962) & 63))
    if true {
        var v5 uint8 = 2
        var v6 uint8 = v5
        var v7 uint = 255
        _ = v5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 453
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    var v3 bool = (!false)
    var v4 bool = v3
    for fi8 := 0; fi8 < 14; fi8++ {
        _ = len(v1)
        for fi9 := 0; fi9 < 1; fi9++ {
            v2 = 90
            _ = fi9
        }
        _ = fi8
    }
    var v10 uint16 = 525
    v4 = (v10 >= v10)
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint8 = 128
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 65535
    v14 := v1[""]
    var v15 int = (^(v14))
    switch (v15 | (-(v14))) {
    case 128:
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                fmt.Println(v15)
                v18 := v1["foo bar"]
                v19 := GenS1{F0: 111, F1: (v14 | (v15 | (-(fi17))))}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 bool = (!false)
            _ = fi16
            _ = v20
        }
    case 32767:
        var v21 uint = (v2 * v2)
        _ = v21
    case 42:
        fmt.Println(v4, v14, v15)
    default:
        switch (v14 << (((v2 ^ (v2 ^ v2))) & 63)) {
        case 4294967295:
            v22 := []int{}
            _ = v22
        case 42:
            switch "foo bar" {
            case "foo bar":
                delete(v1, "hello")
            case "hello":
                v23 := GenS1{F0: uint64(v13), F1: v15}
                _ = v23
            case "":
                var v24 bool = ((1.5 * 0.1) > -2.25)
                _ = v24
            default:
                if true {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        for fi26 := 0; fi26 < 16; fi26++ {
                            v27 := []int{(v14 + fi25), (^(fi26)), fi25, (fi26 ^ (fi25 * fi25))}
                            v10 = v13
                            v28 := GenS2{F0: 1, F1: (v2 * v2)}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        var v29 float32 = 0.1
                        _ = fi25
                        _ = v29
                    }
                    fmt.Println(v15, v14)
                    switch "" {
                    case "\n\t":
                        fmt.Println(v13)
                    case "hello":
                        v15 = 0
                    case "a":
                        v4 = false
                    }
                }
            }
        case -1:
            v14 -= -1
        default:
            if v4 {
                v14 -= (v14 | 32768)
            } else {
                v15 = v14
                for fi30 := 0; fi30 < 12; fi30++ {
                    switch ("" + "") {
                    case "a":
                        fmt.Println(v3)
                    case "foo bar":
                        var v31 float32 = 3.14159
                        _ = v31
                    case "hello":
                        v1["a"] = (v15 & 1)
                    default:
                        var v32 int32 = 2147483647
                        _ = v32
                    }
                    switch ("hello" + "") {
                    case "\n\t":
                        var v33 int8 = -128
                        _ = v33
                    case "hello":
                        v34 := GenS3{F0: int32(v10), F1: 100}
                        _ = v34
                    default:
                        v35 := GenS4{F0: 127, F1: 4294967295}
                        _ = v35
                    }
                    _ = fi30
                }
                var v36 uint64 = 474
                _ = v36
            }
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 uint16 = 827
        switch "hello" {
        case "foo bar":
            var v39 string = "foo bar"
            _ = v39
        case "hello":
            fmt.Println(v3)
        case "":
            switch "hello" {
            case "":
                delete(v1, "hello")
            }
        }
        _ = fi37
        _ = v38
    }
    v4 = (!(("" + ("foo bar" + "foo bar")) == "foo bar"))
    v40 := GenS5{F0: 668, F1: "hello"}
    var v41 float64 = (0.1 * (0.1 / ((((1.5 / 1.0) / -2.25) + -2.25) + (float64(0.1) - 0.1))))
    if (!v3) {
        switch 32768 {
        case 42:
            var v42 uint8 = 146
            _ = v42
        case -1:
            fmt.Println(v3, v2, v41)
        }
        switch (("foo bar" + "") + ("" + "hello")) {
        case "a":
            v15 += v14
        case "foo bar":
            var v43 int = v14
            _ = v43
        case "":
            var v44 uint32 = 973
            _ = v44
        }
        if (!((-2.25 - 0.0) > 1e10)) {
            var v45 float64 = 3.14159
            v14 += -9223372036854775808
            for fi46 := 0; fi46 < 7; fi46++ {
                var v47 int64 = 0
                _ = fi46
                _ = v47
            }
            _ = v45
        }
    } else {
        switch (v14 - v15) {
        case 1:
            v48 := GenS6{F0: 1, F1: v2}
            _ = v48
        default:
            var v49 int = (v15 | (v14 & (^(v15))))
            _ = v49
        }
    }
    fmt.Println(v4)
    var v50 int8 = -128
    var v51 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v50
    _ = v51
}
