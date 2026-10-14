package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = (1.5 - -2.25)
    v2 := []string{}
    v3 := make(map[string]string)
    var v4 uint16 = 2
    v5 := GenS1{F0: 32}
    if false {
        v4 = 42
    }
    if true {
        var v6 uint64 = 42
        _ = v6
    }
    if false {
        v4 += 32767
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v7 uint = 4294967295
        _ = v7
    }
    switch (("foo bar" + "hello") + "a") {
    case "":
        switch "foo bar" {
        case "hello":
            if false {
                for fi8 := 0; fi8 < 4; fi8++ {
                    _ = len(v3)
                    if false {
                        v1 = v1
                    }
                    _ = fi8
                }
            } else {
                switch ("foo bar" + "\n\t") {
                case "a":
                    for fi9 := 0; fi9 < 16; fi9++ {
                        for fi10 := 0; fi10 < 4; fi10++ {
                            v4 += 32767
                            _ = len(v3)
                            var v11 int16 = -920
                            fmt.Println(fi9, v4)
                            v4 = (v4 << ((4294967295) & 63))
                            switch 127 {
                            case 0:
                                v12 := GenS2{F0: 2}
                                _ = v12
                            case -339:
                                var v13 bool = (-2.25 < v1)
                                _ = v13
                            case 2147483647:
                                v4 = (v4 ^ (v4 * (v4 - v4)))
                            }
                            _ = fi10
                            _ = v11
                        }
                        _ = fi9
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "\n\t":
                        var v14 uint32 = 1
                        _ = v14
                    }
                }
            }
        case "foo bar":
            v4 -= v4
        default:
            if false {
                if (!true) {
                    v1 = 0.0
                    if true {
                        v3[""] = "a"
                    } else {
                        _ = len(v3)
                        if true {
                            fmt.Println(v4)
                        }
                        v15 := v3[""]
                        _ = v15
                    }
                }
                fmt.Println(v4)
                var v16 uint8 = 42
                _ = v16
            }
        }
    default:
        var v17 uint64 = 372
        _ = v17
    }
    v4 = (v4 / ((42) | 1))
    v4 = v4
    var v18 uint32 = 119
    var v19 int64 = -9223372036854775808
    if (1 <= 0) {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint = 0
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 bool = true
                fmt.Println(v23)
                v24 := GenS2{F0: 0}
                _ = fi22
                _ = v23
                _ = v24
            }
            v21 += v21
            _ = fi20
            _ = v21
        }
        if len(v2) > 0 {
            v2[0] = ("a" + ("foo bar" + ((("" + "\n\t") + "a") + "hello")))
        }
    }
    var v25 uint64 = 435
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
