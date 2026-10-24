package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint64 }
func genFn0(p0 uint) bool {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 bool = (!false)
        var v7 int16 = -32768
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 += 4294967295
    switch "a" {
    case "hello":
        var v8 int32 = -67
        _ = v8
    default:
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(p0)
            if (585 < 42) {
                if true {
                    p0 += (p0 + p0)
                    p0 = 701
                } else {
                    var v10 int8 = -15
                    p0 = p0
                    _ = v10
                }
            } else {
                var v11 int64 = 0
                p0 = 414
                if true {
                    if true {
                        var v12 int32 = 2147483647
                        fmt.Println(p0, fi9, v12)
                        switch "foo bar" {
                        case "hello":
                            v13 := []int{fi9, fi9}
                            _ = v13
                        }
                        _ = v12
                    }
                    if (fi9 == 32768) {
                        var v14 uint8 = 255
                        if ((!false) && false) {
                            var v15 bool = (42 > -2147483648)
                            _ = v15
                        }
                        p0 = p0
                        _ = v14
                    }
                    fmt.Println(fi9, fi9)
                } else {
                    v11 = -1
                    v16 := GenS1{F0: fi9}
                    _ = v16
                }
                _ = v11
            }
            if (fi9 > fi9) {
                fmt.Println(p0)
                var v17 uint16 = 42
                _ = v17
            }
            _ = fi9
        }
    }
    switch 4294967295 {
    case 2147483647:
        var v18 uint32 = 42
        _ = v18
    case -1:
        v19 := []int{127, -1, 2147483647}
        _ = v19
    case -9223372036854775808:
        var v20 int32 = 1
        _ = v20
    }
    var v21 int16 = 27
    _ = p0
    _ = v21
    return (!(1 < 100))
}
func genFn1(p0 int, p1 int8, p2 int) string {
    fmt.Println(p2, p1, p1)
    switch "a" {
    case "\n\t":
        p2 += p2
    case "a":
        var v22 int = -5
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 int64 = 1
    var v2 int64 = -1
    v3 := make(map[int]bool)
    var v4 int16 = -32768
    fmt.Println(v2)
    var v23 int32 = 42
    v24 := GenS1{F0: 0}
    var v25 bool = true
    v26 := []int{0, -1}
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS1{F0: (fi27 * fi27)}
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: (^(v4)), F1: v1}
    var v30 int32 = -1
    v23 -= (v30 << ((610) & 63))
    for fi31 := 0; fi31 < 19; fi31++ {
        v30 = v30
        if (799 < 4294967295) {
            v4 = 42
            if (42 > 0) {
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 float32 = 1e10
                    v30 = (v30 - (v23 + (v23 | (^(v30)))))
                    v34 := GenS3{F0: (fi31 - fi31), F1: v25}
                    fmt.Println(v25, v25, fi31)
                    v35 := GenS3{F0: 32768, F1: ((fi31 / ((-573) | 1)) == (fi31 - (fi31 << ((442) & 63))))}
                    var v36 int = (fi31 & 0)
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            }
        } else {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            var v37 bool = ("" >= "a")
            _ = v37
        }
        fmt.Println(v4)
        _ = fi31
    }
    if v25 {
        var v38 int = 2147483647
        _ = v38
    } else {
        for fi39 := 0; fi39 < 10; fi39++ {
            delete(v3, fi39)
            _ = fi39
        }
        if v25 {
            for fi40 := 0; fi40 < 14; fi40++ {
                v2 = -1
                _ = fi40
            }
        }
    }
    v1 += v2
    if (42 > 0) {
        v1 = -1
        switch 32768 {
        case -1:
            for fi41 := 0; fi41 < 4; fi41++ {
                delete(v3, (fi41 & fi41))
                _ = fi41
            }
        default:
            var v42 int64 = (v1 + v2)
            _ = v42
        }
        if (2 >= 217) {
            v43 := GenS4{F0: v4}
            _ = v43
        } else {
            if v25 {
                if v25 {
                    v2 += (v2 >> ((42) & 63))
                    for fi44 := 0; fi44 < 15; fi44++ {
                        delete(v3, 127)
                        _ = fi44
                    }
                } else {
                    var v45 string = (("hello" + "\n\t") + "hello")
                    v2 -= (^(v1))
                    _ = v45
                }
            } else {
                v4 = (v4 / ((v4) | 1))
                v46 := GenS5{F0: 32768, F1: 0}
                _ = v46
            }
        }
    }
    _ = len(v26)
    v4 = -353
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
