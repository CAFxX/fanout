package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    var v4 uint32 = 4294967295
    switch 32767 {
    case 876:
        v4 = v4
    case -755:
        v5 := GenS1{F0: 32767, F1: 4294967295}
        _ = v5
    case 32767:
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                var v7 uint16 = 846
                for fi8 := 0; fi8 < 10; fi8++ {
                    fmt.Println(v4)
                    _ = fi8
                }
                var v9 uint = 255
                _ = v7
                _ = v9
            }
            var v10 bool = (255 < 128)
            _ = fi6
            _ = v10
        }
    }
    fmt.Println(v4)
    switch "\n\t" {
    case "a":
        switch -479 {
        case -556:
            if true {
                v4 = v4
                if true {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 int32 = -638
                        _ = fi11
                        _ = v12
                    }
                    var v13 uint16 = 32767
                    var v14 uint8 = 127
                    _ = v13
                    _ = v14
                } else {
                    var v15 int32 = -247
                    fmt.Println(v15, v15)
                    if false {
                        if (-463 >= 1) {
                            delete(v2, (v4 | (^(v4))))
                            fmt.Println(v4, v4)
                            var v16 uint32 = v4
                            _ = v16
                        }
                    }
                    _ = v15
                }
            } else {
                fmt.Println(v4)
                if true {
                    v4 = 4294967295
                }
                var v17 uint32 = 2147483648
                _ = v17
            }
        case 0:
            v18 := GenS1{F0: 311, F1: 753}
            _ = v18
        case 32767:
            var v19 int64 = 0
            _ = v19
        }
    case "foo bar":
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(fi20)
            var v21 uint16 = 42
            _ = fi20
            _ = v21
        }
    }
    v4 -= v4
    v4 = (v4 + (v4 + (v4 ^ v4)))
    v4 = 0
    var v22 int16 = -32768
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
