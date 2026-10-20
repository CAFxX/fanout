package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 int16 }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
func genFn0(p0 int, p1 uint, p2 int8) int8 {
    var v7 string = ((("a" + "") + "\n\t") + "hello")
    if false {
        v8 := []int{}
        var v9 uint = (p1 << ((211) & 63))
        _ = v8
        _ = v9
    }
    if (42 < -2147483648) {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{(p0 ^ (p0 & (fi10 & (p0 ^ (fi10 & (p0 | (fi10 + fi10))))))), 1}
            _ = fi10
            _ = v11
        }
        switch p0 {
        case 1:
            v12 := GenS1{F0: 0, F1: (p1 ^ (p1 << ((0) & 63))), F2: -821}
            _ = v12
        case -1:
            fmt.Println(v7, p1, p0)
        case -695:
            p0 -= p0
        default:
            var v13 uint16 = 1
            _ = v13
        }
        p0 -= (^(p0))
    } else {
        p2 = (p2 % (((p2 / (((p2 << ((p1) & 63))) | 1))) | 1))
        switch "" {
        case "":
            v7 = v7
        default:
            v14 := []int{}
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v7, fi15, v7)
            _ = fi15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 int32, p1 bool, p2 int16) uint16 {
    var v16 uint8 = 222
    switch ("" + "\n\t") {
    case "a":
        var v17 int64 = 1
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := GenS2{F0: 743, F1: 187, F2: (v16 % ((v16) | 1))}
                if (false && true) {
                    fmt.Println(fi18, fi19)
                } else {
                    var v21 uint16 = 2
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 string = ("" + ("a" + ""))
                _ = fi22
                _ = v23
            }
            _ = fi18
        }
    case "foo bar":
        var v24 int = 127
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint32 = 875
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 42, F1: 383, F2: 935}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v27
    return 499
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"\n\t", ((("hello" + "\n\t") + "foo bar") + (("foo bar" + "foo bar") + ("\n\t" + "hello"))), "a"}
    var v3 int32 = 0
    var v4 uint64 = 2
    var v5 bool = (false || false)
    var v6 float64 = 0.0
    switch "" {
    case "a":
        fmt.Println(v1, v3)
    case "foo bar":
        switch ("foo bar" + "hello") {
        case "a":
            switch 42 {
            case 587:
                for fi28 := 0; fi28 < 1; fi28++ {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        var v30 int32 = v3
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
            case -9223372036854775808:
                if (3.14159 < float32((3.14159 - (((float64(0.1) * v6) * v6) / (v6 / v6))))) {
                    var v31 uint32 = (v1 >> ((2) & 63))
                    _ = v31
                }
            default:
                var v32 uint = 886
                _ = v32
            }
        case "":
            var v33 int16 = -467
            _ = v33
        case "hello":
            if len(v2) > 0 {
                v2[0] = ("" + ((("" + ((("a" + "a") + "foo bar") + ("hello" + "a"))) + "hello") + ("" + "hello")))
            }
        default:
            v34 := GenS1{F0: v4, F1: 982, F2: 1000}
            _ = v34
        }
    case "\n\t":
        switch "" {
        case "foo bar":
            v35 := GenS1{F0: (v4 % (((^(v4))) | 1)), F1: 0, F2: -1}
            _ = v35
        case "hello":
            if ((^(v3)) == (^(v3))) {
                var v36 uint = 0
                v4 -= (v4 - v4)
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int8 = 91
                    _ = fi37
                    _ = v38
                }
                _ = v36
            } else {
                v39 := GenS3{F0: -1, F1: 32767, F2: 32767}
                _ = len(v2)
                _ = v39
            }
        }
    }
    var v40 uint32 = (v1 - (^(v1)))
    for fi41 := 0; fi41 < 19; fi41++ {
        var v42 int32 = v3
        _ = fi41
        _ = v42
    }
    if len(v2) > 0 {
        v2[0] = ("hello" + "\n\t")
    }
    _ = len(v2)
    if false {
        v43 := GenS2{F0: 178, F1: 4294967295, F2: 117}
        v44 := GenS4{F0: -29, F1: (v40 * (v1 * v1)), F2: 127}
        v2 = append(v2, ("" + "foo bar"))
        _ = v43
        _ = v44
    }
    var v45 uint8 = 0
    fmt.Println(v45, v1, v3)
    fmt.Println(v40, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
}
