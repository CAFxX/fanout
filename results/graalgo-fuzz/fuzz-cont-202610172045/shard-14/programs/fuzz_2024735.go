package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint8 {
    switch (("a" + (((("hello" + "hello") + "foo bar") + ("\n\t" + "hello")) + "")) + "\n\t") {
    case "hello":
        var v7 int16 = 1
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = -1
            switch ("\n\t" + "\n\t") {
            case "a":
                var v10 int = (^(fi8))
                _ = v10
            }
            _ = fi8
            _ = v9
        }
    case "a":
        var v11 bool = true
        _ = v11
    default:
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
    }
    switch "foo bar" {
    case "hello":
        var v14 int8 = -1
        _ = v14
    default:
        var v15 uint64 = 4294967295
        _ = v15
    }
    return 128
}
func genFn1() uint64 {
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int16 = -807
        v17 -= 42
        _ = fi16
        _ = v17
    }
    if ((!true) || true) {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := []int{(-(fi18)), -1, -444}
            var v20 uint8 = 127
            fmt.Println(fi18)
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint16 = 42
        _ = v21
    }
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = 0.1
    var v3 int8 = 42
    var v4 string = "foo bar"
    v5 := []string{(("foo bar" + "") + v4), "hello"}
    v6 := make(map[uint32]uint32)
    var v22 string = (v4 + "foo bar")
    var v23 int16 = 42
    var v24 string = v4
    _ = len(v6)
    if ((v23 + v23) < 42) {
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS1{F0: 724}
            v1 -= 2147483648
            var v27 uint32 = v1
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v28 float64 = v2
    v2 = ((-2.25 + (((v28 + (v2 * v2)) - v28) + (1e10 - (v28 * v2)))) - v2)
    if false {
        v24 = (v24 + v22)
    } else {
        v29 := GenS1{F0: 1}
        for fi30 := 0; fi30 < 17; fi30++ {
            v4 = "a"
            v23 = (v23 % ((v23) | 1))
            _ = fi30
        }
        v31 := GenS1{F0: 0}
        _ = v29
        _ = v31
    }
    var v32 string = v22
    fmt.Println(v22, v1)
    for fi33 := 0; fi33 < 2; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            fmt.Println(v1)
            v32 = ((v4 + ("\n\t" + v32)) + "a")
            _ = fi34
        }
        _ = fi33
    }
    v35 := v6[65535]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
