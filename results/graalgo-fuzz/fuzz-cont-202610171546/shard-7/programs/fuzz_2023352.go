package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{}
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    v5 := []string{v2, v2, ("" + v2), v2}
    v6 := []int{128, 0, 127, -808}
    var v7 string = (v2 + "\n\t")
    if true {
        var v8 uint64 = 1
        v8 += (v8 * v8)
        v3[347] = 2
        _ = v8
    } else {
        fmt.Println(v7, v2)
        if true {
            fmt.Println(v7, v7)
            switch -9223372036854775808 {
            case -1:
                for fi9 := 0; fi9 < 17; fi9++ {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    for fi10 := 0; fi10 < 17; fi10++ {
                        switch "foo bar" {
                        case "foo bar":
                            v7 = v2
                        case "hello":
                            v11 := GenS1{F0: 742}
                            _ = v11
                        case "":
                            if false {
                                var v12 uint8 = 0
                                v13 := v3[544]
                                _ = v12
                                _ = v13
                            }
                        }
                        v14 := GenS2{F0: (fi10 * 9223372036854775807)}
                        _ = len(v5)
                        _ = fi10
                        _ = v14
                    }
                    _ = fi9
                }
            case 42:
                var v15 uint64 = 1
                _ = v15
            default:
                switch v2 {
                case "a":
                    v16 := GenS2{F0: 127}
                    _ = v16
                default:
                    v2 = (v7 + v2)
                }
            }
            v2 = ((v2 + "a") + "\n\t")
        } else {
            fmt.Println(v7, v2)
        }
    }
    var v17 uint32 = 2147483647
    v7 = "\n\t"
    var v18 uint = 255
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint = v18
        for fi21 := 0; fi21 < 20; fi21++ {
            if true {
                v22 := GenS2{F0: (fi19 >> ((65535) & 63))}
                _ = v22
            } else {
                var v23 int32 = 42
                _ = v23
            }
            _ = fi21
        }
        v24 := v4["hello"]
        _ = fi19
        _ = v20
        _ = v24
    }
    var v25 int8 = -128
    for fi26 := 0; fi26 < 15; fi26++ {
        if true {
            if false {
                fmt.Println(fi26, v17, v18)
            } else {
                if len(v1) > 0 {
                    v1[0] = v2
                }
            }
            var v27 uint = v18
            v17 -= 2
            _ = v27
        } else {
            var v28 bool = false
            _ = v28
        }
        _ = fi26
    }
    var v29 uint16 = 777
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
