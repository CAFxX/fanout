package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 3.14159
    v3 := []string{"a"}
    var v4 uint32 = 65535
    if len(v3) > 0 {
        v3[0] = ("\n\t" + (("a" + ("hello" + "hello")) + (("a" + "hello") + (("foo bar" + "foo bar") + "a"))))
    }
    v2 += (1e10 * 0.0)
    switch -978 {
    case -576:
        v1[("foo bar" + "\n\t")] = ("" + "\n\t")
    case -2147483648:
        if (v4 <= v4) {
            v2 += v2
            v4 += (v4 - (^(v4)))
        } else {
            switch 0 {
            case -2147483648:
                if len(v3) > 0 {
                    v3[0] = ""
                }
            case 681:
                v5 := v1["foo bar"]
                _ = v5
            default:
                v6 := GenS1{F0: ((-1 < 0) || true), F1: "hello", F2: 2147483647}
                _ = v6
            }
            if (542 == 753) {
                var v7 uint32 = v4
                _ = v7
            } else {
                var v8 uint64 = 1
                delete(v1, "hello")
                _ = v8
            }
        }
    case 803:
        if (705 != 42) {
            if true {
                v4 = (^(v4))
            }
            v2 = 1e10
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v2 -= 0.0
        _ = fi9
    }
    v4 -= (v4 << ((42) & 63))
    v10 := v1[""]
    v4 = (v4 - v4)
    fmt.Println(v10)
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v13 uint32 = 65535
            if false {
                v13 = 42
                v4 = (v4 % ((v4) | 1))
                var v14 int8 = -26
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int32 = 2147483647
            var v17 float64 = (3.14159 * 0.1)
            if ((v10 >= (v10 + v10)) || true) {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            }
            v18 := v1["a"]
            var v19 int8 = 42
            switch (v10 + v10) {
            case "a":
                var v20 int = 9223372036854775807
                _ = v20
            case "":
                var v21 int = (fi15 << ((638) & 63))
                _ = v21
            case "hello":
                v10 = v10
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
        }
        v2 += v2
        _ = fi11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
