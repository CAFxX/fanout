package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }

func main() {
    var v1 uint16 = 2
    var v2 uint32 = uint32(v1)
    var v3 uint8 = 2
    var v4 uint32 = (v2 + (v2 ^ 2147483648))
    v5 := GenS1{F0: 42, F1: 0, F2: 734}
    v1 = (v1 / ((v1) | 1))
    v2 += v2
    v6 := GenS1{F0: -1, F1: int32(v3), F2: -2147483648}
    v7 := []int{}
    if (1 < 0) {
        if false {
            _ = len(v7)
            fmt.Println(v1)
        } else {
            v8 := GenS1{F0: -806, F1: 42, F2: 216}
            v9 := GenS1{F0: 0, F1: 312, F2: 155}
            _ = v8
            _ = v9
        }
        _ = len(v7)
    }
    if false {
        _ = len(v7)
        v10 := GenS1{F0: -1, F1: -1, F2: 190}
        _ = v10
    }
    switch ("foo bar" + ("a" + ("hello" + ""))) {
    case "hello":
        var v11 uint32 = (v2 % ((v2) | 1))
        _ = v11
    }
    v12 := GenS1{F0: 299, F1: 0, F2: -1}
    fmt.Println(v2, v4)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v2 -= v2
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int16 = 0
            var v16 int8 = 127
            _ = fi14
            _ = v15
            _ = v16
        }
        if len(v7) > 0 {
            v7[0] = 128
        }
        switch -750 {
        case 32768:
            var v17 int16 = -32768
            _ = v17
        case -2147483648:
            v2 -= 1
        case 1:
            if (380 <= 0) {
                fmt.Println(v1, v3, v1)
            } else {
                v7 = append(v7, fi13)
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 int8 = -119
                    v4 = v2
                    _ = len(v7)
                    _ = fi18
                    _ = v19
                }
            }
        }
        _ = fi13
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int32 = 0
        var v22 int = fi20
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        switch "" {
        case "\n\t":
            switch "" {
            case "\n\t":
                v23 := GenS1{F0: -1, F1: 0, F2: -79}
                _ = v23
            }
        }
        v24 := GenS1{F0: 32767, F1: -260, F2: 98}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
