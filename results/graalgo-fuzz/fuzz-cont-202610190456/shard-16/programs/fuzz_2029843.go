package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 int32) uint32 {
    fmt.Println(p0)
    v7 := GenS1{F0: 394}
    var v8 int = -538
    _ = p0
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 int, p1 int16, p2 int64) bool {
    p2 -= (p2 & -1)
    p1 = (p1 | p1)
    var v9 int32 = 767
    if (!false) {
        var v10 int = p0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (!true)
}

func main() {
    var v1 uint = 65535
    var v2 int = 316
    var v3 uint16 = 233
    v4 := []string{"foo bar", "hello"}
    v5 := make(map[uint32]int)
    var v6 uint16 = 1
    v3 = (v6 >> (((v1 + v1)) & 63))
    var v11 uint16 = v6
    var v12 int64 = 0
    if ((!false) || (0 <= 715)) {
        _ = len(v4)
        v1 = (v1 & v1)
    }
    v5[2] = v2
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: 856}
            var v16 int8 = 0
            v4 = append(v4, "\n\t")
            if false {
                if len(v4) > 0 {
                    v4[0] = ("foo bar" + "foo bar")
                }
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            }
            switch ("\n\t" + ((("" + "hello") + "a") + "")) {
            case "":
                v17 := GenS1{F0: 1}
                _ = v17
            case "hello":
                var v18 uint32 = uint32(v1)
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS1{F0: 700}
            fmt.Println(fi19, v3, v11)
            _ = fi19
            _ = v20
        }
        _ = fi13
    }
    if true {
        var v21 uint16 = (v11 << (((^(v1))) & 63))
        for fi22 := 0; fi22 < 7; fi22++ {
            if false {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v3 = (^(v21))
                    v1 -= v1
                    v24 := GenS1{F0: 759}
                    _ = fi23
                    _ = v24
                }
                if (18446744073709551615 <= 0) {
                    var v25 float64 = (0.1 / 0.1)
                    v4 = append(v4, ((("\n\t" + "foo bar") + "hello") + "a"))
                    _ = v25
                }
                switch "hello" {
                case "":
                    var v26 string = ("\n\t" + "foo bar")
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        v6 += 1
                        var v28 int32 = int32(v3)
                        switch -2147483648 {
                        case -272:
                            v28 = v28
                        case 879:
                            var v29 int32 = (v28 ^ v28)
                            _ = v29
                        default:
                            if ((v12 / ((v12) | 1)) >= 42) {
                                switch (("a" + "hello") + "hello") {
                                case "a":
                                    delete(v5, 4294967295)
                                }
                                switch ("foo bar" + ((("a" + "hello") + "hello") + "a")) {
                                case "":
                                    switch 4294967295 {
                                    case -2147483648:
                                        var v30 float64 = -2.25
                                        _ = v30
                                    case 317:
                                        v12 -= v12
                                    default:
                                        var v31 float64 = (1.5 / ((-2.25 + ((((-2.25 + 0.1) - (((0.1 + 0.1) + (float64(1e10) - float64(-2.25))) + 1.5)) / (1e10 - 1.5)) + 0.1)) - (3.14159 / 3.14159)))
                                        _ = v31
                                    }
                                default:
                                    switch (("foo bar" + "hello") + "") {
                                    case "":
                                        var v32 float32 = 3.14159
                                        _ = v32
                                    default:
                                        var v33 float64 = float64(0.1)
                                        _ = v33
                                    }
                                }
                                v12 = (-(v12))
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            }
            var v34 string = "hello"
            _ = fi22
            _ = v34
        }
        _ = v21
    }
    var v35 int = v2
    v4 = append(v4, ("foo bar" + "a"))
    switch (v2 << ((v1) & 63)) {
    case 0:
        v36 := GenS2{F0: ("a" + ("hello" + "foo bar")), F1: 2}
        _ = v36
    }
    switch 4294967295 {
    case 770:
        var v37 int8 = 0
        _ = v37
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v3, v2, v6)
    }
    if true {
        switch "hello" {
        case "foo bar":
            v38 := v5[1]
            _ = v38
        }
    } else {
        v39 := GenS3{F0: (-634 == (-(v12))), F1: 1000}
        v40 := GenS2{F0: "foo bar", F1: 2147483648}
        _ = v39
        _ = v40
    }
    var v41 uint8 = 200
    v12 = v12
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v5)
    v42 := GenS4{F0: 4294967295, F1: (v1 << (((^(v1))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
