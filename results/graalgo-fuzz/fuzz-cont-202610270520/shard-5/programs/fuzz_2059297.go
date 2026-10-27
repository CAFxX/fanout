package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }

func main() {
    var v1 bool = true
    var v2 uint64 = 0
    var v3 int16 = -32768
    var v4 int64 = 1
    var v5 int16 = v3
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = 0
        var v8 uint = 2
        var v9 uint16 = 32767
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v5 = 1
    v10 := GenS1{F0: 42, F1: "", F2: 41}
    v1 = (("a" + "hello") <= ("foo bar" + ""))
    for fi11 := 0; fi11 < 17; fi11++ {
        v2 = (v2 << ((uint(v2)) & 63))
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    v5 = v3
    v4 += 9223372036854775807
    switch "hello" {
    case "a":
        v13 := []int{9223372036854775807, 127, -841, 32767}
        _ = v13
    case "":
        fmt.Println(v4)
    }
    v4 = v4
    if (v1 || ((1e10 / 1.0) <= 1e10)) {
        v14 := []int{690, -2147483648, -72, 502}
        if v1 {
            var v15 float64 = 3.14159
            for fi16 := 0; fi16 < 3; fi16++ {
                fmt.Println(fi16, v15)
                _ = fi16
            }
            switch ("" + "foo bar") {
            case "":
                switch "\n\t" {
                case "hello":
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v3 += v3
                        _ = fi17
                    }
                case "a":
                    v1 = (835 < 42)
                case "":
                    for fi18 := 0; fi18 < 17; fi18++ {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            fmt.Println(fi19)
                            fmt.Println(v2)
                            switch (fi18 * (^(fi18))) {
                            case 32767:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            case 42:
                                v4 -= 1
                            case -381:
                                v15 = (0.1 * 0.0)
                            }
                            var v20 uint32 = 428
                            var v21 uint16 = 0
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        _ = fi18
                    }
                default:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v1 = v1
                        _ = fi22
                    }
                }
            case "a":
                v15 -= float64(0.1)
            case "hello":
                var v23 uint8 = 2
                _ = v23
            }
            _ = v15
        } else {
            var v24 int8 = 66
            if (963 <= 128) {
                var v25 int32 = 0
                var v26 uint16 = 2
                _ = v25
                _ = v26
            }
            _ = v24
        }
        v3 -= v5
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
