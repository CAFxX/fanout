package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() uint32 {
    var v5 string = ""
    v6 := GenS1{F0: (false && true), F1: (!((!((42 >= 76) && (42 != 65535))) || true))}
    var v7 uint32 = 4294967295
    if false {
        v8 := GenS1{F0: (!(192 != 2)), F1: true}
        fmt.Println(v7)
        v9 := GenS1{F0: false, F1: (false || ((v7 % ((65535) | 1)) == (v7 >> ((0) & 63))))}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
        v5 = v5
        var v12 int8 = 127
        _ = v12
    }
    fmt.Println(v7, v7)
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 int32 = -423
    var v2 string = ""
    var v3 uint32 = 42
    var v4 int16 = -1
    var v13 bool = true
    if (623 <= 42) {
        v14 := GenS1{F0: true, F1: ((false || (!v13)) || v13)}
        if (4294967295 < 373) {
            v1 = 2147483647
        } else {
            v13 = v13
        }
        v15 := []int{2147483647, 2147483647}
        _ = v14
        _ = v15
    }
    if v13 {
        v16 := GenS1{F0: (v13 || (511 <= (v4 | v4))), F1: ((0.0 != 1e10) && (!v13))}
        for fi17 := 0; fi17 < 5; fi17++ {
            v3 += (v3 % ((v3) | 1))
            fmt.Println(v13, v3, v2)
            var v18 int32 = (v1 >> ((524) & 63))
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    switch v2 {
    case "foo bar":
        v19 := GenS2{F0: 65535, F1: 0}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS2{F0: 819, F1: 881}
        _ = fi20
        _ = v21
    }
    var v22 uint16 = 786
    fmt.Println(v1, v1, v1)
    var v23 int32 = (v1 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
