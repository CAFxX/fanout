package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }

func main() {
    var v1 string = ("" + "\n\t")
    var v2 int16 = 1000
    v3 := []int{127, 42, 32767}
    var v4 uint8 = 116
    v1 = ("\n\t" + "a")
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int32 = 1
        _ = fi5
        _ = v6
    }
    if true {
        v2 += (v2 ^ (v2 + 32767))
    } else {
        var v7 uint16 = 1
        _ = v7
    }
    v8 := GenS1{F0: 4294967295, F1: ((!true) || false), F2: uint16(v4)}
    v2 += -1
    v9 := GenS1{F0: 460, F1: true, F2: 73}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v1 = "\n\t"
        for fi10 := 0; fi10 < 18; fi10++ {
            v4 = 128
            var v11 int8 = 0
            _ = fi10
            _ = v11
        }
    }
    v1 = (v1 + "a")
    switch (((("hello" + (v1 + "hello")) + ("hello" + v1)) + v1) + v1) {
    case "foo bar":
        switch 0 {
        case 172:
            fmt.Println(v2)
        case 32768:
            var v12 int8 = -1
            _ = v12
        default:
            v4 = v4
        }
    }
    switch 346 {
    case -837:
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                _ = len(v3)
                _ = fi14
            }
            _ = fi13
        }
    case 32767:
        v1 = v1
    case 0:
        v2 += 422
    default:
        var v15 float32 = float32(3.14159)
        _ = v15
    }
    var v16 uint64 = 1
    v1 = ("" + v1)
    var v17 float32 = float32(1e10)
    if true {
        if false {
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS2{F0: 2, F1: (v4 * (v4 % ((212) | 1))), F2: (v1 + v1)}
                _ = fi18
                _ = v19
            }
            switch v1 {
            case "\n\t":
                v2 = v2
            case "foo bar":
                var v20 bool = true
                _ = v20
            }
        } else {
            v21 := GenS2{F0: 920, F1: v4, F2: ""}
            if (25 <= (v4 - (v4 & v4))) {
                if (4294967295 == uint(v16)) {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v3 = append(v3, (fi22 | 2147483647))
                        _ = fi22
                    }
                }
                fmt.Println(v2)
                v4 += (v4 % ((v4) | 1))
            }
            _ = v21
        }
    } else {
        fmt.Println(v1, v17)
        var v23 string = ((v1 + ((v1 + v1) + v1)) + (("" + "a") + v1))
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
