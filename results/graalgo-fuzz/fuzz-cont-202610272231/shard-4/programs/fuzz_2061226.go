package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64, p1 int64, p2 uint64) int {
    var v6 bool = true
    if true {
        p0 += (p2 ^ (p2 >> ((255) & 63)))
        if (!(9223372036854775807 <= -287)) {
            var v7 int = 32767
            p0 -= (p2 - 2)
            p2 = (p0 + (p0 << ((599) & 63)))
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: (272 <= 268)}
        fmt.Println(fi8)
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(fi8, fi10)
            if (true && ((!false) && true)) {
                var v11 int32 = 1
                var v12 uint64 = 0
                if ((false || true) || false) {
                    v12 = (v12 / ((v12) | 1))
                    var v13 bool = ((-1 <= 127) && true)
                    v14 := GenS1{F0: (!(2147483648 <= 42))}
                    if (1 == 127) {
                        v11 -= v11
                    } else {
                        var v15 uint32 = 1
                        var v16 int16 = 0
                        fmt.Println(v12, v15, v13)
                        _ = v15
                        _ = v16
                    }
                    v13 = v13
                    _ = v13
                    _ = v14
                } else {
                    switch "" {
                    case "foo bar":
                        var v17 int16 = 727
                        _ = v17
                    case "a":
                        v11 -= 42
                    default:
                        var v18 uint64 = v12
                        _ = v18
                    }
                    var v19 float64 = (0.1 - 1e10)
                    _ = v19
                }
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    if (1 >= -2147483648) {
        v20 := GenS1{F0: false}
        _ = v20
    }
    v21 := GenS1{F0: (false && true)}
    _ = v21
    return 255
}

func main() {
    var v1 uint8 = 1
    v2 := []string{"hello", "foo bar", (("" + "foo bar") + (("hello" + ("\n\t" + ("\n\t" + "\n\t"))) + "a")), (("\n\t" + (("\n\t" + "") + "")) + (("hello" + "hello") + "foo bar"))}
    var v3 uint = 91
    var v4 uint16 = 65535
    v5 := make(map[string]int)
    var v22 uint32 = 2147483648
    v3 = (v3 | uint(v1))
    var v23 float32 = 0.0
    switch "a" {
    case "foo bar":
        var v24 uint64 = 2
        _ = v24
    default:
        var v25 uint = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        switch 32767 {
        case -2147483648:
            var v27 uint = v3
            _ = v27
        case -1:
            if (2147483648 < (v22 ^ uint32(v1))) {
                if false {
                    var v28 int = 32768
                    v29 := GenS1{F0: (true && false)}
                    var v30 uint16 = (^(v4))
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
        case 4294967295:
            for fi31 := 0; fi31 < 2; fi31++ {
                fmt.Println(fi26, v22, v23)
                _ = fi31
            }
        }
        _ = fi26
    }
    var v32 uint16 = v4
    var v33 float32 = 0.1
    var v34 uint64 = 4294967295
    v35 := GenS1{F0: (((2147483647 != 9223372036854775807) || true) && true)}
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 string = ("a" + ("\n\t" + ((("hello" + ((("a" + "a") + "\n\t") + ("hello" + ""))) + "\n\t") + "foo bar")))
        _ = len(v2)
        _ = fi36
        _ = v37
    }
    var v38 uint32 = v22
    switch ((("" + "") + ("foo bar" + ((("foo bar" + "a") + "a") + ("hello" + "\n\t")))) + ("hello" + "hello")) {
    case "\n\t":
        var v39 string = (("foo bar" + ("" + "a")) + ("hello" + "foo bar"))
        _ = v39
    case "foo bar":
        v40 := v5["a"]
        _ = v40
    case "hello":
        var v41 uint8 = (v1 >> (((v3 + (v3 - 42))) & 63))
        _ = v41
    default:
        v34 -= v34
    }
    v42 := v5[("foo bar" + (("hello" + ("foo bar" + (("\n\t" + "foo bar") + ""))) + "hello"))]
    var v43 uint64 = (v34 % (((v34 ^ (^(v34)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
