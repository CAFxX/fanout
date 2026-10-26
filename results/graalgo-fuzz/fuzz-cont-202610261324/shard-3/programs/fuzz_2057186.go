package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 string, p1 uint64) int16 {
    p1 = (p1 % ((18446744073709551615) | 1))
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            p1 += p1
            v8 := GenS1{F0: fi7, F1: 2, F2: "hello"}
            v9 := GenS1{F0: (fi7 | (fi7 / ((fi7) | 1))), F1: p1, F2: ((p0 + (p0 + "foo bar")) + "a")}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v10 := GenS1{F0: 128, F1: 18446744073709551615, F2: (((("\n\t" + "\n\t") + p0) + "\n\t") + p0)}
        for fi11 := 0; fi11 < 11; fi11++ {
            p1 = (p1 ^ 42)
            p1 -= (p1 >> ((65535) & 63))
            var v12 int = -242
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    _ = p0
    _ = p1
    return -729
}
func genFn1(p0 int, p1 int64, p2 int8) uint64 {
    var v13 uint = 327
    if true {
        v14 := GenS1{F0: -979, F1: 42, F2: ("hello" + ("hello" + "a"))}
        if (false && false) {
            v15 := []int{1, 2147483647, p0, p0}
            fmt.Println(p2)
            var v16 string = "hello"
            _ = v15
            _ = v16
        } else {
            v17 := []int{(p0 % ((32767) | 1)), (p0 >> (((v13 * (v13 & v13))) & 63))}
            v13 += 42
            _ = v17
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    var v1 int = 32768
    var v2 int8 = 71
    var v3 float64 = 0.0
    var v4 uint8 = 42
    v5 := make(map[string]uint32)
    var v6 uint = uint(v4)
    if true {
        switch (v1 - 42) {
        case 32767:
            v18 := []int{(v1 >> (((v6 << (((v6 & v6)) & 63))) & 63))}
            _ = v18
        }
        v19 := GenS2{F0: true, F1: genFn0("hello", 4294967295), F2: v6}
        var v20 uint64 = 42
        _ = v19
        _ = v20
    } else {
        v5[("foo bar" + "a")] = 2147483647
        var v21 float64 = (1e10 + (v3 - 0.1))
        _ = v21
    }
    v22 := GenS2{F0: false, F1: 1, F2: v6}
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint16 = 115
        _ = fi23
        _ = v24
    }
    switch ("\n\t" + "\n\t") {
    case "":
        v6 = (v6 ^ 4294967295)
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        v5["a"] = 42
    }
    v1 = (-(v1))
    v25 := GenS1{F0: (v1 | v1), F1: 745, F2: "hello"}
    v2 = v2
    var v26 float32 = (3.14159 / 1e10)
    switch "a" {
    case "foo bar":
        v27 := []int{(v1 - (v1 | (v1 * v1))), int(v4), v1, 127}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
