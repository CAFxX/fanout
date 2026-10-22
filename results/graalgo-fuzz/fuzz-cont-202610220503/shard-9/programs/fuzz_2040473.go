package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0() string {
    var v7 uint8 = 251
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int = (fi8 + fi8)
            switch "a" {
            case "\n\t":
                switch "a" {
                case "a":
                    v11 := GenS1{F0: 0, F1: 535, F2: 456}
                    _ = v11
                case "foo bar":
                    v12 := GenS2{F0: v7, F1: 151, F2: "a"}
                    _ = v12
                default:
                    v13 := GenS2{F0: 139, F1: 1, F2: "foo bar"}
                    _ = v13
                }
            case "hello":
                fmt.Println(v10)
            }
            v14 := GenS3{F0: 42}
            v15 := GenS2{F0: (^(v7)), F1: 4294967295, F2: "hello"}
            _ = fi9
            _ = v10
            _ = v14
            _ = v15
        }
        v16 := []int{}
        _ = fi8
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := []int{fi17, -14, fi17, (fi17 ^ fi17)}
        if (366 > 42) {
            v7 -= 0
            v19 := GenS4{F0: v7}
            switch "" {
            case "foo bar":
                v7 -= (v7 % ((v7) | 1))
            }
            _ = v19
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                if ((2147483647 > -954) || false) {
                    var v21 int32 = 42
                    _ = v21
                } else {
                    var v22 int32 = -267
                    v7 = (v7 % (((v7 - (v7 - (v7 * (^(v7)))))) | 1))
                    v23 := GenS5{F0: 322, F1: 610}
                    v24 := GenS5{F0: 1, F1: (fi17 ^ (fi20 / ((-166) | 1)))}
                    switch ("foo bar" + ("hello" + "hello")) {
                    case "\n\t":
                        var v25 int = (fi17 | fi20)
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v26 int64 = 0
                _ = fi20
                _ = v26
            }
            var v27 uint64 = 793
            var v28 uint8 = (v7 >> ((798) & 63))
            _ = v27
            _ = v28
        }
        _ = fi17
        _ = v18
    }
    v7 = v7
    _ = v7
    return "\n\t"
}

func main() {
    var v1 bool = (1 <= 223)
    var v2 uint32 = 42
    v3 := []string{("foo bar" + "foo bar")}
    v4 := []string{}
    var v5 int8 = -112
    var v6 uint = 675
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v29 int16 = 0
    v30 := GenS4{F0: 2}
    _ = len(v4)
    v31 := GenS6{F0: v2}
    var v32 uint32 = v2
    if v1 {
        fmt.Println(v6)
        fmt.Println(v1)
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        if (0 == v5) {
            if len(v3) > 0 {
                v3[0] = ""
            }
        } else {
            var v34 int32 = -1
            var v35 float64 = 0.0
            _ = v34
            _ = v35
        }
        _ = fi33
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 17; fi37++ {
            fmt.Println(v1, v2)
            v38 := GenS7{F0: 882, F1: (("hello" + "hello") + ("\n\t" + "hello")), F2: 0}
            v2 = 667
            _ = fi37
            _ = v38
        }
        var v39 uint32 = 4294967295
        _ = fi36
        _ = v39
    }
    fmt.Println(v29, v1)
    v3 = append(v3, ("hello" + ("foo bar" + (("a" + "hello") + ""))))
    v6 += v6
    v2 = v32
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v40 uint32 = (v32 << ((v6) & 63))
    v41 := GenS4{F0: 127}
    v42 := GenS7{F0: 18446744073709551615, F1: "hello", F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
