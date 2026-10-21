package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int32, p1 uint8) uint16 {
    v4 := []int{32768}
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(p0, fi5, fi5)
        var v6 uint = 844
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 999
}

func main() {
    var v1 uint32 = 42
    var v2 int16 = 0
    var v3 uint64 = 2
    if ((^(v3)) > 1) {
        v1 = v1
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 0}
            var v9 float64 = 1e10
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v1, v2)
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v2 = v2
        var v11 float32 = 1e10
        if (v11 < float32((0.1 - (1.5 - 1.5)))) {
            var v12 bool = true
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v13 := []int{372}
    if true {
        switch 4294967295 {
        case -9223372036854775808:
            v1 += v1
        case 474:
            v3 = v3
        default:
            switch "hello" {
            case "hello":
                fmt.Println(v1, v2)
            case "foo bar":
                if true {
                    var v14 uint8 = 129
                    fmt.Println(v1, v3, v2)
                    switch ("hello" + ("foo bar" + "hello")) {
                    case "a":
                        var v15 uint8 = 58
                        _ = v15
                    case "\n\t":
                        v16 := GenS2{F0: true}
                        _ = v16
                    default:
                        v14 = (v14 | (v14 | v14))
                    }
                    _ = v14
                } else {
                    if false {
                        var v17 bool = ((v1 >> ((2) & 63)) < 325)
                        v3 = v3
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v3 -= 2
                        fmt.Println(fi18)
                        _ = fi18
                    }
                }
            case "a":
                for fi19 := 0; fi19 < 18; fi19++ {
                    v2 -= v2
                    _ = fi19
                }
            default:
                v20 := GenS3{F0: 462, F1: uint32(v3)}
                _ = v20
            }
        }
    }
    fmt.Println(v2, v3)
    v21 := GenS2{F0: false}
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 string = "hello"
        v2 -= 1000
        var v24 int32 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int16 = v2
    var v26 uint64 = (^(v3))
    v3 += (v3 | (v26 >> ((2) & 63)))
    v26 += v26
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 19; fi28++ {
            if len(v13) > 0 {
                v13[0] = fi27
            }
            _ = fi28
        }
        _ = fi27
    }
    v29 := GenS1{F0: 1}
    _ = len(v13)
    if (!(true || true)) {
        v1 -= v1
        fmt.Println(v26, v3)
    } else {
        if (-81 != int32(v1)) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if false {
                if true {
                    v30 := GenS1{F0: genFn0(275, 0)}
                    _ = v30
                }
                fmt.Println(v25, v25, v25)
                v25 = v25
            }
            var v31 int8 = -25
            _ = v31
        }
    }
    var v32 float64 = -2.25
    var v33 int = 1
    for fi34 := 0; fi34 < 13; fi34++ {
        v26 -= (^(v3))
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v32
    _ = v33
}
