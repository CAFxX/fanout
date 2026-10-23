package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 string }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int64, p1 uint, p2 int32) int64 {
    var v5 int32 = p2
    fmt.Println(v5, p0, p1)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 595, F1: -1, F2: 0}
        v8 := GenS2{F0: 2, F1: -128, F2: ("foo bar" + "foo bar")}
        p2 = (p2 & p2)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 bool, p1 uint32) uint8 {
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS2{F0: p1, F1: -1, F2: ""}
        _ = fi9
        _ = v10
    }
    fmt.Println(p1, p1)
    p0 = (121 == 0)
    p1 = (p1 << ((1) & 63))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 83
}

func main() {
    var v1 float64 = 1.5
    var v2 int64 = 0
    var v3 uint64 = 464
    v4 := []string{}
    switch "hello" {
    case "foo bar":
        v3 -= (v3 >> ((4294967295) & 63))
    case "a":
        v2 = 1
    }
    fmt.Println(v2, v1)
    v11 := GenS2{F0: 2147483647, F1: -128, F2: "a"}
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            switch "hello" {
            case "foo bar":
                v13 := GenS3{F0: true, F1: "hello", F2: 255}
                _ = v13
            case "hello":
                if false {
                    v14 := GenS1{F0: 326, F1: 1, F2: (v3 - v3)}
                    _ = len(v4)
                    v3 -= 942
                    _ = v14
                }
            }
            if (!((0 < 2) || ((-2.25 / float32(v1)) == (0.0 / (0.0 - 0.1))))) {
                v2 = 1
            } else {
                if false {
                    var v15 bool = (true && (!false))
                    _ = len(v4)
                    var v16 uint = 42
                    _ = v15
                    _ = v16
                }
            }
            _ = fi12
        }
        v17 := GenS4{F0: 2, F1: 42}
        v3 = v3
        _ = v17
    }
    v18 := GenS3{F0: true, F1: ("hello" + ""), F2: 0}
    fmt.Println(v1, v2)
    var v19 string = ("foo bar" + "foo bar")
    for fi20 := 0; fi20 < 3; fi20++ {
        v1 = (1e10 + v1)
        if (!(v2 >= 928)) {
            var v21 int16 = -498
            switch (fi20 + -56) {
            case -603:
                if true {
                    fmt.Println(v19, v19)
                }
            }
            _ = v21
        } else {
            var v22 uint = 4294967295
            if false {
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 int32 = -933
                    if true {
                        var v25 bool = ((v3 | v3) <= v3)
                        var v26 uint32 = 1
                        v25 = v25
                        var v27 int = (fi23 | 9223372036854775807)
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi23
                    _ = v24
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 float32 = 1e10
                    v30 := GenS5{F0: 2}
                    v4 = append(v4, v19)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            _ = v22
        }
        var v31 float32 = 3.14159
        _ = fi20
        _ = v31
    }
    var v32 string = (v19 + v19)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 42 {
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -418:
        var v33 uint16 = 1
        _ = v33
    default:
        for fi34 := 0; fi34 < 16; fi34++ {
            for fi35 := 0; fi35 < 11; fi35++ {
                _ = len(v4)
                switch "hello" {
                case "hello":
                    var v36 uint = 0
                    _ = v36
                }
                var v37 int8 = 127
                v38 := GenS1{F0: 65535, F1: 939, F2: 968}
                _ = fi35
                _ = v37
                _ = v38
            }
            _ = fi34
        }
    }
    if ((v3 << ((382) & 63)) >= (v3 >> ((217) & 63))) {
        v39 := GenS5{F0: 2}
        for fi40 := 0; fi40 < 3; fi40++ {
            fmt.Println(v19)
            v41 := GenS6{F0: 63, F1: 677}
            v1 -= (1.5 / v1)
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    if false {
        v2 += (v2 * v2)
    } else {
        fmt.Println(v1, v2, v1)
        var v42 float64 = 0.1
        _ = v42
    }
    fmt.Println(v3, v1)
    fmt.Println(v2, v1, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
