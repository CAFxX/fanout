package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = 0
    var v2 uint = 1
    var v3 int32 = 800
    var v4 uint64 = 4294967295
    v5 := []string{"foo bar", "hello", ("foo bar" + ("\n\t" + ""))}
    var v6 string = (("hello" + "a") + "")
    v6 = v6
    _ = len(v5)
    v5 = append(v5, (v6 + ("\n\t" + "\n\t")))
    switch -1 {
    case 127:
        switch -275 {
        case -9223372036854775808:
            v5 = append(v5, ((v6 + (((v6 + (v6 + v6)) + (((v6 + (v6 + v6)) + (v6 + "a")) + v6)) + (("a" + (v6 + "a")) + v6))) + ((v6 + "\n\t") + ((v6 + v6) + ((("\n\t" + (("hello" + ("a" + "hello")) + (v6 + v6))) + (v6 + "\n\t")) + (v6 + "hello"))))))
        case -903:
            fmt.Println(v1)
        default:
            v5 = append(v5, "\n\t")
        }
    }
    var v7 string = v6
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    var v8 uint64 = v4
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint16 = 806
        switch "hello" {
        case "a":
            var v11 float32 = 1.5
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            v7 = v7
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v13 := GenS1{F0: (((("" + "a") + v6) + (v7 + "foo bar")) + ("a" + v7))}
            fmt.Println(v8)
            if true {
                v5 = append(v5, (v6 + v7))
            }
            var v14 bool = true
            v15 := GenS1{F0: v6}
            fmt.Println(v2, v3)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if ((v3 != (v3 % (((^(v3))) | 1))) || false) {
            if (0 >= 128) {
                fmt.Println(v1, v2, v2)
            } else {
                for fi17 := 0; fi17 < 2; fi17++ {
                    _ = len(v5)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 = (v3 / (((v3 / ((2147483647) | 1))) | 1))
                    v6 = v6
                    _ = fi17
                }
            }
        } else {
            fmt.Println(fi16)
            v18 := GenS1{F0: ("a" + ((v6 + v7) + v7))}
            _ = v18
        }
        _ = fi16
    }
    var v19 uint = v2
    v2 = v19
    for fi20 := 0; fi20 < 19; fi20++ {
        v1 = 0
        var v21 uint32 = 2
        _ = fi20
        _ = v21
    }
    v5 = append(v5, v6)
    if len(v5) > 0 {
        v5[0] = "a"
    }
    var v22 bool = true
    v19 = 286
    v5 = append(v5, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
