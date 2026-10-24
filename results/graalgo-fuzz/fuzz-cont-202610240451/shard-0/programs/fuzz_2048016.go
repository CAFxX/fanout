package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint; F2 bool }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int8, p2 bool) int16 {
    if ((p1 ^ (p1 | p1)) != (^(p1))) {
        p2 = true
        v7 := GenS1{F0: "a"}
        _ = v7
    } else {
        v8 := GenS1{F0: (("" + ("a" + "a")) + "\n\t")}
        if (!(42 > 18446744073709551615)) {
            for fi9 := 0; fi9 < 17; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    p1 = -128
                    v11 := GenS1{F0: ("" + "hello")}
                    var v12 uint32 = (p0 | (p0 + 1))
                    fmt.Println(fi9, v12, fi10)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
            if (0 < 1) {
                p1 += p1
            }
        } else {
            v13 := GenS1{F0: "hello"}
            fmt.Println(p2, p0)
            if false {
                v14 := GenS2{F0: 65535}
                _ = v14
            }
            _ = v13
        }
        fmt.Println(p0, p2, p2)
        _ = v8
    }
    v15 := []int{127, 32767, 4294967295, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -1
}
func genFn1(p0 uint, p1 uint64) string {
    if (!(p1 > 4294967295)) {
        var v16 uint32 = 561
        v17 := []int{0}
        v17 = append(v17, 127)
        _ = v16
        _ = v17
    } else {
        p0 = p0
        fmt.Println(p1, p1)
        var v18 uint16 = 42
        _ = v18
    }
    if true {
        fmt.Println(p1)
        p0 = uint(p1)
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int8 = -128
            _ = fi19
            _ = v20
        }
    } else {
        v21 := []int{}
        if len(v21) > 0 {
            v21[0] = 692
        }
        if (true && (9223372036854775807 != -607)) {
            p0 += (^(p0))
        }
        _ = v21
    }
    _ = p0
    _ = p1
    return (("" + "a") + "")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 1000
    var v3 int32 = 99
    var v4 uint32 = 146
    var v5 uint32 = (v4 << ((uint(v4)) & 63))
    v6 := []string{"\n\t", ("hello" + "\n\t"), ("a" + ("a" + "foo bar"))}
    for fi22 := 0; fi22 < 8; fi22++ {
        _ = len(v1)
        _ = len(v1)
        _ = fi22
    }
    switch 0 {
    case -193:
        v23 := GenS3{F0: "\n\t"}
        _ = v23
    default:
        _ = len(v1)
    }
    v5 -= v5
    switch "" {
    case "a":
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS1{F0: "hello"}
            if (1.5 == 0.1) {
                switch (-(fi24)) {
                case 4294967295:
                    v26 := GenS2{F0: 77}
                    _ = v26
                case -507:
                    delete(v1, 439)
                }
                for fi27 := 0; fi27 < 5; fi27++ {
                    v28 := GenS2{F0: 487}
                    v29 := GenS4{F0: (42 <= -398), F1: 4294967295, F2: false}
                    switch ("a" + "\n\t") {
                    case "a":
                        fmt.Println(v3)
                    case "hello":
                        fmt.Println(v5, v2)
                    case "":
                        v3 -= (v3 + (v3 % (((v3 ^ (-(v3)))) | 1)))
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v30 := GenS4{F0: (!(true && (!false))), F1: 1, F2: ((^(v4)) < 2)}
                _ = v30
            }
            _ = fi24
            _ = v25
        }
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v31 int8 = 100
    var v32 int32 = (v3 >> ((629) & 63))
    if (!true) {
        v33 := GenS5{F0: 15}
        var v34 uint8 = 16
        switch "a" {
        case "hello":
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 uint8 = (v34 & 42)
                _ = fi35
                _ = v36
            }
        case "foo bar":
            var v37 int = -866
            _ = v37
        case "":
            if ((!true) || (false && (!true))) {
                v2 = 641
            }
        default:
            v5 += (^(v5))
        }
        _ = v33
        _ = v34
    }
    v38 := GenS5{F0: 1}
    var v39 int16 = v2
    var v40 int32 = -124
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
