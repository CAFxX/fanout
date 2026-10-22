package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 string }
func genFn0(p0 int64, p1 uint32, p2 int32) int {
    p1 -= p1
    v6 := GenS1{F0: 662, F1: 18446744073709551615}
    var v7 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 int, p1 bool) string {
    fmt.Println(p1)
    var v8 int16 = -32768
    var v9 uint = 18446744073709551615
    v10 := GenS1{F0: (-(v8)), F1: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 255
    var v3 uint = 4294967295
    var v4 int = -714
    var v5 int64 = 0
    if (true || true) {
        var v11 int8 = 42
        v5 = 9223372036854775807
        fmt.Println(v5)
        _ = v11
    } else {
        var v12 bool = (65535 > 0)
        fmt.Println(v2)
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v4, v3, v4)
        _ = fi13
    }
    if (((0.0 != 3.14159) || false) && (false && true)) {
        for fi14 := 0; fi14 < 20; fi14++ {
            if (true || true) {
                fmt.Println(v4)
            } else {
                if false {
                    for fi15 := 0; fi15 < 15; fi15++ {
                        delete(v1, "hello")
                        v5 += (v5 & 0)
                        _ = fi15
                    }
                    fmt.Println(v5)
                } else {
                    fmt.Println(v3, v4)
                    v16 := GenS1{F0: 1000, F1: v3}
                    _ = v16
                }
                var v17 uint64 = 2
                fmt.Println(v2)
                _ = v17
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                v5 += (-(v5))
                if ((190 >= -1) && ("foo bar" > "")) {
                    v5 = (v5 / (((-(v5))) | 1))
                }
                _ = fi18
            }
            _ = fi14
        }
        if true {
            for fi19 := 0; fi19 < 18; fi19++ {
                v2 = 2
                _ = fi19
            }
            v20 := GenS1{F0: 1000, F1: v3}
            v21 := []int{(^(v4)), v4, v4, (v4 % (((v4 >> ((v3) & 63))) | 1))}
            _ = v20
            _ = v21
        }
    }
    fmt.Println(v3)
    var v22 float64 = 1e10
    v5 = (v5 * -9223372036854775808)
    v5 += -829
    if true {
        var v23 uint32 = 2
        _ = v23
    } else {
        var v24 int64 = (-(v5))
        var v25 int8 = -99
        if (69 > 940) {
            v26 := GenS1{F0: -32768, F1: v3}
            v27 := GenS2{F0: 2147483647}
            v28 := GenS2{F0: 4294967295}
            _ = v26
            _ = v27
            _ = v28
        }
        _ = v24
        _ = v25
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 uint16 = 32767
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: 1}
    v4 = v4
    var v32 int64 = v5
    v33 := GenS4{F0: 100, F1: 357}
    if ("hello" > ((("hello" + ("\n\t" + "")) + "foo bar") + "")) {
        var v34 int8 = -1
        v35 := GenS5{F0: 42, F1: "hello"}
        v36 := GenS1{F0: -595, F1: v3}
        _ = v34
        _ = v35
        _ = v36
    }
    if true {
        if true {
            v5 -= -788
            fmt.Println(v4)
        } else {
            v1[("foo bar" + "a")] = (v4 & v4)
        }
        var v37 int = (v4 ^ v4)
        _ = v37
    }
    v38 := []int{genFn0((v5 | v32), 850, 1), int(v3), (v4 & 28), (v4 - (-(v4)))}
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 string = "\n\t"
        switch 2147483647 {
        case 128:
            switch v4 {
            case -1:
                v41 := GenS4{F0: -1, F1: 1}
                _ = v41
            case -678:
                for fi42 := 0; fi42 < 4; fi42++ {
                    switch v40 {
                    case "":
                        if len(v38) >= 2 {
                            v38 = v38[1:len(v38)]
                        }
                    }
                    _ = fi42
                }
            case 127:
                var v43 uint64 = 364
                _ = v43
            }
        case -1:
            v44 := GenS1{F0: -32768, F1: 18446744073709551615}
            _ = v44
        }
        v45 := GenS6{F0: (v40 + ("foo bar" + (((((("foo bar" + (v40 + "hello")) + v40) + ("foo bar" + v40)) + "\n\t") + (v40 + v40)) + v40)))}
        _ = fi39
        _ = v40
        _ = v45
    }
    v46 := GenS1{F0: -474, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v46
    _ = v5
}
