package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint8; F1 string; F2 string }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 int16, p1 uint8, p2 string) bool {
    var v5 uint64 = 0
    if ("hello" < p2) {
        v6 := GenS1{F0: 194, F1: 475, F2: 42}
        _ = v6
    } else {
        v7 := GenS2{F0: p1, F1: 9223372036854775807}
        switch p2 {
        case "a":
            fmt.Println(v5, v5, p0)
        case "foo bar":
            v8 := []int{9223372036854775807, -696}
            _ = v8
        }
        var v9 uint16 = uint16(v5)
        _ = v7
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            if ((!false) && false) {
                var v11 int8 = 100
                if true {
                    v12 := []int{fi10, (fi10 + fi10), (fi10 / (((fi10 << ((2) & 63))) | 1)), (fi10 | fi10)}
                    p0 = -32768
                    _ = v12
                } else {
                    switch (p2 + "\n\t") {
                    case "a":
                        v13 := GenS2{F0: 42, F1: 9223372036854775807}
                        _ = v13
                    case "":
                        v14 := []int{-876}
                        _ = v14
                    default:
                        var v15 uint8 = (p1 >> ((700) & 63))
                        _ = v15
                    }
                }
                var v16 string = (p2 + ((p2 + p2) + "foo bar"))
                _ = v11
                _ = v16
            }
            v17 := GenS1{F0: 851, F1: 2147483648, F2: v5}
            v18 := GenS3{F0: p0, F1: -575}
            _ = fi10
            _ = v17
            _ = v18
        }
        fmt.Println(p0, p2)
    }
    if true {
        var v19 float32 = float32((1.5 - float64(((0.0 * -2.25) / 1e10))))
        _ = v19
    } else {
        p0 = 32767
    }
    switch 127 {
    case -956:
        v20 := GenS3{F0: (^(p0)), F1: p0}
        _ = v20
    case 4294967295:
        v5 -= v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}
func genFn1() string {
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(fi21)
        _ = fi21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int64 = 566
        for fi24 := 0; fi24 < 10; fi24++ {
            switch "a" {
            case "hello":
                if true {
                    fmt.Println(fi22, fi22, fi22)
                } else {
                    var v25 uint8 = 2
                    var v26 int16 = 42
                    _ = v25
                    _ = v26
                }
            case "":
                var v27 int32 = -2147483648
                _ = v27
            case "\n\t":
                v28 := []int{(fi22 | 475), (fi24 % (((-(fi24))) | 1))}
                _ = v28
            }
            if (!true) {
                v29 := GenS4{F0: 42, F1: "\n\t", F2: ("foo bar" + "\n\t")}
                switch ("" + "hello") {
                case "a":
                    switch "hello" {
                    case "hello":
                        if (0 > fi22) {
                            var v30 int16 = -32768
                            v23 += 470
                            _ = v30
                        } else {
                            var v31 int8 = 42
                            v23 += (v23 + (v23 << ((4294967295) & 63)))
                            v32 := GenS2{F0: 128, F1: 785}
                            _ = v31
                            _ = v32
                        }
                    case "a":
                        if (true && false) {
                            var v33 string = "hello"
                            if false {
                                v34 := []int{(fi24 >> ((529) & 63)), fi22, fi24, fi22}
                                _ = v34
                            }
                            v23 -= (v23 ^ 9223372036854775807)
                            _ = v33
                        }
                    case "":
                        v23 += 439
                    default:
                        fmt.Println(v23, v23)
                    }
                default:
                    var v35 int8 = 0
                    _ = v35
                }
                var v36 int32 = 42
                v37 := GenS5{F0: 0, F1: (v36 + 475)}
                v38 := []int{(fi24 >> ((42) & 63)), (fi24 * (fi24 ^ 32768)), fi24, fi22}
                _ = v29
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi24
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            v23 -= (-(v23))
            _ = fi39
        }
        _ = fi22
        _ = v23
    }
    var v40 bool = (3.14159 > 3.14159)
    var v41 uint64 = 4294967295
    _ = v40
    _ = v41
    return "\n\t"
}

func main() {
    v1 := []string{"\n\t"}
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    v4 := []int{}
    fmt.Println(v2)
    switch (v2 + v2) {
    case "a":
        var v42 uint = 65535
        _ = v42
    }
    v2 = ""
    for fi43 := 0; fi43 < 4; fi43++ {
        _ = len(v1)
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 int = fi44
            var v46 float32 = 0.0
            v47 := GenS3{F0: -32768, F1: 23}
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        var v48 int16 = 42
        _ = fi43
        _ = v48
    }
    var v49 string = ((("\n\t" + v2) + v2) + v2)
    var v50 uint64 = 2
    v2 = ("a" + (v2 + (v2 + "\n\t")))
    var v51 int64 = -918
    _ = len(v3)
    for fi52 := 0; fi52 < 8; fi52++ {
        v53 := GenS5{F0: 191, F1: -2147483648}
        var v54 uint32 = 2
        _ = fi52
        _ = v53
        _ = v54
    }
    var v55 int8 = int8(v51)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v55
}
