package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64 }
func genFn0() uint64 {
    var v6 uint16 = 472
    v7 := GenS1{F0: ("hello" + ""), F1: 128}
    var v8 int32 = -447
    var v9 int = 9223372036854775807
    var v10 float32 = 0.1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 int8, p1 uint64, p2 int16) string {
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := []int{fi11, -9223372036854775808, (fi11 + -573)}
        _ = fi11
        _ = v12
    }
    if (3.14159 == 1e10) {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint16 = 671
            v15 := []int{fi13}
            if ((!false) || false) {
                var v16 uint8 = 220
                fmt.Println(p2)
                _ = v16
            } else {
                var v17 int8 = (p0 ^ p0)
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        if true {
            if (1 == 42) {
                if (true || true) {
                    var v18 int64 = 1
                    for fi19 := 0; fi19 < 7; fi19++ {
                        var v20 float64 = 3.14159
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                } else {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := []int{(fi21 * fi21)}
                        if true {
                            if ((!(4294967295 <= (fi21 & 128))) && ((!false) || (!(!true)))) {
                                v23 := GenS1{F0: ((((("hello" + "a") + "hello") + "foo bar") + "foo bar") + "\n\t"), F1: 255}
                                var v24 int = (-(fi21))
                                _ = v23
                                _ = v24
                            }
                        }
                        _ = len(v22)
                        _ = fi21
                        _ = v22
                    }
                    p1 = p1
                }
            } else {
                v25 := []int{664, 42}
                var v26 int32 = 1
                _ = v25
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return ("\n\t" + "hello")
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]string)
    var v3 int = 42
    var v4 int = (-(v3))
    var v5 float64 = 3.14159
    v27 := GenS1{F0: "a", F1: 188}
    v4 = (v3 >> ((0) & 63))
    v28 := GenS1{F0: "foo bar", F1: 102}
    v29 := GenS1{F0: ("a" + "a"), F1: 34}
    fmt.Println(v3, v1, v1)
    v30 := GenS1{F0: ("foo bar" + ((("" + ("a" + ("foo bar" + ("hello" + "hello")))) + "\n\t") + ("\n\t" + ("hello" + "\n\t")))), F1: 255}
    var v31 int8 = 100
    switch (-(v4)) {
    case 536:
        fmt.Println(v31, v1)
    case 4294967295:
        switch ("" + "a") {
        case "":
            v32 := GenS2{F0: 695}
            _ = v32
        case "foo bar":
            v5 = -2.25
        }
    default:
        fmt.Println(v4, v1, v31)
    }
    v3 += (v3 - 9223372036854775807)
    if true {
        v3 += v4
        if (805 < 2) {
            v33 := GenS1{F0: "foo bar", F1: 252}
            v31 += 127
            _ = len(v2)
            _ = v33
        } else {
            switch ("\n\t" + "hello") {
            case "a":
                v34 := []int{v3, v4}
                _ = v34
            default:
                v2[("foo bar" + "")] = ("a" + "a")
            }
            v3 += v3
        }
    }
    v5 = v5
    v35 := GenS1{F0: ("a" + "hello"), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
