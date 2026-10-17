package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 bool, p1 string) int16 {
    p0 = (1e10 < 3.14159)
    v4 := GenS1{F0: -128, F1: 128, F2: 538}
    fmt.Println(p1)
    v5 := []int{}
    v6 := GenS1{F0: 100, F1: 101, F2: -577}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 32767
}
func genFn1() int8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = 65535
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 128
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 918
    var v3 int64 = (v2 >> ((18446744073709551615) & 63))
    fmt.Println(v3, v2)
    delete(v1, 603)
    switch "a" {
    case "hello":
        v2 += (v2 - (v3 | v3))
    }
    v2 = v2
    for fi10 := 0; fi10 < 8; fi10++ {
        v2 = v2
        v3 = 843
        _ = fi10
    }
    if false {
        var v11 float64 = -2.25
        fmt.Println(v2)
        _ = v11
    } else {
        if false {
            var v12 int16 = genFn0(false, "foo bar")
            _ = v12
        }
        var v13 float64 = (1.5 - float64((1.5 * 1e10)))
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 97, F1: 2147483647}
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 string = ""
                v3 = (-(v3))
                v13 += float64(0.1)
                _ = fi16
                _ = v17
            }
            var v18 int8 = 100
            _ = fi14
            _ = v15
            _ = v18
        }
        _ = v13
    }
    if true {
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 string = "foo bar"
            var v21 int = fi19
            var v22 uint32 = 2
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        v23 := []int{-1}
        for fi24 := 0; fi24 < 10; fi24++ {
            fmt.Println(fi24, v2)
            var v25 uint64 = 355
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    if (!(!(((1e10 - -2.25) < 0.0) || false))) {
        v3 += v3
        var v26 float64 = 1.5
        fmt.Println(v26, v26)
        _ = v26
    }
    v1[-200] = 221
    v3 = v3
    if true {
        if (false || true) {
            switch (("\n\t" + ((("a" + "foo bar") + "a") + "hello")) + ("hello" + "")) {
            case "a":
                fmt.Println(v2)
            case "":
                v2 = v2
            }
            v2 = -9223372036854775808
        }
    } else {
        v2 = 0
        v3 = (v3 * (v3 / (((v3 % ((9223372036854775807) | 1))) | 1)))
    }
    var v27 int32 = 857
    fmt.Println(v2)
    var v28 uint64 = 42
    fmt.Println(v3, v27)
    var v29 int = -2147483648
    if false {
        var v30 uint8 = 0
        fmt.Println(v30, v29)
        fmt.Println(v2)
        _ = v30
    } else {
        if ((-(v27)) == 1) {
            var v31 int32 = (-(v27))
            var v32 string = ("\n\t" + ("foo bar" + (((("foo bar" + "a") + "hello") + "") + (("" + (("hello" + "") + "a")) + ("foo bar" + (("hello" + "foo bar") + "\n\t"))))))
            _ = v31
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
