package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int, p1 int32, p2 int) string {
    v4 := GenS1{F0: 530, F1: 1, F2: p1}
    v5 := []int{-2147483648, (p0 | p0), 1}
    var v6 float32 = 0.1
    v7 := GenS2{F0: 148}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return ("a" + (("\n\t" + "\n\t") + "foo bar"))
}

func main() {
    v1 := []string{}
    var v2 uint32 = 1
    var v3 int16 = 1000
    if true {
        v2 += 0
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS3{F0: 1}
                v3 = v3
                v3 -= -49
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                v2 += (v2 / ((2147483647) | 1))
                v3 += v3
                v1 = append(v1, "hello")
                _ = fi11
            }
            _ = fi8
        }
        v12 := GenS3{F0: 18446744073709551615}
        var v13 int = 1
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        switch "" {
        case "":
            for fi15 := 0; fi15 < 10; fi15++ {
                v2 += (v2 ^ (v2 | 2147483648))
                var v16 uint = uint(v2)
                if ((("a" != ("foo bar" + "")) && true) && (42 > 0)) {
                    switch fi15 {
                    case 937:
                        v2 = (v2 - v2)
                    }
                    switch fi14 {
                    case 9223372036854775807:
                        var v17 int8 = 127
                        _ = v17
                    }
                    v18 := GenS2{F0: 42}
                    var v19 uint16 = 1
                    _ = v18
                    _ = v19
                }
                v20 := GenS4{F0: 42}
                if (((v3 * v3) >= (v3 & (v3 + 32767))) && (-2147483648 != 0)) {
                    v21 := GenS2{F0: 2}
                    _ = v21
                }
                _ = len(v1)
                _ = fi15
                _ = v16
                _ = v20
            }
        case "hello":
            var v22 uint8 = 0
            _ = v22
        default:
            var v23 int64 = 0
            _ = v23
        }
        _ = fi14
    }
    var v24 int64 = -246
    for fi25 := 0; fi25 < 7; fi25++ {
        if len(v1) > 0 {
            v1[0] = ("foo bar" + (("\n\t" + "") + ((("foo bar" + "") + "") + ("\n\t" + ""))))
        }
        var v26 string = "hello"
        v26 = (("" + "hello") + ((v26 + (v26 + ("a" + ((v26 + ((("" + "a") + v26) + (("a" + v26) + "\n\t"))) + "a")))) + v26))
        _ = fi25
        _ = v26
    }
    if true {
        var v27 uint16 = 2
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = v27
    }
    v2 -= (v2 ^ 4294967295)
    for fi28 := 0; fi28 < 11; fi28++ {
        v3 = 676
        if false {
            fmt.Println(v24)
        } else {
            v24 = (v24 & -358)
            for fi29 := 0; fi29 < 20; fi29++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi29
            }
            var v30 uint8 = 127
            _ = v30
        }
        _ = fi28
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v32 int32 = 416
        for fi33 := 0; fi33 < 18; fi33++ {
            fmt.Println(v2, v24, fi33)
            fmt.Println(v2, v32)
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
