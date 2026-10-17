package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 uint) int16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{fi7, 1, (fi7 ^ fi7)}
        _ = fi7
        _ = v8
    }
    p0 = (p0 >> ((829) & 63))
    _ = p0
    return -1
}
func genFn1(p0 uint8) uint {
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 bool = false
        v12 := []int{}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        v13 := GenS1{F0: ((("\n\t" + (("" + ("a" + "\n\t")) + (("foo bar" + "\n\t") + ("a" + "\n\t")))) + ("a" + ((("hello" + ((("foo bar" + "") + "a") + (("a" + "") + "a"))) + "\n\t") + ("\n\t" + "foo bar")))) + "\n\t"), F1: -112}
        _ = v13
    case "a":
        v9 -= (v9 + v9)
    default:
        v14 := []int{-459, -674, 127, 6}
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: "foo bar", F1: 1}
        var v17 int8 = -54
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v9
    return 611
}

func main() {
    var v1 uint8 = 127
    var v2 uint = 0
    v3 := []uint32{4294967295, 1}
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    var v6 string = ("foo bar" + "")
    switch ("" + v6) {
    case "\n\t":
        v2 = v2
    case "a":
        fmt.Println(v6, v6)
    }
    v18 := GenS1{F0: "foo bar", F1: 127}
    if (!false) {
        v6 = v6
    } else {
        v3 = append(v3, 0)
        for fi19 := 0; fi19 < 7; fi19++ {
            switch fi19 {
            case 42:
                switch (v6 + v6) {
                case "hello":
                    v6 = "a"
                case "a":
                    var v20 int64 = 9223372036854775807
                    _ = v20
                }
            default:
                for fi21 := 0; fi21 < 3; fi21++ {
                    switch 789 {
                    case 864:
                        var v22 uint32 = 2147483648
                        _ = v22
                    default:
                        v23 := GenS2{F0: -469}
                        _ = v23
                    }
                    var v24 uint8 = 42
                    _ = fi21
                    _ = v24
                }
            }
            switch (fi19 - 724) {
            case 32767:
                for fi25 := 0; fi25 < 11; fi25++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi25
                }
            }
            _ = fi19
        }
        var v26 int64 = 9223372036854775807
        _ = v26
    }
    var v27 uint16 = 32767
    if true {
        v3 = append(v3, 4294967295)
        var v28 uint64 = 4294967295
        var v29 bool = (true || ((v2 / ((v2) | 1)) != v2))
        _ = v28
        _ = v29
    } else {
        if false {
            v30 := GenS3{F0: (v1 >> ((v2) & 63)), F1: ((v6 + "\n\t") + "a")}
            v3 = append(v3, 2147483647)
            _ = len(v5)
            _ = v30
        } else {
            _ = len(v3)
            var v31 uint = v2
            v32 := GenS1{F0: "hello", F1: 84}
            _ = v31
            _ = v32
        }
        v5[0] = 0
        switch 42 {
        case 1:
            v27 += v27
        case 42:
            fmt.Println(v6)
        case -2147483648:
            v2 = 525
        default:
            v33 := GenS1{F0: "a", F1: -1}
            _ = v33
        }
    }
    v2 = (v2 ^ (v2 * genFn1((v1 << (((v2 - (^(v2)))) & 63)))))
    delete(v5, -396)
    v34 := GenS4{F0: 0, F1: 32768}
    v6 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
