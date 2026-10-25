package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 uint8; F2 int }
type GenS4 struct { F0 int; F1 uint64; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int; F1 int64 }
func genFn0(p0 uint, p1 uint8, p2 int64) string {
    switch -1 {
    case -685:
        v7 := []int{853, -9223372036854775808, 128, 2147483647}
        _ = v7
    case 127:
        p2 = (p2 >> ((p0) & 63))
    case 32767:
        p2 += -9223372036854775808
    }
    p2 = (p2 / ((p2) | 1))
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 65535
        v10 := GenS1{F0: 128, F1: false}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return ("\n\t" + ("foo bar" + "hello"))
}
func genFn1(p0 uint16, p1 uint, p2 int16) uint64 {
    p2 = (p2 | p2)
    if false {
        var v11 uint64 = 4294967295
        _ = v11
    } else {
        switch 4294967295 {
        case 32768:
            v12 := []int{int(p0)}
            _ = v12
        case 1:
            var v13 bool = false
            _ = v13
        }
        fmt.Println(p2, p0, p1)
    }
    v14 := []int{127, -870, -1}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 float32 = (3.14159 - (0.0 * 1.5))
        switch "a" {
        case "hello":
            if ((!true) || true) {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            } else {
                _ = len(v14)
                switch (^(fi15)) {
                case -921:
                    if false {
                        if true {
                            v17 := GenS2{F0: (1 < (fi15 | 128)), F1: p2, F2: "\n\t"}
                            _ = v17
                        }
                        var v18 float32 = v16
                        _ = v18
                    }
                }
            }
        case "\n\t":
            var v19 uint = 780
            _ = v19
        case "":
            var v20 uint64 = 4294967295
            _ = v20
        }
        v21 := GenS1{F0: p0, F1: (-9223372036854775808 == 42)}
        _ = fi15
        _ = v16
        _ = v21
    }
    switch "a" {
    case "hello":
        v22 := GenS3{F0: 100, F1: 127, F2: 2147483647}
        _ = v22
    default:
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 633
}

func main() {
    var v1 int16 = 471
    var v2 int8 = -128
    var v3 int16 = (v1 % ((-32768) | 1))
    v4 := []string{("" + (("\n\t" + (("\n\t" + "") + ("a" + "hello"))) + ("" + "hello"))), "\n\t", ("foo bar" + "a")}
    v5 := make(map[uint32]int)
    var v6 uint8 = 127
    v2 = int8(v3)
    if false {
        v5[2147483647] = -714
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS4{F0: (fi23 * (fi23 + fi23)), F1: 42, F2: (-(v2))}
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v6 = 109
        _ = len(v5)
        _ = fi25
    }
    if true {
        v6 = (v6 | (v6 + 1))
        if ((("\n\t" + (("hello" + "hello") + "")) + "") >= "a") {
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(v3, v1, v6)
                _ = fi26
            }
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS5{F0: 0}
                _ = fi27
                _ = v28
            }
        }
        _ = len(v5)
    } else {
        var v29 int = 127
        _ = v29
    }
    v3 = v1
    switch "" {
    case "hello":
        v3 = (v3 << ((1) & 63))
    case "\n\t":
        fmt.Println(v1, v6)
    }
    v30 := GenS6{F0: v3, F1: (v2 + (v2 - (-(v2))))}
    if (0 > 699) {
        fmt.Println(v6, v6)
        if (-2.25 < -2.25) {
            fmt.Println(v3, v6)
            switch ("a" + "a") {
            case "a":
                v31 := GenS7{F0: 9223372036854775807, F1: -143}
                _ = v31
            }
        } else {
            v1 = (v3 << ((1) & 63))
            var v32 float32 = float32(3.14159)
            _ = v32
        }
    }
    switch "hello" {
    case "hello":
        for fi33 := 0; fi33 < 2; fi33++ {
            delete(v5, 2147483648)
            for fi34 := 0; fi34 < 11; fi34++ {
                fmt.Println(v1, fi33)
                fmt.Println(v2, v1, v1)
                v35 := GenS5{F0: 196}
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
    case "foo bar":
        if true {
            fmt.Println(v1, v6)
            v3 = v1
            var v36 uint8 = (^(v6))
            _ = v36
        }
    case "\n\t":
        fmt.Println(v1, v3, v1)
    default:
        v37 := GenS2{F0: (!((^(v6)) != (^(v6)))), F1: 1000, F2: ("hello" + "hello")}
        _ = v37
    }
    v1 = v3
    var v38 uint64 = 2
    fmt.Println(v2, v1)
    for fi39 := 0; fi39 < 11; fi39++ {
        if false {
            var v40 int16 = v1
            for fi41 := 0; fi41 < 14; fi41++ {
                var v42 string = ((("foo bar" + "foo bar") + ("a" + ("\n\t" + ("hello" + "a")))) + (("\n\t" + ("" + ("hello" + "a"))) + ""))
                if false {
                    v6 += v6
                    var v43 float32 = 0.1
                    _ = v43
                }
                v6 = v6
                _ = fi41
                _ = v42
            }
            v44 := v5[292]
            _ = v40
            _ = v44
        } else {
            v45 := GenS2{F0: (2147483647 >= 2147483647), F1: v3, F2: "a"}
            for fi46 := 0; fi46 < 3; fi46++ {
                v3 -= (v1 << ((1) & 63))
                _ = fi46
            }
            _ = v45
        }
        _ = fi39
    }
    if false {
        switch "a" {
        case "\n\t":
            var v47 uint = 2
            _ = v47
        case "hello":
            v1 -= (-(v3))
        }
        v1 = (^(v1))
        v48 := GenS6{F0: v3, F1: v2}
        _ = v48
    } else {
        var v49 uint16 = 32767
        _ = len(v4)
        for fi50 := 0; fi50 < 7; fi50++ {
            v6 -= (v6 | 42)
            var v51 int32 = 234
            _ = fi50
            _ = v51
        }
        _ = v49
    }
    v3 -= (v3 / (((-(v3))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
