package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := []uint32{42, 84, 2}
    var v2 uint8 = 0
    var v3 int8 = -51
    v4 := []uint32{0, 65535, 4294967295, 4294967295}
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int32 = 2147483647
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        _ = fi6
        _ = v7
    }
    var v8 float64 = 1.5
    delete(v5, "a")
    v3 -= 1
    v2 += (^(v2))
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v8, v8, v3)
        if true {
            var v10 string = ""
            v11 := GenS1{F0: 42, F1: 122}
            _ = v10
            _ = v11
        } else {
            v2 += (v2 | (v2 ^ v2))
        }
        _ = fi9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            switch (fi13 << ((457) & 63)) {
            case -9223372036854775808:
                v14 := GenS1{F0: 52, F1: -235}
                _ = len(v5)
                v8 = 0.0
                _ = v14
            case 769:
                if len(v1) > 0 {
                    v1[0] = 47
                }
                var v15 int64 = 42
                v8 -= v8
                if true {
                    var v16 uint8 = v2
                    if false {
                        var v17 int8 = (-(v3))
                        _ = v17
                    }
                    delete(v5, ("a" + ""))
                    _ = v16
                }
                delete(v5, "hello")
                _ = v15
            case 9223372036854775807:
                v18 := GenS1{F0: -1, F1: -9223372036854775808}
                _ = v18
            default:
                v2 = 148
            }
            _ = fi13
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint64 = 144
            _ = fi19
            _ = v20
        }
        var v21 float64 = ((0.0 * v8) / (float64(((0.1 * (float32(v8) + 0.0)) * 0.1)) - 1.5))
        _ = fi12
        _ = v21
    }
    v3 = (v3 % ((1) | 1))
    if false {
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v2 += 255
                var v24 uint32 = 4294967295
                if ((!(-2147483648 >= 402)) || true) {
                    v24 -= (v24 & (^(v24)))
                    switch (fi23 * (fi22 * (fi23 | fi23))) {
                    case -2147483648:
                        switch "a" {
                        case "":
                            v24 = v24
                        case "a":
                            var v25 float32 = 1e10
                            _ = v25
                        }
                    case 42:
                        v26 := GenS2{F0: 366, F1: (true && false)}
                        _ = v26
                    }
                } else {
                    var v27 float32 = -2.25
                    _ = v27
                }
                _ = fi23
                _ = v24
            }
            v8 = (v8 * ((v8 - (float64(((0.1 + 0.0) * -2.25)) * v8)) * (v8 * float64((float32(v8) - (0.0 / 0.1))))))
            for fi28 := 0; fi28 < 5; fi28++ {
                fmt.Println(v8)
                v29 := GenS1{F0: 42, F1: 9223372036854775807}
                _ = fi28
                _ = v29
            }
            _ = fi22
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                v3 = (-(v3))
                var v32 float64 = 0.1
                _ = fi31
                _ = v32
            }
            var v33 uint = 4294967295
            switch "hello" {
            case "":
                for fi34 := 0; fi34 < 4; fi34++ {
                    switch (("hello" + "hello") + ((("\n\t" + "a") + ("a" + "a")) + "a")) {
                    case "\n\t":
                        switch (("hello" + "foo bar") + ("hello" + ("\n\t" + ""))) {
                        case "foo bar":
                            var v35 float64 = (((3.14159 + 1e10) + v8) * (0.1 * float64(0.1)))
                            _ = v35
                        }
                    }
                    v5[(("a" + "foo bar") + "foo bar")] = (!((-(v3)) <= 127))
                    _ = fi34
                }
            case "foo bar":
                v33 = (v33 - (v33 ^ v33))
            default:
                v36 := GenS2{F0: 1, F1: true}
                _ = v36
            }
            _ = fi30
            _ = v33
        }
        v37 := GenS2{F0: 3, F1: false}
        _ = v37
    } else {
        v2 += v2
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        for fi39 := 0; fi39 < 20; fi39++ {
            if true {
                var v40 uint16 = 1
                _ = v40
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (801 > 4294967295) {
                    fmt.Println(v2)
                }
                v2 -= v2
            }
            _ = len(v1)
            v41 := GenS3{F0: 2147483648}
            _ = fi39
            _ = v41
        }
        delete(v5, ((("" + "a") + "") + ("foo bar" + "\n\t")))
        _ = fi38
    }
    var v42 float32 = 1.5
    delete(v5, "\n\t")
    switch -9223372036854775808 {
    case 32768:
        var v43 string = ("foo bar" + "")
        _ = v43
    case 42:
        var v44 bool = (126 < 0)
        _ = v44
    case 2147483647:
        var v45 uint32 = 42
        _ = v45
    default:
        var v46 int64 = 1
        _ = v46
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 0 {
    case 171:
        v47 := GenS4{F0: true}
        _ = v47
    case 42:
        _ = len(v4)
    default:
        switch "\n\t" {
        case "foo bar":
            v8 += (float64(-2.25) / 0.1)
        }
    }
    var v48 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v8
}
