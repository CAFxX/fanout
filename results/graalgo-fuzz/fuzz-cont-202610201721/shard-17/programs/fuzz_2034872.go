package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 float32 = 3.14159
    v2 := []string{"", "\n\t"}
    var v3 uint32 = 2
    v1 = -2.25
    var v4 string = ""
    for fi5 := 0; fi5 < 1; fi5++ {
        v1 += v1
        var v6 uint = 18446744073709551615
        if ((float32(0.0) <= v1) || false) {
            var v7 int16 = -32768
            v8 := GenS1{F0: (fi5 >> ((v6) & 63))}
            v9 := GenS1{F0: 32768}
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: fi5}
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    _ = len(v2)
    switch 128 {
    case 912:
        v3 -= 2147483648
    case 32767:
        v11 := GenS1{F0: -1}
        _ = v11
    case 371:
        v4 = v4
    default:
        v12 := GenS1{F0: int(v3)}
        _ = v12
    }
    v13 := GenS2{F0: 128}
    switch v4 {
    case "hello":
        switch 4294967295 {
        case 748:
            fmt.Println(v3, v1)
        case 377:
            for fi14 := 0; fi14 < 5; fi14++ {
                v1 = 3.14159
                fmt.Println(v1)
                _ = fi14
            }
        case 1:
            if true {
                var v15 uint = 65535
                switch (("hello" + ((v4 + (v4 + "\n\t")) + "\n\t")) + v4) {
                case "a":
                    var v16 int32 = -2147483648
                    _ = v16
                case "\n\t":
                    var v17 uint32 = (v3 - v3)
                    _ = v17
                case "foo bar":
                    for fi18 := 0; fi18 < 6; fi18++ {
                        fmt.Println(v1, v3, v3)
                        var v19 uint8 = 12
                        _ = fi18
                        _ = v19
                    }
                default:
                    var v20 bool = (true || ((-2.25 - 0.0) < (1.5 * 0.1)))
                    _ = v20
                }
                _ = v15
            }
        }
    default:
        v21 := GenS2{F0: 128}
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if (40 < 1) {
            var v23 uint = 18446744073709551615
            _ = v23
        } else {
            var v24 int = fi22
            _ = v24
        }
        v25 := GenS2{F0: 0}
        var v26 int32 = 1
        _ = fi22
        _ = v25
        _ = v26
    }
    v2 = append(v2, ((v4 + "") + (v4 + (v4 + "a"))))
    var v27 int = 42
    if true {
        var v28 int16 = 0
        _ = v28
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v3 -= v3
        fmt.Println(v1, v1)
    }
    v2 = append(v2, v4)
    v29 := GenS3{F0: 579, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
