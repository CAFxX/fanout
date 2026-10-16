package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0() int8 {
    var v5 uint8 = 197
    v5 += (v5 - v5)
    v5 -= v5
    var v6 float64 = 1e10
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 string, p1 int32, p2 int) uint64 {
    fmt.Println(p1, p2)
    var v7 uint64 = 0
    var v8 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 2
    var v3 int64 = -9223372036854775808
    var v4 int8 = 1
    v9 := []int{-2147483648, -552, -2147483648, -9223372036854775808}
    if false {
        v10 := GenS1{F0: 977}
        var v11 string = ""
        _ = v10
        _ = v11
    }
    var v12 int16 = 1
    _ = len(v9)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v1 = v1
    var v13 int = 32768
    var v14 int16 = v12
    var v15 uint = 4294967295
    var v16 string = ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))
    if (true || false) {
        fmt.Println(v1, v16, v14)
    }
    v16 = "\n\t"
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint = (v15 | (v15 & (v15 % ((v15) | 1))))
            var v20 uint64 = 18446744073709551615
            var v21 int = fi18
            if len(v9) > 0 {
                v9[0] = -1
            }
            fmt.Println(v3, v19)
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        v1 = (^(v1))
        switch ((v16 + ("a" + "a")) + v16) {
        case "\n\t":
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 int16 = (v14 - (v14 >> ((42) & 63)))
                switch ("a" + v16) {
                case "\n\t":
                    var v24 int32 = 2147483647
                    _ = v24
                case "hello":
                    fmt.Println(fi17)
                case "":
                    if len(v9) > 0 {
                        v9[0] = (fi17 / ((2147483647) | 1))
                    }
                }
                if (1 <= 1) {
                    v25 := GenS2{F0: v1, F1: 2147483648}
                    v14 -= v23
                    if false {
                        _ = len(v9)
                        v9 = append(v9, (-(fi17)))
                        v13 = v13
                        v26 := GenS1{F0: 255}
                        fmt.Println(v16, v4, v14)
                        _ = v26
                    }
                    _ = v25
                } else {
                    var v27 uint8 = v1
                    _ = v27
                }
                _ = fi22
                _ = v23
            }
        case "":
            if true {
                v13 -= int(v12)
                v14 -= v14
            } else {
                v9 = append(v9, 42)
                v15 -= 255
            }
        }
        _ = fi17
    }
    v28 := GenS2{F0: v1, F1: 2147483648}
    if true {
        var v29 int8 = (v4 >> ((v15) & 63))
        _ = v29
    } else {
        var v30 string = "a"
        if false {
            if (!((1000 == v12) && (!false))) {
                _ = len(v9)
                var v31 int32 = 1
                if (genFn0() >= int8(v1)) {
                    switch v30 {
                    case "foo bar":
                        var v32 uint = (v15 * (v15 - 65535))
                        _ = v32
                    case "a":
                        v2 = (^(v2))
                    }
                } else {
                    v33 := GenS3{F0: 32767, F1: -427}
                    _ = v33
                }
                _ = v31
            }
            v13 += v13
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            switch v30 {
            case "\n\t":
                var v35 int8 = (^(v4))
                _ = v35
            }
            v36 := GenS1{F0: 341}
            _ = fi34
            _ = v36
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
