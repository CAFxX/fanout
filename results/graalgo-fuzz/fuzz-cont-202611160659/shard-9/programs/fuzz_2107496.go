package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int32; F2 uint16 }
func genFn0(p0 uint64) int8 {
    fmt.Println(p0)
    var v5 float64 = 0.1
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 string, p1 uint8, p2 int) uint {
    p2 -= p2
    var v6 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 int = 128
    v2 := []string{(("foo bar" + "hello") + ("foo bar" + ("a" + "\n\t")))}
    var v3 int64 = -23
    v4 := make(map[int]uint32)
    var v7 int16 = -171
    for fi8 := 0; fi8 < 11; fi8++ {
        if (0.0 != float64(0.1)) {
            if false {
                v9 := v4[(-(fi8))]
                if true {
                    var v10 int = v1
                    _ = v10
                } else {
                    fmt.Println(v1, v7, v3)
                    fmt.Println(v1)
                    var v11 int16 = v7
                    _ = v11
                }
                _ = v9
            } else {
                var v12 string = (("\n\t" + "\n\t") + "a")
                var v13 int = (v1 - 128)
                fmt.Println(v3)
                _ = v12
                _ = v13
            }
            v3 -= (v3 - 42)
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            switch ("hello" + "\n\t") {
            case "":
                var v15 int = (fi14 >> ((1) & 63))
                _ = v15
            case "\n\t":
                var v16 uint64 = 0
                _ = v16
            }
            _ = fi14
        }
        var v17 string = "a"
        _ = fi8
        _ = v17
    }
    var v18 int16 = v7
    v3 -= v3
    for fi19 := 0; fi19 < 4; fi19++ {
        v2 = append(v2, "\n\t")
        _ = fi19
    }
    v3 = v3
    if false {
        var v20 uint16 = 1
        v3 = (^(v3))
        _ = v20
    } else {
        v21 := GenS1{F0: true, F1: 1, F2: ("hello" + "a")}
        var v22 uint64 = 42
        fmt.Println(v22, v22)
        _ = v21
        _ = v22
    }
    if (42 == 621) {
        v1 = v1
    } else {
        v23 := GenS1{F0: false, F1: 1, F2: ("a" + "a")}
        switch ("a" + "\n\t") {
        case "\n\t":
            fmt.Println(v3, v3, v1)
        case "foo bar":
            var v24 int16 = 42
            _ = v24
        }
        fmt.Println(v1, v3)
        _ = v23
    }
    fmt.Println(v18, v1)
    v25 := GenS2{F0: 945}
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS1{F0: true, F1: 2147483648, F2: "hello"}
            switch 98 {
            case -516:
                v29 := GenS3{F0: "hello"}
                _ = v29
            case 127:
                v30 := GenS4{F0: 1, F1: -717, F2: 611}
                var v31 uint8 = 128
                v32 := v4[fi27]
                _ = v30
                _ = v31
                _ = v32
            case -952:
                _ = len(v4)
                fmt.Println(fi27, v7, v7)
                var v33 uint8 = 230
                switch (("foo bar" + "foo bar") + "foo bar") {
                case "hello":
                    _ = len(v4)
                case "foo bar":
                    var v34 string = ("foo bar" + "\n\t")
                    _ = v34
                case "\n\t":
                    fmt.Println(fi26)
                default:
                    v33 = (v33 >> ((901) & 63))
                }
                _ = v33
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
