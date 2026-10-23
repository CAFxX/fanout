package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello", "\n\t", "a", "hello"}
    v4 := []uint32{42, 673}
    var v5 int = -711
    var v6 string = (("a" + "\n\t") + "foo bar")
    fmt.Println(v5)
    v1 -= (v1 << ((0) & 63))
    for fi7 := 0; fi7 < 19; fi7++ {
        v6 = ("a" + "foo bar")
        _ = fi7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if (true && false) {
            v1 = 0
            v9 := GenS1{F0: (("\n\t" + v6) + "hello")}
            _ = v9
        } else {
            switch ("hello" + v6) {
            case "":
                var v10 int8 = -1
                _ = v10
            case "foo bar":
                var v11 uint = 435
                _ = v11
            case "a":
                if ((fi8 - fi8) < (v5 ^ (^(v5)))) {
                    _ = len(v2)
                }
            }
            delete(v2, 1)
            v1 += (v1 * v1)
        }
        _ = len(v2)
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(v6, v5, v6)
            fmt.Println(fi12, v1)
            _ = fi12
        }
        _ = fi8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 uint64 = 2
    if true {
        fmt.Println(v6, v1)
        var v14 int32 = 1
        var v15 float64 = 1e10
        _ = v14
        _ = v15
    } else {
        v1 += v1
        v1 = 42
        v5 += -1
    }
    if true {
        v13 = (v13 * v13)
        fmt.Println(v5)
    }
    if false {
        if false {
            v16 := GenS2{F0: (255 > 432), F1: 13}
            var v17 int64 = -673
            if true {
                var v18 uint16 = 750
                for fi19 := 0; fi19 < 5; fi19++ {
                    delete(v2, 65535)
                    _ = fi19
                }
                v17 += v17
                _ = v18
            } else {
                v2[42] = 1
            }
            _ = v16
            _ = v17
        }
        var v20 uint64 = v13
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 uint8 = 133
            var v23 int32 = -1
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    } else {
        var v24 uint = 4294967295
        v25 := GenS2{F0: (!true), F1: (v5 * (v5 * v5))}
        _ = v24
        _ = v25
    }
    v5 = v5
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 int8 = 1
            v6 = v6
            v6 = v6
            _ = fi27
            _ = v28
        }
        v29 := GenS2{F0: (!(!true)), F1: 9223372036854775807}
        v30 := GenS1{F0: "\n\t"}
        _ = fi26
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
