package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 int64) int64 {
    if true {
        var v4 bool = (42 > 42)
        _ = v4
    }
    p0 += (p0 % (((p0 / (((p0 | p0)) | 1))) | 1))
    _ = p0
    _ = p1
    return (p1 >> ((464) & 63))
}
func genFn1(p0 string, p1 int16, p2 bool) uint {
    v5 := []int{}
    var v6 int = 46
    v7 := GenS1{F0: (v6 * (v6 >> ((2) & 63)))}
    p2 = (!p2)
    for fi8 := 0; fi8 < 1; fi8++ {
        v6 += v6
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 1
    v2 := []string{("" + "\n\t"), ("\n\t" + "a"), "a", ("\n\t" + (("a" + (("hello" + "\n\t") + "hello")) + ("foo bar" + ("a" + "a"))))}
    var v3 int16 = -446
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi9
        }
    }
    var v10 uint = 2
    var v11 uint64 = 0
    v10 = (v10 + (v10 / (((v10 >> ((932) & 63))) | 1)))
    v2 = append(v2, "hello")
    _ = len(v2)
    v12 := GenS1{F0: 9223372036854775807}
    if true {
        v10 = (v10 * (v10 * v10))
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int = int(v1)
            if false {
                v2 = append(v2, "a")
                var v15 uint32 = 4294967295
                _ = v15
            } else {
                switch 128 {
                case -2147483648:
                    fmt.Println(fi13)
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    if ((^(v10)) < genFn1("", v3, true)) {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        v14 = 2147483647
                        switch "foo bar" {
                        case "\n\t":
                            _ = len(v2)
                        }
                    }
                    _ = fi16
                }
            }
            _ = fi13
            _ = v14
        }
        var v17 int = 378
        _ = v17
    } else {
        var v18 string = ""
        var v19 uint16 = 42
        var v20 bool = ((^(v11)) == v11)
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: -2147483648}
    v22 := GenS1{F0: 420}
    fmt.Println(v11)
    var v23 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
