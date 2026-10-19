package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
func genFn0(p0 int16, p1 int32) string {
    v4 := []int{-686, -456, 127}
    if true {
        v4 = append(v4, -507)
    } else {
        var v5 uint16 = 571
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v4 = append(v4, 127)
            _ = fi7
        }
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 2147483648, F1: (("a" + "") + "a"), F2: 18446744073709551615}
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    return ("\n\t" + "a")
}
func genFn1(p0 uint32) uint32 {
    p0 = p0
    var v10 int = 80
    var v11 bool = true
    if true {
        v12 := []int{v10, v10, -2147483648, 127}
        for fi13 := 0; fi13 < 4; fi13++ {
            v10 = v10
            _ = len(v12)
            p0 -= p0
            _ = fi13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            switch (v10 + (fi14 - (fi14 | (v10 & -2147483648)))) {
            case -2147483648:
                if v11 {
                    switch 2147483647 {
                    case 218:
                        v10 = (^(v10))
                    }
                }
            case -769:
                v10 = -2147483648
            case 1:
                switch ((("" + "") + ("foo bar" + "")) + "a") {
                case "foo bar":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v10 += (fi15 >> ((255) & 63))
                        var v16 uint8 = 42
                        var v17 int = 2147483647
                        fmt.Println(fi14, v17)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case "hello":
                    fmt.Println(v11, v11, p0)
                }
            }
            v18 := GenS1{F0: (p0 + (p0 * p0)), F1: "", F2: 65535}
            _ = fi14
            _ = v18
        }
        _ = v12
    } else {
        v19 := GenS2{F0: -190, F1: 1000, F2: -32768}
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 uint = 484
        v10 = v10
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v10
    _ = v11
    return p0
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            fmt.Println(v1, fi22)
            v23 := GenS2{F0: 42, F1: 22, F2: 42}
            var v24 uint8 = 252
            _ = v23
            _ = v24
        } else {
            var v25 uint32 = 65535
            switch ("" + (("" + "foo bar") + "foo bar")) {
            case "a":
                v1 = v1
            case "foo bar":
                switch (((("a" + "\n\t") + "\n\t") + ("hello" + "a")) + "hello") {
                case "hello":
                    if (v25 > (v25 ^ v25)) {
                        v26 := []int{fi22, fi22, (fi22 - 601), 32768}
                        _ = v26
                    }
                case "\n\t":
                    var v27 uint = 42
                    _ = v27
                case "a":
                    _ = len(v2)
                default:
                    v28 := GenS1{F0: (v25 / (((v25 >> ((1) & 63))) | 1)), F1: "foo bar", F2: 2}
                    _ = v28
                }
            case "hello":
                for fi29 := 0; fi29 < 9; fi29++ {
                    fmt.Println(fi22, v1)
                    var v30 string = ("" + "a")
                    v30 = "\n\t"
                    _ = fi29
                    _ = v30
                }
            }
            _ = v25
        }
        v31 := []int{(fi22 << ((255) & 63)), (fi22 | (fi22 & (fi22 / ((0) | 1)))), fi22, fi22}
        _ = fi22
        _ = v31
    }
    v32 := GenS1{F0: 159, F1: ("a" + "a"), F2: 42}
    v1 = (v1 + 3)
    var v33 string = (("\n\t" + "") + "a")
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 string = v33
        _ = fi34
        _ = v35
    }
    var v36 uint16 = 65535
    if false {
        var v37 int8 = (v1 * v1)
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := []int{(fi38 ^ (-(fi38))), (fi38 % (((fi38 * (-(fi38)))) | 1)), -9223372036854775808, fi38}
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        if true {
            _ = len(v3)
            if false {
                var v40 uint64 = 4294967295
                _ = v40
            }
            _ = len(v3)
        } else {
            var v41 uint = 255
            switch "a" {
            case "foo bar":
                var v42 int = 4294967295
                _ = v42
            case "\n\t":
                var v43 float32 = -2.25
                _ = v43
            case "":
                v36 = (v36 >> ((382) & 63))
            default:
                if (((2 > 913) && (!true)) && true) {
                    var v44 uint32 = 2
                    v45 := []int{1, 0, 32767}
                    _ = v44
                    _ = v45
                }
            }
            _ = v41
        }
    }
    v46 := []int{873, -200}
    var v47 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v46
    _ = v47
}
