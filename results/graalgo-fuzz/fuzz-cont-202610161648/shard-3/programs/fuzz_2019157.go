package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 int16) int64 {
    v4 := GenS1{F0: 1, F1: -1, F2: 848}
    for fi5 := 0; fi5 < 17; fi5++ {
        switch fi5 {
        case -2147483648:
            if true {
                p0 = p0
                v6 := GenS1{F0: 127, F1: 519, F2: 922}
                _ = v6
            }
        }
        _ = fi5
    }
    if true {
        v7 := []int{}
        if (false && true) {
            fmt.Println(p0)
            p0 = p0
        } else {
            var v8 int16 = (p0 - p0)
            _ = v8
        }
        _ = v7
    }
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    var v3 uint = 65535
    var v9 int64 = -1
    var v10 uint16 = 42
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 string = "\n\t"
        _ = fi11
        _ = v12
    }
    var v13 int64 = v9
    v1 += v1
    var v14 uint8 = 127
    switch (("" + (("a" + "a") + "")) + ("hello" + "")) {
    case "hello":
        for fi15 := 0; fi15 < 17; fi15++ {
            v10 = 42
            _ = fi15
        }
    case "\n\t":
        v9 += (v9 + int64(v10))
    }
    var v16 int8 = (-(v1))
    fmt.Println(v3, v13)
    switch 42 {
    case 42:
        if false {
            v17 := GenS1{F0: int8(v10), F1: -1, F2: 2}
            v18 := GenS2{F0: uint(v14), F1: 2}
            _ = v17
            _ = v18
        }
    case -2147483648:
        v2[((("" + "") + "") + "\n\t")] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
