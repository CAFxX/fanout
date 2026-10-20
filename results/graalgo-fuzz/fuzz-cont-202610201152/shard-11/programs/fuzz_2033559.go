package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 int16) uint64 {
    v7 := GenS1{F0: "hello"}
    if false {
        v8 := []int{-273, 2147483647}
        _ = v8
    } else {
        p1 = -808
        var v9 uint = 122
        fmt.Println(p1, v9, v9)
        _ = v9
    }
    v10 := GenS1{F0: ""}
    v11 := []int{-2147483648, -1, -2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    return 4294967295
}

func main() {
    var v1 string = (("foo bar" + "\n\t") + "hello")
    var v2 float32 = 1.5
    v3 := []string{v1, v1, "\n\t"}
    var v4 int = 53
    var v5 uint64 = 1
    v6 := make(map[uint32]string)
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint32 = 65535
        switch v4 {
        case 32767:
            v14 := GenS2{F0: 42}
            _ = v14
        }
        v6[v13] = (("a" + "foo bar") + "hello")
        _ = fi12
        _ = v13
    }
    v4 = (v4 >> ((761) & 63))
    var v15 int32 = 86
    for fi16 := 0; fi16 < 3; fi16++ {
        if true {
            v5 = v5
        }
        v3 = append(v3, v1)
        v1 = v1
        _ = fi16
    }
    switch (v4 ^ (^(v4))) {
    case 160:
        if false {
            _ = len(v6)
            v17 := GenS1{F0: v1}
            if (int(v15) > v4) {
                fmt.Println(v2, v2, v2)
                v1 = v1
            }
            _ = v17
        }
    case 128:
        v18 := GenS1{F0: ((("a" + (v1 + ((v1 + "") + v1))) + "\n\t") + "")}
        _ = v18
    default:
        if false {
            var v19 string = (v1 + v1)
            var v20 int8 = 100
            _ = v19
            _ = v20
        } else {
            v4 += (v4 - v4)
        }
    }
    switch v4 {
    case 2147483647:
        v2 = v2
    case -443:
        if ((1000 <= -27) && ((v4 + v4) >= v4)) {
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 float32 = 1.5
                v6[42] = ""
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 int8 = -128
                    var v25 int16 = 691
                    v26 := GenS2{F0: 42}
                    _ = len(v6)
                    _ = len(v3)
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = fi21
                _ = v22
            }
        }
    case -22:
        for fi27 := 0; fi27 < 10; fi27++ {
            switch "" {
            case "hello":
                switch v1 {
                case "\n\t":
                    var v28 uint32 = 4294967295
                    _ = v28
                }
            default:
                fmt.Println(v5, v5, v4)
            }
            var v29 uint32 = 2147483648
            _ = fi27
            _ = v29
        }
    default:
        v3 = append(v3, v1)
    }
    v15 = (v15 | (v15 ^ (-(v15))))
    v1 = ("foo bar" + "hello")
    v30 := GenS3{F0: 42, F1: false, F2: 0}
    v31 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
