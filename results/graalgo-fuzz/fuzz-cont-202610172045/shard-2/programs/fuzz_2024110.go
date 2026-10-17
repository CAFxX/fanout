package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() int32 {
    var v4 string = "hello"
    v4 = ""
    fmt.Println(v4)
    v5 := GenS1{F0: 32767}
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{}
        v8 := GenS1{F0: 65535}
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int16 = 259
    v2 := []int{}
    v3 := make(map[int]string)
    delete(v3, 32768)
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 bool = (155 > 0)
        if (100 > 42) {
            v10 = v10
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v1)
    switch (("foo bar" + "a") + ("foo bar" + "a")) {
    case "\n\t":
        v1 = (v1 + v1)
    case "hello":
        if false {
            v1 = (v1 % ((v1) | 1))
            var v11 int = -9223372036854775808
            _ = v11
        }
    default:
        fmt.Println(v1)
    }
    v12 := GenS1{F0: 61}
    switch -9223372036854775808 {
    case 1:
        if (2147483647 < 2147483648) {
            v1 = v1
            fmt.Println(v1)
        } else {
            fmt.Println(v1)
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 uint32 = 42
                fmt.Println(v1)
                if (0 <= 0) {
                    if false {
                        var v15 bool = true
                        fmt.Println(v15, v14, v1)
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v1 = (v1 * 0)
                            if (fi16 < -1) {
                                v14 -= 65535
                            }
                            _ = fi16
                        }
                        _ = v15
                    }
                }
                _ = fi13
                _ = v14
            }
            v1 += (-(v1))
        }
    case 42:
        for fi17 := 0; fi17 < 8; fi17++ {
            switch ("a" + "\n\t") {
            case "hello":
                var v18 string = ((("foo bar" + (("" + "\n\t") + "\n\t")) + "foo bar") + ("\n\t" + ("\n\t" + ("" + ""))))
                _ = v18
            }
            switch ("\n\t" + ("" + "\n\t")) {
            case "hello":
                v19 := GenS1{F0: 0}
                _ = v19
            default:
                v1 = v1
            }
            switch "\n\t" {
            case "a":
                for fi20 := 0; fi20 < 8; fi20++ {
                    v1 -= -553
                    v1 = (-(v1))
                    _ = fi20
                }
            }
            _ = fi17
        }
    case 407:
        var v21 int = -1
        _ = v21
    default:
        if (322 != -20) {
            fmt.Println(v1)
            var v22 bool = true
            v23 := GenS1{F0: 255}
            _ = v22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v1 = v1
        if true {
            var v25 int = (fi24 * fi24)
            fmt.Println(v25, fi24, v25)
            var v26 uint32 = 2
            _ = v25
            _ = v26
        }
        _ = fi24
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        fmt.Println(v1, fi27)
        v1 = v1
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
