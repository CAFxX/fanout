package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint8) uint8 {
    p2 -= uint8(p0)
    var v7 string = "foo bar"
    var v8 int32 = -417
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{int(v8), 0}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := []int{}
        if len(v12) > 0 {
            v12[0] = (fi11 / (((fi11 | fi11)) | 1))
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 - p2)
}
func genFn1(p0 int16) uint16 {
    for fi13 := 0; fi13 < 3; fi13++ {
        switch "hello" {
        case "foo bar":
            var v14 uint8 = genFn0(671, 255, 2)
            _ = v14
        case "a":
            switch ("foo bar" + (("hello" + "") + "foo bar")) {
            case "hello":
                fmt.Println(fi13, p0)
            case "foo bar":
                switch "a" {
                case "\n\t":
                    var v15 uint64 = 235
                    _ = v15
                }
            case "\n\t":
                p0 = -1
            }
        case "":
            p0 += 32767
        }
        _ = fi13
    }
    var v16 int8 = 127
    if true {
        var v17 uint64 = 18446744073709551615
        _ = v17
    }
    switch ((("a" + "foo bar") + (((("a" + ("foo bar" + "foo bar")) + "foo bar") + "foo bar") + "a")) + ("a" + ("a" + ""))) {
    case "foo bar":
        fmt.Println(v16, p0)
    case "hello":
        v18 := GenS1{F0: "hello", F1: 390, F2: 127}
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 2; fi19++ {
            p0 -= p0
            _ = fi19
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v16
    return 42
}

func main() {
    var v1 int = 128
    var v2 uint16 = 0
    var v3 uint64 = 687
    var v4 int32 = 2147483647
    var v5 float64 = -2.25
    v6 := []int{}
    v2 += v2
    var v20 int16 = -1
    _ = len(v6)
    var v21 int16 = (v20 - int16(v1))
    fmt.Println(v20, v21)
    for fi22 := 0; fi22 < 7; fi22++ {
        v20 = (v21 % ((v20) | 1))
        fmt.Println(v3, v5, v21)
        if false {
            v21 -= v21
            if (fi22 == 32768) {
                var v23 uint = 1
                if ("hello" > (((("\n\t" + ((("a" + ("\n\t" + "")) + ("hello" + "foo bar")) + "hello")) + ("hello" + ("" + "foo bar"))) + (("\n\t" + ("foo bar" + ("hello" + "hello"))) + "")) + "a")) {
                    fmt.Println(v21, v23, v5)
                }
                var v24 string = "foo bar"
                _ = v23
                _ = v24
            } else {
                var v25 uint = 1
                for fi26 := 0; fi26 < 14; fi26++ {
                    if (true && (1 < -9223372036854775808)) {
                        var v27 float64 = v5
                        v28 := GenS1{F0: "foo bar", F1: (v2 - v2), F2: 114}
                        var v29 uint = 255
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        if len(v6) > 0 {
                            v6[0] = (^(fi22))
                        }
                        v6 = append(v6, (fi26 & -472))
                    }
                    _ = fi26
                }
                switch "" {
                case "\n\t":
                    for fi30 := 0; fi30 < 6; fi30++ {
                        if ((v21 >= v20) || (!((v3 ^ (v3 % (((^(v3))) | 1))) == v3))) {
                            var v31 uint32 = 2147483647
                            _ = v31
                        } else {
                            var v32 uint = uint(v2)
                            _ = v32
                        }
                        var v33 uint16 = genFn1((v20 << (((v25 | 4294967295)) & 63)))
                        _ = fi30
                        _ = v33
                    }
                case "hello":
                    var v34 string = "hello"
                    _ = v34
                case "":
                    var v35 uint64 = uint64(v2)
                    _ = v35
                default:
                    v3 = (v3 << (((v25 & (v25 + (v25 | 496)))) & 63))
                }
                _ = v25
            }
        } else {
            v36 := GenS1{F0: ("a" + "a"), F1: (^(v2)), F2: 2}
            fmt.Println(v1)
            v5 -= v5
            _ = v36
        }
        _ = fi22
    }
    if len(v6) > 0 {
        v6[0] = int(v4)
    }
    _ = len(v6)
    fmt.Println(v4, v20, v5)
    if (2147483647 < (v4 - -2147483648)) {
        fmt.Println(v20, v1)
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        v4 += v4
        if (true && false) {
            v6 = append(v6, (v1 << ((65535) & 63)))
            var v38 float64 = float64(1.5)
            var v39 int = 42
            _ = v38
            _ = v39
        }
        v4 += 0
        _ = fi37
    }
    v40 := GenS2{F0: 127, F1: 0, F2: 0}
    var v41 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
