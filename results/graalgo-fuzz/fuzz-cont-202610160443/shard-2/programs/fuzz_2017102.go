package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 bool; F1 int32; F2 int }
func genFn0(p0 int32, p1 uint32, p2 int) bool {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 string = "a"
        fmt.Println(p1, v7)
        _ = fi6
        _ = v7
    }
    v8 := []int{2147483647, (p2 + p2)}
    p2 += 128
    var v9 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (((false && (((1.5 * float64(v9)) == ((float64(v9) * 1e10) * 1.5)) || false)) || (true && false)) || false)
}

func main() {
    var v1 string = "hello"
    v2 := []string{v1, v1}
    var v3 bool = (true || false)
    var v4 int32 = -1
    var v5 int32 = 0
    v10 := GenS1{F0: v1}
    for fi11 := 0; fi11 < 8; fi11++ {
        v5 -= v5
        _ = fi11
    }
    if len(v2) > 0 {
        v2[0] = (v1 + v1)
    }
    var v12 int64 = 1
    v13 := GenS1{F0: (v1 + ("hello" + "hello"))}
    if (v12 > (v12 - v12)) {
        if false {
            if (65535 < 4294967295) {
                v5 = (v4 % ((-526) | 1))
                var v14 uint64 = 42
                v15 := GenS2{F0: true}
                _ = v14
                _ = v15
            }
            v2 = append(v2, (v1 + "\n\t"))
        }
        v3 = v3
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi17
            }
            v12 = (v12 * 170)
            v4 = v5
            _ = fi16
        }
        _ = len(v2)
    }
    var v18 uint64 = 687
    fmt.Println(v5)
    switch 983 {
    case 0:
        var v19 int32 = -528
        _ = v19
    case -585:
        for fi20 := 0; fi20 < 17; fi20++ {
            switch 631 {
            case 9223372036854775807:
                if ((v5 >> ((4294967295) & 63)) > v4) {
                    var v21 string = v1
                    for fi22 := 0; fi22 < 10; fi22++ {
                        switch (fi20 / ((42) | 1)) {
                        case 32768:
                            switch "hello" {
                            case "":
                                var v23 int64 = 237
                                _ = v23
                            }
                        case 86:
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v3, v21, v21)
                            }
                        case 1:
                            fmt.Println(fi20, v21)
                        default:
                            v1 = v21
                        }
                        _ = fi22
                    }
                    _ = v21
                } else {
                    var v24 uint32 = 4294967295
                    var v25 float32 = (3.14159 / -2.25)
                    _ = v24
                    _ = v25
                }
            case 128:
                v26 := GenS2{F0: v3}
                _ = v26
            case -485:
                switch 447 {
                case -1:
                    v2 = append(v2, v1)
                case 127:
                    if len(v2) > 0 {
                        v2[0] = (v1 + (((v1 + v1) + "hello") + ("hello" + "\n\t")))
                    }
                default:
                    switch v1 {
                    case "hello":
                        if (65535 <= 2) {
                            v27 := GenS3{F0: (v12 | v12), F1: v12}
                            if true {
                                v4 -= (v5 & 42)
                            } else {
                                v28 := GenS3{F0: (v12 << ((2) & 63)), F1: (v12 + v12)}
                                _ = v28
                            }
                            _ = v27
                        } else {
                            for fi29 := 0; fi29 < 4; fi29++ {
                                if v3 {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                }
                                _ = fi29
                            }
                        }
                    case "a":
                        var v30 uint8 = 57
                        _ = v30
                    }
                }
            }
            v5 = v5
            _ = fi20
        }
    default:
        v1 = v1
    }
    v5 = (v4 << ((964) & 63))
    fmt.Println(v12)
    v12 = v12
    v31 := GenS4{F0: 2, F1: true}
    v32 := GenS2{F0: v3}
    switch "foo bar" {
    case "a":
        if v3 {
            var v33 int16 = 1
            _ = v33
        }
    case "":
        fmt.Println(v18, v3)
    case "\n\t":
        fmt.Println(v12)
    }
    switch v1 {
    case "a":
        if (!v3) {
            if true {
                v34 := GenS5{F0: (!v3), F1: v4, F2: 937}
                _ = v34
            }
            var v35 int16 = int16(v12)
            v36 := GenS3{F0: (v12 - v12), F1: v12}
            _ = v35
            _ = v36
        }
    }
    v1 = "hello"
    for fi37 := 0; fi37 < 14; fi37++ {
        fmt.Println(v3)
        switch (fi37 >> ((319) & 63)) {
        case 228:
            fmt.Println(v5)
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
