package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int8 = 100
    v2 := []string{"hello", ("hello" + "foo bar"), "foo bar", ("a" + "")}
    v3 := make(map[string]bool)
    var v4 int64 = -9223372036854775808
    var v5 uint = 922
    var v6 int16 = 0
    var v7 uint8 = 128
    v7 -= v7
    for fi8 := 0; fi8 < 10; fi8++ {
        switch "a" {
        case "hello":
            var v9 int = -250
            _ = v9
        case "":
            for fi10 := 0; fi10 < 15; fi10++ {
                v5 = (v5 | 42)
                v5 = v5
                _ = len(v2)
                _ = fi10
            }
        }
        fmt.Println(v7)
        _ = fi8
    }
    var v11 int8 = (v1 & (v1 - int8(v6)))
    v6 += (-(v6))
    var v12 float64 = 0.0
    if len(v2) > 0 {
        v2[0] = ("a" + ("\n\t" + ("\n\t" + "")))
    }
    v7 += 166
    if true {
        for fi13 := 0; fi13 < 20; fi13++ {
            switch (fi13 << (((v5 << (((^(v5))) & 63))) & 63)) {
            case -9223372036854775808:
                v6 += (v6 << ((v5) & 63))
            case 0:
                for fi14 := 0; fi14 < 18; fi14++ {
                    delete(v3, ("\n\t" + "hello"))
                    v5 -= (v5 * 255)
                    _ = fi14
                }
            default:
                for fi15 := 0; fi15 < 4; fi15++ {
                    fmt.Println(v4, fi15)
                    _ = fi15
                }
            }
            var v16 uint16 = 42
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int32 = 1
                v7 += (v7 & (v7 & v7))
                _ = fi17
                _ = v18
            }
            _ = fi13
            _ = v16
        }
        if false {
            fmt.Println(v11)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            switch "" {
            case "\n\t":
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 int = fi19
                    v20 = -371
                    var v21 int = -827
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 12; fi22++ {
                    switch fi22 {
                    case 352:
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v12 = v12
                            var v24 uint = v5
                            _ = fi23
                            _ = v24
                        }
                    case 4294967295:
                        v6 = v6
                    default:
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v26 := GenS1{F0: (384 < 2147483647)}
                            _ = fi25
                            _ = v26
                        }
                    }
                    v1 -= v11
                    _ = fi22
                }
            }
            v4 += v4
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            v6 = -1
            var v29 int8 = v11
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
