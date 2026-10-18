package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 string; F2 int32 }
func genFn0(p0 int16, p1 int8) int16 {
    var v7 float32 = (0.1 + 1e10)
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "\n\t" {
        case "a":
            var v9 int8 = (p1 ^ (p1 + p1))
            _ = v9
        }
        fmt.Println(p1, v7, v7)
        _ = fi8
    }
    v10 := []int{}
    _ = len(v10)
    v7 = float32(1e10)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return (p0 << ((4294967295) & 63))
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 2
    var v3 float32 = 0.0
    v4 := make(map[string]bool)
    var v5 bool = true
    var v6 int8 = -1
    fmt.Println(v2, v6)
    if (v5 || (!(float32((0.1 * 1.5)) != v3))) {
        v3 += v3
        var v11 int32 = -1
        v12 := GenS1{F0: v2}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS1{F0: (v2 ^ 65535)}
            v15 := v1[("" + "foo bar")]
            for fi16 := 0; fi16 < 10; fi16++ {
                v6 = (v6 * (v6 & v6))
                _ = fi16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int16 = -32768
        var v19 bool = true
        v5 = false
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = int16(v6)
    v20 = -32768
    switch 776 {
    case 42:
        v21 := GenS2{F0: 65535, F1: ("a" + ("hello" + "hello")), F2: 0}
        _ = v21
    case -736:
        delete(v1, "a")
    case 589:
        switch 2147483647 {
        case -145:
            if v5 {
                var v22 uint32 = 4294967295
                var v23 uint64 = 1
                var v24 int = 9223372036854775807
                _ = v22
                _ = v23
                _ = v24
            }
        }
    default:
        switch 1 {
        case -786:
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 int = fi25
                if (181 < 128) {
                    switch (v26 | v26) {
                    case 2147483647:
                        v2 -= 1
                    }
                }
                _ = fi25
                _ = v26
            }
        }
    }
    v5 = v5
    fmt.Println(v6)
    var v27 uint = (v2 << ((v2) & 63))
    for fi28 := 0; fi28 < 3; fi28++ {
        if (0 > 0) {
            v29 := GenS2{F0: 60, F1: "hello", F2: 2147483647}
            _ = v29
        }
        _ = fi28
    }
    v6 -= (^(v6))
    v3 -= 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
