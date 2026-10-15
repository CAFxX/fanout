package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 bool }
func genFn0(p0 bool, p1 string, p2 int8) uint16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch 42 {
        case -1:
            switch (fi5 & int(p2)) {
            case 32767:
                p2 = -67
            case 42:
                var v6 bool = (!false)
                _ = v6
            case -1:
                v7 := []int{fi5, fi5, (fi5 << ((65535) & 63)), fi5}
                _ = v7
            }
        case 437:
            fmt.Println(p2)
        case 42:
            v8 := GenS1{F0: 4294967295, F1: "hello", F2: p1}
            _ = v8
        }
        _ = fi5
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{(fi9 & 127), fi9, 127}
        fmt.Println(p1, fi9)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []int{127}
    var v2 int8 = -128
    v3 := []uint32{2, 65535, 2147483648}
    var v4 string = "a"
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: 577, F1: "a", F2: v4}
        fmt.Println(fi11)
        if (1.5 != 1.5) {
            v13 := GenS1{F0: 65535, F1: v4, F2: v4}
            _ = v13
        } else {
            v14 := GenS1{F0: 1, F1: (v4 + (("a" + (v4 + "")) + v4)), F2: ("\n\t" + ((v4 + v4) + v4))}
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v2 += 24
    v2 += (v2 / ((1) | 1))
    if (4 == 0) {
        var v15 uint32 = 2147483647
        var v16 bool = (!true)
        var v17 int = -2147483648
        _ = v15
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            v2 += (v2 ^ v2)
            fmt.Println(fi18)
            _ = fi18
        }
        if true {
            v4 = v4
        }
        _ = len(v3)
    }
    v4 = v4
    fmt.Println(v4, v2)
    v4 = ("" + "foo bar")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(fi19, v4, v4)
            var v21 bool = true
            var v22 string = v4
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    switch -1 {
    case 127:
        var v23 uint8 = 128
        _ = v23
    default:
        var v24 uint64 = 16
        _ = v24
    }
    if ((1 >= 4294967295) && false) {
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS1{F0: 2147483648, F1: "hello", F2: v4}
                v28 := GenS1{F0: 2, F1: (v4 + "\n\t"), F2: (((v4 + (v4 + v4)) + v4) + "")}
                fmt.Println(fi25, fi25, v4)
                _ = fi26
                _ = v27
                _ = v28
            }
            v29 := GenS2{F0: 9223372036854775807, F1: int8(fi25), F2: 2}
            _ = fi25
            _ = v29
        }
        if (0 >= 111) {
            var v30 float32 = 3.14159
            fmt.Println(v30)
            _ = v30
        }
        v1 = append(v1, 0)
    } else {
        fmt.Println(v2, v4)
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 int = (^(fi31))
            _ = fi31
            _ = v32
        }
    }
    v4 = ((v4 + "\n\t") + "hello")
    if false {
        if false {
            fmt.Println(v4, v2)
            var v33 float64 = (float64(1.5) * 1e10)
            _ = v33
        } else {
            v4 = "a"
            for fi34 := 0; fi34 < 11; fi34++ {
                _ = len(v3)
                switch ("foo bar" + v4) {
                case "\n\t":
                    var v35 string = v4
                    _ = v35
                case "":
                    v2 = (^(v2))
                default:
                    _ = len(v1)
                }
                _ = fi34
            }
        }
        fmt.Println(v4, v2)
        var v36 uint64 = 42
        _ = v36
    } else {
        if false {
            v37 := GenS3{F0: 0, F1: ((((v4 + v4) + v4) + ("foo bar" + v4)) + (v4 + v4)), F2: false}
            if false {
                if false {
                    v38 := GenS1{F0: 2, F1: ((v4 + (v4 + (v4 + "a"))) + ("a" + "hello")), F2: ((("hello" + v4) + "a") + v4)}
                    _ = v38
                } else {
                    var v39 uint64 = 4294967295
                    _ = v39
                }
                v40 := GenS2{F0: 9223372036854775807, F1: (-(v2)), F2: 127}
                var v41 float64 = 0.0
                _ = v40
                _ = v41
            }
            _ = v37
        } else {
            _ = len(v1)
            var v42 int64 = -9223372036854775808
            _ = v42
        }
        v2 = v2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
