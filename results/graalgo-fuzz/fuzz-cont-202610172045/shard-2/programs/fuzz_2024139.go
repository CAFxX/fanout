package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint64 = 1
    var v2 string = ("" + ("foo bar" + ("\n\t" + ("" + ""))))
    var v3 uint16 = 65535
    v4 := []int{}
    var v5 float64 = 3.14159
    v5 += v5
    var v6 float32 = (3.14159 - 0.1)
    fmt.Println(v1)
    if (v2 != ("" + (v2 + ""))) {
        var v7 uint = 139
        _ = v7
    } else {
        if len(v4) > 0 {
            v4[0] = 605
        }
    }
    var v8 uint32 = 2147483648
    fmt.Println(v1, v5, v2)
    fmt.Println(v3, v1)
    var v9 int32 = -1
    v1 -= 0
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v10 bool = (1 < 1)
        var v11 uint64 = (v1 << ((65535) & 63))
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: (v8 - (v8 & (v8 & v8)))}
            if (9223372036854775807 <= int64(fi12)) {
                var v14 uint32 = v8
                switch (v2 + "") {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = (-(fi12))
                    }
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v15 uint16 = 42
                _ = v14
                _ = v15
            }
            var v16 bool = false
            _ = fi12
            _ = v13
            _ = v16
        }
        fmt.Println(v2, v9)
        var v17 int64 = 1
        _ = v17
    }
    if (true && true) {
        v18 := GenS1{F0: (^(v8))}
        if true {
            v8 -= 0
        } else {
            switch "" {
            case "":
                v4 = append(v4, -1)
            case "\n\t":
                v19 := GenS1{F0: uint32(v1)}
                _ = v19
            }
            v6 = v6
            v5 -= v5
        }
        v6 -= v6
        _ = v18
    } else {
        switch (v2 + v2) {
        case "a":
            var v20 uint32 = 2147483648
            _ = v20
        default:
            v21 := GenS2{F0: 100, F1: 42, F2: -671}
            _ = v21
        }
        _ = len(v4)
        v3 = 2
    }
    if true {
        for fi22 := 0; fi22 < 6; fi22++ {
            v6 -= v6
            _ = fi22
        }
        if true {
            var v23 float64 = 3.14159
            v24 := GenS3{F0: (9223372036854775807 >= 2147483647), F1: 44}
            _ = v23
            _ = v24
        }
        var v25 int16 = 32767
        _ = v25
    } else {
        if true {
            v26 := GenS2{F0: 98, F1: 9223372036854775807, F2: int16(v9)}
            _ = v26
        }
    }
    var v27 float64 = v5
    fmt.Println(v2)
    v28 := GenS3{F0: false, F1: 69}
    if (((2 > 310) || false) && false) {
        var v29 string = (v2 + (v2 + "hello"))
        v30 := GenS4{F0: v1}
        for fi31 := 0; fi31 < 18; fi31++ {
            switch ("" + "hello") {
            case "foo bar":
                v9 = -2147483648
            default:
                var v32 float64 = (3.14159 / (v5 * 0.1))
                _ = v32
            }
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 uint = 255
                fmt.Println(v8, v8)
                var v35 uint16 = v3
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi31
        }
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
