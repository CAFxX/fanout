package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 string }
func genFn0(p0 int64, p1 bool, p2 uint8) uint {
    switch (("foo bar" + "hello") + "a") {
    case "\n\t":
        v6 := GenS1{F0: 2, F1: 2}
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 88
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{326}
    var v3 int16 = 32767
    var v4 int = 32768
    var v5 int64 = -1
    var v10 float32 = ((1e10 / 1e10) - -2.25)
    v4 = (-(v4))
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 float32 = v10
            switch 42 {
            case 32768:
                switch int(v5) {
                case 0:
                    var v14 int64 = v5
                    _ = v14
                case -1:
                    v15 := GenS1{F0: 2, F1: 924}
                    _ = v15
                }
            default:
                if ((!false) && false) {
                    v16 := GenS1{F0: 42, F1: 1}
                    var v17 int16 = v3
                    v18 := GenS1{F0: 255, F1: 37}
                    v13 = v10
                    v3 = -832
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            _ = len(v1)
            switch (("" + "a") + "\n\t") {
            case "\n\t":
                switch "foo bar" {
                case "foo bar":
                    var v19 string = ((("\n\t" + "foo bar") + ("foo bar" + "hello")) + "\n\t")
                    _ = v19
                case "a":
                    var v20 int64 = v5
                    _ = v20
                }
            case "hello":
                _ = len(v2)
            case "a":
                v21 := GenS2{F0: 42}
                _ = v21
            }
            var v22 uint = 2
            _ = fi12
            _ = v13
            _ = v22
        }
        v3 += (v3 << ((801) & 63))
        _ = fi11
    }
    v10 += 1e10
    var v23 int16 = v3
    v24 := GenS3{F0: 42, F1: 0, F2: ("a" + ("a" + ""))}
    if false {
        if true {
            for fi25 := 0; fi25 < 9; fi25++ {
                switch (fi25 - -177) {
                case -427:
                    fmt.Println(v10, v10, v3)
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 int8 = 127
                    v2 = append(v2, 2)
                    if true {
                        v28 := v1[(fi25 + 2147483647)]
                        var v29 int64 = (^(v5))
                        fmt.Println(v3, v3)
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        }
        _ = len(v1)
        var v30 int8 = 28
        _ = v30
    } else {
        for fi31 := 0; fi31 < 1; fi31++ {
            delete(v1, 127)
            v32 := GenS1{F0: 255, F1: 2}
            if (!false) {
                v3 -= (-(v23))
                v33 := GenS1{F0: 380, F1: 42}
                _ = v33
            }
            _ = fi31
            _ = v32
        }
    }
    v10 = v10
    var v34 uint16 = 1
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
