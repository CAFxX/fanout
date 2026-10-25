package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 float32 = float32(((-2.25 + 0.1) - 1.5))
    var v4 uint8 = 41
    var v5 uint64 = 0
    var v6 uint8 = 0
    v6 = (v4 & uint8(v5))
    v4 = v4
    v7 := GenS1{F0: v5, F1: "hello", F2: ""}
    if (1 <= 80) {
        v8 := GenS1{F0: 0, F1: ("" + "foo bar"), F2: ""}
        _ = v8
    } else {
        var v9 string = "foo bar"
        _ = v9
    }
    v10 := []int{42, -9223372036854775808, 42}
    if (true && true) {
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 float32 = v3
                _ = fi12
                _ = v13
            }
            var v14 uint32 = (v1 - 2147483647)
            _ = fi11
            _ = v14
        }
    }
    if (150 != 32767) {
        v3 += (1e10 + 1e10)
        v6 = 1
        fmt.Println(v1, v5)
    }
    if true {
        if true {
            if (65535 >= 65535) {
                if (false || false) {
                    _ = len(v10)
                    v15 := GenS1{F0: (v5 + (v5 / (((v5 << ((133) & 63))) | 1))), F1: "\n\t", F2: ("a" + "")}
                    switch 879 {
                    case -1:
                        fmt.Println(v6)
                    default:
                        var v16 uint32 = v1
                        _ = v16
                    }
                    _ = v15
                } else {
                    v17 := GenS1{F0: v5, F1: ("" + (("hello" + "foo bar") + ("hello" + (("foo bar" + "a") + "a")))), F2: ""}
                    v10 = append(v10, -939)
                    var v18 float32 = (1e10 - v3)
                    _ = v17
                    _ = v18
                }
                v1 += (v1 * (^(v1)))
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            delete(v2, v1)
            if (-10 == -75) {
                v2[(v1 & v1)] = false
                for fi19 := 0; fi19 < 11; fi19++ {
                    v1 = (v1 % (((^(v1))) | 1))
                    if len(v10) > 0 {
                        v10[0] = (fi19 | (-(fi19)))
                    }
                    _ = fi19
                }
            }
        }
        v20 := GenS1{F0: (v5 >> ((612) & 63)), F1: ("" + ""), F2: ((("hello" + ("a" + "hello")) + "foo bar") + "\n\t")}
        v4 -= (v6 << ((255) & 63))
        _ = v20
    } else {
        if (((3.14159 / v3) < 0.1) && (18446744073709551615 <= 1)) {
            fmt.Println(v1, v1)
            fmt.Println(v4, v4)
        }
        switch "foo bar" {
        case "":
            v5 += (^(v5))
        case "a":
            var v21 uint16 = 65535
            _ = v21
        case "hello":
            for fi22 := 0; fi22 < 18; fi22++ {
                v6 = 255
                _ = fi22
            }
        }
    }
    v3 += (v3 + -2.25)
    delete(v2, (v1 % ((289) | 1)))
    v23 := v2[v1]
    v24 := GenS2{F0: "foo bar"}
    var v25 float32 = 3.14159
    switch (((("foo bar" + ("\n\t" + "foo bar")) + "\n\t") + "") + "\n\t") {
    case "foo bar":
        v6 = v4
    case "hello":
        if (!(-2147483648 != -9223372036854775808)) {
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int16 = 1
                for fi28 := 0; fi28 < 7; fi28++ {
                    delete(v2, v1)
                    v25 = v25
                    _ = fi28
                }
                v29 := GenS3{F0: 45, F1: (v5 % ((uint64(v1)) | 1))}
                _ = fi26
                _ = v27
                _ = v29
            }
            var v30 int32 = 1
            var v31 uint8 = (^(v4))
            _ = v30
            _ = v31
        }
    case "":
        v6 -= 127
    default:
        v10 = append(v10, -9223372036854775808)
    }
    v10 = append(v10, 804)
    var v32 uint8 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
