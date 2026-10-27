package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint) uint {
    fmt.Println(p0)
    var v4 int16 = 32767
    for fi5 := 0; fi5 < 14; fi5++ {
        v4 = v4
        _ = fi5
    }
    _ = p0
    _ = v4
    return 501
}

func main() {
    var v1 uint32 = 47
    v2 := make(map[string]uint32)
    var v3 string = ("foo bar" + "a")
    var v6 uint = 255
    var v7 int16 = -32768
    var v8 int32 = -2147483648
    v3 = v3
    v2[v3] = v1
    v9 := GenS1{F0: 127, F1: int(v7)}
    var v10 int = -17
    if true {
        for fi11 := 0; fi11 < 6; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint8 = 42
                if (((!(true && false)) || (18446744073709551615 != uint64(v6))) || false) {
                    var v14 string = (v3 + "\n\t")
                    v14 = ("" + "\n\t")
                    var v15 bool = false
                    _ = v14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        var v16 int32 = int32(v10)
        var v17 int16 = (v7 - (v7 - (-(v7))))
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            switch v3 {
            case "a":
                if true {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if true {
                            var v20 uint8 = 189
                            _ = v20
                        } else {
                            fmt.Println(fi18, v1, v6)
                            fmt.Println(v8, v1)
                            v8 = v8
                        }
                        var v21 int32 = v8
                        v22 := GenS2{F0: 1}
                        _ = fi19
                        _ = v21
                        _ = v22
                    }
                    v2[(("" + "foo bar") + (("foo bar" + ((v3 + "foo bar") + (("foo bar" + "foo bar") + v3))) + v3))] = (v1 % ((2147483648) | 1))
                }
            }
            v23 := GenS3{F0: 95, F1: 0}
            _ = fi18
            _ = v23
        }
        v3 = ("hello" + v3)
        switch 620 {
        case 1:
            fmt.Println(v3)
        case 127:
            var v24 string = (((("foo bar" + v3) + (("\n\t" + "hello") + v3)) + v3) + "\n\t")
            _ = v24
        default:
            var v25 uint8 = 128
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
