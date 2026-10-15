package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 int8) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{fi6, fi6}
        if true {
            p0 = (p0 * 42)
            v8 := GenS1{F0: ("" + "\n\t"), F1: false}
            var v9 float64 = (1.5 - 1e10)
            _ = v8
            _ = v9
        } else {
            switch 743 {
            case 32768:
                var v10 int = 9223372036854775807
                _ = v10
            case 0:
                switch "\n\t" {
                case "":
                    var v11 int16 = 658
                    _ = v11
                case "\n\t":
                    v12 := GenS1{F0: ("hello" + "\n\t"), F1: true}
                    _ = v12
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        v13 := GenS1{F0: "foo bar", F1: true}
                        _ = v13
                    }
                }
            default:
                var v14 int8 = (p0 - (p0 ^ (p0 >> ((255) & 63))))
                _ = v14
            }
            v15 := GenS1{F0: ("foo bar" + (("\n\t" + ("a" + "\n\t")) + "a")), F1: false}
            _ = v15
        }
        p0 = p0
        _ = fi6
        _ = v7
    }
    var v16 float64 = 3.14159
    var v17 uint8 = 127
    var v18 int64 = 42
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint64 = 42
        v21 := GenS2{F0: 42, F1: fi19, F2: (-32768 <= -32768)}
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v18)
            switch ("hello" + "hello") {
            case "foo bar":
                if (!(true && ((true || (3.14159 >= ((float64((1.5 - ((1.5 - (1e10 + ((0.0 / 1.5) + 3.14159))) - -2.25))) - ((v16 / ((v16 / float64(1.5)) - float64(1.5))) * v16)) / (float64(1e10) - float64(0.1))))) && (!true)))) {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v20)
                    case "\n\t":
                        v16 = 0.0
                    default:
                        var v23 uint8 = v17
                        _ = v23
                    }
                    v24 := GenS3{F0: 9223372036854775807}
                    _ = v24
                } else {
                    if (false && false) {
                        v25 := GenS1{F0: "foo bar", F1: ((float64((-2.25 - 0.1)) - (((v16 / 3.14159) - float64(float32(v16))) - (v16 / 3.14159))) != (v16 - 0.0))}
                        _ = v25
                    }
                }
            default:
                var v26 float32 = 0.0
                _ = v26
            }
            v27 := GenS4{F0: 65535}
            _ = fi22
            _ = v27
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return 32767
}

func main() {
    var v1 int16 = 32767
    var v2 uint8 = 255
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 int32 = 1
    v28 := GenS3{F0: 1}
    v29 := []int{-2147483648, 32768}
    v30 := v4[785]
    v29 = append(v29, -811)
    var v31 int16 = -32768
    if len(v29) > 0 {
        v29[0] = 127
    }
    fmt.Println(v1, v30, v31)
    v32 := GenS4{F0: 255}
    v33 := GenS1{F0: "", F1: (v30 && false)}
    var v34 string = ("" + "a")
    v29 = append(v29, 127)
    v35 := GenS2{F0: 2, F1: -2147483648, F2: (v31 <= v1)}
    fmt.Println(v31, v31, v31)
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 4; fi37++ {
            switch genFn0(100) {
            case -9223372036854775808:
                fmt.Println(v31, fi36)
            default:
                fmt.Println(v34, fi36)
            }
            if (42 < -332) {
                v3[1] = ((v34 + v34) >= v34)
                v38 := GenS2{F0: 364, F1: -2147483648, F2: false}
                if v30 {
                    var v39 float32 = 0.0
                    fmt.Println(v1, v5, v34)
                    v40 := GenS2{F0: 2, F1: (fi36 >> ((0) & 63)), F2: ((v34 + v34) >= "")}
                    var v41 uint = 489
                    _ = v39
                    _ = v40
                    _ = v41
                }
                _ = len(v29)
                v1 += v31
                fmt.Println(fi37, v2)
                _ = len(v29)
                _ = v38
            }
            v34 = "hello"
            var v42 int16 = (v31 | 191)
            v43 := GenS5{F0: 65535, F1: (v31 / (((v42 << ((438) & 63))) | 1)), F2: v30}
            var v44 uint64 = 357
            if (v2 >= uint8(v44)) {
                switch v34 {
                case "":
                    var v45 uint32 = 81
                    _ = v45
                }
                v1 = (v31 | (v31 / ((v42) | 1)))
            }
            fmt.Println(v42, v44)
            _ = fi37
            _ = v42
            _ = v43
            _ = v44
        }
        _ = fi36
    }
    fmt.Println(v30, v34)
    fmt.Println(v34, v34)
    v46 := GenS3{F0: 87}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v46
    _ = v5
}
