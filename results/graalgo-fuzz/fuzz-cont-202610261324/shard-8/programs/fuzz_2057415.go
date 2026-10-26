package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint16 = 492
    v2 := []uint32{0, 65535, 2, 2}
    v3 := make(map[string]uint32)
    var v4 float32 = 3.14159
    v5 := GenS1{F0: -731}
    for fi6 := 0; fi6 < 4; fi6++ {
        v1 = 32767
        if true {
            v7 := GenS1{F0: 0}
            for fi8 := 0; fi8 < 9; fi8++ {
                fmt.Println(fi6)
                _ = fi8
            }
            _ = v7
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint8 = 255
                _ = fi9
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            switch "hello" {
            case "a":
                switch fi6 {
                case -207:
                    if len(v2) > 0 {
                        v2[0] = 395
                    }
                case 346:
                    var v12 uint16 = v1
                    _ = v12
                }
            case "":
                var v13 int32 = 0
                _ = v13
            case "hello":
                var v14 uint16 = (v1 | (^(v1)))
                _ = v14
            }
            _ = fi11
        }
        _ = fi6
    }
    v3["a"] = 2147483647
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if (!(false || false)) {
            if true {
                fmt.Println(v4)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "hello":
        v15 := GenS2{F0: v1}
        _ = v15
    }
    var v16 uint16 = v1
    v16 = 932
    v17 := v3["hello"]
    var v18 int = 32768
    switch "" {
    case "\n\t":
        if false {
            var v19 uint32 = 2
            _ = v19
        }
    default:
        var v20 int32 = -2147483648
        _ = v20
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
