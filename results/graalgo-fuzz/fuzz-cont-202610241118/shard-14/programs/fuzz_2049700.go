package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 string; F2 uint8 }
type GenS3 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 string, p1 uint16) uint {
    switch p0 {
    case "\n\t":
        p0 = ((p0 + ("" + (p0 + p0))) + "\n\t")
    case "a":
        if true {
            var v7 float64 = (-2.25 / 3.14159)
            _ = v7
        }
    default:
        v8 := GenS1{F0: p0}
        _ = v8
    }
    fmt.Println(p1, p1)
    v9 := GenS1{F0: p0}
    p1 += p1
    p0 = "a"
    _ = p0
    _ = p1
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 bool) uint {
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 float64 = float64(3.14159)
        v12 := []int{}
        v11 -= v11
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{1}
    if (!((!((float32(-2.25) - (1e10 / -2.25)) > (0.1 - float32(1.5)))) || ((!((p0 - p0) == (^(p0)))) || (!((!p1) && (0 > 255)))))) {
        if (!false) {
            switch "\n\t" {
            case "":
                if (2 == 734) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                } else {
                    v14 := GenS2{F0: 18446744073709551615, F1: "a", F2: 127}
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v16 := GenS1{F0: "foo bar"}
                        _ = fi15
                        _ = v16
                    }
                    var v17 uint64 = 0
                    _ = v14
                    _ = v17
                }
            case "hello":
                for fi18 := 0; fi18 < 9; fi18++ {
                    switch "foo bar" {
                    case "":
                        fmt.Println(p0)
                    case "a":
                        p0 += p0
                    case "\n\t":
                        v19 := GenS1{F0: "hello"}
                        _ = v19
                    default:
                        p1 = true
                    }
                    v20 := GenS1{F0: "hello"}
                    _ = fi18
                    _ = v20
                }
            }
            var v21 int64 = -661
            _ = len(v13)
            _ = v21
        } else {
            p0 = p0
        }
        switch "a" {
        case "a":
            var v22 int = -816
            _ = v22
        default:
            p1 = (p1 && (genFn0(("\n\t" + "foo bar"), p0) >= 4294967295))
        }
        var v23 float64 = (1e10 - (float64(3.14159) - (3.14159 + -2.25)))
        _ = v23
    }
    v24 := GenS3{F0: -9223372036854775808, F1: 255, F2: 2}
    _ = p0
    _ = p1
    _ = v13
    _ = v24
    return 42
}

func main() {
    var v1 string = ("" + "\n\t")
    var v2 int8 = -128
    var v3 int16 = 1000
    v4 := make(map[uint32]string)
    var v5 uint8 = 83
    var v6 uint16 = 859
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS3{F0: fi25, F1: (v5 & (v5 + v5)), F2: v5}
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: 675, F1: ("\n\t" + ((v1 + v1) + v1)), F2: 104}
    switch 399 {
    case 660:
        for fi28 := 0; fi28 < 5; fi28++ {
            switch fi28 {
            case 32768:
                var v29 int8 = (v2 << ((65535) & 63))
                _ = v29
            }
            _ = fi28
        }
    default:
        var v30 int8 = 30
        _ = v30
    }
    fmt.Println(v6, v1)
    v2 += v2
    switch "" {
    case "hello":
        var v31 uint16 = 65535
        _ = v31
    case "":
        if (-50 == 36) {
            v32 := GenS3{F0: 737, F1: v5, F2: (v5 ^ v5)}
            _ = v32
        }
    case "\n\t":
        fmt.Println(v6, v6)
    }
    v4[2] = ((v1 + (("" + "") + "a")) + "\n\t")
    var v33 float64 = 0.0
    var v34 bool = (genFn0(v1, uint16(v5)) < genFn1(v6, true))
    switch 42 {
    case 0:
        if v34 {
            switch v1 {
            case "":
                v5 -= v5
            default:
                v6 += (v6 + v6)
            }
            if v34 {
                v3 += 1
                var v35 float32 = -2.25
                v33 = v33
                _ = v35
            }
            switch (v1 + (v1 + (v1 + ""))) {
            case "":
                switch v1 {
                case "\n\t":
                    var v36 uint = 2
                    _ = v36
                case "hello":
                    fmt.Println(v34, v33)
                default:
                    for fi37 := 0; fi37 < 5; fi37++ {
                        var v38 int = 32767
                        fmt.Println(v6, v3)
                        _ = fi37
                        _ = v38
                    }
                }
            }
        } else {
            v39 := GenS1{F0: v1}
            var v40 int64 = -16
            _ = v39
            _ = v40
        }
    default:
        v41 := []int{-474, 4294967295, -446, 1}
        _ = v41
    }
    var v42 int16 = (v3 + 181)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
