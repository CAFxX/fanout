package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
func genFn0() uint16 {
    var v5 int16 = 0
    var v6 bool = false
    v6 = v6
    _ = v5
    _ = v6
    return 87
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 1.5
    var v3 uint64 = 18446744073709551615
    v4 := []string{"\n\t", ""}
    if len(v4) > 0 {
        v4[0] = (("\n\t" + (("a" + "hello") + "")) + "foo bar")
    }
    var v7 int16 = int16(v3)
    switch "foo bar" {
    case "foo bar":
        for fi8 := 0; fi8 < 12; fi8++ {
            v1 -= (v1 >> ((18446744073709551615) & 63))
            fmt.Println(v3, v3, fi8)
            _ = fi8
        }
    }
    v9 := GenS1{F0: "a"}
    for fi10 := 0; fi10 < 16; fi10++ {
        switch fi10 {
        case 32768:
            var v11 int = fi10
            _ = v11
        case 42:
            for fi12 := 0; fi12 < 4; fi12++ {
                if false {
                    v13 := GenS1{F0: "hello"}
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "hello")
                    }
                    _ = v13
                } else {
                    var v14 int = (fi12 / ((fi12) | 1))
                    fmt.Println(v14, fi12)
                    v2 -= (v2 * (float64(0.0) / (v2 / v2)))
                    _ = v14
                }
                fmt.Println(v2, v7)
                fmt.Println(v1)
                _ = fi12
            }
        case 798:
            var v15 int16 = (v7 - (v7 | (v7 % ((-1) | 1))))
            _ = v15
        default:
            if (!(true && ((2 <= 113) && ("hello" <= ("hello" + ("" + "\n\t")))))) {
                fmt.Println(v2, v1, v3)
                fmt.Println(v1)
            }
        }
        var v16 uint64 = v3
        _ = fi10
        _ = v16
    }
    _ = len(v4)
    v17 := GenS1{F0: "hello"}
    if true {
        for fi18 := 0; fi18 < 1; fi18++ {
            switch 32768 {
            case 109:
                switch fi18 {
                case 128:
                    for fi19 := 0; fi19 < 20; fi19++ {
                        if false {
                            if false {
                                var v20 float32 = 3.14159
                                var v21 int32 = 42
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS2{F0: 199}
                            _ = v22
                        }
                        var v23 string = ""
                        _ = fi19
                        _ = v23
                    }
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 1; fi24++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch fi24 {
                        case 42:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi24
                    }
                case 60:
                    v1 = (v1 >> ((18446744073709551615) & 63))
                default:
                    var v25 int16 = v7
                    _ = v25
                }
            }
            _ = fi18
        }
        v3 = 2
    } else {
        v3 = (v3 << ((4294967295) & 63))
    }
    switch 240 {
    case 32768:
        v26 := GenS2{F0: 267}
        _ = v26
    default:
        v2 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
