package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 int, p1 string) uint8 {
    for fi5 := 0; fi5 < 9; fi5++ {
        switch (p1 + p1) {
        case "foo bar":
            var v6 uint8 = 0
            _ = v6
        case "hello":
            if (!(true || (!(255 != 4294967295)))) {
                if false {
                    var v7 float32 = -2.25
                    _ = v7
                }
                var v8 uint32 = 1
                p1 = ("a" + (p1 + p1))
                _ = v8
            } else {
                v9 := []int{42, fi5}
                _ = v9
            }
        }
        p1 = ""
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: 0, F1: 1, F2: 16}
            fmt.Println(p0, p0)
            _ = fi10
            _ = v11
        }
        _ = fi5
    }
    p1 = p1
    fmt.Println(p1, p1)
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v12
    return 128
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]uint32)
    var v3 int = 4294967295
    var v4 uint64 = 18446744073709551615
    if false {
        var v13 float64 = 1.5
        _ = v13
    } else {
        var v14 uint = 4294967295
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int64 = -9223372036854775808
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 string = "a"
            fmt.Println(fi15)
            _ = fi17
            _ = v18
        }
        switch "foo bar" {
        case "hello":
            fmt.Println(v1, v4)
        case "foo bar":
            v19 := GenS2{F0: -32768, F1: -1, F2: 173}
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    if false {
        if (v1 <= v1) {
            v20 := GenS3{F0: true, F1: 0}
            var v21 float64 = (float64(float32(((-2.25 / 1.5) - (0.1 + -2.25)))) + -2.25)
            v22 := GenS3{F0: (!(208 >= 65535)), F1: 79}
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 int16 = -1
        _ = v23
    }
    if (!((v1 << ((1) & 63)) >= v1)) {
        var v24 int16 = -177
        var v25 uint64 = 311
        delete(v2, "a")
        _ = v24
        _ = v25
    } else {
        v26 := []int{4294967295, v3, v3}
        var v27 float32 = 3.14159
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: (!("hello" < "hello")), F1: 255}
    var v29 float32 = 0.1
    v30 := []int{9223372036854775807, 790}
    v31 := GenS4{F0: -827}
    switch "\n\t" {
    case "":
        v32 := GenS3{F0: false, F1: 479}
        _ = v32
    case "\n\t":
        _ = len(v30)
    default:
        for fi33 := 0; fi33 < 5; fi33++ {
            fmt.Println(v1)
            _ = fi33
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        if (true && (42 > 4294967295)) {
            switch "a" {
            case "a":
                fmt.Println(fi34, v1, v4)
            case "":
                fmt.Println(v4)
            default:
                if false {
                    fmt.Println(v29)
                    var v35 bool = (true && (!(false && (true || (false || (false && false))))))
                    v36 := GenS5{F0: 42}
                    _ = v35
                    _ = v36
                } else {
                    if len(v30) > 0 {
                        v30[0] = (fi34 + (-(v3)))
                    }
                }
            }
        }
        _ = fi34
    }
    fmt.Println(v3, v4, v3)
    v37 := GenS6{F0: uint8(v4), F1: 42, F2: -24}
    if ((v29 != float32((((0.1 - (1.5 + -2.25)) - 0.1) + 1e10))) && ((("hello" + "hello") + "a") >= "hello")) {
        if (((0.1 * ((-2.25 + 1.5) / 0.1)) * (3.14159 + -2.25)) > float64(float32(((1e10 - (float64((float32(0.1) - (v29 / v29))) + ((((0.0 * 3.14159) * (0.0 + -2.25)) * ((1.5 - (-2.25 * 1e10)) / 1e10)) / 0.1))) - 0.0)))) {
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 bool = (true && true)
                _ = fi38
                _ = v39
            }
        } else {
            var v40 uint32 = v1
            _ = v40
        }
        v4 = (v4 | v4)
        var v41 int16 = 0
        _ = v41
    } else {
        var v42 int64 = 805
        _ = v42
    }
    var v43 int64 = 42
    var v44 bool = (65535 == 604)
    fmt.Println(v1, v43)
    var v45 string = ""
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
