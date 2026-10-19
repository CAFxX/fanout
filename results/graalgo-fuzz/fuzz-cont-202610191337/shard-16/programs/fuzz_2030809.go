package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0() uint32 {
    var v5 bool = true
    switch 128 {
    case -2147483648:
        for fi6 := 0; fi6 < 9; fi6++ {
            v5 = false
            _ = fi6
        }
    case 32768:
        fmt.Println(v5)
    case 32767:
        v7 := GenS1{F0: 1, F1: "\n\t", F2: 0}
        _ = v7
    }
    v5 = v5
    var v8 bool = v5
    _ = v5
    _ = v8
    return 772
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 128
    var v3 int32 = -17
    var v4 int32 = 2147483647
    if (879 != 4294967295) {
        v9 := []int{(v2 - (v2 - v2)), 114, v2}
        for fi10 := 0; fi10 < 19; fi10++ {
            if (("hello" + ("" + "\n\t")) < "\n\t") {
                if len(v9) > 0 {
                    v9[0] = (v2 | (fi10 % ((fi10) | 1)))
                }
            } else {
                v2 -= v2
            }
            _ = fi10
        }
        _ = v9
    } else {
        var v11 int16 = 899
        _ = v11
    }
    v3 += (v3 - v4)
    switch "a" {
    case "\n\t":
        switch "a" {
        case "a":
            var v12 int = -843
            _ = v12
        case "foo bar":
            _ = len(v1)
        default:
            v13 := []int{}
            _ = v13
        }
    case "":
        v14 := []int{v2}
        _ = v14
    }
    v15 := GenS2{F0: "foo bar"}
    var v16 uint = 4294967295
    v4 = (v4 * (^(v3)))
    var v17 int32 = (v3 << ((2) & 63))
    var v18 uint16 = 1
    switch v2 {
    case 9223372036854775807:
        v19 := GenS3{F0: 1}
        _ = v19
    case 32768:
        var v20 string = "a"
        _ = v20
    default:
        v1["\n\t"] = (127 <= 127)
    }
    v21 := []int{32768, (v2 << ((v16) & 63))}
    for fi22 := 0; fi22 < 8; fi22++ {
        v16 = (^(v16))
        v23 := GenS1{F0: 100, F1: ("\n\t" + ("hello" + "")), F2: (v18 + 0)}
        _ = fi22
        _ = v23
    }
    if true {
        var v24 uint16 = 32767
        _ = v24
    } else {
        var v25 uint8 = 128
        if (v25 < 128) {
            var v26 int8 = 42
            v27 := GenS3{F0: 127}
            var v28 string = "\n\t"
            _ = v26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := v1[("a" + "")]
            _ = len(v1)
            v17 -= (-(v3))
            _ = fi29
            _ = v30
        }
        _ = v25
    }
    switch v2 {
    case 0:
        if (!false) {
            var v31 uint32 = genFn0()
            v32 := GenS3{F0: 1}
            _ = v31
            _ = v32
        } else {
            fmt.Println(v17, v18, v4)
        }
    case -630:
        switch "hello" {
        case "hello":
            var v33 uint8 = uint8(v16)
            _ = v33
        default:
            v34 := GenS4{F0: 278, F1: (v18 / (((v18 ^ v18)) | 1))}
            _ = v34
        }
    case 32768:
        fmt.Println(v3, v4)
    }
    switch ("\n\t" + ("" + (((("\n\t" + "foo bar") + "") + "") + "\n\t"))) {
    case "a":
        v35 := GenS3{F0: 255}
        _ = v35
    }
    fmt.Println(v3, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
