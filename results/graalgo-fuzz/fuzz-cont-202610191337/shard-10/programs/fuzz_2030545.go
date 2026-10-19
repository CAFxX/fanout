package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"hello", (("a" + (("hello" + "a") + "hello")) + ("hello" + ("foo bar" + "hello"))), "foo bar"}
    var v3 int16 = 32767
    var v4 int = 1
    var v5 int = v4
    _ = len(v1)
    v5 -= v4
    if false {
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                if true {
                    var v8 uint = 2
                    v9 := GenS1{F0: 0, F1: 976, F2: 1}
                    if (73 < 956) {
                        v10 := GenS1{F0: 4294967295, F1: 92, F2: 42}
                        _ = len(v2)
                        _ = v10
                    }
                    delete(v1, 2)
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
            var v11 int32 = 1
            v3 += (v3 ^ (v3 / ((v3) | 1)))
            _ = fi6
            _ = v11
        }
        v4 = (-(v5))
    }
    switch (v5 << ((557) & 63)) {
    case -2147483648:
        switch ("a" + ((("" + "a") + "") + (((("a" + ("hello" + "a")) + ("\n\t" + "hello")) + "\n\t") + ""))) {
        case "foo bar":
            fmt.Println(v5, v4, v5)
        default:
            var v12 int = 48
            _ = v12
        }
    case -1:
        switch -1 {
        case -74:
            v13 := GenS1{F0: 4294967295, F1: 596, F2: 0}
            _ = v13
        case -9223372036854775808:
            var v14 uint16 = 167
            _ = v14
        case 127:
            if len(v2) > 0 {
                v2[0] = (((("hello" + "") + ("hello" + "")) + ("\n\t" + "\n\t")) + "a")
            }
        default:
            v4 = v5
        }
    }
    var v15 uint32 = 42
    switch v4 {
    case -2147483648:
        if false {
            v5 = 730
            v15 -= v15
        }
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        _ = len(v1)
        v5 -= fi16
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(fi16, v3, v3)
            _ = fi17
        }
        _ = fi16
    }
    v1[v15] = (false && true)
    fmt.Println(v15, v4, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
