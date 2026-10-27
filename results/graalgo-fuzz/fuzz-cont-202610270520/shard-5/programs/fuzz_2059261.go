package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int }
func genFn0(p0 int64, p1 int16, p2 string) uint16 {
    if (!false) {
        var v6 uint16 = 792
        fmt.Println(p2, v6, p2)
        var v7 uint64 = 2
        _ = v6
        _ = v7
    }
    if true {
        p1 = (p1 | p1)
        v8 := GenS1{F0: 521, F1: (p2 + "a")}
        v9 := GenS1{F0: 0, F1: ("hello" + p2)}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 111
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    var v3 uint16 = 1
    var v4 float32 = (1e10 - 1.5)
    v5 := make(map[int]int)
    if false {
        var v10 uint32 = 0
        fmt.Println(v3, v2, v4)
        if true {
            v3 -= v3
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := v1[(("foo bar" + "a") + "hello")]
                var v13 int = (v12 * -649)
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = v10
    } else {
        if (127 < 1) {
            if false {
                fmt.Println(v3)
            }
        }
        v4 -= (0.0 + 0.0)
        v2 -= (v2 - v2)
    }
    var v14 uint64 = 234
    fmt.Println(v14, v4)
    for fi15 := 0; fi15 < 20; fi15++ {
        v2 -= (v2 & v2)
        switch (("foo bar" + "hello") + ("" + ("\n\t" + (("a" + "hello") + "")))) {
        case "a":
            v16 := GenS2{F0: v3, F1: ""}
            _ = v16
        case "hello":
            switch "foo bar" {
            case "hello":
                if (42 <= 1) {
                    v14 = 1
                    v17 := GenS1{F0: (v3 ^ v3), F1: ""}
                    _ = v17
                } else {
                    fmt.Println(v4, v3, v4)
                    v18 := GenS3{F0: genFn0(-153, -32768, "foo bar"), F1: fi15}
                    v3 = (v3 >> (((v2 >> ((v2) & 63))) & 63))
                    _ = v18
                }
            case "\n\t":
                fmt.Println(v3, v4, v2)
            default:
                var v19 uint8 = 255
                _ = v19
            }
        case "":
            var v20 int32 = 486
            _ = v20
        default:
            if false {
                var v21 float32 = v4
                _ = v21
            } else {
                fmt.Println(v3, fi15)
                v22 := []int{42, (fi15 / ((-1) | 1)), (-(fi15))}
                _ = v22
            }
        }
        _ = fi15
    }
    v23 := []int{-52, 127}
    fmt.Println(v4, v3, v2)
    var v24 uint8 = 2
    v14 += (v14 | (v14 << ((65535) & 63)))
    for fi25 := 0; fi25 < 6; fi25++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        if ((!("\n\t" != "\n\t")) || true) {
            v27 := GenS1{F0: genFn0(-9223372036854775808, 0, "\n\t"), F1: ((((("\n\t" + (("a" + "") + "a")) + ("\n\t" + "a")) + "") + ("" + "\n\t")) + "a")}
            v28 := GenS3{F0: genFn0(1, 0, "a"), F1: fi26}
            _ = v27
            _ = v28
        }
        v29 := GenS2{F0: v3, F1: (("\n\t" + "\n\t") + "foo bar")}
        if true {
            v30 := GenS3{F0: genFn0(1, 416, ("a" + ("\n\t" + ""))), F1: (fi26 - (fi26 ^ (-(fi26))))}
            _ = v30
        }
        _ = fi26
        _ = v29
    }
    var v31 string = ("hello" + "\n\t")
    switch 251 {
    case 32767:
        fmt.Println(v14)
    default:
        if false {
            _ = len(v1)
        }
    }
    v5[32767] = 246
    if (9223372036854775807 > 42) {
        v32 := GenS4{F0: 127}
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 int32 = -1
        v3 += 32767
        for fi35 := 0; fi35 < 10; fi35++ {
            v1[v31] = fi35
            fmt.Println(v3, v3, v2)
            var v36 uint64 = uint64(v2)
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
