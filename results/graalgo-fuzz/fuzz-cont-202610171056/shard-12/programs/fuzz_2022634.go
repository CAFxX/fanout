package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
func genFn0(p0 int64) bool {
    v6 := GenS1{F0: ("" == ""), F1: -32768, F2: "hello"}
    var v7 int64 = p0
    v8 := GenS1{F0: true, F1: 1000, F2: ("foo bar" + "")}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (!(-534 <= -194))
}

func main() {
    var v1 uint8 = 2
    var v2 string = "hello"
    var v3 uint = 340
    var v4 int8 = 1
    var v5 uint64 = uint64(v3)
    switch 0 {
    case 127:
        for fi9 := 0; fi9 < 14; fi9++ {
            v3 -= v3
            _ = fi9
        }
    case -9223372036854775808:
        var v10 uint8 = (v1 >> ((v3) & 63))
        _ = v10
    case 2147483647:
        v2 = "\n\t"
    default:
        v11 := GenS1{F0: (v5 >= (v5 << ((613) & 63))), F1: 42, F2: v2}
        _ = v11
    }
    var v12 uint16 = 0
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: false, F1: 42, F2: ""}
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int = (fi13 ^ (fi15 >> ((uint(v12)) & 63)))
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    } else {
        var v17 int = -767
        _ = v17
    }
    var v18 int64 = 1
    var v19 uint32 = 523
    var v20 int = -469
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint32 = v19
        _ = fi21
        _ = v22
    }
    if ((v3 ^ 255) >= v3) {
        switch v20 {
        case -223:
            fmt.Println(v19)
        }
    } else {
        v3 = (v3 & v3)
        switch v2 {
        case "hello":
            var v23 int = (v20 >> (((v3 % ((v3) | 1))) & 63))
            _ = v23
        default:
            fmt.Println(v2)
        }
        v24 := GenS1{F0: true, F1: -1, F2: "foo bar"}
        _ = v24
    }
    switch -2147483648 {
    case 128:
        var v25 uint = v3
        _ = v25
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
