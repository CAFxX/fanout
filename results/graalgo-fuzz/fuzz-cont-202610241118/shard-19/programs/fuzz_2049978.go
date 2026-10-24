package main

import (
    "fmt"
)

func genFn0(p0 uint16) uint16 {
    v4 := []int{937, -9223372036854775808, 2147483647}
    var v5 string = "a"
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    if false {
        var v6 uint32 = 223
        if false {
            var v7 uint64 = 665
            _ = v7
        } else {
            p0 = (p0 * (p0 & 1))
        }
        var v8 int64 = 0
        _ = v6
        _ = v8
    }
    var v9 float32 = (-2.25 + (float32(1e10) - (0.1 - 1e10)))
    _ = p0
    _ = v4
    _ = v5
    _ = v9
    return p0
}

func main() {
    v1 := []string{(("\n\t" + "foo bar") + "a"), (("a" + "hello") + "a"), ("\n\t" + "foo bar")}
    var v2 string = (("hello" + "") + ((("a" + ("a" + "\n\t")) + (("" + "hello") + "hello")) + ("" + "a")))
    var v3 uint = 42
    _ = len(v1)
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            fmt.Println(fi10, fi10, v3)
        }
        var v11 uint32 = 2147483647
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        switch (fi12 ^ fi12) {
        case 128:
            _ = len(v1)
        case 42:
            var v13 int16 = 717
            _ = v13
        case 32768:
            v3 -= v3
        default:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(v2, v3)
            if ("" == "a") {
                var v15 string = v2
                v1 = append(v1, v15)
                var v16 uint = (^(v3))
                _ = v15
                _ = v16
            }
            _ = fi14
        }
        v3 -= 255
        _ = fi12
    }
    var v17 uint16 = 143
    var v18 int16 = 1000
    switch ((v2 + "") + (("" + ("a" + v2)) + ("a" + ("a" + "\n\t")))) {
    case "\n\t":
        v3 = v3
    case "a":
        v18 -= v18
    }
    var v19 bool = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v19 {
        if ((v19 || v19) && ((^(v17)) < genFn0(v17))) {
            var v20 string = v2
            var v21 int8 = 42
            var v22 uint16 = v17
            _ = v20
            _ = v21
            _ = v22
        }
        switch v2 {
        case "":
            var v23 int8 = 127
            _ = v23
        case "hello":
            v3 = v3
        default:
            v18 -= (v18 + 1)
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
