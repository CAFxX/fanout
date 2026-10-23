package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 int64; F2 uint }

func main() {
    v1 := []string{("" + "foo bar"), ((("foo bar" + "foo bar") + "a") + ("" + "a")), ("" + "a"), "\n\t"}
    v2 := make(map[string]uint32)
    var v3 bool = (false && true)
    var v4 int64 = -9223372036854775808
    v1 = append(v1, "\n\t")
    if true {
        v4 = 9223372036854775807
        if (v3 || v3) {
            switch ("" + "a") {
            case "":
                fmt.Println(v3, v4)
            case "foo bar":
                for fi5 := 0; fi5 < 7; fi5++ {
                    _ = len(v1)
                    delete(v2, ("a" + "hello"))
                    _ = fi5
                }
            }
            var v6 uint16 = 2
            delete(v2, "hello")
            _ = v6
        }
        var v7 uint16 = 65535
        _ = v7
    }
    v3 = v3
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint16 = 255
        v3 = v3
        v3 = v3
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: 255, F1: 18446744073709551615}
            v4 += -1
            switch ("" + "hello") {
            case "hello":
                v13 := GenS1{F0: 2, F1: 269}
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        if (1e10 != 0.1) {
            var v14 int16 = 1
            v14 = 1000
            _ = v14
        } else {
            v15 := GenS2{F0: "foo bar", F1: v4, F2: 42}
            var v16 uint64 = 349
            _ = v15
            _ = v16
        }
        var v17 string = ""
        _ = fi10
        _ = v17
    }
    fmt.Println(v4, v3)
    v4 = (v4 >> ((42) & 63))
    fmt.Println(v3, v3)
    fmt.Println(v3, v3)
    v3 = v3
    v3 = (!false)
    delete(v2, "foo bar")
    v4 = v4
    if (!(-1 >= -41)) {
        var v18 int32 = -1
        switch 2147483647 {
        case -866:
            var v19 int32 = v18
            _ = v19
        }
        _ = v18
    } else {
        fmt.Println(v3)
        var v20 int16 = -1
        if v3 {
            fmt.Println(v3)
            switch -604 {
            case -2147483648:
                v4 = -9223372036854775808
            default:
                switch "hello" {
                case "":
                    switch 99 {
                    case 32768:
                        fmt.Println(v20, v20)
                    case 4294967295:
                        fmt.Println(v4, v20, v20)
                    default:
                        var v21 uint = 42
                        _ = v21
                    }
                case "a":
                    if false {
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v3 = v3
                            _ = fi22
                        }
                        v3 = true
                    } else {
                        var v23 uint32 = 661
                        if ((0 >= -373) && v3) {
                            var v24 bool = false
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 uint8 = 2
                                _ = fi25
                                _ = v26
                            }
                            var v27 float64 = -2.25
                            _ = v24
                            _ = v27
                        }
                        _ = v23
                    }
                }
            }
            var v28 int8 = -1
            _ = v28
        }
        _ = v20
    }
    switch ((("foo bar" + ("a" + ("foo bar" + "hello"))) + ("a" + ("\n\t" + "\n\t"))) + ("a" + "a")) {
    case "\n\t":
        fmt.Println(v4, v4)
    case "":
        fmt.Println(v3)
    }
    v29 := GenS1{F0: 32767, F1: 0}
    var v30 int16 = -415
    v4 += (-(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
