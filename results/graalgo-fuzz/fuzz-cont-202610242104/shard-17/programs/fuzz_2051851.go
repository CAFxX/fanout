package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 string; F2 int32 }
func genFn0() uint32 {
    v4 := []int{42}
    var v5 string = "hello"
    v6 := GenS1{F0: (v5 == ""), F1: -2147483648}
    _ = v4
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1(p0 uint64, p1 int64, p2 int8) int16 {
    p1 = p1
    var v7 int32 = -94
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint8 = 0
    var v3 float64 = 1e10
    v1 = v1
    v8 := []int{}
    v9 := GenS2{F0: -808}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v8 = append(v8, 42)
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int = (-(fi10))
        v1 = v1
        switch ("hello" + ("foo bar" + "hello")) {
        case "hello":
            v12 := GenS1{F0: v1, F1: 0}
            _ = v12
        case "foo bar":
            v13 := GenS2{F0: -637}
            _ = v13
        case "a":
            var v14 int = 730
            _ = v14
        default:
            var v15 uint64 = 18446744073709551615
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    v16 := GenS3{F0: ((("hello" + "a") + ("" + "")) + "a")}
    var v17 int32 = 0
    if ((v2 | (^(v2))) == (v2 | 247)) {
        if (true && false) {
            fmt.Println(v2, v17, v1)
            var v18 string = ((((("" + (("a" + "") + "foo bar")) + "hello") + "foo bar") + ("" + "\n\t")) + "foo bar")
            v8 = append(v8, 1)
            _ = v18
        } else {
            v19 := GenS4{F0: 2147483647}
            fmt.Println(v3, v17)
            _ = v19
        }
        if false {
            var v20 bool = true
            if (v3 >= 0.0) {
                var v21 int16 = -45
                if (v3 <= (1e10 / v3)) {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v1 = v20
                        var v23 uint8 = (v2 | v2)
                        if (1 > 156) {
                            for fi24 := 0; fi24 < 18; fi24++ {
                                if v20 {
                                    v25 := GenS4{F0: fi24}
                                    v17 = (v17 - (v17 * v17))
                                    fmt.Println(v3)
                                    _ = v25
                                }
                                v26 := GenS3{F0: "\n\t"}
                                _ = fi24
                                _ = v26
                            }
                            switch "a" {
                            case "foo bar":
                                for fi27 := 0; fi27 < 13; fi27++ {
                                    var v28 uint32 = 42
                                    _ = fi27
                                    _ = v28
                                }
                            case "hello":
                                v29 := GenS3{F0: "a"}
                                _ = v29
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                _ = v21
            }
            _ = v20
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        var v30 float32 = (0.1 * (1.5 * float32(float64(0.1))))
        _ = v30
    }
    v17 -= 2147483647
    switch 128 {
    case 32768:
        for fi31 := 0; fi31 < 2; fi31++ {
            switch (fi31 - 0) {
            case 32768:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case -693:
                if true {
                    v1 = v1
                } else {
                    _ = len(v8)
                    v32 := GenS5{F0: genFn1(597, 1, -64), F1: "a", F2: v17}
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v1 = (2 > 1)
                        _ = fi33
                    }
                    _ = v32
                }
            default:
                switch fi31 {
                case 1:
                    v8 = append(v8, fi31)
                }
            }
            switch "foo bar" {
            case "hello":
                fmt.Println(fi31)
            }
            v8 = append(v8, fi31)
            _ = fi31
        }
    case -2147483648:
        var v34 float64 = v3
        _ = v34
    default:
        v35 := GenS1{F0: v1, F1: v17}
        _ = v35
    }
    var v36 uint8 = v2
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 int32 = 1
        if v1 {
            fmt.Println(v17, fi37, v38)
            _ = len(v8)
        }
        var v39 uint64 = 2
        _ = fi37
        _ = v38
        _ = v39
    }
    v2 -= (v36 | v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v8
    _ = v9
}
