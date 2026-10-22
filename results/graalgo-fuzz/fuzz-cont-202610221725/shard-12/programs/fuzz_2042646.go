package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 int64, p1 int) uint32 {
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int = 1
    var v2 uint8 = 0
    var v3 float32 = 3.14159
    v4 := make(map[string]uint32)
    v7 := GenS1{F0: false}
    v8 := []int{(v1 % (((v1 ^ (v1 | (v1 - v1)))) | 1)), (v1 / (((^(v1))) | 1))}
    var v9 int = (v1 * v1)
    switch (-(v1)) {
    case 4294967295:
        switch v9 {
        case 578:
            v2 += v2
        case 0:
            if true {
                v9 = v9
            } else {
                v1 = (v9 * v9)
                var v10 int64 = -943
                v4["foo bar"] = 2147483647
                _ = v10
            }
        case -1:
            var v11 string = ("\n\t" + (((("foo bar" + "hello") + "") + ("" + "foo bar")) + "\n\t"))
            _ = v11
        }
    case 32768:
        if (42 <= 2147483648) {
            v12 := GenS1{F0: true}
            var v13 uint64 = 0
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v2)
    var v14 bool = (401 > -1)
    v15 := GenS1{F0: v14}
    var v16 string = ("hello" + ("hello" + "a"))
    var v17 int8 = -3
    if (v16 > ((v16 + (v16 + v16)) + v16)) {
        v18 := GenS2{F0: 957, F1: -150, F2: v17}
        _ = v18
    } else {
        if v14 {
            fmt.Println(v16, v2)
            for fi19 := 0; fi19 < 11; fi19++ {
                v4[(v16 + v16)] = 2147483647
                v16 = (v16 + ("hello" + (v16 + (v16 + "\n\t"))))
                _ = fi19
            }
        }
        if (!(!false)) {
            var v20 uint64 = 667
            _ = v20
        } else {
            switch -348 {
            case 127:
                fmt.Println(v14, v9, v16)
            case -379:
                v21 := GenS3{F0: v1, F1: 877}
                _ = v21
            case -9223372036854775808:
                for fi22 := 0; fi22 < 8; fi22++ {
                    v3 -= v3
                    var v23 float32 = float32(0.0)
                    _ = fi22
                    _ = v23
                }
            default:
                var v24 uint = 184
                _ = v24
            }
            v1 = v1
            v17 = (v17 ^ v17)
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            if true {
                fmt.Println(v9, v9, v1)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = fi25
        }
    }
    v26 := GenS2{F0: 255, F1: 259, F2: v17}
    for fi27 := 0; fi27 < 17; fi27++ {
        if v14 {
            v4[(v16 + v16)] = 0
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        v17 = v17
        _ = fi27
    }
    var v28 string = ("\n\t" + (("\n\t" + ((("\n\t" + v16) + (("\n\t" + "\n\t") + (("\n\t" + v16) + (v16 + v16)))) + "a")) + v16))
    fmt.Println(v16)
    switch 32767 {
    case -560:
        if v14 {
            fmt.Println(v17)
        } else {
            var v29 uint8 = v2
            switch v16 {
            case "":
                if (0 >= 0) {
                    if len(v8) > 0 {
                        v8[0] = v9
                    }
                } else {
                    v30 := GenS2{F0: 445, F1: 0, F2: v17}
                    if true {
                        v17 = (v17 << ((65535) & 63))
                    }
                    var v31 uint = 65535
                    _ = v30
                    _ = v31
                }
            }
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
