package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 int8; F2 bool }
func genFn0(p0 string) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 = ("a" + (p0 + p0))
        _ = fi7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        switch (fi8 - fi8) {
        case 32768:
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 int16 = -50
                v11 := GenS1{F0: -406}
                p0 = p0
                var v12 float64 = 0.1
                v12 = ((v12 - v12) - float64(0.1))
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        v13 := GenS1{F0: 9223372036854775807}
        switch (^(fi8)) {
        case -545:
            var v14 float32 = float32(float64(-2.25))
            _ = v14
        case 127:
            for fi15 := 0; fi15 < 13; fi15++ {
                p0 = (((p0 + "") + (p0 + ((("foo bar" + ("" + p0)) + ("hello" + p0)) + p0))) + p0)
                v16 := []int{32768, -1}
                if len(v16) > 0 {
                    v16[0] = -376
                }
                _ = fi15
                _ = v16
            }
        }
        _ = fi8
        _ = v13
    }
    switch 108 {
    case -227:
        switch 637 {
        case -1:
            var v17 uint64 = 977
            _ = v17
        case -2147483648:
            fmt.Println(p0)
        case -32:
            fmt.Println(p0)
        }
    case 32768:
        var v18 int = -1
        _ = v18
    default:
        v19 := GenS1{F0: 42}
        _ = v19
    }
    p0 = "foo bar"
    if ((-940 < 1) || (false || false)) {
        if (!(42 >= 255)) {
            if true {
                p0 = p0
            }
            v20 := GenS2{F0: 32767, F1: -1}
            switch 127 {
            case 173:
                p0 = (((p0 + (p0 + p0)) + p0) + p0)
            }
            _ = v20
        } else {
            v21 := GenS3{F0: 4294967295, F1: false, F2: 0}
            v22 := []int{127, -1, -163, 128}
            _ = v21
            _ = v22
        }
    }
    _ = p0
    return -32768
}
func genFn1(p0 int64, p1 uint, p2 uint64) uint32 {
    fmt.Println(p2)
    if (!false) {
        if true {
            p2 -= (p2 | (p2 >> ((p1) & 63)))
        }
        fmt.Println(p1, p0)
        v23 := GenS4{F0: 91}
        _ = v23
    } else {
        v24 := GenS1{F0: p0}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    var v4 int16 = int16(v1)
    v5 := []string{"hello", "hello", ("a" + ""), "foo bar"}
    v6 := []uint32{2147483647, 42}
    fmt.Println(v4, v1)
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := GenS5{F0: 2}
        if ("a" <= ("" + "")) {
            for fi27 := 0; fi27 < 20; fi27++ {
                if false {
                    fmt.Println(v4)
                }
                v28 := GenS6{F0: 2}
                var v29 bool = (42 != 0)
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            v4 = 1000
            v1 = v1
        }
        var v30 string = (("" + "hello") + "hello")
        _ = fi25
        _ = v26
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 uint = v1
        v2[950] = uint32(v1)
        var v33 int32 = 0
        _ = fi31
        _ = v32
        _ = v33
    }
    if ((0.0 < ((0.1 - (((1e10 + 1e10) * ((3.14159 / (3.14159 - float64(0.1))) - (3.14159 / ((0.0 / 1.5) - 1e10)))) + 0.1)) / 1.5)) && true) {
        switch 9223372036854775807 {
        case -1:
            v34 := GenS1{F0: 9223372036854775807}
            _ = v34
        case -9223372036854775808:
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 uint16 = 42
                _ = fi35
                _ = v36
            }
        default:
            fmt.Println(v4, v4)
        }
        v1 -= (v1 ^ (v1 + (v1 << ((v1) & 63))))
        v1 = (^(v1))
    }
    switch "" {
    case "\n\t":
        switch 1 {
        case 127:
            fmt.Println(v4, v4)
        case 0:
            fmt.Println(v4, v1)
        }
    case "hello":
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 float64 = (0.0 / 1.0)
            var v39 uint16 = 255
            _ = fi37
            _ = v38
            _ = v39
        }
    case "foo bar":
        v40 := GenS7{F0: (v4 & genFn0("hello"))}
        _ = v40
    }
    if false {
        if true {
            var v41 uint32 = 2
            var v42 int8 = -128
            _ = v41
            _ = v42
        } else {
            var v43 float32 = (0.1 - (1.5 * 3.14159))
            v44 := GenS2{F0: (v4 + v4), F1: -9223372036854775808}
            _ = v43
            _ = v44
        }
    }
    if (9223372036854775807 >= -277) {
        var v45 float32 = (-2.25 - 1.5)
        _ = v45
    }
    v1 -= v1
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 int8 = 86
        if (!false) {
            var v48 int = 4294967295
            fmt.Println(v4, v48)
            if true {
                for fi49 := 0; fi49 < 3; fi49++ {
                    var v50 float64 = 1e10
                    v47 = (v47 >> ((703) & 63))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi49
                    _ = v50
                }
                delete(v2, 715)
                v51 := GenS5{F0: 255}
                _ = v51
            }
            _ = v48
        } else {
            for fi52 := 0; fi52 < 13; fi52++ {
                v3 = append(v3, (fi46 & fi52))
                v47 = (^(v47))
                v53 := GenS8{F0: (v47 & v47), F1: int8(fi52), F2: true}
                _ = fi52
                _ = v53
            }
        }
        _ = fi46
        _ = v47
    }
    fmt.Println(v4, v1)
    switch (("hello" + "hello") + "a") {
    case "\n\t":
        for fi54 := 0; fi54 < 4; fi54++ {
            v55 := GenS6{F0: 42}
            _ = fi54
            _ = v55
        }
    case "hello":
        _ = len(v2)
    case "foo bar":
        switch 936 {
        case -912:
            v1 = v1
        case -1:
            v1 += v1
        case 32767:
            for fi56 := 0; fi56 < 9; fi56++ {
                v1 += (^(v1))
                _ = fi56
            }
        }
    default:
        v1 += 609
    }
    fmt.Println(v1)
    for fi57 := 0; fi57 < 3; fi57++ {
        for fi58 := 0; fi58 < 1; fi58++ {
            v4 = (v4 ^ v4)
            fmt.Println(v4, v4)
            var v59 uint16 = 1
            _ = fi58
            _ = v59
        }
        _ = fi57
    }
    for fi60 := 0; fi60 < 6; fi60++ {
        v4 -= 0
        _ = fi60
    }
    for fi61 := 0; fi61 < 7; fi61++ {
        _ = len(v2)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v62 uint32 = 2
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
