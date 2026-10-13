package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int32, p1 int8, p2 uint32) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{-1}
        var v7 uint32 = (p2 >> ((42) & 63))
        _ = fi5
        _ = v6
        _ = v7
    }
    if ((84 <= 121) || (1 < 2)) {
        switch "" {
        case "\n\t":
            p0 -= p0
        case "":
            switch 2147483647 {
            case 32767:
                var v8 uint64 = 474
                _ = v8
            case 2147483647:
                var v9 int8 = (-(p1))
                _ = v9
            case -9223372036854775808:
                fmt.Println(p1, p0)
            }
        case "foo bar":
            var v10 bool = (uint16(p2) < 65535)
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: (-(fi11)), F1: 0}
            if true {
                if false {
                    v13 := []int{fi11}
                    _ = v13
                } else {
                    v14 := []int{int(p1), (fi11 * (fi11 - (-(fi11)))), fi11}
                    p1 -= (p1 << ((255) & 63))
                    _ = v14
                }
                p2 += 782
                var v15 uint8 = 128
                _ = v15
            } else {
                var v16 int = fi11
                p1 += p1
                switch (("\n\t" + ("\n\t" + "\n\t")) + "a") {
                case "a":
                    if false {
                        fmt.Println(p2, p0)
                        var v17 uint32 = 374
                        p2 = v17
                        _ = v17
                    } else {
                        v16 = fi11
                        fmt.Println(p2, p2)
                    }
                }
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(p1, p1, p2)
            p1 -= (p1 ^ p1)
            _ = fi18
        }
    }
    p0 -= 42
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() int {
    var v19 int64 = 42
    if (4294967295 <= 348) {
        var v20 uint64 = 4294967295
        fmt.Println(v20)
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := GenS1{F0: (fi21 - (fi21 + fi21)), F1: -957}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 2147483647, F1: 463}
    for fi24 := 0; fi24 < 5; fi24++ {
        if false {
            v25 := GenS1{F0: -2147483648, F1: -775}
            _ = v25
        }
        v26 := []int{9223372036854775807, (fi24 | fi24), (^(fi24)), fi24}
        _ = fi24
        _ = v26
    }
    _ = v19
    _ = v23
    return 2147483647
}

func main() {
    v1 := []string{}
    var v2 int64 = 754
    v3 := make(map[int]int)
    v4 := []string{"hello", (((("\n\t" + "a") + ("\n\t" + (("foo bar" + "foo bar") + ("\n\t" + ((("\n\t" + ("\n\t" + "a")) + "a") + "a"))))) + "foo bar") + "\n\t")}
    _ = len(v3)
    _ = len(v3)
    for fi27 := 0; fi27 < 9; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            v29 := v3[genFn1()]
            v2 = 615
            _ = fi28
            _ = v29
        }
        var v30 uint32 = 42
        _ = fi27
        _ = v30
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        v2 -= 8
        v2 = v2
        v32 := GenS1{F0: genFn1(), F1: 1}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        if false {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            var v34 uint32 = 2
            var v35 int = 0
            _ = v34
            _ = v35
        }
        var v36 int32 = 0
        _ = fi33
        _ = v36
    }
    fmt.Println(v2)
    v2 -= (v2 * 0)
    v3[2147483647] = 2147483647
    v37 := GenS1{F0: 32767, F1: -474}
    v38 := GenS1{F0: genFn1(), F1: -383}
    v2 = v2
    v2 = v2
    v2 = v2
    v39 := GenS2{F0: 2147483647, F1: 306}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
