package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := []int{42, 32768, -594}
    var v5 bool = (!false)
    v1[2147483648] = -35
    switch ((("\n\t" + "a") + ("\n\t" + "\n\t")) + "a") {
    case "":
        var v6 uint = (v2 | v2)
        _ = v6
    case "\n\t":
        v5 = false
    case "foo bar":
        _ = len(v4)
    }
    if v5 {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint = v2
            _ = fi7
            _ = v8
        }
    } else {
        var v9 int32 = -327
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        switch (-(fi10)) {
        case -2147483648:
            v2 = (v2 | 0)
        case 832:
            v11 := GenS1{F0: 77, F1: -1, F2: ""}
            _ = v11
        case 32767:
            fmt.Println(v2, fi10, fi10)
        }
        v12 := GenS1{F0: 4294967295, F1: 42, F2: (("foo bar" + ("hello" + "a")) + "a")}
        _ = fi10
        _ = v12
    }
    switch ("foo bar" + "") {
    case "a":
        switch 1 {
        case -2147483648:
            v2 += (v2 % ((1) | 1))
        case 32767:
            v2 -= 775
        case 0:
            switch 0 {
            case -1:
                v2 += v2
            }
        }
    default:
        switch "foo bar" {
        case "\n\t":
            var v13 string = ("a" + (("\n\t" + "a") + "hello"))
            _ = v13
        case "a":
            delete(v3, 2)
        case "":
            var v14 float32 = 0.0
            _ = v14
        default:
            var v15 uint8 = 58
            _ = v15
        }
    }
    _ = len(v4)
    v5 = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
