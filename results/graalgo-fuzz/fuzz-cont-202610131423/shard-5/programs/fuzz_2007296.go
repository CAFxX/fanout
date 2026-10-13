package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
func genFn0(p0 int, p1 uint16, p2 int32) string {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint64 = 629
        switch (("\n\t" + (("a" + "foo bar") + ("\n\t" + "foo bar"))) + "") {
        case "a":
            var v8 int = (p0 * p0)
            _ = v8
        }
        fmt.Println(p0, p2, fi6)
        _ = fi6
        _ = v7
    }
    v9 := GenS1{F0: -11, F1: 351, F2: "foo bar"}
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: 42, F1: uint64(p1), F2: (("hello" + "foo bar") + (("a" + "\n\t") + "a"))}
        v12 := []int{}
        p1 = 2
        _ = fi10
        _ = v11
        _ = v12
    }
    if ((("foo bar" + "") + "\n\t") <= "hello") {
        var v13 float64 = float64(3.14159)
        _ = v13
    } else {
        if true {
            p2 += (^(p2))
            fmt.Println(p2)
            switch (p0 % ((p0) | 1)) {
            case 9223372036854775807:
                var v14 uint64 = 683
                _ = v14
            case 42:
                for fi15 := 0; fi15 < 2; fi15++ {
                    p0 = (fi15 - (p0 % ((p0) | 1)))
                    _ = fi15
                }
            default:
                var v16 string = ("a" + "a")
                _ = v16
            }
        }
        fmt.Println(p1, p2)
        var v17 uint64 = 1
        _ = v17
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 float64 = 0.0
    var v5 uint = 1
    v4 += (v4 * float64((3.14159 * (1.5 / (-2.25 * 1.5)))))
    fmt.Println(v5, v5)
    v18 := v3["foo bar"]
    v5 -= (v5 % ((v5) | 1))
    v19 := GenS1{F0: 42, F1: 0, F2: "hello"}
    switch ("\n\t" + "foo bar") {
    case "":
        var v20 uint16 = uint16(v5)
        _ = v20
    case "a":
        v5 = (v5 >> (((v5 | 76)) & 63))
    }
    var v21 int32 = 0
    var v22 int = 1
    if false {
        var v23 string = ""
        v24 := v3[((((((("a" + "foo bar") + (v23 + v23)) + v23) + v23) + "foo bar") + (v23 + v23)) + "hello")]
        _ = v23
        _ = v24
    } else {
        var v25 uint64 = 443
        v26 := []int{2147483647}
        if ((!v18) && v18) {
            v26 = append(v26, 0)
            v22 = 9223372036854775807
        }
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
