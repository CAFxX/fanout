package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint64 = 459
    v2 := []uint32{42, 435, 42}
    var v3 uint64 = (^(v1))
    var v4 uint32 = 4294967295
    fmt.Println(v4, v1, v1)
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: ("foo bar" + (("" + "a") + ("hello" + "hello"))), F1: (("hello" + ("\n\t" + "hello")) + ("a" + "foo bar")), F2: fi5}
        _ = fi5
        _ = v6
    }
    if (255 <= 65535) {
        switch "\n\t" {
        case "hello":
            switch 0 {
            case 696:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "":
            v7 := GenS2{F0: "", F1: 2, F2: 396}
            _ = v7
        case "\n\t":
            v3 += (v1 << ((42) & 63))
        default:
            v8 := GenS3{F0: 2147483647, F1: 1000}
            _ = v8
        }
    }
    fmt.Println(v1, v1)
    fmt.Println(v3)
    v1 = (v3 - v1)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 209
        _ = fi9
        _ = v10
    }
    if (true || false) {
        var v11 uint32 = 2147483648
        _ = v11
    } else {
        v12 := GenS4{F0: 42}
        v13 := GenS1{F0: "", F1: "\n\t", F2: 42}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        if (true || false) {
            v1 = 42
            v1 -= v3
        } else {
            var v15 uint16 = 255
            if (true && (!((("a" + ("a" + "foo bar")) + "hello") >= "foo bar"))) {
                v16 := GenS2{F0: "hello", F1: (v4 << ((125) & 63)), F2: 0}
                v3 -= (v1 << ((772) & 63))
                v15 -= v15
                _ = v16
            } else {
                v4 -= v4
            }
            _ = v15
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            switch (("a" + "foo bar") + "\n\t") {
            case "a":
                var v18 uint64 = (v1 * v1)
                _ = v18
            default:
                v19 := GenS1{F0: "foo bar", F1: "a", F2: fi17}
                _ = v19
            }
            v1 = (v1 * v3)
            if false {
                v20 := GenS1{F0: "a", F1: "\n\t", F2: (fi14 - fi17)}
                _ = v20
            }
            _ = fi17
        }
        _ = fi14
    }
    v3 -= (v3 / (((v3 >> ((42) & 63))) | 1))
    fmt.Println(v4)
    for fi21 := 0; fi21 < 16; fi21++ {
        v4 -= (v4 ^ uint32(v1))
        var v22 uint8 = 0
        _ = fi21
        _ = v22
    }
    var v23 int = 4294967295
    switch 32767 {
    case -918:
        if len(v2) > 0 {
            v2[0] = (v4 ^ (^(v4)))
        }
    case 4294967295:
        var v24 uint8 = 42
        _ = v24
    case 0:
        for fi25 := 0; fi25 < 2; fi25++ {
            v4 = v4
            for fi26 := 0; fi26 < 3; fi26++ {
                _ = len(v2)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if len(v2) > 0 {
                    v2[0] = 838
                }
                _ = fi26
            }
            _ = fi25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
