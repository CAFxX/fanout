package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "foo bar"
    var v3 uint16 = 255
    v1[224] = 954
    v4 := []int{42, 32767, 42, 1}
    _ = len(v1)
    var v5 int = 42
    switch (^(v5)) {
    case -765:
        var v6 uint32 = 2
        _ = v6
    }
    if false {
        switch (v5 / (((^(v5))) | 1)) {
        case -601:
            if (3.14159 > 0.1) {
                var v7 uint16 = (^(v3))
                v5 = 315
                if ((121 > 108) && true) {
                    v7 += (v3 % (((v3 + (v3 + 65535))) | 1))
                    _ = len(v4)
                }
                _ = v7
            }
        case 962:
            v5 = v5
        }
    } else {
        var v8 uint16 = v3
        var v9 uint64 = 498
        v2 = ("hello" + (v2 + v2))
        _ = v8
        _ = v9
    }
    _ = len(v1)
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            for fi11 := 0; fi11 < 10; fi11++ {
                if false {
                    v2 = ""
                    v12 := GenS1{F0: 798, F1: 1, F2: fi10}
                    _ = v12
                } else {
                    v5 = (v5 / (((fi10 - (fi10 + (-(fi10))))) | 1))
                    v5 = 859
                }
                _ = fi11
            }
            v2 = v2
        }
        switch v2 {
        case "a":
            for fi13 := 0; fi13 < 14; fi13++ {
                if true {
                    fmt.Println(fi13)
                } else {
                    v5 = v5
                    v5 = fi13
                    v14 := GenS2{F0: v2}
                    _ = v14
                }
                _ = fi13
            }
        case "hello":
            v4 = append(v4, -2147483648)
        case "":
            if len(v4) > 0 {
                v4[0] = fi10
            }
        }
        v5 += (fi10 | (^(v5)))
        _ = fi10
    }
    switch (("\n\t" + "\n\t") + "\n\t") {
    case "\n\t":
        fmt.Println(v3, v5, v2)
    default:
        var v15 int = 127
        _ = v15
    }
    var v16 int64 = 933
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
