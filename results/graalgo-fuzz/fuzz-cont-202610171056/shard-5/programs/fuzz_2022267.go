package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
func genFn0() int16 {
    var v6 uint = 0
    if (!(0 <= -9223372036854775808)) {
        switch ("\n\t" + ("foo bar" + "")) {
        case "":
            var v7 uint64 = 2
            _ = v7
        }
        if (61 <= -128) {
            var v8 int32 = -2147483648
            _ = v8
        } else {
            if false {
                var v9 uint32 = 65535
                var v10 uint = (^(v6))
                _ = v10
                _ = v9
            }
        }
    }
    var v11 int32 = 2147483647
    v11 = 42
    fmt.Println(v11)
    _ = v11
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{230, 42}
    var v2 uint = 4294967295
    var v3 int32 = -625
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint8 = 238
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v2, fi14, v2)
        _ = fi14
    }
    v3 -= v3
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint16 = 42
        _ = fi15
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v17 uint64 = 42
        if (517 >= -1) {
            if (true || (100 != -128)) {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v2 += (v2 % ((v2) | 1))
                    for fi19 := 0; fi19 < 11; fi19++ {
                        if (100 != 42) {
                            var v20 int = fi19
                            fmt.Println(v3, v2, v20)
                            v21 := GenS1{F0: 235}
                            _ = v20
                            _ = v21
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v22 uint32 = 2147483648
                            _ = v22
                        case "a":
                            switch ("\n\t" + ("hello" + (("\n\t" + "hello") + "hello"))) {
                            case "\n\t":
                                v17 -= (v17 & 1)
                            case "foo bar":
                                fmt.Println(v2, v17)
                            case "":
                                v2 -= (v2 ^ (v2 << ((v2) & 63)))
                            }
                        case "\n\t":
                            v3 = v3
                        default:
                            v2 = 1
                            var v23 int8 = 127
                            if false {
                                var v24 uint = (v2 / ((437) | 1))
                                var v25 bool = (3.14159 >= -2.25)
                                _ = v24
                                _ = v25
                            }
                            _ = v23
                        }
                        _ = fi19
                    }
                    _ = fi18
                }
            }
            var v26 int32 = (v3 / ((v3) | 1))
            if len(v1) > 0 {
                v1[0] = 919
            }
            _ = v26
        }
        _ = v17
    }
    switch (("" + "hello") + "foo bar") {
    case "hello":
        fmt.Println(v3)
    case "a":
        fmt.Println(v2)
    case "foo bar":
        var v27 string = "hello"
        _ = v27
    }
    v28 := GenS1{F0: 241}
    if (576 == 65535) {
        switch "foo bar" {
        case "foo bar":
            var v29 bool = (float64(((0.1 * (0.0 / 0.1)) * (0.0 / -2.25))) != float64(-2.25))
            _ = v29
        }
        var v30 int16 = 1
        var v31 int32 = (v3 << ((v2) & 63))
        _ = v30
        _ = v31
    }
    if false {
        var v32 string = "a"
        fmt.Println(v32, v3, v2)
        _ = v32
    }
    var v33 int16 = 1000
    fmt.Println(v33)
    fmt.Println(v2, v2)
    v34 := GenS2{F0: 128, F1: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
