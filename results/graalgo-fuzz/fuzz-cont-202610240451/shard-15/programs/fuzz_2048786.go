package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 string; F2 int8 }
func genFn0(p0 bool, p1 uint, p2 int) bool {
    v5 := GenS1{F0: 1, F1: 609, F2: 42}
    var v6 uint8 = 173
    switch (("foo bar" + "foo bar") + "") {
    case "foo bar":
        p1 = 1
    case "":
        switch p2 {
        case 127:
            var v7 int32 = -938
            _ = v7
        case 1:
            p0 = p0
        default:
            if false {
                if p0 {
                    v6 = (^(v6))
                    var v8 uint8 = 128
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            v11 := []int{p2, fi10, (-(fi10)), fi9}
                            p2 -= fi10
                            _ = fi10
                            _ = v11
                        }
                        v12 := []int{(-(fi9)), fi9, int(p1)}
                        _ = fi9
                        _ = v12
                    }
                    if (!(!(p0 || p0))) {
                        v6 += (v6 | v6)
                        v13 := []int{p2, p2, -9223372036854775808, 1}
                        if true {
                            _ = len(v13)
                        }
                        _ = v13
                    }
                }
                v14 := GenS1{F0: 1, F1: 42, F2: 681}
                _ = v14
            }
        }
    default:
        if (float64(3.14159) > ((0.0 + -2.25) * -2.25)) {
            for fi15 := 0; fi15 < 14; fi15++ {
                if (!(!p0)) {
                    v16 := GenS1{F0: 60, F1: 4294967295, F2: -925}
                    switch fi15 {
                    case 9223372036854775807:
                        var v17 float32 = (-2.25 / 0.1)
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 4; fi18++ {
                        p0 = p0
                        var v19 int16 = 0
                        var v20 uint = p1
                        v21 := GenS2{F0: 0, F1: 3, F2: p0}
                        if p0 {
                            fmt.Println(p2, v6, p1)
                            v6 = 42
                            v22 := []int{fi18, (-(fi15)), (p2 >> ((p1) & 63))}
                            var v23 uint64 = 42
                            p0 = p0
                            var v24 int16 = (-(v19))
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = v16
                }
                _ = fi15
            }
            switch "foo bar" {
            case "hello":
                var v25 int32 = -915
                _ = v25
            case "\n\t":
                v26 := GenS2{F0: 42, F1: 65535, F2: p0}
                _ = v26
            case "a":
                switch (("foo bar" + "\n\t") + "") {
                case "a":
                    v27 := GenS2{F0: 2, F1: 1, F2: (1 > 578)}
                    _ = v27
                case "hello":
                    v28 := []int{(-(p2))}
                    _ = v28
                }
            }
        }
    }
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 uint64 = 705
    var v2 uint64 = v1
    v3 := make(map[uint32]uint32)
    var v4 float64 = 0.0
    switch (("hello" + ("foo bar" + (("hello" + "\n\t") + "foo bar"))) + (("hello" + "hello") + "hello")) {
    case "foo bar":
        v1 = v1
    case "a":
        var v29 uint16 = 32767
        _ = v29
    }
    v30 := GenS1{F0: 9223372036854775807, F1: v1, F2: -699}
    _ = len(v3)
    v31 := GenS2{F0: 32767, F1: 42, F2: false}
    v32 := GenS1{F0: 1, F1: (v2 + (v2 ^ v2)), F2: 32767}
    v33 := GenS1{F0: 42, F1: (^(v2)), F2: -771}
    v34 := GenS3{F0: 2}
    if (v1 >= 1) {
        v4 = 1e10
        if (!(true && ("hello" <= "foo bar"))) {
            var v35 uint16 = uint16(v1)
            _ = len(v3)
            if (v2 != 799) {
                v36 := GenS3{F0: 255}
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 string = "hello"
                    v39 := GenS1{F0: 1, F1: (^(v2)), F2: -1}
                    var v40 int32 = -2147483648
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = v36
            } else {
                v1 = (v1 - 428)
            }
            _ = v35
        } else {
            fmt.Println(v1, v1)
            fmt.Println(v2, v2)
        }
        v4 += float64(0.1)
    } else {
        switch "\n\t" {
        case "hello":
            v41 := GenS1{F0: 393, F1: v2, F2: 308}
            _ = v41
        case "":
            var v42 uint8 = 255
            _ = v42
        default:
            v1 = v1
        }
        switch "a" {
        case "":
            v2 -= (v2 % ((709) | 1))
        default:
            v43 := GenS4{F0: 2147483648, F1: "foo bar", F2: 119}
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
