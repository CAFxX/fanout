package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0(p0 int8, p1 int8, p2 uint8) uint64 {
    if (false || true) {
        var v7 bool = (false || (65535 >= 42))
        p1 += p0
        v8 := GenS1{F0: 207, F1: p1, F2: 1}
        _ = v7
        _ = v8
    }
    var v9 int16 = 32767
    fmt.Println(v9, p1)
    v10 := []int{549}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint = 1
    var v2 uint16 = 155
    var v3 int32 = 0
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 uint8 = 6
    v11 := []int{-2147483648}
    delete(v4, "hello")
    v11 = append(v11, 2147483647)
    v2 += (v2 & 2)
    v12 := GenS2{F0: -107, F1: 2}
    if false {
        if (true || true) {
            switch ("" + ("hello" + "hello")) {
            case "":
                v1 -= 1
            default:
                v11 = append(v11, -423)
            }
        }
        v6 = 128
    }
    v6 = 128
    if true {
        var v13 bool = ((v3 / ((148) | 1)) <= (v3 ^ (v3 / ((v3) | 1))))
        v14 := GenS1{F0: genFn0(-1, -128, (v6 | (v6 ^ 88))), F1: int8(v3), F2: v2}
        _ = v13
        _ = v14
    } else {
        v4["a"] = false
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    switch ("\n\t" + "a") {
    case "a":
        v3 = v3
    }
    var v15 int8 = 1
    v16 := GenS1{F0: 4294967295, F1: -125, F2: v2}
    switch 42 {
    case 4294967295:
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int64 = -1
            fmt.Println(v15, v15, fi17)
            _ = fi17
            _ = v18
        }
    case 729:
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 float64 = -2.25
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 3; fi22++ {
                if (float32((1e10 * 1.5)) < 0.0) {
                    if (((v3 >> ((18446744073709551615) & 63)) != (v3 * (v3 & (v3 << ((578) & 63))))) && (974 >= 127)) {
                        if true {
                            var v23 int16 = 907
                            _ = v23
                        } else {
                            if false {
                                var v24 float64 = -2.25
                                fmt.Println(fi22, v3)
                                _ = v24
                            } else {
                                switch fi19 {
                                case 32767:
                                    if (((!(!(false || (int64(v3) > -9223372036854775808)))) || (false && true)) || (3.14159 > 3.14159)) {
                                        v1 -= 269
                                        v1 -= uint(v6)
                                        var v25 uint32 = 42
                                        _ = v25
                                    }
                                }
                                v26 := GenS2{F0: (v15 - (v15 ^ -40)), F1: 42}
                                v6 = (v6 + 198)
                                _ = v26
                            }
                            switch "hello" {
                            case "hello":
                                fmt.Println(fi19, fi22)
                            case "":
                                if (!(!false)) {
                                    v2 = v2
                                    v11 = append(v11, fi19)
                                    v6 += 0
                                }
                            }
                        }
                    }
                }
                switch fi19 {
                case -466:
                    v5[(-(fi19))] = (2 <= v1)
                default:
                    fmt.Println(v3)
                    fmt.Println(v3)
                    v11 = append(v11, (fi19 << ((v1) & 63)))
                }
                _ = fi22
            }
            if true {
                switch "a" {
                case "hello":
                    var v27 uint32 = 42
                    _ = v27
                case "a":
                    v28 := GenS3{F0: 233, F1: (v15 << ((255) & 63))}
                    _ = v28
                case "\n\t":
                    if (false && ((fi19 == fi19) || false)) {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v30 := v4["a"]
                            v31 := GenS4{F0: (fi19 % (((-(fi29))) | 1)), F1: (fi19 - fi29)}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        delete(v5, 42)
                        v15 = (v15 * 42)
                    }
                default:
                    var v32 int32 = v3
                    _ = v32
                }
            }
            _ = fi19
        }
    default:
        switch ((("a" + "foo bar") + "\n\t") + ("" + "foo bar")) {
        case "\n\t":
            var v33 uint64 = 4294967295
            _ = v33
        case "a":
            for fi34 := 0; fi34 < 4; fi34++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v1 += (v1 + (v1 + 346))
                var v35 uint8 = (v6 - v6)
                _ = fi34
                _ = v35
            }
        case "":
            v36 := GenS5{F0: 2, F1: (!((true || (2147483648 == 2)) || true))}
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
