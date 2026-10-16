package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
func genFn0() uint16 {
    var v4 uint64 = 4294967295
    v4 -= v4
    var v5 uint8 = 0
    v6 := GenS1{F0: 65535}
    v4 = (v4 & (v4 + 2))
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int, p1 int, p2 uint32) int {
    v7 := []int{p1, p0, p0, 32768}
    if true {
        var v8 uint64 = 457
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v8
    } else {
        var v9 uint32 = (^(p2))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 664
}

func main() {
    var v1 float64 = float64(3.14159)
    var v2 uint = 18446744073709551615
    var v3 uint32 = 323
    v10 := []int{}
    fmt.Println(v1)
    var v11 int64 = 42
    if true {
        v12 := GenS1{F0: v3}
        if true {
            v1 = ((v1 + 0.0) * ((((((0.1 * (float64((3.14159 + -2.25)) - (v1 + v1))) * float64(-2.25)) - v1) - float64(-2.25)) * ((3.14159 - (1.5 - v1)) + v1)) * v1))
            if (true || (!(false && false))) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        var v15 int64 = (v11 + v11)
                        _ = len(v10)
                        v2 = (v2 << (((v2 - (^(v2)))) & 63))
                        _ = fi14
                        _ = v15
                    }
                    var v16 int16 = 0
                    _ = fi13
                    _ = v16
                }
                fmt.Println(v1, v2)
            } else {
                for fi17 := 0; fi17 < 15; fi17++ {
                    _ = len(v10)
                    _ = fi17
                }
                if len(v10) > 0 {
                    v10[0] = -788
                }
                if false {
                    _ = len(v10)
                    v2 -= (v2 ^ v2)
                }
            }
        } else {
            var v18 bool = (65535 <= v3)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v19 := GenS2{F0: 589}
            _ = v18
            _ = v19
        }
        var v20 int8 = 127
        _ = v12
        _ = v20
    } else {
        v1 = (v1 / float64(1e10))
        for fi21 := 0; fi21 < 6; fi21++ {
            switch 128 {
            case -2147483648:
                for fi22 := 0; fi22 < 4; fi22++ {
                    var v23 int16 = -1
                    _ = fi22
                    _ = v23
                }
            case 664:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            default:
                switch 127 {
                case 2147483647:
                    v24 := GenS2{F0: 42}
                    _ = v24
                default:
                    switch "\n\t" {
                    case "a":
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    case "hello":
                        switch ("a" + "\n\t") {
                        case "hello":
                            for fi25 := 0; fi25 < 14; fi25++ {
                                fmt.Println(fi25)
                                var v26 float64 = ((float64((3.14159 / 1.5)) + float64(((-2.25 / 3.14159) * 0.1))) / (v1 + -2.25))
                                var v27 uint8 = 201
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        case "a":
                            v2 = 2
                        case "":
                            if (true || true) {
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    var v29 bool = false
                                    var v30 int32 = -568
                                    _ = len(v10)
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                var v31 int32 = -2147483648
                                _ = v31
                            }
                        }
                    case "foo bar":
                        var v32 float64 = (v1 + -2.25)
                        _ = v32
                    default:
                        v33 := GenS3{F0: 1, F1: (("foo bar" + "a") + "\n\t"), F2: 105}
                        _ = v33
                    }
                }
            }
            _ = fi21
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 float32 = 0.1
            _ = fi34
            _ = v35
        }
    }
    switch 1 {
    case 0:
        fmt.Println(v2, v11, v2)
    case 574:
        var v36 string = "hello"
        _ = v36
    }
    switch (("foo bar" + ("" + "")) + "foo bar") {
    case "":
        var v37 int8 = -28
        _ = v37
    case "a":
        fmt.Println(v2, v11)
    case "hello":
        var v38 float32 = -2.25
        _ = v38
    default:
        for fi39 := 0; fi39 < 15; fi39++ {
            if len(v10) > 0 {
                v10[0] = fi39
            }
            var v40 float64 = -2.25
            var v41 bool = false
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v11 -= v11
    v1 = float64((3.14159 + 1.5))
    v3 = 2147483647
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch ("" + (("\n\t" + ("hello" + "a")) + "foo bar")) {
    case "foo bar":
        _ = len(v10)
    case "\n\t":
        var v42 int8 = -24
        _ = v42
    case "hello":
        v43 := GenS2{F0: 1}
        _ = v43
    default:
        v2 += (^(v2))
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
