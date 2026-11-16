package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int64 = -380
    var v2 float64 = ((((0.0 - 0.1) - float64((1.5 * (-2.25 + 3.14159)))) * 0.0) / 1e10)
    v3 := []uint32{2, 65535}
    v1 = v1
    for fi4 := 0; fi4 < 1; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            fmt.Println(v1, fi4, fi4)
            _ = fi5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        v2 = v2
        switch (("foo bar" + "hello") + ("foo bar" + "foo bar")) {
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            var v7 uint32 = 1
            _ = v7
        }
        _ = fi6
    }
    if ("" > (("hello" + "hello") + "\n\t")) {
        _ = len(v3)
    } else {
        if len(v3) > 0 {
            v3[0] = 65535
        }
    }
    fmt.Println(v2)
    v2 = (3.14159 - (v2 / 1.5))
    switch 2147483647 {
    case -1:
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(fi8, fi8)
            _ = fi8
        }
    }
    var v9 string = ""
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v1, fi10)
        v9 = v9
        if false {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            if false {
                fmt.Println(v2)
                if true {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v12 := GenS1{F0: 159, F1: 74}
                        _ = fi11
                        _ = v12
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 891
                    }
                    fmt.Println(v1, v2)
                    var v13 float64 = v2
                    _ = v13
                }
            } else {
                var v14 uint = 0
                _ = v14
            }
            switch 32768 {
            case 2147483647:
                var v15 int32 = 42
                _ = v15
            case 1:
                var v16 int8 = 0
                _ = v16
            default:
                var v17 int64 = (v1 - v1)
                _ = v17
            }
        } else {
            v18 := GenS1{F0: 1, F1: 100}
            v19 := GenS1{F0: 42, F1: 0}
            var v20 int32 = 0
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi10
    }
    var v21 uint64 = 42
    v9 = v9
    v22 := GenS1{F0: v21, F1: -51}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v23 uint64 = v21
    fmt.Println(v21, v21)
    v3 = append(v3, 65535)
    if false {
        for fi24 := 0; fi24 < 5; fi24++ {
            if true {
                switch v9 {
                case "\n\t":
                    v21 += v23
                case "a":
                    v2 = float64(((1.5 - float32((-2.25 + (v2 * float64(-2.25))))) * 0.1))
                case "foo bar":
                    v1 = (v1 >> ((42) & 63))
                }
                fmt.Println(v23, fi24, v1)
            } else {
                switch (^(fi24)) {
                case 42:
                    v25 := GenS1{F0: (v23 >> ((507) & 63)), F1: 100}
                    _ = v25
                default:
                    if (!false) {
                        if len(v3) > 0 {
                            v3[0] = 280
                        }
                        fmt.Println(v1)
                    }
                }
            }
            _ = fi24
        }
        var v26 uint16 = 255
        var v27 int = 451
        _ = v26
        _ = v27
    } else {
        v28 := GenS2{F0: (v1 - v1), F1: 885}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
