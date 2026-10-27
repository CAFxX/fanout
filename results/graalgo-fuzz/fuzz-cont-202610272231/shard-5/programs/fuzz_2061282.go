package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int16; F2 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int8 }
type GenS4 struct { F0 uint64 }
func genFn0() int32 {
    switch -1 {
    case 128:
        var v4 int32 = -2147483648
        _ = v4
    }
    var v5 uint16 = 1
    _ = v5
    return -522
}

func main() {
    var v1 int32 = -202
    var v2 int16 = 1
    var v3 string = ("hello" + "\n\t")
    v6 := []int{2147483647}
    var v7 uint32 = 2147483647
    var v8 int32 = (^(v1))
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v1)
        v10 := GenS1{F0: 0}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS2{F0: (v2 & v2), F1: (v2 + (v2 + (v2 ^ (v2 | int16(v7))))), F2: v3}
        v1 += (-(v1))
        v8 -= v1
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        _ = len(v6)
        _ = fi13
    }
    if (false && (-300 >= 84)) {
        switch 2147483647 {
        case 2147483647:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        switch ("\n\t" + v3) {
        case "a":
            switch -9223372036854775808 {
            case 9223372036854775807:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 32767:
                fmt.Println(v8, v2, v2)
            case 526:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v1 = (v8 << ((18446744073709551615) & 63))
                    _ = fi14
                }
            }
        case "foo bar":
            v3 = (("\n\t" + v3) + "hello")
        case "hello":
            switch -901 {
            case -1:
                v15 := GenS3{F0: 293, F1: 754, F2: int8(v7)}
                _ = v15
            }
        }
        _ = len(v6)
    } else {
        v16 := GenS4{F0: 794}
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
        _ = v16
    }
    v6 = append(v6, -205)
    if ((-2.25 * 0.1) <= 0.1) {
        var v17 int32 = (v1 & v8)
        var v18 uint = 65535
        _ = v17
        _ = v18
    }
    v3 = v3
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v19 bool = (false && (genFn0() != v1))
    var v20 float64 = 1.5
    fmt.Println(v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
