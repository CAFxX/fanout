package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0(p0 bool, p1 uint32) int64 {
    p1 -= (p1 | 405)
    switch "a" {
    case "hello":
        for fi5 := 0; fi5 < 17; fi5++ {
            switch (fi5 >> ((1) & 63)) {
            case -2147483648:
                for fi6 := 0; fi6 < 16; fi6++ {
                    var v7 uint = 1
                    p1 = (p1 & p1)
                    v8 := []int{fi6, (fi5 + fi5), 32768}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            }
            _ = fi5
        }
    case "":
        if p0 {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := []int{(fi9 - -110), 32768, (-(fi9))}
                fmt.Println(fi9, p0)
                _ = fi9
                _ = v10
            }
        }
    case "\n\t":
        var v11 int = 127
        _ = v11
    }
    var v12 bool = p0
    var v13 bool = p0
    v13 = v12
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 556
}
func genFn1(p0 uint64, p1 uint32, p2 uint32) string {
    v14 := GenS1{F0: 127}
    var v15 uint64 = p0
    var v16 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return ""
}

func main() {
    v1 := []uint32{1}
    v2 := []string{"a", "hello", ("foo bar" + ("\n\t" + ("" + "foo bar"))), ("hello" + "hello")}
    v3 := make(map[string]string)
    v4 := []int{-774, 42}
    var v17 uint64 = 2
    var v18 int32 = -668
    if ((374 != 4294967295) || true) {
        if (2 < 127) {
            switch 0 {
            case -642:
                var v19 uint64 = (v17 % (((^(v17))) | 1))
                _ = v19
            case 127:
                var v20 int = 32768
                _ = v20
            }
            var v21 int64 = 42
            _ = v21
        } else {
            v17 = (v17 * 2)
        }
        var v22 bool = (("hello" + ("" + "\n\t")) <= "")
        var v23 uint16 = 1
        _ = v22
        _ = v23
    } else {
        v17 = (v17 ^ v17)
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v18 = v18
        _ = fi24
    }
    v18 = (v18 + 0)
    v17 = 0
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 bool = (fi25 < (^(fi25)))
        var v27 uint64 = v17
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: (!false), F1: 2, F2: -128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
