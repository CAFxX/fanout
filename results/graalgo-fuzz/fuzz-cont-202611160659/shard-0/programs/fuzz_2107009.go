package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint8 = 65
    v2 := []string{}
    v3 := []string{("" + "hello"), "foo bar", ""}
    var v4 int32 = 0
    v5 := GenS1{F0: 32767, F1: 2}
    v1 += v1
    v6 := GenS1{F0: -9223372036854775808, F1: 42}
    if true {
        v4 -= v4
        fmt.Println(v1, v4)
        if (42 != 255) {
            for fi7 := 0; fi7 < 9; fi7++ {
                if (255 != 533) {
                    v4 = v4
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 int16 = 32767
                        v10 := GenS1{F0: -868, F1: 255}
                        v11 := GenS2{F0: 237, F1: 460, F2: (!true)}
                        var v12 string = ("hello" + ("" + ""))
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    }
                    var v13 float32 = 0.0
                    _ = v13
                }
                _ = fi7
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v3) > 0 {
                v3[0] = (("foo bar" + "hello") + ("" + ""))
            }
        }
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v1 = (^(v1))
        _ = fi14
    }
    var v15 int16 = 32767
    v15 += v15
    v16 := GenS3{F0: false}
    v17 := GenS4{F0: v1}
    switch -239 {
    case 484:
        var v18 int32 = int32(v15)
        _ = v18
    case 32768:
        for fi19 := 0; fi19 < 13; fi19++ {
            v4 = (v4 | v4)
            for fi20 := 0; fi20 < 20; fi20++ {
                v15 = v15
                _ = fi20
            }
            v4 += 1
            _ = fi19
        }
    case 2147483647:
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 int8 = 100
            _ = fi21
            _ = v22
        }
    }
    var v23 uint32 = 354
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS5{F0: 68, F1: v23}
        v4 += v4
        if len(v3) > 0 {
            v3[0] = ("hello" + ("" + ("hello" + "foo bar")))
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v15, v23, v23)
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: fi26, F1: uint(v1)}
        v1 += v1
        _ = fi26
        _ = v27
    }
    _ = len(v2)
    for fi28 := 0; fi28 < 10; fi28++ {
        if true {
            for fi29 := 0; fi29 < 5; fi29++ {
                v23 += v23
                switch ("hello" + "") {
                case "a":
                    var v30 uint = 838
                    _ = v30
                default:
                    var v31 uint32 = v23
                    _ = v31
                }
                _ = fi29
            }
            fmt.Println(fi28, v1, fi28)
            fmt.Println(v15)
        }
        if len(v3) > 0 {
            v3[0] = ("a" + "a")
        }
        _ = len(v2)
        _ = fi28
    }
    if (0 == 316) {
        var v32 uint32 = v23
        switch ("foo bar" + ("a" + "hello")) {
        case "a":
            for fi33 := 0; fi33 < 17; fi33++ {
                if len(v2) > 0 {
                    v2[0] = ("" + "")
                }
                v3 = append(v3, "a")
                fmt.Println(v4, v4, fi33)
                _ = fi33
            }
        }
        var v34 uint16 = 255
        _ = v32
        _ = v34
    } else {
        v35 := GenS6{F0: 436}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
