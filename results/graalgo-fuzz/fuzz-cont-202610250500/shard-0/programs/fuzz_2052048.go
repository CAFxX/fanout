package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 float64 = (1e10 / 1e10)
    v2 := make(map[string]bool)
    var v3 int8 = 42
    if true {
        for fi4 := 0; fi4 < 19; fi4++ {
            if (!((!true) || true)) {
                v5 := GenS1{F0: 2147483647}
                var v6 bool = true
                delete(v2, "\n\t")
                _ = v5
                _ = v6
            } else {
                delete(v2, "a")
                v1 = v1
                for fi7 := 0; fi7 < 2; fi7++ {
                    var v8 bool = true
                    _ = fi7
                    _ = v8
                }
            }
            var v9 uint64 = 2
            _ = fi4
            _ = v9
        }
        v10 := GenS1{F0: 2}
        v3 -= v3
        _ = v10
    } else {
        var v11 uint = 4294967295
        v1 -= 1.5
        v12 := GenS2{F0: 2, F1: 398}
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: true, F1: 18446744073709551615}
    if ("\n\t" <= (("foo bar" + ("a" + "")) + "a")) {
        fmt.Println(v1, v3)
    } else {
        var v14 int32 = 1
        switch 575 {
        case 42:
            fmt.Println(v1, v3)
        case 128:
            for fi15 := 0; fi15 < 3; fi15++ {
                if false {
                    var v16 int = (fi15 >> ((42) & 63))
                    if (42 < 270) {
                        fmt.Println(v14)
                        fmt.Println(v3, v16, fi15)
                        v1 = (v1 + (0.0 + v1))
                    }
                    _ = v16
                }
                _ = fi15
            }
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            switch (fi17 & fi17) {
            case 2147483647:
                v1 = (1e10 + (0.1 * v1))
            case 9223372036854775807:
                fmt.Println(v1, v3)
            default:
                switch "\n\t" {
                case "":
                    v2[("foo bar" + (("" + ((("a" + ("hello" + "a")) + "\n\t") + "foo bar")) + "hello"))] = true
                case "a":
                    v14 += v14
                default:
                    var v18 int32 = (v14 - 299)
                    _ = v18
                }
            }
            var v19 uint = 255
            _ = fi17
            _ = v19
        }
        _ = v14
    }
    v3 += 42
    for fi20 := 0; fi20 < 4; fi20++ {
        v1 += float64(-2.25)
        var v21 bool = (true || true)
        _ = fi20
        _ = v21
    }
    v22 := GenS4{F0: 208}
    v23 := []int{-455}
    v3 += v3
    var v24 int64 = 9223372036854775807
    var v25 int = 9223372036854775807
    v23 = append(v23, int(v24))
    v24 = (v24 >> ((381) & 63))
    var v26 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
