package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{0}
    var v3 bool = true
    v4 := make(map[string]uint32)
    v3 = true
    fmt.Println(v3)
    v3 = (1 < 0)
    fmt.Println(v3)
    for fi5 := 0; fi5 < 11; fi5++ {
        _ = len(v4)
        var v6 float64 = -2.25
        _ = fi5
        _ = v6
    }
    fmt.Println(v3)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 bool = v3
        var v9 int16 = 607
        fmt.Println(v8, fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "hello" {
        case "":
            fmt.Println(v3, fi10)
        default:
            var v11 uint = 2
            _ = v11
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            if len(v2) > 0 {
                v2[0] = (fi12 >> ((42) & 63))
            }
            _ = len(v4)
            _ = fi12
        }
        _ = fi10
    }
    if true {
        if v3 {
            v2 = append(v2, -1)
            var v13 uint16 = 935
            if (18446744073709551615 > 711) {
                if v3 {
                    v3 = (!((true || (255 <= 119)) && (!v3)))
                    switch "" {
                    case "foo bar":
                        switch ("hello" + ("a" + "foo bar")) {
                        case "foo bar":
                            v3 = false
                        case "\n\t":
                            fmt.Println(v3, v13)
                        }
                    case "":
                        v14 := GenS1{F0: -2147483648}
                        _ = v14
                    case "a":
                        for fi15 := 0; fi15 < 15; fi15++ {
                            for fi16 := 0; fi16 < 8; fi16++ {
                                v2 = append(v2, (fi16 >> ((214) & 63)))
                                _ = fi16
                            }
                            _ = fi15
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                    }
                } else {
                    fmt.Println(v3, v13)
                }
                fmt.Println(v13)
                v3 = (42 >= 490)
            } else {
                v3 = v3
                for fi17 := 0; fi17 < 9; fi17++ {
                    fmt.Println(fi17, v13, v3)
                    _ = fi17
                }
                switch "hello" {
                case "hello":
                    if v3 {
                        v18 := GenS1{F0: 0}
                        _ = v18
                    }
                case "a":
                    var v19 uint = 2
                    _ = v19
                default:
                    delete(v1, 65535)
                }
            }
            _ = v13
        } else {
            var v20 uint32 = 2
            delete(v4, ("a" + "hello"))
            _ = v20
        }
        var v21 int16 = 32767
        var v22 uint16 = 1
        _ = v21
        _ = v22
    } else {
        fmt.Println(v3)
        v23 := GenS2{F0: 1}
        _ = v23
    }
    if (-415 <= 218) {
        for fi24 := 0; fi24 < 14; fi24++ {
            if len(v2) > 0 {
                v2[0] = fi24
            }
            var v25 string = ""
            _ = fi24
            _ = v25
        }
        var v26 int = -2147483648
        v27 := GenS2{F0: (v26 ^ (v26 * v26))}
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
