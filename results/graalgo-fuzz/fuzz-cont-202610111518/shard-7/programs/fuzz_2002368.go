package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 string; F2 uint16 }
type GenS7 struct { F0 int }

func main() {
    var v1 string = "hello"
    var v2 uint32 = 901
    var v3 uint = 255
    v4 := make(map[int]bool)
    var v5 int32 = 1
    if (((("a" + "hello") + v1) + v1) <= v1) {
        v6 := []int{32767, -423, -314}
        for fi7 := 0; fi7 < 6; fi7++ {
            switch ((v1 + v1) + (v1 + "\n\t")) {
            case "foo bar":
                v8 := GenS1{F0: true, F1: 32767, F2: int16(fi7)}
                _ = v8
            case "\n\t":
                if ((1e10 * float32(-2.25)) > 3.14159) {
                    v5 -= v5
                    if false {
                        v2 += v2
                        v9 := GenS2{F0: -643}
                        var v10 int64 = 42
                        _ = v10
                        _ = v9
                    }
                }
            case "a":
                var v11 int = (fi7 << ((v3) & 63))
                _ = v11
            default:
                v1 = "\n\t"
            }
            _ = fi7
        }
        var v12 uint32 = v2
        _ = v12
        _ = v6
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := []int{612}
        var v15 float64 = 1e10
        switch "\n\t" {
        case "\n\t":
            switch (-(fi13)) {
            case -228:
                v4[(fi13 & (fi13 << ((65535) & 63)))] = (false && (!(!true)))
            }
        case "foo bar":
            v2 = v2
        default:
            _ = len(v14)
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 -= (v2 & 25)
    switch 292 {
    case 583:
        _ = len(v4)
    default:
        var v16 uint64 = 466
        _ = v16
    }
    if ((("\n\t" + ("a" + ((v1 + "hello") + v1))) + (("foo bar" + v1) + v1)) != "hello") {
        v17 := GenS3{F0: v2, F1: v5, F2: 32767}
        if false {
            for fi18 := 0; fi18 < 6; fi18++ {
                fmt.Println(v5, v5)
                _ = fi18
            }
            v19 := []int{}
            v3 -= v3
            _ = v19
        }
        fmt.Println(v3)
        _ = v17
    }
    if true {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int16 = -32768
            _ = fi20
            _ = v21
        }
        var v22 float32 = (float32(((float64(3.14159) - (((-2.25 * ((((((0.1 * ((3.14159 / float64(3.14159)) * 0.1)) * ((0.0 * float64((1.5 * 0.0))) - (0.0 * 0.1))) + (1e10 + 1.5)) * float64((3.14159 + 3.14159))) / 1.0) + 0.1)) / 0.1) / 1.5)) / float64(0.1))) * (0.1 + (-2.25 / float32(((1.5 - -2.25) - (3.14159 + 1e10))))))
        _ = v22
    }
    if (true && false) {
        v3 = (^(v3))
        fmt.Println(v5)
        var v23 int = 390
        _ = v23
    }
    switch ((v1 + v1) + "") {
    case "":
        for fi24 := 0; fi24 < 4; fi24++ {
            if true {
                v25 := GenS2{F0: -807}
                v26 := GenS3{F0: uint32(v3), F1: 2147483647, F2: 2}
                v27 := GenS3{F0: (^(v2)), F1: (-(v5)), F2: 0}
                _ = v25
                _ = v26
                _ = v27
            } else {
                var v28 float32 = 0.1
                fmt.Println(v28)
                v29 := GenS4{F0: ((float32((0.1 * 3.14159)) + float32(0.1)) >= ((1e10 + 0.0) + ((v28 - v28) / v28))), F1: 0}
                _ = v28
                _ = v29
            }
            switch fi24 {
            case -877:
                var v30 string = v1
                _ = v30
            case -9223372036854775808:
                for fi31 := 0; fi31 < 15; fi31++ {
                    v32 := GenS4{F0: false, F1: -1}
                    _ = fi31
                    _ = v32
                }
            case -562:
                var v33 uint32 = v2
                _ = v33
            }
            _ = fi24
        }
    case "hello":
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 uint16 = 42
            for fi36 := 0; fi36 < 6; fi36++ {
                v37 := v4[(fi36 / ((42) | 1))]
                var v38 int64 = int64(v2)
                v39 := GenS5{F0: 567}
                var v40 int8 = 127
                var v41 float64 = 0.0
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi34
            _ = v35
        }
    }
    v5 += v5
    fmt.Println(v5, v1)
    v42 := GenS6{F0: 0, F1: v1, F2: uint16(v3)}
    var v43 uint8 = 1
    v44 := GenS7{F0: 636}
    var v45 uint32 = (v2 - (v2 & v2))
    var v46 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
