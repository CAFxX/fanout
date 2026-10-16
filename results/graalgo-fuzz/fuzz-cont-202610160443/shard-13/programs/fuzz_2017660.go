package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{-2147483648, 353, 9223372036854775807, -9223372036854775808}
    v2 := []uint32{42, 0}
    var v3 uint64 = 1
    var v4 int8 = 100
    var v5 uint16 = 65535
    v6 := make(map[string]bool)
    v3 = (v3 ^ 232)
    for fi7 := 0; fi7 < 8; fi7++ {
        v4 = (v4 << ((uint(v5)) & 63))
        _ = fi7
    }
    if false {
        var v8 uint32 = 320
        if false {
            var v9 int = 1
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if (!true) {
                if false {
                    delete(v6, "a")
                    var v10 uint32 = (^(v8))
                    _ = v10
                } else {
                    var v11 int = v9
                    _ = v11
                }
                v3 -= v3
            } else {
                delete(v6, ("hello" + (("foo bar" + "") + ((("foo bar" + ("hello" + "hello")) + "hello") + "hello"))))
                var v12 uint64 = v3
                fmt.Println(v8)
                _ = v12
            }
            _ = v9
        } else {
            v8 = (v8 ^ v8)
        }
        var v13 uint8 = 127
        _ = v13
        _ = v8
    }
    switch 0 {
    case -2147483648:
        switch "a" {
        case "":
            var v14 int32 = 1
            _ = v14
        case "foo bar":
            v5 -= v5
        }
    case 9223372036854775807:
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(fi15)
            var v16 float64 = -2.25
            _ = fi15
            _ = v16
        }
    }
    switch "\n\t" {
    case "":
        if (v3 == uint64(v5)) {
            var v17 bool = true
            switch "\n\t" {
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    v18 := GenS1{F0: 18}
                    _ = v18
                }
            }
            _ = v17
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS2{F0: -9223372036854775808}
                v3 = (v3 & 355)
                if len(v1) > 0 {
                    v1[0] = fi19
                }
                _ = fi19
                _ = v20
            }
            v5 -= 42
        }
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(fi21)
        _ = fi21
    }
    var v22 uint = 4294967295
    switch "" {
    case "a":
        v4 = (v4 / ((v4) | 1))
    case "hello":
        v4 = (v4 ^ -128)
    case "foo bar":
        if (true && true) {
            var v23 uint64 = (v3 + uint64(v22))
            v23 = (v3 & 1)
            _ = v23
        } else {
            var v24 uint16 = v5
            for fi25 := 0; fi25 < 2; fi25++ {
                if (!true) {
                    v26 := GenS1{F0: 0}
                    if (697 == -1) {
                        fmt.Println(v3, v5)
                        fmt.Println(v3, fi25)
                    }
                    _ = v26
                }
                _ = fi25
            }
            _ = v24
        }
    default:
        if false {
            var v27 uint32 = 698
            var v28 float32 = 0.0
            _ = v27
            _ = v28
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
