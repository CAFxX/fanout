package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 uint16 = 834
    v4 += v4
    for fi5 := 0; fi5 < 11; fi5++ {
        if (!(true && false)) {
            v4 += (v4 << ((828) & 63))
            fmt.Println(fi5, fi5)
        } else {
            v6 := GenS1{F0: 42, F1: 128, F2: 32767}
            _ = v6
        }
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{(fi7 & (fi5 ^ 32767)), fi5, fi7, int(v4)}
            v4 = 65535
            var v9 int32 = -239
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    var v10 uint8 = 255
    if false {
        v10 += 1
    } else {
        fmt.Println(v4, v10)
    }
    var v11 bool = (true && (1.5 >= (((0.1 * (1e10 - (float32(1e10) * -2.25))) / 1e10) / 3.14159)))
    var v12 uint = 589
    fmt.Println(v12)
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := v3[2]
        v1[1] = 0
        switch (fi13 >> (((v12 | v12)) & 63)) {
        case 32767:
            var v15 uint16 = 65535
            _ = v15
        case 127:
            v14 = ""
        default:
            switch (v14 + ((((((("\n\t" + "\n\t") + (("foo bar" + v14) + ("\n\t" + ("foo bar" + v14)))) + v14) + "") + v14) + "\n\t") + v14)) {
            case "a":
                var v16 uint16 = v4
                _ = v16
            case "hello":
                v11 = v11
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v4)
    if false {
        v17 := GenS1{F0: 883, F1: 2147483647, F2: 705}
        v18 := v3[304]
        fmt.Println(v12, v11)
        _ = v17
        _ = v18
    }
    v11 = ("\n\t" <= ("hello" + "hello"))
    switch -9223372036854775808 {
    case 32768:
        v11 = v11
    case 0:
        if v11 {
            if true {
                if (127 < 4294967295) {
                    if v11 {
                        delete(v2, 63)
                    } else {
                        v19 := []int{128, -680}
                        _ = v19
                    }
                    v20 := []int{73, 9223372036854775807, -1}
                    fmt.Println(v10, v11)
                    _ = v20
                }
                fmt.Println(v10, v4)
                var v21 int = int(v10)
                _ = v21
            } else {
                switch ("hello" + ("" + "a")) {
                case "hello":
                    var v22 float64 = (1.5 - 0.1)
                    _ = v22
                default:
                    if (-1 != 100) {
                        v12 += (v12 >> ((v12) & 63))
                        v12 += v12
                    }
                }
                v23 := GenS1{F0: 96, F1: -705, F2: v4}
                var v24 bool = (v11 && (false || v11))
                _ = v23
                _ = v24
            }
        }
    }
    v25 := GenS1{F0: 0, F1: 42, F2: (v4 - (v4 | (v4 / ((728) | 1))))}
    for fi26 := 0; fi26 < 9; fi26++ {
        v11 = v11
        _ = fi26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v11, v11)
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
