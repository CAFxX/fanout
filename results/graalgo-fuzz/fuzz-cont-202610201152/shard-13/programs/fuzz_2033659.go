package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS2 struct { F0 string; F1 string; F2 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 string }
func genFn0() bool {
    var v5 int8 = -1
    v6 := GenS1{F0: 852, F1: 0, F2: 2}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "" {
        case "foo bar":
            fmt.Println(v5)
        case "a":
            for fi8 := 0; fi8 < 13; fi8++ {
                v5 = (v5 >> ((42) & 63))
                v9 := GenS2{F0: "\n\t", F1: ("\n\t" + "a"), F2: 1}
                _ = fi8
                _ = v9
            }
        default:
            v10 := []int{}
            _ = v10
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 uint32 = 65535
    var v2 uint16 = 42
    var v3 int = 42
    var v4 uint = uint(v2)
    v2 = 2
    v11 := GenS3{F0: uint(v2), F1: ("hello" + ("" + "hello")), F2: ("a" + "hello")}
    if (!(!false)) {
        var v12 int = (-(v3))
        v1 = (v1 << (((v4 % ((v4) | 1))) & 63))
        v13 := []int{v12}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 2147483647, F2: v1}
    v15 := GenS3{F0: v4, F1: "\n\t", F2: "hello"}
    if ("" >= "") {
        v16 := []int{}
        v1 = v1
        var v17 float64 = -2.25
        _ = v16
        _ = v17
    }
    v4 = (^(v4))
    var v18 uint16 = (v2 | 834)
    fmt.Println(v4)
    v18 += 65535
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint64 = uint64(v1)
        if ((!((("" + "") <= "") || (-2.25 >= 1e10))) && true) {
            fmt.Println(v2, v4)
            switch fi19 {
            case 32768:
                v21 := GenS2{F0: "hello", F1: "", F2: 243}
                _ = v21
            case 387:
                v22 := []int{int(v1), 9223372036854775807, v3}
                _ = v22
            }
        } else {
            if (0 >= 552) {
                v23 := []int{(v3 << ((234) & 63))}
                _ = v23
            }
        }
        var v24 uint8 = 255
        _ = fi19
        _ = v20
        _ = v24
    }
    var v25 uint8 = 127
    v26 := GenS2{F0: "hello", F1: "", F2: v25}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
