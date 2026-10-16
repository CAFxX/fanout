package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 int8, p1 string) int16 {
    if false {
        var v6 float32 = 3.14159
        _ = v6
    } else {
        fmt.Println(p0)
    }
    fmt.Println(p0)
    var v7 int64 = -1
    _ = p0
    _ = p1
    _ = v7
    return -917
}

func main() {
    v1 := []uint32{42, 42, 485}
    var v2 int16 = -1
    var v3 int = 42
    v4 := make(map[int]uint32)
    var v5 int64 = 979
    for fi8 := 0; fi8 < 13; fi8++ {
        if true {
            v3 -= v3
            v3 = (v3 % (((fi8 - v3)) | 1))
        } else {
            var v9 int8 = 100
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v3, v5, fi10)
            v11 := GenS1{F0: v5, F1: 42, F2: 2}
            v2 = 32767
            var v12 uint16 = 65535
            if true {
                v12 = (v12 - (v12 << ((261) & 63)))
                var v13 int = (v3 >> ((18446744073709551615) & 63))
                switch "\n\t" {
                case "":
                    var v14 int16 = (v2 % ((genFn0(100, ("hello" + ""))) | 1))
                    _ = v14
                case "foo bar":
                    var v15 string = "foo bar"
                    _ = v15
                }
                _ = v13
            } else {
                if false {
                    var v16 float32 = (0.0 / 1.5)
                    v17 := GenS1{F0: (v5 ^ (v5 - (v5 << ((578) & 63)))), F1: (^(v12)), F2: v12}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v18 int = fi8
                    v12 = (v12 << ((139) & 63))
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v19 int32 = -2147483648
    v2 = (v2 & (v2 % ((-32768) | 1)))
    for fi20 := 0; fi20 < 2; fi20++ {
        v2 += v2
        fmt.Println(v3, v3)
        _ = fi20
    }
    _ = len(v1)
    if (("" + "foo bar") != ("foo bar" + "hello")) {
        var v21 uint16 = 406
        _ = v21
    }
    var v22 float64 = ((-2.25 / 1e10) * ((0.1 - float64((3.14159 + 1.5))) - (1e10 / 0.1)))
    var v23 uint32 = 4294967295
    v2 += (-(v2))
    switch (v3 << ((426) & 63)) {
    case -1:
        var v24 int32 = (v19 | v19)
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 string = (("hello" + "") + "\n\t")
            v28 := GenS1{F0: -1, F1: 2, F2: 2}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    switch ("\n\t" + (("\n\t" + (("" + "foo bar") + "")) + "")) {
    case "a":
        v23 -= v23
    default:
        v3 -= int(v23)
    }
    delete(v4, (v3 - v3))
    _ = len(v4)
    switch ("" + "a") {
    case "":
        var v29 string = ("foo bar" + (("foo bar" + (("hello" + ("foo bar" + "")) + "a")) + "a"))
        _ = v29
    case "\n\t":
        var v30 uint = 40
        _ = v30
    case "hello":
        v31 := GenS2{F0: 127, F1: int8(v3)}
        _ = v31
    }
    switch v3 {
    case -2147483648:
        if (!(!true)) {
            var v32 uint32 = v23
            _ = v32
        } else {
            if true {
                v33 := GenS1{F0: v5, F1: 1, F2: 854}
                _ = len(v1)
                _ = v33
            } else {
                for fi34 := 0; fi34 < 6; fi34++ {
                    _ = len(v1)
                    _ = fi34
                }
                v35 := GenS1{F0: v5, F1: 1, F2: 637}
                _ = v35
            }
            v5 -= v5
        }
    case 719:
        v3 = int(v19)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
