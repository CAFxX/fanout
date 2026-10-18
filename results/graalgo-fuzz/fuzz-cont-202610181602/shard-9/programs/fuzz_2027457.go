package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 1
    var v2 uint32 = 0
    var v3 int64 = -1
    var v4 uint = 42
    var v5 float64 = 1e10
    var v6 int = 42
    var v7 uint8 = 0
    var v8 int8 = -1
    if false {
        var v9 int32 = 42
        _ = v9
    } else {
        v10 := GenS1{F0: true}
        var v11 uint8 = (v7 * (^(v7)))
        v12 := GenS1{F0: ((((v6 >> ((442) & 63)) >= v6) || true) && (!true))}
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v2 -= v2
        _ = fi13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: false}
        if false {
            v5 = float64(0.0)
            v16 := []int{fi14, 127, 2147483647, (v6 << (((v4 * 2)) & 63))}
            _ = v16
        }
        var v17 int = (v6 << (((^(v4))) & 63))
        _ = fi14
        _ = v15
        _ = v17
    }
    var v18 uint8 = (v7 % ((49) | 1))
    v19 := []int{}
    if len(v19) > 0 {
        v19[0] = v6
    }
    var v20 int64 = -9223372036854775808
    for fi21 := 0; fi21 < 19; fi21++ {
        v3 -= v3
        switch fi21 {
        case 112:
            var v22 uint64 = 516
            _ = v22
        case 127:
            switch "hello" {
            case "a":
                if (false && (("" + "") <= "hello")) {
                    v4 += v4
                    if (true && ((^(v3)) >= v20)) {
                        v23 := GenS1{F0: (-9223372036854775808 < v20)}
                        _ = v23
                    }
                    var v24 uint8 = 127
                    _ = v24
                }
            case "foo bar":
                v19 = append(v19, (v6 ^ fi21))
            case "":
                v2 -= (v2 | (^(v2)))
            }
        case -693:
            switch (v6 | v6) {
            case 353:
                fmt.Println(v3)
            case -308:
                if (!(fi21 <= int(v8))) {
                    fmt.Println(v18)
                    fmt.Println(v18, v20)
                    v19 = append(v19, fi21)
                } else {
                    var v25 float64 = v5
                    _ = v25
                }
            default:
                var v26 float32 = 1e10
                _ = v26
            }
        }
        var v27 uint16 = 65535
        _ = fi21
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        _ = len(v19)
        _ = fi28
    }
    switch ("foo bar" + "a") {
    case "hello":
        var v29 string = ("\n\t" + ("foo bar" + "foo bar"))
        _ = v29
    case "a":
        if (-9223372036854775808 <= -9223372036854775808) {
            v30 := GenS1{F0: false}
            _ = v30
        }
    case "":
        if false {
            var v31 uint8 = v18
            v31 += v7
            _ = v31
        }
    default:
        v3 -= -9223372036854775808
    }
    var v32 int16 = -361
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
