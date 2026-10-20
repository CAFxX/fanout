package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }
func genFn0(p0 uint16) string {
    v4 := []int{}
    for fi5 := 0; fi5 < 20; fi5++ {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        _ = fi5
    }
    var v6 string = "foo bar"
    _ = p0
    _ = v4
    _ = v6
    return (("foo bar" + "\n\t") + (v6 + ((v6 + "\n\t") + ("a" + ("" + v6)))))
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "a"
    var v3 int = 2147483647
    v7 := GenS1{F0: 1}
    var v8 uint64 = 2
    v2 = ("\n\t" + v2)
    var v9 uint32 = 1
    v8 = uint64(v9)
    var v10 uint32 = v9
    v11 := GenS2{F0: 1, F1: (v9 << ((2) & 63)), F2: v9}
    var v12 uint32 = (^(v10))
    if false {
        fmt.Println(v2, v12, v2)
        switch "\n\t" {
        case "hello":
            var v13 float32 = 1e10
            _ = v13
        }
    }
    switch v2 {
    case "foo bar":
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                if (!(true || true)) {
                    v3 -= 128
                    switch "a" {
                    case "hello":
                        v16 := GenS1{F0: 504}
                        _ = v16
                    case "foo bar":
                        var v17 int = (v3 % (((^(v3))) | 1))
                        _ = v17
                    case "a":
                        switch (v2 + v2) {
                        case "hello":
                            var v18 float64 = float64(1.5)
                            _ = v18
                        default:
                            var v19 int64 = -9223372036854775808
                            _ = v19
                        }
                    }
                    var v20 int64 = 824
                    v9 = v9
                    _ = v20
                } else {
                    var v21 bool = false
                    if (1 < 0) {
                        v2 = (("a" + (("a" + (v2 + v2)) + ("\n\t" + "a"))) + (v2 + "hello"))
                    }
                    _ = v21
                }
                _ = fi15
            }
            _ = fi14
        }
    default:
        fmt.Println(v10, v10)
    }
    switch 32768 {
    case 0:
        switch ("\n\t" + "foo bar") {
        case "a":
            v22 := GenS3{F0: (v10 * (v12 & (v9 >> ((18446744073709551615) & 63)))), F1: (v2 + "a"), F2: 1}
            _ = v22
        case "hello":
            var v23 bool = false
            _ = v23
        case "\n\t":
            if (v2 == "") {
                v12 += (v9 << ((670) & 63))
                fmt.Println(v10)
            }
        }
    default:
        v24 := []int{(-(v3))}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
