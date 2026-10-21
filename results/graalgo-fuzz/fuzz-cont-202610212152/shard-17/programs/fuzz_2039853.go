package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -2147483648
    v3 := make(map[string]int)
    var v4 uint64 = 18446744073709551615
    var v5 int16 = 884
    switch "a" {
    case "a":
        v4 = 4294967295
    case "":
        v6 := GenS1{F0: "a", F1: 1, F2: (-90 != -867)}
        _ = v6
    default:
        var v7 uint16 = 42
        _ = v7
    }
    var v8 int16 = (^(v5))
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 = fi9
            _ = fi10
        }
        var v11 uint32 = 2
        v12 := GenS2{F0: v4, F1: 42, F2: 42}
        _ = fi9
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 0, F1: 2147483647, F2: 127}
    for fi14 := 0; fi14 < 9; fi14++ {
        v5 += v8
        if false {
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 = fi15
                if (true || false) {
                    var v16 float32 = 0.0
                    var v17 bool = ((false && (127 <= -128)) || false)
                    _ = v16
                    _ = v17
                }
                v8 += 643
                v2 = (fi15 - v2)
                _ = fi15
            }
            var v18 float64 = float64(((((0.0 + 3.14159) + 1.5) + 1e10) + 0.1))
            if true {
                if true {
                    switch v2 {
                    case 4294967295:
                        var v19 float32 = -2.25
                        _ = v19
                    }
                    v20 := []int{int(v5), -1, -604}
                    _ = v20
                }
                v4 = (v4 & (^(v4)))
                if true {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        fmt.Println(v4, v18, fi14)
                        v22 := GenS3{F0: -12}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    if (true && true) {
                        delete(v1, 2147483648)
                    } else {
                        v23 := GenS4{F0: 255}
                        v24 := GenS4{F0: 0}
                        _ = v23
                        _ = v24
                    }
                    v5 -= v5
                    v25 := []int{v2, 127}
                    _ = v25
                }
            }
            _ = v18
        }
        if ((true || (v4 >= (v4 << ((1) & 63)))) && false) {
            var v26 string = ("" + (("\n\t" + "hello") + "\n\t"))
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 float64 = (3.14159 + 1.5)
                v8 = v5
                v2 = (fi14 % ((1) | 1))
                _ = fi27
                _ = v28
            }
            switch fi14 {
            case 9223372036854775807:
                if false {
                    var v29 uint8 = 133
                    var v30 uint16 = 65535
                    _ = v29
                    _ = v30
                }
            default:
                switch (v26 + (v26 + "\n\t")) {
                case "a":
                    var v31 int32 = -1
                    _ = v31
                case "":
                    var v32 int64 = 42
                    _ = v32
                case "foo bar":
                    v4 += (v4 >> ((4294967295) & 63))
                default:
                    if (!(v26 != v26)) {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v5 -= -239
                            fmt.Println(v4, v4)
                            fmt.Println(v4)
                            _ = fi33
                        }
                    }
                }
            }
            _ = v26
        }
        _ = fi14
    }
    v34 := GenS5{F0: ((("\n\t" + "") + ((("foo bar" + "a") + "") + ("hello" + ((("foo bar" + "foo bar") + "") + (("foo bar" + "foo bar") + ("a" + "hello")))))) + "hello"), F1: ""}
    for fi35 := 0; fi35 < 10; fi35++ {
        delete(v3, ("foo bar" + (((("" + "a") + "") + "foo bar") + "")))
        for fi36 := 0; fi36 < 5; fi36++ {
            var v37 int8 = int8(v5)
            v2 = v2
            _ = fi36
            _ = v37
        }
        var v38 int64 = 42
        _ = fi35
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v8 = (v5 | (v5 ^ (-(v5))))
        _ = fi39
    }
    v40 := []int{}
    v2 = v2
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2, v8, v5)
    case "":
        v8 += v8
    case "a":
        _ = len(v40)
    default:
        if (65535 > 654) {
            var v41 string = "hello"
            fmt.Println(v4)
            v41 = "foo bar"
            _ = v41
        }
    }
    switch "a" {
    case "":
        for fi42 := 0; fi42 < 19; fi42++ {
            var v43 uint32 = 0
            _ = fi42
            _ = v43
        }
    case "foo bar":
        v44 := GenS6{F0: 1, F1: 97}
        _ = v44
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        var v46 bool = (!true)
        for fi47 := 0; fi47 < 7; fi47++ {
            fmt.Println(fi47)
            fmt.Println(v2)
            v8 = (v5 ^ 1)
            _ = fi47
        }
        var v48 string = ("" + "a")
        _ = fi45
        _ = v46
        _ = v48
    }
    v4 += (^(v4))
    var v49 bool = true
    var v50 float64 = ((-2.25 + 1.5) + 3.14159)
    if v49 {
        fmt.Println(v4, v49, v8)
        if (-119 >= -128) {
            var v51 int8 = -128
            var v52 uint64 = v4
            v50 = 3.14159
            _ = v51
            _ = v52
        } else {
            v53 := GenS1{F0: (("" + "\n\t") + ""), F1: 1, F2: (int8(v8) > -128)}
            for fi54 := 0; fi54 < 6; fi54++ {
                if true {
                    v1[65535] = 1
                } else {
                    v55 := v1[2147483647]
                    var v56 uint32 = v55
                    _ = v55
                    _ = v56
                }
                var v57 bool = (!v49)
                v58 := GenS7{F0: v4}
                _ = fi54
                _ = v57
                _ = v58
            }
            _ = v53
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v49
    _ = v5
    _ = v50
    _ = v8
}
