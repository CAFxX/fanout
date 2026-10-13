package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "\n\t"
    var v3 uint64 = 135
    var v4 float64 = -2.25
    v5 := make(map[uint32]int)
    var v6 uint32 = 2147483647
    v6 = v6
    v7 := []int{-322}
    var v8 string = v2
    v2 = "hello"
    if (74 <= 42) {
        v9 := GenS1{F0: 65535, F1: 9223372036854775807}
        for fi10 := 0; fi10 < 18; fi10++ {
            if (false || (42 == 2)) {
                if false {
                    v11 := GenS1{F0: 18446744073709551615, F1: fi10}
                    if false {
                        _ = len(v7)
                        fmt.Println(v8, v8, v2)
                    } else {
                        v8 = ((v8 + "\n\t") + v8)
                        v2 = v8
                        if (v8 == ("foo bar" + ((v8 + (v8 + v8)) + "foo bar"))) {
                            var v12 int = fi10
                            _ = v12
                        } else {
                            v5[v6] = -2147483648
                            v13 := GenS2{F0: (v6 | (v6 << ((18446744073709551615) & 63)))}
                            for fi14 := 0; fi14 < 10; fi14++ {
                                var v15 int64 = -563
                                var v16 uint32 = v6
                                switch 982 {
                                case -739:
                                    v8 = "foo bar"
                                }
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            _ = v13
                        }
                    }
                    var v17 int = (fi10 % (((-(fi10))) | 1))
                    _ = v11
                    _ = v17
                }
            }
            switch ((v8 + (v8 + v2)) + v2) {
            case "\n\t":
                fmt.Println(v4, v8, v4)
            case "foo bar":
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 string = v8
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi10
        }
        v8 = "\n\t"
        _ = v9
    }
    var v20 int32 = -867
    switch ((v2 + v2) + v8) {
    case "hello":
        _ = len(v5)
    case "a":
        if (false || true) {
            v20 += v20
            v6 -= v6
        } else {
            fmt.Println(v2)
        }
    case "foo bar":
        v4 = ((v4 - ((0.0 / v4) - (float64((0.1 - 3.14159)) * v4))) - v4)
    }
    var v21 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
