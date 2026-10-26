package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int; F1 int8; F2 uint64 }
func genFn0(p0 int8, p1 uint64, p2 uint8) uint16 {
    switch ("a" + "foo bar") {
    case "\n\t":
        v5 := GenS1{F0: true, F1: 1, F2: 648}
        _ = v5
    }
    var v6 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 float64 = float64(0.1)
    v2 := []int{}
    var v3 uint64 = 2
    v4 := make(map[int]bool)
    if true {
        v7 := GenS2{F0: 389, F1: 100, F2: 793}
        if false {
            v8 := GenS3{F0: 0, F1: 2147483647}
            _ = v8
        } else {
            _ = len(v2)
        }
        _ = v7
    } else {
        fmt.Println(v3)
    }
    v3 = v3
    if (!((128 >= 794) || (true && true))) {
        v4[42] = (-1 >= 574)
        fmt.Println(v3, v3)
        v2 = append(v2, 4294967295)
    }
    var v9 float32 = 3.14159
    if true {
        if (!true) {
            var v10 uint16 = 2
            _ = v10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint = 509
            _ = fi11
            _ = v12
        }
    } else {
        switch "\n\t" {
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "a":
            if (!false) {
                _ = len(v2)
                v9 = v9
                var v13 uint32 = 1
                _ = v13
            }
        }
        _ = len(v2)
    }
    switch "a" {
    case "\n\t":
        var v14 int64 = -9223372036854775808
        _ = v14
    case "hello":
        v1 = ((v1 * v1) / 0.1)
    default:
        v15 := GenS4{F0: 128, F1: 100, F2: (^(v3))}
        _ = v15
    }
    delete(v4, int(v3))
    var v16 string = (("foo bar" + ("a" + "a")) + ("\n\t" + ("a" + ((("\n\t" + ("" + "foo bar")) + ("\n\t" + "\n\t")) + ""))))
    v16 = ((("a" + (v16 + ((("a" + "\n\t") + v16) + "\n\t"))) + (v16 + v16)) + ((v16 + ("hello" + "foo bar")) + v16))
    v9 -= v9
    v3 = (v3 ^ 42)
    switch (v16 + (v16 + "\n\t")) {
    case "hello":
        v3 = v3
    case "a":
        v9 = -2.25
    }
    v3 -= (v3 >> ((4294967295) & 63))
    v17 := GenS4{F0: 32768, F1: 1, F2: (^(v3))}
    for fi18 := 0; fi18 < 13; fi18++ {
        if ("hello" == ((v16 + "a") + v16)) {
            fmt.Println(fi18)
        }
        switch 128 {
        case 127:
            v4[(-(fi18))] = false
        }
        _ = fi18
    }
    if false {
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS4{F0: 2147483647, F1: 1, F2: v3}
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int16 = -1
                v16 = "a"
                v23 := GenS1{F0: true, F1: 2, F2: 0}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 string = (v16 + v16)
            _ = fi19
            _ = v20
            _ = v24
        }
        var v25 uint16 = 1
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        delete(v4, (-(fi26)))
        var v27 uint16 = 42
        var v28 bool = false
        _ = fi26
        _ = v27
        _ = v28
    }
    switch (v16 + "") {
    case "\n\t":
        var v29 uint64 = v3
        _ = v29
    case "foo bar":
        fmt.Println(v9, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
