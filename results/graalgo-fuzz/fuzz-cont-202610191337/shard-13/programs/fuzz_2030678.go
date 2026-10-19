package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
type GenS5 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0(p0 uint, p1 uint64) int {
    switch "\n\t" {
    case "\n\t":
        p1 -= 658
    case "":
        v5 := GenS1{F0: 36}
        _ = v5
    case "a":
        var v6 string = "\n\t"
        _ = v6
    default:
        if (42 < 42) {
            var v7 uint8 = 1
            v7 = 0
            fmt.Println(v7, p1, p1)
            _ = v7
        } else {
            v8 := GenS1{F0: (p0 + p0)}
            p1 += p1
            var v9 uint64 = p1
            _ = v8
            _ = v9
        }
    }
    var v10 uint64 = (p1 | (^(p1)))
    v11 := GenS2{F0: 0, F1: 1000}
    fmt.Println(v10)
    fmt.Println(v10, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 bool = (true && false)
    v2 := make(map[string]string)
    v3 := []string{"\n\t", ("" + "a"), "foo bar"}
    var v4 bool = false
    if false {
        if v1 {
            v1 = v1
            for fi12 := 0; fi12 < 20; fi12++ {
                if (560 < 4294967295) {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 int16 = 42
                        _ = fi13
                        _ = v14
                    }
                    v1 = (298 != 255)
                } else {
                    v4 = (v1 || (-2147483648 < 1))
                }
                v1 = (true || ((-466 < -1) || (!((-1 == 63) || v4))))
                _ = fi12
            }
        }
        v15 := GenS3{F0: 39}
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint8 = 127
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            switch "a" {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v19 bool = (!v1)
            fmt.Println(fi18)
            _ = fi18
            _ = v19
        }
        var v20 int64 = 0
        v21 := GenS4{F0: 2, F1: 70, F2: true}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 float64 = 0.0
        _ = fi22
        _ = v23
    }
    var v24 int8 = 127
    if ((v24 != v24) || ((v24 & v24) <= v24)) {
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS5{F0: 18446744073709551615, F1: 169, F2: -261}
            var v27 uint = 42
            _ = fi25
            _ = v26
            _ = v27
        }
        if v4 {
            switch 9223372036854775807 {
            case 580:
                v28 := GenS4{F0: 486, F1: -32768, F2: v4}
                _ = v28
            case 42:
                for fi29 := 0; fi29 < 17; fi29++ {
                    var v30 float64 = 3.14159
                    _ = fi29
                    _ = v30
                }
            default:
                v24 += v24
            }
            var v31 uint64 = 0
            fmt.Println(v24)
            _ = v31
        } else {
            var v32 uint = 2
            var v33 uint32 = 404
            fmt.Println(v32, v4, v33)
            _ = v32
            _ = v33
        }
        v3 = append(v3, "")
    } else {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    _ = len(v3)
    if (("\n\t" + ("\n\t" + "a")) == "a") {
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS3{F0: v24}
            var v36 float32 = 0.1
            v36 = 0.1
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    var v37 int64 = -9223372036854775808
    var v38 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
