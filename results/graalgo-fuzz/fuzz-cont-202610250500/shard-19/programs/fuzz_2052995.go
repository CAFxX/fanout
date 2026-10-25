package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int16 = 0
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v4 := GenS1{F0: 0}
        _ = v4
    case 2147483647:
        delete(v2, "\n\t")
    }
    switch "foo bar" {
    case "hello":
        if (1 < 370) {
            switch 32768 {
            case 128:
                var v5 float32 = (0.1 + (1.5 * 3.14159))
                _ = v5
            default:
                for fi6 := 0; fi6 < 2; fi6++ {
                    if (108 > 2147483648) {
                        if true {
                            var v7 int = (fi6 ^ fi6)
                            v1 = 0
                            _ = v7
                        }
                        v1 -= 678
                        fmt.Println(v1, v1)
                    }
                    for fi8 := 0; fi8 < 12; fi8++ {
                        v1 = v1
                        v9 := GenS1{F0: 498}
                        v1 -= (v1 * v1)
                        _ = fi8
                        _ = v9
                    }
                    _ = fi6
                }
            }
        } else {
            v10 := GenS2{F0: (v1 ^ 42)}
            _ = v10
        }
    case "\n\t":
        if false {
            var v11 uint64 = 42
            fmt.Println(v11)
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 1; fi12++ {
            v1 = v1
            var v13 uint8 = 1
            _ = fi12
            _ = v13
        }
    }
    var v14 string = ("" + ("a" + ((("\n\t" + "hello") + "") + "")))
    switch 2147483647 {
    case 4294967295:
        v15 := []int{-1, 128, 9223372036854775807, -1}
        _ = v15
    }
    v1 = v1
    for fi16 := 0; fi16 < 4; fi16++ {
        v1 -= (v1 * (v1 << ((76) & 63)))
        switch (^(fi16)) {
        case -28:
            for fi17 := 0; fi17 < 10; fi17++ {
                switch v14 {
                case "hello":
                    v18 := v2[v14]
                    _ = v18
                case "foo bar":
                    var v19 int = (fi17 << ((280) & 63))
                    _ = v19
                }
                var v20 float64 = (0.0 + -2.25)
                if false {
                    v21 := []int{fi16}
                    var v22 int64 = 0
                    switch v14 {
                    case "foo bar":
                        var v23 bool = false
                        _ = v23
                    case "a":
                        fmt.Println(v1, v1, fi17)
                    }
                    switch "hello" {
                    case "":
                        fmt.Println(fi16, v1)
                    default:
                        var v24 uint64 = 0
                        _ = v24
                    }
                    fmt.Println(fi16)
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v20
            }
        case 42:
            for fi25 := 0; fi25 < 14; fi25++ {
                fmt.Println(fi25, fi25)
                _ = fi25
            }
        }
        v1 = v1
        _ = fi16
    }
    var v26 int = 0
    switch 127 {
    case 4294967295:
        v1 += 758
    case 128:
        switch v26 {
        case 128:
            v14 = v14
        case 127:
            for fi27 := 0; fi27 < 13; fi27++ {
                v1 = (v1 << ((1) & 63))
                _ = fi27
            }
        case -592:
            var v28 int16 = 1
            _ = v28
        }
    case 646:
        fmt.Println(v14)
    }
    if (false || ((4294967295 <= 65535) || (false && true))) {
        var v29 uint32 = 4294967295
        _ = v29
    } else {
        if true {
            var v30 uint = 327
            if ((0.0 - 0.0) >= float64(1.5)) {
                var v31 uint32 = 824
                var v32 uint8 = 1
                var v33 uint32 = v31
                _ = v31
                _ = v32
                _ = v33
            }
            v34 := GenS2{F0: (v1 % ((535) | 1))}
            _ = v30
            _ = v34
        }
        var v35 uint32 = 2
        v35 -= v35
        _ = v35
    }
    v14 = "hello"
    fmt.Println(v14, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
}
