package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0(p0 uint64, p1 uint, p2 uint) int64 {
    fmt.Println(p1, p1, p1)
    fmt.Println(p2, p2, p2)
    for fi7 := 0; fi7 < 8; fi7++ {
        switch "a" {
        case "foo bar":
            var v8 int64 = 42
            _ = v8
        case "hello":
            p1 -= (^(p1))
        case "a":
            if true {
                var v9 int8 = -114
                _ = v9
            }
        default:
            switch "\n\t" {
            case "foo bar":
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 int64 = 9223372036854775807
                    var v12 int = fi10
                    v13 := []int{4294967295, (v12 * (^(fi7))), (v12 + (fi10 - int(v11))), fi7}
                    var v14 string = ("\n\t" + "hello")
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            case "hello":
                v15 := GenS1{F0: 100, F1: 32767}
                _ = v15
            case "a":
                v16 := []int{int(p0), (-(fi7)), fi7, fi7}
                _ = v16
            default:
                var v17 bool = true
                _ = v17
            }
        }
        p1 += (p1 * uint(p0))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint16, p1 uint32, p2 int16) int {
    p1 += 4294967295
    if false {
        p2 += -201
    } else {
        p2 = (p2 ^ 1000)
        if true {
            v18 := []int{-509, 127}
            _ = v18
        }
        fmt.Println(p2, p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]uint32)
    var v3 int = -1
    var v4 int32 = -599
    var v5 uint = 255
    v6 := []string{""}
    for fi19 := 0; fi19 < 16; fi19++ {
        if true {
            fmt.Println(fi19, v5)
            var v20 float32 = 0.0
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(fi21)
        v22 := GenS1{F0: 0, F1: 1}
        _ = fi21
        _ = v22
    }
    switch ("foo bar" + "") {
    case "\n\t":
        var v23 uint8 = 255
        _ = v23
    case "hello":
        var v24 float32 = (1.5 - 3.14159)
        _ = v24
    case "a":
        if true {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 float32 = 1e10
                switch ("foo bar" + ("\n\t" + "")) {
                case "a":
                    switch ("\n\t" + "") {
                    case "":
                        var v27 int8 = -128
                        _ = v27
                    case "a":
                        var v28 int = v3
                        _ = v28
                    case "foo bar":
                        fmt.Println(v1, v1, v3)
                    }
                case "":
                    if false {
                        v29 := GenS1{F0: 84, F1: 1000}
                        v30 := GenS2{F0: 329}
                        _ = v29
                        _ = v30
                    }
                default:
                    if (genFn1(0, 42, 0) != genFn1(65535, 2147483647, 1)) {
                        fmt.Println(v3)
                        v3 -= (fi25 >> (((^(v5))) & 63))
                    }
                }
                _ = fi25
                _ = v26
            }
            var v31 int16 = -1
            _ = v31
        } else {
            v1 -= -697
        }
    default:
        for fi32 := 0; fi32 < 18; fi32++ {
            if (!false) {
                var v33 bool = true
                if ((v1 - (^(v1))) >= v1) {
                    v34 := GenS3{F0: 2}
                    for fi35 := 0; fi35 < 8; fi35++ {
                        fmt.Println(v4, fi35)
                        _ = fi35
                    }
                    _ = len(v6)
                    _ = v34
                }
                _ = v33
            }
            if false {
                fmt.Println(v4, fi32)
            }
            v3 = (fi32 >> (((v5 * 65535)) & 63))
            _ = fi32
        }
    }
    switch (("foo bar" + "") + ("foo bar" + "hello")) {
    case "foo bar":
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 int = (fi36 + v3)
            fmt.Println(v37, v3)
            _ = fi36
            _ = v37
        }
    default:
        var v38 int16 = 0
        _ = v38
    }
    fmt.Println(v3, v5)
    switch "hello" {
    case "":
        var v39 uint64 = 42
        _ = v39
    }
    var v40 int16 = 175
    _ = len(v6)
    fmt.Println(v40)
    if (80 == 250) {
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS4{F0: ("a" + "hello"), F1: 493, F2: (-(v1))}
            if true {
                v5 = (v5 & v5)
                v43 := GenS5{F0: 1, F1: ("a" <= (("\n\t" + ("hello" + ("a" + (((("a" + ("hello" + "\n\t")) + "foo bar") + ((("hello" + "hello") + ("a" + "foo bar")) + "foo bar")) + "\n\t")))) + "a"))}
                _ = v43
            } else {
                var v44 float64 = float64((0.1 * 0.0))
                v45 := GenS5{F0: uint64(v5), F1: (!false)}
                _ = v44
                _ = v45
            }
            if (!(0.0 >= 1e10)) {
                for fi46 := 0; fi46 < 11; fi46++ {
                    var v47 float32 = ((float32(1e10) - (3.14159 / 1e10)) + 0.0)
                    v40 = (v40 / ((v40) | 1))
                    v40 = v40
                    _ = fi46
                    _ = v47
                }
                var v48 uint64 = 4294967295
                _ = v48
            }
            _ = fi41
            _ = v42
        }
        var v49 int32 = v4
        v4 += 2147483647
        _ = v49
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            v5 -= 996
            v4 = -1
            _ = fi50
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v51 := GenS1{F0: -38, F1: (v40 + v40)}
    for fi52 := 0; fi52 < 4; fi52++ {
        v53 := v2[4294967295]
        switch (v3 ^ (fi52 - fi52)) {
        case -1:
            switch (^(v3)) {
            case 9223372036854775807:
                var v54 float32 = (0.0 * 1.5)
                _ = v54
            default:
                fmt.Println(v4, v4, v53)
            }
        default:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi52
        _ = v53
    }
    v1 = (v1 % ((v1) | 1))
    var v55 int64 = (v1 + v1)
    v4 = 42
    var v56 int32 = 42
    var v57 uint8 = 0
    var v58 int16 = (v40 | v40)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v51
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
