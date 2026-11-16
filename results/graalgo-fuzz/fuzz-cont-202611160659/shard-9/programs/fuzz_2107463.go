package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint) uint8 {
    var v5 uint64 = 42
    var v6 int8 = -94
    _ = p0
    _ = v5
    _ = v6
    return 88
}
func genFn1(p0 int, p1 int8) int16 {
    var v7 int32 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 32767}
        var v10 uint32 = 579
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = -800
    if (!(-9223372036854775808 <= 564)) {
        v12 := []int{p0, p0, p0}
        fmt.Println(p1, v11, p1)
        v13 := GenS1{F0: -32768}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return v11
}

func main() {
    var v1 string = (((("foo bar" + "a") + (("" + (("hello" + "a") + "\n\t")) + ("\n\t" + ("foo bar" + "hello")))) + ((("a" + "a") + "") + "a")) + ("" + "\n\t"))
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []string{v1, (((("a" + "\n\t") + (v1 + (("\n\t" + "foo bar") + (("foo bar" + v1) + ("" + v1))))) + v1) + ("hello" + v1)), "a", v1}
    if (true && (false || ((false && false) || false))) {
        v14 := GenS1{F0: -528}
        v2[32768] = 762
        _ = v14
    } else {
        v1 = ""
        var v15 int32 = 2147483647
        _ = v15
    }
    if true {
        fmt.Println(v1)
        var v16 bool = (((-2147483648 >= -1) && true) && (552 >= 65535))
        _ = v16
    } else {
        fmt.Println(v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v17 := GenS2{F0: 0}
        _ = v17
    }
    switch "foo bar" {
    case "\n\t":
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v1 = v1
                if len(v4) > 0 {
                    v4[0] = v1
                }
                _ = fi19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS2{F0: 727}
                switch (v1 + ((("hello" + "\n\t") + "\n\t") + ("foo bar" + ("foo bar" + v1)))) {
                case "":
                    _ = len(v4)
                case "\n\t":
                    v22 := GenS1{F0: 1000}
                    _ = v22
                case "a":
                    if false {
                        var v23 uint16 = 255
                        v1 = v1
                        _ = v23
                    } else {
                        var v24 uint16 = 267
                        _ = v24
                    }
                }
                v3[49] = 218
                _ = fi20
                _ = v21
            }
            switch "a" {
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = v1
                }
            }
            _ = fi18
        }
    case "foo bar":
        var v25 int64 = 42
        _ = v25
    case "":
        v1 = "hello"
    }
    fmt.Println(v1)
    switch v1 {
    case "foo bar":
        v1 = (v1 + v1)
    case "hello":
        if ("a" >= v1) {
            fmt.Println(v1)
            if len(v4) > 0 {
                v4[0] = v1
            }
        }
    case "":
        var v26 int = 42
        _ = v26
    }
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 1; fi28++ {
        v1 = v1
        v1 = v1
        var v29 uint8 = 29
        _ = fi28
        _ = v29
    }
    v27 -= (^(v27))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
