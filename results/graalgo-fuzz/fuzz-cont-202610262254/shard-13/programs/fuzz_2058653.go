package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{932, 2147483647, 0}
    var v3 bool = false
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            v3 = true
            _ = len(v2)
            if v3 {
                v6 := GenS1{F0: 4294967295}
                if v3 {
                    v3 = (18446744073709551615 == 255)
                    switch "" {
                    case "foo bar":
                        v7 := GenS1{F0: 4294967295}
                        _ = v7
                    default:
                        fmt.Println(fi4, fi5, fi5)
                    }
                    var v8 uint32 = 2147483647
                    _ = v8
                }
                switch 32768 {
                case 32767:
                    var v9 bool = (42 == 42)
                    _ = v9
                case 159:
                    switch fi4 {
                    case 93:
                        var v10 int = fi5
                        _ = v10
                    }
                case 128:
                    fmt.Println(v1, v1)
                default:
                    var v11 int8 = 0
                    _ = v11
                }
                _ = v6
            }
            _ = fi5
        }
        _ = fi4
    }
    v1 += v1
    var v12 int8 = 100
    v12 -= v12
    var v13 uint = 1
    v3 = v3
    v14 := GenS1{F0: 484}
    if v3 {
        var v15 uint32 = 42
        v3 = (v3 && v3)
        if (-32768 >= -32768) {
            for fi16 := 0; fi16 < 11; fi16++ {
                _ = len(v2)
                fmt.Println(v1, v3)
                var v17 uint16 = 32767
                _ = fi16
                _ = v17
            }
            v15 = v15
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(fi18)
                for fi19 := 0; fi19 < 3; fi19++ {
                    fmt.Println(v15, v3, v12)
                    if v3 {
                        var v20 string = "a"
                        v21 := GenS2{F0: v20}
                        var v22 uint = v13
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v23 int16 = 0
                    _ = fi19
                    _ = v23
                }
                _ = fi18
            }
        } else {
            v24 := GenS2{F0: "a"}
            var v25 uint64 = uint64(v13)
            _ = v24
            _ = v25
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
