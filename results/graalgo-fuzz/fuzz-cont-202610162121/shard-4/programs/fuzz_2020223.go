package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint) bool {
    var v6 int = -596
    fmt.Println(v6)
    if (127 >= v6) {
        var v7 uint8 = 223
        _ = v7
    }
    _ = p0
    _ = v6
    return ((true || true) || false)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -53
    var v3 uint16 = 703
    var v4 int = int(v3)
    var v5 int32 = 1
    _ = len(v1)
    switch "foo bar" {
    case "":
        switch (v4 - v4) {
        case -323:
            fmt.Println(v5, v2, v2)
        case -2147483648:
            if (-2.25 > 3.14159) {
                fmt.Println(v4, v2, v3)
                v8 := []int{(^(v4)), v4}
                v3 = v3
                _ = v8
            } else {
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch (fi9 >> ((255) & 63)) {
                    case 42:
                        if false {
                            v10 := GenS1{F0: 243}
                            _ = v10
                        } else {
                            if true {
                                fmt.Println(v4, v5, v5)
                            }
                        }
                    }
                    fmt.Println(v3)
                    fmt.Println(v2, v2, v4)
                    _ = fi9
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    if (4294967295 < 1) {
                        var v12 string = "a"
                        v13 := v1["hello"]
                        fmt.Println(v12, v5)
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(fi11)
                        var v14 int32 = v5
                        _ = v14
                    }
                    _ = fi11
                }
            }
        case 2147483647:
            fmt.Println(v2)
        default:
            v15 := []int{32768, (v4 >> ((0) & 63))}
            _ = v15
        }
    case "hello":
        v2 = v2
    case "\n\t":
        var v16 uint32 = 2
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v4 -= fi17
        var v18 float64 = (1e10 / (1.5 + 0.0))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v5 -= v5
        v1["hello"] = true
        v4 = (-(v4))
        _ = fi19
    }
    v2 += (-(v2))
    v3 -= v3
    fmt.Println(v5)
    v5 -= (-(v5))
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS1{F0: 936}
        _ = fi20
        _ = v21
    }
    v4 = (v4 / ((319) | 1))
    var v22 int8 = v2
    _ = len(v1)
    switch (v4 | 0) {
    case 128:
        v5 += v5
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v24 := []int{fi23, v4, v4}
        v25 := GenS2{F0: true, F1: (fi23 & fi23)}
        switch (-(fi23)) {
        case -612:
            switch "a" {
            case "\n\t":
                var v26 int16 = 1000
                _ = v26
            default:
                switch 992 {
                case 32767:
                    switch ("foo bar" + "\n\t") {
                    case "":
                        var v27 int16 = -32768
                        _ = v27
                    default:
                        var v28 string = (("hello" + "foo bar") + "")
                        _ = v28
                    }
                case -2147483648:
                    for fi29 := 0; fi29 < 10; fi29++ {
                        var v30 int8 = (v2 ^ (^(v22)))
                        v31 := GenS2{F0: (true || (((!false) && ("hello" == "\n\t")) || false)), F1: v4}
                        v32 := GenS3{F0: 42, F1: ("a" + "\n\t"), F2: 255}
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                default:
                    var v33 uint64 = 18446744073709551615
                    _ = v33
                }
            }
        case 779:
            if true {
                v34 := GenS4{F0: "foo bar", F1: v2}
                if (false || true) {
                    v3 += v3
                    v35 := GenS5{F0: -118}
                    _ = v35
                }
                _ = v34
            } else {
                var v36 bool = (83 <= 2)
                _ = v36
            }
        case 0:
            fmt.Println(v4, v2)
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        fmt.Println(v22, v3)
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
