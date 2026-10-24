package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int64, p1 int32) int {
    v7 := GenS1{F0: 380}
    v8 := []int{1}
    v9 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 35
}
func genFn1(p0 uint64, p1 uint32, p2 int) bool {
    for fi10 := 0; fi10 < 1; fi10++ {
        p2 += (p2 & -893)
        _ = fi10
    }
    p1 -= (p1 | p1)
    p0 = p0
    p0 = p0
    for fi11 := 0; fi11 < 17; fi11++ {
        if false {
            var v12 float32 = ((0.1 / 1.0) / 1e10)
            p2 = 1
            fmt.Println(fi11, p2, fi11)
            _ = v12
        }
        if true {
            if true {
                if false {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        p0 -= p0
                        _ = fi13
                    }
                }
                v14 := GenS1{F0: 42}
                v15 := []int{}
                _ = v14
                _ = v15
            } else {
                if false {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        p2 -= 32768
                        var v17 string = "hello"
                        _ = fi16
                        _ = v17
                    }
                }
                var v18 uint = 42
                p0 = p0
                _ = v18
            }
            if false {
                p2 += -168
            } else {
                p1 -= p1
            }
            if true {
                v19 := []int{4294967295, 127, genFn0(1, 15)}
                if len(v19) > 0 {
                    v19[0] = (fi11 << ((643) & 63))
                }
                p0 = (p0 >> ((4294967295) & 63))
                _ = v19
            }
        }
        p1 = 1
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return ((p0 << ((65535) & 63)) < (p0 << ((1) & 63)))
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "hello"
    v3 := make(map[string]bool)
    var v4 string = ""
    var v5 int32 = 1
    v6 := []int{-168, -9223372036854775808, 652, 32767}
    switch genFn0(757, (v5 << ((0) & 63))) {
    case 32767:
        v1 = v1
    case -9223372036854775808:
        if (42 > 4294967295) {
            if true {
                switch v2 {
                case "a":
                    var v20 int8 = 0
                    _ = v20
                case "foo bar":
                    v4 = (v4 + "foo bar")
                }
            } else {
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
                for fi21 := 0; fi21 < 14; fi21++ {
                    v3[v4] = false
                    v4 = v2
                    _ = fi21
                }
                v22 := GenS1{F0: 0}
                _ = v22
            }
            v2 = v4
        }
    default:
        v23 := GenS1{F0: 221}
        _ = v23
    }
    switch v2 {
    case "":
        for fi24 := 0; fi24 < 5; fi24++ {
            if true {
                delete(v3, "foo bar")
                v4 = v2
                switch v2 {
                case "\n\t":
                    switch 0 {
                    case 2147483647:
                        v5 += v5
                    case -1:
                        v2 = ""
                    case -828:
                        fmt.Println(v4)
                    }
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi24
        }
    case "a":
        v6 = append(v6, genFn0(-1, v5))
    default:
        v3[v4] = (false && (false || (!true)))
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            v1 = 4294967295
            _ = fi26
        }
        switch (v2 + v2) {
        case "\n\t":
            var v27 uint16 = 1
            _ = v27
        case "foo bar":
            var v28 int = (fi25 & -459)
            _ = v28
        default:
            var v29 uint8 = 42
            _ = v29
        }
        _ = fi25
    }
    _ = len(v6)
    delete(v3, v2)
    var v30 float32 = 0.0
    v4 = "\n\t"
    v5 -= v5
    v30 += v30
    var v31 uint = 2
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 bool = (!false)
        fmt.Println(v33)
        var v34 uint8 = 2
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: 32767}
    v3[v4] = (v1 == v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
