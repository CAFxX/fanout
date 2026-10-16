package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]int)
    var v6 int8 = 42
    v4 = v4
    switch -2147483648 {
    case 388:
        var v7 float32 = 3.14159
        _ = v7
    }
    v4 = v4
    v1[("hello" + "\n\t")] = 2147483647
    if true {
        if ((v6 | (v6 % ((100) | 1))) <= (v6 + (v6 + (v6 << ((255) & 63))))) {
            var v8 uint64 = 20
            _ = v8
        }
        v9 := GenS1{F0: ""}
        _ = v9
    }
    fmt.Println(v6, v4)
    switch "" {
    case "\n\t":
        fmt.Println(v6)
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            v6 = v6
            var v11 uint16 = 0
            for fi12 := 0; fi12 < 8; fi12++ {
                v1[""] = (v4 + (^(v4)))
                var v13 uint = 549
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    default:
        v6 = 112
    }
    v4 += 0
    fmt.Println(v6, v4)
    if true {
        v14 := GenS1{F0: ("\n\t" + "a")}
        _ = v14
    } else {
        switch ("foo bar" + "") {
        case "hello":
            v6 -= -4
        case "":
            v15 := []int{}
            _ = v15
        default:
            var v16 int = -1
            _ = v16
        }
        v17 := GenS2{F0: true, F1: 32767}
        _ = v17
    }
    v18 := []int{340, 4294967295}
    switch ("\n\t" + "foo bar") {
    case "":
        if (-168 != 760) {
            v6 = 1
            v6 = -23
            v6 += (v6 / (((-(v6))) | 1))
        }
    case "a":
        for fi19 := 0; fi19 < 2; fi19++ {
            v5[v4] = fi19
            v20 := GenS3{F0: ("hello" + "foo bar")}
            _ = fi19
            _ = v20
        }
    default:
        var v21 uint32 = v4
        _ = v21
    }
    v6 = (v6 >> ((18446744073709551615) & 63))
    var v22 uint8 = 1
    _ = len(v1)
    var v23 int = 32767
    v18 = append(v18, -113)
    var v24 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
