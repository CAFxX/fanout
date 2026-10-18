package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint32, p1 uint16, p2 bool) int16 {
    v5 := []int{32767, 127}
    p0 = 2147483648
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(p1, p0)
        _ = fi6
    }
    fmt.Println(p1)
    switch "a" {
    case "":
        _ = len(v5)
    case "hello":
        var v7 string = ("\n\t" + "foo bar")
        _ = v7
    case "foo bar":
        var v8 uint16 = (p1 | p1)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1000
}
func genFn1(p0 int, p1 int8, p2 string) bool {
    v9 := GenS1{F0: 32767, F1: p2}
    switch (p0 << ((382) & 63)) {
    case 32767:
        if (!(!(0.0 > 0.0))) {
            for fi10 := 0; fi10 < 17; fi10++ {
                p0 -= (fi10 ^ (p0 ^ (fi10 * (fi10 % ((fi10) | 1)))))
                _ = fi10
            }
        } else {
            if false {
                v11 := GenS1{F0: 6, F1: p2}
                _ = v11
            }
            v12 := []int{p0, p0, (p0 % ((-1) | 1))}
            _ = v12
        }
    case 4294967295:
        p0 = (p0 | p0)
    default:
        var v13 int8 = 127
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (1 >= p0)
}

func main() {
    var v1 string = ("a" + ("foo bar" + (("" + ((("a" + "hello") + "foo bar") + "\n\t")) + "")))
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 string = v1
    for fi14 := 0; fi14 < 1; fi14++ {
        v1 = (v1 + v1)
        var v15 uint64 = 0
        v1 = v1
        _ = fi14
        _ = v15
    }
    if (true && false) {
        fmt.Println(v1, v4)
        v16 := []int{}
        _ = v16
    } else {
        v17 := GenS1{F0: 65535, F1: v4}
        v18 := GenS1{F0: 735, F1: "a"}
        _ = v17
        _ = v18
    }
    v1 = ""
    delete(v2, ("\n\t" + v1))
    if (0.0 <= 0.0) {
        v1 = v4
    } else {
        var v19 int8 = 127
        if (42 > -1) {
            var v20 int32 = 13
            v21 := GenS2{F0: 581, F1: 354, F2: 744}
            _ = v20
            _ = v21
        }
        v19 += (v19 << ((255) & 63))
        _ = v19
    }
    if false {
        for fi22 := 0; fi22 < 16; fi22++ {
            v1 = v1
            v4 = (v1 + "foo bar")
            _ = fi22
        }
        if (true && true) {
            if (!false) {
                v1 = v1
                v1 = ""
                v1 = v1
            } else {
                v23 := GenS1{F0: 32767, F1: "a"}
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 uint8 = 0
                    var v26 uint16 = 65535
                    fmt.Println(v1, v25)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v27 := GenS1{F0: 65535, F1: ((("hello" + ("a" + (v1 + v1))) + ("" + v4)) + ((v1 + "hello") + (((v4 + "hello") + v4) + ((("" + "hello") + v4) + v1))))}
                _ = v23
                _ = v27
            }
            v28 := GenS3{F0: 933, F1: 250}
            v29 := []int{2147483647, 42}
            _ = v28
            _ = v29
        } else {
            v1 = v4
            v30 := GenS4{F0: 255, F1: 65535}
            var v31 float64 = 0.0
            _ = v30
            _ = v31
        }
    } else {
        if (1e10 > 0.1) {
            v2["\n\t"] = true
            v1 = "\n\t"
            v32 := GenS1{F0: 65535, F1: "\n\t"}
            _ = v32
        }
    }
    fmt.Println(v4, v1)
    var v33 int32 = 322
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 uint32 = 0
        switch v4 {
        case "":
            var v36 uint64 = 383
            _ = v36
        case "a":
            var v37 int = (fi34 * fi34)
            _ = v37
        default:
            v38 := GenS1{F0: 0, F1: v1}
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
