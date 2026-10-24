package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int; F2 uint16 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int32, p1 string, p2 int32) int8 {
    var v7 float32 = 0.1
    var v8 uint = 18446744073709551615
    if (!false) {
        v9 := GenS1{F0: 32767, F1: 9223372036854775807, F2: 1}
        var v10 uint16 = 0
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int64 = -553
    var v2 float64 = ((0.1 / 0.1) + float64(1.5))
    var v3 int64 = v1
    v4 := make(map[string]int)
    var v5 uint32 = 2147483647
    v6 := []string{"hello", ("hello" + ""), ("hello" + "foo bar")}
    for fi11 := 0; fi11 < 4; fi11++ {
        switch "foo bar" {
        case "":
            switch (^(fi11)) {
            case 32768:
                v12 := GenS1{F0: 474, F1: (fi11 / (((-(fi11))) | 1)), F2: -304}
                _ = v12
            }
        case "\n\t":
            switch "hello" {
            case "foo bar":
                if false {
                    v13 := v4[("a" + "hello")]
                    switch fi11 {
                    case -2147483648:
                        for fi14 := 0; fi14 < 4; fi14++ {
                            var v15 int8 = genFn0(-398, "a", 0)
                            var v16 int = 832
                            _ = len(v4)
                            v3 = (v1 << ((787) & 63))
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = v13
                } else {
                    v17 := GenS2{F0: v1}
                    _ = v17
                }
            case "a":
                var v18 int = fi11
                _ = v18
            }
        }
        v1 = v1
        _ = fi11
    }
    if len(v6) > 0 {
        v6[0] = ("foo bar" + (("a" + "") + (("hello" + "foo bar") + "a")))
    }
    if (!(false || false)) {
        v3 = (-(v3))
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS3{F0: false}
            var v21 string = "hello"
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    delete(v4, "")
    if len(v6) > 0 {
        v6[0] = "a"
    }
    v22 := GenS4{F0: 1, F1: 110, F2: 32767}
    v2 = ((0.0 / v2) - v2)
    delete(v4, "a")
    v1 -= (-(v1))
    v23 := v4["\n\t"]
    var v24 uint8 = 255
    if true {
        if (((v24 & v24) >= (v24 << ((42) & 63))) || true) {
            v3 = 42
            var v25 uint = 0
            _ = v25
        } else {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 string = "\n\t"
                _ = fi26
                _ = v27
            }
            fmt.Println(v5, v1, v24)
        }
    }
    v24 += v24
    var v28 int32 = 1
    fmt.Println(v28, v23, v5)
    for fi29 := 0; fi29 < 2; fi29++ {
        if (v23 >= (fi29 >> ((955) & 63))) {
            if true {
                var v30 uint64 = 933
                var v31 float64 = (v2 / 3.14159)
                _ = v30
                _ = v31
            }
            v32 := GenS5{F0: int8(fi29), F1: v1}
            _ = v32
        } else {
            switch "foo bar" {
            case "foo bar":
                var v33 float64 = v2
                _ = v33
            case "\n\t":
                switch "" {
                case "hello":
                    if true {
                        v24 = v24
                    }
                case "a":
                    var v34 int64 = -873
                    _ = v34
                }
            case "hello":
                if false {
                    v1 = v3
                }
            default:
                v6 = append(v6, "a")
            }
            switch fi29 {
            case 102:
                fmt.Println(v24)
            }
        }
        v35 := GenS6{F0: v5}
        var v36 uint32 = v5
        _ = fi29
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
