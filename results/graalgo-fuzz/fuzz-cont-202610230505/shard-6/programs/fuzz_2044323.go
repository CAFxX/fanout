package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    v1 := []string{(("foo bar" + ("\n\t" + ("foo bar" + "hello"))) + "foo bar"), ("a" + ("hello" + "a")), "\n\t", "a"}
    v2 := []string{"a", "hello"}
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 17; fi4++ {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        if false {
            _ = len(v2)
        } else {
            _ = len(v1)
        }
        _ = fi4
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v5 uint32 = 38
    switch "hello" {
    case "":
        switch "foo bar" {
        case "a":
            fmt.Println(v5)
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            v5 = (^(v5))
        }
    case "foo bar":
        fmt.Println(v5)
    case "a":
        v6 := v3["\n\t"]
        _ = v6
    }
    switch -9223372036854775808 {
    case 32767:
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int = fi7
            switch "\n\t" {
            case "foo bar":
                _ = len(v3)
            case "hello":
                v9 := GenS1{F0: false}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    case -282:
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                v5 += (v5 + (v5 + 242))
                fmt.Println(fi10, fi10, fi10)
                var v12 uint8 = 174
                _ = fi11
                _ = v12
            }
            v1 = append(v1, "\n\t")
            _ = fi10
        }
    }
    var v13 int32 = 1
    v14 := GenS2{F0: 42, F1: v13}
    var v15 uint64 = 2
    fmt.Println(v5, v5)
    v1 = append(v1, "\n\t")
    if ((188 <= 9) || true) {
        switch (("a" + "") + (("hello" + (("\n\t" + "") + "hello")) + ("hello" + ((("a" + "hello") + (("hello" + "\n\t") + "\n\t")) + "hello")))) {
        case "\n\t":
            v16 := GenS3{F0: -670, F1: v13}
            _ = v16
        case "foo bar":
            switch "foo bar" {
            case "a":
                var v17 int8 = -1
                _ = v17
            }
        case "":
            if len(v2) > 0 {
                v2[0] = ("a" + "\n\t")
            }
        default:
            switch "\n\t" {
            case "hello":
                v5 += (^(v5))
            case "":
                v18 := GenS3{F0: -1, F1: (v13 - v13)}
                _ = v18
            }
        }
        var v19 uint16 = 444
        fmt.Println(v5, v15, v13)
        _ = v19
    } else {
        if (193 > v15) {
            v3["\n\t"] = 0
            fmt.Println(v13)
            var v20 int64 = 42
            _ = v20
        }
    }
    var v21 uint8 = 2
    v5 = (v5 - v5)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 string = "hello"
        for fi24 := 0; fi24 < 8; fi24++ {
            v15 -= (v15 * (v15 << ((870) & 63)))
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    if true {
        var v25 uint64 = 2
        _ = v25
    } else {
        switch 4294967295 {
        case 32767:
            switch "\n\t" {
            case "a":
                v5 = (v5 + 555)
            case "":
                fmt.Println(v13, v13)
            }
        case -75:
            v15 = 4294967295
        }
    }
    if true {
        v1 = append(v1, (("a" + "hello") + "\n\t"))
    }
    var v26 uint16 = uint16(v5)
    fmt.Println(v13, v21, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v5
}
