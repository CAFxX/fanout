package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{0, 0}
    var v3 uint16 = 246
    var v4 int8 = -10
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    _ = len(v2)
    v7 := GenS1{F0: "foo bar", F1: 0, F2: v4}
    v8 := GenS1{F0: ("" + (("foo bar" + ("" + "\n\t")) + "")), F1: 837, F2: (v4 - (^(v4)))}
    v9 := GenS1{F0: "foo bar", F1: 42, F2: (v4 - v4)}
    if (32768 == -865) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v4 = v4
            var v11 uint32 = 2
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS2{F0: ("" + "foo bar"), F1: (v11 | uint32(v3))}
                v3 = 255
                if ((fi10 % ((fi12) | 1)) <= fi10) {
                    delete(v6, 128)
                    var v14 uint = 4294967295
                    if (fi12 <= (fi12 - (fi12 / (((fi10 / ((9223372036854775807) | 1))) | 1)))) {
                        v15 := GenS1{F0: "foo bar", F1: uint(v3), F2: (v4 << ((v14) & 63))}
                        _ = v15
                    }
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        v4 += v4
        fmt.Println(v3, v4)
    } else {
        v2 = append(v2, 2)
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint16 = (v3 << ((65535) & 63))
            v6[-847] = ("hello" + ("foo bar" + ""))
            fmt.Println(fi17, v4, v18)
            _ = fi17
            _ = v18
        }
        v3 = (^(v3))
        _ = fi16
    }
    fmt.Println(v4, v3)
    v6[822] = (("foo bar" + "") + ("a" + ("" + "hello")))
    var v19 string = "\n\t"
    if (399 < 639) {
        if true {
            if true {
                v20 := GenS2{F0: "foo bar", F1: 2}
                _ = v20
            } else {
                v21 := v1[911]
                v22 := GenS3{F0: 2147483647}
                v4 += (v4 + 100)
                _ = v21
                _ = v22
            }
            v23 := GenS4{F0: 580}
            var v24 int32 = -1
            _ = v23
            _ = v24
        }
    } else {
        var v25 int8 = v4
        _ = v25
    }
    switch (((v19 + v19) + v19) + v19) {
    case "foo bar":
        switch 1 {
        case -2147483648:
            var v26 uint32 = 647
            _ = v26
        default:
            fmt.Println(v4)
        }
    default:
        _ = len(v2)
    }
    fmt.Println(v19, v4)
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
