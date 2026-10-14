package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint }
func genFn0(p0 uint8) uint64 {
    p0 += p0
    switch ("hello" + ((("a" + "hello") + ("a" + ("" + "\n\t"))) + "foo bar")) {
    case "foo bar":
        p0 += p0
    case "hello":
        fmt.Println(p0)
    default:
        fmt.Println(p0)
    }
    switch "hello" {
    case "":
        p0 = (p0 - 127)
    case "foo bar":
        var v7 uint = 2
        _ = v7
    default:
        var v8 uint8 = 0
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(p0)
        if true {
            fmt.Println(fi9, fi9)
        } else {
            var v10 bool = true
            switch ("\n\t" + "a") {
            case "foo bar":
                v11 := GenS1{F0: 100, F1: 42, F2: 1}
                _ = v11
            case "hello":
                var v12 int16 = 42
                _ = v12
            }
            _ = v10
        }
        _ = fi9
    }
    var v13 uint64 = 42
    _ = p0
    _ = v13
    return 2
}
func genFn1(p0 string, p1 int16) uint64 {
    var v14 uint64 = 1
    v15 := GenS1{F0: 80, F1: 114, F2: 18446744073709551615}
    p1 = p1
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int8 = 0
        _ = fi16
        _ = v17
    }
    v18 := []int{}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v18
    return (v14 / ((v14) | 1))
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := []string{"a", "a", "a", "hello"}
    var v6 uint = 42
    var v19 int8 = 42
    if (72 > 127) {
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint64 = v2
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 bool = true
                if v23 {
                    if (!(-1 == -1)) {
                        _ = len(v4)
                    } else {
                        var v24 uint16 = 65535
                        _ = v24
                    }
                    if v23 {
                        var v25 uint = (^(v6))
                        _ = v25
                    }
                }
                var v26 uint8 = 20
                _ = fi22
                _ = v23
                _ = v26
            }
            fmt.Println(v21, v1)
            _ = fi20
            _ = v21
        }
        v1 = v1
    }
    var v27 int64 = -636
    for fi28 := 0; fi28 < 9; fi28++ {
        switch (("hello" + "a") + "a") {
        case "foo bar":
            switch (((("hello" + "hello") + "a") + ("hello" + ("a" + ((("foo bar" + "foo bar") + "hello") + "\n\t")))) + ("\n\t" + ("\n\t" + "\n\t"))) {
            case "foo bar":
                var v29 uint64 = (v2 & (^(v2)))
                _ = v29
            }
        }
        _ = fi28
    }
    if false {
        switch 4294967295 {
        case 42:
            if len(v5) > 0 {
                v5[0] = "a"
            }
        case 25:
            var v30 uint64 = (v2 * v2)
            _ = v30
        }
    } else {
        var v31 bool = (("hello" <= ("foo bar" + "")) || ((v1 & (v1 & v1)) < (v1 * 645)))
        _ = v31
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        delete(v4, 932)
    case "a":
        var v32 int8 = (v19 << ((530) & 63))
        _ = v32
    case "\n\t":
        fmt.Println(v2)
    default:
        for fi33 := 0; fi33 < 20; fi33++ {
            switch "\n\t" {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = ((("\n\t" + "foo bar") + "foo bar") + "foo bar")
                }
            }
            _ = fi33
        }
    }
    v1 = v1
    delete(v4, -563)
    v1 = 0
    var v34 int32 = 1
    var v35 int8 = (v19 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
