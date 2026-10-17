package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0(p0 uint8, p1 uint8) uint8 {
    v7 := GenS1{F0: 42, F1: -206}
    switch "a" {
    case "":
        var v8 string = ((("foo bar" + (("hello" + ("a" + "foo bar")) + "foo bar")) + "foo bar") + "\n\t")
        _ = v8
    }
    p1 -= (p1 + p1)
    _ = p0
    _ = p1
    _ = v7
    return (p1 + (^(p1)))
}

func main() {
    var v1 int64 = 39
    var v2 int16 = 1000
    var v3 uint8 = 68
    v4 := []string{(("" + (("hello" + "hello") + "")) + "\n\t"), (("foo bar" + "foo bar") + ("a" + "foo bar")), "", "a"}
    var v5 string = ""
    v6 := []int{0, 127}
    var v9 int = 9223372036854775807
    v10 := GenS1{F0: 18446744073709551615, F1: v1}
    v2 = -761
    fmt.Println(v9, v5)
    switch (^(v9)) {
    case -494:
        v2 = 1000
    case 335:
        fmt.Println(v1, v9)
    case 293:
        if true {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int64 = (v1 >> ((1) & 63))
                _ = fi11
                _ = v12
            }
            _ = len(v4)
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                v5 = v5
                _ = fi13
            }
            if ("\n\t" <= ("" + (v5 + v5))) {
                _ = len(v4)
                if false {
                    v9 += (v9 & 1)
                }
                var v14 bool = (!(1 != 65535))
                _ = v14
            } else {
                v15 := GenS2{F0: 546, F1: (v1 >> ((255) & 63)), F2: 42}
                switch (v9 / ((v9) | 1)) {
                case -9223372036854775808:
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v1 += (v1 + v1)
                        var v17 int = v9
                        _ = fi16
                        _ = v17
                    }
                case 40:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v18 := GenS1{F0: 396, F1: v1}
                _ = v15
                _ = v18
            }
        }
    default:
        var v19 int64 = (v1 >> ((65535) & 63))
        _ = v19
    }
    v2 += (-(v2))
    switch v9 {
    case -1:
        if false {
            var v20 uint32 = 2
            _ = v20
        }
    default:
        v1 = (v1 / ((9223372036854775807) | 1))
    }
    var v21 int64 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
