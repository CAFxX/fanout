package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 255
    v3 := []string{(((v1 + v1) + v1) + ((("\n\t" + v1) + v1) + "foo bar")), "hello"}
    v4 := make(map[int]bool)
    var v5 int64 = 0
    var v6 float64 = 1.5
    v1 = "hello"
    if (false || (568 <= 2)) {
        delete(v4, 32768)
        v7 := GenS1{F0: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if (465 < 18446744073709551615) {
            v9 := GenS1{F0: (v5 * v5)}
            _ = v9
        }
        v5 += (v5 * int64(v2))
        v2 += v2
        _ = fi8
    }
    v1 = "a"
    fmt.Println(v5, v1, v5)
    v2 -= 4294967295
    var v10 float32 = 1e10
    var v11 uint16 = 65535
    v2 = v2
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 bool = true
        if ((0 < 0) && v13) {
            v4[(fi12 ^ fi12)] = (int16(v5) <= int16(v5))
        }
        _ = fi12
        _ = v13
    }
    switch 127 {
    case -9223372036854775808:
        if false {
            v2 -= (v2 ^ v2)
            v10 = v10
            for fi14 := 0; fi14 < 10; fi14++ {
                switch (v1 + (("foo bar" + v1) + "\n\t")) {
                case "a":
                    v15 := GenS1{F0: 42}
                    _ = v15
                }
                if (true || (!true)) {
                    v16 := GenS1{F0: 706}
                    _ = v16
                } else {
                    v10 = v10
                    fmt.Println(v5, v11)
                    var v17 uint8 = 42
                    _ = v17
                }
                _ = fi14
            }
        } else {
            if len(v3) > 0 {
                v3[0] = ("a" + (v1 + v1))
            }
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
