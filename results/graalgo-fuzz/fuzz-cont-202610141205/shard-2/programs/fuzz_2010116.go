package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []uint32{}
    var v2 int64 = 9223372036854775807
    var v3 int16 = 32767
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 13; fi5++ {
        if (232 > 55) {
            var v6 uint16 = v4
            _ = v6
        }
        _ = fi5
    }
    var v7 float32 = float32(0.0)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = 1000
    v8 := GenS1{F0: false}
    v4 = (v4 % (((v4 << ((2) & 63))) | 1))
    if false {
        var v9 string = "\n\t"
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint32 = 0
            var v12 uint64 = 42
            switch "" {
            case "hello":
                if (!true) {
                    if true {
                        var v13 int32 = -2147483648
                        switch 9223372036854775807 {
                        case 144:
                            v11 -= 542
                        }
                        _ = v13
                    }
                    fmt.Println(fi10)
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v2 = v2
    }
    switch "foo bar" {
    case "":
        var v14 int = -1
        _ = v14
    case "hello":
        v15 := GenS1{F0: true}
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
