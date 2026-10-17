package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }

func main() {
    var v1 uint64 = 1
    v2 := []uint32{}
    var v3 uint32 = 42
    fmt.Println(v1, v1)
    var v4 string = "foo bar"
    v3 = (^(v3))
    if false {
        v3 += 676
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 uint = 42
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 255, F1: 42, F2: 446}
        _ = v7
    }
    var v8 uint64 = 42
    switch v4 {
    case "\n\t":
        v4 = "hello"
    case "":
        v4 = "\n\t"
    }
    var v9 int32 = 181
    fmt.Println(v4)
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint64 = v1
        _ = fi10
        _ = v11
    }
    switch v4 {
    case "":
        v2 = append(v2, v3)
    case "a":
        for fi12 := 0; fi12 < 2; fi12++ {
            v3 = v3
            fmt.Println(v3)
            v13 := GenS1{F0: 698, F1: 2, F2: (^(v9))}
            _ = fi12
            _ = v13
        }
    default:
        v14 := GenS1{F0: 42, F1: 2, F2: v9}
        _ = v14
    }
    var v15 int = int(v9)
    for fi16 := 0; fi16 < 17; fi16++ {
        switch ("hello" + (((v4 + "foo bar") + "") + (v4 + v4))) {
        case "foo bar":
            if true {
                var v17 uint64 = 336
                for fi18 := 0; fi18 < 19; fi18++ {
                    fmt.Println(v3, fi18, v1)
                    _ = fi18
                }
                var v19 bool = false
                _ = v17
                _ = v19
            }
        case "a":
            fmt.Println(v4, v15, v3)
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            if len(v2) > 0 {
                v2[0] = v3
            }
            _ = fi20
        }
        _ = fi16
    }
    v8 -= 1
    var v21 string = ("\n\t" + v4)
    if true {
        var v22 int16 = 1
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(v15, v8, v15)
            v24 := GenS1{F0: 469, F1: 65535, F2: v9}
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 int32 = v9
            if (false && false) {
                var v27 uint8 = 2
                _ = v27
            }
            _ = fi25
            _ = v26
        }
    }
    if true {
        var v28 uint8 = 55
        _ = v28
    } else {
        var v29 uint = 775
        for fi30 := 0; fi30 < 20; fi30++ {
            switch "foo bar" {
            case "":
                switch 2147483647 {
                case 32767:
                    fmt.Println(v4)
                }
            default:
                if (2 >= 1) {
                    v21 = v4
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v31 uint8 = 2
                    _ = v31
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v21 = (v4 + v4)
                switch v4 {
                case "hello":
                    v3 = (v3 << (((^(v29))) & 63))
                }
                _ = fi32
            }
            _ = fi30
        }
        _ = v29
    }
    v33 := GenS1{F0: 0, F1: 0, F2: v9}
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 float64 = (1e10 * 1e10)
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
