package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 string }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    var v1 string = (((("a" + ("hello" + "a")) + "a") + "\n\t") + "\n\t")
    var v2 float32 = 3.14159
    var v3 int8 = 57
    var v4 uint = 521
    var v5 uint16 = 65535
    v6 := []int{127}
    v7 := GenS1{F0: 42, F1: v4, F2: (("\n\t" + v1) + (v1 + ("\n\t" + (v1 + v1))))}
    if true {
        var v8 uint8 = 128
        v2 = v2
        _ = v8
    }
    var v9 uint32 = 2147483648
    var v10 uint32 = v9
    v11 := GenS2{F0: "hello", F1: -1}
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(v10, v4, v10)
    }
    if (0 == -9223372036854775808) {
        fmt.Println(v10, v5)
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int8 = v3
        switch (fi12 >> ((0) & 63)) {
        case 2147483647:
            if (!false) {
                if len(v6) > 0 {
                    v6[0] = fi12
                }
                if len(v6) > 0 {
                    v6[0] = int(v13)
                }
            } else {
                for fi14 := 0; fi14 < 18; fi14++ {
                    fmt.Println(v13, v3, v13)
                    _ = fi14
                }
                v15 := GenS2{F0: v1, F1: -690}
                for fi16 := 0; fi16 < 14; fi16++ {
                    switch v1 {
                    case "\n\t":
                        v9 += v9
                    }
                    v17 := GenS2{F0: "foo bar", F1: -189}
                    var v18 float32 = v2
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        }
        v19 := GenS1{F0: 1, F1: (^(v4)), F2: ""}
        _ = fi12
        _ = v13
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint16 = v5
        switch (v1 + "a") {
        case "":
            fmt.Println(fi20)
        case "hello":
            v6 = append(v6, fi20)
        case "\n\t":
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 int = fi20
                _ = fi22
                _ = v23
            }
        default:
            fmt.Println(v4, v5, v9)
        }
        _ = fi20
        _ = v21
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            v5 += (v5 / ((v5) | 1))
            _ = fi25
        }
        _ = fi24
    }
    var v26 uint8 = uint8(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
