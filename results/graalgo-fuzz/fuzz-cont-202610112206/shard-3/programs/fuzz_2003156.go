package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0(p0 uint32) uint16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            p0 = p0
            p0 = 1
            _ = fi8
        }
        _ = fi7
    }
    v9 := GenS1{F0: 42, F1: 0}
    _ = p0
    _ = v9
    return 2
}

func main() {
    var v1 int64 = -32
    var v2 uint8 = 2
    v3 := []int{1}
    v4 := []string{"foo bar"}
    v5 := make(map[string]string)
    var v6 int = 127
    var v10 uint8 = v2
    if true {
        v1 = (v1 - (-(v1)))
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                if ("a" <= "a") {
                    switch "" {
                    case "\n\t":
                        fmt.Println(v1, v1)
                    case "":
                        var v12 int64 = v1
                        _ = v12
                    case "a":
                        for fi13 := 0; fi13 < 5; fi13++ {
                            var v14 float32 = 0.0
                            v2 -= 2
                            _ = fi13
                            _ = v14
                        }
                    }
                }
                var v15 int64 = (v1 >> ((461) & 63))
                _ = v15
            } else {
                v16 := GenS2{F0: -1, F1: 598, F2: 18446744073709551615}
                for fi17 := 0; fi17 < 10; fi17++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) > 0 {
                        v3[0] = (-(fi11))
                    }
                    v1 += v1
                    if (1 > 817) {
                        v18 := GenS2{F0: -2147483648, F1: 325, F2: 42}
                        _ = len(v5)
                        v10 -= v2
                        _ = v18
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (v6 * fi11)
                        }
                        v19 := GenS1{F0: int8(v6), F1: 42}
                        v5[""] = ""
                        _ = v19
                    }
                    var v20 int8 = 0
                    _ = fi17
                    _ = v20
                }
                _ = v16
            }
            var v21 bool = true
            var v22 bool = v21
            _ = fi11
            _ = v21
            _ = v22
        }
    } else {
        v23 := v5[("hello" + ("" + "a"))]
        v24 := GenS2{F0: -2147483648, F1: 2147483648, F2: 1}
        fmt.Println(v1)
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 int32 = -462
        var v27 string = (("hello" + "\n\t") + (("foo bar" + "hello") + "a"))
        _ = fi25
        _ = v26
        _ = v27
    }
    if true {
        var v28 string = ("hello" + "foo bar")
        _ = len(v5)
        _ = v28
    }
    v1 = (v1 << ((255) & 63))
    if ((18446744073709551615 <= 65535) && true) {
        v1 = (-(v1))
    }
    if ((false || true) || (32767 >= 42)) {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int = fi29
            _ = fi29
            _ = v30
        }
    } else {
        var v31 uint8 = (v2 % ((v10) | 1))
        _ = v31
    }
    if (!false) {
        var v32 int64 = 42
        _ = v32
    }
    var v33 string = "hello"
    var v34 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
