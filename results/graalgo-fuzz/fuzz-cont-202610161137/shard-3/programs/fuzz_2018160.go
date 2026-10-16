package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float64 = (-2.25 / 1.0)
    var v2 int = 1
    var v3 float32 = 1.5
    v4 := []int{-1, 128, (v2 << ((4294967295) & 63))}
    v5 := GenS1{F0: 332}
    var v6 string = ((("a" + "hello") + ("" + "\n\t")) + "a")
    for fi7 := 0; fi7 < 5; fi7++ {
        v2 += fi7
        v6 = v6
        _ = fi7
    }
    v6 = "foo bar"
    v1 += (v1 + v1)
    switch v2 {
    case 32768:
        switch (^(v2)) {
        case 4294967295:
            if (412 <= 0) {
                var v8 float32 = 3.14159
                _ = len(v4)
                switch v2 {
                case 2147483647:
                    for fi9 := 0; fi9 < 14; fi9++ {
                        fmt.Println(fi9, v1, fi9)
                        for fi10 := 0; fi10 < 16; fi10++ {
                            fmt.Println(v2, fi9)
                            _ = fi10
                        }
                        var v11 uint = 2
                        _ = fi9
                        _ = v11
                    }
                case 32768:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 uint32 = 0
                        v4 = append(v4, v2)
                        v14 := GenS1{F0: 0}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 51:
                    var v15 uint16 = 0
                    _ = v15
                }
                _ = v8
            } else {
                var v16 uint = 255
                var v17 uint16 = 2
                var v18 uint = 75
                _ = v16
                _ = v17
                _ = v18
            }
        case 127:
            for fi19 := 0; fi19 < 1; fi19++ {
                v6 = (v6 + v6)
                var v20 int8 = 127
                if false {
                    switch v6 {
                    case "":
                        var v21 uint64 = 18446744073709551615
                        _ = v21
                    case "foo bar":
                        var v22 uint8 = 255
                        _ = v22
                    }
                }
                _ = fi19
                _ = v20
            }
        case -2147483648:
            for fi23 := 0; fi23 < 7; fi23++ {
                if (0 >= 42) {
                    var v24 int16 = 173
                    _ = v24
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v2 = (fi23 + fi23)
                }
                switch ((v6 + (((((v6 + v6) + ("foo bar" + (("\n\t" + "a") + v6))) + "hello") + v6) + v6)) + v6) {
                case "hello":
                    v3 = ((1e10 * (float32(v1) - v3)) / v3)
                case "foo bar":
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v6 = v6
                        _ = fi25
                    }
                }
                var v26 int = (fi23 * v2)
                _ = fi23
                _ = v26
            }
        default:
            switch v6 {
            case "a":
                v3 -= 1e10
            case "foo bar":
                v1 = (-2.25 * v1)
            case "\n\t":
                _ = len(v4)
            }
        }
    case 32767:
        fmt.Println(v2, v2)
    case 282:
        for fi27 := 0; fi27 < 3; fi27++ {
            switch v6 {
            case "\n\t":
                v3 -= (v3 * (v3 + ((v3 + v3) + 1.5)))
            case "foo bar":
                fmt.Println(v3, v6)
            }
            _ = fi27
        }
    default:
        var v28 float32 = (-2.25 / v3)
        _ = v28
    }
    var v29 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
