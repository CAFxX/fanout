package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 int }
type GenS4 struct { F0 int8; F1 int32; F2 int8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint; F1 bool }

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{}
    v3 := make(map[uint32]int)
    v4 := []string{}
    var v5 float64 = 0.1
    v6 := GenS1{F0: 32767}
    var v7 uint8 = 132
    v7 = (v7 | (^(v7)))
    for fi8 := 0; fi8 < 18; fi8++ {
        _ = len(v2)
        var v9 bool = true
        _ = fi8
        _ = v9
    }
    if false {
        var v10 int16 = -1
        if (int32(v10) < 42) {
            for fi11 := 0; fi11 < 9; fi11++ {
                v1 = v1
                _ = fi11
            }
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS2{F0: -389}
            if (148 == 623) {
                if false {
                    fmt.Println(v7)
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 uint64 = v1
                    _ = fi14
                    _ = v15
                }
            }
            v3[0] = (fi12 - (fi12 & 4294967295))
            _ = fi12
            _ = v13
        }
        _ = v10
    }
    if (1 != 32767) {
        if false {
            v16 := GenS3{F0: false, F1: 2147483648, F2: -4}
            for fi17 := 0; fi17 < 14; fi17++ {
                if (32767 < 158) {
                    var v18 uint8 = v7
                    _ = v18
                }
                v3[2147483647] = (fi17 % ((fi17) | 1))
                _ = fi17
            }
            var v19 int8 = 64
            _ = v16
            _ = v19
        } else {
            if false {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v1 -= v1
                    fmt.Println(v1, v7)
                    _ = fi20
                }
                _ = len(v4)
                v21 := GenS2{F0: 0}
                _ = v21
            }
            switch "hello" {
            case "foo bar":
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 int = 32767
                    _ = fi22
                    _ = v23
                }
            case "\n\t":
                var v24 uint = 167
                _ = v24
            }
            switch -1 {
            case 42:
                v7 = 1
            }
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(fi26, v7)
                v27 := GenS3{F0: true, F1: 2, F2: (-(fi25))}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        if true {
            v1 -= 247
            var v28 uint8 = (^(v7))
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 int = (fi29 ^ fi29)
                _ = len(v3)
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            delete(v3, 4294967295)
        }
    } else {
        v31 := GenS4{F0: -1, F1: -1, F2: 0}
        switch ("a" + (("hello" + "") + ("hello" + "a"))) {
        case "hello":
            switch 4294967295 {
            case 128:
                var v32 float64 = 1.5
                _ = v32
            case 2147483647:
                v33 := GenS5{F0: 42, F1: v7}
                _ = v33
            }
        }
        var v34 int32 = 1
        _ = v31
        _ = v34
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if (!((0.1 / -2.25) <= -2.25)) {
        var v35 int8 = 42
        var v36 uint16 = 0
        _ = v35
        _ = v36
    }
    v37 := GenS3{F0: false, F1: 42, F2: -9223372036854775808}
    if true {
        var v38 uint16 = 406
        switch 2147483647 {
        case 4294967295:
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        case 32767:
            _ = len(v2)
        case 2147483647:
            fmt.Println(v5, v1)
        default:
            if false {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
        }
        var v39 string = ("" + "a")
        _ = v38
        _ = v39
    }
    if false {
        v40 := GenS6{F0: 255, F1: (!false)}
        fmt.Println(v1)
        _ = v40
    }
    var v41 float32 = -2.25
    v1 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
